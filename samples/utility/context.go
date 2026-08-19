package utility

import (
	"fmt"
	"runtime"
	"unsafe"

	"go.cld.moe/vk_google/vulkan"
)

// VulkanContext manages instance, device, queues, memory, and command pools.
type VulkanContext struct {
	Instance       vulkan.Instance
	PhysicalDevice vulkan.PhysicalDevice
	Device         vulkan.Device

	GraphicsQueueFamily uint32
	ComputeQueueFamily  uint32

	GraphicsQueue vulkan.Queue
	ComputeQueue  vulkan.Queue

	CommandPool        vulkan.CommandPool
	ComputeCommandPool vulkan.CommandPool

	MemoryProperties vulkan.PhysicalDeviceMemoryProperties
	DeviceProperties vulkan.PhysicalDeviceProperties
	DeviceFeatures   vulkan.PhysicalDeviceFeatures

	SupportsMDI                 bool
	SupportsFirstInstance       bool
	SupportsBufferDeviceAddress bool
}

// ContextConfig specifies options for initializing the Vulkan context.
type ContextConfig struct {
	AppName                     string
	RequireMDI                  bool
	RequireCompute              bool
	EnableBufferDeviceAddress   bool
	EnableValidation            bool
	ValidationCallback          vulkan.PFN_vkDebugUtilsMessengerCallbackEXT
	InstanceExtensions          []string
	DeviceExtensions            []string
}

