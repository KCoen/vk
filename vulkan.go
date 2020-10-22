package main
//----------------------------------------
//--Build-in
//----------------------------------------
type unk = int
type external_type = unk
type funcptr = uintptr
type handle = uintptr
//----------------------------------------
//--Types
//----------------------------------------
type Display = external_type
type VisualID = external_type
type Window = external_type
type RROutput = external_type
type wl_display = external_type
type wl_surface = external_type
type HINSTANCE = external_type
type HWND = external_type
type HMONITOR = external_type
type HANDLE = external_type
type SECURITY_ATTRIBUTES = external_type
type DWORD = external_type
type LPCWSTR = external_type
type xcb_connection_t = external_type
type xcb_visualid_t = external_type
type xcb_window_t = external_type
type IDirectFB = external_type
type IDirectFBSurface = external_type
type zx_handle_t = external_type
type GgpStreamDescriptor = external_type
type GgpFrameToken = external_type
type ANativeWindow = unk
type AHardwareBuffer = unk
type CAMetalLayer = unk
type VkSampleMask = uint32
type VkBool32 = uint32
type VkFlags = uint32
type VkDeviceSize = uint64
type VkDeviceAddress = uint64
type void = external_type
type char = external_type
type float = external_type
type double = external_type
type uint8_t = external_type
type uint16_t = external_type
type uint32_t = external_type
type uint64_t = external_type
type int32_t = external_type
type int64_t = external_type
type size_t = external_type
type VkFramebufferCreateFlags = uint32
type VkQueryPoolCreateFlags = uint32
type VkRenderPassCreateFlags = uint32
type VkSamplerCreateFlags = uint32
type VkPipelineLayoutCreateFlags = uint32
type VkPipelineCacheCreateFlags = uint32
type VkPipelineDepthStencilStateCreateFlags = uint32
type VkPipelineDynamicStateCreateFlags = uint32
type VkPipelineColorBlendStateCreateFlags = uint32
type VkPipelineMultisampleStateCreateFlags = uint32
type VkPipelineRasterizationStateCreateFlags = uint32
type VkPipelineViewportStateCreateFlags = uint32
type VkPipelineTessellationStateCreateFlags = uint32
type VkPipelineInputAssemblyStateCreateFlags = uint32
type VkPipelineVertexInputStateCreateFlags = uint32
type VkPipelineShaderStageCreateFlags = uint32
type VkDescriptorSetLayoutCreateFlags = uint32
type VkBufferViewCreateFlags = uint32
type VkInstanceCreateFlags = uint32
type VkDeviceCreateFlags = uint32
type VkDeviceQueueCreateFlags = uint32
type VkQueueFlags = uint32
type VkMemoryPropertyFlags = uint32
type VkMemoryHeapFlags = uint32
type VkAccessFlags = uint32
type VkBufferUsageFlags = uint32
type VkBufferCreateFlags = uint32
type VkShaderStageFlags = uint32
type VkImageUsageFlags = uint32
type VkImageCreateFlags = uint32
type VkImageViewCreateFlags = uint32
type VkPipelineCreateFlags = uint32
type VkColorComponentFlags = uint32
type VkFenceCreateFlags = uint32
type VkSemaphoreCreateFlags = uint32
type VkFormatFeatureFlags = uint32
type VkQueryControlFlags = uint32
type VkQueryResultFlags = uint32
type VkShaderModuleCreateFlags = uint32
type VkEventCreateFlags = uint32
type VkCommandPoolCreateFlags = uint32
type VkCommandPoolResetFlags = uint32
type VkCommandBufferResetFlags = uint32
type VkCommandBufferUsageFlags = uint32
type VkQueryPipelineStatisticFlags = uint32
type VkMemoryMapFlags = uint32
type VkImageAspectFlags = uint32
type VkSparseMemoryBindFlags = uint32
type VkSparseImageFormatFlags = uint32
type VkSubpassDescriptionFlags = uint32
type VkPipelineStageFlags = uint32
type VkSampleCountFlags = uint32
type VkAttachmentDescriptionFlags = uint32
type VkStencilFaceFlags = uint32
type VkCullModeFlags = uint32
type VkDescriptorPoolCreateFlags = uint32
type VkDescriptorPoolResetFlags = uint32
type VkDependencyFlags = uint32
type VkSubgroupFeatureFlags = uint32
type VkIndirectCommandsLayoutUsageFlagsNV = uint32
type VkIndirectStateFlagsNV = uint32
type VkGeometryFlagsKHR = uint32
type VkGeometryFlagsNV = uint32
type VkGeometryInstanceFlagsKHR = uint32
type VkGeometryInstanceFlagsNV = uint32
type VkBuildAccelerationStructureFlagsKHR = uint32
type VkBuildAccelerationStructureFlagsNV = uint32
type VkPrivateDataSlotCreateFlagsEXT = uint32
type VkDescriptorUpdateTemplateCreateFlags = uint32
type VkDescriptorUpdateTemplateCreateFlagsKHR = uint32
type VkPipelineCreationFeedbackFlagsEXT = uint32
type VkPerformanceCounterDescriptionFlagsKHR = uint32
type VkAcquireProfilingLockFlagsKHR = uint32
type VkSemaphoreWaitFlags = uint32
type VkSemaphoreWaitFlagsKHR = uint32
type VkPipelineCompilerControlFlagsAMD = uint32
type VkShaderCorePropertiesFlagsAMD = uint32
type VkDeviceDiagnosticsConfigFlagsNV = uint32
type VkCompositeAlphaFlagsKHR = uint32
type VkDisplayPlaneAlphaFlagsKHR = uint32
type VkSurfaceTransformFlagsKHR = uint32
type VkSwapchainCreateFlagsKHR = uint32
type VkDisplayModeCreateFlagsKHR = uint32
type VkDisplaySurfaceCreateFlagsKHR = uint32
type VkAndroidSurfaceCreateFlagsKHR = uint32
type VkViSurfaceCreateFlagsNN = uint32
type VkWaylandSurfaceCreateFlagsKHR = uint32
type VkWin32SurfaceCreateFlagsKHR = uint32
type VkXlibSurfaceCreateFlagsKHR = uint32
type VkXcbSurfaceCreateFlagsKHR = uint32
type VkDirectFBSurfaceCreateFlagsEXT = uint32
type VkIOSSurfaceCreateFlagsMVK = uint32
type VkMacOSSurfaceCreateFlagsMVK = uint32
type VkMetalSurfaceCreateFlagsEXT = uint32
type VkImagePipeSurfaceCreateFlagsFUCHSIA = uint32
type VkStreamDescriptorSurfaceCreateFlagsGGP = uint32
type VkHeadlessSurfaceCreateFlagsEXT = uint32
type VkPeerMemoryFeatureFlags = uint32
type VkPeerMemoryFeatureFlagsKHR = uint32
type VkMemoryAllocateFlags = uint32
type VkMemoryAllocateFlagsKHR = uint32
type VkDeviceGroupPresentModeFlagsKHR = uint32
type VkDebugReportFlagsEXT = uint32
type VkCommandPoolTrimFlags = uint32
type VkCommandPoolTrimFlagsKHR = uint32
type VkExternalMemoryHandleTypeFlagsNV = uint32
type VkExternalMemoryFeatureFlagsNV = uint32
type VkExternalMemoryHandleTypeFlags = uint32
type VkExternalMemoryHandleTypeFlagsKHR = uint32
type VkExternalMemoryFeatureFlags = uint32
type VkExternalMemoryFeatureFlagsKHR = uint32
type VkExternalSemaphoreHandleTypeFlags = uint32
type VkExternalSemaphoreHandleTypeFlagsKHR = uint32
type VkExternalSemaphoreFeatureFlags = uint32
type VkExternalSemaphoreFeatureFlagsKHR = uint32
type VkSemaphoreImportFlags = uint32
type VkSemaphoreImportFlagsKHR = uint32
type VkExternalFenceHandleTypeFlags = uint32
type VkExternalFenceHandleTypeFlagsKHR = uint32
type VkExternalFenceFeatureFlags = uint32
type VkExternalFenceFeatureFlagsKHR = uint32
type VkFenceImportFlags = uint32
type VkFenceImportFlagsKHR = uint32
type VkSurfaceCounterFlagsEXT = uint32
type VkPipelineViewportSwizzleStateCreateFlagsNV = uint32
type VkPipelineDiscardRectangleStateCreateFlagsEXT = uint32
type VkPipelineCoverageToColorStateCreateFlagsNV = uint32
type VkPipelineCoverageModulationStateCreateFlagsNV = uint32
type VkPipelineCoverageReductionStateCreateFlagsNV = uint32
type VkValidationCacheCreateFlagsEXT = uint32
type VkDebugUtilsMessageSeverityFlagsEXT = uint32
type VkDebugUtilsMessageTypeFlagsEXT = uint32
type VkDebugUtilsMessengerCreateFlagsEXT = uint32
type VkDebugUtilsMessengerCallbackDataFlagsEXT = uint32
type VkDeviceMemoryReportFlagsEXT = uint32
type VkPipelineRasterizationConservativeStateCreateFlagsEXT = uint32
type VkDescriptorBindingFlags = uint32
type VkDescriptorBindingFlagsEXT = uint32
type VkConditionalRenderingFlagsEXT = uint32
type VkResolveModeFlags = uint32
type VkResolveModeFlagsKHR = uint32
type VkPipelineRasterizationStateStreamCreateFlagsEXT = uint32
type VkPipelineRasterizationDepthClipStateCreateFlagsEXT = uint32
type VkSwapchainImageUsageFlagsANDROID = uint32
type VkToolPurposeFlagsEXT = uint32
type VkInstance = uint64
type VkPhysicalDevice = uint64
type VkDevice = uint64
type VkQueue = uint64
type VkCommandBuffer = uint64
type VkDeviceMemory = uint64
type VkCommandPool = uint64
type VkBuffer = uint64
type VkBufferView = uint64
type VkImage = uint64
type VkImageView = uint64
type VkShaderModule = uint64
type VkPipeline = uint64
type VkPipelineLayout = uint64
type VkSampler = uint64
type VkDescriptorSet = uint64
type VkDescriptorSetLayout = uint64
type VkDescriptorPool = uint64
type VkFence = uint64
type VkSemaphore = uint64
type VkEvent = uint64
type VkQueryPool = uint64
type VkFramebuffer = uint64
type VkRenderPass = uint64
type VkPipelineCache = uint64
type VkIndirectCommandsLayoutNV = uint64
type VkDescriptorUpdateTemplate = uint64
type VkDescriptorUpdateTemplateKHR = uint64
type VkSamplerYcbcrConversion = uint64
type VkSamplerYcbcrConversionKHR = uint64
type VkValidationCacheEXT = uint64
type VkAccelerationStructureKHR = uint64
type VkAccelerationStructureNV = uint64
type VkPerformanceConfigurationINTEL = uint64
type VkDeferredOperationKHR = uint64
type VkPrivateDataSlotEXT = uint64
type VkDisplayKHR = uint64
type VkDisplayModeKHR = uint64
type VkSurfaceKHR = uint64
type VkSwapchainKHR = uint64
type VkDebugReportCallbackEXT = uint64
type VkDebugUtilsMessengerEXT = uint64
type VkDescriptorUpdateTemplateTypeKHR = VkDescriptorUpdateTemplateType
type VkPointClippingBehaviorKHR = VkPointClippingBehavior
type VkResolveModeFlagBitsKHR = VkResolveModeFlagBits
type VkDescriptorBindingFlagBitsEXT = VkDescriptorBindingFlagBits
type VkSemaphoreTypeKHR = VkSemaphoreType
type VkGeometryFlagBitsNV = VkGeometryFlagBitsKHR
type VkGeometryInstanceFlagBitsNV = VkGeometryInstanceFlagBitsKHR
type VkBuildAccelerationStructureFlagBitsNV = VkBuildAccelerationStructureFlagBitsKHR
type VkCopyAccelerationStructureModeNV = VkCopyAccelerationStructureModeKHR
type VkAccelerationStructureTypeNV = VkAccelerationStructureTypeKHR
type VkGeometryTypeNV = VkGeometryTypeKHR
type VkRayTracingShaderGroupTypeNV = VkRayTracingShaderGroupTypeKHR
type VkAccelerationStructureMemoryRequirementsTypeNV = VkAccelerationStructureMemoryRequirementsTypeKHR
type VkSemaphoreWaitFlagBitsKHR = VkSemaphoreWaitFlagBits
type VkExternalMemoryHandleTypeFlagBitsKHR = VkExternalMemoryHandleTypeFlagBits
type VkExternalMemoryFeatureFlagBitsKHR = VkExternalMemoryFeatureFlagBits
type VkExternalSemaphoreHandleTypeFlagBitsKHR = VkExternalSemaphoreHandleTypeFlagBits
type VkExternalSemaphoreFeatureFlagBitsKHR = VkExternalSemaphoreFeatureFlagBits
type VkSemaphoreImportFlagBitsKHR = VkSemaphoreImportFlagBits
type VkExternalFenceHandleTypeFlagBitsKHR = VkExternalFenceHandleTypeFlagBits
type VkExternalFenceFeatureFlagBitsKHR = VkExternalFenceFeatureFlagBits
type VkFenceImportFlagBitsKHR = VkFenceImportFlagBits
type VkPeerMemoryFeatureFlagBitsKHR = VkPeerMemoryFeatureFlagBits
type VkMemoryAllocateFlagBitsKHR = VkMemoryAllocateFlagBits
type VkTessellationDomainOriginKHR = VkTessellationDomainOrigin
type VkSamplerYcbcrModelConversionKHR = VkSamplerYcbcrModelConversion
type VkSamplerYcbcrRangeKHR = VkSamplerYcbcrRange
type VkChromaLocationKHR = VkChromaLocation
type VkSamplerReductionModeEXT = VkSamplerReductionMode
type VkShaderFloatControlsIndependenceKHR = VkShaderFloatControlsIndependence
type VkDriverIdKHR = VkDriverId
type PFN_vkInternalAllocationNotification = funcptr
type PFN_vkInternalFreeNotification = funcptr
type PFN_vkReallocationFunction = funcptr
type PFN_vkAllocationFunction = funcptr
type PFN_vkFreeFunction = funcptr
type PFN_vkVoidFunction = funcptr
type PFN_vkDebugReportCallbackEXT = funcptr
type PFN_vkDebugUtilsMessengerCallbackEXT = funcptr
type PFN_vkDeviceMemoryReportCallbackEXT = funcptr
type VkBaseOutStructure struct {
	sType VkStructureType
	pNext *VkBaseOutStructure
}
type VkBaseInStructure struct {
	sType VkStructureType
	pNext *VkBaseInStructure
}
type VkOffset2D struct {
	x int32
	y int32
}
type VkOffset3D struct {
	x int32
	y int32
	z int32
}
type VkExtent2D struct {
	width uint32
	height uint32
}
type VkExtent3D struct {
	width uint32
	height uint32
	depth uint32
}
type VkViewport struct {
	x float32
	y float32
	width float32
	height float32
	minDepth float32
	maxDepth float32
}
type VkRect2D struct {
	offset VkOffset2D
	extent VkExtent2D
}
type VkClearRect struct {
	rect VkRect2D
	baseArrayLayer uint32
	layerCount uint32
}
type VkComponentMapping struct {
	r VkComponentSwizzle
	g VkComponentSwizzle
	b VkComponentSwizzle
	a VkComponentSwizzle
}
type VkPhysicalDeviceProperties struct {
	apiVersion uint32
	driverVersion uint32
	vendorID uint32
	deviceID uint32
	deviceType VkPhysicalDeviceType
	deviceName byte
	pipelineCacheUUID byte
	limits VkPhysicalDeviceLimits
	sparseProperties VkPhysicalDeviceSparseProperties
}
type VkExtensionProperties struct {
	extensionName byte
	specVersion uint32
}
type VkLayerProperties struct {
	layerName byte
	specVersion uint32
	implementationVersion uint32
	description byte
}
type VkApplicationInfo struct {
	sType VkStructureType
	pNext uintptr
	pApplicationName *byte
	applicationVersion uint32
	pEngineName *byte
	engineVersion uint32
	apiVersion uint32
}
type VkAllocationCallbacks struct {
	pUserData uintptr
	pfnAllocation PFN_vkAllocationFunction
	pfnReallocation PFN_vkReallocationFunction
	pfnFree PFN_vkFreeFunction
	pfnInternalAllocation PFN_vkInternalAllocationNotification
	pfnInternalFree PFN_vkInternalFreeNotification
}
type VkDeviceQueueCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkDeviceQueueCreateFlags
	queueFamilyIndex uint32
	queueCount uint32
	pQueuePriorities *float32
}
type VkDeviceCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkDeviceCreateFlags
	queueCreateInfoCount uint32
	pQueueCreateInfos *VkDeviceQueueCreateInfo
	enabledLayerCount uint32
	ppEnabledLayerNames **byte
	enabledExtensionCount uint32
	ppEnabledExtensionNames **byte
	pEnabledFeatures *VkPhysicalDeviceFeatures
}
type VkInstanceCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkInstanceCreateFlags
	pApplicationInfo *VkApplicationInfo
	enabledLayerCount uint32
	ppEnabledLayerNames **byte
	enabledExtensionCount uint32
	ppEnabledExtensionNames **byte
}
type VkQueueFamilyProperties struct {
	queueFlags VkQueueFlags
	queueCount uint32
	timestampValidBits uint32
	minImageTransferGranularity VkExtent3D
}
type VkPhysicalDeviceMemoryProperties struct {
	memoryTypeCount uint32
	memoryTypes VkMemoryType
	memoryHeapCount uint32
	memoryHeaps VkMemoryHeap
}
type VkMemoryAllocateInfo struct {
	sType VkStructureType
	pNext uintptr
	allocationSize VkDeviceSize
	memoryTypeIndex uint32
}
type VkMemoryRequirements struct {
	size VkDeviceSize
	alignment VkDeviceSize
	memoryTypeBits uint32
}
type VkSparseImageFormatProperties struct {
	aspectMask VkImageAspectFlags
	imageGranularity VkExtent3D
	flags VkSparseImageFormatFlags
}
type VkSparseImageMemoryRequirements struct {
	formatProperties VkSparseImageFormatProperties
	imageMipTailFirstLod uint32
	imageMipTailSize VkDeviceSize
	imageMipTailOffset VkDeviceSize
	imageMipTailStride VkDeviceSize
}
type VkMemoryType struct {
	propertyFlags VkMemoryPropertyFlags
	heapIndex uint32
}
type VkMemoryHeap struct {
	size VkDeviceSize
	flags VkMemoryHeapFlags
}
type VkMappedMemoryRange struct {
	sType VkStructureType
	pNext uintptr
	memory VkDeviceMemory
	offset VkDeviceSize
	size VkDeviceSize
}
type VkFormatProperties struct {
	linearTilingFeatures VkFormatFeatureFlags
	optimalTilingFeatures VkFormatFeatureFlags
	bufferFeatures VkFormatFeatureFlags
}
type VkImageFormatProperties struct {
	maxExtent VkExtent3D
	maxMipLevels uint32
	maxArrayLayers uint32
	sampleCounts VkSampleCountFlags
	maxResourceSize VkDeviceSize
}
type VkDescriptorBufferInfo struct {
	buffer VkBuffer
	offset VkDeviceSize
	range0 VkDeviceSize
}
type VkDescriptorImageInfo struct {
	sampler VkSampler
	imageView VkImageView
	imageLayout VkImageLayout
}
type VkWriteDescriptorSet struct {
	sType VkStructureType
	pNext uintptr
	dstSet VkDescriptorSet
	dstBinding uint32
	dstArrayElement uint32
	descriptorCount uint32
	descriptorType VkDescriptorType
	pImageInfo *VkDescriptorImageInfo
	pBufferInfo *VkDescriptorBufferInfo
	pTexelBufferView *VkBufferView
}
type VkCopyDescriptorSet struct {
	sType VkStructureType
	pNext uintptr
	srcSet VkDescriptorSet
	srcBinding uint32
	srcArrayElement uint32
	dstSet VkDescriptorSet
	dstBinding uint32
	dstArrayElement uint32
	descriptorCount uint32
}
type VkBufferCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkBufferCreateFlags
	size VkDeviceSize
	usage VkBufferUsageFlags
	sharingMode VkSharingMode
	queueFamilyIndexCount uint32
	pQueueFamilyIndices *uint32
}
type VkBufferViewCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkBufferViewCreateFlags
	buffer VkBuffer
	format VkFormat
	offset VkDeviceSize
	range0 VkDeviceSize
}
type VkImageSubresource struct {
	aspectMask VkImageAspectFlags
	mipLevel uint32
	arrayLayer uint32
}
type VkImageSubresourceLayers struct {
	aspectMask VkImageAspectFlags
	mipLevel uint32
	baseArrayLayer uint32
	layerCount uint32
}
type VkImageSubresourceRange struct {
	aspectMask VkImageAspectFlags
	baseMipLevel uint32
	levelCount uint32
	baseArrayLayer uint32
	layerCount uint32
}
type VkMemoryBarrier struct {
	sType VkStructureType
	pNext uintptr
	srcAccessMask VkAccessFlags
	dstAccessMask VkAccessFlags
}
type VkBufferMemoryBarrier struct {
	sType VkStructureType
	pNext uintptr
	srcAccessMask VkAccessFlags
	dstAccessMask VkAccessFlags
	srcQueueFamilyIndex uint32
	dstQueueFamilyIndex uint32
	buffer VkBuffer
	offset VkDeviceSize
	size VkDeviceSize
}
type VkImageMemoryBarrier struct {
	sType VkStructureType
	pNext uintptr
	srcAccessMask VkAccessFlags
	dstAccessMask VkAccessFlags
	oldLayout VkImageLayout
	newLayout VkImageLayout
	srcQueueFamilyIndex uint32
	dstQueueFamilyIndex uint32
	image VkImage
	subresourceRange VkImageSubresourceRange
}
type VkImageCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkImageCreateFlags
	imageType VkImageType
	format VkFormat
	extent VkExtent3D
	mipLevels uint32
	arrayLayers uint32
	samples VkSampleCountFlagBits
	tiling VkImageTiling
	usage VkImageUsageFlags
	sharingMode VkSharingMode
	queueFamilyIndexCount uint32
	pQueueFamilyIndices *uint32
	initialLayout VkImageLayout
}
type VkSubresourceLayout struct {
	offset VkDeviceSize
	size VkDeviceSize
	rowPitch VkDeviceSize
	arrayPitch VkDeviceSize
	depthPitch VkDeviceSize
}
type VkImageViewCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkImageViewCreateFlags
	image VkImage
	viewType VkImageViewType
	format VkFormat
	components VkComponentMapping
	subresourceRange VkImageSubresourceRange
}
type VkBufferCopy struct {
	srcOffset VkDeviceSize
	dstOffset VkDeviceSize
	size VkDeviceSize
}
type VkSparseMemoryBind struct {
	resourceOffset VkDeviceSize
	size VkDeviceSize
	memory VkDeviceMemory
	memoryOffset VkDeviceSize
	flags VkSparseMemoryBindFlags
}
type VkSparseImageMemoryBind struct {
	subresource VkImageSubresource
	offset VkOffset3D
	extent VkExtent3D
	memory VkDeviceMemory
	memoryOffset VkDeviceSize
	flags VkSparseMemoryBindFlags
}
type VkSparseBufferMemoryBindInfo struct {
	buffer VkBuffer
	bindCount uint32
	pBinds *VkSparseMemoryBind
}
type VkSparseImageOpaqueMemoryBindInfo struct {
	image VkImage
	bindCount uint32
	pBinds *VkSparseMemoryBind
}
type VkSparseImageMemoryBindInfo struct {
	image VkImage
	bindCount uint32
	pBinds *VkSparseImageMemoryBind
}
type VkBindSparseInfo struct {
	sType VkStructureType
	pNext uintptr
	waitSemaphoreCount uint32
	pWaitSemaphores *VkSemaphore
	bufferBindCount uint32
	pBufferBinds *VkSparseBufferMemoryBindInfo
	imageOpaqueBindCount uint32
	pImageOpaqueBinds *VkSparseImageOpaqueMemoryBindInfo
	imageBindCount uint32
	pImageBinds *VkSparseImageMemoryBindInfo
	signalSemaphoreCount uint32
	pSignalSemaphores *VkSemaphore
}
type VkImageCopy struct {
	srcSubresource VkImageSubresourceLayers
	srcOffset VkOffset3D
	dstSubresource VkImageSubresourceLayers
	dstOffset VkOffset3D
	extent VkExtent3D
}
type VkImageBlit struct {
	srcSubresource VkImageSubresourceLayers
	srcOffsets VkOffset3D
	dstSubresource VkImageSubresourceLayers
	dstOffsets VkOffset3D
}
type VkBufferImageCopy struct {
	bufferOffset VkDeviceSize
	bufferRowLength uint32
	bufferImageHeight uint32
	imageSubresource VkImageSubresourceLayers
	imageOffset VkOffset3D
	imageExtent VkExtent3D
}
type VkImageResolve struct {
	srcSubresource VkImageSubresourceLayers
	srcOffset VkOffset3D
	dstSubresource VkImageSubresourceLayers
	dstOffset VkOffset3D
	extent VkExtent3D
}
type VkShaderModuleCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkShaderModuleCreateFlags
	codeSize uint
	pCode *uint32
}
type VkDescriptorSetLayoutBinding struct {
	binding uint32
	descriptorType VkDescriptorType
	descriptorCount uint32
	stageFlags VkShaderStageFlags
	pImmutableSamplers *VkSampler
}
type VkDescriptorSetLayoutCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkDescriptorSetLayoutCreateFlags
	bindingCount uint32
	pBindings *VkDescriptorSetLayoutBinding
}
type VkDescriptorPoolSize struct {
	type0 VkDescriptorType
	descriptorCount uint32
}
type VkDescriptorPoolCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkDescriptorPoolCreateFlags
	maxSets uint32
	poolSizeCount uint32
	pPoolSizes *VkDescriptorPoolSize
}
type VkDescriptorSetAllocateInfo struct {
	sType VkStructureType
	pNext uintptr
	descriptorPool VkDescriptorPool
	descriptorSetCount uint32
	pSetLayouts *VkDescriptorSetLayout
}
type VkSpecializationMapEntry struct {
	constantID uint32
	offset uint32
	size uint
}
type VkSpecializationInfo struct {
	mapEntryCount uint32
	pMapEntries *VkSpecializationMapEntry
	dataSize uint
	pData uintptr
}
type VkPipelineShaderStageCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineShaderStageCreateFlags
	stage VkShaderStageFlagBits
	module VkShaderModule
	pName *byte
	pSpecializationInfo *VkSpecializationInfo
}
type VkComputePipelineCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineCreateFlags
	stage VkPipelineShaderStageCreateInfo
	layout VkPipelineLayout
	basePipelineHandle VkPipeline
	basePipelineIndex int32
}
type VkVertexInputBindingDescription struct {
	binding uint32
	stride uint32
	inputRate VkVertexInputRate
}
type VkVertexInputAttributeDescription struct {
	location uint32
	binding uint32
	format VkFormat
	offset uint32
}
type VkPipelineVertexInputStateCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineVertexInputStateCreateFlags
	vertexBindingDescriptionCount uint32
	pVertexBindingDescriptions *VkVertexInputBindingDescription
	vertexAttributeDescriptionCount uint32
	pVertexAttributeDescriptions *VkVertexInputAttributeDescription
}
type VkPipelineInputAssemblyStateCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineInputAssemblyStateCreateFlags
	topology VkPrimitiveTopology
	primitiveRestartEnable VkBool32
}
type VkPipelineTessellationStateCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineTessellationStateCreateFlags
	patchControlPoints uint32
}
type VkPipelineViewportStateCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineViewportStateCreateFlags
	viewportCount uint32
	pViewports *VkViewport
	scissorCount uint32
	pScissors *VkRect2D
}
type VkPipelineRasterizationStateCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineRasterizationStateCreateFlags
	depthClampEnable VkBool32
	rasterizerDiscardEnable VkBool32
	polygonMode VkPolygonMode
	cullMode VkCullModeFlags
	frontFace VkFrontFace
	depthBiasEnable VkBool32
	depthBiasConstantFactor float32
	depthBiasClamp float32
	depthBiasSlopeFactor float32
	lineWidth float32
}
type VkPipelineMultisampleStateCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineMultisampleStateCreateFlags
	rasterizationSamples VkSampleCountFlagBits
	sampleShadingEnable VkBool32
	minSampleShading float32
	pSampleMask *VkSampleMask
	alphaToCoverageEnable VkBool32
	alphaToOneEnable VkBool32
}
type VkPipelineColorBlendAttachmentState struct {
	blendEnable VkBool32
	srcColorBlendFactor VkBlendFactor
	dstColorBlendFactor VkBlendFactor
	colorBlendOp VkBlendOp
	srcAlphaBlendFactor VkBlendFactor
	dstAlphaBlendFactor VkBlendFactor
	alphaBlendOp VkBlendOp
	colorWriteMask VkColorComponentFlags
}
type VkPipelineColorBlendStateCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineColorBlendStateCreateFlags
	logicOpEnable VkBool32
	logicOp VkLogicOp
	attachmentCount uint32
	pAttachments *VkPipelineColorBlendAttachmentState
	blendConstants float32
}
type VkPipelineDynamicStateCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineDynamicStateCreateFlags
	dynamicStateCount uint32
	pDynamicStates *VkDynamicState
}
type VkStencilOpState struct {
	failOp VkStencilOp
	passOp VkStencilOp
	depthFailOp VkStencilOp
	compareOp VkCompareOp
	compareMask uint32
	writeMask uint32
	reference uint32
}
type VkPipelineDepthStencilStateCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineDepthStencilStateCreateFlags
	depthTestEnable VkBool32
	depthWriteEnable VkBool32
	depthCompareOp VkCompareOp
	depthBoundsTestEnable VkBool32
	stencilTestEnable VkBool32
	front VkStencilOpState
	back VkStencilOpState
	minDepthBounds float32
	maxDepthBounds float32
}
type VkGraphicsPipelineCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineCreateFlags
	stageCount uint32
	pStages *VkPipelineShaderStageCreateInfo
	pVertexInputState *VkPipelineVertexInputStateCreateInfo
	pInputAssemblyState *VkPipelineInputAssemblyStateCreateInfo
	pTessellationState *VkPipelineTessellationStateCreateInfo
	pViewportState *VkPipelineViewportStateCreateInfo
	pRasterizationState *VkPipelineRasterizationStateCreateInfo
	pMultisampleState *VkPipelineMultisampleStateCreateInfo
	pDepthStencilState *VkPipelineDepthStencilStateCreateInfo
	pColorBlendState *VkPipelineColorBlendStateCreateInfo
	pDynamicState *VkPipelineDynamicStateCreateInfo
	layout VkPipelineLayout
	renderPass VkRenderPass
	subpass uint32
	basePipelineHandle VkPipeline
	basePipelineIndex int32
}
type VkPipelineCacheCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineCacheCreateFlags
	initialDataSize uint
	pInitialData uintptr
}
type VkPushConstantRange struct {
	stageFlags VkShaderStageFlags
	offset uint32
	size uint32
}
type VkPipelineLayoutCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineLayoutCreateFlags
	setLayoutCount uint32
	pSetLayouts *VkDescriptorSetLayout
	pushConstantRangeCount uint32
	pPushConstantRanges *VkPushConstantRange
}
type VkSamplerCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkSamplerCreateFlags
	magFilter VkFilter
	minFilter VkFilter
	mipmapMode VkSamplerMipmapMode
	addressModeU VkSamplerAddressMode
	addressModeV VkSamplerAddressMode
	addressModeW VkSamplerAddressMode
	mipLodBias float32
	anisotropyEnable VkBool32
	maxAnisotropy float32
	compareEnable VkBool32
	compareOp VkCompareOp
	minLod float32
	maxLod float32
	borderColor VkBorderColor
	unnormalizedCoordinates VkBool32
}
type VkCommandPoolCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkCommandPoolCreateFlags
	queueFamilyIndex uint32
}
type VkCommandBufferAllocateInfo struct {
	sType VkStructureType
	pNext uintptr
	commandPool VkCommandPool
	level VkCommandBufferLevel
	commandBufferCount uint32
}
type VkCommandBufferInheritanceInfo struct {
	sType VkStructureType
	pNext uintptr
	renderPass VkRenderPass
	subpass uint32
	framebuffer VkFramebuffer
	occlusionQueryEnable VkBool32
	queryFlags VkQueryControlFlags
	pipelineStatistics VkQueryPipelineStatisticFlags
}
type VkCommandBufferBeginInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkCommandBufferUsageFlags
	pInheritanceInfo *VkCommandBufferInheritanceInfo
}
type VkRenderPassBeginInfo struct {
	sType VkStructureType
	pNext uintptr
	renderPass VkRenderPass
	framebuffer VkFramebuffer
	renderArea VkRect2D
	clearValueCount uint32
	pClearValues *VkClearValue
}
type VkClearColorValue = struct{uintptr}
type VkClearDepthStencilValue struct {
	depth float32
	stencil uint32
}
type VkClearValue = struct{uintptr}
type VkClearAttachment struct {
	aspectMask VkImageAspectFlags
	colorAttachment uint32
	clearValue VkClearValue
}
type VkAttachmentDescription struct {
	flags VkAttachmentDescriptionFlags
	format VkFormat
	samples VkSampleCountFlagBits
	loadOp VkAttachmentLoadOp
	storeOp VkAttachmentStoreOp
	stencilLoadOp VkAttachmentLoadOp
	stencilStoreOp VkAttachmentStoreOp
	initialLayout VkImageLayout
	finalLayout VkImageLayout
}
type VkAttachmentReference struct {
	attachment uint32
	layout VkImageLayout
}
type VkSubpassDescription struct {
	flags VkSubpassDescriptionFlags
	pipelineBindPoint VkPipelineBindPoint
	inputAttachmentCount uint32
	pInputAttachments *VkAttachmentReference
	colorAttachmentCount uint32
	pColorAttachments *VkAttachmentReference
	pResolveAttachments *VkAttachmentReference
	pDepthStencilAttachment *VkAttachmentReference
	preserveAttachmentCount uint32
	pPreserveAttachments *uint32
}
type VkSubpassDependency struct {
	srcSubpass uint32
	dstSubpass uint32
	srcStageMask VkPipelineStageFlags
	dstStageMask VkPipelineStageFlags
	srcAccessMask VkAccessFlags
	dstAccessMask VkAccessFlags
	dependencyFlags VkDependencyFlags
}
type VkRenderPassCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkRenderPassCreateFlags
	attachmentCount uint32
	pAttachments *VkAttachmentDescription
	subpassCount uint32
	pSubpasses *VkSubpassDescription
	dependencyCount uint32
	pDependencies *VkSubpassDependency
}
type VkEventCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkEventCreateFlags
}
type VkFenceCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkFenceCreateFlags
}
type VkPhysicalDeviceFeatures struct {
	robustBufferAccess VkBool32
	fullDrawIndexUint32 VkBool32
	imageCubeArray VkBool32
	independentBlend VkBool32
	geometryShader VkBool32
	tessellationShader VkBool32
	sampleRateShading VkBool32
	dualSrcBlend VkBool32
	logicOp VkBool32
	multiDrawIndirect VkBool32
	drawIndirectFirstInstance VkBool32
	depthClamp VkBool32
	depthBiasClamp VkBool32
	fillModeNonSolid VkBool32
	depthBounds VkBool32
	wideLines VkBool32
	largePoints VkBool32
	alphaToOne VkBool32
	multiViewport VkBool32
	samplerAnisotropy VkBool32
	textureCompressionETC2 VkBool32
	textureCompressionASTC_LDR VkBool32
	textureCompressionBC VkBool32
	occlusionQueryPrecise VkBool32
	pipelineStatisticsQuery VkBool32
	vertexPipelineStoresAndAtomics VkBool32
	fragmentStoresAndAtomics VkBool32
	shaderTessellationAndGeometryPointSize VkBool32
	shaderImageGatherExtended VkBool32
	shaderStorageImageExtendedFormats VkBool32
	shaderStorageImageMultisample VkBool32
	shaderStorageImageReadWithoutFormat VkBool32
	shaderStorageImageWriteWithoutFormat VkBool32
	shaderUniformBufferArrayDynamicIndexing VkBool32
	shaderSampledImageArrayDynamicIndexing VkBool32
	shaderStorageBufferArrayDynamicIndexing VkBool32
	shaderStorageImageArrayDynamicIndexing VkBool32
	shaderClipDistance VkBool32
	shaderCullDistance VkBool32
	shaderFloat64 VkBool32
	shaderInt64 VkBool32
	shaderInt16 VkBool32
	shaderResourceResidency VkBool32
	shaderResourceMinLod VkBool32
	sparseBinding VkBool32
	sparseResidencyBuffer VkBool32
	sparseResidencyImage2D VkBool32
	sparseResidencyImage3D VkBool32
	sparseResidency2Samples VkBool32
	sparseResidency4Samples VkBool32
	sparseResidency8Samples VkBool32
	sparseResidency16Samples VkBool32
	sparseResidencyAliased VkBool32
	variableMultisampleRate VkBool32
	inheritedQueries VkBool32
}
type VkPhysicalDeviceSparseProperties struct {
	residencyStandard2DBlockShape VkBool32
	residencyStandard2DMultisampleBlockShape VkBool32
	residencyStandard3DBlockShape VkBool32
	residencyAlignedMipSize VkBool32
	residencyNonResidentStrict VkBool32
}
type VkPhysicalDeviceLimits struct {
	maxImageDimension1D uint32
	maxImageDimension2D uint32
	maxImageDimension3D uint32
	maxImageDimensionCube uint32
	maxImageArrayLayers uint32
	maxTexelBufferElements uint32
	maxUniformBufferRange uint32
	maxStorageBufferRange uint32
	maxPushConstantsSize uint32
	maxMemoryAllocationCount uint32
	maxSamplerAllocationCount uint32
	bufferImageGranularity VkDeviceSize
	sparseAddressSpaceSize VkDeviceSize
	maxBoundDescriptorSets uint32
	maxPerStageDescriptorSamplers uint32
	maxPerStageDescriptorUniformBuffers uint32
	maxPerStageDescriptorStorageBuffers uint32
	maxPerStageDescriptorSampledImages uint32
	maxPerStageDescriptorStorageImages uint32
	maxPerStageDescriptorInputAttachments uint32
	maxPerStageResources uint32
	maxDescriptorSetSamplers uint32
	maxDescriptorSetUniformBuffers uint32
	maxDescriptorSetUniformBuffersDynamic uint32
	maxDescriptorSetStorageBuffers uint32
	maxDescriptorSetStorageBuffersDynamic uint32
	maxDescriptorSetSampledImages uint32
	maxDescriptorSetStorageImages uint32
	maxDescriptorSetInputAttachments uint32
	maxVertexInputAttributes uint32
	maxVertexInputBindings uint32
	maxVertexInputAttributeOffset uint32
	maxVertexInputBindingStride uint32
	maxVertexOutputComponents uint32
	maxTessellationGenerationLevel uint32
	maxTessellationPatchSize uint32
	maxTessellationControlPerVertexInputComponents uint32
	maxTessellationControlPerVertexOutputComponents uint32
	maxTessellationControlPerPatchOutputComponents uint32
	maxTessellationControlTotalOutputComponents uint32
	maxTessellationEvaluationInputComponents uint32
	maxTessellationEvaluationOutputComponents uint32
	maxGeometryShaderInvocations uint32
	maxGeometryInputComponents uint32
	maxGeometryOutputComponents uint32
	maxGeometryOutputVertices uint32
	maxGeometryTotalOutputComponents uint32
	maxFragmentInputComponents uint32
	maxFragmentOutputAttachments uint32
	maxFragmentDualSrcAttachments uint32
	maxFragmentCombinedOutputResources uint32
	maxComputeSharedMemorySize uint32
	maxComputeWorkGroupCount uint32
	maxComputeWorkGroupInvocations uint32
	maxComputeWorkGroupSize uint32
	subPixelPrecisionBits uint32
	subTexelPrecisionBits uint32
	mipmapPrecisionBits uint32
	maxDrawIndexedIndexValue uint32
	maxDrawIndirectCount uint32
	maxSamplerLodBias float32
	maxSamplerAnisotropy float32
	maxViewports uint32
	maxViewportDimensions uint32
	viewportBoundsRange float32
	viewportSubPixelBits uint32
	minMemoryMapAlignment uint
	minTexelBufferOffsetAlignment VkDeviceSize
	minUniformBufferOffsetAlignment VkDeviceSize
	minStorageBufferOffsetAlignment VkDeviceSize
	minTexelOffset int32
	maxTexelOffset uint32
	minTexelGatherOffset int32
	maxTexelGatherOffset uint32
	minInterpolationOffset float32
	maxInterpolationOffset float32
	subPixelInterpolationOffsetBits uint32
	maxFramebufferWidth uint32
	maxFramebufferHeight uint32
	maxFramebufferLayers uint32
	framebufferColorSampleCounts VkSampleCountFlags
	framebufferDepthSampleCounts VkSampleCountFlags
	framebufferStencilSampleCounts VkSampleCountFlags
	framebufferNoAttachmentsSampleCounts VkSampleCountFlags
	maxColorAttachments uint32
	sampledImageColorSampleCounts VkSampleCountFlags
	sampledImageIntegerSampleCounts VkSampleCountFlags
	sampledImageDepthSampleCounts VkSampleCountFlags
	sampledImageStencilSampleCounts VkSampleCountFlags
	storageImageSampleCounts VkSampleCountFlags
	maxSampleMaskWords uint32
	timestampComputeAndGraphics VkBool32
	timestampPeriod float32
	maxClipDistances uint32
	maxCullDistances uint32
	maxCombinedClipAndCullDistances uint32
	discreteQueuePriorities uint32
	pointSizeRange float32
	lineWidthRange float32
	pointSizeGranularity float32
	lineWidthGranularity float32
	strictLines VkBool32
	standardSampleLocations VkBool32
	optimalBufferCopyOffsetAlignment VkDeviceSize
	optimalBufferCopyRowPitchAlignment VkDeviceSize
	nonCoherentAtomSize VkDeviceSize
}
type VkSemaphoreCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkSemaphoreCreateFlags
}
type VkQueryPoolCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkQueryPoolCreateFlags
	queryType VkQueryType
	queryCount uint32
	pipelineStatistics VkQueryPipelineStatisticFlags
}
type VkFramebufferCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkFramebufferCreateFlags
	renderPass VkRenderPass
	attachmentCount uint32
	pAttachments *VkImageView
	width uint32
	height uint32
	layers uint32
}
type VkDrawIndirectCommand struct {
	vertexCount uint32
	instanceCount uint32
	firstVertex uint32
	firstInstance uint32
}
type VkDrawIndexedIndirectCommand struct {
	indexCount uint32
	instanceCount uint32
	firstIndex uint32
	vertexOffset int32
	firstInstance uint32
}
type VkDispatchIndirectCommand struct {
	x uint32
	y uint32
	z uint32
}
type VkSubmitInfo struct {
	sType VkStructureType
	pNext uintptr
	waitSemaphoreCount uint32
	pWaitSemaphores *VkSemaphore
	pWaitDstStageMask *VkPipelineStageFlags
	commandBufferCount uint32
	pCommandBuffers *VkCommandBuffer
	signalSemaphoreCount uint32
	pSignalSemaphores *VkSemaphore
}
type VkDisplayPropertiesKHR struct {
	display VkDisplayKHR
	displayName *byte
	physicalDimensions VkExtent2D
	physicalResolution VkExtent2D
	supportedTransforms VkSurfaceTransformFlagsKHR
	planeReorderPossible VkBool32
	persistentContent VkBool32
}
type VkDisplayPlanePropertiesKHR struct {
	currentDisplay VkDisplayKHR
	currentStackIndex uint32
}
type VkDisplayModeParametersKHR struct {
	visibleRegion VkExtent2D
	refreshRate uint32
}
type VkDisplayModePropertiesKHR struct {
	displayMode VkDisplayModeKHR
	parameters VkDisplayModeParametersKHR
}
type VkDisplayModeCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	flags VkDisplayModeCreateFlagsKHR
	parameters VkDisplayModeParametersKHR
}
type VkDisplayPlaneCapabilitiesKHR struct {
	supportedAlpha VkDisplayPlaneAlphaFlagsKHR
	minSrcPosition VkOffset2D
	maxSrcPosition VkOffset2D
	minSrcExtent VkExtent2D
	maxSrcExtent VkExtent2D
	minDstPosition VkOffset2D
	maxDstPosition VkOffset2D
	minDstExtent VkExtent2D
	maxDstExtent VkExtent2D
}
type VkDisplaySurfaceCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	flags VkDisplaySurfaceCreateFlagsKHR
	displayMode VkDisplayModeKHR
	planeIndex uint32
	planeStackIndex uint32
	transform VkSurfaceTransformFlagBitsKHR
	globalAlpha float32
	alphaMode VkDisplayPlaneAlphaFlagBitsKHR
	imageExtent VkExtent2D
}
type VkDisplayPresentInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	srcRect VkRect2D
	dstRect VkRect2D
	persistent VkBool32
}
type VkSurfaceCapabilitiesKHR struct {
	minImageCount uint32
	maxImageCount uint32
	currentExtent VkExtent2D
	minImageExtent VkExtent2D
	maxImageExtent VkExtent2D
	maxImageArrayLayers uint32
	supportedTransforms VkSurfaceTransformFlagsKHR
	currentTransform VkSurfaceTransformFlagBitsKHR
	supportedCompositeAlpha VkCompositeAlphaFlagsKHR
	supportedUsageFlags VkImageUsageFlags
}
type VkAndroidSurfaceCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	flags VkAndroidSurfaceCreateFlagsKHR
	window *ANativeWindow
}
type VkViSurfaceCreateInfoNN struct {
	sType VkStructureType
	pNext uintptr
	flags VkViSurfaceCreateFlagsNN
	window uintptr
}
type VkWaylandSurfaceCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	flags VkWaylandSurfaceCreateFlagsKHR
	display *wl_display
	surface *wl_surface
}
type VkWin32SurfaceCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	flags VkWin32SurfaceCreateFlagsKHR
	hinstance HINSTANCE
	hwnd HWND
}
type VkXlibSurfaceCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	flags VkXlibSurfaceCreateFlagsKHR
	dpy *Display
	window Window
}
type VkXcbSurfaceCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	flags VkXcbSurfaceCreateFlagsKHR
	connection *xcb_connection_t
	window xcb_window_t
}
type VkDirectFBSurfaceCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	flags VkDirectFBSurfaceCreateFlagsEXT
	dfb *IDirectFB
	surface *IDirectFBSurface
}
type VkImagePipeSurfaceCreateInfoFUCHSIA struct {
	sType VkStructureType
	pNext uintptr
	flags VkImagePipeSurfaceCreateFlagsFUCHSIA
	imagePipeHandle zx_handle_t
}
type VkStreamDescriptorSurfaceCreateInfoGGP struct {
	sType VkStructureType
	pNext uintptr
	flags VkStreamDescriptorSurfaceCreateFlagsGGP
	streamDescriptor GgpStreamDescriptor
}
type VkSurfaceFormatKHR struct {
	format VkFormat
	colorSpace VkColorSpaceKHR
}
type VkSwapchainCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	flags VkSwapchainCreateFlagsKHR
	surface VkSurfaceKHR
	minImageCount uint32
	imageFormat VkFormat
	imageColorSpace VkColorSpaceKHR
	imageExtent VkExtent2D
	imageArrayLayers uint32
	imageUsage VkImageUsageFlags
	imageSharingMode VkSharingMode
	queueFamilyIndexCount uint32
	pQueueFamilyIndices *uint32
	preTransform VkSurfaceTransformFlagBitsKHR
	compositeAlpha VkCompositeAlphaFlagBitsKHR
	presentMode VkPresentModeKHR
	clipped VkBool32
	oldSwapchain VkSwapchainKHR
}
type VkPresentInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	waitSemaphoreCount uint32
	pWaitSemaphores *VkSemaphore
	swapchainCount uint32
	pSwapchains *VkSwapchainKHR
	pImageIndices *uint32
	pResults *VkResult
}
type VkDebugReportCallbackCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	flags VkDebugReportFlagsEXT
	pfnCallback PFN_vkDebugReportCallbackEXT
	pUserData uintptr
}
type VkValidationFlagsEXT struct {
	sType VkStructureType
	pNext uintptr
	disabledValidationCheckCount uint32
	pDisabledValidationChecks *VkValidationCheckEXT
}
type VkValidationFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	enabledValidationFeatureCount uint32
	pEnabledValidationFeatures *VkValidationFeatureEnableEXT
	disabledValidationFeatureCount uint32
	pDisabledValidationFeatures *VkValidationFeatureDisableEXT
}
type VkPipelineRasterizationStateRasterizationOrderAMD struct {
	sType VkStructureType
	pNext uintptr
	rasterizationOrder VkRasterizationOrderAMD
}
type VkDebugMarkerObjectNameInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	objectType VkDebugReportObjectTypeEXT
	object uint64
	pObjectName *byte
}
type VkDebugMarkerObjectTagInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	objectType VkDebugReportObjectTypeEXT
	object uint64
	tagName uint64
	tagSize uint
	pTag uintptr
}
type VkDebugMarkerMarkerInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	pMarkerName *byte
	color float32
}
type VkDedicatedAllocationImageCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	dedicatedAllocation VkBool32
}
type VkDedicatedAllocationBufferCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	dedicatedAllocation VkBool32
}
type VkDedicatedAllocationMemoryAllocateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	image VkImage
	buffer VkBuffer
}
type VkExternalImageFormatPropertiesNV struct {
	imageFormatProperties VkImageFormatProperties
	externalMemoryFeatures VkExternalMemoryFeatureFlagsNV
	exportFromImportedHandleTypes VkExternalMemoryHandleTypeFlagsNV
	compatibleHandleTypes VkExternalMemoryHandleTypeFlagsNV
}
type VkExternalMemoryImageCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	handleTypes VkExternalMemoryHandleTypeFlagsNV
}
type VkExportMemoryAllocateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	handleTypes VkExternalMemoryHandleTypeFlagsNV
}
type VkImportMemoryWin32HandleInfoNV struct {
	sType VkStructureType
	pNext uintptr
	handleType VkExternalMemoryHandleTypeFlagsNV
	handle HANDLE
}
type VkExportMemoryWin32HandleInfoNV struct {
	sType VkStructureType
	pNext uintptr
	pAttributes *SECURITY_ATTRIBUTES
	dwAccess DWORD
}
type VkWin32KeyedMutexAcquireReleaseInfoNV struct {
	sType VkStructureType
	pNext uintptr
	acquireCount uint32
	pAcquireSyncs *VkDeviceMemory
	pAcquireKeys *uint64
	pAcquireTimeoutMilliseconds *uint32
	releaseCount uint32
	pReleaseSyncs *VkDeviceMemory
	pReleaseKeys *uint64
}
type VkPhysicalDeviceDeviceGeneratedCommandsFeaturesNV struct {
	sType VkStructureType
	pNext uintptr
	deviceGeneratedCommands VkBool32
}
type VkDevicePrivateDataCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	privateDataSlotRequestCount uint32
}
type VkPrivateDataSlotCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	flags VkPrivateDataSlotCreateFlagsEXT
}
type VkPhysicalDevicePrivateDataFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	privateData VkBool32
}
type VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV struct {
	sType VkStructureType
	pNext uintptr
	maxGraphicsShaderGroupCount uint32
	maxIndirectSequenceCount uint32
	maxIndirectCommandsTokenCount uint32
	maxIndirectCommandsStreamCount uint32
	maxIndirectCommandsTokenOffset uint32
	maxIndirectCommandsStreamStride uint32
	minSequencesCountBufferOffsetAlignment uint32
	minSequencesIndexBufferOffsetAlignment uint32
	minIndirectCommandsBufferOffsetAlignment uint32
}
type VkGraphicsShaderGroupCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	stageCount uint32
	pStages *VkPipelineShaderStageCreateInfo
	pVertexInputState *VkPipelineVertexInputStateCreateInfo
	pTessellationState *VkPipelineTessellationStateCreateInfo
}
type VkGraphicsPipelineShaderGroupsCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	groupCount uint32
	pGroups *VkGraphicsShaderGroupCreateInfoNV
	pipelineCount uint32
	pPipelines *VkPipeline
}
type VkBindShaderGroupIndirectCommandNV struct {
	groupIndex uint32
}
type VkBindIndexBufferIndirectCommandNV struct {
	bufferAddress VkDeviceAddress
	size uint32
	indexType VkIndexType
}
type VkBindVertexBufferIndirectCommandNV struct {
	bufferAddress VkDeviceAddress
	size uint32
	stride uint32
}
type VkSetStateFlagsIndirectCommandNV struct {
	data uint32
}
type VkIndirectCommandsStreamNV struct {
	buffer VkBuffer
	offset VkDeviceSize
}
type VkIndirectCommandsLayoutTokenNV struct {
	sType VkStructureType
	pNext uintptr
	tokenType VkIndirectCommandsTokenTypeNV
	stream uint32
	offset uint32
	vertexBindingUnit uint32
	vertexDynamicStride VkBool32
	pushconstantPipelineLayout VkPipelineLayout
	pushconstantShaderStageFlags VkShaderStageFlags
	pushconstantOffset uint32
	pushconstantSize uint32
	indirectStateFlags VkIndirectStateFlagsNV
	indexTypeCount uint32
	pIndexTypes *VkIndexType
	pIndexTypeValues *uint32
}
type VkIndirectCommandsLayoutCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	flags VkIndirectCommandsLayoutUsageFlagsNV
	pipelineBindPoint VkPipelineBindPoint
	tokenCount uint32
	pTokens *VkIndirectCommandsLayoutTokenNV
	streamCount uint32
	pStreamStrides *uint32
}
type VkGeneratedCommandsInfoNV struct {
	sType VkStructureType
	pNext uintptr
	pipelineBindPoint VkPipelineBindPoint
	pipeline VkPipeline
	indirectCommandsLayout VkIndirectCommandsLayoutNV
	streamCount uint32
	pStreams *VkIndirectCommandsStreamNV
	sequencesCount uint32
	preprocessBuffer VkBuffer
	preprocessOffset VkDeviceSize
	preprocessSize VkDeviceSize
	sequencesCountBuffer VkBuffer
	sequencesCountOffset VkDeviceSize
	sequencesIndexBuffer VkBuffer
	sequencesIndexOffset VkDeviceSize
}
type VkGeneratedCommandsMemoryRequirementsInfoNV struct {
	sType VkStructureType
	pNext uintptr
	pipelineBindPoint VkPipelineBindPoint
	pipeline VkPipeline
	indirectCommandsLayout VkIndirectCommandsLayoutNV
	maxSequencesCount uint32
}
type VkPhysicalDeviceFeatures2 struct {
	sType VkStructureType
	pNext uintptr
	features VkPhysicalDeviceFeatures
}
type VkPhysicalDeviceFeatures2KHR struct {
}
type VkPhysicalDeviceProperties2 struct {
	sType VkStructureType
	pNext uintptr
	properties VkPhysicalDeviceProperties
}
type VkPhysicalDeviceProperties2KHR struct {
}
type VkFormatProperties2 struct {
	sType VkStructureType
	pNext uintptr
	formatProperties VkFormatProperties
}
type VkFormatProperties2KHR struct {
}
type VkImageFormatProperties2 struct {
	sType VkStructureType
	pNext uintptr
	imageFormatProperties VkImageFormatProperties
}
type VkImageFormatProperties2KHR struct {
}
type VkPhysicalDeviceImageFormatInfo2 struct {
	sType VkStructureType
	pNext uintptr
	format VkFormat
	type0 VkImageType
	tiling VkImageTiling
	usage VkImageUsageFlags
	flags VkImageCreateFlags
}
type VkPhysicalDeviceImageFormatInfo2KHR struct {
}
type VkQueueFamilyProperties2 struct {
	sType VkStructureType
	pNext uintptr
	queueFamilyProperties VkQueueFamilyProperties
}
type VkQueueFamilyProperties2KHR struct {
}
type VkPhysicalDeviceMemoryProperties2 struct {
	sType VkStructureType
	pNext uintptr
	memoryProperties VkPhysicalDeviceMemoryProperties
}
type VkPhysicalDeviceMemoryProperties2KHR struct {
}
type VkSparseImageFormatProperties2 struct {
	sType VkStructureType
	pNext uintptr
	properties VkSparseImageFormatProperties
}
type VkSparseImageFormatProperties2KHR struct {
}
type VkPhysicalDeviceSparseImageFormatInfo2 struct {
	sType VkStructureType
	pNext uintptr
	format VkFormat
	type0 VkImageType
	samples VkSampleCountFlagBits
	usage VkImageUsageFlags
	tiling VkImageTiling
}
type VkPhysicalDeviceSparseImageFormatInfo2KHR struct {
}
type VkPhysicalDevicePushDescriptorPropertiesKHR struct {
	sType VkStructureType
	pNext uintptr
	maxPushDescriptors uint32
}
type VkConformanceVersion struct {
	major byte
	minor byte
	subminor byte
	patch byte
}
type VkConformanceVersionKHR struct {
}
type VkPhysicalDeviceDriverProperties struct {
	sType VkStructureType
	pNext uintptr
	driverID VkDriverId
	driverName byte
	driverInfo byte
	conformanceVersion VkConformanceVersion
}
type VkPhysicalDeviceDriverPropertiesKHR struct {
}
type VkPresentRegionsKHR struct {
	sType VkStructureType
	pNext uintptr
	swapchainCount uint32
	pRegions *VkPresentRegionKHR
}
type VkPresentRegionKHR struct {
	rectangleCount uint32
	pRectangles *VkRectLayerKHR
}
type VkRectLayerKHR struct {
	offset VkOffset2D
	extent VkExtent2D
	layer uint32
}
type VkPhysicalDeviceVariablePointersFeatures struct {
	sType VkStructureType
	pNext uintptr
	variablePointersStorageBuffer VkBool32
	variablePointers VkBool32
}
type VkPhysicalDeviceVariablePointersFeaturesKHR struct {
}
type VkPhysicalDeviceVariablePointerFeaturesKHR struct {
}
type VkPhysicalDeviceVariablePointerFeatures struct {
}
type VkExternalMemoryProperties struct {
	externalMemoryFeatures VkExternalMemoryFeatureFlags
	exportFromImportedHandleTypes VkExternalMemoryHandleTypeFlags
	compatibleHandleTypes VkExternalMemoryHandleTypeFlags
}
type VkExternalMemoryPropertiesKHR struct {
}
type VkPhysicalDeviceExternalImageFormatInfo struct {
	sType VkStructureType
	pNext uintptr
	handleType VkExternalMemoryHandleTypeFlagBits
}
type VkPhysicalDeviceExternalImageFormatInfoKHR struct {
}
type VkExternalImageFormatProperties struct {
	sType VkStructureType
	pNext uintptr
	externalMemoryProperties VkExternalMemoryProperties
}
type VkExternalImageFormatPropertiesKHR struct {
}
type VkPhysicalDeviceExternalBufferInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkBufferCreateFlags
	usage VkBufferUsageFlags
	handleType VkExternalMemoryHandleTypeFlagBits
}
type VkPhysicalDeviceExternalBufferInfoKHR struct {
}
type VkExternalBufferProperties struct {
	sType VkStructureType
	pNext uintptr
	externalMemoryProperties VkExternalMemoryProperties
}
type VkExternalBufferPropertiesKHR struct {
}
type VkPhysicalDeviceIDProperties struct {
	sType VkStructureType
	pNext uintptr
	deviceUUID byte
	driverUUID byte
	deviceLUID byte
	deviceNodeMask uint32
	deviceLUIDValid VkBool32
}
type VkPhysicalDeviceIDPropertiesKHR struct {
}
type VkExternalMemoryImageCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	handleTypes VkExternalMemoryHandleTypeFlags
}
type VkExternalMemoryImageCreateInfoKHR struct {
}
type VkExternalMemoryBufferCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	handleTypes VkExternalMemoryHandleTypeFlags
}
type VkExternalMemoryBufferCreateInfoKHR struct {
}
type VkExportMemoryAllocateInfo struct {
	sType VkStructureType
	pNext uintptr
	handleTypes VkExternalMemoryHandleTypeFlags
}
type VkExportMemoryAllocateInfoKHR struct {
}
type VkImportMemoryWin32HandleInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	handleType VkExternalMemoryHandleTypeFlagBits
	handle HANDLE
	name LPCWSTR
}
type VkExportMemoryWin32HandleInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	pAttributes *SECURITY_ATTRIBUTES
	dwAccess DWORD
	name LPCWSTR
}
type VkMemoryWin32HandlePropertiesKHR struct {
	sType VkStructureType
	pNext uintptr
	memoryTypeBits uint32
}
type VkMemoryGetWin32HandleInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	memory VkDeviceMemory
	handleType VkExternalMemoryHandleTypeFlagBits
}
type VkImportMemoryFdInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	handleType VkExternalMemoryHandleTypeFlagBits
	fd int
}
type VkMemoryFdPropertiesKHR struct {
	sType VkStructureType
	pNext uintptr
	memoryTypeBits uint32
}
type VkMemoryGetFdInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	memory VkDeviceMemory
	handleType VkExternalMemoryHandleTypeFlagBits
}
type VkWin32KeyedMutexAcquireReleaseInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	acquireCount uint32
	pAcquireSyncs *VkDeviceMemory
	pAcquireKeys *uint64
	pAcquireTimeouts *uint32
	releaseCount uint32
	pReleaseSyncs *VkDeviceMemory
	pReleaseKeys *uint64
}
type VkPhysicalDeviceExternalSemaphoreInfo struct {
	sType VkStructureType
	pNext uintptr
	handleType VkExternalSemaphoreHandleTypeFlagBits
}
type VkPhysicalDeviceExternalSemaphoreInfoKHR struct {
}
type VkExternalSemaphoreProperties struct {
	sType VkStructureType
	pNext uintptr
	exportFromImportedHandleTypes VkExternalSemaphoreHandleTypeFlags
	compatibleHandleTypes VkExternalSemaphoreHandleTypeFlags
	externalSemaphoreFeatures VkExternalSemaphoreFeatureFlags
}
type VkExternalSemaphorePropertiesKHR struct {
}
type VkExportSemaphoreCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	handleTypes VkExternalSemaphoreHandleTypeFlags
}
type VkExportSemaphoreCreateInfoKHR struct {
}
type VkImportSemaphoreWin32HandleInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	semaphore VkSemaphore
	flags VkSemaphoreImportFlags
	handleType VkExternalSemaphoreHandleTypeFlagBits
	handle HANDLE
	name LPCWSTR
}
type VkExportSemaphoreWin32HandleInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	pAttributes *SECURITY_ATTRIBUTES
	dwAccess DWORD
	name LPCWSTR
}
type VkD3D12FenceSubmitInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	waitSemaphoreValuesCount uint32
	pWaitSemaphoreValues *uint64
	signalSemaphoreValuesCount uint32
	pSignalSemaphoreValues *uint64
}
type VkSemaphoreGetWin32HandleInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	semaphore VkSemaphore
	handleType VkExternalSemaphoreHandleTypeFlagBits
}
type VkImportSemaphoreFdInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	semaphore VkSemaphore
	flags VkSemaphoreImportFlags
	handleType VkExternalSemaphoreHandleTypeFlagBits
	fd int
}
type VkSemaphoreGetFdInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	semaphore VkSemaphore
	handleType VkExternalSemaphoreHandleTypeFlagBits
}
type VkPhysicalDeviceExternalFenceInfo struct {
	sType VkStructureType
	pNext uintptr
	handleType VkExternalFenceHandleTypeFlagBits
}
type VkPhysicalDeviceExternalFenceInfoKHR struct {
}
type VkExternalFenceProperties struct {
	sType VkStructureType
	pNext uintptr
	exportFromImportedHandleTypes VkExternalFenceHandleTypeFlags
	compatibleHandleTypes VkExternalFenceHandleTypeFlags
	externalFenceFeatures VkExternalFenceFeatureFlags
}
type VkExternalFencePropertiesKHR struct {
}
type VkExportFenceCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	handleTypes VkExternalFenceHandleTypeFlags
}
type VkExportFenceCreateInfoKHR struct {
}
type VkImportFenceWin32HandleInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	fence VkFence
	flags VkFenceImportFlags
	handleType VkExternalFenceHandleTypeFlagBits
	handle HANDLE
	name LPCWSTR
}
type VkExportFenceWin32HandleInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	pAttributes *SECURITY_ATTRIBUTES
	dwAccess DWORD
	name LPCWSTR
}
type VkFenceGetWin32HandleInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	fence VkFence
	handleType VkExternalFenceHandleTypeFlagBits
}
type VkImportFenceFdInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	fence VkFence
	flags VkFenceImportFlags
	handleType VkExternalFenceHandleTypeFlagBits
	fd int
}
type VkFenceGetFdInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	fence VkFence
	handleType VkExternalFenceHandleTypeFlagBits
}
type VkPhysicalDeviceMultiviewFeatures struct {
	sType VkStructureType
	pNext uintptr
	multiview VkBool32
	multiviewGeometryShader VkBool32
	multiviewTessellationShader VkBool32
}
type VkPhysicalDeviceMultiviewFeaturesKHR struct {
}
type VkPhysicalDeviceMultiviewProperties struct {
	sType VkStructureType
	pNext uintptr
	maxMultiviewViewCount uint32
	maxMultiviewInstanceIndex uint32
}
type VkPhysicalDeviceMultiviewPropertiesKHR struct {
}
type VkRenderPassMultiviewCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	subpassCount uint32
	pViewMasks *uint32
	dependencyCount uint32
	pViewOffsets *int32
	correlationMaskCount uint32
	pCorrelationMasks *uint32
}
type VkRenderPassMultiviewCreateInfoKHR struct {
}
type VkSurfaceCapabilities2EXT struct {
	sType VkStructureType
	pNext uintptr
	minImageCount uint32
	maxImageCount uint32
	currentExtent VkExtent2D
	minImageExtent VkExtent2D
	maxImageExtent VkExtent2D
	maxImageArrayLayers uint32
	supportedTransforms VkSurfaceTransformFlagsKHR
	currentTransform VkSurfaceTransformFlagBitsKHR
	supportedCompositeAlpha VkCompositeAlphaFlagsKHR
	supportedUsageFlags VkImageUsageFlags
	supportedSurfaceCounters VkSurfaceCounterFlagsEXT
}
type VkDisplayPowerInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	powerState VkDisplayPowerStateEXT
}
type VkDeviceEventInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	deviceEvent VkDeviceEventTypeEXT
}
type VkDisplayEventInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	displayEvent VkDisplayEventTypeEXT
}
type VkSwapchainCounterCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	surfaceCounters VkSurfaceCounterFlagsEXT
}
type VkPhysicalDeviceGroupProperties struct {
	sType VkStructureType
	pNext uintptr
	physicalDeviceCount uint32
	physicalDevices VkPhysicalDevice
	subsetAllocation VkBool32
}
type VkPhysicalDeviceGroupPropertiesKHR struct {
}
type VkMemoryAllocateFlagsInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkMemoryAllocateFlags
	deviceMask uint32
}
type VkMemoryAllocateFlagsInfoKHR struct {
}
type VkBindBufferMemoryInfo struct {
	sType VkStructureType
	pNext uintptr
	buffer VkBuffer
	memory VkDeviceMemory
	memoryOffset VkDeviceSize
}
type VkBindBufferMemoryInfoKHR struct {
}
type VkBindBufferMemoryDeviceGroupInfo struct {
	sType VkStructureType
	pNext uintptr
	deviceIndexCount uint32
	pDeviceIndices *uint32
}
type VkBindBufferMemoryDeviceGroupInfoKHR struct {
}
type VkBindImageMemoryInfo struct {
	sType VkStructureType
	pNext uintptr
	image VkImage
	memory VkDeviceMemory
	memoryOffset VkDeviceSize
}
type VkBindImageMemoryInfoKHR struct {
}
type VkBindImageMemoryDeviceGroupInfo struct {
	sType VkStructureType
	pNext uintptr
	deviceIndexCount uint32
	pDeviceIndices *uint32
	splitInstanceBindRegionCount uint32
	pSplitInstanceBindRegions *VkRect2D
}
type VkBindImageMemoryDeviceGroupInfoKHR struct {
}
type VkDeviceGroupRenderPassBeginInfo struct {
	sType VkStructureType
	pNext uintptr
	deviceMask uint32
	deviceRenderAreaCount uint32
	pDeviceRenderAreas *VkRect2D
}
type VkDeviceGroupRenderPassBeginInfoKHR struct {
}
type VkDeviceGroupCommandBufferBeginInfo struct {
	sType VkStructureType
	pNext uintptr
	deviceMask uint32
}
type VkDeviceGroupCommandBufferBeginInfoKHR struct {
}
type VkDeviceGroupSubmitInfo struct {
	sType VkStructureType
	pNext uintptr
	waitSemaphoreCount uint32
	pWaitSemaphoreDeviceIndices *uint32
	commandBufferCount uint32
	pCommandBufferDeviceMasks *uint32
	signalSemaphoreCount uint32
	pSignalSemaphoreDeviceIndices *uint32
}
type VkDeviceGroupSubmitInfoKHR struct {
}
type VkDeviceGroupBindSparseInfo struct {
	sType VkStructureType
	pNext uintptr
	resourceDeviceIndex uint32
	memoryDeviceIndex uint32
}
type VkDeviceGroupBindSparseInfoKHR struct {
}
type VkDeviceGroupPresentCapabilitiesKHR struct {
	sType VkStructureType
	pNext uintptr
	presentMask uint32
	modes VkDeviceGroupPresentModeFlagsKHR
}
type VkImageSwapchainCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	swapchain VkSwapchainKHR
}
type VkBindImageMemorySwapchainInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	swapchain VkSwapchainKHR
	imageIndex uint32
}
type VkAcquireNextImageInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	swapchain VkSwapchainKHR
	timeout uint64
	semaphore VkSemaphore
	fence VkFence
	deviceMask uint32
}
type VkDeviceGroupPresentInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	swapchainCount uint32
	pDeviceMasks *uint32
	mode VkDeviceGroupPresentModeFlagBitsKHR
}
type VkDeviceGroupDeviceCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	physicalDeviceCount uint32
	pPhysicalDevices *VkPhysicalDevice
}
type VkDeviceGroupDeviceCreateInfoKHR struct {
}
type VkDeviceGroupSwapchainCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	modes VkDeviceGroupPresentModeFlagsKHR
}
type VkDescriptorUpdateTemplateEntry struct {
	dstBinding uint32
	dstArrayElement uint32
	descriptorCount uint32
	descriptorType VkDescriptorType
	offset uint
	stride uint
}
type VkDescriptorUpdateTemplateEntryKHR struct {
}
type VkDescriptorUpdateTemplateCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkDescriptorUpdateTemplateCreateFlags
	descriptorUpdateEntryCount uint32
	pDescriptorUpdateEntries *VkDescriptorUpdateTemplateEntry
	templateType VkDescriptorUpdateTemplateType
	descriptorSetLayout VkDescriptorSetLayout
	pipelineBindPoint VkPipelineBindPoint
	pipelineLayout VkPipelineLayout
	set uint32
}
type VkDescriptorUpdateTemplateCreateInfoKHR struct {
}
type VkXYColorEXT struct {
	x float32
	y float32
}
type VkHdrMetadataEXT struct {
	sType VkStructureType
	pNext uintptr
	displayPrimaryRed VkXYColorEXT
	displayPrimaryGreen VkXYColorEXT
	displayPrimaryBlue VkXYColorEXT
	whitePoint VkXYColorEXT
	maxLuminance float32
	minLuminance float32
	maxContentLightLevel float32
	maxFrameAverageLightLevel float32
}
type VkDisplayNativeHdrSurfaceCapabilitiesAMD struct {
	sType VkStructureType
	pNext uintptr
	localDimmingSupport VkBool32
}
type VkSwapchainDisplayNativeHdrCreateInfoAMD struct {
	sType VkStructureType
	pNext uintptr
	localDimmingEnable VkBool32
}
type VkRefreshCycleDurationGOOGLE struct {
	refreshDuration uint64
}
type VkPastPresentationTimingGOOGLE struct {
	presentID uint32
	desiredPresentTime uint64
	actualPresentTime uint64
	earliestPresentTime uint64
	presentMargin uint64
}
type VkPresentTimesInfoGOOGLE struct {
	sType VkStructureType
	pNext uintptr
	swapchainCount uint32
	pTimes *VkPresentTimeGOOGLE
}
type VkPresentTimeGOOGLE struct {
	presentID uint32
	desiredPresentTime uint64
}
type VkIOSSurfaceCreateInfoMVK struct {
	sType VkStructureType
	pNext uintptr
	flags VkIOSSurfaceCreateFlagsMVK
	pView uintptr
}
type VkMacOSSurfaceCreateInfoMVK struct {
	sType VkStructureType
	pNext uintptr
	flags VkMacOSSurfaceCreateFlagsMVK
	pView uintptr
}
type VkMetalSurfaceCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	flags VkMetalSurfaceCreateFlagsEXT
	pLayer *CAMetalLayer
}
type VkViewportWScalingNV struct {
	xcoeff float32
	ycoeff float32
}
type VkPipelineViewportWScalingStateCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	viewportWScalingEnable VkBool32
	viewportCount uint32
	pViewportWScalings *VkViewportWScalingNV
}
type VkViewportSwizzleNV struct {
	x VkViewportCoordinateSwizzleNV
	y VkViewportCoordinateSwizzleNV
	z VkViewportCoordinateSwizzleNV
	w VkViewportCoordinateSwizzleNV
}
type VkPipelineViewportSwizzleStateCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineViewportSwizzleStateCreateFlagsNV
	viewportCount uint32
	pViewportSwizzles *VkViewportSwizzleNV
}
type VkPhysicalDeviceDiscardRectanglePropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	maxDiscardRectangles uint32
}
type VkPipelineDiscardRectangleStateCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineDiscardRectangleStateCreateFlagsEXT
	discardRectangleMode VkDiscardRectangleModeEXT
	discardRectangleCount uint32
	pDiscardRectangles *VkRect2D
}
type VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX struct {
	sType VkStructureType
	pNext uintptr
	perViewPositionAllComponents VkBool32
}
type VkInputAttachmentAspectReference struct {
	subpass uint32
	inputAttachmentIndex uint32
	aspectMask VkImageAspectFlags
}
type VkInputAttachmentAspectReferenceKHR struct {
}
type VkRenderPassInputAttachmentAspectCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	aspectReferenceCount uint32
	pAspectReferences *VkInputAttachmentAspectReference
}
type VkRenderPassInputAttachmentAspectCreateInfoKHR struct {
}
type VkPhysicalDeviceSurfaceInfo2KHR struct {
	sType VkStructureType
	pNext uintptr
	surface VkSurfaceKHR
}
type VkSurfaceCapabilities2KHR struct {
	sType VkStructureType
	pNext uintptr
	surfaceCapabilities VkSurfaceCapabilitiesKHR
}
type VkSurfaceFormat2KHR struct {
	sType VkStructureType
	pNext uintptr
	surfaceFormat VkSurfaceFormatKHR
}
type VkDisplayProperties2KHR struct {
	sType VkStructureType
	pNext uintptr
	displayProperties VkDisplayPropertiesKHR
}
type VkDisplayPlaneProperties2KHR struct {
	sType VkStructureType
	pNext uintptr
	displayPlaneProperties VkDisplayPlanePropertiesKHR
}
type VkDisplayModeProperties2KHR struct {
	sType VkStructureType
	pNext uintptr
	displayModeProperties VkDisplayModePropertiesKHR
}
type VkDisplayPlaneInfo2KHR struct {
	sType VkStructureType
	pNext uintptr
	mode VkDisplayModeKHR
	planeIndex uint32
}
type VkDisplayPlaneCapabilities2KHR struct {
	sType VkStructureType
	pNext uintptr
	capabilities VkDisplayPlaneCapabilitiesKHR
}
type VkSharedPresentSurfaceCapabilitiesKHR struct {
	sType VkStructureType
	pNext uintptr
	sharedPresentSupportedUsageFlags VkImageUsageFlags
}
type VkPhysicalDevice16BitStorageFeatures struct {
	sType VkStructureType
	pNext uintptr
	storageBuffer16BitAccess VkBool32
	uniformAndStorageBuffer16BitAccess VkBool32
	storagePushConstant16 VkBool32
	storageInputOutput16 VkBool32
}
type VkPhysicalDevice16BitStorageFeaturesKHR struct {
}
type VkPhysicalDeviceSubgroupProperties struct {
	sType VkStructureType
	pNext uintptr
	subgroupSize uint32
	supportedStages VkShaderStageFlags
	supportedOperations VkSubgroupFeatureFlags
	quadOperationsInAllStages VkBool32
}
type VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures struct {
	sType VkStructureType
	pNext uintptr
	shaderSubgroupExtendedTypes VkBool32
}
type VkPhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR struct {
}
type VkBufferMemoryRequirementsInfo2 struct {
	sType VkStructureType
	pNext uintptr
	buffer VkBuffer
}
type VkBufferMemoryRequirementsInfo2KHR struct {
}
type VkImageMemoryRequirementsInfo2 struct {
	sType VkStructureType
	pNext uintptr
	image VkImage
}
type VkImageMemoryRequirementsInfo2KHR struct {
}
type VkImageSparseMemoryRequirementsInfo2 struct {
	sType VkStructureType
	pNext uintptr
	image VkImage
}
type VkImageSparseMemoryRequirementsInfo2KHR struct {
}
type VkMemoryRequirements2 struct {
	sType VkStructureType
	pNext uintptr
	memoryRequirements VkMemoryRequirements
}
type VkMemoryRequirements2KHR struct {
}
type VkSparseImageMemoryRequirements2 struct {
	sType VkStructureType
	pNext uintptr
	memoryRequirements VkSparseImageMemoryRequirements
}
type VkSparseImageMemoryRequirements2KHR struct {
}
type VkPhysicalDevicePointClippingProperties struct {
	sType VkStructureType
	pNext uintptr
	pointClippingBehavior VkPointClippingBehavior
}
type VkPhysicalDevicePointClippingPropertiesKHR struct {
}
type VkMemoryDedicatedRequirements struct {
	sType VkStructureType
	pNext uintptr
	prefersDedicatedAllocation VkBool32
	requiresDedicatedAllocation VkBool32
}
type VkMemoryDedicatedRequirementsKHR struct {
}
type VkMemoryDedicatedAllocateInfo struct {
	sType VkStructureType
	pNext uintptr
	image VkImage
	buffer VkBuffer
}
type VkMemoryDedicatedAllocateInfoKHR struct {
}
type VkImageViewUsageCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	usage VkImageUsageFlags
}
type VkImageViewUsageCreateInfoKHR struct {
}
type VkPipelineTessellationDomainOriginStateCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	domainOrigin VkTessellationDomainOrigin
}
type VkPipelineTessellationDomainOriginStateCreateInfoKHR struct {
}
type VkSamplerYcbcrConversionInfo struct {
	sType VkStructureType
	pNext uintptr
	conversion VkSamplerYcbcrConversion
}
type VkSamplerYcbcrConversionInfoKHR struct {
}
type VkSamplerYcbcrConversionCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	format VkFormat
	ycbcrModel VkSamplerYcbcrModelConversion
	ycbcrRange VkSamplerYcbcrRange
	components VkComponentMapping
	xChromaOffset VkChromaLocation
	yChromaOffset VkChromaLocation
	chromaFilter VkFilter
	forceExplicitReconstruction VkBool32
}
type VkSamplerYcbcrConversionCreateInfoKHR struct {
}
type VkBindImagePlaneMemoryInfo struct {
	sType VkStructureType
	pNext uintptr
	planeAspect VkImageAspectFlagBits
}
type VkBindImagePlaneMemoryInfoKHR struct {
}
type VkImagePlaneMemoryRequirementsInfo struct {
	sType VkStructureType
	pNext uintptr
	planeAspect VkImageAspectFlagBits
}
type VkImagePlaneMemoryRequirementsInfoKHR struct {
}
type VkPhysicalDeviceSamplerYcbcrConversionFeatures struct {
	sType VkStructureType
	pNext uintptr
	samplerYcbcrConversion VkBool32
}
type VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR struct {
}
type VkSamplerYcbcrConversionImageFormatProperties struct {
	sType VkStructureType
	pNext uintptr
	combinedImageSamplerDescriptorCount uint32
}
type VkSamplerYcbcrConversionImageFormatPropertiesKHR struct {
}
type VkTextureLODGatherFormatPropertiesAMD struct {
	sType VkStructureType
	pNext uintptr
	supportsTextureGatherLODBiasAMD VkBool32
}
type VkConditionalRenderingBeginInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	buffer VkBuffer
	offset VkDeviceSize
	flags VkConditionalRenderingFlagsEXT
}
type VkProtectedSubmitInfo struct {
	sType VkStructureType
	pNext uintptr
	protectedSubmit VkBool32
}
type VkPhysicalDeviceProtectedMemoryFeatures struct {
	sType VkStructureType
	pNext uintptr
	protectedMemory VkBool32
}
type VkPhysicalDeviceProtectedMemoryProperties struct {
	sType VkStructureType
	pNext uintptr
	protectedNoFault VkBool32
}
type VkDeviceQueueInfo2 struct {
	sType VkStructureType
	pNext uintptr
	flags VkDeviceQueueCreateFlags
	queueFamilyIndex uint32
	queueIndex uint32
}
type VkPipelineCoverageToColorStateCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineCoverageToColorStateCreateFlagsNV
	coverageToColorEnable VkBool32
	coverageToColorLocation uint32
}
type VkPhysicalDeviceSamplerFilterMinmaxProperties struct {
	sType VkStructureType
	pNext uintptr
	filterMinmaxSingleComponentFormats VkBool32
	filterMinmaxImageComponentMapping VkBool32
}
type VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT struct {
}
type VkSampleLocationEXT struct {
	x float32
	y float32
}
type VkSampleLocationsInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	sampleLocationsPerPixel VkSampleCountFlagBits
	sampleLocationGridSize VkExtent2D
	sampleLocationsCount uint32
	pSampleLocations *VkSampleLocationEXT
}
type VkAttachmentSampleLocationsEXT struct {
	attachmentIndex uint32
	sampleLocationsInfo VkSampleLocationsInfoEXT
}
type VkSubpassSampleLocationsEXT struct {
	subpassIndex uint32
	sampleLocationsInfo VkSampleLocationsInfoEXT
}
type VkRenderPassSampleLocationsBeginInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	attachmentInitialSampleLocationsCount uint32
	pAttachmentInitialSampleLocations *VkAttachmentSampleLocationsEXT
	postSubpassSampleLocationsCount uint32
	pPostSubpassSampleLocations *VkSubpassSampleLocationsEXT
}
type VkPipelineSampleLocationsStateCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	sampleLocationsEnable VkBool32
	sampleLocationsInfo VkSampleLocationsInfoEXT
}
type VkPhysicalDeviceSampleLocationsPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	sampleLocationSampleCounts VkSampleCountFlags
	maxSampleLocationGridSize VkExtent2D
	sampleLocationCoordinateRange float32
	sampleLocationSubPixelBits uint32
	variableSampleLocations VkBool32
}
type VkMultisamplePropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	maxSampleLocationGridSize VkExtent2D
}
type VkSamplerReductionModeCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	reductionMode VkSamplerReductionMode
}
type VkSamplerReductionModeCreateInfoEXT struct {
}
type VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	advancedBlendCoherentOperations VkBool32
}
type VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	advancedBlendMaxColorAttachments uint32
	advancedBlendIndependentBlend VkBool32
	advancedBlendNonPremultipliedSrcColor VkBool32
	advancedBlendNonPremultipliedDstColor VkBool32
	advancedBlendCorrelatedOverlap VkBool32
	advancedBlendAllOperations VkBool32
}
type VkPipelineColorBlendAdvancedStateCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	srcPremultiplied VkBool32
	dstPremultiplied VkBool32
	blendOverlap VkBlendOverlapEXT
}
type VkPhysicalDeviceInlineUniformBlockFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	inlineUniformBlock VkBool32
	descriptorBindingInlineUniformBlockUpdateAfterBind VkBool32
}
type VkPhysicalDeviceInlineUniformBlockPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	maxInlineUniformBlockSize uint32
	maxPerStageDescriptorInlineUniformBlocks uint32
	maxPerStageDescriptorUpdateAfterBindInlineUniformBlocks uint32
	maxDescriptorSetInlineUniformBlocks uint32
	maxDescriptorSetUpdateAfterBindInlineUniformBlocks uint32
}
type VkWriteDescriptorSetInlineUniformBlockEXT struct {
	sType VkStructureType
	pNext uintptr
	dataSize uint32
	pData uintptr
}
type VkDescriptorPoolInlineUniformBlockCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	maxInlineUniformBlockBindings uint32
}
type VkPipelineCoverageModulationStateCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineCoverageModulationStateCreateFlagsNV
	coverageModulationMode VkCoverageModulationModeNV
	coverageModulationTableEnable VkBool32
	coverageModulationTableCount uint32
	pCoverageModulationTable *float32
}
type VkImageFormatListCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	viewFormatCount uint32
	pViewFormats *VkFormat
}
type VkImageFormatListCreateInfoKHR struct {
}
type VkValidationCacheCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	flags VkValidationCacheCreateFlagsEXT
	initialDataSize uint
	pInitialData uintptr
}
type VkShaderModuleValidationCacheCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	validationCache VkValidationCacheEXT
}
type VkPhysicalDeviceMaintenance3Properties struct {
	sType VkStructureType
	pNext uintptr
	maxPerSetDescriptors uint32
	maxMemoryAllocationSize VkDeviceSize
}
type VkPhysicalDeviceMaintenance3PropertiesKHR struct {
}
type VkDescriptorSetLayoutSupport struct {
	sType VkStructureType
	pNext uintptr
	supported VkBool32
}
type VkDescriptorSetLayoutSupportKHR struct {
}
type VkPhysicalDeviceShaderDrawParametersFeatures struct {
	sType VkStructureType
	pNext uintptr
	shaderDrawParameters VkBool32
}
type VkPhysicalDeviceShaderDrawParameterFeatures struct {
}
type VkPhysicalDeviceShaderFloat16Int8Features struct {
	sType VkStructureType
	pNext uintptr
	shaderFloat16 VkBool32
	shaderInt8 VkBool32
}
type VkPhysicalDeviceShaderFloat16Int8FeaturesKHR struct {
}
type VkPhysicalDeviceFloat16Int8FeaturesKHR struct {
}
type VkPhysicalDeviceFloatControlsProperties struct {
	sType VkStructureType
	pNext uintptr
	denormBehaviorIndependence VkShaderFloatControlsIndependence
	roundingModeIndependence VkShaderFloatControlsIndependence
	shaderSignedZeroInfNanPreserveFloat16 VkBool32
	shaderSignedZeroInfNanPreserveFloat32 VkBool32
	shaderSignedZeroInfNanPreserveFloat64 VkBool32
	shaderDenormPreserveFloat16 VkBool32
	shaderDenormPreserveFloat32 VkBool32
	shaderDenormPreserveFloat64 VkBool32
	shaderDenormFlushToZeroFloat16 VkBool32
	shaderDenormFlushToZeroFloat32 VkBool32
	shaderDenormFlushToZeroFloat64 VkBool32
	shaderRoundingModeRTEFloat16 VkBool32
	shaderRoundingModeRTEFloat32 VkBool32
	shaderRoundingModeRTEFloat64 VkBool32
	shaderRoundingModeRTZFloat16 VkBool32
	shaderRoundingModeRTZFloat32 VkBool32
	shaderRoundingModeRTZFloat64 VkBool32
}
type VkPhysicalDeviceFloatControlsPropertiesKHR struct {
}
type VkPhysicalDeviceHostQueryResetFeatures struct {
	sType VkStructureType
	pNext uintptr
	hostQueryReset VkBool32
}
type VkPhysicalDeviceHostQueryResetFeaturesEXT struct {
}
type VkNativeBufferUsage2ANDROID struct {
	consumer uint64
	producer uint64
}
type VkNativeBufferANDROID struct {
	sType VkStructureType
	pNext uintptr
	handle uintptr
	stride int
	format int
	usage int
	usage2 VkNativeBufferUsage2ANDROID
}
type VkSwapchainImageCreateInfoANDROID struct {
	sType VkStructureType
	pNext uintptr
	usage VkSwapchainImageUsageFlagsANDROID
}
type VkPhysicalDevicePresentationPropertiesANDROID struct {
	sType VkStructureType
	pNext uintptr
	sharedImage VkBool32
}
type VkShaderResourceUsageAMD struct {
	numUsedVgprs uint32
	numUsedSgprs uint32
	ldsSizePerLocalWorkGroup uint32
	ldsUsageSizeInBytes uint
	scratchMemUsageInBytes uint
}
type VkShaderStatisticsInfoAMD struct {
	shaderStageMask VkShaderStageFlags
	resourceUsage VkShaderResourceUsageAMD
	numPhysicalVgprs uint32
	numPhysicalSgprs uint32
	numAvailableVgprs uint32
	numAvailableSgprs uint32
	computeWorkGroupSize uint32
}
type VkDeviceQueueGlobalPriorityCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	globalPriority VkQueueGlobalPriorityEXT
}
type VkDebugUtilsObjectNameInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	objectType VkObjectType
	objectHandle uint64
	pObjectName *byte
}
type VkDebugUtilsObjectTagInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	objectType VkObjectType
	objectHandle uint64
	tagName uint64
	tagSize uint
	pTag uintptr
}
type VkDebugUtilsLabelEXT struct {
	sType VkStructureType
	pNext uintptr
	pLabelName *byte
	color float32
}
type VkDebugUtilsMessengerCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	flags VkDebugUtilsMessengerCreateFlagsEXT
	messageSeverity VkDebugUtilsMessageSeverityFlagsEXT
	messageType VkDebugUtilsMessageTypeFlagsEXT
	pfnUserCallback PFN_vkDebugUtilsMessengerCallbackEXT
	pUserData uintptr
}
type VkDebugUtilsMessengerCallbackDataEXT struct {
	sType VkStructureType
	pNext uintptr
	flags VkDebugUtilsMessengerCallbackDataFlagsEXT
	pMessageIdName *byte
	messageIdNumber int32
	pMessage *byte
	queueLabelCount uint32
	pQueueLabels *VkDebugUtilsLabelEXT
	cmdBufLabelCount uint32
	pCmdBufLabels *VkDebugUtilsLabelEXT
	objectCount uint32
	pObjects *VkDebugUtilsObjectNameInfoEXT
}
type VkPhysicalDeviceDeviceMemoryReportFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	deviceMemoryReport VkBool32
}
type VkDeviceDeviceMemoryReportCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	flags VkDeviceMemoryReportFlagsEXT
	pfnUserCallback PFN_vkDeviceMemoryReportCallbackEXT
	pUserData uintptr
}
type VkDeviceMemoryReportCallbackDataEXT struct {
	sType VkStructureType
	pNext uintptr
	flags VkDeviceMemoryReportFlagsEXT
	type0 VkDeviceMemoryReportEventTypeEXT
	memoryObjectId uint64
	size VkDeviceSize
	objectType VkObjectType
	objectHandle uint64
	heapIndex uint32
}
type VkImportMemoryHostPointerInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	handleType VkExternalMemoryHandleTypeFlagBits
	pHostPointer uintptr
}
type VkMemoryHostPointerPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	memoryTypeBits uint32
}
type VkPhysicalDeviceExternalMemoryHostPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	minImportedHostPointerAlignment VkDeviceSize
}
type VkPhysicalDeviceConservativeRasterizationPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	primitiveOverestimationSize float32
	maxExtraPrimitiveOverestimationSize float32
	extraPrimitiveOverestimationSizeGranularity float32
	primitiveUnderestimation VkBool32
	conservativePointAndLineRasterization VkBool32
	degenerateTrianglesRasterized VkBool32
	degenerateLinesRasterized VkBool32
	fullyCoveredFragmentShaderInputVariable VkBool32
	conservativeRasterizationPostDepthCoverage VkBool32
}
type VkCalibratedTimestampInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	timeDomain VkTimeDomainEXT
}
type VkPhysicalDeviceShaderCorePropertiesAMD struct {
	sType VkStructureType
	pNext uintptr
	shaderEngineCount uint32
	shaderArraysPerEngineCount uint32
	computeUnitsPerShaderArray uint32
	simdPerComputeUnit uint32
	wavefrontsPerSimd uint32
	wavefrontSize uint32
	sgprsPerSimd uint32
	minSgprAllocation uint32
	maxSgprAllocation uint32
	sgprAllocationGranularity uint32
	vgprsPerSimd uint32
	minVgprAllocation uint32
	maxVgprAllocation uint32
	vgprAllocationGranularity uint32
}
type VkPhysicalDeviceShaderCoreProperties2AMD struct {
	sType VkStructureType
	pNext uintptr
	shaderCoreFeatures VkShaderCorePropertiesFlagsAMD
	activeComputeUnitCount uint32
}
type VkPipelineRasterizationConservativeStateCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineRasterizationConservativeStateCreateFlagsEXT
	conservativeRasterizationMode VkConservativeRasterizationModeEXT
	extraPrimitiveOverestimationSize float32
}
type VkPhysicalDeviceDescriptorIndexingFeatures struct {
	sType VkStructureType
	pNext uintptr
	shaderInputAttachmentArrayDynamicIndexing VkBool32
	shaderUniformTexelBufferArrayDynamicIndexing VkBool32
	shaderStorageTexelBufferArrayDynamicIndexing VkBool32
	shaderUniformBufferArrayNonUniformIndexing VkBool32
	shaderSampledImageArrayNonUniformIndexing VkBool32
	shaderStorageBufferArrayNonUniformIndexing VkBool32
	shaderStorageImageArrayNonUniformIndexing VkBool32
	shaderInputAttachmentArrayNonUniformIndexing VkBool32
	shaderUniformTexelBufferArrayNonUniformIndexing VkBool32
	shaderStorageTexelBufferArrayNonUniformIndexing VkBool32
	descriptorBindingUniformBufferUpdateAfterBind VkBool32
	descriptorBindingSampledImageUpdateAfterBind VkBool32
	descriptorBindingStorageImageUpdateAfterBind VkBool32
	descriptorBindingStorageBufferUpdateAfterBind VkBool32
	descriptorBindingUniformTexelBufferUpdateAfterBind VkBool32
	descriptorBindingStorageTexelBufferUpdateAfterBind VkBool32
	descriptorBindingUpdateUnusedWhilePending VkBool32
	descriptorBindingPartiallyBound VkBool32
	descriptorBindingVariableDescriptorCount VkBool32
	runtimeDescriptorArray VkBool32
}
type VkPhysicalDeviceDescriptorIndexingFeaturesEXT struct {
}
type VkPhysicalDeviceDescriptorIndexingProperties struct {
	sType VkStructureType
	pNext uintptr
	maxUpdateAfterBindDescriptorsInAllPools uint32
	shaderUniformBufferArrayNonUniformIndexingNative VkBool32
	shaderSampledImageArrayNonUniformIndexingNative VkBool32
	shaderStorageBufferArrayNonUniformIndexingNative VkBool32
	shaderStorageImageArrayNonUniformIndexingNative VkBool32
	shaderInputAttachmentArrayNonUniformIndexingNative VkBool32
	robustBufferAccessUpdateAfterBind VkBool32
	quadDivergentImplicitLod VkBool32
	maxPerStageDescriptorUpdateAfterBindSamplers uint32
	maxPerStageDescriptorUpdateAfterBindUniformBuffers uint32
	maxPerStageDescriptorUpdateAfterBindStorageBuffers uint32
	maxPerStageDescriptorUpdateAfterBindSampledImages uint32
	maxPerStageDescriptorUpdateAfterBindStorageImages uint32
	maxPerStageDescriptorUpdateAfterBindInputAttachments uint32
	maxPerStageUpdateAfterBindResources uint32
	maxDescriptorSetUpdateAfterBindSamplers uint32
	maxDescriptorSetUpdateAfterBindUniformBuffers uint32
	maxDescriptorSetUpdateAfterBindUniformBuffersDynamic uint32
	maxDescriptorSetUpdateAfterBindStorageBuffers uint32
	maxDescriptorSetUpdateAfterBindStorageBuffersDynamic uint32
	maxDescriptorSetUpdateAfterBindSampledImages uint32
	maxDescriptorSetUpdateAfterBindStorageImages uint32
	maxDescriptorSetUpdateAfterBindInputAttachments uint32
}
type VkPhysicalDeviceDescriptorIndexingPropertiesEXT struct {
}
type VkDescriptorSetLayoutBindingFlagsCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	bindingCount uint32
	pBindingFlags *VkDescriptorBindingFlags
}
type VkDescriptorSetLayoutBindingFlagsCreateInfoEXT struct {
}
type VkDescriptorSetVariableDescriptorCountAllocateInfo struct {
	sType VkStructureType
	pNext uintptr
	descriptorSetCount uint32
	pDescriptorCounts *uint32
}
type VkDescriptorSetVariableDescriptorCountAllocateInfoEXT struct {
}
type VkDescriptorSetVariableDescriptorCountLayoutSupport struct {
	sType VkStructureType
	pNext uintptr
	maxVariableDescriptorCount uint32
}
type VkDescriptorSetVariableDescriptorCountLayoutSupportEXT struct {
}
type VkAttachmentDescription2 struct {
	sType VkStructureType
	pNext uintptr
	flags VkAttachmentDescriptionFlags
	format VkFormat
	samples VkSampleCountFlagBits
	loadOp VkAttachmentLoadOp
	storeOp VkAttachmentStoreOp
	stencilLoadOp VkAttachmentLoadOp
	stencilStoreOp VkAttachmentStoreOp
	initialLayout VkImageLayout
	finalLayout VkImageLayout
}
type VkAttachmentDescription2KHR struct {
}
type VkAttachmentReference2 struct {
	sType VkStructureType
	pNext uintptr
	attachment uint32
	layout VkImageLayout
	aspectMask VkImageAspectFlags
}
type VkAttachmentReference2KHR struct {
}
type VkSubpassDescription2 struct {
	sType VkStructureType
	pNext uintptr
	flags VkSubpassDescriptionFlags
	pipelineBindPoint VkPipelineBindPoint
	viewMask uint32
	inputAttachmentCount uint32
	pInputAttachments *VkAttachmentReference2
	colorAttachmentCount uint32
	pColorAttachments *VkAttachmentReference2
	pResolveAttachments *VkAttachmentReference2
	pDepthStencilAttachment *VkAttachmentReference2
	preserveAttachmentCount uint32
	pPreserveAttachments *uint32
}
type VkSubpassDescription2KHR struct {
}
type VkSubpassDependency2 struct {
	sType VkStructureType
	pNext uintptr
	srcSubpass uint32
	dstSubpass uint32
	srcStageMask VkPipelineStageFlags
	dstStageMask VkPipelineStageFlags
	srcAccessMask VkAccessFlags
	dstAccessMask VkAccessFlags
	dependencyFlags VkDependencyFlags
	viewOffset int32
}
type VkSubpassDependency2KHR struct {
}
type VkRenderPassCreateInfo2 struct {
	sType VkStructureType
	pNext uintptr
	flags VkRenderPassCreateFlags
	attachmentCount uint32
	pAttachments *VkAttachmentDescription2
	subpassCount uint32
	pSubpasses *VkSubpassDescription2
	dependencyCount uint32
	pDependencies *VkSubpassDependency2
	correlatedViewMaskCount uint32
	pCorrelatedViewMasks *uint32
}
type VkRenderPassCreateInfo2KHR struct {
}
type VkSubpassBeginInfo struct {
	sType VkStructureType
	pNext uintptr
	contents VkSubpassContents
}
type VkSubpassBeginInfoKHR struct {
}
type VkSubpassEndInfo struct {
	sType VkStructureType
	pNext uintptr
}
type VkSubpassEndInfoKHR struct {
}
type VkPhysicalDeviceTimelineSemaphoreFeatures struct {
	sType VkStructureType
	pNext uintptr
	timelineSemaphore VkBool32
}
type VkPhysicalDeviceTimelineSemaphoreFeaturesKHR struct {
}
type VkPhysicalDeviceTimelineSemaphoreProperties struct {
	sType VkStructureType
	pNext uintptr
	maxTimelineSemaphoreValueDifference uint64
}
type VkPhysicalDeviceTimelineSemaphorePropertiesKHR struct {
}
type VkSemaphoreTypeCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	semaphoreType VkSemaphoreType
	initialValue uint64
}
type VkSemaphoreTypeCreateInfoKHR struct {
}
type VkTimelineSemaphoreSubmitInfo struct {
	sType VkStructureType
	pNext uintptr
	waitSemaphoreValueCount uint32
	pWaitSemaphoreValues *uint64
	signalSemaphoreValueCount uint32
	pSignalSemaphoreValues *uint64
}
type VkTimelineSemaphoreSubmitInfoKHR struct {
}
type VkSemaphoreWaitInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkSemaphoreWaitFlags
	semaphoreCount uint32
	pSemaphores *VkSemaphore
	pValues *uint64
}
type VkSemaphoreWaitInfoKHR struct {
}
type VkSemaphoreSignalInfo struct {
	sType VkStructureType
	pNext uintptr
	semaphore VkSemaphore
	value uint64
}
type VkSemaphoreSignalInfoKHR struct {
}
type VkVertexInputBindingDivisorDescriptionEXT struct {
	binding uint32
	divisor uint32
}
type VkPipelineVertexInputDivisorStateCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	vertexBindingDivisorCount uint32
	pVertexBindingDivisors *VkVertexInputBindingDivisorDescriptionEXT
}
type VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	maxVertexAttribDivisor uint32
}
type VkPhysicalDevicePCIBusInfoPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	pciDomain uint32
	pciBus uint32
	pciDevice uint32
	pciFunction uint32
}
type VkImportAndroidHardwareBufferInfoANDROID struct {
	sType VkStructureType
	pNext uintptr
	buffer *AHardwareBuffer
}
type VkAndroidHardwareBufferUsageANDROID struct {
	sType VkStructureType
	pNext uintptr
	androidHardwareBufferUsage uint64
}
type VkAndroidHardwareBufferPropertiesANDROID struct {
	sType VkStructureType
	pNext uintptr
	allocationSize VkDeviceSize
	memoryTypeBits uint32
}
type VkMemoryGetAndroidHardwareBufferInfoANDROID struct {
	sType VkStructureType
	pNext uintptr
	memory VkDeviceMemory
}
type VkAndroidHardwareBufferFormatPropertiesANDROID struct {
	sType VkStructureType
	pNext uintptr
	format VkFormat
	externalFormat uint64
	formatFeatures VkFormatFeatureFlags
	samplerYcbcrConversionComponents VkComponentMapping
	suggestedYcbcrModel VkSamplerYcbcrModelConversion
	suggestedYcbcrRange VkSamplerYcbcrRange
	suggestedXChromaOffset VkChromaLocation
	suggestedYChromaOffset VkChromaLocation
}
type VkCommandBufferInheritanceConditionalRenderingInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	conditionalRenderingEnable VkBool32
}
type VkExternalFormatANDROID struct {
	sType VkStructureType
	pNext uintptr
	externalFormat uint64
}
type VkPhysicalDevice8BitStorageFeatures struct {
	sType VkStructureType
	pNext uintptr
	storageBuffer8BitAccess VkBool32
	uniformAndStorageBuffer8BitAccess VkBool32
	storagePushConstant8 VkBool32
}
type VkPhysicalDevice8BitStorageFeaturesKHR struct {
}
type VkPhysicalDeviceConditionalRenderingFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	conditionalRendering VkBool32
	inheritedConditionalRendering VkBool32
}
type VkPhysicalDeviceVulkanMemoryModelFeatures struct {
	sType VkStructureType
	pNext uintptr
	vulkanMemoryModel VkBool32
	vulkanMemoryModelDeviceScope VkBool32
	vulkanMemoryModelAvailabilityVisibilityChains VkBool32
}
type VkPhysicalDeviceVulkanMemoryModelFeaturesKHR struct {
}
type VkPhysicalDeviceShaderAtomicInt64Features struct {
	sType VkStructureType
	pNext uintptr
	shaderBufferInt64Atomics VkBool32
	shaderSharedInt64Atomics VkBool32
}
type VkPhysicalDeviceShaderAtomicInt64FeaturesKHR struct {
}
type VkPhysicalDeviceShaderAtomicFloatFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	shaderBufferFloat32Atomics VkBool32
	shaderBufferFloat32AtomicAdd VkBool32
	shaderBufferFloat64Atomics VkBool32
	shaderBufferFloat64AtomicAdd VkBool32
	shaderSharedFloat32Atomics VkBool32
	shaderSharedFloat32AtomicAdd VkBool32
	shaderSharedFloat64Atomics VkBool32
	shaderSharedFloat64AtomicAdd VkBool32
	shaderImageFloat32Atomics VkBool32
	shaderImageFloat32AtomicAdd VkBool32
	sparseImageFloat32Atomics VkBool32
	sparseImageFloat32AtomicAdd VkBool32
}
type VkPhysicalDeviceVertexAttributeDivisorFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	vertexAttributeInstanceRateDivisor VkBool32
	vertexAttributeInstanceRateZeroDivisor VkBool32
}
type VkQueueFamilyCheckpointPropertiesNV struct {
	sType VkStructureType
	pNext uintptr
	checkpointExecutionStageMask VkPipelineStageFlags
}
type VkCheckpointDataNV struct {
	sType VkStructureType
	pNext uintptr
	stage VkPipelineStageFlagBits
	pCheckpointMarker uintptr
}
type VkPhysicalDeviceDepthStencilResolveProperties struct {
	sType VkStructureType
	pNext uintptr
	supportedDepthResolveModes VkResolveModeFlags
	supportedStencilResolveModes VkResolveModeFlags
	independentResolveNone VkBool32
	independentResolve VkBool32
}
type VkPhysicalDeviceDepthStencilResolvePropertiesKHR struct {
}
type VkSubpassDescriptionDepthStencilResolve struct {
	sType VkStructureType
	pNext uintptr
	depthResolveMode VkResolveModeFlagBits
	stencilResolveMode VkResolveModeFlagBits
	pDepthStencilResolveAttachment *VkAttachmentReference2
}
type VkSubpassDescriptionDepthStencilResolveKHR struct {
}
type VkImageViewASTCDecodeModeEXT struct {
	sType VkStructureType
	pNext uintptr
	decodeMode VkFormat
}
type VkPhysicalDeviceASTCDecodeFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	decodeModeSharedExponent VkBool32
}
type VkPhysicalDeviceTransformFeedbackFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	transformFeedback VkBool32
	geometryStreams VkBool32
}
type VkPhysicalDeviceTransformFeedbackPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	maxTransformFeedbackStreams uint32
	maxTransformFeedbackBuffers uint32
	maxTransformFeedbackBufferSize VkDeviceSize
	maxTransformFeedbackStreamDataSize uint32
	maxTransformFeedbackBufferDataSize uint32
	maxTransformFeedbackBufferDataStride uint32
	transformFeedbackQueries VkBool32
	transformFeedbackStreamsLinesTriangles VkBool32
	transformFeedbackRasterizationStreamSelect VkBool32
	transformFeedbackDraw VkBool32
}
type VkPipelineRasterizationStateStreamCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineRasterizationStateStreamCreateFlagsEXT
	rasterizationStream uint32
}
type VkPhysicalDeviceRepresentativeFragmentTestFeaturesNV struct {
	sType VkStructureType
	pNext uintptr
	representativeFragmentTest VkBool32
}
type VkPipelineRepresentativeFragmentTestStateCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	representativeFragmentTestEnable VkBool32
}
type VkPhysicalDeviceExclusiveScissorFeaturesNV struct {
	sType VkStructureType
	pNext uintptr
	exclusiveScissor VkBool32
}
type VkPipelineViewportExclusiveScissorStateCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	exclusiveScissorCount uint32
	pExclusiveScissors *VkRect2D
}
type VkPhysicalDeviceCornerSampledImageFeaturesNV struct {
	sType VkStructureType
	pNext uintptr
	cornerSampledImage VkBool32
}
type VkPhysicalDeviceComputeShaderDerivativesFeaturesNV struct {
	sType VkStructureType
	pNext uintptr
	computeDerivativeGroupQuads VkBool32
	computeDerivativeGroupLinear VkBool32
}
type VkPhysicalDeviceFragmentShaderBarycentricFeaturesNV struct {
	sType VkStructureType
	pNext uintptr
	fragmentShaderBarycentric VkBool32
}
type VkPhysicalDeviceShaderImageFootprintFeaturesNV struct {
	sType VkStructureType
	pNext uintptr
	imageFootprint VkBool32
}
type VkPhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV struct {
	sType VkStructureType
	pNext uintptr
	dedicatedAllocationImageAliasing VkBool32
}
type VkShadingRatePaletteNV struct {
	shadingRatePaletteEntryCount uint32
	pShadingRatePaletteEntries *VkShadingRatePaletteEntryNV
}
type VkPipelineViewportShadingRateImageStateCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	shadingRateImageEnable VkBool32
	viewportCount uint32
	pShadingRatePalettes *VkShadingRatePaletteNV
}
type VkPhysicalDeviceShadingRateImageFeaturesNV struct {
	sType VkStructureType
	pNext uintptr
	shadingRateImage VkBool32
	shadingRateCoarseSampleOrder VkBool32
}
type VkPhysicalDeviceShadingRateImagePropertiesNV struct {
	sType VkStructureType
	pNext uintptr
	shadingRateTexelSize VkExtent2D
	shadingRatePaletteSize uint32
	shadingRateMaxCoarseSamples uint32
}
type VkCoarseSampleLocationNV struct {
	pixelX uint32
	pixelY uint32
	sample uint32
}
type VkCoarseSampleOrderCustomNV struct {
	shadingRate VkShadingRatePaletteEntryNV
	sampleCount uint32
	sampleLocationCount uint32
	pSampleLocations *VkCoarseSampleLocationNV
}
type VkPipelineViewportCoarseSampleOrderStateCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	sampleOrderType VkCoarseSampleOrderTypeNV
	customSampleOrderCount uint32
	pCustomSampleOrders *VkCoarseSampleOrderCustomNV
}
type VkPhysicalDeviceMeshShaderFeaturesNV struct {
	sType VkStructureType
	pNext uintptr
	taskShader VkBool32
	meshShader VkBool32
}
type VkPhysicalDeviceMeshShaderPropertiesNV struct {
	sType VkStructureType
	pNext uintptr
	maxDrawMeshTasksCount uint32
	maxTaskWorkGroupInvocations uint32
	maxTaskWorkGroupSize uint32
	maxTaskTotalMemorySize uint32
	maxTaskOutputCount uint32
	maxMeshWorkGroupInvocations uint32
	maxMeshWorkGroupSize uint32
	maxMeshTotalMemorySize uint32
	maxMeshOutputVertices uint32
	maxMeshOutputPrimitives uint32
	maxMeshMultiviewViewCount uint32
	meshOutputPerVertexGranularity uint32
	meshOutputPerPrimitiveGranularity uint32
}
type VkDrawMeshTasksIndirectCommandNV struct {
	taskCount uint32
	firstTask uint32
}
type VkRayTracingShaderGroupCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	type0 VkRayTracingShaderGroupTypeKHR
	generalShader uint32
	closestHitShader uint32
	anyHitShader uint32
	intersectionShader uint32
}
type VkRayTracingShaderGroupCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	type0 VkRayTracingShaderGroupTypeKHR
	generalShader uint32
	closestHitShader uint32
	anyHitShader uint32
	intersectionShader uint32
	pShaderGroupCaptureReplayHandle uintptr
}
type VkRayTracingPipelineCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineCreateFlags
	stageCount uint32
	pStages *VkPipelineShaderStageCreateInfo
	groupCount uint32
	pGroups *VkRayTracingShaderGroupCreateInfoNV
	maxRecursionDepth uint32
	layout VkPipelineLayout
	basePipelineHandle VkPipeline
	basePipelineIndex int32
}
type VkRayTracingPipelineCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineCreateFlags
	stageCount uint32
	pStages *VkPipelineShaderStageCreateInfo
	groupCount uint32
	pGroups *VkRayTracingShaderGroupCreateInfoKHR
	maxRecursionDepth uint32
	libraries VkPipelineLibraryCreateInfoKHR
	pLibraryInterface *VkRayTracingPipelineInterfaceCreateInfoKHR
	layout VkPipelineLayout
	basePipelineHandle VkPipeline
	basePipelineIndex int32
}
type VkGeometryTrianglesNV struct {
	sType VkStructureType
	pNext uintptr
	vertexData VkBuffer
	vertexOffset VkDeviceSize
	vertexCount uint32
	vertexStride VkDeviceSize
	vertexFormat VkFormat
	indexData VkBuffer
	indexOffset VkDeviceSize
	indexCount uint32
	indexType VkIndexType
	transformData VkBuffer
	transformOffset VkDeviceSize
}
type VkGeometryAABBNV struct {
	sType VkStructureType
	pNext uintptr
	aabbData VkBuffer
	numAABBs uint32
	stride uint32
	offset VkDeviceSize
}
type VkGeometryDataNV struct {
	triangles VkGeometryTrianglesNV
	aabbs VkGeometryAABBNV
}
type VkGeometryNV struct {
	sType VkStructureType
	pNext uintptr
	geometryType VkGeometryTypeKHR
	geometry VkGeometryDataNV
	flags VkGeometryFlagsKHR
}
type VkAccelerationStructureInfoNV struct {
	sType VkStructureType
	pNext uintptr
	type0 VkAccelerationStructureTypeNV
	flags VkBuildAccelerationStructureFlagsNV
	instanceCount uint32
	geometryCount uint32
	pGeometries *VkGeometryNV
}
type VkAccelerationStructureCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	compactedSize VkDeviceSize
	info VkAccelerationStructureInfoNV
}
type VkBindAccelerationStructureMemoryInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	accelerationStructure VkAccelerationStructureKHR
	memory VkDeviceMemory
	memoryOffset VkDeviceSize
	deviceIndexCount uint32
	pDeviceIndices *uint32
}
type VkBindAccelerationStructureMemoryInfoNV struct {
}
type VkWriteDescriptorSetAccelerationStructureKHR struct {
	sType VkStructureType
	pNext uintptr
	accelerationStructureCount uint32
	pAccelerationStructures *VkAccelerationStructureKHR
}
type VkWriteDescriptorSetAccelerationStructureNV struct {
}
type VkAccelerationStructureMemoryRequirementsInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	type0 VkAccelerationStructureMemoryRequirementsTypeKHR
	buildType VkAccelerationStructureBuildTypeKHR
	accelerationStructure VkAccelerationStructureKHR
}
type VkAccelerationStructureMemoryRequirementsInfoNV struct {
	sType VkStructureType
	pNext uintptr
	type0 VkAccelerationStructureMemoryRequirementsTypeNV
	accelerationStructure VkAccelerationStructureNV
}
type VkPhysicalDeviceRayTracingFeaturesKHR struct {
	sType VkStructureType
	pNext uintptr
	rayTracing VkBool32
	rayTracingShaderGroupHandleCaptureReplay VkBool32
	rayTracingShaderGroupHandleCaptureReplayMixed VkBool32
	rayTracingAccelerationStructureCaptureReplay VkBool32
	rayTracingIndirectTraceRays VkBool32
	rayTracingIndirectAccelerationStructureBuild VkBool32
	rayTracingHostAccelerationStructureCommands VkBool32
	rayQuery VkBool32
	rayTracingPrimitiveCulling VkBool32
}
type VkPhysicalDeviceRayTracingPropertiesKHR struct {
	sType VkStructureType
	pNext uintptr
	shaderGroupHandleSize uint32
	maxRecursionDepth uint32
	maxShaderGroupStride uint32
	shaderGroupBaseAlignment uint32
	maxGeometryCount uint64
	maxInstanceCount uint64
	maxPrimitiveCount uint64
	maxDescriptorSetAccelerationStructures uint32
	shaderGroupHandleCaptureReplaySize uint32
}
type VkPhysicalDeviceRayTracingPropertiesNV struct {
	sType VkStructureType
	pNext uintptr
	shaderGroupHandleSize uint32
	maxRecursionDepth uint32
	maxShaderGroupStride uint32
	shaderGroupBaseAlignment uint32
	maxGeometryCount uint64
	maxInstanceCount uint64
	maxTriangleCount uint64
	maxDescriptorSetAccelerationStructures uint32
}
type VkStridedBufferRegionKHR struct {
	buffer VkBuffer
	offset VkDeviceSize
	stride VkDeviceSize
	size VkDeviceSize
}
type VkTraceRaysIndirectCommandKHR struct {
	width uint32
	height uint32
	depth uint32
}
type VkDrmFormatModifierPropertiesListEXT struct {
	sType VkStructureType
	pNext uintptr
	drmFormatModifierCount uint32
	pDrmFormatModifierProperties *VkDrmFormatModifierPropertiesEXT
}
type VkDrmFormatModifierPropertiesEXT struct {
	drmFormatModifier uint64
	drmFormatModifierPlaneCount uint32
	drmFormatModifierTilingFeatures VkFormatFeatureFlags
}
type VkPhysicalDeviceImageDrmFormatModifierInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	drmFormatModifier uint64
	sharingMode VkSharingMode
	queueFamilyIndexCount uint32
	pQueueFamilyIndices *uint32
}
type VkImageDrmFormatModifierListCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	drmFormatModifierCount uint32
	pDrmFormatModifiers *uint64
}
type VkImageDrmFormatModifierExplicitCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	drmFormatModifier uint64
	drmFormatModifierPlaneCount uint32
	pPlaneLayouts *VkSubresourceLayout
}
type VkImageDrmFormatModifierPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	drmFormatModifier uint64
}
type VkImageStencilUsageCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	stencilUsage VkImageUsageFlags
}
type VkImageStencilUsageCreateInfoEXT struct {
}
type VkDeviceMemoryOverallocationCreateInfoAMD struct {
	sType VkStructureType
	pNext uintptr
	overallocationBehavior VkMemoryOverallocationBehaviorAMD
}
type VkPhysicalDeviceFragmentDensityMapFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	fragmentDensityMap VkBool32
	fragmentDensityMapDynamic VkBool32
	fragmentDensityMapNonSubsampledImages VkBool32
}
type VkPhysicalDeviceFragmentDensityMap2FeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	fragmentDensityMapDeferred VkBool32
}
type VkPhysicalDeviceFragmentDensityMapPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	minFragmentDensityTexelSize VkExtent2D
	maxFragmentDensityTexelSize VkExtent2D
	fragmentDensityInvocations VkBool32
}
type VkPhysicalDeviceFragmentDensityMap2PropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	subsampledLoads VkBool32
	subsampledCoarseReconstructionEarlyAccess VkBool32
	maxSubsampledArrayLayers uint32
	maxDescriptorSetSubsampledSamplers uint32
}
type VkRenderPassFragmentDensityMapCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	fragmentDensityMapAttachment VkAttachmentReference
}
type VkPhysicalDeviceScalarBlockLayoutFeatures struct {
	sType VkStructureType
	pNext uintptr
	scalarBlockLayout VkBool32
}
type VkPhysicalDeviceScalarBlockLayoutFeaturesEXT struct {
}
type VkSurfaceProtectedCapabilitiesKHR struct {
	sType VkStructureType
	pNext uintptr
	supportsProtected VkBool32
}
type VkPhysicalDeviceUniformBufferStandardLayoutFeatures struct {
	sType VkStructureType
	pNext uintptr
	uniformBufferStandardLayout VkBool32
}
type VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR struct {
}
type VkPhysicalDeviceDepthClipEnableFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	depthClipEnable VkBool32
}
type VkPipelineRasterizationDepthClipStateCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineRasterizationDepthClipStateCreateFlagsEXT
	depthClipEnable VkBool32
}
type VkPhysicalDeviceMemoryBudgetPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	heapBudget VkDeviceSize
	heapUsage VkDeviceSize
}
type VkPhysicalDeviceMemoryPriorityFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	memoryPriority VkBool32
}
type VkMemoryPriorityAllocateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	priority float32
}
type VkPhysicalDeviceBufferDeviceAddressFeatures struct {
	sType VkStructureType
	pNext uintptr
	bufferDeviceAddress VkBool32
	bufferDeviceAddressCaptureReplay VkBool32
	bufferDeviceAddressMultiDevice VkBool32
}
type VkPhysicalDeviceBufferDeviceAddressFeaturesKHR struct {
}
type VkPhysicalDeviceBufferDeviceAddressFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	bufferDeviceAddress VkBool32
	bufferDeviceAddressCaptureReplay VkBool32
	bufferDeviceAddressMultiDevice VkBool32
}
type VkPhysicalDeviceBufferAddressFeaturesEXT struct {
}
type VkBufferDeviceAddressInfo struct {
	sType VkStructureType
	pNext uintptr
	buffer VkBuffer
}
type VkBufferDeviceAddressInfoKHR struct {
}
type VkBufferDeviceAddressInfoEXT struct {
}
type VkBufferOpaqueCaptureAddressCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	opaqueCaptureAddress uint64
}
type VkBufferOpaqueCaptureAddressCreateInfoKHR struct {
}
type VkBufferDeviceAddressCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	deviceAddress VkDeviceAddress
}
type VkPhysicalDeviceImageViewImageFormatInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	imageViewType VkImageViewType
}
type VkFilterCubicImageViewImageFormatPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	filterCubic VkBool32
	filterCubicMinmax VkBool32
}
type VkPhysicalDeviceImagelessFramebufferFeatures struct {
	sType VkStructureType
	pNext uintptr
	imagelessFramebuffer VkBool32
}
type VkPhysicalDeviceImagelessFramebufferFeaturesKHR struct {
}
type VkFramebufferAttachmentsCreateInfo struct {
	sType VkStructureType
	pNext uintptr
	attachmentImageInfoCount uint32
	pAttachmentImageInfos *VkFramebufferAttachmentImageInfo
}
type VkFramebufferAttachmentsCreateInfoKHR struct {
}
type VkFramebufferAttachmentImageInfo struct {
	sType VkStructureType
	pNext uintptr
	flags VkImageCreateFlags
	usage VkImageUsageFlags
	width uint32
	height uint32
	layerCount uint32
	viewFormatCount uint32
	pViewFormats *VkFormat
}
type VkFramebufferAttachmentImageInfoKHR struct {
}
type VkRenderPassAttachmentBeginInfo struct {
	sType VkStructureType
	pNext uintptr
	attachmentCount uint32
	pAttachments *VkImageView
}
type VkRenderPassAttachmentBeginInfoKHR struct {
}
type VkPhysicalDeviceTextureCompressionASTCHDRFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	textureCompressionASTC_HDR VkBool32
}
type VkPhysicalDeviceCooperativeMatrixFeaturesNV struct {
	sType VkStructureType
	pNext uintptr
	cooperativeMatrix VkBool32
	cooperativeMatrixRobustBufferAccess VkBool32
}
type VkPhysicalDeviceCooperativeMatrixPropertiesNV struct {
	sType VkStructureType
	pNext uintptr
	cooperativeMatrixSupportedStages VkShaderStageFlags
}
type VkCooperativeMatrixPropertiesNV struct {
	sType VkStructureType
	pNext uintptr
	MSize uint32
	NSize uint32
	KSize uint32
	AType VkComponentTypeNV
	BType VkComponentTypeNV
	CType VkComponentTypeNV
	DType VkComponentTypeNV
	scope VkScopeNV
}
type VkPhysicalDeviceYcbcrImageArraysFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	ycbcrImageArrays VkBool32
}
type VkImageViewHandleInfoNVX struct {
	sType VkStructureType
	pNext uintptr
	imageView VkImageView
	descriptorType VkDescriptorType
	sampler VkSampler
}
type VkImageViewAddressPropertiesNVX struct {
	sType VkStructureType
	pNext uintptr
	deviceAddress VkDeviceAddress
	size VkDeviceSize
}
type VkPresentFrameTokenGGP struct {
	sType VkStructureType
	pNext uintptr
	frameToken GgpFrameToken
}
type VkPipelineCreationFeedbackEXT struct {
	flags VkPipelineCreationFeedbackFlagsEXT
	duration uint64
}
type VkPipelineCreationFeedbackCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	pPipelineCreationFeedback *VkPipelineCreationFeedbackEXT
	pipelineStageCreationFeedbackCount uint32
	pPipelineStageCreationFeedbacks *VkPipelineCreationFeedbackEXT
}
type VkSurfaceFullScreenExclusiveInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	fullScreenExclusive VkFullScreenExclusiveEXT
}
type VkSurfaceFullScreenExclusiveWin32InfoEXT struct {
	sType VkStructureType
	pNext uintptr
	hmonitor HMONITOR
}
type VkSurfaceCapabilitiesFullScreenExclusiveEXT struct {
	sType VkStructureType
	pNext uintptr
	fullScreenExclusiveSupported VkBool32
}
type VkPhysicalDevicePerformanceQueryFeaturesKHR struct {
	sType VkStructureType
	pNext uintptr
	performanceCounterQueryPools VkBool32
	performanceCounterMultipleQueryPools VkBool32
}
type VkPhysicalDevicePerformanceQueryPropertiesKHR struct {
	sType VkStructureType
	pNext uintptr
	allowCommandBufferQueryCopies VkBool32
}
type VkPerformanceCounterKHR struct {
	sType VkStructureType
	pNext uintptr
	unit VkPerformanceCounterUnitKHR
	scope VkPerformanceCounterScopeKHR
	storage VkPerformanceCounterStorageKHR
	uuid byte
}
type VkPerformanceCounterDescriptionKHR struct {
	sType VkStructureType
	pNext uintptr
	flags VkPerformanceCounterDescriptionFlagsKHR
	name byte
	category byte
	description byte
}
type VkQueryPoolPerformanceCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	queueFamilyIndex uint32
	counterIndexCount uint32
	pCounterIndices *uint32
}
type VkPerformanceCounterResultKHR = struct{uintptr}
type VkAcquireProfilingLockInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	flags VkAcquireProfilingLockFlagsKHR
	timeout uint64
}
type VkPerformanceQuerySubmitInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	counterPassIndex uint32
}
type VkHeadlessSurfaceCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	flags VkHeadlessSurfaceCreateFlagsEXT
}
type VkPhysicalDeviceCoverageReductionModeFeaturesNV struct {
	sType VkStructureType
	pNext uintptr
	coverageReductionMode VkBool32
}
type VkPipelineCoverageReductionStateCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	flags VkPipelineCoverageReductionStateCreateFlagsNV
	coverageReductionMode VkCoverageReductionModeNV
}
type VkFramebufferMixedSamplesCombinationNV struct {
	sType VkStructureType
	pNext uintptr
	coverageReductionMode VkCoverageReductionModeNV
	rasterizationSamples VkSampleCountFlagBits
	depthStencilSamples VkSampleCountFlags
	colorSamples VkSampleCountFlags
}
type VkPhysicalDeviceShaderIntegerFunctions2FeaturesINTEL struct {
	sType VkStructureType
	pNext uintptr
	shaderIntegerFunctions2 VkBool32
}
type VkPerformanceValueDataINTEL = struct{uintptr}
type VkPerformanceValueINTEL struct {
	type0 VkPerformanceValueTypeINTEL
	data VkPerformanceValueDataINTEL
}
type VkInitializePerformanceApiInfoINTEL struct {
	sType VkStructureType
	pNext uintptr
	pUserData uintptr
}
type VkQueryPoolPerformanceQueryCreateInfoINTEL struct {
	sType VkStructureType
	pNext uintptr
	performanceCountersSampling VkQueryPoolSamplingModeINTEL
}
type VkQueryPoolCreateInfoINTEL struct {
}
type VkPerformanceMarkerInfoINTEL struct {
	sType VkStructureType
	pNext uintptr
	marker uint64
}
type VkPerformanceStreamMarkerInfoINTEL struct {
	sType VkStructureType
	pNext uintptr
	marker uint32
}
type VkPerformanceOverrideInfoINTEL struct {
	sType VkStructureType
	pNext uintptr
	type0 VkPerformanceOverrideTypeINTEL
	enable VkBool32
	parameter uint64
}
type VkPerformanceConfigurationAcquireInfoINTEL struct {
	sType VkStructureType
	pNext uintptr
	type0 VkPerformanceConfigurationTypeINTEL
}
type VkPhysicalDeviceShaderClockFeaturesKHR struct {
	sType VkStructureType
	pNext uintptr
	shaderSubgroupClock VkBool32
	shaderDeviceClock VkBool32
}
type VkPhysicalDeviceIndexTypeUint8FeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	indexTypeUint8 VkBool32
}
type VkPhysicalDeviceShaderSMBuiltinsPropertiesNV struct {
	sType VkStructureType
	pNext uintptr
	shaderSMCount uint32
	shaderWarpsPerSM uint32
}
type VkPhysicalDeviceShaderSMBuiltinsFeaturesNV struct {
	sType VkStructureType
	pNext uintptr
	shaderSMBuiltins VkBool32
}
type VkPhysicalDeviceFragmentShaderInterlockFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	fragmentShaderSampleInterlock VkBool32
	fragmentShaderPixelInterlock VkBool32
	fragmentShaderShadingRateInterlock VkBool32
}
type VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures struct {
	sType VkStructureType
	pNext uintptr
	separateDepthStencilLayouts VkBool32
}
type VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR struct {
}
type VkAttachmentReferenceStencilLayout struct {
	sType VkStructureType
	pNext uintptr
	stencilLayout VkImageLayout
}
type VkAttachmentReferenceStencilLayoutKHR struct {
}
type VkAttachmentDescriptionStencilLayout struct {
	sType VkStructureType
	pNext uintptr
	stencilInitialLayout VkImageLayout
	stencilFinalLayout VkImageLayout
}
type VkAttachmentDescriptionStencilLayoutKHR struct {
}
type VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR struct {
	sType VkStructureType
	pNext uintptr
	pipelineExecutableInfo VkBool32
}
type VkPipelineInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	pipeline VkPipeline
}
type VkPipelineExecutablePropertiesKHR struct {
	sType VkStructureType
	pNext uintptr
	stages VkShaderStageFlags
	name byte
	description byte
	subgroupSize uint32
}
type VkPipelineExecutableInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	pipeline VkPipeline
	executableIndex uint32
}
type VkPipelineExecutableStatisticValueKHR = struct{uintptr}
type VkPipelineExecutableStatisticKHR struct {
	sType VkStructureType
	pNext uintptr
	name byte
	description byte
	format VkPipelineExecutableStatisticFormatKHR
	value VkPipelineExecutableStatisticValueKHR
}
type VkPipelineExecutableInternalRepresentationKHR struct {
	sType VkStructureType
	pNext uintptr
	name byte
	description byte
	isText VkBool32
	dataSize uint
	pData uintptr
}
type VkPhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	shaderDemoteToHelperInvocation VkBool32
}
type VkPhysicalDeviceTexelBufferAlignmentFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	texelBufferAlignment VkBool32
}
type VkPhysicalDeviceTexelBufferAlignmentPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	storageTexelBufferOffsetAlignmentBytes VkDeviceSize
	storageTexelBufferOffsetSingleTexelAlignment VkBool32
	uniformTexelBufferOffsetAlignmentBytes VkDeviceSize
	uniformTexelBufferOffsetSingleTexelAlignment VkBool32
}
type VkPhysicalDeviceSubgroupSizeControlFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	subgroupSizeControl VkBool32
	computeFullSubgroups VkBool32
}
type VkPhysicalDeviceSubgroupSizeControlPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	minSubgroupSize uint32
	maxSubgroupSize uint32
	maxComputeWorkgroupSubgroups uint32
	requiredSubgroupSizeStages VkShaderStageFlags
}
type VkPipelineShaderStageRequiredSubgroupSizeCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	requiredSubgroupSize uint32
}
type VkMemoryOpaqueCaptureAddressAllocateInfo struct {
	sType VkStructureType
	pNext uintptr
	opaqueCaptureAddress uint64
}
type VkMemoryOpaqueCaptureAddressAllocateInfoKHR struct {
}
type VkDeviceMemoryOpaqueCaptureAddressInfo struct {
	sType VkStructureType
	pNext uintptr
	memory VkDeviceMemory
}
type VkDeviceMemoryOpaqueCaptureAddressInfoKHR struct {
}
type VkPhysicalDeviceLineRasterizationFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	rectangularLines VkBool32
	bresenhamLines VkBool32
	smoothLines VkBool32
	stippledRectangularLines VkBool32
	stippledBresenhamLines VkBool32
	stippledSmoothLines VkBool32
}
type VkPhysicalDeviceLineRasterizationPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	lineSubPixelPrecisionBits uint32
}
type VkPipelineRasterizationLineStateCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	lineRasterizationMode VkLineRasterizationModeEXT
	stippledLineEnable VkBool32
	lineStippleFactor uint32
	lineStipplePattern uint16_t
}
type VkPhysicalDevicePipelineCreationCacheControlFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	pipelineCreationCacheControl VkBool32
}
type VkPhysicalDeviceVulkan11Features struct {
	sType VkStructureType
	pNext uintptr
	storageBuffer16BitAccess VkBool32
	uniformAndStorageBuffer16BitAccess VkBool32
	storagePushConstant16 VkBool32
	storageInputOutput16 VkBool32
	multiview VkBool32
	multiviewGeometryShader VkBool32
	multiviewTessellationShader VkBool32
	variablePointersStorageBuffer VkBool32
	variablePointers VkBool32
	protectedMemory VkBool32
	samplerYcbcrConversion VkBool32
	shaderDrawParameters VkBool32
}
type VkPhysicalDeviceVulkan11Properties struct {
	sType VkStructureType
	pNext uintptr
	deviceUUID byte
	driverUUID byte
	deviceLUID byte
	deviceNodeMask uint32
	deviceLUIDValid VkBool32
	subgroupSize uint32
	subgroupSupportedStages VkShaderStageFlags
	subgroupSupportedOperations VkSubgroupFeatureFlags
	subgroupQuadOperationsInAllStages VkBool32
	pointClippingBehavior VkPointClippingBehavior
	maxMultiviewViewCount uint32
	maxMultiviewInstanceIndex uint32
	protectedNoFault VkBool32
	maxPerSetDescriptors uint32
	maxMemoryAllocationSize VkDeviceSize
}
type VkPhysicalDeviceVulkan12Features struct {
	sType VkStructureType
	pNext uintptr
	samplerMirrorClampToEdge VkBool32
	drawIndirectCount VkBool32
	storageBuffer8BitAccess VkBool32
	uniformAndStorageBuffer8BitAccess VkBool32
	storagePushConstant8 VkBool32
	shaderBufferInt64Atomics VkBool32
	shaderSharedInt64Atomics VkBool32
	shaderFloat16 VkBool32
	shaderInt8 VkBool32
	descriptorIndexing VkBool32
	shaderInputAttachmentArrayDynamicIndexing VkBool32
	shaderUniformTexelBufferArrayDynamicIndexing VkBool32
	shaderStorageTexelBufferArrayDynamicIndexing VkBool32
	shaderUniformBufferArrayNonUniformIndexing VkBool32
	shaderSampledImageArrayNonUniformIndexing VkBool32
	shaderStorageBufferArrayNonUniformIndexing VkBool32
	shaderStorageImageArrayNonUniformIndexing VkBool32
	shaderInputAttachmentArrayNonUniformIndexing VkBool32
	shaderUniformTexelBufferArrayNonUniformIndexing VkBool32
	shaderStorageTexelBufferArrayNonUniformIndexing VkBool32
	descriptorBindingUniformBufferUpdateAfterBind VkBool32
	descriptorBindingSampledImageUpdateAfterBind VkBool32
	descriptorBindingStorageImageUpdateAfterBind VkBool32
	descriptorBindingStorageBufferUpdateAfterBind VkBool32
	descriptorBindingUniformTexelBufferUpdateAfterBind VkBool32
	descriptorBindingStorageTexelBufferUpdateAfterBind VkBool32
	descriptorBindingUpdateUnusedWhilePending VkBool32
	descriptorBindingPartiallyBound VkBool32
	descriptorBindingVariableDescriptorCount VkBool32
	runtimeDescriptorArray VkBool32
	samplerFilterMinmax VkBool32
	scalarBlockLayout VkBool32
	imagelessFramebuffer VkBool32
	uniformBufferStandardLayout VkBool32
	shaderSubgroupExtendedTypes VkBool32
	separateDepthStencilLayouts VkBool32
	hostQueryReset VkBool32
	timelineSemaphore VkBool32
	bufferDeviceAddress VkBool32
	bufferDeviceAddressCaptureReplay VkBool32
	bufferDeviceAddressMultiDevice VkBool32
	vulkanMemoryModel VkBool32
	vulkanMemoryModelDeviceScope VkBool32
	vulkanMemoryModelAvailabilityVisibilityChains VkBool32
	shaderOutputViewportIndex VkBool32
	shaderOutputLayer VkBool32
	subgroupBroadcastDynamicId VkBool32
}
type VkPhysicalDeviceVulkan12Properties struct {
	sType VkStructureType
	pNext uintptr
	driverID VkDriverId
	driverName byte
	driverInfo byte
	conformanceVersion VkConformanceVersion
	denormBehaviorIndependence VkShaderFloatControlsIndependence
	roundingModeIndependence VkShaderFloatControlsIndependence
	shaderSignedZeroInfNanPreserveFloat16 VkBool32
	shaderSignedZeroInfNanPreserveFloat32 VkBool32
	shaderSignedZeroInfNanPreserveFloat64 VkBool32
	shaderDenormPreserveFloat16 VkBool32
	shaderDenormPreserveFloat32 VkBool32
	shaderDenormPreserveFloat64 VkBool32
	shaderDenormFlushToZeroFloat16 VkBool32
	shaderDenormFlushToZeroFloat32 VkBool32
	shaderDenormFlushToZeroFloat64 VkBool32
	shaderRoundingModeRTEFloat16 VkBool32
	shaderRoundingModeRTEFloat32 VkBool32
	shaderRoundingModeRTEFloat64 VkBool32
	shaderRoundingModeRTZFloat16 VkBool32
	shaderRoundingModeRTZFloat32 VkBool32
	shaderRoundingModeRTZFloat64 VkBool32
	maxUpdateAfterBindDescriptorsInAllPools uint32
	shaderUniformBufferArrayNonUniformIndexingNative VkBool32
	shaderSampledImageArrayNonUniformIndexingNative VkBool32
	shaderStorageBufferArrayNonUniformIndexingNative VkBool32
	shaderStorageImageArrayNonUniformIndexingNative VkBool32
	shaderInputAttachmentArrayNonUniformIndexingNative VkBool32
	robustBufferAccessUpdateAfterBind VkBool32
	quadDivergentImplicitLod VkBool32
	maxPerStageDescriptorUpdateAfterBindSamplers uint32
	maxPerStageDescriptorUpdateAfterBindUniformBuffers uint32
	maxPerStageDescriptorUpdateAfterBindStorageBuffers uint32
	maxPerStageDescriptorUpdateAfterBindSampledImages uint32
	maxPerStageDescriptorUpdateAfterBindStorageImages uint32
	maxPerStageDescriptorUpdateAfterBindInputAttachments uint32
	maxPerStageUpdateAfterBindResources uint32
	maxDescriptorSetUpdateAfterBindSamplers uint32
	maxDescriptorSetUpdateAfterBindUniformBuffers uint32
	maxDescriptorSetUpdateAfterBindUniformBuffersDynamic uint32
	maxDescriptorSetUpdateAfterBindStorageBuffers uint32
	maxDescriptorSetUpdateAfterBindStorageBuffersDynamic uint32
	maxDescriptorSetUpdateAfterBindSampledImages uint32
	maxDescriptorSetUpdateAfterBindStorageImages uint32
	maxDescriptorSetUpdateAfterBindInputAttachments uint32
	supportedDepthResolveModes VkResolveModeFlags
	supportedStencilResolveModes VkResolveModeFlags
	independentResolveNone VkBool32
	independentResolve VkBool32
	filterMinmaxSingleComponentFormats VkBool32
	filterMinmaxImageComponentMapping VkBool32
	maxTimelineSemaphoreValueDifference uint64
	framebufferIntegerColorSampleCounts VkSampleCountFlags
}
type VkPipelineCompilerControlCreateInfoAMD struct {
	sType VkStructureType
	pNext uintptr
	compilerControlFlags VkPipelineCompilerControlFlagsAMD
}
type VkPhysicalDeviceCoherentMemoryFeaturesAMD struct {
	sType VkStructureType
	pNext uintptr
	deviceCoherentMemory VkBool32
}
type VkPhysicalDeviceToolPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	name byte
	version byte
	purposes VkToolPurposeFlagsEXT
	description byte
	layer byte
}
type VkSamplerCustomBorderColorCreateInfoEXT struct {
	sType VkStructureType
	pNext uintptr
	customBorderColor VkClearColorValue
	format VkFormat
}
type VkPhysicalDeviceCustomBorderColorPropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	maxCustomBorderColorSamplers uint32
}
type VkPhysicalDeviceCustomBorderColorFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	customBorderColors VkBool32
	customBorderColorWithoutFormat VkBool32
}
type VkDeviceOrHostAddressKHR = struct{uintptr}
type VkDeviceOrHostAddressConstKHR = struct{uintptr}
type VkAccelerationStructureGeometryTrianglesDataKHR struct {
	sType VkStructureType
	pNext uintptr
	vertexFormat VkFormat
	vertexData VkDeviceOrHostAddressConstKHR
	vertexStride VkDeviceSize
	indexType VkIndexType
	indexData VkDeviceOrHostAddressConstKHR
	transformData VkDeviceOrHostAddressConstKHR
}
type VkAccelerationStructureGeometryAabbsDataKHR struct {
	sType VkStructureType
	pNext uintptr
	data VkDeviceOrHostAddressConstKHR
	stride VkDeviceSize
}
type VkAccelerationStructureGeometryInstancesDataKHR struct {
	sType VkStructureType
	pNext uintptr
	arrayOfPointers VkBool32
	data VkDeviceOrHostAddressConstKHR
}
type VkAccelerationStructureGeometryDataKHR = struct{uintptr}
type VkAccelerationStructureGeometryKHR struct {
	sType VkStructureType
	pNext uintptr
	geometryType VkGeometryTypeKHR
	geometry VkAccelerationStructureGeometryDataKHR
	flags VkGeometryFlagsKHR
}
type VkAccelerationStructureBuildGeometryInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	type0 VkAccelerationStructureTypeKHR
	flags VkBuildAccelerationStructureFlagsKHR
	update VkBool32
	srcAccelerationStructure VkAccelerationStructureKHR
	dstAccelerationStructure VkAccelerationStructureKHR
	geometryArrayOfPointers VkBool32
	geometryCount uint32
	ppGeometries **VkAccelerationStructureGeometryKHR
	scratchData VkDeviceOrHostAddressKHR
}
type VkAccelerationStructureBuildOffsetInfoKHR struct {
	primitiveCount uint32
	primitiveOffset uint32
	firstVertex uint32
	transformOffset uint32
}
type VkAccelerationStructureCreateGeometryTypeInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	geometryType VkGeometryTypeKHR
	maxPrimitiveCount uint32
	indexType VkIndexType
	maxVertexCount uint32
	vertexFormat VkFormat
	allowsTransforms VkBool32
}
type VkAccelerationStructureCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	compactedSize VkDeviceSize
	type0 VkAccelerationStructureTypeKHR
	flags VkBuildAccelerationStructureFlagsKHR
	maxGeometryCount uint32
	pGeometryInfos *VkAccelerationStructureCreateGeometryTypeInfoKHR
	deviceAddress VkDeviceAddress
}
type VkAabbPositionsKHR struct {
	minX float32
	minY float32
	minZ float32
	maxX float32
	maxY float32
	maxZ float32
}
type VkAabbPositionsNV struct {
}
type VkTransformMatrixKHR struct {
	matrix float32
}
type VkTransformMatrixNV struct {
}
type VkAccelerationStructureInstanceKHR struct {
	transform VkTransformMatrixKHR
	instanceCustomIndex uint32
	mask uint32
	instanceShaderBindingTableRecordOffset uint32
	flags VkGeometryInstanceFlagsKHR
	accelerationStructureReference uint64
}
type VkAccelerationStructureInstanceNV struct {
}
type VkAccelerationStructureDeviceAddressInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	accelerationStructure VkAccelerationStructureKHR
}
type VkAccelerationStructureVersionKHR struct {
	sType VkStructureType
	pNext uintptr
	versionData *byte
}
type VkCopyAccelerationStructureInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	src VkAccelerationStructureKHR
	dst VkAccelerationStructureKHR
	mode VkCopyAccelerationStructureModeKHR
}
type VkCopyAccelerationStructureToMemoryInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	src VkAccelerationStructureKHR
	dst VkDeviceOrHostAddressKHR
	mode VkCopyAccelerationStructureModeKHR
}
type VkCopyMemoryToAccelerationStructureInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	src VkDeviceOrHostAddressConstKHR
	dst VkAccelerationStructureKHR
	mode VkCopyAccelerationStructureModeKHR
}
type VkRayTracingPipelineInterfaceCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	maxPayloadSize uint32
	maxAttributeSize uint32
	maxCallableSize uint32
}
type VkDeferredOperationInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	operationHandle VkDeferredOperationKHR
}
type VkPipelineLibraryCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	libraryCount uint32
	pLibraries *VkPipeline
}
type VkPhysicalDeviceExtendedDynamicStateFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	extendedDynamicState VkBool32
}
type VkRenderPassTransformBeginInfoQCOM struct {
	sType VkStructureType
	pNext uintptr
	transform VkSurfaceTransformFlagBitsKHR
}
type VkCommandBufferInheritanceRenderPassTransformInfoQCOM struct {
	sType VkStructureType
	pNext uintptr
	transform VkSurfaceTransformFlagBitsKHR
	renderArea VkRect2D
}
type VkPhysicalDeviceDiagnosticsConfigFeaturesNV struct {
	sType VkStructureType
	pNext uintptr
	diagnosticsConfig VkBool32
}
type VkDeviceDiagnosticsConfigCreateInfoNV struct {
	sType VkStructureType
	pNext uintptr
	flags VkDeviceDiagnosticsConfigFlagsNV
}
type VkPhysicalDeviceRobustness2FeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	robustBufferAccess2 VkBool32
	robustImageAccess2 VkBool32
	nullDescriptor VkBool32
}
type VkPhysicalDeviceRobustness2PropertiesEXT struct {
	sType VkStructureType
	pNext uintptr
	robustStorageBufferAccessSizeAlignment VkDeviceSize
	robustUniformBufferAccessSizeAlignment VkDeviceSize
}
type VkPhysicalDeviceImageRobustnessFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	robustImageAccess VkBool32
}
type VkPhysicalDevicePortabilitySubsetFeaturesKHR struct {
	sType VkStructureType
	pNext uintptr
	constantAlphaColorBlendFactors VkBool32
	events VkBool32
	imageViewFormatReinterpretation VkBool32
	imageViewFormatSwizzle VkBool32
	imageView2DOn3DImage VkBool32
	multisampleArrayImage VkBool32
	mutableComparisonSamplers VkBool32
	pointPolygons VkBool32
	samplerMipLodBias VkBool32
	separateStencilMaskRef VkBool32
	shaderSampleRateInterpolationFunctions VkBool32
	tessellationIsolines VkBool32
	tessellationPointMode VkBool32
	triangleFans VkBool32
	vertexAttributeAccessBeyondStride VkBool32
}
type VkPhysicalDevicePortabilitySubsetPropertiesKHR struct {
	sType VkStructureType
	pNext uintptr
	minVertexInputBindingStrideAlignment uint32
}
type VkPhysicalDevice4444FormatsFeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	formatA4R4G4B4 VkBool32
	formatA4B4G4R4 VkBool32
}
type VkBufferCopy2KHR struct {
	sType VkStructureType
	pNext uintptr
	srcOffset VkDeviceSize
	dstOffset VkDeviceSize
	size VkDeviceSize
}
type VkImageCopy2KHR struct {
	sType VkStructureType
	pNext uintptr
	srcSubresource VkImageSubresourceLayers
	srcOffset VkOffset3D
	dstSubresource VkImageSubresourceLayers
	dstOffset VkOffset3D
	extent VkExtent3D
}
type VkImageBlit2KHR struct {
	sType VkStructureType
	pNext uintptr
	srcSubresource VkImageSubresourceLayers
	srcOffsets VkOffset3D
	dstSubresource VkImageSubresourceLayers
	dstOffsets VkOffset3D
}
type VkBufferImageCopy2KHR struct {
	sType VkStructureType
	pNext uintptr
	bufferOffset VkDeviceSize
	bufferRowLength uint32
	bufferImageHeight uint32
	imageSubresource VkImageSubresourceLayers
	imageOffset VkOffset3D
	imageExtent VkExtent3D
}
type VkImageResolve2KHR struct {
	sType VkStructureType
	pNext uintptr
	srcSubresource VkImageSubresourceLayers
	srcOffset VkOffset3D
	dstSubresource VkImageSubresourceLayers
	dstOffset VkOffset3D
	extent VkExtent3D
}
type VkCopyBufferInfo2KHR struct {
	sType VkStructureType
	pNext uintptr
	srcBuffer VkBuffer
	dstBuffer VkBuffer
	regionCount uint32
	pRegions *VkBufferCopy2KHR
}
type VkCopyImageInfo2KHR struct {
	sType VkStructureType
	pNext uintptr
	srcImage VkImage
	srcImageLayout VkImageLayout
	dstImage VkImage
	dstImageLayout VkImageLayout
	regionCount uint32
	pRegions *VkImageCopy2KHR
}
type VkBlitImageInfo2KHR struct {
	sType VkStructureType
	pNext uintptr
	srcImage VkImage
	srcImageLayout VkImageLayout
	dstImage VkImage
	dstImageLayout VkImageLayout
	regionCount uint32
	pRegions *VkImageBlit2KHR
	filter VkFilter
}
type VkCopyBufferToImageInfo2KHR struct {
	sType VkStructureType
	pNext uintptr
	srcBuffer VkBuffer
	dstImage VkImage
	dstImageLayout VkImageLayout
	regionCount uint32
	pRegions *VkBufferImageCopy2KHR
}
type VkCopyImageToBufferInfo2KHR struct {
	sType VkStructureType
	pNext uintptr
	srcImage VkImage
	srcImageLayout VkImageLayout
	dstBuffer VkBuffer
	regionCount uint32
	pRegions *VkBufferImageCopy2KHR
}
type VkResolveImageInfo2KHR struct {
	sType VkStructureType
	pNext uintptr
	srcImage VkImage
	srcImageLayout VkImageLayout
	dstImage VkImage
	dstImageLayout VkImageLayout
	regionCount uint32
	pRegions *VkImageResolve2KHR
}
type VkPhysicalDeviceShaderImageAtomicInt64FeaturesEXT struct {
	sType VkStructureType
	pNext uintptr
	shaderImageInt64Atomics VkBool32
	sparseImageInt64Atomics VkBool32
}
type VkFragmentShadingRateAttachmentInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	pFragmentShadingRateAttachment *VkAttachmentReference2
	shadingRateAttachmentTexelSize VkExtent2D
}
type VkPipelineFragmentShadingRateStateCreateInfoKHR struct {
	sType VkStructureType
	pNext uintptr
	fragmentSize VkExtent2D
	combinerOps VkFragmentShadingRateCombinerOpKHR
}
type VkPhysicalDeviceFragmentShadingRateFeaturesKHR struct {
	sType VkStructureType
	pNext uintptr
	pipelineFragmentShadingRate VkBool32
	primitiveFragmentShadingRate VkBool32
	attachmentFragmentShadingRate VkBool32
}
type VkPhysicalDeviceFragmentShadingRatePropertiesKHR struct {
	sType VkStructureType
	pNext uintptr
	minFragmentShadingRateAttachmentTexelSize VkExtent2D
	maxFragmentShadingRateAttachmentTexelSize VkExtent2D
	maxFragmentShadingRateAttachmentTexelSizeAspectRatio uint32
	primitiveFragmentShadingRateWithMultipleViewports VkBool32
	layeredShadingRateAttachments VkBool32
	fragmentShadingRateNonTrivialCombinerOps VkBool32
	maxFragmentSize VkExtent2D
	maxFragmentSizeAspectRatio uint32
	maxFragmentShadingRateCoverageSamples uint32
	maxFragmentShadingRateRasterizationSamples VkSampleCountFlagBits
	fragmentShadingRateWithShaderDepthStencilWrites VkBool32
	fragmentShadingRateWithSampleMask VkBool32
	fragmentShadingRateWithShaderSampleMask VkBool32
	fragmentShadingRateWithConservativeRasterization VkBool32
	fragmentShadingRateWithFragmentShaderInterlock VkBool32
	fragmentShadingRateWithCustomSampleLocations VkBool32
	fragmentShadingRateStrictMultiplyCombiner VkBool32
}
type VkPhysicalDeviceFragmentShadingRateKHR struct {
	sType VkStructureType
	pNext uintptr
	sampleCounts VkSampleCountFlags
	fragmentSize VkExtent2D
}
type VkPhysicalDeviceShaderTerminateInvocationFeaturesKHR struct {
	sType VkStructureType
	pNext uintptr
	shaderTerminateInvocation VkBool32
}
//----------------------------------------
//--Enums
//----------------------------------------
const VK_MAX_PHYSICAL_DEVICE_NAME_SIZE  = 256
const VK_UUID_SIZE  = 16
const VK_LUID_SIZE  = 8
const VK_LUID_SIZE_KHR  = VK_LUID_SIZE
const VK_MAX_EXTENSION_NAME_SIZE  = 256
const VK_MAX_DESCRIPTION_SIZE  = 256
const VK_MAX_MEMORY_TYPES  = 32
const VK_MAX_MEMORY_HEAPS  = 16
const VK_LOD_CLAMP_NONE  = 1000.000000
const VK_REMAINING_MIP_LEVELS  = "(~0U)"
const VK_REMAINING_ARRAY_LAYERS  = "(~0U)"
const VK_WHOLE_SIZE  = "(~0ULL)"
const VK_ATTACHMENT_UNUSED  = "(~0U)"
const VK_TRUE  = 1
const VK_FALSE  = 0
const VK_QUEUE_FAMILY_IGNORED  = "(~0U)"
const VK_QUEUE_FAMILY_EXTERNAL  = "(~0U-1)"
const VK_QUEUE_FAMILY_EXTERNAL_KHR  = VK_QUEUE_FAMILY_EXTERNAL
const VK_QUEUE_FAMILY_FOREIGN_EXT  = "(~0U-2)"
const VK_SUBPASS_EXTERNAL  = "(~0U)"
const VK_MAX_DEVICE_GROUP_SIZE  = 32
const VK_MAX_DEVICE_GROUP_SIZE_KHR  = VK_MAX_DEVICE_GROUP_SIZE
const VK_MAX_DRIVER_NAME_SIZE  = 256
const VK_MAX_DRIVER_NAME_SIZE_KHR  = VK_MAX_DRIVER_NAME_SIZE
const VK_MAX_DRIVER_INFO_SIZE  = 256
const VK_MAX_DRIVER_INFO_SIZE_KHR  = VK_MAX_DRIVER_INFO_SIZE
const VK_SHADER_UNUSED_KHR  = "(~0U)"
const VK_SHADER_UNUSED_NV  = VK_SHADER_UNUSED_KHR
type VkImageLayout = int32
const VK_IMAGE_LAYOUT_UNDEFINED VkImageLayout = 0
const VK_IMAGE_LAYOUT_GENERAL VkImageLayout = 1
const VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL VkImageLayout = 2
const VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL VkImageLayout = 3
const VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL VkImageLayout = 4
const VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL VkImageLayout = 5
const VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL VkImageLayout = 6
const VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL VkImageLayout = 7
const VK_IMAGE_LAYOUT_PREINITIALIZED VkImageLayout = 8
type VkAttachmentLoadOp = int32
const VK_ATTACHMENT_LOAD_OP_LOAD VkAttachmentLoadOp = 0
const VK_ATTACHMENT_LOAD_OP_CLEAR VkAttachmentLoadOp = 1
const VK_ATTACHMENT_LOAD_OP_DONT_CARE VkAttachmentLoadOp = 2
type VkAttachmentStoreOp = int32
const VK_ATTACHMENT_STORE_OP_STORE VkAttachmentStoreOp = 0
const VK_ATTACHMENT_STORE_OP_DONT_CARE VkAttachmentStoreOp = 1
type VkImageType = int32
const VK_IMAGE_TYPE_1D VkImageType = 0
const VK_IMAGE_TYPE_2D VkImageType = 1
const VK_IMAGE_TYPE_3D VkImageType = 2
type VkImageTiling = int32
const VK_IMAGE_TILING_OPTIMAL VkImageTiling = 0
const VK_IMAGE_TILING_LINEAR VkImageTiling = 1
type VkImageViewType = int32
const VK_IMAGE_VIEW_TYPE_1D VkImageViewType = 0
const VK_IMAGE_VIEW_TYPE_2D VkImageViewType = 1
const VK_IMAGE_VIEW_TYPE_3D VkImageViewType = 2
const VK_IMAGE_VIEW_TYPE_CUBE VkImageViewType = 3
const VK_IMAGE_VIEW_TYPE_1D_ARRAY VkImageViewType = 4
const VK_IMAGE_VIEW_TYPE_2D_ARRAY VkImageViewType = 5
const VK_IMAGE_VIEW_TYPE_CUBE_ARRAY VkImageViewType = 6
type VkCommandBufferLevel = int32
const VK_COMMAND_BUFFER_LEVEL_PRIMARY VkCommandBufferLevel = 0
const VK_COMMAND_BUFFER_LEVEL_SECONDARY VkCommandBufferLevel = 1
type VkComponentSwizzle = int32
const VK_COMPONENT_SWIZZLE_IDENTITY VkComponentSwizzle = 0
const VK_COMPONENT_SWIZZLE_ZERO VkComponentSwizzle = 1
const VK_COMPONENT_SWIZZLE_ONE VkComponentSwizzle = 2
const VK_COMPONENT_SWIZZLE_R VkComponentSwizzle = 3
const VK_COMPONENT_SWIZZLE_G VkComponentSwizzle = 4
const VK_COMPONENT_SWIZZLE_B VkComponentSwizzle = 5
const VK_COMPONENT_SWIZZLE_A VkComponentSwizzle = 6
type VkDescriptorType = int32
const VK_DESCRIPTOR_TYPE_SAMPLER VkDescriptorType = 0
const VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER VkDescriptorType = 1
const VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE VkDescriptorType = 2
const VK_DESCRIPTOR_TYPE_STORAGE_IMAGE VkDescriptorType = 3
const VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER VkDescriptorType = 4
const VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER VkDescriptorType = 5
const VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER VkDescriptorType = 6
const VK_DESCRIPTOR_TYPE_STORAGE_BUFFER VkDescriptorType = 7
const VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC VkDescriptorType = 8
const VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC VkDescriptorType = 9
const VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT VkDescriptorType = 10
type VkQueryType = int32
const VK_QUERY_TYPE_OCCLUSION VkQueryType = 0
const VK_QUERY_TYPE_PIPELINE_STATISTICS VkQueryType = 1
const VK_QUERY_TYPE_TIMESTAMP VkQueryType = 2
type VkBorderColor = int32
const VK_BORDER_COLOR_FLOAT_TRANSPARENT_BLACK VkBorderColor = 0
const VK_BORDER_COLOR_INT_TRANSPARENT_BLACK VkBorderColor = 1
const VK_BORDER_COLOR_FLOAT_OPAQUE_BLACK VkBorderColor = 2
const VK_BORDER_COLOR_INT_OPAQUE_BLACK VkBorderColor = 3
const VK_BORDER_COLOR_FLOAT_OPAQUE_WHITE VkBorderColor = 4
const VK_BORDER_COLOR_INT_OPAQUE_WHITE VkBorderColor = 5
type VkPipelineBindPoint = int32
const VK_PIPELINE_BIND_POINT_GRAPHICS VkPipelineBindPoint = 0
const VK_PIPELINE_BIND_POINT_COMPUTE VkPipelineBindPoint = 1
type VkPipelineCacheHeaderVersion = int32
const VK_PIPELINE_CACHE_HEADER_VERSION_ONE VkPipelineCacheHeaderVersion = 1
type VkPipelineCacheCreateFlagBits = int64
type VkPrimitiveTopology = int32
const VK_PRIMITIVE_TOPOLOGY_POINT_LIST VkPrimitiveTopology = 0
const VK_PRIMITIVE_TOPOLOGY_LINE_LIST VkPrimitiveTopology = 1
const VK_PRIMITIVE_TOPOLOGY_LINE_STRIP VkPrimitiveTopology = 2
const VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST VkPrimitiveTopology = 3
const VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP VkPrimitiveTopology = 4
const VK_PRIMITIVE_TOPOLOGY_TRIANGLE_FAN VkPrimitiveTopology = 5
const VK_PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY VkPrimitiveTopology = 6
const VK_PRIMITIVE_TOPOLOGY_LINE_STRIP_WITH_ADJACENCY VkPrimitiveTopology = 7
const VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY VkPrimitiveTopology = 8
const VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP_WITH_ADJACENCY VkPrimitiveTopology = 9
const VK_PRIMITIVE_TOPOLOGY_PATCH_LIST VkPrimitiveTopology = 10
type VkSharingMode = int32
const VK_SHARING_MODE_EXCLUSIVE VkSharingMode = 0
const VK_SHARING_MODE_CONCURRENT VkSharingMode = 1
type VkIndexType = int32
const VK_INDEX_TYPE_UINT16 VkIndexType = 0
const VK_INDEX_TYPE_UINT32 VkIndexType = 1
type VkFilter = int32
const VK_FILTER_NEAREST VkFilter = 0
const VK_FILTER_LINEAR VkFilter = 1
type VkSamplerMipmapMode = int32
const VK_SAMPLER_MIPMAP_MODE_NEAREST VkSamplerMipmapMode = 0
const VK_SAMPLER_MIPMAP_MODE_LINEAR VkSamplerMipmapMode = 1
type VkSamplerAddressMode = int32
const VK_SAMPLER_ADDRESS_MODE_REPEAT VkSamplerAddressMode = 0
const VK_SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT VkSamplerAddressMode = 1
const VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE VkSamplerAddressMode = 2
const VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER VkSamplerAddressMode = 3
type VkCompareOp = int32
const VK_COMPARE_OP_NEVER VkCompareOp = 0
const VK_COMPARE_OP_LESS VkCompareOp = 1
const VK_COMPARE_OP_EQUAL VkCompareOp = 2
const VK_COMPARE_OP_LESS_OR_EQUAL VkCompareOp = 3
const VK_COMPARE_OP_GREATER VkCompareOp = 4
const VK_COMPARE_OP_NOT_EQUAL VkCompareOp = 5
const VK_COMPARE_OP_GREATER_OR_EQUAL VkCompareOp = 6
const VK_COMPARE_OP_ALWAYS VkCompareOp = 7
type VkPolygonMode = int32
const VK_POLYGON_MODE_FILL VkPolygonMode = 0
const VK_POLYGON_MODE_LINE VkPolygonMode = 1
const VK_POLYGON_MODE_POINT VkPolygonMode = 2
type VkFrontFace = int32
const VK_FRONT_FACE_COUNTER_CLOCKWISE VkFrontFace = 0
const VK_FRONT_FACE_CLOCKWISE VkFrontFace = 1
type VkBlendFactor = int32
const VK_BLEND_FACTOR_ZERO VkBlendFactor = 0
const VK_BLEND_FACTOR_ONE VkBlendFactor = 1
const VK_BLEND_FACTOR_SRC_COLOR VkBlendFactor = 2
const VK_BLEND_FACTOR_ONE_MINUS_SRC_COLOR VkBlendFactor = 3
const VK_BLEND_FACTOR_DST_COLOR VkBlendFactor = 4
const VK_BLEND_FACTOR_ONE_MINUS_DST_COLOR VkBlendFactor = 5
const VK_BLEND_FACTOR_SRC_ALPHA VkBlendFactor = 6
const VK_BLEND_FACTOR_ONE_MINUS_SRC_ALPHA VkBlendFactor = 7
const VK_BLEND_FACTOR_DST_ALPHA VkBlendFactor = 8
const VK_BLEND_FACTOR_ONE_MINUS_DST_ALPHA VkBlendFactor = 9
const VK_BLEND_FACTOR_CONSTANT_COLOR VkBlendFactor = 10
const VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR VkBlendFactor = 11
const VK_BLEND_FACTOR_CONSTANT_ALPHA VkBlendFactor = 12
const VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA VkBlendFactor = 13
const VK_BLEND_FACTOR_SRC_ALPHA_SATURATE VkBlendFactor = 14
const VK_BLEND_FACTOR_SRC1_COLOR VkBlendFactor = 15
const VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR VkBlendFactor = 16
const VK_BLEND_FACTOR_SRC1_ALPHA VkBlendFactor = 17
const VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA VkBlendFactor = 18
type VkBlendOp = int32
const VK_BLEND_OP_ADD VkBlendOp = 0
const VK_BLEND_OP_SUBTRACT VkBlendOp = 1
const VK_BLEND_OP_REVERSE_SUBTRACT VkBlendOp = 2
const VK_BLEND_OP_MIN VkBlendOp = 3
const VK_BLEND_OP_MAX VkBlendOp = 4
type VkStencilOp = int32
const VK_STENCIL_OP_KEEP VkStencilOp = 0
const VK_STENCIL_OP_ZERO VkStencilOp = 1
const VK_STENCIL_OP_REPLACE VkStencilOp = 2
const VK_STENCIL_OP_INCREMENT_AND_CLAMP VkStencilOp = 3
const VK_STENCIL_OP_DECREMENT_AND_CLAMP VkStencilOp = 4
const VK_STENCIL_OP_INVERT VkStencilOp = 5
const VK_STENCIL_OP_INCREMENT_AND_WRAP VkStencilOp = 6
const VK_STENCIL_OP_DECREMENT_AND_WRAP VkStencilOp = 7
type VkLogicOp = int32
const VK_LOGIC_OP_CLEAR VkLogicOp = 0
const VK_LOGIC_OP_AND VkLogicOp = 1
const VK_LOGIC_OP_AND_REVERSE VkLogicOp = 2
const VK_LOGIC_OP_COPY VkLogicOp = 3
const VK_LOGIC_OP_AND_INVERTED VkLogicOp = 4
const VK_LOGIC_OP_NO_OP VkLogicOp = 5
const VK_LOGIC_OP_XOR VkLogicOp = 6
const VK_LOGIC_OP_OR VkLogicOp = 7
const VK_LOGIC_OP_NOR VkLogicOp = 8
const VK_LOGIC_OP_EQUIVALENT VkLogicOp = 9
const VK_LOGIC_OP_INVERT VkLogicOp = 10
const VK_LOGIC_OP_OR_REVERSE VkLogicOp = 11
const VK_LOGIC_OP_COPY_INVERTED VkLogicOp = 12
const VK_LOGIC_OP_OR_INVERTED VkLogicOp = 13
const VK_LOGIC_OP_NAND VkLogicOp = 14
const VK_LOGIC_OP_SET VkLogicOp = 15
type VkInternalAllocationType = int32
const VK_INTERNAL_ALLOCATION_TYPE_EXECUTABLE VkInternalAllocationType = 0
type VkSystemAllocationScope = int32
const VK_SYSTEM_ALLOCATION_SCOPE_COMMAND VkSystemAllocationScope = 0
const VK_SYSTEM_ALLOCATION_SCOPE_OBJECT VkSystemAllocationScope = 1
const VK_SYSTEM_ALLOCATION_SCOPE_CACHE VkSystemAllocationScope = 2
const VK_SYSTEM_ALLOCATION_SCOPE_DEVICE VkSystemAllocationScope = 3
const VK_SYSTEM_ALLOCATION_SCOPE_INSTANCE VkSystemAllocationScope = 4
type VkPhysicalDeviceType = int32
const VK_PHYSICAL_DEVICE_TYPE_OTHER VkPhysicalDeviceType = 0
const VK_PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU VkPhysicalDeviceType = 1
const VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU VkPhysicalDeviceType = 2
const VK_PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU VkPhysicalDeviceType = 3
const VK_PHYSICAL_DEVICE_TYPE_CPU VkPhysicalDeviceType = 4
type VkVertexInputRate = int32
const VK_VERTEX_INPUT_RATE_VERTEX VkVertexInputRate = 0
const VK_VERTEX_INPUT_RATE_INSTANCE VkVertexInputRate = 1
type VkFormat = int32
const VK_FORMAT_UNDEFINED VkFormat = 0
const VK_FORMAT_R4G4_UNORM_PACK8 VkFormat = 1
const VK_FORMAT_R4G4B4A4_UNORM_PACK16 VkFormat = 2
const VK_FORMAT_B4G4R4A4_UNORM_PACK16 VkFormat = 3
const VK_FORMAT_R5G6B5_UNORM_PACK16 VkFormat = 4
const VK_FORMAT_B5G6R5_UNORM_PACK16 VkFormat = 5
const VK_FORMAT_R5G5B5A1_UNORM_PACK16 VkFormat = 6
const VK_FORMAT_B5G5R5A1_UNORM_PACK16 VkFormat = 7
const VK_FORMAT_A1R5G5B5_UNORM_PACK16 VkFormat = 8
const VK_FORMAT_R8_UNORM VkFormat = 9
const VK_FORMAT_R8_SNORM VkFormat = 10
const VK_FORMAT_R8_USCALED VkFormat = 11
const VK_FORMAT_R8_SSCALED VkFormat = 12
const VK_FORMAT_R8_UINT VkFormat = 13
const VK_FORMAT_R8_SINT VkFormat = 14
const VK_FORMAT_R8_SRGB VkFormat = 15
const VK_FORMAT_R8G8_UNORM VkFormat = 16
const VK_FORMAT_R8G8_SNORM VkFormat = 17
const VK_FORMAT_R8G8_USCALED VkFormat = 18
const VK_FORMAT_R8G8_SSCALED VkFormat = 19
const VK_FORMAT_R8G8_UINT VkFormat = 20
const VK_FORMAT_R8G8_SINT VkFormat = 21
const VK_FORMAT_R8G8_SRGB VkFormat = 22
const VK_FORMAT_R8G8B8_UNORM VkFormat = 23
const VK_FORMAT_R8G8B8_SNORM VkFormat = 24
const VK_FORMAT_R8G8B8_USCALED VkFormat = 25
const VK_FORMAT_R8G8B8_SSCALED VkFormat = 26
const VK_FORMAT_R8G8B8_UINT VkFormat = 27
const VK_FORMAT_R8G8B8_SINT VkFormat = 28
const VK_FORMAT_R8G8B8_SRGB VkFormat = 29
const VK_FORMAT_B8G8R8_UNORM VkFormat = 30
const VK_FORMAT_B8G8R8_SNORM VkFormat = 31
const VK_FORMAT_B8G8R8_USCALED VkFormat = 32
const VK_FORMAT_B8G8R8_SSCALED VkFormat = 33
const VK_FORMAT_B8G8R8_UINT VkFormat = 34
const VK_FORMAT_B8G8R8_SINT VkFormat = 35
const VK_FORMAT_B8G8R8_SRGB VkFormat = 36
const VK_FORMAT_R8G8B8A8_UNORM VkFormat = 37
const VK_FORMAT_R8G8B8A8_SNORM VkFormat = 38
const VK_FORMAT_R8G8B8A8_USCALED VkFormat = 39
const VK_FORMAT_R8G8B8A8_SSCALED VkFormat = 40
const VK_FORMAT_R8G8B8A8_UINT VkFormat = 41
const VK_FORMAT_R8G8B8A8_SINT VkFormat = 42
const VK_FORMAT_R8G8B8A8_SRGB VkFormat = 43
const VK_FORMAT_B8G8R8A8_UNORM VkFormat = 44
const VK_FORMAT_B8G8R8A8_SNORM VkFormat = 45
const VK_FORMAT_B8G8R8A8_USCALED VkFormat = 46
const VK_FORMAT_B8G8R8A8_SSCALED VkFormat = 47
const VK_FORMAT_B8G8R8A8_UINT VkFormat = 48
const VK_FORMAT_B8G8R8A8_SINT VkFormat = 49
const VK_FORMAT_B8G8R8A8_SRGB VkFormat = 50
const VK_FORMAT_A8B8G8R8_UNORM_PACK32 VkFormat = 51
const VK_FORMAT_A8B8G8R8_SNORM_PACK32 VkFormat = 52
const VK_FORMAT_A8B8G8R8_USCALED_PACK32 VkFormat = 53
const VK_FORMAT_A8B8G8R8_SSCALED_PACK32 VkFormat = 54
const VK_FORMAT_A8B8G8R8_UINT_PACK32 VkFormat = 55
const VK_FORMAT_A8B8G8R8_SINT_PACK32 VkFormat = 56
const VK_FORMAT_A8B8G8R8_SRGB_PACK32 VkFormat = 57
const VK_FORMAT_A2R10G10B10_UNORM_PACK32 VkFormat = 58
const VK_FORMAT_A2R10G10B10_SNORM_PACK32 VkFormat = 59
const VK_FORMAT_A2R10G10B10_USCALED_PACK32 VkFormat = 60
const VK_FORMAT_A2R10G10B10_SSCALED_PACK32 VkFormat = 61
const VK_FORMAT_A2R10G10B10_UINT_PACK32 VkFormat = 62
const VK_FORMAT_A2R10G10B10_SINT_PACK32 VkFormat = 63
const VK_FORMAT_A2B10G10R10_UNORM_PACK32 VkFormat = 64
const VK_FORMAT_A2B10G10R10_SNORM_PACK32 VkFormat = 65
const VK_FORMAT_A2B10G10R10_USCALED_PACK32 VkFormat = 66
const VK_FORMAT_A2B10G10R10_SSCALED_PACK32 VkFormat = 67
const VK_FORMAT_A2B10G10R10_UINT_PACK32 VkFormat = 68
const VK_FORMAT_A2B10G10R10_SINT_PACK32 VkFormat = 69
const VK_FORMAT_R16_UNORM VkFormat = 70
const VK_FORMAT_R16_SNORM VkFormat = 71
const VK_FORMAT_R16_USCALED VkFormat = 72
const VK_FORMAT_R16_SSCALED VkFormat = 73
const VK_FORMAT_R16_UINT VkFormat = 74
const VK_FORMAT_R16_SINT VkFormat = 75
const VK_FORMAT_R16_SFLOAT VkFormat = 76
const VK_FORMAT_R16G16_UNORM VkFormat = 77
const VK_FORMAT_R16G16_SNORM VkFormat = 78
const VK_FORMAT_R16G16_USCALED VkFormat = 79
const VK_FORMAT_R16G16_SSCALED VkFormat = 80
const VK_FORMAT_R16G16_UINT VkFormat = 81
const VK_FORMAT_R16G16_SINT VkFormat = 82
const VK_FORMAT_R16G16_SFLOAT VkFormat = 83
const VK_FORMAT_R16G16B16_UNORM VkFormat = 84
const VK_FORMAT_R16G16B16_SNORM VkFormat = 85
const VK_FORMAT_R16G16B16_USCALED VkFormat = 86
const VK_FORMAT_R16G16B16_SSCALED VkFormat = 87
const VK_FORMAT_R16G16B16_UINT VkFormat = 88
const VK_FORMAT_R16G16B16_SINT VkFormat = 89
const VK_FORMAT_R16G16B16_SFLOAT VkFormat = 90
const VK_FORMAT_R16G16B16A16_UNORM VkFormat = 91
const VK_FORMAT_R16G16B16A16_SNORM VkFormat = 92
const VK_FORMAT_R16G16B16A16_USCALED VkFormat = 93
const VK_FORMAT_R16G16B16A16_SSCALED VkFormat = 94
const VK_FORMAT_R16G16B16A16_UINT VkFormat = 95
const VK_FORMAT_R16G16B16A16_SINT VkFormat = 96
const VK_FORMAT_R16G16B16A16_SFLOAT VkFormat = 97
const VK_FORMAT_R32_UINT VkFormat = 98
const VK_FORMAT_R32_SINT VkFormat = 99
const VK_FORMAT_R32_SFLOAT VkFormat = 100
const VK_FORMAT_R32G32_UINT VkFormat = 101
const VK_FORMAT_R32G32_SINT VkFormat = 102
const VK_FORMAT_R32G32_SFLOAT VkFormat = 103
const VK_FORMAT_R32G32B32_UINT VkFormat = 104
const VK_FORMAT_R32G32B32_SINT VkFormat = 105
const VK_FORMAT_R32G32B32_SFLOAT VkFormat = 106
const VK_FORMAT_R32G32B32A32_UINT VkFormat = 107
const VK_FORMAT_R32G32B32A32_SINT VkFormat = 108
const VK_FORMAT_R32G32B32A32_SFLOAT VkFormat = 109
const VK_FORMAT_R64_UINT VkFormat = 110
const VK_FORMAT_R64_SINT VkFormat = 111
const VK_FORMAT_R64_SFLOAT VkFormat = 112
const VK_FORMAT_R64G64_UINT VkFormat = 113
const VK_FORMAT_R64G64_SINT VkFormat = 114
const VK_FORMAT_R64G64_SFLOAT VkFormat = 115
const VK_FORMAT_R64G64B64_UINT VkFormat = 116
const VK_FORMAT_R64G64B64_SINT VkFormat = 117
const VK_FORMAT_R64G64B64_SFLOAT VkFormat = 118
const VK_FORMAT_R64G64B64A64_UINT VkFormat = 119
const VK_FORMAT_R64G64B64A64_SINT VkFormat = 120
const VK_FORMAT_R64G64B64A64_SFLOAT VkFormat = 121
const VK_FORMAT_B10G11R11_UFLOAT_PACK32 VkFormat = 122
const VK_FORMAT_E5B9G9R9_UFLOAT_PACK32 VkFormat = 123
const VK_FORMAT_D16_UNORM VkFormat = 124
const VK_FORMAT_X8_D24_UNORM_PACK32 VkFormat = 125
const VK_FORMAT_D32_SFLOAT VkFormat = 126
const VK_FORMAT_S8_UINT VkFormat = 127
const VK_FORMAT_D16_UNORM_S8_UINT VkFormat = 128
const VK_FORMAT_D24_UNORM_S8_UINT VkFormat = 129
const VK_FORMAT_D32_SFLOAT_S8_UINT VkFormat = 130
const VK_FORMAT_BC1_RGB_UNORM_BLOCK VkFormat = 131
const VK_FORMAT_BC1_RGB_SRGB_BLOCK VkFormat = 132
const VK_FORMAT_BC1_RGBA_UNORM_BLOCK VkFormat = 133
const VK_FORMAT_BC1_RGBA_SRGB_BLOCK VkFormat = 134
const VK_FORMAT_BC2_UNORM_BLOCK VkFormat = 135
const VK_FORMAT_BC2_SRGB_BLOCK VkFormat = 136
const VK_FORMAT_BC3_UNORM_BLOCK VkFormat = 137
const VK_FORMAT_BC3_SRGB_BLOCK VkFormat = 138
const VK_FORMAT_BC4_UNORM_BLOCK VkFormat = 139
const VK_FORMAT_BC4_SNORM_BLOCK VkFormat = 140
const VK_FORMAT_BC5_UNORM_BLOCK VkFormat = 141
const VK_FORMAT_BC5_SNORM_BLOCK VkFormat = 142
const VK_FORMAT_BC6H_UFLOAT_BLOCK VkFormat = 143
const VK_FORMAT_BC6H_SFLOAT_BLOCK VkFormat = 144
const VK_FORMAT_BC7_UNORM_BLOCK VkFormat = 145
const VK_FORMAT_BC7_SRGB_BLOCK VkFormat = 146
const VK_FORMAT_ETC2_R8G8B8_UNORM_BLOCK VkFormat = 147
const VK_FORMAT_ETC2_R8G8B8_SRGB_BLOCK VkFormat = 148
const VK_FORMAT_ETC2_R8G8B8A1_UNORM_BLOCK VkFormat = 149
const VK_FORMAT_ETC2_R8G8B8A1_SRGB_BLOCK VkFormat = 150
const VK_FORMAT_ETC2_R8G8B8A8_UNORM_BLOCK VkFormat = 151
const VK_FORMAT_ETC2_R8G8B8A8_SRGB_BLOCK VkFormat = 152
const VK_FORMAT_EAC_R11_UNORM_BLOCK VkFormat = 153
const VK_FORMAT_EAC_R11_SNORM_BLOCK VkFormat = 154
const VK_FORMAT_EAC_R11G11_UNORM_BLOCK VkFormat = 155
const VK_FORMAT_EAC_R11G11_SNORM_BLOCK VkFormat = 156
const VK_FORMAT_ASTC_4x4_UNORM_BLOCK VkFormat = 157
const VK_FORMAT_ASTC_4x4_SRGB_BLOCK VkFormat = 158
const VK_FORMAT_ASTC_5x4_UNORM_BLOCK VkFormat = 159
const VK_FORMAT_ASTC_5x4_SRGB_BLOCK VkFormat = 160
const VK_FORMAT_ASTC_5x5_UNORM_BLOCK VkFormat = 161
const VK_FORMAT_ASTC_5x5_SRGB_BLOCK VkFormat = 162
const VK_FORMAT_ASTC_6x5_UNORM_BLOCK VkFormat = 163
const VK_FORMAT_ASTC_6x5_SRGB_BLOCK VkFormat = 164
const VK_FORMAT_ASTC_6x6_UNORM_BLOCK VkFormat = 165
const VK_FORMAT_ASTC_6x6_SRGB_BLOCK VkFormat = 166
const VK_FORMAT_ASTC_8x5_UNORM_BLOCK VkFormat = 167
const VK_FORMAT_ASTC_8x5_SRGB_BLOCK VkFormat = 168
const VK_FORMAT_ASTC_8x6_UNORM_BLOCK VkFormat = 169
const VK_FORMAT_ASTC_8x6_SRGB_BLOCK VkFormat = 170
const VK_FORMAT_ASTC_8x8_UNORM_BLOCK VkFormat = 171
const VK_FORMAT_ASTC_8x8_SRGB_BLOCK VkFormat = 172
const VK_FORMAT_ASTC_10x5_UNORM_BLOCK VkFormat = 173
const VK_FORMAT_ASTC_10x5_SRGB_BLOCK VkFormat = 174
const VK_FORMAT_ASTC_10x6_UNORM_BLOCK VkFormat = 175
const VK_FORMAT_ASTC_10x6_SRGB_BLOCK VkFormat = 176
const VK_FORMAT_ASTC_10x8_UNORM_BLOCK VkFormat = 177
const VK_FORMAT_ASTC_10x8_SRGB_BLOCK VkFormat = 178
const VK_FORMAT_ASTC_10x10_UNORM_BLOCK VkFormat = 179
const VK_FORMAT_ASTC_10x10_SRGB_BLOCK VkFormat = 180
const VK_FORMAT_ASTC_12x10_UNORM_BLOCK VkFormat = 181
const VK_FORMAT_ASTC_12x10_SRGB_BLOCK VkFormat = 182
const VK_FORMAT_ASTC_12x12_UNORM_BLOCK VkFormat = 183
const VK_FORMAT_ASTC_12x12_SRGB_BLOCK VkFormat = 184
type VkStructureType = int32
const VK_STRUCTURE_TYPE_APPLICATION_INFO VkStructureType = 0
const VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO VkStructureType = 1
const VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO VkStructureType = 2
const VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO VkStructureType = 3
const VK_STRUCTURE_TYPE_SUBMIT_INFO VkStructureType = 4
const VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO VkStructureType = 5
const VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE VkStructureType = 6
const VK_STRUCTURE_TYPE_BIND_SPARSE_INFO VkStructureType = 7
const VK_STRUCTURE_TYPE_FENCE_CREATE_INFO VkStructureType = 8
const VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO VkStructureType = 9
const VK_STRUCTURE_TYPE_EVENT_CREATE_INFO VkStructureType = 10
const VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO VkStructureType = 11
const VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO VkStructureType = 12
const VK_STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO VkStructureType = 13
const VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO VkStructureType = 14
const VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO VkStructureType = 15
const VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO VkStructureType = 16
const VK_STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO VkStructureType = 17
const VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO VkStructureType = 18
const VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_STATE_CREATE_INFO VkStructureType = 19
const VK_STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO VkStructureType = 20
const VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO VkStructureType = 21
const VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO VkStructureType = 22
const VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO VkStructureType = 23
const VK_STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO VkStructureType = 24
const VK_STRUCTURE_TYPE_PIPELINE_DEPTH_STENCIL_STATE_CREATE_INFO VkStructureType = 25
const VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO VkStructureType = 26
const VK_STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO VkStructureType = 27
const VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO VkStructureType = 28
const VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO VkStructureType = 29
const VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO VkStructureType = 30
const VK_STRUCTURE_TYPE_SAMPLER_CREATE_INFO VkStructureType = 31
const VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO VkStructureType = 32
const VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO VkStructureType = 33
const VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO VkStructureType = 34
const VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET VkStructureType = 35
const VK_STRUCTURE_TYPE_COPY_DESCRIPTOR_SET VkStructureType = 36
const VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO VkStructureType = 37
const VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO VkStructureType = 38
const VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO VkStructureType = 39
const VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO VkStructureType = 40
const VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO VkStructureType = 41
const VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO VkStructureType = 42
const VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO VkStructureType = 43
const VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER VkStructureType = 44
const VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER VkStructureType = 45
const VK_STRUCTURE_TYPE_MEMORY_BARRIER VkStructureType = 46
const VK_STRUCTURE_TYPE_LOADER_INSTANCE_CREATE_INFO VkStructureType = 47
const VK_STRUCTURE_TYPE_LOADER_DEVICE_CREATE_INFO VkStructureType = 48
type VkSubpassContents = int32
const VK_SUBPASS_CONTENTS_INLINE VkSubpassContents = 0
const VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS VkSubpassContents = 1
type VkResult = int32
const VK_SUCCESS VkResult = 0
const VK_NOT_READY VkResult = 1
const VK_TIMEOUT VkResult = 2
const VK_EVENT_SET VkResult = 3
const VK_EVENT_RESET VkResult = 4
const VK_INCOMPLETE VkResult = 5
const VK_ERROR_OUT_OF_HOST_MEMORY VkResult = -1
const VK_ERROR_OUT_OF_DEVICE_MEMORY VkResult = -2
const VK_ERROR_INITIALIZATION_FAILED VkResult = -3
const VK_ERROR_DEVICE_LOST VkResult = -4
const VK_ERROR_MEMORY_MAP_FAILED VkResult = -5
const VK_ERROR_LAYER_NOT_PRESENT VkResult = -6
const VK_ERROR_EXTENSION_NOT_PRESENT VkResult = -7
const VK_ERROR_FEATURE_NOT_PRESENT VkResult = -8
const VK_ERROR_INCOMPATIBLE_DRIVER VkResult = -9
const VK_ERROR_TOO_MANY_OBJECTS VkResult = -10
const VK_ERROR_FORMAT_NOT_SUPPORTED VkResult = -11
const VK_ERROR_FRAGMENTED_POOL VkResult = -12
const VK_ERROR_UNKNOWN VkResult = -13
type VkDynamicState = int32
const VK_DYNAMIC_STATE_VIEWPORT VkDynamicState = 0
const VK_DYNAMIC_STATE_SCISSOR VkDynamicState = 1
const VK_DYNAMIC_STATE_LINE_WIDTH VkDynamicState = 2
const VK_DYNAMIC_STATE_DEPTH_BIAS VkDynamicState = 3
const VK_DYNAMIC_STATE_BLEND_CONSTANTS VkDynamicState = 4
const VK_DYNAMIC_STATE_DEPTH_BOUNDS VkDynamicState = 5
const VK_DYNAMIC_STATE_STENCIL_COMPARE_MASK VkDynamicState = 6
const VK_DYNAMIC_STATE_STENCIL_WRITE_MASK VkDynamicState = 7
const VK_DYNAMIC_STATE_STENCIL_REFERENCE VkDynamicState = 8
type VkDescriptorUpdateTemplateType = int32
const VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET VkDescriptorUpdateTemplateType = 0
type VkObjectType = int32
const VK_OBJECT_TYPE_UNKNOWN VkObjectType = 0
const VK_OBJECT_TYPE_INSTANCE VkObjectType = 1
const VK_OBJECT_TYPE_PHYSICAL_DEVICE VkObjectType = 2
const VK_OBJECT_TYPE_DEVICE VkObjectType = 3
const VK_OBJECT_TYPE_QUEUE VkObjectType = 4
const VK_OBJECT_TYPE_SEMAPHORE VkObjectType = 5
const VK_OBJECT_TYPE_COMMAND_BUFFER VkObjectType = 6
const VK_OBJECT_TYPE_FENCE VkObjectType = 7
const VK_OBJECT_TYPE_DEVICE_MEMORY VkObjectType = 8
const VK_OBJECT_TYPE_BUFFER VkObjectType = 9
const VK_OBJECT_TYPE_IMAGE VkObjectType = 10
const VK_OBJECT_TYPE_EVENT VkObjectType = 11
const VK_OBJECT_TYPE_QUERY_POOL VkObjectType = 12
const VK_OBJECT_TYPE_BUFFER_VIEW VkObjectType = 13
const VK_OBJECT_TYPE_IMAGE_VIEW VkObjectType = 14
const VK_OBJECT_TYPE_SHADER_MODULE VkObjectType = 15
const VK_OBJECT_TYPE_PIPELINE_CACHE VkObjectType = 16
const VK_OBJECT_TYPE_PIPELINE_LAYOUT VkObjectType = 17
const VK_OBJECT_TYPE_RENDER_PASS VkObjectType = 18
const VK_OBJECT_TYPE_PIPELINE VkObjectType = 19
const VK_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT VkObjectType = 20
const VK_OBJECT_TYPE_SAMPLER VkObjectType = 21
const VK_OBJECT_TYPE_DESCRIPTOR_POOL VkObjectType = 22
const VK_OBJECT_TYPE_DESCRIPTOR_SET VkObjectType = 23
const VK_OBJECT_TYPE_FRAMEBUFFER VkObjectType = 24
const VK_OBJECT_TYPE_COMMAND_POOL VkObjectType = 25
type VkQueueFlagBits = int64
const VK_QUEUE_GRAPHICS_BIT VkQueueFlagBits = 0x1 << 0
const VK_QUEUE_COMPUTE_BIT VkQueueFlagBits = 0x1 << 1
const VK_QUEUE_TRANSFER_BIT VkQueueFlagBits = 0x1 << 2
const VK_QUEUE_SPARSE_BINDING_BIT VkQueueFlagBits = 0x1 << 3
type VkCullModeFlagBits = int64
const VK_CULL_MODE_NONE VkCullModeFlagBits = 0
const VK_CULL_MODE_FRONT_BIT VkCullModeFlagBits = 0x1 << 0
const VK_CULL_MODE_BACK_BIT VkCullModeFlagBits = 0x1 << 1
const VK_CULL_MODE_FRONT_AND_BACK VkCullModeFlagBits = 0x00000003
type VkRenderPassCreateFlagBits = int64
type VkDeviceQueueCreateFlagBits = int64
type VkMemoryPropertyFlagBits = int64
const VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT VkMemoryPropertyFlagBits = 0x1 << 0
const VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT VkMemoryPropertyFlagBits = 0x1 << 1
const VK_MEMORY_PROPERTY_HOST_COHERENT_BIT VkMemoryPropertyFlagBits = 0x1 << 2
const VK_MEMORY_PROPERTY_HOST_CACHED_BIT VkMemoryPropertyFlagBits = 0x1 << 3
const VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT VkMemoryPropertyFlagBits = 0x1 << 4
type VkMemoryHeapFlagBits = int64
const VK_MEMORY_HEAP_DEVICE_LOCAL_BIT VkMemoryHeapFlagBits = 0x1 << 0
type VkAccessFlagBits = int64
const VK_ACCESS_INDIRECT_COMMAND_READ_BIT VkAccessFlagBits = 0x1 << 0
const VK_ACCESS_INDEX_READ_BIT VkAccessFlagBits = 0x1 << 1
const VK_ACCESS_VERTEX_ATTRIBUTE_READ_BIT VkAccessFlagBits = 0x1 << 2
const VK_ACCESS_UNIFORM_READ_BIT VkAccessFlagBits = 0x1 << 3
const VK_ACCESS_INPUT_ATTACHMENT_READ_BIT VkAccessFlagBits = 0x1 << 4
const VK_ACCESS_SHADER_READ_BIT VkAccessFlagBits = 0x1 << 5
const VK_ACCESS_SHADER_WRITE_BIT VkAccessFlagBits = 0x1 << 6
const VK_ACCESS_COLOR_ATTACHMENT_READ_BIT VkAccessFlagBits = 0x1 << 7
const VK_ACCESS_COLOR_ATTACHMENT_WRITE_BIT VkAccessFlagBits = 0x1 << 8
const VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT VkAccessFlagBits = 0x1 << 9
const VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT VkAccessFlagBits = 0x1 << 10
const VK_ACCESS_TRANSFER_READ_BIT VkAccessFlagBits = 0x1 << 11
const VK_ACCESS_TRANSFER_WRITE_BIT VkAccessFlagBits = 0x1 << 12
const VK_ACCESS_HOST_READ_BIT VkAccessFlagBits = 0x1 << 13
const VK_ACCESS_HOST_WRITE_BIT VkAccessFlagBits = 0x1 << 14
const VK_ACCESS_MEMORY_READ_BIT VkAccessFlagBits = 0x1 << 15
const VK_ACCESS_MEMORY_WRITE_BIT VkAccessFlagBits = 0x1 << 16
type VkBufferUsageFlagBits = int64
const VK_BUFFER_USAGE_TRANSFER_SRC_BIT VkBufferUsageFlagBits = 0x1 << 0
const VK_BUFFER_USAGE_TRANSFER_DST_BIT VkBufferUsageFlagBits = 0x1 << 1
const VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT VkBufferUsageFlagBits = 0x1 << 2
const VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT VkBufferUsageFlagBits = 0x1 << 3
const VK_BUFFER_USAGE_UNIFORM_BUFFER_BIT VkBufferUsageFlagBits = 0x1 << 4
const VK_BUFFER_USAGE_STORAGE_BUFFER_BIT VkBufferUsageFlagBits = 0x1 << 5
const VK_BUFFER_USAGE_INDEX_BUFFER_BIT VkBufferUsageFlagBits = 0x1 << 6
const VK_BUFFER_USAGE_VERTEX_BUFFER_BIT VkBufferUsageFlagBits = 0x1 << 7
const VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT VkBufferUsageFlagBits = 0x1 << 8
type VkBufferCreateFlagBits = int64
const VK_BUFFER_CREATE_SPARSE_BINDING_BIT VkBufferCreateFlagBits = 0x1 << 0
const VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT VkBufferCreateFlagBits = 0x1 << 1
const VK_BUFFER_CREATE_SPARSE_ALIASED_BIT VkBufferCreateFlagBits = 0x1 << 2
type VkShaderStageFlagBits = int64
const VK_SHADER_STAGE_VERTEX_BIT VkShaderStageFlagBits = 0x1 << 0
const VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT VkShaderStageFlagBits = 0x1 << 1
const VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT VkShaderStageFlagBits = 0x1 << 2
const VK_SHADER_STAGE_GEOMETRY_BIT VkShaderStageFlagBits = 0x1 << 3
const VK_SHADER_STAGE_FRAGMENT_BIT VkShaderStageFlagBits = 0x1 << 4
const VK_SHADER_STAGE_COMPUTE_BIT VkShaderStageFlagBits = 0x1 << 5
const VK_SHADER_STAGE_ALL_GRAPHICS VkShaderStageFlagBits = 0x0000001F
const VK_SHADER_STAGE_ALL VkShaderStageFlagBits = 0x7FFFFFFF
type VkImageUsageFlagBits = int64
const VK_IMAGE_USAGE_TRANSFER_SRC_BIT VkImageUsageFlagBits = 0x1 << 0
const VK_IMAGE_USAGE_TRANSFER_DST_BIT VkImageUsageFlagBits = 0x1 << 1
const VK_IMAGE_USAGE_SAMPLED_BIT VkImageUsageFlagBits = 0x1 << 2
const VK_IMAGE_USAGE_STORAGE_BIT VkImageUsageFlagBits = 0x1 << 3
const VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT VkImageUsageFlagBits = 0x1 << 4
const VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT VkImageUsageFlagBits = 0x1 << 5
const VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT VkImageUsageFlagBits = 0x1 << 6
const VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT VkImageUsageFlagBits = 0x1 << 7
type VkImageCreateFlagBits = int64
const VK_IMAGE_CREATE_SPARSE_BINDING_BIT VkImageCreateFlagBits = 0x1 << 0
const VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT VkImageCreateFlagBits = 0x1 << 1
const VK_IMAGE_CREATE_SPARSE_ALIASED_BIT VkImageCreateFlagBits = 0x1 << 2
const VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT VkImageCreateFlagBits = 0x1 << 3
const VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT VkImageCreateFlagBits = 0x1 << 4
type VkImageViewCreateFlagBits = int64
type VkSamplerCreateFlagBits = int64
type VkPipelineCreateFlagBits = int64
const VK_PIPELINE_CREATE_DISABLE_OPTIMIZATION_BIT VkPipelineCreateFlagBits = 0x1 << 0
const VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT VkPipelineCreateFlagBits = 0x1 << 1
const VK_PIPELINE_CREATE_DERIVATIVE_BIT VkPipelineCreateFlagBits = 0x1 << 2
type VkPipelineShaderStageCreateFlagBits = int64
type VkColorComponentFlagBits = int64
const VK_COLOR_COMPONENT_R_BIT VkColorComponentFlagBits = 0x1 << 0
const VK_COLOR_COMPONENT_G_BIT VkColorComponentFlagBits = 0x1 << 1
const VK_COLOR_COMPONENT_B_BIT VkColorComponentFlagBits = 0x1 << 2
const VK_COLOR_COMPONENT_A_BIT VkColorComponentFlagBits = 0x1 << 3
type VkFenceCreateFlagBits = int64
const VK_FENCE_CREATE_SIGNALED_BIT VkFenceCreateFlagBits = 0x1 << 0
type VkSemaphoreCreateFlagBits = int64
type VkFormatFeatureFlagBits = int64
const VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT VkFormatFeatureFlagBits = 0x1 << 0
const VK_FORMAT_FEATURE_STORAGE_IMAGE_BIT VkFormatFeatureFlagBits = 0x1 << 1
const VK_FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT VkFormatFeatureFlagBits = 0x1 << 2
const VK_FORMAT_FEATURE_UNIFORM_TEXEL_BUFFER_BIT VkFormatFeatureFlagBits = 0x1 << 3
const VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT VkFormatFeatureFlagBits = 0x1 << 4
const VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT VkFormatFeatureFlagBits = 0x1 << 5
const VK_FORMAT_FEATURE_VERTEX_BUFFER_BIT VkFormatFeatureFlagBits = 0x1 << 6
const VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT VkFormatFeatureFlagBits = 0x1 << 7
const VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT VkFormatFeatureFlagBits = 0x1 << 8
const VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT VkFormatFeatureFlagBits = 0x1 << 9
const VK_FORMAT_FEATURE_BLIT_SRC_BIT VkFormatFeatureFlagBits = 0x1 << 10
const VK_FORMAT_FEATURE_BLIT_DST_BIT VkFormatFeatureFlagBits = 0x1 << 11
const VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT VkFormatFeatureFlagBits = 0x1 << 12
type VkQueryControlFlagBits = int64
const VK_QUERY_CONTROL_PRECISE_BIT VkQueryControlFlagBits = 0x1 << 0
type VkQueryResultFlagBits = int64
const VK_QUERY_RESULT_64_BIT VkQueryResultFlagBits = 0x1 << 0
const VK_QUERY_RESULT_WAIT_BIT VkQueryResultFlagBits = 0x1 << 1
const VK_QUERY_RESULT_WITH_AVAILABILITY_BIT VkQueryResultFlagBits = 0x1 << 2
const VK_QUERY_RESULT_PARTIAL_BIT VkQueryResultFlagBits = 0x1 << 3
type VkCommandBufferUsageFlagBits = int64
const VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT VkCommandBufferUsageFlagBits = 0x1 << 0
const VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT VkCommandBufferUsageFlagBits = 0x1 << 1
const VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT VkCommandBufferUsageFlagBits = 0x1 << 2
type VkQueryPipelineStatisticFlagBits = int64
const VK_QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_VERTICES_BIT VkQueryPipelineStatisticFlagBits = 0x1 << 0
const VK_QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_PRIMITIVES_BIT VkQueryPipelineStatisticFlagBits = 0x1 << 1
const VK_QUERY_PIPELINE_STATISTIC_VERTEX_SHADER_INVOCATIONS_BIT VkQueryPipelineStatisticFlagBits = 0x1 << 2
const VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_INVOCATIONS_BIT VkQueryPipelineStatisticFlagBits = 0x1 << 3
const VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_PRIMITIVES_BIT VkQueryPipelineStatisticFlagBits = 0x1 << 4
const VK_QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT VkQueryPipelineStatisticFlagBits = 0x1 << 5
const VK_QUERY_PIPELINE_STATISTIC_CLIPPING_PRIMITIVES_BIT VkQueryPipelineStatisticFlagBits = 0x1 << 6
const VK_QUERY_PIPELINE_STATISTIC_FRAGMENT_SHADER_INVOCATIONS_BIT VkQueryPipelineStatisticFlagBits = 0x1 << 7
const VK_QUERY_PIPELINE_STATISTIC_TESSELLATION_CONTROL_SHADER_PATCHES_BIT VkQueryPipelineStatisticFlagBits = 0x1 << 8
const VK_QUERY_PIPELINE_STATISTIC_TESSELLATION_EVALUATION_SHADER_INVOCATIONS_BIT VkQueryPipelineStatisticFlagBits = 0x1 << 9
const VK_QUERY_PIPELINE_STATISTIC_COMPUTE_SHADER_INVOCATIONS_BIT VkQueryPipelineStatisticFlagBits = 0x1 << 10
type VkImageAspectFlagBits = int64
const VK_IMAGE_ASPECT_COLOR_BIT VkImageAspectFlagBits = 0x1 << 0
const VK_IMAGE_ASPECT_DEPTH_BIT VkImageAspectFlagBits = 0x1 << 1
const VK_IMAGE_ASPECT_STENCIL_BIT VkImageAspectFlagBits = 0x1 << 2
const VK_IMAGE_ASPECT_METADATA_BIT VkImageAspectFlagBits = 0x1 << 3
type VkSparseImageFormatFlagBits = int64
const VK_SPARSE_IMAGE_FORMAT_SINGLE_MIPTAIL_BIT VkSparseImageFormatFlagBits = 0x1 << 0
const VK_SPARSE_IMAGE_FORMAT_ALIGNED_MIP_SIZE_BIT VkSparseImageFormatFlagBits = 0x1 << 1
const VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT VkSparseImageFormatFlagBits = 0x1 << 2
type VkSparseMemoryBindFlagBits = int64
const VK_SPARSE_MEMORY_BIND_METADATA_BIT VkSparseMemoryBindFlagBits = 0x1 << 0
type VkPipelineStageFlagBits = int64
const VK_PIPELINE_STAGE_TOP_OF_PIPE_BIT VkPipelineStageFlagBits = 0x1 << 0
const VK_PIPELINE_STAGE_DRAW_INDIRECT_BIT VkPipelineStageFlagBits = 0x1 << 1
const VK_PIPELINE_STAGE_VERTEX_INPUT_BIT VkPipelineStageFlagBits = 0x1 << 2
const VK_PIPELINE_STAGE_VERTEX_SHADER_BIT VkPipelineStageFlagBits = 0x1 << 3
const VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT VkPipelineStageFlagBits = 0x1 << 4
const VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT VkPipelineStageFlagBits = 0x1 << 5
const VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT VkPipelineStageFlagBits = 0x1 << 6
const VK_PIPELINE_STAGE_FRAGMENT_SHADER_BIT VkPipelineStageFlagBits = 0x1 << 7
const VK_PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT VkPipelineStageFlagBits = 0x1 << 8
const VK_PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT VkPipelineStageFlagBits = 0x1 << 9
const VK_PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT VkPipelineStageFlagBits = 0x1 << 10
const VK_PIPELINE_STAGE_COMPUTE_SHADER_BIT VkPipelineStageFlagBits = 0x1 << 11
const VK_PIPELINE_STAGE_TRANSFER_BIT VkPipelineStageFlagBits = 0x1 << 12
const VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT VkPipelineStageFlagBits = 0x1 << 13
const VK_PIPELINE_STAGE_HOST_BIT VkPipelineStageFlagBits = 0x1 << 14
const VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT VkPipelineStageFlagBits = 0x1 << 15
const VK_PIPELINE_STAGE_ALL_COMMANDS_BIT VkPipelineStageFlagBits = 0x1 << 16
type VkCommandPoolCreateFlagBits = int64
const VK_COMMAND_POOL_CREATE_TRANSIENT_BIT VkCommandPoolCreateFlagBits = 0x1 << 0
const VK_COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT VkCommandPoolCreateFlagBits = 0x1 << 1
type VkCommandPoolResetFlagBits = int64
const VK_COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT VkCommandPoolResetFlagBits = 0x1 << 0
type VkCommandBufferResetFlagBits = int64
const VK_COMMAND_BUFFER_RESET_RELEASE_RESOURCES_BIT VkCommandBufferResetFlagBits = 0x1 << 0
type VkSampleCountFlagBits = int64
const VK_SAMPLE_COUNT_1_BIT VkSampleCountFlagBits = 0x1 << 0
const VK_SAMPLE_COUNT_2_BIT VkSampleCountFlagBits = 0x1 << 1
const VK_SAMPLE_COUNT_4_BIT VkSampleCountFlagBits = 0x1 << 2
const VK_SAMPLE_COUNT_8_BIT VkSampleCountFlagBits = 0x1 << 3
const VK_SAMPLE_COUNT_16_BIT VkSampleCountFlagBits = 0x1 << 4
const VK_SAMPLE_COUNT_32_BIT VkSampleCountFlagBits = 0x1 << 5
const VK_SAMPLE_COUNT_64_BIT VkSampleCountFlagBits = 0x1 << 6
type VkAttachmentDescriptionFlagBits = int64
const VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT VkAttachmentDescriptionFlagBits = 0x1 << 0
type VkStencilFaceFlagBits = int64
const VK_STENCIL_FACE_FRONT_BIT VkStencilFaceFlagBits = 0x1 << 0
const VK_STENCIL_FACE_BACK_BIT VkStencilFaceFlagBits = 0x1 << 1
const VK_STENCIL_FACE_FRONT_AND_BACK VkStencilFaceFlagBits = 0x00000003
const VK_STENCIL_FRONT_AND_BACK VkStencilFaceFlagBits = VK_STENCIL_FACE_FRONT_AND_BACK
type VkDescriptorPoolCreateFlagBits = int64
const VK_DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT VkDescriptorPoolCreateFlagBits = 0x1 << 0
type VkDependencyFlagBits = int64
const VK_DEPENDENCY_BY_REGION_BIT VkDependencyFlagBits = 0x1 << 0
type VkSemaphoreType = int32
const VK_SEMAPHORE_TYPE_BINARY VkSemaphoreType = 0
const VK_SEMAPHORE_TYPE_TIMELINE VkSemaphoreType = 1
type VkSemaphoreWaitFlagBits = int64
const VK_SEMAPHORE_WAIT_ANY_BIT VkSemaphoreWaitFlagBits = 0x1 << 0
type VkPresentModeKHR = int32
const VK_PRESENT_MODE_IMMEDIATE_KHR VkPresentModeKHR = 0
const VK_PRESENT_MODE_MAILBOX_KHR VkPresentModeKHR = 1
const VK_PRESENT_MODE_FIFO_KHR VkPresentModeKHR = 2
const VK_PRESENT_MODE_FIFO_RELAXED_KHR VkPresentModeKHR = 3
type VkColorSpaceKHR = int32
const VK_COLOR_SPACE_SRGB_NONLINEAR_KHR VkColorSpaceKHR = 0
const VK_COLORSPACE_SRGB_NONLINEAR_KHR VkColorSpaceKHR = VK_COLOR_SPACE_SRGB_NONLINEAR_KHR
type VkDisplayPlaneAlphaFlagBitsKHR = int64
const VK_DISPLAY_PLANE_ALPHA_OPAQUE_BIT_KHR VkDisplayPlaneAlphaFlagBitsKHR = 0x1 << 0
const VK_DISPLAY_PLANE_ALPHA_GLOBAL_BIT_KHR VkDisplayPlaneAlphaFlagBitsKHR = 0x1 << 1
const VK_DISPLAY_PLANE_ALPHA_PER_PIXEL_BIT_KHR VkDisplayPlaneAlphaFlagBitsKHR = 0x1 << 2
const VK_DISPLAY_PLANE_ALPHA_PER_PIXEL_PREMULTIPLIED_BIT_KHR VkDisplayPlaneAlphaFlagBitsKHR = 0x1 << 3
type VkCompositeAlphaFlagBitsKHR = int64
const VK_COMPOSITE_ALPHA_OPAQUE_BIT_KHR VkCompositeAlphaFlagBitsKHR = 0x1 << 0
const VK_COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR VkCompositeAlphaFlagBitsKHR = 0x1 << 1
const VK_COMPOSITE_ALPHA_POST_MULTIPLIED_BIT_KHR VkCompositeAlphaFlagBitsKHR = 0x1 << 2
const VK_COMPOSITE_ALPHA_INHERIT_BIT_KHR VkCompositeAlphaFlagBitsKHR = 0x1 << 3
type VkSurfaceTransformFlagBitsKHR = int64
const VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR VkSurfaceTransformFlagBitsKHR = 0x1 << 0
const VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR VkSurfaceTransformFlagBitsKHR = 0x1 << 1
const VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR VkSurfaceTransformFlagBitsKHR = 0x1 << 2
const VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR VkSurfaceTransformFlagBitsKHR = 0x1 << 3
const VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_BIT_KHR VkSurfaceTransformFlagBitsKHR = 0x1 << 4
const VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_90_BIT_KHR VkSurfaceTransformFlagBitsKHR = 0x1 << 5
const VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_180_BIT_KHR VkSurfaceTransformFlagBitsKHR = 0x1 << 6
const VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_270_BIT_KHR VkSurfaceTransformFlagBitsKHR = 0x1 << 7
const VK_SURFACE_TRANSFORM_INHERIT_BIT_KHR VkSurfaceTransformFlagBitsKHR = 0x1 << 8
type VkSwapchainImageUsageFlagBitsANDROID = int64
const VK_SWAPCHAIN_IMAGE_USAGE_SHARED_BIT_ANDROID VkSwapchainImageUsageFlagBitsANDROID = 0x1 << 0
type VkTimeDomainEXT = int32
const VK_TIME_DOMAIN_DEVICE_EXT VkTimeDomainEXT = 0
const VK_TIME_DOMAIN_CLOCK_MONOTONIC_EXT VkTimeDomainEXT = 1
const VK_TIME_DOMAIN_CLOCK_MONOTONIC_RAW_EXT VkTimeDomainEXT = 2
const VK_TIME_DOMAIN_QUERY_PERFORMANCE_COUNTER_EXT VkTimeDomainEXT = 3
type VkDebugReportFlagBitsEXT = int64
const VK_DEBUG_REPORT_INFORMATION_BIT_EXT VkDebugReportFlagBitsEXT = 0x1 << 0
const VK_DEBUG_REPORT_WARNING_BIT_EXT VkDebugReportFlagBitsEXT = 0x1 << 1
const VK_DEBUG_REPORT_PERFORMANCE_WARNING_BIT_EXT VkDebugReportFlagBitsEXT = 0x1 << 2
const VK_DEBUG_REPORT_ERROR_BIT_EXT VkDebugReportFlagBitsEXT = 0x1 << 3
const VK_DEBUG_REPORT_DEBUG_BIT_EXT VkDebugReportFlagBitsEXT = 0x1 << 4
type VkDebugReportObjectTypeEXT = int32
const VK_DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT VkDebugReportObjectTypeEXT = 0
const VK_DEBUG_REPORT_OBJECT_TYPE_INSTANCE_EXT VkDebugReportObjectTypeEXT = 1
const VK_DEBUG_REPORT_OBJECT_TYPE_PHYSICAL_DEVICE_EXT VkDebugReportObjectTypeEXT = 2
const VK_DEBUG_REPORT_OBJECT_TYPE_DEVICE_EXT VkDebugReportObjectTypeEXT = 3
const VK_DEBUG_REPORT_OBJECT_TYPE_QUEUE_EXT VkDebugReportObjectTypeEXT = 4
const VK_DEBUG_REPORT_OBJECT_TYPE_SEMAPHORE_EXT VkDebugReportObjectTypeEXT = 5
const VK_DEBUG_REPORT_OBJECT_TYPE_COMMAND_BUFFER_EXT VkDebugReportObjectTypeEXT = 6
const VK_DEBUG_REPORT_OBJECT_TYPE_FENCE_EXT VkDebugReportObjectTypeEXT = 7
const VK_DEBUG_REPORT_OBJECT_TYPE_DEVICE_MEMORY_EXT VkDebugReportObjectTypeEXT = 8
const VK_DEBUG_REPORT_OBJECT_TYPE_BUFFER_EXT VkDebugReportObjectTypeEXT = 9
const VK_DEBUG_REPORT_OBJECT_TYPE_IMAGE_EXT VkDebugReportObjectTypeEXT = 10
const VK_DEBUG_REPORT_OBJECT_TYPE_EVENT_EXT VkDebugReportObjectTypeEXT = 11
const VK_DEBUG_REPORT_OBJECT_TYPE_QUERY_POOL_EXT VkDebugReportObjectTypeEXT = 12
const VK_DEBUG_REPORT_OBJECT_TYPE_BUFFER_VIEW_EXT VkDebugReportObjectTypeEXT = 13
const VK_DEBUG_REPORT_OBJECT_TYPE_IMAGE_VIEW_EXT VkDebugReportObjectTypeEXT = 14
const VK_DEBUG_REPORT_OBJECT_TYPE_SHADER_MODULE_EXT VkDebugReportObjectTypeEXT = 15
const VK_DEBUG_REPORT_OBJECT_TYPE_PIPELINE_CACHE_EXT VkDebugReportObjectTypeEXT = 16
const VK_DEBUG_REPORT_OBJECT_TYPE_PIPELINE_LAYOUT_EXT VkDebugReportObjectTypeEXT = 17
const VK_DEBUG_REPORT_OBJECT_TYPE_RENDER_PASS_EXT VkDebugReportObjectTypeEXT = 18
const VK_DEBUG_REPORT_OBJECT_TYPE_PIPELINE_EXT VkDebugReportObjectTypeEXT = 19
const VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT_EXT VkDebugReportObjectTypeEXT = 20
const VK_DEBUG_REPORT_OBJECT_TYPE_SAMPLER_EXT VkDebugReportObjectTypeEXT = 21
const VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_POOL_EXT VkDebugReportObjectTypeEXT = 22
const VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_EXT VkDebugReportObjectTypeEXT = 23
const VK_DEBUG_REPORT_OBJECT_TYPE_FRAMEBUFFER_EXT VkDebugReportObjectTypeEXT = 24
const VK_DEBUG_REPORT_OBJECT_TYPE_COMMAND_POOL_EXT VkDebugReportObjectTypeEXT = 25
const VK_DEBUG_REPORT_OBJECT_TYPE_SURFACE_KHR_EXT VkDebugReportObjectTypeEXT = 26
const VK_DEBUG_REPORT_OBJECT_TYPE_SWAPCHAIN_KHR_EXT VkDebugReportObjectTypeEXT = 27
const VK_DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT_EXT VkDebugReportObjectTypeEXT = 28
const VK_DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_EXT VkDebugReportObjectTypeEXT = VK_DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT_EXT
const VK_DEBUG_REPORT_OBJECT_TYPE_DISPLAY_KHR_EXT VkDebugReportObjectTypeEXT = 29
const VK_DEBUG_REPORT_OBJECT_TYPE_DISPLAY_MODE_KHR_EXT VkDebugReportObjectTypeEXT = 30
const VK_DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT_EXT VkDebugReportObjectTypeEXT = 33
const VK_DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT VkDebugReportObjectTypeEXT = VK_DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT_EXT
type VkDeviceMemoryReportEventTypeEXT = int32
const VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT VkDeviceMemoryReportEventTypeEXT = 0
const VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_FREE_EXT VkDeviceMemoryReportEventTypeEXT = 1
const VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_IMPORT_EXT VkDeviceMemoryReportEventTypeEXT = 2
const VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_UNIMPORT_EXT VkDeviceMemoryReportEventTypeEXT = 3
const VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATION_FAILED_EXT VkDeviceMemoryReportEventTypeEXT = 4
type VkRasterizationOrderAMD = int32
const VK_RASTERIZATION_ORDER_STRICT_AMD VkRasterizationOrderAMD = 0
const VK_RASTERIZATION_ORDER_RELAXED_AMD VkRasterizationOrderAMD = 1
type VkExternalMemoryHandleTypeFlagBitsNV = int64
const VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_NV VkExternalMemoryHandleTypeFlagBitsNV = 0x1 << 0
const VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_NV VkExternalMemoryHandleTypeFlagBitsNV = 0x1 << 1
const VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_IMAGE_BIT_NV VkExternalMemoryHandleTypeFlagBitsNV = 0x1 << 2
const VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_IMAGE_KMT_BIT_NV VkExternalMemoryHandleTypeFlagBitsNV = 0x1 << 3
type VkExternalMemoryFeatureFlagBitsNV = int64
const VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT_NV VkExternalMemoryFeatureFlagBitsNV = 0x1 << 0
const VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT_NV VkExternalMemoryFeatureFlagBitsNV = 0x1 << 1
const VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_NV VkExternalMemoryFeatureFlagBitsNV = 0x1 << 2
type VkValidationCheckEXT = int32
const VK_VALIDATION_CHECK_ALL_EXT VkValidationCheckEXT = 0
const VK_VALIDATION_CHECK_SHADERS_EXT VkValidationCheckEXT = 1
type VkValidationFeatureEnableEXT = int32
const VK_VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_EXT VkValidationFeatureEnableEXT = 0
const VK_VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_RESERVE_BINDING_SLOT_EXT VkValidationFeatureEnableEXT = 1
const VK_VALIDATION_FEATURE_ENABLE_BEST_PRACTICES_EXT VkValidationFeatureEnableEXT = 2
const VK_VALIDATION_FEATURE_ENABLE_DEBUG_PRINTF_EXT VkValidationFeatureEnableEXT = 3
const VK_VALIDATION_FEATURE_ENABLE_SYNCHRONIZATION_VALIDATION_EXT VkValidationFeatureEnableEXT = 4
type VkValidationFeatureDisableEXT = int32
const VK_VALIDATION_FEATURE_DISABLE_ALL_EXT VkValidationFeatureDisableEXT = 0
const VK_VALIDATION_FEATURE_DISABLE_SHADERS_EXT VkValidationFeatureDisableEXT = 1
const VK_VALIDATION_FEATURE_DISABLE_THREAD_SAFETY_EXT VkValidationFeatureDisableEXT = 2
const VK_VALIDATION_FEATURE_DISABLE_API_PARAMETERS_EXT VkValidationFeatureDisableEXT = 3
const VK_VALIDATION_FEATURE_DISABLE_OBJECT_LIFETIMES_EXT VkValidationFeatureDisableEXT = 4
const VK_VALIDATION_FEATURE_DISABLE_CORE_CHECKS_EXT VkValidationFeatureDisableEXT = 5
const VK_VALIDATION_FEATURE_DISABLE_UNIQUE_HANDLES_EXT VkValidationFeatureDisableEXT = 6
type VkSubgroupFeatureFlagBits = int64
const VK_SUBGROUP_FEATURE_BASIC_BIT VkSubgroupFeatureFlagBits = 0x1 << 0
const VK_SUBGROUP_FEATURE_VOTE_BIT VkSubgroupFeatureFlagBits = 0x1 << 1
const VK_SUBGROUP_FEATURE_ARITHMETIC_BIT VkSubgroupFeatureFlagBits = 0x1 << 2
const VK_SUBGROUP_FEATURE_BALLOT_BIT VkSubgroupFeatureFlagBits = 0x1 << 3
const VK_SUBGROUP_FEATURE_SHUFFLE_BIT VkSubgroupFeatureFlagBits = 0x1 << 4
const VK_SUBGROUP_FEATURE_SHUFFLE_RELATIVE_BIT VkSubgroupFeatureFlagBits = 0x1 << 5
const VK_SUBGROUP_FEATURE_CLUSTERED_BIT VkSubgroupFeatureFlagBits = 0x1 << 6
const VK_SUBGROUP_FEATURE_QUAD_BIT VkSubgroupFeatureFlagBits = 0x1 << 7
type VkIndirectCommandsLayoutUsageFlagBitsNV = int64
const VK_INDIRECT_COMMANDS_LAYOUT_USAGE_EXPLICIT_PREPROCESS_BIT_NV VkIndirectCommandsLayoutUsageFlagBitsNV = 0x1 << 0
const VK_INDIRECT_COMMANDS_LAYOUT_USAGE_INDEXED_SEQUENCES_BIT_NV VkIndirectCommandsLayoutUsageFlagBitsNV = 0x1 << 1
const VK_INDIRECT_COMMANDS_LAYOUT_USAGE_UNORDERED_SEQUENCES_BIT_NV VkIndirectCommandsLayoutUsageFlagBitsNV = 0x1 << 2
type VkIndirectStateFlagBitsNV = int64
const VK_INDIRECT_STATE_FLAG_FRONTFACE_BIT_NV VkIndirectStateFlagBitsNV = 0x1 << 0
type VkIndirectCommandsTokenTypeNV = int32
const VK_INDIRECT_COMMANDS_TOKEN_TYPE_SHADER_GROUP_NV VkIndirectCommandsTokenTypeNV = 0
const VK_INDIRECT_COMMANDS_TOKEN_TYPE_STATE_FLAGS_NV VkIndirectCommandsTokenTypeNV = 1
const VK_INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_NV VkIndirectCommandsTokenTypeNV = 2
const VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_NV VkIndirectCommandsTokenTypeNV = 3
const VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NV VkIndirectCommandsTokenTypeNV = 4
const VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_NV VkIndirectCommandsTokenTypeNV = 5
const VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_NV VkIndirectCommandsTokenTypeNV = 6
const VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_TASKS_NV VkIndirectCommandsTokenTypeNV = 7
type VkPrivateDataSlotCreateFlagBitsEXT = int64
type VkDescriptorSetLayoutCreateFlagBits = int64
type VkExternalMemoryHandleTypeFlagBits = int64
const VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT VkExternalMemoryHandleTypeFlagBits = 0x1 << 0
const VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT VkExternalMemoryHandleTypeFlagBits = 0x1 << 1
const VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT VkExternalMemoryHandleTypeFlagBits = 0x1 << 2
const VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT VkExternalMemoryHandleTypeFlagBits = 0x1 << 3
const VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT VkExternalMemoryHandleTypeFlagBits = 0x1 << 4
const VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT VkExternalMemoryHandleTypeFlagBits = 0x1 << 5
const VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT VkExternalMemoryHandleTypeFlagBits = 0x1 << 6
type VkExternalMemoryFeatureFlagBits = int64
const VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT VkExternalMemoryFeatureFlagBits = 0x1 << 0
const VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT VkExternalMemoryFeatureFlagBits = 0x1 << 1
const VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT VkExternalMemoryFeatureFlagBits = 0x1 << 2
type VkExternalSemaphoreHandleTypeFlagBits = int64
const VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT VkExternalSemaphoreHandleTypeFlagBits = 0x1 << 0
const VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT VkExternalSemaphoreHandleTypeFlagBits = 0x1 << 1
const VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT VkExternalSemaphoreHandleTypeFlagBits = 0x1 << 2
const VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT VkExternalSemaphoreHandleTypeFlagBits = 0x1 << 3
const VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D11_FENCE_BIT VkExternalSemaphoreHandleTypeFlagBits = VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT
const VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT VkExternalSemaphoreHandleTypeFlagBits = 0x1 << 4
type VkExternalSemaphoreFeatureFlagBits = int64
const VK_EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT VkExternalSemaphoreFeatureFlagBits = 0x1 << 0
const VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT VkExternalSemaphoreFeatureFlagBits = 0x1 << 1
type VkSemaphoreImportFlagBits = int64
const VK_SEMAPHORE_IMPORT_TEMPORARY_BIT VkSemaphoreImportFlagBits = 0x1 << 0
type VkExternalFenceHandleTypeFlagBits = int64
const VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT VkExternalFenceHandleTypeFlagBits = 0x1 << 0
const VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT VkExternalFenceHandleTypeFlagBits = 0x1 << 1
const VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT VkExternalFenceHandleTypeFlagBits = 0x1 << 2
const VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT VkExternalFenceHandleTypeFlagBits = 0x1 << 3
type VkExternalFenceFeatureFlagBits = int64
const VK_EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT VkExternalFenceFeatureFlagBits = 0x1 << 0
const VK_EXTERNAL_FENCE_FEATURE_IMPORTABLE_BIT VkExternalFenceFeatureFlagBits = 0x1 << 1
type VkFenceImportFlagBits = int64
const VK_FENCE_IMPORT_TEMPORARY_BIT VkFenceImportFlagBits = 0x1 << 0
type VkSurfaceCounterFlagBitsEXT = int64
const VK_SURFACE_COUNTER_VBLANK_EXT VkSurfaceCounterFlagBitsEXT = 0x1 << 0
type VkDisplayPowerStateEXT = int32
const VK_DISPLAY_POWER_STATE_OFF_EXT VkDisplayPowerStateEXT = 0
const VK_DISPLAY_POWER_STATE_SUSPEND_EXT VkDisplayPowerStateEXT = 1
const VK_DISPLAY_POWER_STATE_ON_EXT VkDisplayPowerStateEXT = 2
type VkDeviceEventTypeEXT = int32
const VK_DEVICE_EVENT_TYPE_DISPLAY_HOTPLUG_EXT VkDeviceEventTypeEXT = 0
type VkDisplayEventTypeEXT = int32
const VK_DISPLAY_EVENT_TYPE_FIRST_PIXEL_OUT_EXT VkDisplayEventTypeEXT = 0
type VkPeerMemoryFeatureFlagBits = int64
const VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT VkPeerMemoryFeatureFlagBits = 0x1 << 0
const VK_PEER_MEMORY_FEATURE_COPY_DST_BIT VkPeerMemoryFeatureFlagBits = 0x1 << 1
const VK_PEER_MEMORY_FEATURE_GENERIC_SRC_BIT VkPeerMemoryFeatureFlagBits = 0x1 << 2
const VK_PEER_MEMORY_FEATURE_GENERIC_DST_BIT VkPeerMemoryFeatureFlagBits = 0x1 << 3
type VkMemoryAllocateFlagBits = int64
const VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT VkMemoryAllocateFlagBits = 0x1 << 0
type VkDeviceGroupPresentModeFlagBitsKHR = int64
const VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR VkDeviceGroupPresentModeFlagBitsKHR = 0x1 << 0
const VK_DEVICE_GROUP_PRESENT_MODE_REMOTE_BIT_KHR VkDeviceGroupPresentModeFlagBitsKHR = 0x1 << 1
const VK_DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR VkDeviceGroupPresentModeFlagBitsKHR = 0x1 << 2
const VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_MULTI_DEVICE_BIT_KHR VkDeviceGroupPresentModeFlagBitsKHR = 0x1 << 3
type VkSwapchainCreateFlagBitsKHR = int64
type VkViewportCoordinateSwizzleNV = int32
const VK_VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_X_NV VkViewportCoordinateSwizzleNV = 0
const VK_VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_X_NV VkViewportCoordinateSwizzleNV = 1
const VK_VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_Y_NV VkViewportCoordinateSwizzleNV = 2
const VK_VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_Y_NV VkViewportCoordinateSwizzleNV = 3
const VK_VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_Z_NV VkViewportCoordinateSwizzleNV = 4
const VK_VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_Z_NV VkViewportCoordinateSwizzleNV = 5
const VK_VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_W_NV VkViewportCoordinateSwizzleNV = 6
const VK_VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_W_NV VkViewportCoordinateSwizzleNV = 7
type VkDiscardRectangleModeEXT = int32
const VK_DISCARD_RECTANGLE_MODE_INCLUSIVE_EXT VkDiscardRectangleModeEXT = 0
const VK_DISCARD_RECTANGLE_MODE_EXCLUSIVE_EXT VkDiscardRectangleModeEXT = 1
type VkSubpassDescriptionFlagBits = int64
type VkPointClippingBehavior = int32
const VK_POINT_CLIPPING_BEHAVIOR_ALL_CLIP_PLANES VkPointClippingBehavior = 0
const VK_POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY VkPointClippingBehavior = 1
type VkSamplerReductionMode = int32
const VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE VkSamplerReductionMode = 0
const VK_SAMPLER_REDUCTION_MODE_MIN VkSamplerReductionMode = 1
const VK_SAMPLER_REDUCTION_MODE_MAX VkSamplerReductionMode = 2
type VkTessellationDomainOrigin = int32
const VK_TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT VkTessellationDomainOrigin = 0
const VK_TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT VkTessellationDomainOrigin = 1
type VkSamplerYcbcrModelConversion = int32
const VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY VkSamplerYcbcrModelConversion = 0
const VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY VkSamplerYcbcrModelConversion = 1
const VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709 VkSamplerYcbcrModelConversion = 2
const VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601 VkSamplerYcbcrModelConversion = 3
const VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020 VkSamplerYcbcrModelConversion = 4
type VkSamplerYcbcrRange = int32
const VK_SAMPLER_YCBCR_RANGE_ITU_FULL VkSamplerYcbcrRange = 0
const VK_SAMPLER_YCBCR_RANGE_ITU_NARROW VkSamplerYcbcrRange = 1
type VkChromaLocation = int32
const VK_CHROMA_LOCATION_COSITED_EVEN VkChromaLocation = 0
const VK_CHROMA_LOCATION_MIDPOINT VkChromaLocation = 1
type VkBlendOverlapEXT = int32
const VK_BLEND_OVERLAP_UNCORRELATED_EXT VkBlendOverlapEXT = 0
const VK_BLEND_OVERLAP_DISJOINT_EXT VkBlendOverlapEXT = 1
const VK_BLEND_OVERLAP_CONJOINT_EXT VkBlendOverlapEXT = 2
type VkCoverageModulationModeNV = int32
const VK_COVERAGE_MODULATION_MODE_NONE_NV VkCoverageModulationModeNV = 0
const VK_COVERAGE_MODULATION_MODE_RGB_NV VkCoverageModulationModeNV = 1
const VK_COVERAGE_MODULATION_MODE_ALPHA_NV VkCoverageModulationModeNV = 2
const VK_COVERAGE_MODULATION_MODE_RGBA_NV VkCoverageModulationModeNV = 3
type VkCoverageReductionModeNV = int32
const VK_COVERAGE_REDUCTION_MODE_MERGE_NV VkCoverageReductionModeNV = 0
const VK_COVERAGE_REDUCTION_MODE_TRUNCATE_NV VkCoverageReductionModeNV = 1
type VkValidationCacheHeaderVersionEXT = int32
const VK_VALIDATION_CACHE_HEADER_VERSION_ONE_EXT VkValidationCacheHeaderVersionEXT = 1
type VkShaderInfoTypeAMD = int32
const VK_SHADER_INFO_TYPE_STATISTICS_AMD VkShaderInfoTypeAMD = 0
const VK_SHADER_INFO_TYPE_BINARY_AMD VkShaderInfoTypeAMD = 1
const VK_SHADER_INFO_TYPE_DISASSEMBLY_AMD VkShaderInfoTypeAMD = 2
type VkQueueGlobalPriorityEXT = int32
const VK_QUEUE_GLOBAL_PRIORITY_LOW_EXT VkQueueGlobalPriorityEXT = 128
const VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_EXT VkQueueGlobalPriorityEXT = 256
const VK_QUEUE_GLOBAL_PRIORITY_HIGH_EXT VkQueueGlobalPriorityEXT = 512
const VK_QUEUE_GLOBAL_PRIORITY_REALTIME_EXT VkQueueGlobalPriorityEXT = 1024
type VkDebugUtilsMessageSeverityFlagBitsEXT = int64
const VK_DEBUG_UTILS_MESSAGE_SEVERITY_VERBOSE_BIT_EXT VkDebugUtilsMessageSeverityFlagBitsEXT = 0x1 << 0
const VK_DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT VkDebugUtilsMessageSeverityFlagBitsEXT = 0x1 << 4
const VK_DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT VkDebugUtilsMessageSeverityFlagBitsEXT = 0x1 << 8
const VK_DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT VkDebugUtilsMessageSeverityFlagBitsEXT = 0x1 << 12
type VkDebugUtilsMessageTypeFlagBitsEXT = int64
const VK_DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT VkDebugUtilsMessageTypeFlagBitsEXT = 0x1 << 0
const VK_DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT VkDebugUtilsMessageTypeFlagBitsEXT = 0x1 << 1
const VK_DEBUG_UTILS_MESSAGE_TYPE_PERFORMANCE_BIT_EXT VkDebugUtilsMessageTypeFlagBitsEXT = 0x1 << 2
type VkConservativeRasterizationModeEXT = int32
const VK_CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXT VkConservativeRasterizationModeEXT = 0
const VK_CONSERVATIVE_RASTERIZATION_MODE_OVERESTIMATE_EXT VkConservativeRasterizationModeEXT = 1
const VK_CONSERVATIVE_RASTERIZATION_MODE_UNDERESTIMATE_EXT VkConservativeRasterizationModeEXT = 2
type VkDescriptorBindingFlagBits = int64
const VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT VkDescriptorBindingFlagBits = 0x1 << 0
const VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT VkDescriptorBindingFlagBits = 0x1 << 1
const VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT VkDescriptorBindingFlagBits = 0x1 << 2
const VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT VkDescriptorBindingFlagBits = 0x1 << 3
type VkVendorId = int32
const VK_VENDOR_ID_VIV VkVendorId = 0x10001
const VK_VENDOR_ID_VSI VkVendorId = 0x10002
const VK_VENDOR_ID_KAZAN VkVendorId = 0x10003
const VK_VENDOR_ID_CODEPLAY VkVendorId = 0x10004
const VK_VENDOR_ID_MESA VkVendorId = 0x10005
type VkDriverId = int32
const VK_DRIVER_ID_AMD_PROPRIETARY VkDriverId = 1
const VK_DRIVER_ID_AMD_OPEN_SOURCE VkDriverId = 2
const VK_DRIVER_ID_MESA_RADV VkDriverId = 3
const VK_DRIVER_ID_NVIDIA_PROPRIETARY VkDriverId = 4
const VK_DRIVER_ID_INTEL_PROPRIETARY_WINDOWS VkDriverId = 5
const VK_DRIVER_ID_INTEL_OPEN_SOURCE_MESA VkDriverId = 6
const VK_DRIVER_ID_IMAGINATION_PROPRIETARY VkDriverId = 7
const VK_DRIVER_ID_QUALCOMM_PROPRIETARY VkDriverId = 8
const VK_DRIVER_ID_ARM_PROPRIETARY VkDriverId = 9
const VK_DRIVER_ID_GOOGLE_SWIFTSHADER VkDriverId = 10
const VK_DRIVER_ID_GGP_PROPRIETARY VkDriverId = 11
const VK_DRIVER_ID_BROADCOM_PROPRIETARY VkDriverId = 12
const VK_DRIVER_ID_MESA_LLVMPIPE VkDriverId = 13
const VK_DRIVER_ID_MOLTENVK VkDriverId = 14
type VkConditionalRenderingFlagBitsEXT = int64
const VK_CONDITIONAL_RENDERING_INVERTED_BIT_EXT VkConditionalRenderingFlagBitsEXT = 0x1 << 0
type VkResolveModeFlagBits = int64
const VK_RESOLVE_MODE_NONE VkResolveModeFlagBits = 0
const VK_RESOLVE_MODE_SAMPLE_ZERO_BIT VkResolveModeFlagBits = 0x1 << 0
const VK_RESOLVE_MODE_AVERAGE_BIT VkResolveModeFlagBits = 0x1 << 1
const VK_RESOLVE_MODE_MIN_BIT VkResolveModeFlagBits = 0x1 << 2
const VK_RESOLVE_MODE_MAX_BIT VkResolveModeFlagBits = 0x1 << 3
type VkShadingRatePaletteEntryNV = int32
const VK_SHADING_RATE_PALETTE_ENTRY_NO_INVOCATIONS_NV VkShadingRatePaletteEntryNV = 0
const VK_SHADING_RATE_PALETTE_ENTRY_16_INVOCATIONS_PER_PIXEL_NV VkShadingRatePaletteEntryNV = 1
const VK_SHADING_RATE_PALETTE_ENTRY_8_INVOCATIONS_PER_PIXEL_NV VkShadingRatePaletteEntryNV = 2
const VK_SHADING_RATE_PALETTE_ENTRY_4_INVOCATIONS_PER_PIXEL_NV VkShadingRatePaletteEntryNV = 3
const VK_SHADING_RATE_PALETTE_ENTRY_2_INVOCATIONS_PER_PIXEL_NV VkShadingRatePaletteEntryNV = 4
const VK_SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_PIXEL_NV VkShadingRatePaletteEntryNV = 5
const VK_SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_2X1_PIXELS_NV VkShadingRatePaletteEntryNV = 6
const VK_SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_1X2_PIXELS_NV VkShadingRatePaletteEntryNV = 7
const VK_SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_2X2_PIXELS_NV VkShadingRatePaletteEntryNV = 8
const VK_SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_4X2_PIXELS_NV VkShadingRatePaletteEntryNV = 9
const VK_SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_2X4_PIXELS_NV VkShadingRatePaletteEntryNV = 10
const VK_SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_4X4_PIXELS_NV VkShadingRatePaletteEntryNV = 11
type VkCoarseSampleOrderTypeNV = int32
const VK_COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV VkCoarseSampleOrderTypeNV = 0
const VK_COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV VkCoarseSampleOrderTypeNV = 1
const VK_COARSE_SAMPLE_ORDER_TYPE_PIXEL_MAJOR_NV VkCoarseSampleOrderTypeNV = 2
const VK_COARSE_SAMPLE_ORDER_TYPE_SAMPLE_MAJOR_NV VkCoarseSampleOrderTypeNV = 3
type VkGeometryInstanceFlagBitsKHR = int64
const VK_GEOMETRY_INSTANCE_TRIANGLE_FACING_CULL_DISABLE_BIT_KHR VkGeometryInstanceFlagBitsKHR = 0x1 << 0
const VK_GEOMETRY_INSTANCE_TRIANGLE_FRONT_COUNTERCLOCKWISE_BIT_KHR VkGeometryInstanceFlagBitsKHR = 0x1 << 1
const VK_GEOMETRY_INSTANCE_FORCE_OPAQUE_BIT_KHR VkGeometryInstanceFlagBitsKHR = 0x1 << 2
const VK_GEOMETRY_INSTANCE_FORCE_NO_OPAQUE_BIT_KHR VkGeometryInstanceFlagBitsKHR = 0x1 << 3
type VkGeometryFlagBitsKHR = int64
const VK_GEOMETRY_OPAQUE_BIT_KHR VkGeometryFlagBitsKHR = 0x1 << 0
const VK_GEOMETRY_NO_DUPLICATE_ANY_HIT_INVOCATION_BIT_KHR VkGeometryFlagBitsKHR = 0x1 << 1
type VkBuildAccelerationStructureFlagBitsKHR = int64
const VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_KHR VkBuildAccelerationStructureFlagBitsKHR = 0x1 << 0
const VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR VkBuildAccelerationStructureFlagBitsKHR = 0x1 << 1
const VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_TRACE_BIT_KHR VkBuildAccelerationStructureFlagBitsKHR = 0x1 << 2
const VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_BUILD_BIT_KHR VkBuildAccelerationStructureFlagBitsKHR = 0x1 << 3
const VK_BUILD_ACCELERATION_STRUCTURE_LOW_MEMORY_BIT_KHR VkBuildAccelerationStructureFlagBitsKHR = 0x1 << 4
type VkCopyAccelerationStructureModeKHR = int32
const VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_KHR VkCopyAccelerationStructureModeKHR = 0
const VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR VkCopyAccelerationStructureModeKHR = 1
const VK_COPY_ACCELERATION_STRUCTURE_MODE_SERIALIZE_KHR VkCopyAccelerationStructureModeKHR = 2
const VK_COPY_ACCELERATION_STRUCTURE_MODE_DESERIALIZE_KHR VkCopyAccelerationStructureModeKHR = 3
type VkAccelerationStructureTypeKHR = int32
const VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR VkAccelerationStructureTypeKHR = 0
const VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR VkAccelerationStructureTypeKHR = 1
type VkGeometryTypeKHR = int32
const VK_GEOMETRY_TYPE_TRIANGLES_KHR VkGeometryTypeKHR = 0
const VK_GEOMETRY_TYPE_AABBS_KHR VkGeometryTypeKHR = 1
type VkAccelerationStructureMemoryRequirementsTypeKHR = int32
const VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_KHR VkAccelerationStructureMemoryRequirementsTypeKHR = 0
const VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_BUILD_SCRATCH_KHR VkAccelerationStructureMemoryRequirementsTypeKHR = 1
const VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_UPDATE_SCRATCH_KHR VkAccelerationStructureMemoryRequirementsTypeKHR = 2
type VkAccelerationStructureBuildTypeKHR = int32
const VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_KHR VkAccelerationStructureBuildTypeKHR = 0
const VK_ACCELERATION_STRUCTURE_BUILD_TYPE_DEVICE_KHR VkAccelerationStructureBuildTypeKHR = 1
const VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_OR_DEVICE_KHR VkAccelerationStructureBuildTypeKHR = 2
type VkRayTracingShaderGroupTypeKHR = int32
const VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_KHR VkRayTracingShaderGroupTypeKHR = 0
const VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR VkRayTracingShaderGroupTypeKHR = 1
const VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR VkRayTracingShaderGroupTypeKHR = 2
type VkMemoryOverallocationBehaviorAMD = int32
const VK_MEMORY_OVERALLOCATION_BEHAVIOR_DEFAULT_AMD VkMemoryOverallocationBehaviorAMD = 0
const VK_MEMORY_OVERALLOCATION_BEHAVIOR_ALLOWED_AMD VkMemoryOverallocationBehaviorAMD = 1
const VK_MEMORY_OVERALLOCATION_BEHAVIOR_DISALLOWED_AMD VkMemoryOverallocationBehaviorAMD = 2
type VkFramebufferCreateFlagBits = int64
type VkScopeNV = int32
const VK_SCOPE_DEVICE_NV VkScopeNV = 1
const VK_SCOPE_WORKGROUP_NV VkScopeNV = 2
const VK_SCOPE_SUBGROUP_NV VkScopeNV = 3
const VK_SCOPE_QUEUE_FAMILY_NV VkScopeNV = 5
type VkComponentTypeNV = int32
const VK_COMPONENT_TYPE_FLOAT16_NV VkComponentTypeNV = 0
const VK_COMPONENT_TYPE_FLOAT32_NV VkComponentTypeNV = 1
const VK_COMPONENT_TYPE_FLOAT64_NV VkComponentTypeNV = 2
const VK_COMPONENT_TYPE_SINT8_NV VkComponentTypeNV = 3
const VK_COMPONENT_TYPE_SINT16_NV VkComponentTypeNV = 4
const VK_COMPONENT_TYPE_SINT32_NV VkComponentTypeNV = 5
const VK_COMPONENT_TYPE_SINT64_NV VkComponentTypeNV = 6
const VK_COMPONENT_TYPE_UINT8_NV VkComponentTypeNV = 7
const VK_COMPONENT_TYPE_UINT16_NV VkComponentTypeNV = 8
const VK_COMPONENT_TYPE_UINT32_NV VkComponentTypeNV = 9
const VK_COMPONENT_TYPE_UINT64_NV VkComponentTypeNV = 10
type VkDeviceDiagnosticsConfigFlagBitsNV = int64
const VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_DEBUG_INFO_BIT_NV VkDeviceDiagnosticsConfigFlagBitsNV = 0x1 << 0
const VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_RESOURCE_TRACKING_BIT_NV VkDeviceDiagnosticsConfigFlagBitsNV = 0x1 << 1
const VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_AUTOMATIC_CHECKPOINTS_BIT_NV VkDeviceDiagnosticsConfigFlagBitsNV = 0x1 << 2
type VkPipelineCreationFeedbackFlagBitsEXT = int64
const VK_PIPELINE_CREATION_FEEDBACK_VALID_BIT_EXT VkPipelineCreationFeedbackFlagBitsEXT = 0x1 << 0
const VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT_EXT VkPipelineCreationFeedbackFlagBitsEXT = 0x1 << 1
const VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT_EXT VkPipelineCreationFeedbackFlagBitsEXT = 0x1 << 2
type VkFullScreenExclusiveEXT = int32
const VK_FULL_SCREEN_EXCLUSIVE_DEFAULT_EXT VkFullScreenExclusiveEXT = 0
const VK_FULL_SCREEN_EXCLUSIVE_ALLOWED_EXT VkFullScreenExclusiveEXT = 1
const VK_FULL_SCREEN_EXCLUSIVE_DISALLOWED_EXT VkFullScreenExclusiveEXT = 2
const VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT VkFullScreenExclusiveEXT = 3
type VkPerformanceCounterScopeKHR = int32
const VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_BUFFER_KHR VkPerformanceCounterScopeKHR = 0
const VK_PERFORMANCE_COUNTER_SCOPE_RENDER_PASS_KHR VkPerformanceCounterScopeKHR = 1
const VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_KHR VkPerformanceCounterScopeKHR = 2
const VK_QUERY_SCOPE_COMMAND_BUFFER_KHR VkPerformanceCounterScopeKHR = VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_BUFFER_KHR
const VK_QUERY_SCOPE_RENDER_PASS_KHR VkPerformanceCounterScopeKHR = VK_PERFORMANCE_COUNTER_SCOPE_RENDER_PASS_KHR
const VK_QUERY_SCOPE_COMMAND_KHR VkPerformanceCounterScopeKHR = VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_KHR
type VkPerformanceCounterUnitKHR = int32
const VK_PERFORMANCE_COUNTER_UNIT_GENERIC_KHR VkPerformanceCounterUnitKHR = 0
const VK_PERFORMANCE_COUNTER_UNIT_PERCENTAGE_KHR VkPerformanceCounterUnitKHR = 1
const VK_PERFORMANCE_COUNTER_UNIT_NANOSECONDS_KHR VkPerformanceCounterUnitKHR = 2
const VK_PERFORMANCE_COUNTER_UNIT_BYTES_KHR VkPerformanceCounterUnitKHR = 3
const VK_PERFORMANCE_COUNTER_UNIT_BYTES_PER_SECOND_KHR VkPerformanceCounterUnitKHR = 4
const VK_PERFORMANCE_COUNTER_UNIT_KELVIN_KHR VkPerformanceCounterUnitKHR = 5
const VK_PERFORMANCE_COUNTER_UNIT_WATTS_KHR VkPerformanceCounterUnitKHR = 6
const VK_PERFORMANCE_COUNTER_UNIT_VOLTS_KHR VkPerformanceCounterUnitKHR = 7
const VK_PERFORMANCE_COUNTER_UNIT_AMPS_KHR VkPerformanceCounterUnitKHR = 8
const VK_PERFORMANCE_COUNTER_UNIT_HERTZ_KHR VkPerformanceCounterUnitKHR = 9
const VK_PERFORMANCE_COUNTER_UNIT_CYCLES_KHR VkPerformanceCounterUnitKHR = 10
type VkPerformanceCounterStorageKHR = int32
const VK_PERFORMANCE_COUNTER_STORAGE_INT32_KHR VkPerformanceCounterStorageKHR = 0
const VK_PERFORMANCE_COUNTER_STORAGE_INT64_KHR VkPerformanceCounterStorageKHR = 1
const VK_PERFORMANCE_COUNTER_STORAGE_UINT32_KHR VkPerformanceCounterStorageKHR = 2
const VK_PERFORMANCE_COUNTER_STORAGE_UINT64_KHR VkPerformanceCounterStorageKHR = 3
const VK_PERFORMANCE_COUNTER_STORAGE_FLOAT32_KHR VkPerformanceCounterStorageKHR = 4
const VK_PERFORMANCE_COUNTER_STORAGE_FLOAT64_KHR VkPerformanceCounterStorageKHR = 5
type VkPerformanceCounterDescriptionFlagBitsKHR = int64
const VK_PERFORMANCE_COUNTER_DESCRIPTION_PERFORMANCE_IMPACTING_KHR VkPerformanceCounterDescriptionFlagBitsKHR = 0x1 << 0
const VK_PERFORMANCE_COUNTER_DESCRIPTION_CONCURRENTLY_IMPACTED_KHR VkPerformanceCounterDescriptionFlagBitsKHR = 0x1 << 1
type VkAcquireProfilingLockFlagBitsKHR = int64
type VkShaderCorePropertiesFlagBitsAMD = int64
type VkPerformanceConfigurationTypeINTEL = int32
const VK_PERFORMANCE_CONFIGURATION_TYPE_COMMAND_QUEUE_METRICS_DISCOVERY_ACTIVATED_INTEL VkPerformanceConfigurationTypeINTEL = 0
type VkQueryPoolSamplingModeINTEL = int32
const VK_QUERY_POOL_SAMPLING_MODE_MANUAL_INTEL VkQueryPoolSamplingModeINTEL = 0
type VkPerformanceOverrideTypeINTEL = int32
const VK_PERFORMANCE_OVERRIDE_TYPE_NULL_HARDWARE_INTEL VkPerformanceOverrideTypeINTEL = 0
const VK_PERFORMANCE_OVERRIDE_TYPE_FLUSH_GPU_CACHES_INTEL VkPerformanceOverrideTypeINTEL = 1
type VkPerformanceParameterTypeINTEL = int32
const VK_PERFORMANCE_PARAMETER_TYPE_HW_COUNTERS_SUPPORTED_INTEL VkPerformanceParameterTypeINTEL = 0
const VK_PERFORMANCE_PARAMETER_TYPE_STREAM_MARKER_VALID_BITS_INTEL VkPerformanceParameterTypeINTEL = 1
type VkPerformanceValueTypeINTEL = int32
const VK_PERFORMANCE_VALUE_TYPE_UINT32_INTEL VkPerformanceValueTypeINTEL = 0
const VK_PERFORMANCE_VALUE_TYPE_UINT64_INTEL VkPerformanceValueTypeINTEL = 1
const VK_PERFORMANCE_VALUE_TYPE_FLOAT_INTEL VkPerformanceValueTypeINTEL = 2
const VK_PERFORMANCE_VALUE_TYPE_BOOL_INTEL VkPerformanceValueTypeINTEL = 3
const VK_PERFORMANCE_VALUE_TYPE_STRING_INTEL VkPerformanceValueTypeINTEL = 4
type VkShaderFloatControlsIndependence = int32
const VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY VkShaderFloatControlsIndependence = 0
const VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL VkShaderFloatControlsIndependence = 1
const VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE VkShaderFloatControlsIndependence = 2
type VkPipelineExecutableStatisticFormatKHR = int32
const VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_BOOL32_KHR VkPipelineExecutableStatisticFormatKHR = 0
const VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_INT64_KHR VkPipelineExecutableStatisticFormatKHR = 1
const VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_UINT64_KHR VkPipelineExecutableStatisticFormatKHR = 2
const VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_FLOAT64_KHR VkPipelineExecutableStatisticFormatKHR = 3
type VkLineRasterizationModeEXT = int32
const VK_LINE_RASTERIZATION_MODE_DEFAULT_EXT VkLineRasterizationModeEXT = 0
const VK_LINE_RASTERIZATION_MODE_RECTANGULAR_EXT VkLineRasterizationModeEXT = 1
const VK_LINE_RASTERIZATION_MODE_BRESENHAM_EXT VkLineRasterizationModeEXT = 2
const VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_EXT VkLineRasterizationModeEXT = 3
type VkShaderModuleCreateFlagBits = int64
type VkPipelineCompilerControlFlagBitsAMD = int64
type VkToolPurposeFlagBitsEXT = int64
const VK_TOOL_PURPOSE_VALIDATION_BIT_EXT VkToolPurposeFlagBitsEXT = 0x1 << 0
const VK_TOOL_PURPOSE_PROFILING_BIT_EXT VkToolPurposeFlagBitsEXT = 0x1 << 1
const VK_TOOL_PURPOSE_TRACING_BIT_EXT VkToolPurposeFlagBitsEXT = 0x1 << 2
const VK_TOOL_PURPOSE_ADDITIONAL_FEATURES_BIT_EXT VkToolPurposeFlagBitsEXT = 0x1 << 3
const VK_TOOL_PURPOSE_MODIFYING_FEATURES_BIT_EXT VkToolPurposeFlagBitsEXT = 0x1 << 4
type VkFragmentShadingRateCombinerOpKHR = int32
const VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR VkFragmentShadingRateCombinerOpKHR = 0
const VK_FRAGMENT_SHADING_RATE_COMBINER_OP_REPLACE_KHR VkFragmentShadingRateCombinerOpKHR = 1
const VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MIN_KHR VkFragmentShadingRateCombinerOpKHR = 2
const VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MAX_KHR VkFragmentShadingRateCombinerOpKHR = 3
const VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MUL_KHR VkFragmentShadingRateCombinerOpKHR = 4
//----------------------------------------
//--Functions
//----------------------------------------
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INITIALIZATION_FAILED,VK_ERROR_LAYER_NOT_PRESENT,VK_ERROR_EXTENSION_NOT_PRESENT,VK_ERROR_INCOMPATIBLE_DRIVER
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INITIALIZATION_FAILED
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pPhysicalDeviceCount*//*[len]=null-terminated*//*[optional]=true*//*[len]=null-terminated*//*[optional]=false,true*//*[optional]=true*//*[len]=pQueueFamilyPropertyCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_FORMAT_NOT_SUPPORTED
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INITIALIZATION_FAILED,VK_ERROR_EXTENSION_NOT_PRESENT,VK_ERROR_FEATURE_NOT_PRESENT,VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_DEVICE_LOST
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pPropertyCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_LAYER_NOT_PRESENT
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=true*//*[len]=null-terminated*//*[optional]=false,true*//*[optional]=true*//*[len]=pPropertyCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pPropertyCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_LAYER_NOT_PRESENT
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=true*//*[len]=null-terminated*//*[optional]=false,true*//*[optional]=true*//*[len]=pPropertyCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_DEVICE_LOST
//On success: VK_SUCCESS
/*[optional]=true*//*[len]=submitCount*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_DEVICE_LOST
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_DEVICE_LOST
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INVALID_EXTERNAL_HANDLE,VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_MEMORY_MAP_FAILED
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=false,true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[len]=memoryRangeCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[len]=memoryRangeCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=false,true*//*[optional]=true*//*[len]=pSparseMemoryRequirementCount*//*[optional]=false,true*//*[optional]=true*//*[len]=pPropertyCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_DEVICE_LOST
//On success: VK_SUCCESS
//queues: sparse_binding
/*[optional]=true*//*[len]=bindInfoCount*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[len]=fenceCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_DEVICE_LOST
//On success: VK_SUCCESS,VK_NOT_READY
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_DEVICE_LOST
//On success: VK_SUCCESS,VK_TIMEOUT
/*[len]=fenceCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_DEVICE_LOST
//On success: VK_EVENT_SET,VK_EVENT_RESET
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_DEVICE_LOST
//On success: VK_SUCCESS,VK_NOT_READY
/*[len]=dataSize*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INVALID_SHADER_NV
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pDataSize*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[len]=srcCacheCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INVALID_SHADER_NV
//On success: VK_SUCCESS,VK_PIPELINE_COMPILE_REQUIRED_EXT
/*[optional]=true*//*[len]=createInfoCount*//*[optional]=true*//*[len]=createInfoCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INVALID_SHADER_NV
//On success: VK_SUCCESS,VK_PIPELINE_COMPILE_REQUIRED_EXT
/*[optional]=true*//*[len]=createInfoCount*//*[optional]=true*//*[len]=createInfoCount*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_FRAGMENTATION_EXT
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_FRAGMENTED_POOL,VK_ERROR_OUT_OF_POOL_MEMORY
//On success: VK_SUCCESS
/*[len]=pAllocateInfo->descriptorSetCount*///On success: VK_SUCCESS
/*[len]=descriptorSetCount*//*[optional]=true*//*[len]=descriptorWriteCount*//*[optional]=true*//*[len]=descriptorCopyCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[len]=pAllocateInfo->commandBufferCount*//*[len]=commandBufferCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*///queues: graphics,compute
//renderpass: both
//queues: graphics
//renderpass: both
/*[len]=viewportCount*///queues: graphics
//renderpass: both
/*[len]=scissorCount*///queues: graphics
//renderpass: both
//queues: graphics
//renderpass: both
//queues: graphics
//renderpass: both
//queues: graphics
//renderpass: both
//queues: graphics
//renderpass: both
//queues: graphics
//renderpass: both
//queues: graphics
//renderpass: both
//queues: graphics,compute
//renderpass: both
/*[len]=descriptorSetCount*//*[optional]=true*//*[len]=dynamicOffsetCount*///queues: graphics
//renderpass: both
//queues: graphics
//renderpass: both
/*[optional]=false,true*//*[len]=bindingCount*//*[len]=bindingCount*///queues: graphics
//renderpass: inside
//pipeline: graphics
//queues: graphics
//renderpass: inside
//pipeline: graphics
//queues: graphics
//renderpass: inside
//pipeline: graphics
//queues: graphics
//renderpass: inside
//pipeline: graphics
//queues: compute
//renderpass: outside
//pipeline: compute
//queues: compute
//renderpass: outside
//pipeline: compute
//queues: transfer,graphics,compute
//renderpass: outside
//pipeline: transfer
/*[len]=regionCount*///queues: transfer,graphics,compute
//renderpass: outside
//pipeline: transfer
/*[len]=regionCount*///queues: graphics
//renderpass: outside
//pipeline: transfer
/*[len]=regionCount*///queues: transfer,graphics,compute
//renderpass: outside
//pipeline: transfer
/*[len]=regionCount*///queues: transfer,graphics,compute
//renderpass: outside
//pipeline: transfer
/*[len]=regionCount*///queues: transfer,graphics,compute
//renderpass: outside
//pipeline: transfer
/*[len]=dataSize*///queues: transfer,graphics,compute
//renderpass: outside
//pipeline: transfer
//queues: graphics,compute
//renderpass: outside
//pipeline: transfer
/*[len]=rangeCount*///queues: graphics
//renderpass: outside
//pipeline: transfer
/*[len]=rangeCount*///queues: graphics
//renderpass: inside
//pipeline: graphics
/*[len]=attachmentCount*//*[len]=rectCount*///queues: graphics
//renderpass: outside
//pipeline: transfer
/*[len]=regionCount*///queues: graphics,compute
//renderpass: outside
//queues: graphics,compute
//renderpass: outside
//queues: graphics,compute
//renderpass: both
/*[len]=eventCount*//*[optional]=true*//*[len]=memoryBarrierCount*//*[optional]=true*//*[len]=bufferMemoryBarrierCount*//*[optional]=true*//*[len]=imageMemoryBarrierCount*///queues: transfer,graphics,compute
//renderpass: both
/*[optional]=true*//*[optional]=true*//*[len]=memoryBarrierCount*//*[optional]=true*//*[len]=bufferMemoryBarrierCount*//*[optional]=true*//*[len]=imageMemoryBarrierCount*///queues: graphics,compute
//renderpass: both
/*[optional]=true*///queues: graphics,compute
//renderpass: both
//queues: graphics,compute
//renderpass: both
//queues: graphics,compute
//renderpass: both
//queues: graphics,compute
//renderpass: outside
//queues: transfer,graphics,compute
//renderpass: both
//pipeline: transfer
//queues: graphics,compute
//renderpass: outside
//pipeline: transfer
/*[optional]=true*///queues: graphics,compute
//renderpass: both
/*[len]=size*///queues: graphics
//renderpass: outside
//pipeline: graphics
//queues: graphics
//renderpass: inside
//pipeline: graphics
//queues: graphics
//renderpass: inside
//pipeline: graphics
//queues: transfer,graphics,compute
//renderpass: both
/*[len]=commandBufferCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pPropertyCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pPropertyCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pDisplayCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pPropertyCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INITIALIZATION_FAILED
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INCOMPATIBLE_DISPLAY_KHR,VK_ERROR_DEVICE_LOST,VK_ERROR_SURFACE_LOST_KHR
//On success: VK_SUCCESS
/*[len]=swapchainCount*//*[optional]=true*//*[len]=swapchainCount*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_SURFACE_LOST_KHR
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_SURFACE_LOST_KHR
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_SURFACE_LOST_KHR
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pSurfaceFormatCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_SURFACE_LOST_KHR
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pPresentModeCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_DEVICE_LOST,VK_ERROR_SURFACE_LOST_KHR,VK_ERROR_NATIVE_WINDOW_IN_USE_KHR,VK_ERROR_INITIALIZATION_FAILED
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pSwapchainImageCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_DEVICE_LOST,VK_ERROR_OUT_OF_DATE_KHR,VK_ERROR_SURFACE_LOST_KHR,VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT
//On success: VK_SUCCESS,VK_TIMEOUT,VK_NOT_READY,VK_SUBOPTIMAL_KHR
/*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_DEVICE_LOST,VK_ERROR_OUT_OF_DATE_KHR,VK_ERROR_SURFACE_LOST_KHR,VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT
//On success: VK_SUCCESS,VK_SUBOPTIMAL_KHR
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*//*[len]=null-terminated*//*[len]=null-terminated*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
//queues: graphics,compute
//renderpass: both
//queues: graphics,compute
//renderpass: both
//queues: graphics,compute
//renderpass: both
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_FORMAT_NOT_SUPPORTED
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//queues: graphics,compute
//renderpass: inside
//queues: graphics,compute
//renderpass: outside
//queues: graphics,compute
//renderpass: both
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_FORMAT_NOT_SUPPORTED
//On success: VK_SUCCESS
/*[optional]=false,true*//*[optional]=true*//*[len]=pQueueFamilyPropertyCount*//*[optional]=false,true*//*[optional]=true*//*[len]=pPropertyCount*///queues: graphics,compute
//renderpass: both
/*[len]=descriptorWriteCount*//*[optional]=true*///Errors: VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_INVALID_EXTERNAL_HANDLE
//On success: VK_SUCCESS
//Errors: VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_INVALID_EXTERNAL_HANDLE
//On success: VK_SUCCESS
//Errors: VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_INVALID_EXTERNAL_HANDLE
//On success: VK_SUCCESS
//Errors: VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_INVALID_EXTERNAL_HANDLE
//On success: VK_SUCCESS
//Errors: VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_INVALID_EXTERNAL_HANDLE
//On success: VK_SUCCESS
//Errors: VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_INVALID_EXTERNAL_HANDLE
//On success: VK_SUCCESS
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_INITIALIZATION_FAILED
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_DEVICE_LOST,VK_ERROR_OUT_OF_DATE_KHR
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_SURFACE_LOST_KHR
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INITIALIZATION_FAILED
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pPhysicalDeviceGroupCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR
//On success: VK_SUCCESS
/*[len]=bindInfoCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[len]=bindInfoCount*///queues: graphics,compute,transfer
//renderpass: both
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_SURFACE_LOST_KHR
//On success: VK_SUCCESS
/*[optional]=false,true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_DEVICE_LOST,VK_ERROR_OUT_OF_DATE_KHR,VK_ERROR_SURFACE_LOST_KHR,VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT
//On success: VK_SUCCESS,VK_TIMEOUT,VK_NOT_READY,VK_SUBOPTIMAL_KHR
//queues: compute
//renderpass: outside
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pRectCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///queues: graphics,compute
//renderpass: both
/*[len]=swapchainCount*//*[len]=swapchainCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_DEVICE_LOST,VK_ERROR_OUT_OF_DATE_KHR,VK_ERROR_SURFACE_LOST_KHR,VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT
//On success: VK_SUCCESS,VK_SUBOPTIMAL_KHR
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_DEVICE_LOST,VK_ERROR_SURFACE_LOST_KHR
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_DEVICE_LOST,VK_ERROR_OUT_OF_DATE_KHR,VK_ERROR_SURFACE_LOST_KHR
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pPresentationTimingCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
//On success: VK_SUCCESS
/*[optional]=true*///queues: graphics
//renderpass: both
/*[len]=viewportCount*///queues: graphics
//renderpass: both
/*[len]=discardRectangleCount*///queues: graphics
//renderpass: both
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_SURFACE_LOST_KHR
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_SURFACE_LOST_KHR
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pSurfaceFormatCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pPropertyCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pPropertyCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pPropertyCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=false,true*//*[optional]=true*//*[len]=pSparseMemoryRequirementCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pDataSize*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[len]=srcCacheCount*//*[len]=waitSemaphoreCount*///Errors: VK_ERROR_FEATURE_NOT_PRESENT,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pInfoSize*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pTimeDomainCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[len]=timestampCount*//*[len]=timestampCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
//queues: graphics,compute
//renderpass: both
//queues: graphics,compute
//renderpass: both
//queues: graphics,compute
//renderpass: both
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_INVALID_EXTERNAL_HANDLE
//On success: VK_SUCCESS
/*[optional]=false*///queues: transfer,graphics,compute
//renderpass: both
//pipeline: transfer
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*///queues: graphics
//renderpass: outside
//pipeline: graphics
//queues: graphics
//renderpass: inside
//pipeline: graphics
//queues: graphics
//renderpass: inside
//pipeline: graphics
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_DEVICE_LOST
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_DEVICE_LOST
//On success: VK_SUCCESS,VK_TIMEOUT
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_INVALID_EXTERNAL_HANDLE_KHR
//On success: VK_SUCCESS
//Errors: VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//queues: graphics
//renderpass: inside
//pipeline: graphics
//queues: graphics
//renderpass: inside
//pipeline: graphics
//queues: graphics,compute,transfer
//renderpass: both
/*[optional]=false,true*//*[optional]=true*//*[len]=pCheckpointDataCount*///queues: graphics
//renderpass: both
/*[len]=bindingCount*//*[len]=bindingCount*//*[optional]=true*//*[len]=bindingCount*///queues: graphics
//renderpass: inside
/*[optional]=true*//*[len]=counterBufferCount*//*[optional]=true*//*[len]=counterBufferCount*///queues: graphics
//renderpass: inside
/*[optional]=true*//*[len]=counterBufferCount*//*[optional]=true*//*[len]=counterBufferCount*///queues: graphics,compute
//renderpass: both
/*[optional]=true*///queues: graphics,compute
//renderpass: both
//queues: graphics
//renderpass: inside
//pipeline: graphics
//queues: graphics
//renderpass: both
/*[len]=exclusiveScissorCount*///queues: graphics
//renderpass: both
/*[optional]=true*///queues: graphics
//renderpass: both
/*[len]=viewportCount*///queues: graphics
//renderpass: both
/*[optional]=true*//*[len]=customSampleOrderCount*///queues: graphics
//renderpass: inside
//pipeline: graphics
//queues: graphics
//renderpass: inside
//pipeline: graphics
//queues: graphics
//renderpass: inside
//pipeline: graphics
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[len]=bindInfoCount*///queues: compute
//renderpass: outside
//queues: compute
//renderpass: outside
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_OPERATION_DEFERRED_KHR,VK_OPERATION_NOT_DEFERRED_KHR
//queues: compute
//renderpass: outside
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_OPERATION_DEFERRED_KHR,VK_OPERATION_NOT_DEFERRED_KHR
//queues: compute
//renderpass: outside
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_OPERATION_DEFERRED_KHR,VK_OPERATION_NOT_DEFERRED_KHR
//queues: compute
//renderpass: outside
/*[len]=accelerationStructureCount*///queues: compute
//renderpass: outside
/*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[len]=accelerationStructureCount*//*[len]=dataSize*///queues: compute
//renderpass: outside
//queues: compute
//renderpass: outside
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[len]=dataSize*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[len]=dataSize*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[len]=dataSize*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INVALID_SHADER_NV
//On success: VK_SUCCESS,VK_PIPELINE_COMPILE_REQUIRED_EXT
/*[optional]=true*//*[len]=createInfoCount*//*[optional]=true*//*[len]=createInfoCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS
//On success: VK_SUCCESS,VK_OPERATION_DEFERRED_KHR,VK_OPERATION_NOT_DEFERRED_KHR,VK_PIPELINE_COMPILE_REQUIRED_EXT
/*[optional]=true*//*[len]=createInfoCount*//*[optional]=true*//*[len]=createInfoCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pPropertyCount*///queues: compute
//renderpass: outside
//Errors: VK_ERROR_INCOMPATIBLE_VERSION_KHR
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_UNKNOWN
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_SURFACE_LOST_KHR
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pPresentModeCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_SURFACE_LOST_KHR
//On success: VK_SUCCESS
/*[optional]=false,true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INITIALIZATION_FAILED,VK_ERROR_SURFACE_LOST_KHR
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_SURFACE_LOST_KHR
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY,VK_ERROR_INITIALIZATION_FAILED
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pCounterCount*//*[optional]=true*//*[len]=pCounterCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_TIMEOUT
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pCombinationCount*///Errors: VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//queues: graphics,compute,transfer
//renderpass: both
//Errors: VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//queues: graphics,compute,transfer
//renderpass: both
//Errors: VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//queues: graphics,compute,transfer
//renderpass: both
//Errors: VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*///Errors: VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_TOO_MANY_OBJECTS,VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pExecutableCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pStatisticCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pInternalRepresentationCount*///queues: graphics
//renderpass: both
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pToolCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR
//On success: VK_SUCCESS
/*[optional]=true*///queues: compute
//renderpass: outside
/*[len]=infoCount*//*[len]=infoCount*///queues: compute
//renderpass: outside
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_OPERATION_DEFERRED_KHR,VK_OPERATION_NOT_DEFERRED_KHR
/*[len]=infoCount*//*[len]=infoCount*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///On success: VK_SUCCESS,VK_NOT_READY
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY,VK_ERROR_OUT_OF_DEVICE_MEMORY
//On success: VK_SUCCESS,VK_THREAD_DONE_KHR,VK_THREAD_IDLE_KHR
//queues: graphics
//renderpass: both
/*[optional]=true*///queues: graphics
//renderpass: both
//queues: graphics
//renderpass: both
//queues: graphics
//renderpass: both
/*[len]=viewportCount*///queues: graphics
//renderpass: both
/*[len]=scissorCount*///queues: graphics
//renderpass: both
/*[len]=bindingCount*//*[len]=bindingCount*//*[optional]=true*//*[len]=bindingCount*//*[optional]=true*//*[len]=bindingCount*///queues: graphics
//renderpass: both
//queues: graphics
//renderpass: both
//queues: graphics
//renderpass: both
//queues: graphics
//renderpass: both
//queues: graphics
//renderpass: both
//queues: graphics
//renderpass: both
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
/*[optional]=true*//*[optional]=true*//*[optional]=true*///Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS
//queues: transfer,graphics,compute
//renderpass: outside
//pipeline: transfer
//queues: transfer,graphics,compute
//renderpass: outside
//pipeline: transfer
//queues: graphics
//renderpass: outside
//pipeline: transfer
//queues: transfer,graphics,compute
//renderpass: outside
//pipeline: transfer
//queues: transfer,graphics,compute
//renderpass: outside
//pipeline: transfer
//queues: graphics
//renderpass: outside
//pipeline: transfer
//queues: graphics
//renderpass: both
//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On success: VK_SUCCESS,VK_INCOMPLETE
/*[optional]=false,true*//*[optional]=true*//*[len]=pFragmentShadingRateCount*/