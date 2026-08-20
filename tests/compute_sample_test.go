package tests

import (
	"runtime"
	"testing"
	"unsafe"

	"github.com/KCoen/vk/vulkan"
)

func bailOnBadResult(res vulkan.Result) {
	if res != vulkan.SUCCESS {
		panic(res)
	}
}

func getBestTransferQueue(physicalDevice vulkan.PhysicalDevice) (uint32, bool) {
	props, res := vulkan.GetPhysicalDeviceQueueFamilyProperties(physicalDevice)
	if res != vulkan.SUCCESS {
		return 0, false
	}

	for i, p := range props {
		if p.QueueFlags.Has(vulkan.QUEUE_GRAPHICS_BIT) || p.QueueFlags.Has(vulkan.QUEUE_COMPUTE_BIT) || p.QueueFlags.Has(vulkan.QUEUE_TRANSFER_BIT) {
			return uint32(i), true
		}
	}
	return 0, false
}

func getBestComputeQueue(physicalDevice vulkan.PhysicalDevice) (uint32, bool) {
	props, res := vulkan.GetPhysicalDeviceQueueFamilyProperties(physicalDevice)
	if res != vulkan.SUCCESS {
		return 0, false
	}

	for i, p := range props {
		if p.QueueFlags.Has(vulkan.QUEUE_COMPUTE_BIT) {
			return uint32(i), true
		}
	}
	return 0, false
}

func TestHelloWorld(t *testing.T) {
	if err := vulkan.Init(); err != nil {
		t.Skipf("Vulkan loader not available on host: %v", err)
	}

	appInfo := vulkan.ApplicationInfo{
		ApplicationName:    "Hello Triangle",
		ApplicationVersion: vulkan.MakeVersion(1, 0, 0),
		EngineName:         "go/none",
		EngineVersion:      vulkan.MakeVersion(1, 0, 0),
		ApiVersion:         vulkan.MakeVersion(1, 0, 0),
	}

	createInfo := vulkan.InstanceCreateInfo{
		ApplicationInfo: &appInfo,
	}

	exts, res := vulkan.EnumerateInstanceExtensionProperties("")
	if res != vulkan.SUCCESS {
		t.Fatalf("Failed to enumerate extension properties: %v", res)
	}
	t.Logf("Num extensions: %d", len(exts))
	for i, ext := range exts {
		extName := vulkan.ByteSliceToString(ext.ExtensionName[:])
		t.Logf("Ext[%d]\tname: %s (spec %d)", i, extName, ext.SpecVersion)
	}

	instance, res := vulkan.CreateInstance(&createInfo, nil)
	if res != vulkan.SUCCESS {
		t.Fatalf("Failed to create instance: %v", res)
	}
	t.Logf("Created instance: %v", instance)

	vulkan.InitCommands(instance, 0)
	vulkan.DestroyInstance(instance, nil)
	t.Logf("Destroyed instance successfully")
}