// NewVulkanContext initializes a VulkanContext based on the provided configuration.
func NewVulkanContext(cfg ContextConfig) (*VulkanContext, error) {
	runtime.LockOSThread()

	if err := vulkan.Init(); err != nil {
		return nil, fmt.Errorf("failed to init vulkan loader: %w", err)
	}
	vulkan.InitCommands(0, 0)

	ctx := &VulkanContext{}

	// 1. Check available instance layers
	var enabledLayers []string
	if cfg.EnableValidation {
		layerProps, _ := vulkan.EnumerateInstanceLayerProperties()
		for _, lp := range layerProps {
			name := vulkan.ByteSliceToString(lp.LayerName[:])
			if name == "VK_LAYER_KHRONOS_validation" {
				enabledLayers = append(enabledLayers, name)
				break
			}
		}
	}

	// 2. Check available instance extensions
	var enabledExtensions []string
	enabledExtensions = append(enabledExtensions, cfg.InstanceExtensions...)
	extProps, res := vulkan.EnumerateInstanceExtensionProperties("")
	if res == vulkan.SUCCESS {
		for _, ep := range extProps {
			name := vulkan.ByteSliceToString(ep.ExtensionName[:])
			if name == "VK_EXT_debug_utils" && cfg.EnableValidation {
				// Don't duplicate if already in cfg.InstanceExtensions
				found := false
				for _, ext := range cfg.InstanceExtensions {
					if ext == name {
						found = true
						break
					}
				}
				if !found {
					enabledExtensions = append(enabledExtensions, name)
				}
			}
		}
	}

	appInfo := vulkan.ApplicationInfo{
		ApplicationName:    cfg.AppName,
		ApplicationVersion: vulkan.MakeVersion(1, 0, 0),
		EngineName:         "vk_google_samples",
		EngineVersion:      vulkan.MakeVersion(1, 0, 0),
		ApiVersion:         vulkan.API_VERSION_1_3,
	}

	instanceCreateInfo := vulkan.InstanceCreateInfo{
		ApplicationInfo:       &appInfo,
		EnabledLayerNames:     enabledLayers,
		EnabledExtensionNames: enabledExtensions,
	}

	var debugMessengerInfo *vulkan.DebugUtilsMessengerCreateInfoEXT
	if cfg.ValidationCallback != nil {
		debugMessengerInfo = vulkan.NewDebugUtilsMessengerCreateInfoEXT()
		debugMessengerInfo.MessageSeverity = vulkan.DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT | vulkan.DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT
		debugMessengerInfo.MessageType = vulkan.DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT | vulkan.DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT
		debugMessengerInfo.UserCallback = cfg.ValidationCallback
		instanceCreateInfo.SetNext(debugMessengerInfo)
	}

	instance, res := vulkan.CreateInstance(&instanceCreateInfo, nil)
	if res != vulkan.SUCCESS {
		return nil, fmt.Errorf("vkCreateInstance failed: %v", res)
	}
	ctx.Instance = instance
	vulkan.InitCommands(instance, 0)

	// 3. Pick Physical Device
	physicalDevices, res := vulkan.EnumeratePhysicalDevices(instance)
	if res != vulkan.SUCCESS || len(physicalDevices) == 0 {
		vulkan.DestroyInstance(instance, nil)
		return nil, fmt.Errorf("no Vulkan physical devices found: %v", res)
	}

	ctx.PhysicalDevice = physicalDevices[0]
	ctx.DeviceProperties = vulkan.GetPhysicalDeviceProperties(ctx.PhysicalDevice)
	ctx.DeviceFeatures = vulkan.GetPhysicalDeviceFeatures(ctx.PhysicalDevice)
	ctx.MemoryProperties = vulkan.GetPhysicalDeviceMemoryProperties(ctx.PhysicalDevice)

	ctx.SupportsMDI = (ctx.DeviceFeatures.MultiDrawIndirect != 0)
	ctx.SupportsFirstInstance = (ctx.DeviceFeatures.DrawIndirectFirstInstance != 0)

	// 4. Find Queue Families
	queueFamilies, _ := vulkan.GetPhysicalDeviceQueueFamilyProperties(ctx.PhysicalDevice)
	ctx.GraphicsQueueFamily = vulkan.QUEUE_FAMILY_IGNORED
	ctx.ComputeQueueFamily = vulkan.QUEUE_FAMILY_IGNORED

	for i, qf := range queueFamilies {
		if (qf.QueueFlags & vulkan.QUEUE_GRAPHICS_BIT) != 0 {
			if ctx.GraphicsQueueFamily == vulkan.QUEUE_FAMILY_IGNORED {
				ctx.GraphicsQueueFamily = uint32(i)
			}
		}
		if (qf.QueueFlags & vulkan.QUEUE_COMPUTE_BIT) != 0 {
			if ctx.ComputeQueueFamily == vulkan.QUEUE_FAMILY_IGNORED {
				ctx.ComputeQueueFamily = uint32(i)
			}
		}
	}

	if ctx.GraphicsQueueFamily == vulkan.QUEUE_FAMILY_IGNORED {
		vulkan.DestroyInstance(instance, nil)
		return nil, fmt.Errorf("no graphics queue family found")
	}
	if ctx.ComputeQueueFamily == vulkan.QUEUE_FAMILY_IGNORED {
		ctx.ComputeQueueFamily = ctx.GraphicsQueueFamily
	}

	// 5. Setup Device Queue Create Infos
	queuePriority := float32(1.0)
	var queueCreateInfos []vulkan.DeviceQueueCreateInfo
	queueFamilyMap := map[uint32]bool{
		ctx.GraphicsQueueFamily: true,
		ctx.ComputeQueueFamily:  true,
	}

	for qf := range queueFamilyMap {
		qci := vulkan.NewDeviceQueueCreateInfo()
		qci.QueueFamilyIndex = qf
		qci.QueuePriorities = []float32{queuePriority}
		queueCreateInfos = append(queueCreateInfos, *qci)
	}

	// 6. Setup Enabled Features
	enabledFeatures := vulkan.PhysicalDeviceFeatures{}
	if ctx.SupportsMDI {
		enabledFeatures.MultiDrawIndirect = vulkan.True
	}
	if ctx.SupportsFirstInstance {
		enabledFeatures.DrawIndirectFirstInstance = vulkan.True
	}
	enabledFeatures.SamplerAnisotropy = ctx.DeviceFeatures.SamplerAnisotropy

	// Enable Vulkan 1.3 dynamic rendering feature
	vk13Features := vulkan.NewPhysicalDeviceVulkan13Features()
	vk13Features.DynamicRendering = vulkan.True
	vk13Features.Synchronization2 = vulkan.True

	deviceCreateInfo := vulkan.NewDeviceCreateInfo()
	deviceCreateInfo.QueueCreateInfos = queueCreateInfos
	deviceCreateInfo.EnabledFeatures = &enabledFeatures
	deviceCreateInfo.EnabledExtensionNames = cfg.DeviceExtensions
	deviceCreateInfo.SetNext(vk13Features)

	device, res := vulkan.CreateDevice(ctx.PhysicalDevice, deviceCreateInfo, nil)
	if res != vulkan.SUCCESS {
		vulkan.DestroyInstance(instance, nil)
		return nil, fmt.Errorf("vkCreateDevice failed: %v", res)
	}
	ctx.Device = device
	vulkan.InitCommands(instance, device)

	// 7. Get Queues
	ctx.GraphicsQueue = vulkan.GetDeviceQueue(device, ctx.GraphicsQueueFamily, 0)
	ctx.ComputeQueue = vulkan.GetDeviceQueue(device, ctx.ComputeQueueFamily, 0)

	// 8. Create Command Pools
	poolInfo := vulkan.NewCommandPoolCreateInfo()
	poolInfo.Flags = vulkan.COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT
	poolInfo.QueueFamilyIndex = ctx.GraphicsQueueFamily

	cmdPool, res := vulkan.CreateCommandPool(device, poolInfo, nil)
	if res != vulkan.SUCCESS {
		ctx.Destroy()
		return nil, fmt.Errorf("vkCreateCommandPool graphics failed: %v", res)
	}
	ctx.CommandPool = cmdPool

	if ctx.ComputeQueueFamily != ctx.GraphicsQueueFamily {
		computePoolInfo := vulkan.NewCommandPoolCreateInfo()
		computePoolInfo.Flags = vulkan.COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT
		computePoolInfo.QueueFamilyIndex = ctx.ComputeQueueFamily
		computePool, res := vulkan.CreateCommandPool(device, computePoolInfo, nil)
		if res == vulkan.SUCCESS {
			ctx.ComputeCommandPool = computePool
		} else {
			ctx.ComputeCommandPool = ctx.CommandPool
		}
	} else {
		ctx.ComputeCommandPool = ctx.CommandPool
	}

	return ctx, nil
}

