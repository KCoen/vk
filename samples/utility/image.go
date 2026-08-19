package utility

import (
	"fmt"

	"go.cld.moe/vk_google/vulkan"
)

// Texture wraps an Image, DeviceMemory, ImageView, and Sampler.
type Texture struct {
	Image   vulkan.Image
	Memory  vulkan.DeviceMemory
	View    vulkan.ImageView
	Sampler vulkan.Sampler
	Width   uint32
	Height  uint32
	Format  vulkan.Format
}

// CreateTexture2D creates a 2D texture from raw RGBA pixel data.
func CreateTexture2D(ctx *VulkanContext, width, height uint32, format vulkan.Format, pixels []byte) (*Texture, error) {
	imageSize := uint64(len(pixels))
	if imageSize == 0 {
		imageSize = uint64(width * height * 4)
	}

	// 1. Create Staging Buffer
	staging, err := CreateBuffer(ctx, imageSize, vulkan.BUFFER_USAGE_TRANSFER_SRC_BIT, vulkan.MEMORY_PROPERTY_HOST_VISIBLE_BIT|vulkan.MEMORY_PROPERTY_HOST_COHERENT_BIT)
	if err != nil {
		return nil, fmt.Errorf("failed to create staging buffer for texture: %w", err)
	}
	defer staging.Destroy(ctx)

	if len(pixels) > 0 {
		if err := staging.Upload(ctx, pixels, 0); err != nil {
			return nil, fmt.Errorf("failed to upload texture pixels: %w", err)
		}
	}

	// 2. Create VkImage
	imgInfo := vulkan.NewImageCreateInfo()
	imgInfo.ImageType = vulkan.IMAGE_TYPE_2D
	imgInfo.Extent = vulkan.Extent3D{Width: width, Height: height, Depth: 1}
	imgInfo.MipLevels = 1
	imgInfo.ArrayLayers = 1
	imgInfo.Format = format
	imgInfo.Tiling = vulkan.IMAGE_TILING_OPTIMAL
	imgInfo.InitialLayout = vulkan.IMAGE_LAYOUT_UNDEFINED
	imgInfo.Usage = vulkan.IMAGE_USAGE_TRANSFER_DST_BIT | vulkan.IMAGE_USAGE_SAMPLED_BIT
	imgInfo.SharingMode = vulkan.SHARING_MODE_EXCLUSIVE
	imgInfo.Samples = vulkan.SAMPLE_COUNT_1_BIT

	image, res := vulkan.CreateImage(ctx.Device, imgInfo, nil)
	if res != vulkan.SUCCESS {
		return nil, fmt.Errorf("vkCreateImage failed: %v", res)
	}

	memReqs := vulkan.GetImageMemoryRequirements(ctx.Device, image)
	memTypeIndex, err := ctx.FindMemoryType(memReqs.MemoryTypeBits, vulkan.MEMORY_PROPERTY_DEVICE_LOCAL_BIT)
	if err != nil {
		vulkan.DestroyImage(ctx.Device, image, nil)
		return nil, fmt.Errorf("FindMemoryType for image failed: %w", err)
	}

	allocInfo := vulkan.NewMemoryAllocateInfo()
	allocInfo.AllocationSize = memReqs.Size
	allocInfo.MemoryTypeIndex = memTypeIndex

	memory, res := vulkan.AllocateMemory(ctx.Device, allocInfo, nil)
	if res != vulkan.SUCCESS {
		vulkan.DestroyImage(ctx.Device, image, nil)
		return nil, fmt.Errorf("vkAllocateMemory for image failed: %v", res)
	}

	vulkan.BindImageMemory(ctx.Device, image, memory, 0)

	// 3. Transition image to TRANSFER_DST_OPTIMAL, copy buffer, then transition to SHADER_READ_ONLY_OPTIMAL
	err = ctx.ExecuteOneTimeCommands(func(cmd vulkan.CommandBuffer) {
		TransitionImageLayout(cmd, image, vulkan.IMAGE_LAYOUT_UNDEFINED, vulkan.IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL)

		region := vulkan.BufferImageCopy{
			BufferOffset:      0,
			BufferRowLength:   0,
			BufferImageHeight: 0,
			ImageSubresource: vulkan.ImageSubresourceLayers{
				AspectMask:     vulkan.IMAGE_ASPECT_COLOR_BIT,
				MipLevel:       0,
				BaseArrayLayer: 0,
				LayerCount:     1,
			},
			ImageOffset: vulkan.Offset3D{X: 0, Y: 0, Z: 0},
			ImageExtent: vulkan.Extent3D{Width: width, Height: height, Depth: 1},
		}

		vulkan.CmdCopyBufferToImage(cmd, staging.Handle, image, vulkan.IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL, []vulkan.BufferImageCopy{region})

		TransitionImageLayout(cmd, image, vulkan.IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL, vulkan.IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL)
	})
	if err != nil {
		vulkan.FreeMemory(ctx.Device, memory, nil)
		vulkan.DestroyImage(ctx.Device, image, nil)
		return nil, fmt.Errorf("texture upload failed: %w", err)
	}

	// 4. Create ImageView
	viewInfo := vulkan.NewImageViewCreateInfo()
	viewInfo.Image = image
	viewInfo.ViewType = vulkan.IMAGE_VIEW_TYPE_2D
	viewInfo.Format = format
	viewInfo.SubresourceRange = vulkan.ImageSubresourceRange{
		AspectMask:     vulkan.IMAGE_ASPECT_COLOR_BIT,
		BaseMipLevel:   0,
		LevelCount:     1,
		BaseArrayLayer: 0,
		LayerCount:     1,
	}

	view, res := vulkan.CreateImageView(ctx.Device, viewInfo, nil)
	if res != vulkan.SUCCESS {
		vulkan.FreeMemory(ctx.Device, memory, nil)
		vulkan.DestroyImage(ctx.Device, image, nil)
		return nil, fmt.Errorf("vkCreateImageView failed: %v", res)
	}

	// 5. Create Sampler
	samplerInfo := vulkan.NewSamplerCreateInfo()
	samplerInfo.MagFilter = vulkan.FILTER_LINEAR
	samplerInfo.MinFilter = vulkan.FILTER_LINEAR
	samplerInfo.AddressModeU = vulkan.SAMPLER_ADDRESS_MODE_REPEAT
	samplerInfo.AddressModeV = vulkan.SAMPLER_ADDRESS_MODE_REPEAT
	samplerInfo.AddressModeW = vulkan.SAMPLER_ADDRESS_MODE_REPEAT
	samplerInfo.MipmapMode = vulkan.SAMPLER_MIPMAP_MODE_LINEAR

	sampler, res := vulkan.CreateSampler(ctx.Device, samplerInfo, nil)
	if res != vulkan.SUCCESS {
		vulkan.DestroyImageView(ctx.Device, view, nil)
		vulkan.FreeMemory(ctx.Device, memory, nil)
		vulkan.DestroyImage(ctx.Device, image, nil)
		return nil, fmt.Errorf("vkCreateSampler failed: %v", res)
	}

	return &Texture{
		Image:   image,
		Memory:  memory,
		View:    view,
		Sampler: sampler,
		Width:   width,
		Height:  height,
		Format:  format,
	}, nil
}