func TestCompute(t *testing.T) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if err := vulkan.Init(); err != nil {
		t.Skipf("Vulkan loader not available on host: %v", err)
	}

	applicationInfo := vulkan.ApplicationInfo{
		ApplicationName:    "VKComputeSample",
		ApplicationVersion: 0,
		EngineName:         "",
		EngineVersion:      0,
		ApiVersion:         vulkan.MakeVersion(1, 0, 9),
	}

	instanceCreateInfo := vulkan.InstanceCreateInfo{
		ApplicationInfo: &applicationInfo,
	}

	instance, res := vulkan.CreateInstance(&instanceCreateInfo, nil)
	bailOnBadResult(res)
	defer vulkan.DestroyInstance(instance, nil)

	vulkan.InitCommands(instance, 0)

	physicalDevices, res := vulkan.EnumeratePhysicalDevices(instance)
	if res == vulkan.ERROR_INITIALIZATION_FAILED || len(physicalDevices) == 0 {
		t.Skip("No valid GPUs or physical devices detected on this host")
	}
	bailOnBadResult(res)

	for _, physicalDevice := range physicalDevices {
		queueFamilyIndex, ok := getBestComputeQueue(physicalDevice)
		if !ok {
			t.Log("Physical device does not support compute queue")
			continue
		}

		queuePriority := float32(1.0)
		deviceQueueCreateInfo := vulkan.DeviceQueueCreateInfo{
			QueueFamilyIndex: queueFamilyIndex,
			QueuePriorities:  []float32{queuePriority},
		}

		deviceCreateInfo := vulkan.DeviceCreateInfo{
			QueueCreateInfos: []vulkan.DeviceQueueCreateInfo{deviceQueueCreateInfo},
		}

		device, res := vulkan.CreateDevice(physicalDevice, &deviceCreateInfo, nil)
		bailOnBadResult(res)
		defer vulkan.DestroyDevice(device, nil)

		vulkan.InitCommands(instance, device)

		properties := vulkan.GetPhysicalDeviceMemoryProperties(physicalDevice)

		bufferLength := int32(16384)
		bufferSize := vulkan.DeviceSize(4 * bufferLength)
		memorySize := bufferSize * 2

		memoryTypeIndex := uint32(vulkan.MAX_MEMORY_TYPES)
		for k := uint32(0); k < properties.MemoryTypeCount; k++ {
			memType := properties.MemoryTypes[k]
			if memType.PropertyFlags.Has(vulkan.MEMORY_PROPERTY_HOST_VISIBLE_BIT) &&
				memType.PropertyFlags.Has(vulkan.MEMORY_PROPERTY_HOST_COHERENT_BIT) &&
				memorySize < properties.MemoryHeaps[memType.HeapIndex].Size {
				memoryTypeIndex = k
				break
			}
		}

		if memoryTypeIndex == vulkan.MAX_MEMORY_TYPES {
			panic(vulkan.ERROR_OUT_OF_HOST_MEMORY)
		}

		memoryAllocateInfo := vulkan.MemoryAllocateInfo{
			AllocationSize:  memorySize,
			MemoryTypeIndex: memoryTypeIndex,
		}

		memory, res := vulkan.AllocateMemory(device, &memoryAllocateInfo, nil)
		bailOnBadResult(res)
		defer vulkan.FreeMemory(device, memory, nil)

		payloadPtr, res := vulkan.MapMemory(device, memory, 0, memorySize, 0)
		bailOnBadResult(res)
		if payloadPtr == nil {
			panic("mapped pointer is nil")
		}

		payload := unsafe.Slice((*int32)(payloadPtr), memorySize/4)
		for k := uint64(1); k < uint64(memorySize/4); k++ {
			payload[k] = 1
		}
		vulkan.UnmapMemory(device, memory)

		bufferCreateInfo := vulkan.BufferCreateInfo{
			Size:               bufferSize,
			Usage:              vulkan.BUFFER_USAGE_STORAGE_BUFFER_BIT,
			SharingMode:        vulkan.SHARING_MODE_EXCLUSIVE,
			QueueFamilyIndices: []uint32{queueFamilyIndex},
		}

		inBuffer, res := vulkan.CreateBuffer(device, &bufferCreateInfo, nil)
		bailOnBadResult(res)
		defer vulkan.DestroyBuffer(device, inBuffer, nil)
		bailOnBadResult(vulkan.BindBufferMemory(device, inBuffer, memory, 0))

		outBuffer, res := vulkan.CreateBuffer(device, &bufferCreateInfo, nil)
		bailOnBadResult(res)
		defer vulkan.DestroyBuffer(device, outBuffer, nil)
		bailOnBadResult(vulkan.BindBufferMemory(device, outBuffer, memory, bufferSize))

		const (
			RESERVED_ID = iota
			FUNC_ID
			IN_ID
			OUT_ID
			GLOBAL_INVOCATION_ID
			VOID_TYPE_ID
			FUNC_TYPE_ID
			INT_TYPE_ID
			INT_ARRAY_TYPE_ID
			STRUCT_ID
			POINTER_TYPE_ID
			ELEMENT_POINTER_TYPE_ID
			INT_VECTOR_TYPE_ID
			INT_VECTOR_POINTER_TYPE_ID
			INT_POINTER_TYPE_ID
			CONSTANT_ZERO_ID
			CONSTANT_ARRAY_LENGTH_ID
			LABEL_ID
			IN_ELEMENT_ID
			OUT_ELEMENT_ID
			GLOBAL_INVOCATION_X_ID
			GLOBAL_INVOCATION_X_PTR_ID
			TEMP_LOADED_ID
			BOUND
		)

		const (
			INPUT                = 1
			UNIFORM              = 2
			BUFFER_BLOCK         = 3
			ARRAY_STRIDE         = 6
			BUILTIN              = 11
			BINDING              = 33
			OFFSET               = 35
			DESCRIPTOR_SET       = 34
			GLOBAL_INVOCATION    = 28
			OP_TYPE_VOID         = 19
			OP_TYPE_FUNCTION     = 33
			OP_TYPE_INT          = 21
			OP_TYPE_VECTOR       = 23
			OP_TYPE_ARRAY        = 28
			OP_TYPE_STRUCT       = 30
			OP_TYPE_POINTER      = 32
			OP_VARIABLE          = 59
			OP_DECORATE          = 71
			OP_MEMBER_DECORATE   = 72
			OP_FUNCTION          = 54
			OP_LABEL             = 248
			OP_ACCESS_CHAIN      = 65
			OP_CONSTANT          = 43
			OP_LOAD              = 61
			OP_STORE             = 62
			OP_RETURN            = 253
			OP_FUNCTION_END      = 56
			OP_CAPABILITY        = 17
			OP_MEMORY_MODEL      = 14
			OP_ENTRY_POINT       = 15
			OP_EXECUTION_MODE    = 16
			OP_COMPOSITE_EXTRACT = 81
		)

		shader := []uint32{
			0x07230203,
			0x00010000,
			0,
			BOUND,
			0,

			(2 << 16) | OP_CAPABILITY, 1,
			(3 << 16) | OP_MEMORY_MODEL, 0, 0,
			(4 << 16) | OP_ENTRY_POINT, 5, FUNC_ID, 0x00000066,
			(6 << 16) | OP_EXECUTION_MODE, FUNC_ID, 17, 1, 1, 1,
			(3 << 16) | OP_DECORATE, STRUCT_ID, BUFFER_BLOCK,
			(4 << 16) | OP_DECORATE, GLOBAL_INVOCATION_ID, BUILTIN, GLOBAL_INVOCATION,
			(4 << 16) | OP_DECORATE, IN_ID, DESCRIPTOR_SET, 0,
			(4 << 16) | OP_DECORATE, IN_ID, BINDING, 0,
			(4 << 16) | OP_DECORATE, OUT_ID, DESCRIPTOR_SET, 0,
			(4 << 16) | OP_DECORATE, OUT_ID, BINDING, 1,
			(4 << 16) | OP_DECORATE, INT_ARRAY_TYPE_ID, ARRAY_STRIDE, 4,
			(5 << 16) | OP_MEMBER_DECORATE, STRUCT_ID, 0, OFFSET, 0,
			(2 << 16) | OP_TYPE_VOID, VOID_TYPE_ID,
			(3 << 16) | OP_TYPE_FUNCTION, FUNC_TYPE_ID, VOID_TYPE_ID,
			(4 << 16) | OP_TYPE_INT, INT_TYPE_ID, 32, 1,
			(4 << 16) | OP_CONSTANT, INT_TYPE_ID, CONSTANT_ARRAY_LENGTH_ID, uint32(bufferLength),
			(4 << 16) | OP_TYPE_ARRAY, INT_ARRAY_TYPE_ID, INT_TYPE_ID, CONSTANT_ARRAY_LENGTH_ID,
			(3 << 16) | OP_TYPE_STRUCT, STRUCT_ID, INT_ARRAY_TYPE_ID,
			(4 << 16) | OP_TYPE_POINTER, POINTER_TYPE_ID, UNIFORM, STRUCT_ID,
			(4 << 16) | OP_TYPE_POINTER, ELEMENT_POINTER_TYPE_ID, UNIFORM, INT_TYPE_ID,
			(4 << 16) | OP_TYPE_VECTOR, INT_VECTOR_TYPE_ID, INT_TYPE_ID, 3,
			(4 << 16) | OP_TYPE_POINTER, INT_VECTOR_POINTER_TYPE_ID, INPUT, INT_VECTOR_TYPE_ID,
			(4 << 16) | OP_TYPE_POINTER, INT_POINTER_TYPE_ID, INPUT, INT_TYPE_ID,
			(4 << 16) | OP_CONSTANT, INT_TYPE_ID, CONSTANT_ZERO_ID, 0,
			(4 << 16) | OP_VARIABLE, POINTER_TYPE_ID, IN_ID, UNIFORM,
			(4 << 16) | OP_VARIABLE, POINTER_TYPE_ID, OUT_ID, UNIFORM,
			(4 << 16) | OP_VARIABLE, INT_VECTOR_POINTER_TYPE_ID, GLOBAL_INVOCATION_ID, INPUT,
			(5 << 16) | OP_FUNCTION, VOID_TYPE_ID, FUNC_ID, 0, FUNC_TYPE_ID,
			(2 << 16) | OP_LABEL, LABEL_ID,
			(5 << 16) | OP_ACCESS_CHAIN, INT_POINTER_TYPE_ID, GLOBAL_INVOCATION_X_PTR_ID, GLOBAL_INVOCATION_ID, CONSTANT_ZERO_ID,
			(4 << 16) | OP_LOAD, INT_TYPE_ID, GLOBAL_INVOCATION_X_ID, GLOBAL_INVOCATION_X_PTR_ID,
			(6 << 16) | OP_ACCESS_CHAIN, ELEMENT_POINTER_TYPE_ID, IN_ELEMENT_ID, IN_ID, CONSTANT_ZERO_ID, GLOBAL_INVOCATION_X_ID,
			(4 << 16) | OP_LOAD, INT_TYPE_ID, TEMP_LOADED_ID, IN_ELEMENT_ID,
			(6 << 16) | OP_ACCESS_CHAIN, ELEMENT_POINTER_TYPE_ID, OUT_ELEMENT_ID, OUT_ID, CONSTANT_ZERO_ID, GLOBAL_INVOCATION_X_ID,
			(3 << 16) | OP_STORE, OUT_ELEMENT_ID, TEMP_LOADED_ID,
			(1 << 16) | OP_RETURN,
			(1 << 16) | OP_FUNCTION_END,
		}

		shaderModuleCreateInfo := vulkan.ShaderModuleCreateInfo{
			Code: shader,
		}

		shaderModule, res := vulkan.CreateShaderModule(device, &shaderModuleCreateInfo, nil)
		bailOnBadResult(res)
		defer vulkan.DestroyShaderModule(device, shaderModule, nil)

		descriptorSetLayoutBindings := []vulkan.DescriptorSetLayoutBinding{
			{
				Binding:         0,
				DescriptorType:  vulkan.DESCRIPTOR_TYPE_STORAGE_BUFFER,
				DescriptorCount: 1,
				StageFlags:      vulkan.SHADER_STAGE_COMPUTE_BIT,
			},
			{
				Binding:         1,
				DescriptorType:  vulkan.DESCRIPTOR_TYPE_STORAGE_BUFFER,
				DescriptorCount: 1,
				StageFlags:      vulkan.SHADER_STAGE_COMPUTE_BIT,
			},
		}

		descriptorSetLayoutCreateInfo := vulkan.DescriptorSetLayoutCreateInfo{
			Bindings: descriptorSetLayoutBindings,
		}

		descriptorSetLayout, res := vulkan.CreateDescriptorSetLayout(device, &descriptorSetLayoutCreateInfo, nil)
		bailOnBadResult(res)
		defer vulkan.DestroyDescriptorSetLayout(device, descriptorSetLayout, nil)

		pipelineLayoutCreateInfo := vulkan.PipelineLayoutCreateInfo{
			SetLayouts: []vulkan.DescriptorSetLayout{descriptorSetLayout},
		}

		pipelineLayout, res := vulkan.CreatePipelineLayout(device, &pipelineLayoutCreateInfo, nil)
		bailOnBadResult(res)
		defer vulkan.DestroyPipelineLayout(device, pipelineLayout, nil)

		computePipelineCreateInfo := vulkan.ComputePipelineCreateInfo{
			Stage: vulkan.PipelineShaderStageCreateInfo{
				Stage:  vulkan.SHADER_STAGE_COMPUTE_BIT,
				Module: shaderModule,
				Name:   "f",
			},
			Layout: pipelineLayout,
		}

		pipeline, res := vulkan.CreateComputePipelines(device, 0, []vulkan.ComputePipelineCreateInfo{computePipelineCreateInfo}, nil)
		bailOnBadResult(res)
		defer vulkan.DestroyPipeline(device, pipeline, nil)

		commandPoolCreateInfo := vulkan.CommandPoolCreateInfo{
			QueueFamilyIndex: queueFamilyIndex,
		}

		descriptorPoolCreateInfo := vulkan.DescriptorPoolCreateInfo{
			MaxSets: 1,
			PoolSizes: []vulkan.DescriptorPoolSize{
				{
					Type:            vulkan.DESCRIPTOR_TYPE_STORAGE_BUFFER,
					DescriptorCount: 2,
				},
			},
		}

		descriptorPool, res := vulkan.CreateDescriptorPool(device, &descriptorPoolCreateInfo, nil)
		bailOnBadResult(res)
		defer vulkan.DestroyDescriptorPool(device, descriptorPool, nil)

		descriptorSetAllocateInfo := vulkan.DescriptorSetAllocateInfo{
			DescriptorPool: descriptorPool,
			SetLayouts:     []vulkan.DescriptorSetLayout{descriptorSetLayout},
		}

		descriptorSet, res := vulkan.AllocateDescriptorSets(device, &descriptorSetAllocateInfo)
		bailOnBadResult(res)

		inDescriptorBufferInfo := vulkan.DescriptorBufferInfo{
			Buffer: inBuffer,
			Offset: 0,
			Range:  0xFFFF_FFFF_FFFF_FFFF,
		}

		outDescriptorBufferInfo := vulkan.DescriptorBufferInfo{
			Buffer: outBuffer,
			Offset: 0,
			Range:  0xFFFF_FFFF_FFFF_FFFF,
		}

		writeDescriptorSets := []vulkan.WriteDescriptorSet{
			{
				DstSet:         descriptorSet,
				DstBinding:     0,
				DescriptorType: vulkan.DESCRIPTOR_TYPE_STORAGE_BUFFER,
				BufferInfo:     []vulkan.DescriptorBufferInfo{inDescriptorBufferInfo},
			},
			{
				DstSet:         descriptorSet,
				DstBinding:     1,
				DescriptorType: vulkan.DESCRIPTOR_TYPE_STORAGE_BUFFER,
				BufferInfo:     []vulkan.DescriptorBufferInfo{outDescriptorBufferInfo},
			},
		}

		vulkan.UpdateDescriptorSets(device, writeDescriptorSets, nil)

		commandPool, res := vulkan.CreateCommandPool(device, &commandPoolCreateInfo, nil)
		bailOnBadResult(res)
		defer vulkan.DestroyCommandPool(device, commandPool, nil)

		commandBufferAllocateInfo := vulkan.CommandBufferAllocateInfo{
			CommandPool:        commandPool,
			Level:              vulkan.COMMAND_BUFFER_LEVEL_PRIMARY,
			CommandBufferCount: 1,
		}

		commandBuffer, res := vulkan.AllocateCommandBuffers(device, &commandBufferAllocateInfo)
		bailOnBadResult(res)

		commandBufferBeginInfo := vulkan.CommandBufferBeginInfo{
			Flags: vulkan.COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT,
		}

		bailOnBadResult(vulkan.BeginCommandBuffer(commandBuffer, &commandBufferBeginInfo))
		vulkan.CmdBindPipeline(commandBuffer, vulkan.PIPELINE_BIND_POINT_COMPUTE, pipeline)
		vulkan.CmdBindDescriptorSets(commandBuffer, vulkan.PIPELINE_BIND_POINT_COMPUTE, pipelineLayout, 0, []vulkan.DescriptorSet{descriptorSet}, nil)
		vulkan.CmdDispatch(commandBuffer, uint32(bufferSize/4), 1, 1)
		bailOnBadResult(vulkan.EndCommandBuffer(commandBuffer))

		queue := vulkan.GetDeviceQueue(device, queueFamilyIndex, 0)

		submitInfo := vulkan.SubmitInfo{
			CommandBuffers: []vulkan.CommandBuffer{commandBuffer},
		}

		bailOnBadResult(vulkan.QueueSubmit(queue, []vulkan.SubmitInfo{submitInfo}, 0))
		bailOnBadResult(vulkan.QueueWaitIdle(queue))

		payloadPtr, res = vulkan.MapMemory(device, memory, 0, memorySize, 0)
		bailOnBadResult(res)
		payload = unsafe.Slice((*int32)(payloadPtr), memorySize/4)

		k := uint64(0)
		e := uint64(bufferSize / 4)
		for ; k < e; k++ {
			if payload[k+e] != payload[k] {
				t.Fatalf("Mismatch at %d: expected %d, got %d", k, payload[k], payload[k+e])
			}
		}
		vulkan.UnmapMemory(device, memory)
		t.Log("Compute shader execution and verification succeeded!")
	}
}
