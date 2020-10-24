package vk

import (
	"runtime"
	"testing"
	"unsafe"

	"cld.moe/vk/generator/util"
)

var min = util.Min
var max = util.Max

func QuickCString(s string) *byte {	
	a := make([]byte, len(s) + 1)
	copy(a, s)
	return &a[0]
}

func BAIL_ON_BAD_RESULT(res int32) {
	if res != VK_SUCCESS {
		panic(res)
	}
}

func vkGetBestTransferQueueNPH(physicalDevice VkPhysicalDevice, queueFamilyIndex *uint32) VkResult {
	queueFamilyPropertiesCount := uint32(0)
	vkGetPhysicalDeviceQueueFamilyProperties(physicalDevice, &queueFamilyPropertiesCount, nil)

	queueFamilyProperties := make([]VkQueueFamilyProperties, queueFamilyPropertiesCount)
	vkGetPhysicalDeviceQueueFamilyProperties(physicalDevice, &queueFamilyPropertiesCount, &queueFamilyProperties[0])

	// lastly get any queue that'll work for us (graphics, compute or transfer bit set)
	for i := uint32(0); i < queueFamilyPropertiesCount; i++ {

		// mask out the sparse binding bit that we aren't caring about (yet!)
		a := ^VK_QUEUE_SPARSE_BINDING_BIT // Bypass Contant Evaluator
		maskedFlags := (uint32(a) & queueFamilyProperties[i].queueFlags)

		if (VK_QUEUE_GRAPHICS_BIT|VK_QUEUE_COMPUTE_BIT|VK_QUEUE_TRANSFER_BIT)&maskedFlags != 0 {
			*queueFamilyIndex = i
			return VK_SUCCESS
		}
	}

	return VK_ERROR_INITIALIZATION_FAILED
}

func vkGetBestComputeQueueNPH(physicalDevice VkPhysicalDevice, queueFamilyIndex *uint32) VkResult {
	queueFamilyPropertiesCount := uint32(0)
	vkGetPhysicalDeviceQueueFamilyProperties(physicalDevice, &queueFamilyPropertiesCount, nil)

	queueFamilyProperties := make([]VkQueueFamilyProperties, queueFamilyPropertiesCount)

	vkGetPhysicalDeviceQueueFamilyProperties(physicalDevice, &queueFamilyPropertiesCount, &queueFamilyProperties[0])

	// lastly get any queue that'll work for us
	for i := uint32(0); i < queueFamilyPropertiesCount; i++ {
		// mask out the sparse binding bit that we aren't caring about (yet!) and the transfer bit
		a := ^(VK_QUEUE_SPARSE_BINDING_BIT | VK_QUEUE_SPARSE_BINDING_BIT) // Bypass Contant Evaluator
		maskedFlags := uint32(a) & queueFamilyProperties[i].queueFlags

		if VK_QUEUE_COMPUTE_BIT & maskedFlags != 0 {
			*queueFamilyIndex = i
			return VK_SUCCESS
		}
	}

	return VK_ERROR_INITIALIZATION_FAILED
}