// TransitionImageLayout records an image memory barrier for layout transitions.
func TransitionImageLayout(cmd vulkan.CommandBuffer, image vulkan.Image, oldLayout, newLayout vulkan.ImageLayout) {
	barrier := vulkan.NewImageMemoryBarrier()
	barrier.OldLayout = oldLayout
	barrier.NewLayout = newLayout
	barrier.SrcQueueFamilyIndex = vulkan.QUEUE_FAMILY_IGNORED
	barrier.DstQueueFamilyIndex = vulkan.QUEUE_FAMILY_IGNORED
	barrier.Image = image
	barrier.SubresourceRange = vulkan.ImageSubresourceRange{
		AspectMask:     vulkan.IMAGE_ASPECT_COLOR_BIT,
		BaseMipLevel:   0,
		LevelCount:     1,
		BaseArrayLayer: 0,
		LayerCount:     1,
	}

	var srcStage, dstStage vulkan.PipelineStageFlags

	if oldLayout == vulkan.IMAGE_LAYOUT_UNDEFINED && newLayout == vulkan.IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL {
		barrier.SrcAccessMask = 0
		barrier.DstAccessMask = vulkan.ACCESS_TRANSFER_WRITE_BIT
		srcStage = vulkan.PIPELINE_STAGE_TOP_OF_PIPE_BIT
		dstStage = vulkan.PIPELINE_STAGE_TRANSFER_BIT
	} else if oldLayout == vulkan.IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL && newLayout == vulkan.IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL {
		barrier.SrcAccessMask = vulkan.ACCESS_TRANSFER_WRITE_BIT
		barrier.DstAccessMask = vulkan.ACCESS_SHADER_READ_BIT
		srcStage = vulkan.PIPELINE_STAGE_TRANSFER_BIT
		dstStage = vulkan.PIPELINE_STAGE_FRAGMENT_SHADER_BIT
	} else {
		barrier.SrcAccessMask = 0
		barrier.DstAccessMask = vulkan.ACCESS_MEMORY_READ_BIT | vulkan.ACCESS_MEMORY_WRITE_BIT
		srcStage = vulkan.PIPELINE_STAGE_ALL_COMMANDS_BIT
		dstStage = vulkan.PIPELINE_STAGE_ALL_COMMANDS_BIT
	}

	vulkan.CmdPipelineBarrier(
		cmd,
		srcStage,
		dstStage,
		0,
		nil,
		nil,
		[]vulkan.ImageMemoryBarrier{*barrier},
	)
}

// Destroy frees the image, view, sampler, and backing memory.
func (t *Texture) Destroy(ctx *VulkanContext) {
	if t == nil {
		return
	}
	if t.Sampler != 0 {
		vulkan.DestroySampler(ctx.Device, t.Sampler, nil)
		t.Sampler = 0
	}
	if t.View != 0 {
		vulkan.DestroyImageView(ctx.Device, t.View, nil)
		t.View = 0
	}
	if t.Image != 0 {
		vulkan.DestroyImage(ctx.Device, t.Image, nil)
		t.Image = 0
	}
	if t.Memory != 0 {
		vulkan.FreeMemory(ctx.Device, t.Memory, nil)
		t.Memory = 0
	}
}
