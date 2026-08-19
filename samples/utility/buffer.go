package utility

import (
	"fmt"
	"unsafe"

	"github.com/KCoen/vk/vulkan"
)

// Buffer wraps a Vulkan Buffer and its backing DeviceMemory.
type Buffer struct {
	Handle vulkan.Buffer
	Memory vulkan.DeviceMemory
	Size   uint64
	Mapped unsafe.Pointer
}

// CreateBuffer creates a VkBuffer and allocates/binds VkDeviceMemory for it.
func CreateBuffer(ctx *VulkanContext, size uint64, usage vulkan.BufferUsageFlags, properties vulkan.MemoryPropertyFlags) (*Buffer, error) {
	if size == 0 {
		return nil, fmt.Errorf("cannot create 0-sized buffer")
	}

	bufInfo := vulkan.NewBufferCreateInfo()
	bufInfo.Size = vulkan.DeviceSize(size)
	bufInfo.Usage = usage
	bufInfo.SharingMode = vulkan.SHARING_MODE_EXCLUSIVE

	handle, res := vulkan.CreateBuffer(ctx.Device, bufInfo, nil)
	if res != vulkan.SUCCESS {
		return nil, fmt.Errorf("vkCreateBuffer failed: %v", res)
	}

	memReqs := vulkan.GetBufferMemoryRequirements(ctx.Device, handle)
	memTypeIndex, err := ctx.FindMemoryType(memReqs.MemoryTypeBits, properties)
	if err != nil {
		vulkan.DestroyBuffer(ctx.Device, handle, nil)
		return nil, fmt.Errorf("FindMemoryType failed: %w", err)
	}

	allocInfo := vulkan.NewMemoryAllocateInfo()
	allocInfo.AllocationSize = memReqs.Size
	allocInfo.MemoryTypeIndex = memTypeIndex

	memory, res := vulkan.AllocateMemory(ctx.Device, allocInfo, nil)
	if res != vulkan.SUCCESS {
		vulkan.DestroyBuffer(ctx.Device, handle, nil)
		return nil, fmt.Errorf("vkAllocateMemory failed: %v", res)
	}

	res = vulkan.BindBufferMemory(ctx.Device, handle, memory, 0)
	if res != vulkan.SUCCESS {
		vulkan.FreeMemory(ctx.Device, memory, nil)
		vulkan.DestroyBuffer(ctx.Device, handle, nil)
		return nil, fmt.Errorf("vkBindBufferMemory failed: %v", res)
	}

	b := &Buffer{
		Handle: handle,
		Memory: memory,
		Size:   size,
	}

	// Auto-map if host-visible
	if (properties & vulkan.MEMORY_PROPERTY_HOST_VISIBLE_BIT) != 0 {
		b.Map(ctx)
	}

	return b, nil
}

// Map maps the device memory for host access.
func (b *Buffer) Map(ctx *VulkanContext) (unsafe.Pointer, error) {
	if b.Mapped != nil {
		return b.Mapped, nil
	}
	ptr, res := vulkan.MapMemory(ctx.Device, b.Memory, 0, vulkan.DeviceSize(b.Size), 0)
	if res != vulkan.SUCCESS {
		return nil, fmt.Errorf("vkMapMemory failed: %v", res)
	}
	b.Mapped = ptr
	return ptr, nil
}

// Unmap unmaps the host memory.
func (b *Buffer) Unmap(ctx *VulkanContext) {
	if b.Mapped != nil {
		vulkan.UnmapMemory(ctx.Device, b.Memory)
		b.Mapped = nil
	}
}

// Upload writes raw data bytes into the buffer memory (requires HOST_VISIBLE memory).
func (b *Buffer) Upload(ctx *VulkanContext, data []byte, offset uint64) error {
	if offset+uint64(len(data)) > b.Size {
		return fmt.Errorf("upload out of range: offset %d + data %d > size %d", offset, len(data), b.Size)
	}
	ptr, err := b.Map(ctx)
	if err != nil {
		return err
	}
	dst := unsafe.Add(ptr, offset)
	copy(unsafe.Slice((*byte)(dst), len(data)), data)
	return nil
}