func TestCompute(tst *testing.T) {
	runtime.LockOSThread()

	applicationInfo := VkApplicationInfo{
		VK_STRUCTURE_TYPE_APPLICATION_INFO,
		0,
		QuickCString("VKComputeSample"),
		0,
		QuickCString(""),
		0,
		VK_MAKE_VERSION(1, 0, 9),
	}

	instanceCreateInfo := VkInstanceCreateInfo{
		VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO,
		0,
		0,
		&applicationInfo,
		0,
		nil,
		0,
		nil,
	}

	var instance VkInstance
	BAIL_ON_BAD_RESULT(vkCreateInstance(&instanceCreateInfo, nil, &instance))

	var physicalDeviceCount uint32
	BAIL_ON_BAD_RESULT(vkEnumeratePhysicalDevices(instance, &physicalDeviceCount, nil))

	physicalDevices := make([]VkPhysicalDevice, physicalDeviceCount)

	BAIL_ON_BAD_RESULT(vkEnumeratePhysicalDevices(instance, &physicalDeviceCount, &physicalDevices[0]))

	for i := uint32(0); i < physicalDeviceCount; i++ {
		queueFamilyIndex := uint32(0)
		BAIL_ON_BAD_RESULT(vkGetBestComputeQueueNPH(physicalDevices[i], &queueFamilyIndex))

		queuePrioritory := float32(1.0)
		deviceQueueCreateInfo := VkDeviceQueueCreateInfo{
			VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO,
			0,
			0,
			queueFamilyIndex,
			1,
			&queuePrioritory,
		}

		deviceCreateInfo := VkDeviceCreateInfo{
			VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO,
			0,
			0,
			1,
			&deviceQueueCreateInfo,
			0,
			nil,
			0,
			nil,
			nil,
		}

		var device VkDevice
		BAIL_ON_BAD_RESULT(vkCreateDevice(physicalDevices[i], &deviceCreateInfo, nil, &device))

		var properties VkPhysicalDeviceMemoryProperties
		vkGetPhysicalDeviceMemoryProperties(physicalDevices[i], &properties)
		
		tst.Log(util.Vardump(deviceCreateInfo))

		bufferLength := int32(16384)
		bufferSize := VkDeviceSize(4 * bufferLength)

		// we are going to need two buffers from this one memory
		memorySize := VkDeviceSize(bufferSize * 2)

		// set memoryTypeIndex to an invalid entry in the properties.memoryTypes array
		memoryTypeIndex := uint32(VK_MAX_MEMORY_TYPES)

		for k := uint32(0); k < properties.memoryTypeCount; k++ {
			if (VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT & properties.memoryTypes[k].propertyFlags != 0) &&
				(VK_MEMORY_PROPERTY_HOST_COHERENT_BIT & properties.memoryTypes[k].propertyFlags != 0) &&
				(memorySize < properties.memoryHeaps[properties.memoryTypes[k].heapIndex].size) {
				memoryTypeIndex = k
				break
			}
		}

		if memoryTypeIndex == VK_MAX_MEMORY_TYPES {
			panic(VK_ERROR_OUT_OF_HOST_MEMORY)
		}

		memoryAllocateInfo := VkMemoryAllocateInfo{
			VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO,
			0,
			memorySize,
			memoryTypeIndex,
		}

		var memory VkDeviceMemory
		BAIL_ON_BAD_RESULT(vkAllocateMemory(device, &memoryAllocateInfo, nil, &memory))

		var _payload uintptr
		BAIL_ON_BAD_RESULT(vkMapMemory(device, memory, 0, memorySize, 0, &_payload))
		if _payload == 0 {
			panic(0)
		}

		payload := (*[1 << 20]int32)(unsafe.Pointer(_payload))

		for k := uint64(1); k < (memorySize/4); k++ {
			payload[k] = int32(1) //rand()
		}

		vkUnmapMemory(device, memory)

		bufferCreateInfo := VkBufferCreateInfo{
			VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO,
			0,
			0,
			uint64(bufferSize),
			uint32(VK_BUFFER_USAGE_STORAGE_BUFFER_BIT),
			VK_SHARING_MODE_EXCLUSIVE,
			1,
			&queueFamilyIndex,
		}

		var in_buffer VkBuffer
		BAIL_ON_BAD_RESULT(vkCreateBuffer(device, &bufferCreateInfo, nil, &in_buffer))

		BAIL_ON_BAD_RESULT(vkBindBufferMemory(device, in_buffer, memory, 0))

		var out_buffer VkBuffer
		BAIL_ON_BAD_RESULT(vkCreateBuffer(device, &bufferCreateInfo, nil, &out_buffer))

		BAIL_ON_BAD_RESULT(vkBindBufferMemory(device, out_buffer, memory, bufferSize))

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
			// first is the SPIR-V header
			0x07230203, // magic header ID
			0x00010000, // version 1.0.0
			0,          // generator (optional)
			BOUND,      // bound
			0,          // schema

			// OpCapability Shader
			(2 << 16) | OP_CAPABILITY, 1,
			// OpMemoryModel Logical Simple
			(3 << 16) | OP_MEMORY_MODEL, 0, 0,
			// OpEntryPoint GLCompute %FUNC_ID "f" %IN_ID %OUT_ID
			(4 << 16) | OP_ENTRY_POINT, 5, FUNC_ID, 0x00000066,
			// OpExecutionMode %FUNC_ID LocalSize 1 1 1
			(6 << 16) | OP_EXECUTION_MODE, FUNC_ID, 17, 1, 1, 1,
			// next declare decorations
			(3 << 16) | OP_DECORATE, STRUCT_ID, BUFFER_BLOCK,
			(4 << 16) | OP_DECORATE, GLOBAL_INVOCATION_ID, BUILTIN, GLOBAL_INVOCATION,
			(4 << 16) | OP_DECORATE, IN_ID, DESCRIPTOR_SET, 0,
			(4 << 16) | OP_DECORATE, IN_ID, BINDING, 0,
			(4 << 16) | OP_DECORATE, OUT_ID, DESCRIPTOR_SET, 0,
			(4 << 16) | OP_DECORATE, OUT_ID, BINDING, 1,
			(4 << 16) | OP_DECORATE, INT_ARRAY_TYPE_ID, ARRAY_STRIDE, 4,
			(5 << 16) | OP_MEMBER_DECORATE, STRUCT_ID, 0, OFFSET, 0,
			// next declare types
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
			// then declare constants
			(4 << 16) | OP_CONSTANT, INT_TYPE_ID, CONSTANT_ZERO_ID, 0,
			// then declare variables
			(4 << 16) | OP_VARIABLE, POINTER_TYPE_ID, IN_ID, UNIFORM,
			(4 << 16) | OP_VARIABLE, POINTER_TYPE_ID, OUT_ID, UNIFORM,
			(4 << 16) | OP_VARIABLE, INT_VECTOR_POINTER_TYPE_ID, GLOBAL_INVOCATION_ID, INPUT,
			// then declare function
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

		shaderModuleCreateInfo := VkShaderModuleCreateInfo{
			VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO,
			0,
			0,
			uint(len(shader) * 4),
			&shader[0],
		}

		var shader_module VkShaderModule

		BAIL_ON_BAD_RESULT(vkCreateShaderModule(device, &shaderModuleCreateInfo, nil, &shader_module))

		descriptorSetLayoutBindings := [2]VkDescriptorSetLayoutBinding{
			{
				0,
				VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
				1,
				uint32(VK_SHADER_STAGE_COMPUTE_BIT),
				nil,
			},
			{
				1,
				VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
				1,
				uint32(VK_SHADER_STAGE_COMPUTE_BIT),
				nil,
			},
		}

		descriptorSetLayoutCreateInfo := VkDescriptorSetLayoutCreateInfo{
			VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO,
			0,
			0,
			2,
			&descriptorSetLayoutBindings[0],
		}

		var descriptorSetLayout VkDescriptorSetLayout
		BAIL_ON_BAD_RESULT(vkCreateDescriptorSetLayout(device, &descriptorSetLayoutCreateInfo, nil, &descriptorSetLayout))

		pipelineLayoutCreateInfo := VkPipelineLayoutCreateInfo{
			VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO,
			0,
			0,
			1,
			&descriptorSetLayout,
			0,
			nil,
		}

		var pipelineLayout VkPipelineLayout
		BAIL_ON_BAD_RESULT(vkCreatePipelineLayout(device, &pipelineLayoutCreateInfo, nil, &pipelineLayout))

		computePipelineCreateInfo := VkComputePipelineCreateInfo{
			VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO,
			0,
			0,
			VkPipelineShaderStageCreateInfo{
				VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO,
				0,
				0,
				VK_SHADER_STAGE_COMPUTE_BIT,
				shader_module,
				QuickCString("f"),
				nil,
			},
			pipelineLayout,
			0,
			0,
		}

		var pipeline VkPipeline
		BAIL_ON_BAD_RESULT(vkCreateComputePipelines(device, 0, 1, &computePipelineCreateInfo, nil, &pipeline))

		commandPoolCreateInfo := VkCommandPoolCreateInfo{
			VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO,
			0,
			0,
			queueFamilyIndex,
		}

		descriptorPoolSize := VkDescriptorPoolSize{
			VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
			2,
		}

		descriptorPoolCreateInfo := VkDescriptorPoolCreateInfo{
			VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO,
			0,
			0,
			1,
			1,
			&descriptorPoolSize,
		}

		var descriptorPool VkDescriptorPool
		BAIL_ON_BAD_RESULT(vkCreateDescriptorPool(device, &descriptorPoolCreateInfo, nil, &descriptorPool))

		descriptorSetAllocateInfo := VkDescriptorSetAllocateInfo{
			VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO,
			0,
			descriptorPool,
			1,
			&descriptorSetLayout,
		}

		var descriptorSet VkDescriptorSet
		BAIL_ON_BAD_RESULT(vkAllocateDescriptorSets(device, &descriptorSetAllocateInfo, &descriptorSet))

		in_descriptorBufferInfo := VkDescriptorBufferInfo{
			in_buffer,
			0,
			0xFFFF_FFFF_FFFF_FFFF,
		}

		out_descriptorBufferInfo := VkDescriptorBufferInfo{
			out_buffer,
			0,
			0xFFFF_FFFF_FFFF_FFFF,
		}

		writeDescriptorSet := [2]VkWriteDescriptorSet{
			{
				VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET,
				0,
				descriptorSet,
				0,
				0,
				1,
				VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
				nil,
				&in_descriptorBufferInfo,
				nil,
			},
			{
				VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET,
				0,
				descriptorSet,
				1,
				0,
				1,
				VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
				nil,
				&out_descriptorBufferInfo,
				nil,
			},
		}

		vkUpdateDescriptorSets(device, 2, &writeDescriptorSet[0], 0, nil)

		var commandPool VkCommandPool
		BAIL_ON_BAD_RESULT(vkCreateCommandPool(device, &commandPoolCreateInfo, nil, &commandPool))

		commandBufferAllocateInfo := VkCommandBufferAllocateInfo{
			VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO,
			0,
			commandPool,
			VK_COMMAND_BUFFER_LEVEL_PRIMARY,
			1,
		}

		var commandBuffer VkCommandBuffer
		BAIL_ON_BAD_RESULT(vkAllocateCommandBuffers(device, &commandBufferAllocateInfo, &commandBuffer))

		commandBufferBeginInfo := VkCommandBufferBeginInfo{
			VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO,
			0,
			uint32(VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT),
			nil,
		}

		BAIL_ON_BAD_RESULT(vkBeginCommandBuffer(commandBuffer, &commandBufferBeginInfo))

		vkCmdBindPipeline(commandBuffer, VK_PIPELINE_BIND_POINT_COMPUTE, pipeline)

		vkCmdBindDescriptorSets(commandBuffer, VK_PIPELINE_BIND_POINT_COMPUTE,
			pipelineLayout, 0, 1, &descriptorSet, 0, nil)

		vkCmdDispatch(commandBuffer, uint32(bufferSize/4), 1, 1)

		BAIL_ON_BAD_RESULT(vkEndCommandBuffer(commandBuffer))

		var queue VkQueue
		vkGetDeviceQueue(device, queueFamilyIndex, 0, &queue)

		submitInfo := VkSubmitInfo{
			VK_STRUCTURE_TYPE_SUBMIT_INFO,
			0,
			0,
			nil,
			nil,
			1,
			&commandBuffer,
			0,
			nil,
		}

		BAIL_ON_BAD_RESULT(vkQueueSubmit(queue, 1, &submitInfo, 0))
		BAIL_ON_BAD_RESULT(vkQueueWaitIdle(queue))
		BAIL_ON_BAD_RESULT(vkMapMemory(device, memory, 0, memorySize, 0, (*uintptr)(unsafe.Pointer(&payload))))
		k := uint64(0)
		e := bufferSize / 4

		for ;k < e; k++ {
			if payload[k+e] != payload[k] {
				panic(VK_ERROR_OUT_OF_HOST_MEMORY)
			}
		}
	}
}