// Destroy frees all Vulkan resources associated with the context.
func (ctx *VulkanContext) Destroy() {
	if ctx.Device != 0 {
		vulkan.DeviceWaitIdle(ctx.Device)
		if ctx.ComputeCommandPool != 0 && ctx.ComputeCommandPool != ctx.CommandPool {
			vulkan.DestroyCommandPool(ctx.Device, ctx.ComputeCommandPool, nil)
			ctx.ComputeCommandPool = 0
		}
		if ctx.CommandPool != 0 {
			vulkan.DestroyCommandPool(ctx.Device, ctx.CommandPool, nil)
			ctx.CommandPool = 0
		}
		vulkan.DestroyDevice(ctx.Device, nil)
		ctx.Device = 0
	}
	if ctx.Instance != 0 {
		vulkan.DestroyInstance(ctx.Instance, nil)
		ctx.Instance = 0
	}
	runtime.UnlockOSThread()
}

// FindMemoryType finds a suitable memory type index for given typeFilter and memory property flags.
func (ctx *VulkanContext) FindMemoryType(typeFilter uint32, properties vulkan.MemoryPropertyFlags) (uint32, error) {
	for i := uint32(0); i < ctx.MemoryProperties.MemoryTypeCount; i++ {
		if (typeFilter & (1 << i)) != 0 {
			if (ctx.MemoryProperties.MemoryTypes[i].PropertyFlags & properties) == properties {
				return i, nil
			}
		}
	}
	return 0, fmt.Errorf("failed to find suitable memory type for filter 0x%x, properties 0x%x", typeFilter, properties)
}

// ExecuteOneTimeCommands allocates a command buffer, runs the provided callback, submits, and waits.
func (ctx *VulkanContext) ExecuteOneTimeCommands(fn func(cmd vulkan.CommandBuffer)) error {
	allocInfo := vulkan.NewCommandBufferAllocateInfo()
	allocInfo.CommandPool = ctx.CommandPool
	allocInfo.Level = vulkan.COMMAND_BUFFER_LEVEL_PRIMARY
	allocInfo.CommandBufferCount = 1

	cmd, res := vulkan.AllocateCommandBuffers(ctx.Device, allocInfo)
	if res != vulkan.SUCCESS || cmd == 0 {
		return fmt.Errorf("failed to allocate one-time command buffer: %v", res)
	}

	beginInfo := vulkan.NewCommandBufferBeginInfo()
	beginInfo.Flags = vulkan.COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT
	vulkan.BeginCommandBuffer(cmd, beginInfo)

	fn(cmd)

	vulkan.EndCommandBuffer(cmd)

	submitInfo := vulkan.NewSubmitInfo()
	submitInfo.CommandBuffers = []vulkan.CommandBuffer{cmd}

	vulkan.QueueSubmit(ctx.GraphicsQueue, []vulkan.SubmitInfo{*submitInfo}, 0)
	vulkan.QueueWaitIdle(ctx.GraphicsQueue)

	vulkan.FreeCommandBuffers(ctx.Device, ctx.CommandPool, []vulkan.CommandBuffer{cmd})
	return nil
}

// AsPointer converts a pointer to unsafe.Pointer.
func AsPointer[T any](ptr *T) unsafe.Pointer {
	return unsafe.Pointer(ptr)
}