// UploadSlice writes a slice of any type T into the buffer.
func UploadSlice[T any](ctx *VulkanContext, b *Buffer, data []T, offset uint64) error {
	if len(data) == 0 {
		return nil
	}
	var dummy T
	elemSize := uint64(unsafe.Sizeof(dummy))
	totalBytes := elemSize * uint64(len(data))
	srcBytes := unsafe.Slice((*byte)(unsafe.Pointer(&data[0])), totalBytes)
	return b.Upload(ctx, srcBytes, offset)
}

// ReadSlice reads back a slice of type T from a host-visible buffer.
func ReadSlice[T any](ctx *VulkanContext, b *Buffer, count int, offset uint64) ([]T, error) {
	if count == 0 {
		return nil, nil
	}
	ptr, err := b.Map(ctx)
	if err != nil {
		return nil, err
	}
	var dummy T
	elemSize := uint64(unsafe.Sizeof(dummy))
	totalBytes := elemSize * uint64(count)
	if offset+totalBytes > b.Size {
		return nil, fmt.Errorf("read out of range: offset %d + bytes %d > size %d", offset, totalBytes, b.Size)
	}

	src := unsafe.Add(ptr, offset)
	out := make([]T, count)
	dstBytes := unsafe.Slice((*byte)(unsafe.Pointer(&out[0])), totalBytes)
	srcBytes := unsafe.Slice((*byte)(src), totalBytes)
	copy(dstBytes, srcBytes)
	return out, nil
}

// CopyTo copies content from this buffer to dst buffer using a command buffer.
func (b *Buffer) CopyTo(ctx *VulkanContext, dst *Buffer, size uint64) error {
	return ctx.ExecuteOneTimeCommands(func(cmd vulkan.CommandBuffer) {
		copyRegion := vulkan.BufferCopy{
			SrcOffset: 0,
			DstOffset: 0,
			Size:      vulkan.DeviceSize(size),
		}
		vulkan.CmdCopyBuffer(cmd, b.Handle, dst.Handle, []vulkan.BufferCopy{copyRegion})
	})
}

// CreateDeviceLocalBufferFromData creates a GPU-only buffer and populates it via a temporary staging buffer.
func CreateDeviceLocalBufferFromData(ctx *VulkanContext, data []byte, usage vulkan.BufferUsageFlags) (*Buffer, error) {
	size := uint64(len(data))
	staging, err := CreateBuffer(ctx, size, vulkan.BUFFER_USAGE_TRANSFER_SRC_BIT, vulkan.MEMORY_PROPERTY_HOST_VISIBLE_BIT|vulkan.MEMORY_PROPERTY_HOST_COHERENT_BIT)
	if err != nil {
		return nil, fmt.Errorf("failed to create staging buffer: %w", err)
	}
	defer staging.Destroy(ctx)

	if err := staging.Upload(ctx, data, 0); err != nil {
		return nil, fmt.Errorf("failed to upload staging data: %w", err)
	}

	gpuBuf, err := CreateBuffer(ctx, size, usage|vulkan.BUFFER_USAGE_TRANSFER_DST_BIT, vulkan.MEMORY_PROPERTY_DEVICE_LOCAL_BIT)
	if err != nil {
		return nil, fmt.Errorf("failed to create GPU buffer: %w", err)
	}

	if err := staging.CopyTo(ctx, gpuBuf, size); err != nil {
		gpuBuf.Destroy(ctx)
		return nil, fmt.Errorf("failed to copy staging buffer to GPU buffer: %w", err)
	}

	return gpuBuf, nil
}

// Destroy destroys the buffer handle and frees its device memory.
func (b *Buffer) Destroy(ctx *VulkanContext) {
	if b == nil {
		return
	}
	b.Unmap(ctx)
	if b.Handle != 0 {
		vulkan.DestroyBuffer(ctx.Device, b.Handle, nil)
		b.Handle = 0
	}
	if b.Memory != 0 {
		vulkan.FreeMemory(ctx.Device, b.Memory, nil)
		b.Memory = 0
	}
}
