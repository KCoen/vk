package vk


import(
	"unsafe"
	"syscall"
	"golang.org/x/sys/windows"
)
var _ unsafe.Pointer

var(
	vulkan = windows.NewLazySystemDLL("vulkan-1.dll")
	procvkCreateInstance = vulkan.NewProc("vkCreateInstance")
	procvkDestroyInstance = vulkan.NewProc("vkDestroyInstance")
	procvkEnumeratePhysicalDevices = vulkan.NewProc("vkEnumeratePhysicalDevices")
	procvkGetDeviceProcAddr = vulkan.NewProc("vkGetDeviceProcAddr")
	procvkGetInstanceProcAddr = vulkan.NewProc("vkGetInstanceProcAddr")
	procvkGetPhysicalDeviceProperties = vulkan.NewProc("vkGetPhysicalDeviceProperties")
	procvkGetPhysicalDeviceQueueFamilyProperties = vulkan.NewProc("vkGetPhysicalDeviceQueueFamilyProperties")
	procvkGetPhysicalDeviceMemoryProperties = vulkan.NewProc("vkGetPhysicalDeviceMemoryProperties")
	procvkGetPhysicalDeviceFeatures = vulkan.NewProc("vkGetPhysicalDeviceFeatures")
	procvkGetPhysicalDeviceFormatProperties = vulkan.NewProc("vkGetPhysicalDeviceFormatProperties")
	procvkGetPhysicalDeviceImageFormatProperties = vulkan.NewProc("vkGetPhysicalDeviceImageFormatProperties")
	procvkCreateDevice = vulkan.NewProc("vkCreateDevice")
	procvkDestroyDevice = vulkan.NewProc("vkDestroyDevice")
	procvkEnumerateInstanceVersion = vulkan.NewProc("vkEnumerateInstanceVersion")
	procvkEnumerateInstanceLayerProperties = vulkan.NewProc("vkEnumerateInstanceLayerProperties")
	procvkEnumerateInstanceExtensionProperties = vulkan.NewProc("vkEnumerateInstanceExtensionProperties")
	procvkEnumerateDeviceLayerProperties = vulkan.NewProc("vkEnumerateDeviceLayerProperties")
	procvkEnumerateDeviceExtensionProperties = vulkan.NewProc("vkEnumerateDeviceExtensionProperties")
	procvkGetDeviceQueue = vulkan.NewProc("vkGetDeviceQueue")
	procvkQueueSubmit = vulkan.NewProc("vkQueueSubmit")
	procvkQueueWaitIdle = vulkan.NewProc("vkQueueWaitIdle")
	procvkDeviceWaitIdle = vulkan.NewProc("vkDeviceWaitIdle")
	procvkAllocateMemory = vulkan.NewProc("vkAllocateMemory")
	procvkFreeMemory = vulkan.NewProc("vkFreeMemory")
	procvkMapMemory = vulkan.NewProc("vkMapMemory")
	procvkUnmapMemory = vulkan.NewProc("vkUnmapMemory")
	procvkFlushMappedMemoryRanges = vulkan.NewProc("vkFlushMappedMemoryRanges")
	procvkInvalidateMappedMemoryRanges = vulkan.NewProc("vkInvalidateMappedMemoryRanges")
	procvkGetDeviceMemoryCommitment = vulkan.NewProc("vkGetDeviceMemoryCommitment")
	procvkGetBufferMemoryRequirements = vulkan.NewProc("vkGetBufferMemoryRequirements")
	procvkBindBufferMemory = vulkan.NewProc("vkBindBufferMemory")
	procvkGetImageMemoryRequirements = vulkan.NewProc("vkGetImageMemoryRequirements")
	procvkBindImageMemory = vulkan.NewProc("vkBindImageMemory")
	procvkGetImageSparseMemoryRequirements = vulkan.NewProc("vkGetImageSparseMemoryRequirements")
	procvkGetPhysicalDeviceSparseImageFormatProperties = vulkan.NewProc("vkGetPhysicalDeviceSparseImageFormatProperties")
	procvkQueueBindSparse = vulkan.NewProc("vkQueueBindSparse")
	procvkCreateFence = vulkan.NewProc("vkCreateFence")
	procvkDestroyFence = vulkan.NewProc("vkDestroyFence")
	procvkResetFences = vulkan.NewProc("vkResetFences")
	procvkGetFenceStatus = vulkan.NewProc("vkGetFenceStatus")
	procvkWaitForFences = vulkan.NewProc("vkWaitForFences")
	procvkCreateSemaphore = vulkan.NewProc("vkCreateSemaphore")
	procvkDestroySemaphore = vulkan.NewProc("vkDestroySemaphore")
	procvkCreateEvent = vulkan.NewProc("vkCreateEvent")
	procvkDestroyEvent = vulkan.NewProc("vkDestroyEvent")
	procvkGetEventStatus = vulkan.NewProc("vkGetEventStatus")
	procvkSetEvent = vulkan.NewProc("vkSetEvent")
	procvkResetEvent = vulkan.NewProc("vkResetEvent")
	procvkCreateQueryPool = vulkan.NewProc("vkCreateQueryPool")
	procvkDestroyQueryPool = vulkan.NewProc("vkDestroyQueryPool")
	procvkGetQueryPoolResults = vulkan.NewProc("vkGetQueryPoolResults")
	procvkResetQueryPool = vulkan.NewProc("vkResetQueryPool")
	procvkResetQueryPoolEXT = vulkan.NewProc("vkResetQueryPoolEXT")
	procvkCreateBuffer = vulkan.NewProc("vkCreateBuffer")
	procvkDestroyBuffer = vulkan.NewProc("vkDestroyBuffer")
	procvkCreateBufferView = vulkan.NewProc("vkCreateBufferView")
	procvkDestroyBufferView = vulkan.NewProc("vkDestroyBufferView")
	procvkCreateImage = vulkan.NewProc("vkCreateImage")
	procvkDestroyImage = vulkan.NewProc("vkDestroyImage")
	procvkGetImageSubresourceLayout = vulkan.NewProc("vkGetImageSubresourceLayout")
	procvkCreateImageView = vulkan.NewProc("vkCreateImageView")
	procvkDestroyImageView = vulkan.NewProc("vkDestroyImageView")
	procvkCreateShaderModule = vulkan.NewProc("vkCreateShaderModule")
	procvkDestroyShaderModule = vulkan.NewProc("vkDestroyShaderModule")
	procvkCreatePipelineCache = vulkan.NewProc("vkCreatePipelineCache")
	procvkDestroyPipelineCache = vulkan.NewProc("vkDestroyPipelineCache")
	procvkGetPipelineCacheData = vulkan.NewProc("vkGetPipelineCacheData")
	procvkMergePipelineCaches = vulkan.NewProc("vkMergePipelineCaches")
	procvkCreateGraphicsPipelines = vulkan.NewProc("vkCreateGraphicsPipelines")
	procvkCreateComputePipelines = vulkan.NewProc("vkCreateComputePipelines")
	procvkDestroyPipeline = vulkan.NewProc("vkDestroyPipeline")
	procvkCreatePipelineLayout = vulkan.NewProc("vkCreatePipelineLayout")
	procvkDestroyPipelineLayout = vulkan.NewProc("vkDestroyPipelineLayout")
	procvkCreateSampler = vulkan.NewProc("vkCreateSampler")
	procvkDestroySampler = vulkan.NewProc("vkDestroySampler")
	procvkCreateDescriptorSetLayout = vulkan.NewProc("vkCreateDescriptorSetLayout")
	procvkDestroyDescriptorSetLayout = vulkan.NewProc("vkDestroyDescriptorSetLayout")
	procvkCreateDescriptorPool = vulkan.NewProc("vkCreateDescriptorPool")
	procvkDestroyDescriptorPool = vulkan.NewProc("vkDestroyDescriptorPool")
	procvkResetDescriptorPool = vulkan.NewProc("vkResetDescriptorPool")
	procvkAllocateDescriptorSets = vulkan.NewProc("vkAllocateDescriptorSets")
	procvkFreeDescriptorSets = vulkan.NewProc("vkFreeDescriptorSets")
	procvkUpdateDescriptorSets = vulkan.NewProc("vkUpdateDescriptorSets")
	procvkCreateFramebuffer = vulkan.NewProc("vkCreateFramebuffer")
	procvkDestroyFramebuffer = vulkan.NewProc("vkDestroyFramebuffer")
	procvkCreateRenderPass = vulkan.NewProc("vkCreateRenderPass")
	procvkDestroyRenderPass = vulkan.NewProc("vkDestroyRenderPass")
	procvkGetRenderAreaGranularity = vulkan.NewProc("vkGetRenderAreaGranularity")
	procvkCreateCommandPool = vulkan.NewProc("vkCreateCommandPool")
	procvkDestroyCommandPool = vulkan.NewProc("vkDestroyCommandPool")
	procvkResetCommandPool = vulkan.NewProc("vkResetCommandPool")
	procvkAllocateCommandBuffers = vulkan.NewProc("vkAllocateCommandBuffers")
	procvkFreeCommandBuffers = vulkan.NewProc("vkFreeCommandBuffers")
	procvkBeginCommandBuffer = vulkan.NewProc("vkBeginCommandBuffer")
	procvkEndCommandBuffer = vulkan.NewProc("vkEndCommandBuffer")
	procvkResetCommandBuffer = vulkan.NewProc("vkResetCommandBuffer")
	procvkCmdBindPipeline = vulkan.NewProc("vkCmdBindPipeline")
	procvkCmdSetViewport = vulkan.NewProc("vkCmdSetViewport")
	procvkCmdSetScissor = vulkan.NewProc("vkCmdSetScissor")
	procvkCmdSetLineWidth = vulkan.NewProc("vkCmdSetLineWidth")
	procvkCmdSetDepthBias = vulkan.NewProc("vkCmdSetDepthBias")
	procvkCmdSetBlendConstants = vulkan.NewProc("vkCmdSetBlendConstants")
	procvkCmdSetDepthBounds = vulkan.NewProc("vkCmdSetDepthBounds")
	procvkCmdSetStencilCompareMask = vulkan.NewProc("vkCmdSetStencilCompareMask")
	procvkCmdSetStencilWriteMask = vulkan.NewProc("vkCmdSetStencilWriteMask")
	procvkCmdSetStencilReference = vulkan.NewProc("vkCmdSetStencilReference")
	procvkCmdBindDescriptorSets = vulkan.NewProc("vkCmdBindDescriptorSets")
	procvkCmdBindIndexBuffer = vulkan.NewProc("vkCmdBindIndexBuffer")
	procvkCmdBindVertexBuffers = vulkan.NewProc("vkCmdBindVertexBuffers")
	procvkCmdDraw = vulkan.NewProc("vkCmdDraw")
	procvkCmdDrawIndexed = vulkan.NewProc("vkCmdDrawIndexed")
	procvkCmdDrawIndirect = vulkan.NewProc("vkCmdDrawIndirect")
	procvkCmdDrawIndexedIndirect = vulkan.NewProc("vkCmdDrawIndexedIndirect")
	procvkCmdDispatch = vulkan.NewProc("vkCmdDispatch")
	procvkCmdDispatchIndirect = vulkan.NewProc("vkCmdDispatchIndirect")
	procvkCmdCopyBuffer = vulkan.NewProc("vkCmdCopyBuffer")
	procvkCmdCopyImage = vulkan.NewProc("vkCmdCopyImage")
	procvkCmdBlitImage = vulkan.NewProc("vkCmdBlitImage")
	procvkCmdCopyBufferToImage = vulkan.NewProc("vkCmdCopyBufferToImage")
	procvkCmdCopyImageToBuffer = vulkan.NewProc("vkCmdCopyImageToBuffer")
	procvkCmdUpdateBuffer = vulkan.NewProc("vkCmdUpdateBuffer")
	procvkCmdFillBuffer = vulkan.NewProc("vkCmdFillBuffer")
	procvkCmdClearColorImage = vulkan.NewProc("vkCmdClearColorImage")
	procvkCmdClearDepthStencilImage = vulkan.NewProc("vkCmdClearDepthStencilImage")
	procvkCmdClearAttachments = vulkan.NewProc("vkCmdClearAttachments")
	procvkCmdResolveImage = vulkan.NewProc("vkCmdResolveImage")
	procvkCmdSetEvent = vulkan.NewProc("vkCmdSetEvent")
	procvkCmdResetEvent = vulkan.NewProc("vkCmdResetEvent")
	procvkCmdWaitEvents = vulkan.NewProc("vkCmdWaitEvents")
	procvkCmdPipelineBarrier = vulkan.NewProc("vkCmdPipelineBarrier")
	procvkCmdBeginQuery = vulkan.NewProc("vkCmdBeginQuery")
	procvkCmdEndQuery = vulkan.NewProc("vkCmdEndQuery")
	procvkCmdBeginConditionalRenderingEXT = vulkan.NewProc("vkCmdBeginConditionalRenderingEXT")
	procvkCmdEndConditionalRenderingEXT = vulkan.NewProc("vkCmdEndConditionalRenderingEXT")
	procvkCmdResetQueryPool = vulkan.NewProc("vkCmdResetQueryPool")
	procvkCmdWriteTimestamp = vulkan.NewProc("vkCmdWriteTimestamp")
	procvkCmdCopyQueryPoolResults = vulkan.NewProc("vkCmdCopyQueryPoolResults")
	procvkCmdPushConstants = vulkan.NewProc("vkCmdPushConstants")
	procvkCmdBeginRenderPass = vulkan.NewProc("vkCmdBeginRenderPass")
	procvkCmdNextSubpass = vulkan.NewProc("vkCmdNextSubpass")
	procvkCmdEndRenderPass = vulkan.NewProc("vkCmdEndRenderPass")
	procvkCmdExecuteCommands = vulkan.NewProc("vkCmdExecuteCommands")
	procvkCreateAndroidSurfaceKHR = vulkan.NewProc("vkCreateAndroidSurfaceKHR")
	procvkGetPhysicalDeviceDisplayPropertiesKHR = vulkan.NewProc("vkGetPhysicalDeviceDisplayPropertiesKHR")
	procvkGetPhysicalDeviceDisplayPlanePropertiesKHR = vulkan.NewProc("vkGetPhysicalDeviceDisplayPlanePropertiesKHR")
	procvkGetDisplayPlaneSupportedDisplaysKHR = vulkan.NewProc("vkGetDisplayPlaneSupportedDisplaysKHR")
	procvkGetDisplayModePropertiesKHR = vulkan.NewProc("vkGetDisplayModePropertiesKHR")
	procvkCreateDisplayModeKHR = vulkan.NewProc("vkCreateDisplayModeKHR")
	procvkGetDisplayPlaneCapabilitiesKHR = vulkan.NewProc("vkGetDisplayPlaneCapabilitiesKHR")
	procvkCreateDisplayPlaneSurfaceKHR = vulkan.NewProc("vkCreateDisplayPlaneSurfaceKHR")
	procvkCreateSharedSwapchainsKHR = vulkan.NewProc("vkCreateSharedSwapchainsKHR")
	procvkDestroySurfaceKHR = vulkan.NewProc("vkDestroySurfaceKHR")
	procvkGetPhysicalDeviceSurfaceSupportKHR = vulkan.NewProc("vkGetPhysicalDeviceSurfaceSupportKHR")
	procvkGetPhysicalDeviceSurfaceCapabilitiesKHR = vulkan.NewProc("vkGetPhysicalDeviceSurfaceCapabilitiesKHR")
	procvkGetPhysicalDeviceSurfaceFormatsKHR = vulkan.NewProc("vkGetPhysicalDeviceSurfaceFormatsKHR")
	procvkGetPhysicalDeviceSurfacePresentModesKHR = vulkan.NewProc("vkGetPhysicalDeviceSurfacePresentModesKHR")
	procvkCreateSwapchainKHR = vulkan.NewProc("vkCreateSwapchainKHR")
	procvkDestroySwapchainKHR = vulkan.NewProc("vkDestroySwapchainKHR")
	procvkGetSwapchainImagesKHR = vulkan.NewProc("vkGetSwapchainImagesKHR")
	procvkAcquireNextImageKHR = vulkan.NewProc("vkAcquireNextImageKHR")
	procvkQueuePresentKHR = vulkan.NewProc("vkQueuePresentKHR")
	procvkCreateViSurfaceNN = vulkan.NewProc("vkCreateViSurfaceNN")
	procvkCreateWaylandSurfaceKHR = vulkan.NewProc("vkCreateWaylandSurfaceKHR")
	procvkGetPhysicalDeviceWaylandPresentationSupportKHR = vulkan.NewProc("vkGetPhysicalDeviceWaylandPresentationSupportKHR")
	procvkCreateWin32SurfaceKHR = vulkan.NewProc("vkCreateWin32SurfaceKHR")
	procvkGetPhysicalDeviceWin32PresentationSupportKHR = vulkan.NewProc("vkGetPhysicalDeviceWin32PresentationSupportKHR")
	procvkCreateXlibSurfaceKHR = vulkan.NewProc("vkCreateXlibSurfaceKHR")
	procvkGetPhysicalDeviceXlibPresentationSupportKHR = vulkan.NewProc("vkGetPhysicalDeviceXlibPresentationSupportKHR")
	procvkCreateXcbSurfaceKHR = vulkan.NewProc("vkCreateXcbSurfaceKHR")
	procvkGetPhysicalDeviceXcbPresentationSupportKHR = vulkan.NewProc("vkGetPhysicalDeviceXcbPresentationSupportKHR")
	procvkCreateDirectFBSurfaceEXT = vulkan.NewProc("vkCreateDirectFBSurfaceEXT")
	procvkGetPhysicalDeviceDirectFBPresentationSupportEXT = vulkan.NewProc("vkGetPhysicalDeviceDirectFBPresentationSupportEXT")
	procvkCreateImagePipeSurfaceFUCHSIA = vulkan.NewProc("vkCreateImagePipeSurfaceFUCHSIA")
	procvkCreateStreamDescriptorSurfaceGGP = vulkan.NewProc("vkCreateStreamDescriptorSurfaceGGP")
	procvkCreateDebugReportCallbackEXT = vulkan.NewProc("vkCreateDebugReportCallbackEXT")
	procvkDestroyDebugReportCallbackEXT = vulkan.NewProc("vkDestroyDebugReportCallbackEXT")
	procvkDebugReportMessageEXT = vulkan.NewProc("vkDebugReportMessageEXT")
	procvkDebugMarkerSetObjectNameEXT = vulkan.NewProc("vkDebugMarkerSetObjectNameEXT")
	procvkDebugMarkerSetObjectTagEXT = vulkan.NewProc("vkDebugMarkerSetObjectTagEXT")
	procvkCmdDebugMarkerBeginEXT = vulkan.NewProc("vkCmdDebugMarkerBeginEXT")
	procvkCmdDebugMarkerEndEXT = vulkan.NewProc("vkCmdDebugMarkerEndEXT")
	procvkCmdDebugMarkerInsertEXT = vulkan.NewProc("vkCmdDebugMarkerInsertEXT")
	procvkGetPhysicalDeviceExternalImageFormatPropertiesNV = vulkan.NewProc("vkGetPhysicalDeviceExternalImageFormatPropertiesNV")
	procvkGetMemoryWin32HandleNV = vulkan.NewProc("vkGetMemoryWin32HandleNV")
	procvkCmdExecuteGeneratedCommandsNV = vulkan.NewProc("vkCmdExecuteGeneratedCommandsNV")
	procvkCmdPreprocessGeneratedCommandsNV = vulkan.NewProc("vkCmdPreprocessGeneratedCommandsNV")
	procvkCmdBindPipelineShaderGroupNV = vulkan.NewProc("vkCmdBindPipelineShaderGroupNV")
	procvkGetGeneratedCommandsMemoryRequirementsNV = vulkan.NewProc("vkGetGeneratedCommandsMemoryRequirementsNV")
	procvkCreateIndirectCommandsLayoutNV = vulkan.NewProc("vkCreateIndirectCommandsLayoutNV")
	procvkDestroyIndirectCommandsLayoutNV = vulkan.NewProc("vkDestroyIndirectCommandsLayoutNV")
	procvkGetPhysicalDeviceFeatures2 = vulkan.NewProc("vkGetPhysicalDeviceFeatures2")
	procvkGetPhysicalDeviceFeatures2KHR = vulkan.NewProc("vkGetPhysicalDeviceFeatures2KHR")
	procvkGetPhysicalDeviceProperties2 = vulkan.NewProc("vkGetPhysicalDeviceProperties2")
	procvkGetPhysicalDeviceProperties2KHR = vulkan.NewProc("vkGetPhysicalDeviceProperties2KHR")
	procvkGetPhysicalDeviceFormatProperties2 = vulkan.NewProc("vkGetPhysicalDeviceFormatProperties2")
	procvkGetPhysicalDeviceFormatProperties2KHR = vulkan.NewProc("vkGetPhysicalDeviceFormatProperties2KHR")
	procvkGetPhysicalDeviceImageFormatProperties2 = vulkan.NewProc("vkGetPhysicalDeviceImageFormatProperties2")
	procvkGetPhysicalDeviceImageFormatProperties2KHR = vulkan.NewProc("vkGetPhysicalDeviceImageFormatProperties2KHR")
	procvkGetPhysicalDeviceQueueFamilyProperties2 = vulkan.NewProc("vkGetPhysicalDeviceQueueFamilyProperties2")
	procvkGetPhysicalDeviceQueueFamilyProperties2KHR = vulkan.NewProc("vkGetPhysicalDeviceQueueFamilyProperties2KHR")
	procvkGetPhysicalDeviceMemoryProperties2 = vulkan.NewProc("vkGetPhysicalDeviceMemoryProperties2")
	procvkGetPhysicalDeviceMemoryProperties2KHR = vulkan.NewProc("vkGetPhysicalDeviceMemoryProperties2KHR")
	procvkGetPhysicalDeviceSparseImageFormatProperties2 = vulkan.NewProc("vkGetPhysicalDeviceSparseImageFormatProperties2")
	procvkGetPhysicalDeviceSparseImageFormatProperties2KHR = vulkan.NewProc("vkGetPhysicalDeviceSparseImageFormatProperties2KHR")
	procvkCmdPushDescriptorSetKHR = vulkan.NewProc("vkCmdPushDescriptorSetKHR")
	procvkTrimCommandPool = vulkan.NewProc("vkTrimCommandPool")
	procvkTrimCommandPoolKHR = vulkan.NewProc("vkTrimCommandPoolKHR")
	procvkGetPhysicalDeviceExternalBufferProperties = vulkan.NewProc("vkGetPhysicalDeviceExternalBufferProperties")
	procvkGetPhysicalDeviceExternalBufferPropertiesKHR = vulkan.NewProc("vkGetPhysicalDeviceExternalBufferPropertiesKHR")
	procvkGetMemoryWin32HandleKHR = vulkan.NewProc("vkGetMemoryWin32HandleKHR")
	procvkGetMemoryWin32HandlePropertiesKHR = vulkan.NewProc("vkGetMemoryWin32HandlePropertiesKHR")
	procvkGetMemoryFdKHR = vulkan.NewProc("vkGetMemoryFdKHR")
	procvkGetMemoryFdPropertiesKHR = vulkan.NewProc("vkGetMemoryFdPropertiesKHR")
	procvkGetPhysicalDeviceExternalSemaphoreProperties = vulkan.NewProc("vkGetPhysicalDeviceExternalSemaphoreProperties")
	procvkGetPhysicalDeviceExternalSemaphorePropertiesKHR = vulkan.NewProc("vkGetPhysicalDeviceExternalSemaphorePropertiesKHR")
	procvkGetSemaphoreWin32HandleKHR = vulkan.NewProc("vkGetSemaphoreWin32HandleKHR")
	procvkImportSemaphoreWin32HandleKHR = vulkan.NewProc("vkImportSemaphoreWin32HandleKHR")
	procvkGetSemaphoreFdKHR = vulkan.NewProc("vkGetSemaphoreFdKHR")
	procvkImportSemaphoreFdKHR = vulkan.NewProc("vkImportSemaphoreFdKHR")
	procvkGetPhysicalDeviceExternalFenceProperties = vulkan.NewProc("vkGetPhysicalDeviceExternalFenceProperties")
	procvkGetPhysicalDeviceExternalFencePropertiesKHR = vulkan.NewProc("vkGetPhysicalDeviceExternalFencePropertiesKHR")
	procvkGetFenceWin32HandleKHR = vulkan.NewProc("vkGetFenceWin32HandleKHR")
	procvkImportFenceWin32HandleKHR = vulkan.NewProc("vkImportFenceWin32HandleKHR")
	procvkGetFenceFdKHR = vulkan.NewProc("vkGetFenceFdKHR")
	procvkImportFenceFdKHR = vulkan.NewProc("vkImportFenceFdKHR")
	procvkReleaseDisplayEXT = vulkan.NewProc("vkReleaseDisplayEXT")
	procvkAcquireXlibDisplayEXT = vulkan.NewProc("vkAcquireXlibDisplayEXT")
	procvkGetRandROutputDisplayEXT = vulkan.NewProc("vkGetRandROutputDisplayEXT")
	procvkDisplayPowerControlEXT = vulkan.NewProc("vkDisplayPowerControlEXT")
	procvkRegisterDeviceEventEXT = vulkan.NewProc("vkRegisterDeviceEventEXT")
	procvkRegisterDisplayEventEXT = vulkan.NewProc("vkRegisterDisplayEventEXT")
	procvkGetSwapchainCounterEXT = vulkan.NewProc("vkGetSwapchainCounterEXT")
	procvkGetPhysicalDeviceSurfaceCapabilities2EXT = vulkan.NewProc("vkGetPhysicalDeviceSurfaceCapabilities2EXT")
	procvkEnumeratePhysicalDeviceGroups = vulkan.NewProc("vkEnumeratePhysicalDeviceGroups")
	procvkEnumeratePhysicalDeviceGroupsKHR = vulkan.NewProc("vkEnumeratePhysicalDeviceGroupsKHR")
	procvkGetDeviceGroupPeerMemoryFeatures = vulkan.NewProc("vkGetDeviceGroupPeerMemoryFeatures")
	procvkGetDeviceGroupPeerMemoryFeaturesKHR = vulkan.NewProc("vkGetDeviceGroupPeerMemoryFeaturesKHR")
	procvkBindBufferMemory2 = vulkan.NewProc("vkBindBufferMemory2")
	procvkBindBufferMemory2KHR = vulkan.NewProc("vkBindBufferMemory2KHR")
	procvkBindImageMemory2 = vulkan.NewProc("vkBindImageMemory2")
	procvkBindImageMemory2KHR = vulkan.NewProc("vkBindImageMemory2KHR")
	procvkCmdSetDeviceMask = vulkan.NewProc("vkCmdSetDeviceMask")
	procvkCmdSetDeviceMaskKHR = vulkan.NewProc("vkCmdSetDeviceMaskKHR")
	procvkGetDeviceGroupPresentCapabilitiesKHR = vulkan.NewProc("vkGetDeviceGroupPresentCapabilitiesKHR")
	procvkGetDeviceGroupSurfacePresentModesKHR = vulkan.NewProc("vkGetDeviceGroupSurfacePresentModesKHR")
	procvkAcquireNextImage2KHR = vulkan.NewProc("vkAcquireNextImage2KHR")
	procvkCmdDispatchBase = vulkan.NewProc("vkCmdDispatchBase")
	procvkCmdDispatchBaseKHR = vulkan.NewProc("vkCmdDispatchBaseKHR")
	procvkGetPhysicalDevicePresentRectanglesKHR = vulkan.NewProc("vkGetPhysicalDevicePresentRectanglesKHR")
	procvkCreateDescriptorUpdateTemplate = vulkan.NewProc("vkCreateDescriptorUpdateTemplate")
	procvkCreateDescriptorUpdateTemplateKHR = vulkan.NewProc("vkCreateDescriptorUpdateTemplateKHR")
	procvkDestroyDescriptorUpdateTemplate = vulkan.NewProc("vkDestroyDescriptorUpdateTemplate")
	procvkDestroyDescriptorUpdateTemplateKHR = vulkan.NewProc("vkDestroyDescriptorUpdateTemplateKHR")
	procvkUpdateDescriptorSetWithTemplate = vulkan.NewProc("vkUpdateDescriptorSetWithTemplate")
	procvkUpdateDescriptorSetWithTemplateKHR = vulkan.NewProc("vkUpdateDescriptorSetWithTemplateKHR")
	procvkCmdPushDescriptorSetWithTemplateKHR = vulkan.NewProc("vkCmdPushDescriptorSetWithTemplateKHR")
	procvkSetHdrMetadataEXT = vulkan.NewProc("vkSetHdrMetadataEXT")
	procvkGetSwapchainStatusKHR = vulkan.NewProc("vkGetSwapchainStatusKHR")
	procvkGetRefreshCycleDurationGOOGLE = vulkan.NewProc("vkGetRefreshCycleDurationGOOGLE")
	procvkGetPastPresentationTimingGOOGLE = vulkan.NewProc("vkGetPastPresentationTimingGOOGLE")
	procvkCreateIOSSurfaceMVK = vulkan.NewProc("vkCreateIOSSurfaceMVK")
	procvkCreateMacOSSurfaceMVK = vulkan.NewProc("vkCreateMacOSSurfaceMVK")
	procvkCreateMetalSurfaceEXT = vulkan.NewProc("vkCreateMetalSurfaceEXT")
	procvkCmdSetViewportWScalingNV = vulkan.NewProc("vkCmdSetViewportWScalingNV")
	procvkCmdSetDiscardRectangleEXT = vulkan.NewProc("vkCmdSetDiscardRectangleEXT")
	procvkCmdSetSampleLocationsEXT = vulkan.NewProc("vkCmdSetSampleLocationsEXT")
	procvkGetPhysicalDeviceMultisamplePropertiesEXT = vulkan.NewProc("vkGetPhysicalDeviceMultisamplePropertiesEXT")
	procvkGetPhysicalDeviceSurfaceCapabilities2KHR = vulkan.NewProc("vkGetPhysicalDeviceSurfaceCapabilities2KHR")
	procvkGetPhysicalDeviceSurfaceFormats2KHR = vulkan.NewProc("vkGetPhysicalDeviceSurfaceFormats2KHR")
	procvkGetPhysicalDeviceDisplayProperties2KHR = vulkan.NewProc("vkGetPhysicalDeviceDisplayProperties2KHR")
	procvkGetPhysicalDeviceDisplayPlaneProperties2KHR = vulkan.NewProc("vkGetPhysicalDeviceDisplayPlaneProperties2KHR")
	procvkGetDisplayModeProperties2KHR = vulkan.NewProc("vkGetDisplayModeProperties2KHR")
	procvkGetDisplayPlaneCapabilities2KHR = vulkan.NewProc("vkGetDisplayPlaneCapabilities2KHR")
	procvkGetBufferMemoryRequirements2 = vulkan.NewProc("vkGetBufferMemoryRequirements2")
	procvkGetBufferMemoryRequirements2KHR = vulkan.NewProc("vkGetBufferMemoryRequirements2KHR")
	procvkGetImageMemoryRequirements2 = vulkan.NewProc("vkGetImageMemoryRequirements2")
	procvkGetImageMemoryRequirements2KHR = vulkan.NewProc("vkGetImageMemoryRequirements2KHR")
	procvkGetImageSparseMemoryRequirements2 = vulkan.NewProc("vkGetImageSparseMemoryRequirements2")
	procvkGetImageSparseMemoryRequirements2KHR = vulkan.NewProc("vkGetImageSparseMemoryRequirements2KHR")
	procvkCreateSamplerYcbcrConversion = vulkan.NewProc("vkCreateSamplerYcbcrConversion")
	procvkCreateSamplerYcbcrConversionKHR = vulkan.NewProc("vkCreateSamplerYcbcrConversionKHR")
	procvkDestroySamplerYcbcrConversion = vulkan.NewProc("vkDestroySamplerYcbcrConversion")
	procvkDestroySamplerYcbcrConversionKHR = vulkan.NewProc("vkDestroySamplerYcbcrConversionKHR")
	procvkGetDeviceQueue2 = vulkan.NewProc("vkGetDeviceQueue2")
	procvkCreateValidationCacheEXT = vulkan.NewProc("vkCreateValidationCacheEXT")
	procvkDestroyValidationCacheEXT = vulkan.NewProc("vkDestroyValidationCacheEXT")
	procvkGetValidationCacheDataEXT = vulkan.NewProc("vkGetValidationCacheDataEXT")
	procvkMergeValidationCachesEXT = vulkan.NewProc("vkMergeValidationCachesEXT")
	procvkGetDescriptorSetLayoutSupport = vulkan.NewProc("vkGetDescriptorSetLayoutSupport")
	procvkGetDescriptorSetLayoutSupportKHR = vulkan.NewProc("vkGetDescriptorSetLayoutSupportKHR")
	procvkGetSwapchainGrallocUsageANDROID = vulkan.NewProc("vkGetSwapchainGrallocUsageANDROID")
	procvkGetSwapchainGrallocUsage2ANDROID = vulkan.NewProc("vkGetSwapchainGrallocUsage2ANDROID")
	procvkAcquireImageANDROID = vulkan.NewProc("vkAcquireImageANDROID")
	procvkQueueSignalReleaseImageANDROID = vulkan.NewProc("vkQueueSignalReleaseImageANDROID")
	procvkGetShaderInfoAMD = vulkan.NewProc("vkGetShaderInfoAMD")
	procvkSetLocalDimmingAMD = vulkan.NewProc("vkSetLocalDimmingAMD")
	procvkGetPhysicalDeviceCalibrateableTimeDomainsEXT = vulkan.NewProc("vkGetPhysicalDeviceCalibrateableTimeDomainsEXT")
	procvkGetCalibratedTimestampsEXT = vulkan.NewProc("vkGetCalibratedTimestampsEXT")
	procvkSetDebugUtilsObjectNameEXT = vulkan.NewProc("vkSetDebugUtilsObjectNameEXT")
	procvkSetDebugUtilsObjectTagEXT = vulkan.NewProc("vkSetDebugUtilsObjectTagEXT")
	procvkQueueBeginDebugUtilsLabelEXT = vulkan.NewProc("vkQueueBeginDebugUtilsLabelEXT")
	procvkQueueEndDebugUtilsLabelEXT = vulkan.NewProc("vkQueueEndDebugUtilsLabelEXT")
	procvkQueueInsertDebugUtilsLabelEXT = vulkan.NewProc("vkQueueInsertDebugUtilsLabelEXT")
	procvkCmdBeginDebugUtilsLabelEXT = vulkan.NewProc("vkCmdBeginDebugUtilsLabelEXT")
	procvkCmdEndDebugUtilsLabelEXT = vulkan.NewProc("vkCmdEndDebugUtilsLabelEXT")
	procvkCmdInsertDebugUtilsLabelEXT = vulkan.NewProc("vkCmdInsertDebugUtilsLabelEXT")
	procvkCreateDebugUtilsMessengerEXT = vulkan.NewProc("vkCreateDebugUtilsMessengerEXT")
	procvkDestroyDebugUtilsMessengerEXT = vulkan.NewProc("vkDestroyDebugUtilsMessengerEXT")
	procvkSubmitDebugUtilsMessageEXT = vulkan.NewProc("vkSubmitDebugUtilsMessageEXT")
	procvkGetMemoryHostPointerPropertiesEXT = vulkan.NewProc("vkGetMemoryHostPointerPropertiesEXT")
	procvkCmdWriteBufferMarkerAMD = vulkan.NewProc("vkCmdWriteBufferMarkerAMD")
	procvkCreateRenderPass2 = vulkan.NewProc("vkCreateRenderPass2")
	procvkCreateRenderPass2KHR = vulkan.NewProc("vkCreateRenderPass2KHR")
	procvkCmdBeginRenderPass2 = vulkan.NewProc("vkCmdBeginRenderPass2")
	procvkCmdBeginRenderPass2KHR = vulkan.NewProc("vkCmdBeginRenderPass2KHR")
	procvkCmdNextSubpass2 = vulkan.NewProc("vkCmdNextSubpass2")
	procvkCmdNextSubpass2KHR = vulkan.NewProc("vkCmdNextSubpass2KHR")
	procvkCmdEndRenderPass2 = vulkan.NewProc("vkCmdEndRenderPass2")
	procvkCmdEndRenderPass2KHR = vulkan.NewProc("vkCmdEndRenderPass2KHR")
	procvkGetSemaphoreCounterValue = vulkan.NewProc("vkGetSemaphoreCounterValue")
	procvkGetSemaphoreCounterValueKHR = vulkan.NewProc("vkGetSemaphoreCounterValueKHR")
	procvkWaitSemaphores = vulkan.NewProc("vkWaitSemaphores")
	procvkWaitSemaphoresKHR = vulkan.NewProc("vkWaitSemaphoresKHR")
	procvkSignalSemaphore = vulkan.NewProc("vkSignalSemaphore")
	procvkSignalSemaphoreKHR = vulkan.NewProc("vkSignalSemaphoreKHR")
	procvkGetAndroidHardwareBufferPropertiesANDROID = vulkan.NewProc("vkGetAndroidHardwareBufferPropertiesANDROID")
	procvkGetMemoryAndroidHardwareBufferANDROID = vulkan.NewProc("vkGetMemoryAndroidHardwareBufferANDROID")
	procvkCmdDrawIndirectCount = vulkan.NewProc("vkCmdDrawIndirectCount")
	procvkCmdDrawIndirectCountKHR = vulkan.NewProc("vkCmdDrawIndirectCountKHR")
	procvkCmdDrawIndirectCountAMD = vulkan.NewProc("vkCmdDrawIndirectCountAMD")
	procvkCmdDrawIndexedIndirectCount = vulkan.NewProc("vkCmdDrawIndexedIndirectCount")
	procvkCmdDrawIndexedIndirectCountKHR = vulkan.NewProc("vkCmdDrawIndexedIndirectCountKHR")
	procvkCmdDrawIndexedIndirectCountAMD = vulkan.NewProc("vkCmdDrawIndexedIndirectCountAMD")
	procvkCmdSetCheckpointNV = vulkan.NewProc("vkCmdSetCheckpointNV")
	procvkGetQueueCheckpointDataNV = vulkan.NewProc("vkGetQueueCheckpointDataNV")
	procvkCmdBindTransformFeedbackBuffersEXT = vulkan.NewProc("vkCmdBindTransformFeedbackBuffersEXT")
	procvkCmdBeginTransformFeedbackEXT = vulkan.NewProc("vkCmdBeginTransformFeedbackEXT")
	procvkCmdEndTransformFeedbackEXT = vulkan.NewProc("vkCmdEndTransformFeedbackEXT")
	procvkCmdBeginQueryIndexedEXT = vulkan.NewProc("vkCmdBeginQueryIndexedEXT")
	procvkCmdEndQueryIndexedEXT = vulkan.NewProc("vkCmdEndQueryIndexedEXT")
	procvkCmdDrawIndirectByteCountEXT = vulkan.NewProc("vkCmdDrawIndirectByteCountEXT")
	procvkCmdSetExclusiveScissorNV = vulkan.NewProc("vkCmdSetExclusiveScissorNV")
	procvkCmdBindShadingRateImageNV = vulkan.NewProc("vkCmdBindShadingRateImageNV")
	procvkCmdSetViewportShadingRatePaletteNV = vulkan.NewProc("vkCmdSetViewportShadingRatePaletteNV")
	procvkCmdSetCoarseSampleOrderNV = vulkan.NewProc("vkCmdSetCoarseSampleOrderNV")
	procvkCmdDrawMeshTasksNV = vulkan.NewProc("vkCmdDrawMeshTasksNV")
	procvkCmdDrawMeshTasksIndirectNV = vulkan.NewProc("vkCmdDrawMeshTasksIndirectNV")
	procvkCmdDrawMeshTasksIndirectCountNV = vulkan.NewProc("vkCmdDrawMeshTasksIndirectCountNV")
	procvkCompileDeferredNV = vulkan.NewProc("vkCompileDeferredNV")
	procvkCreateAccelerationStructureNV = vulkan.NewProc("vkCreateAccelerationStructureNV")
	procvkDestroyAccelerationStructureKHR = vulkan.NewProc("vkDestroyAccelerationStructureKHR")
	procvkDestroyAccelerationStructureNV = vulkan.NewProc("vkDestroyAccelerationStructureNV")
	procvkGetAccelerationStructureMemoryRequirementsKHR = vulkan.NewProc("vkGetAccelerationStructureMemoryRequirementsKHR")
	procvkGetAccelerationStructureMemoryRequirementsNV = vulkan.NewProc("vkGetAccelerationStructureMemoryRequirementsNV")
	procvkBindAccelerationStructureMemoryKHR = vulkan.NewProc("vkBindAccelerationStructureMemoryKHR")
	procvkBindAccelerationStructureMemoryNV = vulkan.NewProc("vkBindAccelerationStructureMemoryNV")
	procvkCmdCopyAccelerationStructureNV = vulkan.NewProc("vkCmdCopyAccelerationStructureNV")
	procvkCmdCopyAccelerationStructureKHR = vulkan.NewProc("vkCmdCopyAccelerationStructureKHR")
	procvkCopyAccelerationStructureKHR = vulkan.NewProc("vkCopyAccelerationStructureKHR")
	procvkCmdCopyAccelerationStructureToMemoryKHR = vulkan.NewProc("vkCmdCopyAccelerationStructureToMemoryKHR")
	procvkCopyAccelerationStructureToMemoryKHR = vulkan.NewProc("vkCopyAccelerationStructureToMemoryKHR")
	procvkCmdCopyMemoryToAccelerationStructureKHR = vulkan.NewProc("vkCmdCopyMemoryToAccelerationStructureKHR")
	procvkCopyMemoryToAccelerationStructureKHR = vulkan.NewProc("vkCopyMemoryToAccelerationStructureKHR")
	procvkCmdWriteAccelerationStructuresPropertiesKHR = vulkan.NewProc("vkCmdWriteAccelerationStructuresPropertiesKHR")
	procvkCmdWriteAccelerationStructuresPropertiesNV = vulkan.NewProc("vkCmdWriteAccelerationStructuresPropertiesNV")
	procvkCmdBuildAccelerationStructureNV = vulkan.NewProc("vkCmdBuildAccelerationStructureNV")
	procvkWriteAccelerationStructuresPropertiesKHR = vulkan.NewProc("vkWriteAccelerationStructuresPropertiesKHR")
	procvkCmdTraceRaysKHR = vulkan.NewProc("vkCmdTraceRaysKHR")
	procvkCmdTraceRaysNV = vulkan.NewProc("vkCmdTraceRaysNV")
	procvkGetRayTracingShaderGroupHandlesKHR = vulkan.NewProc("vkGetRayTracingShaderGroupHandlesKHR")
	procvkGetRayTracingShaderGroupHandlesNV = vulkan.NewProc("vkGetRayTracingShaderGroupHandlesNV")
	procvkGetRayTracingCaptureReplayShaderGroupHandlesKHR = vulkan.NewProc("vkGetRayTracingCaptureReplayShaderGroupHandlesKHR")
	procvkGetAccelerationStructureHandleNV = vulkan.NewProc("vkGetAccelerationStructureHandleNV")
	procvkCreateRayTracingPipelinesNV = vulkan.NewProc("vkCreateRayTracingPipelinesNV")
	procvkCreateRayTracingPipelinesKHR = vulkan.NewProc("vkCreateRayTracingPipelinesKHR")
	procvkGetPhysicalDeviceCooperativeMatrixPropertiesNV = vulkan.NewProc("vkGetPhysicalDeviceCooperativeMatrixPropertiesNV")
	procvkCmdTraceRaysIndirectKHR = vulkan.NewProc("vkCmdTraceRaysIndirectKHR")
	procvkGetDeviceAccelerationStructureCompatibilityKHR = vulkan.NewProc("vkGetDeviceAccelerationStructureCompatibilityKHR")
	procvkGetImageViewHandleNVX = vulkan.NewProc("vkGetImageViewHandleNVX")
	procvkGetImageViewAddressNVX = vulkan.NewProc("vkGetImageViewAddressNVX")
	procvkGetPhysicalDeviceSurfacePresentModes2EXT = vulkan.NewProc("vkGetPhysicalDeviceSurfacePresentModes2EXT")
	procvkGetDeviceGroupSurfacePresentModes2EXT = vulkan.NewProc("vkGetDeviceGroupSurfacePresentModes2EXT")
	procvkAcquireFullScreenExclusiveModeEXT = vulkan.NewProc("vkAcquireFullScreenExclusiveModeEXT")
	procvkReleaseFullScreenExclusiveModeEXT = vulkan.NewProc("vkReleaseFullScreenExclusiveModeEXT")
	procvkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR = vulkan.NewProc("vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR")
	procvkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR = vulkan.NewProc("vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR")
	procvkAcquireProfilingLockKHR = vulkan.NewProc("vkAcquireProfilingLockKHR")
	procvkReleaseProfilingLockKHR = vulkan.NewProc("vkReleaseProfilingLockKHR")
	procvkGetImageDrmFormatModifierPropertiesEXT = vulkan.NewProc("vkGetImageDrmFormatModifierPropertiesEXT")
	procvkGetBufferOpaqueCaptureAddress = vulkan.NewProc("vkGetBufferOpaqueCaptureAddress")
	procvkGetBufferOpaqueCaptureAddressKHR = vulkan.NewProc("vkGetBufferOpaqueCaptureAddressKHR")
	procvkGetBufferDeviceAddress = vulkan.NewProc("vkGetBufferDeviceAddress")
	procvkGetBufferDeviceAddressKHR = vulkan.NewProc("vkGetBufferDeviceAddressKHR")
	procvkGetBufferDeviceAddressEXT = vulkan.NewProc("vkGetBufferDeviceAddressEXT")
	procvkCreateHeadlessSurfaceEXT = vulkan.NewProc("vkCreateHeadlessSurfaceEXT")
	procvkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV = vulkan.NewProc("vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV")
	procvkInitializePerformanceApiINTEL = vulkan.NewProc("vkInitializePerformanceApiINTEL")
	procvkUninitializePerformanceApiINTEL = vulkan.NewProc("vkUninitializePerformanceApiINTEL")
	procvkCmdSetPerformanceMarkerINTEL = vulkan.NewProc("vkCmdSetPerformanceMarkerINTEL")
	procvkCmdSetPerformanceStreamMarkerINTEL = vulkan.NewProc("vkCmdSetPerformanceStreamMarkerINTEL")
	procvkCmdSetPerformanceOverrideINTEL = vulkan.NewProc("vkCmdSetPerformanceOverrideINTEL")
	procvkAcquirePerformanceConfigurationINTEL = vulkan.NewProc("vkAcquirePerformanceConfigurationINTEL")
	procvkReleasePerformanceConfigurationINTEL = vulkan.NewProc("vkReleasePerformanceConfigurationINTEL")
	procvkQueueSetPerformanceConfigurationINTEL = vulkan.NewProc("vkQueueSetPerformanceConfigurationINTEL")
	procvkGetPerformanceParameterINTEL = vulkan.NewProc("vkGetPerformanceParameterINTEL")
	procvkGetDeviceMemoryOpaqueCaptureAddress = vulkan.NewProc("vkGetDeviceMemoryOpaqueCaptureAddress")
	procvkGetDeviceMemoryOpaqueCaptureAddressKHR = vulkan.NewProc("vkGetDeviceMemoryOpaqueCaptureAddressKHR")
	procvkGetPipelineExecutablePropertiesKHR = vulkan.NewProc("vkGetPipelineExecutablePropertiesKHR")
	procvkGetPipelineExecutableStatisticsKHR = vulkan.NewProc("vkGetPipelineExecutableStatisticsKHR")
	procvkGetPipelineExecutableInternalRepresentationsKHR = vulkan.NewProc("vkGetPipelineExecutableInternalRepresentationsKHR")
	procvkCmdSetLineStippleEXT = vulkan.NewProc("vkCmdSetLineStippleEXT")
	procvkGetPhysicalDeviceToolPropertiesEXT = vulkan.NewProc("vkGetPhysicalDeviceToolPropertiesEXT")
	procvkCreateAccelerationStructureKHR = vulkan.NewProc("vkCreateAccelerationStructureKHR")
	procvkCmdBuildAccelerationStructureKHR = vulkan.NewProc("vkCmdBuildAccelerationStructureKHR")
	procvkCmdBuildAccelerationStructureIndirectKHR = vulkan.NewProc("vkCmdBuildAccelerationStructureIndirectKHR")
	procvkBuildAccelerationStructureKHR = vulkan.NewProc("vkBuildAccelerationStructureKHR")
	procvkGetAccelerationStructureDeviceAddressKHR = vulkan.NewProc("vkGetAccelerationStructureDeviceAddressKHR")
	procvkCreateDeferredOperationKHR = vulkan.NewProc("vkCreateDeferredOperationKHR")
	procvkDestroyDeferredOperationKHR = vulkan.NewProc("vkDestroyDeferredOperationKHR")
	procvkGetDeferredOperationMaxConcurrencyKHR = vulkan.NewProc("vkGetDeferredOperationMaxConcurrencyKHR")
	procvkGetDeferredOperationResultKHR = vulkan.NewProc("vkGetDeferredOperationResultKHR")
	procvkDeferredOperationJoinKHR = vulkan.NewProc("vkDeferredOperationJoinKHR")
	procvkCmdSetCullModeEXT = vulkan.NewProc("vkCmdSetCullModeEXT")
	procvkCmdSetFrontFaceEXT = vulkan.NewProc("vkCmdSetFrontFaceEXT")
	procvkCmdSetPrimitiveTopologyEXT = vulkan.NewProc("vkCmdSetPrimitiveTopologyEXT")
	procvkCmdSetViewportWithCountEXT = vulkan.NewProc("vkCmdSetViewportWithCountEXT")
	procvkCmdSetScissorWithCountEXT = vulkan.NewProc("vkCmdSetScissorWithCountEXT")
	procvkCmdBindVertexBuffers2EXT = vulkan.NewProc("vkCmdBindVertexBuffers2EXT")
	procvkCmdSetDepthTestEnableEXT = vulkan.NewProc("vkCmdSetDepthTestEnableEXT")
	procvkCmdSetDepthWriteEnableEXT = vulkan.NewProc("vkCmdSetDepthWriteEnableEXT")
	procvkCmdSetDepthCompareOpEXT = vulkan.NewProc("vkCmdSetDepthCompareOpEXT")
	procvkCmdSetDepthBoundsTestEnableEXT = vulkan.NewProc("vkCmdSetDepthBoundsTestEnableEXT")
	procvkCmdSetStencilTestEnableEXT = vulkan.NewProc("vkCmdSetStencilTestEnableEXT")
	procvkCmdSetStencilOpEXT = vulkan.NewProc("vkCmdSetStencilOpEXT")
	procvkCreatePrivateDataSlotEXT = vulkan.NewProc("vkCreatePrivateDataSlotEXT")
	procvkDestroyPrivateDataSlotEXT = vulkan.NewProc("vkDestroyPrivateDataSlotEXT")
	procvkSetPrivateDataEXT = vulkan.NewProc("vkSetPrivateDataEXT")
	procvkGetPrivateDataEXT = vulkan.NewProc("vkGetPrivateDataEXT")
	procvkCmdCopyBuffer2KHR = vulkan.NewProc("vkCmdCopyBuffer2KHR")
	procvkCmdCopyImage2KHR = vulkan.NewProc("vkCmdCopyImage2KHR")
	procvkCmdBlitImage2KHR = vulkan.NewProc("vkCmdBlitImage2KHR")
	procvkCmdCopyBufferToImage2KHR = vulkan.NewProc("vkCmdCopyBufferToImage2KHR")
	procvkCmdCopyImageToBuffer2KHR = vulkan.NewProc("vkCmdCopyImageToBuffer2KHR")
	procvkCmdResolveImage2KHR = vulkan.NewProc("vkCmdResolveImage2KHR")
	procvkCmdSetFragmentShadingRateKHR = vulkan.NewProc("vkCmdSetFragmentShadingRateKHR")
	procvkGetPhysicalDeviceFragmentShadingRatesKHR = vulkan.NewProc("vkGetPhysicalDeviceFragmentShadingRatesKHR")

)


//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INITIALIZATION_FAILED, VK_ERROR_LAYER_NOT_PRESENT, VK_ERROR_EXTENSION_NOT_PRESENT, VK_ERROR_INCOMPATIBLE_DRIVER
//On Success: VK_SUCCESS
func vkCreateInstance(pCreateInfo *VkInstanceCreateInfo, pAllocator *VkAllocationCallbacks, pInstance *VkInstance) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkCreateInstance.Addr(), 3, uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pInstance)))	
	ret0 = VkResult(r0)
	return
}


func vkDestroyInstance(instance VkInstance, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyInstance.Addr(), 2, uintptr(instance), uintptr(unsafe.Pointer(pAllocator)), 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INITIALIZATION_FAILED
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkEnumeratePhysicalDevices(instance VkInstance, pPhysicalDeviceCount *uint32, pPhysicalDevices *VkPhysicalDevice) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkEnumeratePhysicalDevices.Addr(), 3, uintptr(instance), uintptr(unsafe.Pointer(pPhysicalDeviceCount)), uintptr(unsafe.Pointer(pPhysicalDevices)))	
	ret0 = VkResult(r0)
	return
}


func vkGetDeviceProcAddr(device VkDevice, pName *byte) (ret0 PFN_vkVoidFunction) {
	r0, _, _ := syscall.Syscall(procvkGetDeviceProcAddr.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pName)), 0)	
	ret0 = PFN_vkVoidFunction(r0)
	return
}


func vkGetInstanceProcAddr(instance VkInstance, pName *byte) (ret0 PFN_vkVoidFunction) {
	r0, _, _ := syscall.Syscall(procvkGetInstanceProcAddr.Addr(), 2, uintptr(instance), uintptr(unsafe.Pointer(pName)), 0)	
	ret0 = PFN_vkVoidFunction(r0)
	return
}


func vkGetPhysicalDeviceProperties(physicalDevice VkPhysicalDevice, pProperties *VkPhysicalDeviceProperties) {
	syscall.Syscall(procvkGetPhysicalDeviceProperties.Addr(), 2, uintptr(physicalDevice), uintptr(unsafe.Pointer(pProperties)), 0)	
	return
}


func vkGetPhysicalDeviceQueueFamilyProperties(physicalDevice VkPhysicalDevice, pQueueFamilyPropertyCount *uint32, pQueueFamilyProperties *VkQueueFamilyProperties) {
	syscall.Syscall(procvkGetPhysicalDeviceQueueFamilyProperties.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pQueueFamilyPropertyCount)), uintptr(unsafe.Pointer(pQueueFamilyProperties)))	
	return
}


func vkGetPhysicalDeviceMemoryProperties(physicalDevice VkPhysicalDevice, pMemoryProperties *VkPhysicalDeviceMemoryProperties) {
	syscall.Syscall(procvkGetPhysicalDeviceMemoryProperties.Addr(), 2, uintptr(physicalDevice), uintptr(unsafe.Pointer(pMemoryProperties)), 0)	
	return
}


func vkGetPhysicalDeviceFeatures(physicalDevice VkPhysicalDevice, pFeatures *VkPhysicalDeviceFeatures) {
	syscall.Syscall(procvkGetPhysicalDeviceFeatures.Addr(), 2, uintptr(physicalDevice), uintptr(unsafe.Pointer(pFeatures)), 0)	
	return
}


func vkGetPhysicalDeviceFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, pFormatProperties *VkFormatProperties) {
	syscall.Syscall(procvkGetPhysicalDeviceFormatProperties.Addr(), 3, uintptr(physicalDevice), uintptr(format), uintptr(unsafe.Pointer(pFormatProperties)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_FORMAT_NOT_SUPPORTED
//On Success: VK_SUCCESS
func vkGetPhysicalDeviceImageFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, type0 VkImageType, tiling VkImageTiling, usage VkImageUsageFlags, flags VkImageCreateFlags, pImageFormatProperties *VkImageFormatProperties) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall9(procvkGetPhysicalDeviceImageFormatProperties.Addr(), 7, uintptr(physicalDevice), uintptr(format), uintptr(type0), uintptr(tiling), uintptr(usage), uintptr(flags), uintptr(unsafe.Pointer(pImageFormatProperties)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INITIALIZATION_FAILED, VK_ERROR_EXTENSION_NOT_PRESENT, VK_ERROR_FEATURE_NOT_PRESENT, VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_DEVICE_LOST
//On Success: VK_SUCCESS
func vkCreateDevice(physicalDevice VkPhysicalDevice, pCreateInfo *VkDeviceCreateInfo, pAllocator *VkAllocationCallbacks, pDevice *VkDevice) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateDevice.Addr(), 4, uintptr(physicalDevice), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pDevice)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyDevice(device VkDevice, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyDevice.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pAllocator)), 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkEnumerateInstanceVersion(pApiVersion *uint32) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkEnumerateInstanceVersion.Addr(), 1, uintptr(unsafe.Pointer(pApiVersion)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkEnumerateInstanceLayerProperties(pPropertyCount *uint32, pProperties *VkLayerProperties) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkEnumerateInstanceLayerProperties.Addr(), 2, uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_LAYER_NOT_PRESENT
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkEnumerateInstanceExtensionProperties(pLayerName *byte, pPropertyCount *uint32, pProperties *VkExtensionProperties) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkEnumerateInstanceExtensionProperties.Addr(), 3, uintptr(unsafe.Pointer(pLayerName)), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkEnumerateDeviceLayerProperties(physicalDevice VkPhysicalDevice, pPropertyCount *uint32, pProperties *VkLayerProperties) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkEnumerateDeviceLayerProperties.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_LAYER_NOT_PRESENT
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkEnumerateDeviceExtensionProperties(physicalDevice VkPhysicalDevice, pLayerName *byte, pPropertyCount *uint32, pProperties *VkExtensionProperties) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkEnumerateDeviceExtensionProperties.Addr(), 4, uintptr(physicalDevice), uintptr(unsafe.Pointer(pLayerName)), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkGetDeviceQueue(device VkDevice, queueFamilyIndex uint32, queueIndex uint32, pQueue *VkQueue) {
	syscall.Syscall6(procvkGetDeviceQueue.Addr(), 4, uintptr(device), uintptr(queueFamilyIndex), uintptr(queueIndex), uintptr(unsafe.Pointer(pQueue)), 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_DEVICE_LOST
//On Success: VK_SUCCESS
func vkQueueSubmit(queue VkQueue, submitCount uint32, pSubmits *VkSubmitInfo, fence VkFence) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkQueueSubmit.Addr(), 4, uintptr(queue), uintptr(submitCount), uintptr(unsafe.Pointer(pSubmits)), uintptr(fence), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_DEVICE_LOST
//On Success: VK_SUCCESS
func vkQueueWaitIdle(queue VkQueue) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkQueueWaitIdle.Addr(), 1, uintptr(queue), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_DEVICE_LOST
//On Success: VK_SUCCESS
func vkDeviceWaitIdle(device VkDevice) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkDeviceWaitIdle.Addr(), 1, uintptr(device), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INVALID_EXTERNAL_HANDLE, VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR
//On Success: VK_SUCCESS
func vkAllocateMemory(device VkDevice, pAllocateInfo *VkMemoryAllocateInfo, pAllocator *VkAllocationCallbacks, pMemory *VkDeviceMemory) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkAllocateMemory.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pAllocateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pMemory)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkFreeMemory(device VkDevice, memory VkDeviceMemory, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkFreeMemory.Addr(), 3, uintptr(device), uintptr(memory), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_MEMORY_MAP_FAILED
//On Success: VK_SUCCESS
func vkMapMemory(device VkDevice, memory VkDeviceMemory, offset VkDeviceSize, size VkDeviceSize, flags VkMemoryMapFlags, ppData *uintptr) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkMapMemory.Addr(), 6, uintptr(device), uintptr(memory), uintptr(offset), uintptr(size), uintptr(flags), uintptr(unsafe.Pointer(ppData)))	
	ret0 = VkResult(r0)
	return
}


func vkUnmapMemory(device VkDevice, memory VkDeviceMemory) {
	syscall.Syscall(procvkUnmapMemory.Addr(), 2, uintptr(device), uintptr(memory), 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkFlushMappedMemoryRanges(device VkDevice, memoryRangeCount uint32, pMemoryRanges *VkMappedMemoryRange) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkFlushMappedMemoryRanges.Addr(), 3, uintptr(device), uintptr(memoryRangeCount), uintptr(unsafe.Pointer(pMemoryRanges)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkInvalidateMappedMemoryRanges(device VkDevice, memoryRangeCount uint32, pMemoryRanges *VkMappedMemoryRange) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkInvalidateMappedMemoryRanges.Addr(), 3, uintptr(device), uintptr(memoryRangeCount), uintptr(unsafe.Pointer(pMemoryRanges)))	
	ret0 = VkResult(r0)
	return
}


func vkGetDeviceMemoryCommitment(device VkDevice, memory VkDeviceMemory, pCommittedMemoryInBytes *VkDeviceSize) {
	syscall.Syscall(procvkGetDeviceMemoryCommitment.Addr(), 3, uintptr(device), uintptr(memory), uintptr(unsafe.Pointer(pCommittedMemoryInBytes)))	
	return
}


func vkGetBufferMemoryRequirements(device VkDevice, buffer VkBuffer, pMemoryRequirements *VkMemoryRequirements) {
	syscall.Syscall(procvkGetBufferMemoryRequirements.Addr(), 3, uintptr(device), uintptr(buffer), uintptr(unsafe.Pointer(pMemoryRequirements)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR
//On Success: VK_SUCCESS
func vkBindBufferMemory(device VkDevice, buffer VkBuffer, memory VkDeviceMemory, memoryOffset VkDeviceSize) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkBindBufferMemory.Addr(), 4, uintptr(device), uintptr(buffer), uintptr(memory), uintptr(memoryOffset), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkGetImageMemoryRequirements(device VkDevice, image VkImage, pMemoryRequirements *VkMemoryRequirements) {
	syscall.Syscall(procvkGetImageMemoryRequirements.Addr(), 3, uintptr(device), uintptr(image), uintptr(unsafe.Pointer(pMemoryRequirements)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkBindImageMemory(device VkDevice, image VkImage, memory VkDeviceMemory, memoryOffset VkDeviceSize) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkBindImageMemory.Addr(), 4, uintptr(device), uintptr(image), uintptr(memory), uintptr(memoryOffset), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkGetImageSparseMemoryRequirements(device VkDevice, image VkImage, pSparseMemoryRequirementCount *uint32, pSparseMemoryRequirements *VkSparseImageMemoryRequirements) {
	syscall.Syscall6(procvkGetImageSparseMemoryRequirements.Addr(), 4, uintptr(device), uintptr(image), uintptr(unsafe.Pointer(pSparseMemoryRequirementCount)), uintptr(unsafe.Pointer(pSparseMemoryRequirements)), 0, 0)	
	return
}


func vkGetPhysicalDeviceSparseImageFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, type0 VkImageType, samples VkSampleCountFlagBits, usage VkImageUsageFlags, tiling VkImageTiling, pPropertyCount *uint32, pProperties *VkSparseImageFormatProperties) {
	syscall.Syscall9(procvkGetPhysicalDeviceSparseImageFormatProperties.Addr(), 8, uintptr(physicalDevice), uintptr(format), uintptr(type0), uintptr(samples), uintptr(usage), uintptr(tiling), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)), 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_DEVICE_LOST
//On Success: VK_SUCCESS
func vkQueueBindSparse(queue VkQueue, bindInfoCount uint32, pBindInfo *VkBindSparseInfo, fence VkFence) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkQueueBindSparse.Addr(), 4, uintptr(queue), uintptr(bindInfoCount), uintptr(unsafe.Pointer(pBindInfo)), uintptr(fence), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateFence(device VkDevice, pCreateInfo *VkFenceCreateInfo, pAllocator *VkAllocationCallbacks, pFence *VkFence) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateFence.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pFence)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyFence(device VkDevice, fence VkFence, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyFence.Addr(), 3, uintptr(device), uintptr(fence), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkResetFences(device VkDevice, fenceCount uint32, pFences *VkFence) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkResetFences.Addr(), 3, uintptr(device), uintptr(fenceCount), uintptr(unsafe.Pointer(pFences)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_DEVICE_LOST
//On Success: VK_SUCCESS, VK_NOT_READY
func vkGetFenceStatus(device VkDevice, fence VkFence) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetFenceStatus.Addr(), 2, uintptr(device), uintptr(fence), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_DEVICE_LOST
//On Success: VK_SUCCESS, VK_TIMEOUT
func vkWaitForFences(device VkDevice, fenceCount uint32, pFences *VkFence, waitAll VkBool32, timeout uint64) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkWaitForFences.Addr(), 5, uintptr(device), uintptr(fenceCount), uintptr(unsafe.Pointer(pFences)), uintptr(waitAll), uintptr(timeout), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateSemaphore(device VkDevice, pCreateInfo *VkSemaphoreCreateInfo, pAllocator *VkAllocationCallbacks, pSemaphore *VkSemaphore) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateSemaphore.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSemaphore)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroySemaphore(device VkDevice, semaphore VkSemaphore, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroySemaphore.Addr(), 3, uintptr(device), uintptr(semaphore), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateEvent(device VkDevice, pCreateInfo *VkEventCreateInfo, pAllocator *VkAllocationCallbacks, pEvent *VkEvent) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateEvent.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pEvent)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyEvent(device VkDevice, event VkEvent, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyEvent.Addr(), 3, uintptr(device), uintptr(event), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_DEVICE_LOST
//On Success: VK_EVENT_SET, VK_EVENT_RESET
func vkGetEventStatus(device VkDevice, event VkEvent) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetEventStatus.Addr(), 2, uintptr(device), uintptr(event), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkSetEvent(device VkDevice, event VkEvent) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkSetEvent.Addr(), 2, uintptr(device), uintptr(event), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkResetEvent(device VkDevice, event VkEvent) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkResetEvent.Addr(), 2, uintptr(device), uintptr(event), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateQueryPool(device VkDevice, pCreateInfo *VkQueryPoolCreateInfo, pAllocator *VkAllocationCallbacks, pQueryPool *VkQueryPool) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateQueryPool.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pQueryPool)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyQueryPool(device VkDevice, queryPool VkQueryPool, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyQueryPool.Addr(), 3, uintptr(device), uintptr(queryPool), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_DEVICE_LOST
//On Success: VK_SUCCESS, VK_NOT_READY
func vkGetQueryPoolResults(device VkDevice, queryPool VkQueryPool, firstQuery uint32, queryCount uint32, dataSize uint, pData uintptr, stride VkDeviceSize, flags VkQueryResultFlags) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall9(procvkGetQueryPoolResults.Addr(), 8, uintptr(device), uintptr(queryPool), uintptr(firstQuery), uintptr(queryCount), uintptr(dataSize), uintptr(pData), uintptr(stride), uintptr(flags), 0)	
	ret0 = VkResult(r0)
	return
}


func vkResetQueryPool(device VkDevice, queryPool VkQueryPool, firstQuery uint32, queryCount uint32) {
	syscall.Syscall6(procvkResetQueryPool.Addr(), 4, uintptr(device), uintptr(queryPool), uintptr(firstQuery), uintptr(queryCount), 0, 0)	
	return
}


func vkResetQueryPoolEXT() {
	syscall.Syscall(procvkResetQueryPoolEXT.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR
//On Success: VK_SUCCESS
func vkCreateBuffer(device VkDevice, pCreateInfo *VkBufferCreateInfo, pAllocator *VkAllocationCallbacks, pBuffer *VkBuffer) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateBuffer.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pBuffer)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyBuffer(device VkDevice, buffer VkBuffer, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyBuffer.Addr(), 3, uintptr(device), uintptr(buffer), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateBufferView(device VkDevice, pCreateInfo *VkBufferViewCreateInfo, pAllocator *VkAllocationCallbacks, pView *VkBufferView) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateBufferView.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pView)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyBufferView(device VkDevice, bufferView VkBufferView, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyBufferView.Addr(), 3, uintptr(device), uintptr(bufferView), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateImage(device VkDevice, pCreateInfo *VkImageCreateInfo, pAllocator *VkAllocationCallbacks, pImage *VkImage) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateImage.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pImage)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyImage(device VkDevice, image VkImage, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyImage.Addr(), 3, uintptr(device), uintptr(image), uintptr(unsafe.Pointer(pAllocator)))	
	return
}


func vkGetImageSubresourceLayout(device VkDevice, image VkImage, pSubresource *VkImageSubresource, pLayout *VkSubresourceLayout) {
	syscall.Syscall6(procvkGetImageSubresourceLayout.Addr(), 4, uintptr(device), uintptr(image), uintptr(unsafe.Pointer(pSubresource)), uintptr(unsafe.Pointer(pLayout)), 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateImageView(device VkDevice, pCreateInfo *VkImageViewCreateInfo, pAllocator *VkAllocationCallbacks, pView *VkImageView) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateImageView.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pView)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyImageView(device VkDevice, imageView VkImageView, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyImageView.Addr(), 3, uintptr(device), uintptr(imageView), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INVALID_SHADER_NV
//On Success: VK_SUCCESS
func vkCreateShaderModule(device VkDevice, pCreateInfo *VkShaderModuleCreateInfo, pAllocator *VkAllocationCallbacks, pShaderModule *VkShaderModule) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateShaderModule.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pShaderModule)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyShaderModule(device VkDevice, shaderModule VkShaderModule, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyShaderModule.Addr(), 3, uintptr(device), uintptr(shaderModule), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreatePipelineCache(device VkDevice, pCreateInfo *VkPipelineCacheCreateInfo, pAllocator *VkAllocationCallbacks, pPipelineCache *VkPipelineCache) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreatePipelineCache.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pPipelineCache)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyPipelineCache(device VkDevice, pipelineCache VkPipelineCache, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyPipelineCache.Addr(), 3, uintptr(device), uintptr(pipelineCache), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPipelineCacheData(device VkDevice, pipelineCache VkPipelineCache, pDataSize *uint, pData uintptr) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetPipelineCacheData.Addr(), 4, uintptr(device), uintptr(pipelineCache), uintptr(unsafe.Pointer(pDataSize)), uintptr(pData), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkMergePipelineCaches(device VkDevice, dstCache VkPipelineCache, srcCacheCount uint32, pSrcCaches *VkPipelineCache) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkMergePipelineCaches.Addr(), 4, uintptr(device), uintptr(dstCache), uintptr(srcCacheCount), uintptr(unsafe.Pointer(pSrcCaches)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INVALID_SHADER_NV
//On Success: VK_SUCCESS, VK_PIPELINE_COMPILE_REQUIRED_EXT
func vkCreateGraphicsPipelines(device VkDevice, pipelineCache VkPipelineCache, createInfoCount uint32, pCreateInfos *VkGraphicsPipelineCreateInfo, pAllocator *VkAllocationCallbacks, pPipelines *VkPipeline) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateGraphicsPipelines.Addr(), 6, uintptr(device), uintptr(pipelineCache), uintptr(createInfoCount), uintptr(unsafe.Pointer(pCreateInfos)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pPipelines)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INVALID_SHADER_NV
//On Success: VK_SUCCESS, VK_PIPELINE_COMPILE_REQUIRED_EXT
func vkCreateComputePipelines(device VkDevice, pipelineCache VkPipelineCache, createInfoCount uint32, pCreateInfos *VkComputePipelineCreateInfo, pAllocator *VkAllocationCallbacks, pPipelines *VkPipeline) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateComputePipelines.Addr(), 6, uintptr(device), uintptr(pipelineCache), uintptr(createInfoCount), uintptr(unsafe.Pointer(pCreateInfos)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pPipelines)))	
	ret0 = VkResult(r0)
	return
}


func vkDestroyPipeline(device VkDevice, pipeline VkPipeline, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyPipeline.Addr(), 3, uintptr(device), uintptr(pipeline), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreatePipelineLayout(device VkDevice, pCreateInfo *VkPipelineLayoutCreateInfo, pAllocator *VkAllocationCallbacks, pPipelineLayout *VkPipelineLayout) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreatePipelineLayout.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pPipelineLayout)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyPipelineLayout(device VkDevice, pipelineLayout VkPipelineLayout, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyPipelineLayout.Addr(), 3, uintptr(device), uintptr(pipelineLayout), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateSampler(device VkDevice, pCreateInfo *VkSamplerCreateInfo, pAllocator *VkAllocationCallbacks, pSampler *VkSampler) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateSampler.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSampler)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroySampler(device VkDevice, sampler VkSampler, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroySampler.Addr(), 3, uintptr(device), uintptr(sampler), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateDescriptorSetLayout(device VkDevice, pCreateInfo *VkDescriptorSetLayoutCreateInfo, pAllocator *VkAllocationCallbacks, pSetLayout *VkDescriptorSetLayout) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateDescriptorSetLayout.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSetLayout)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyDescriptorSetLayout(device VkDevice, descriptorSetLayout VkDescriptorSetLayout, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyDescriptorSetLayout.Addr(), 3, uintptr(device), uintptr(descriptorSetLayout), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_FRAGMENTATION_EXT
//On Success: VK_SUCCESS
func vkCreateDescriptorPool(device VkDevice, pCreateInfo *VkDescriptorPoolCreateInfo, pAllocator *VkAllocationCallbacks, pDescriptorPool *VkDescriptorPool) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateDescriptorPool.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pDescriptorPool)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyDescriptorPool(device VkDevice, descriptorPool VkDescriptorPool, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyDescriptorPool.Addr(), 3, uintptr(device), uintptr(descriptorPool), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//On Success: VK_SUCCESS
func vkResetDescriptorPool(device VkDevice, descriptorPool VkDescriptorPool, flags VkDescriptorPoolResetFlags) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkResetDescriptorPool.Addr(), 3, uintptr(device), uintptr(descriptorPool), uintptr(flags))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_FRAGMENTED_POOL, VK_ERROR_OUT_OF_POOL_MEMORY
//On Success: VK_SUCCESS
func vkAllocateDescriptorSets(device VkDevice, pAllocateInfo *VkDescriptorSetAllocateInfo, pDescriptorSets *VkDescriptorSet) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkAllocateDescriptorSets.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pAllocateInfo)), uintptr(unsafe.Pointer(pDescriptorSets)))	
	ret0 = VkResult(r0)
	return
}

//On Success: VK_SUCCESS
func vkFreeDescriptorSets(device VkDevice, descriptorPool VkDescriptorPool, descriptorSetCount uint32, pDescriptorSets *VkDescriptorSet) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkFreeDescriptorSets.Addr(), 4, uintptr(device), uintptr(descriptorPool), uintptr(descriptorSetCount), uintptr(unsafe.Pointer(pDescriptorSets)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkUpdateDescriptorSets(device VkDevice, descriptorWriteCount uint32, pDescriptorWrites *VkWriteDescriptorSet, descriptorCopyCount uint32, pDescriptorCopies *VkCopyDescriptorSet) {
	syscall.Syscall6(procvkUpdateDescriptorSets.Addr(), 5, uintptr(device), uintptr(descriptorWriteCount), uintptr(unsafe.Pointer(pDescriptorWrites)), uintptr(descriptorCopyCount), uintptr(unsafe.Pointer(pDescriptorCopies)), 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateFramebuffer(device VkDevice, pCreateInfo *VkFramebufferCreateInfo, pAllocator *VkAllocationCallbacks, pFramebuffer *VkFramebuffer) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateFramebuffer.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pFramebuffer)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyFramebuffer(device VkDevice, framebuffer VkFramebuffer, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyFramebuffer.Addr(), 3, uintptr(device), uintptr(framebuffer), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateRenderPass(device VkDevice, pCreateInfo *VkRenderPassCreateInfo, pAllocator *VkAllocationCallbacks, pRenderPass *VkRenderPass) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateRenderPass.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pRenderPass)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyRenderPass(device VkDevice, renderPass VkRenderPass, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyRenderPass.Addr(), 3, uintptr(device), uintptr(renderPass), uintptr(unsafe.Pointer(pAllocator)))	
	return
}


func vkGetRenderAreaGranularity(device VkDevice, renderPass VkRenderPass, pGranularity *VkExtent2D) {
	syscall.Syscall(procvkGetRenderAreaGranularity.Addr(), 3, uintptr(device), uintptr(renderPass), uintptr(unsafe.Pointer(pGranularity)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateCommandPool(device VkDevice, pCreateInfo *VkCommandPoolCreateInfo, pAllocator *VkAllocationCallbacks, pCommandPool *VkCommandPool) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateCommandPool.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pCommandPool)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyCommandPool(device VkDevice, commandPool VkCommandPool, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyCommandPool.Addr(), 3, uintptr(device), uintptr(commandPool), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkResetCommandPool(device VkDevice, commandPool VkCommandPool, flags VkCommandPoolResetFlags) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkResetCommandPool.Addr(), 3, uintptr(device), uintptr(commandPool), uintptr(flags))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkAllocateCommandBuffers(device VkDevice, pAllocateInfo *VkCommandBufferAllocateInfo, pCommandBuffers *VkCommandBuffer) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkAllocateCommandBuffers.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pAllocateInfo)), uintptr(unsafe.Pointer(pCommandBuffers)))	
	ret0 = VkResult(r0)
	return
}


func vkFreeCommandBuffers(device VkDevice, commandPool VkCommandPool, commandBufferCount uint32, pCommandBuffers *VkCommandBuffer) {
	syscall.Syscall6(procvkFreeCommandBuffers.Addr(), 4, uintptr(device), uintptr(commandPool), uintptr(commandBufferCount), uintptr(unsafe.Pointer(pCommandBuffers)), 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkBeginCommandBuffer(commandBuffer VkCommandBuffer, pBeginInfo *VkCommandBufferBeginInfo) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkBeginCommandBuffer.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pBeginInfo)), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkEndCommandBuffer(commandBuffer VkCommandBuffer) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkEndCommandBuffer.Addr(), 1, uintptr(commandBuffer), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkResetCommandBuffer(commandBuffer VkCommandBuffer, flags VkCommandBufferResetFlags) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkResetCommandBuffer.Addr(), 2, uintptr(commandBuffer), uintptr(flags), 0)	
	ret0 = VkResult(r0)
	return
}


func vkCmdBindPipeline(commandBuffer VkCommandBuffer, pipelineBindPoint VkPipelineBindPoint, pipeline VkPipeline) {
	syscall.Syscall(procvkCmdBindPipeline.Addr(), 3, uintptr(commandBuffer), uintptr(pipelineBindPoint), uintptr(pipeline))	
	return
}


func vkCmdSetViewport(commandBuffer VkCommandBuffer, firstViewport uint32, viewportCount uint32, pViewports *VkViewport) {
	syscall.Syscall6(procvkCmdSetViewport.Addr(), 4, uintptr(commandBuffer), uintptr(firstViewport), uintptr(viewportCount), uintptr(unsafe.Pointer(pViewports)), 0, 0)	
	return
}


func vkCmdSetScissor(commandBuffer VkCommandBuffer, firstScissor uint32, scissorCount uint32, pScissors *VkRect2D) {
	syscall.Syscall6(procvkCmdSetScissor.Addr(), 4, uintptr(commandBuffer), uintptr(firstScissor), uintptr(scissorCount), uintptr(unsafe.Pointer(pScissors)), 0, 0)	
	return
}


func vkCmdSetLineWidth(commandBuffer VkCommandBuffer, lineWidth float32) {
	syscall.Syscall(procvkCmdSetLineWidth.Addr(), 2, uintptr(commandBuffer), uintptr(lineWidth), 0)	
	return
}


func vkCmdSetDepthBias(commandBuffer VkCommandBuffer, depthBiasConstantFactor float32, depthBiasClamp float32, depthBiasSlopeFactor float32) {
	syscall.Syscall6(procvkCmdSetDepthBias.Addr(), 4, uintptr(commandBuffer), uintptr(depthBiasConstantFactor), uintptr(depthBiasClamp), uintptr(depthBiasSlopeFactor), 0, 0)	
	return
}


func vkCmdSetBlendConstants(commandBuffer VkCommandBuffer, blendConstants4 float32, blendConstants3 float32, blendConstants2 float32, blendConstants1 float32) {
	syscall.Syscall6(procvkCmdSetBlendConstants.Addr(), 5, uintptr(commandBuffer), uintptr(blendConstants4), uintptr(blendConstants3), uintptr(blendConstants2), uintptr(blendConstants1), 0)	
	return
}


func vkCmdSetDepthBounds(commandBuffer VkCommandBuffer, minDepthBounds float32, maxDepthBounds float32) {
	syscall.Syscall(procvkCmdSetDepthBounds.Addr(), 3, uintptr(commandBuffer), uintptr(minDepthBounds), uintptr(maxDepthBounds))	
	return
}


func vkCmdSetStencilCompareMask(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, compareMask uint32) {
	syscall.Syscall(procvkCmdSetStencilCompareMask.Addr(), 3, uintptr(commandBuffer), uintptr(faceMask), uintptr(compareMask))	
	return
}


func vkCmdSetStencilWriteMask(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, writeMask uint32) {
	syscall.Syscall(procvkCmdSetStencilWriteMask.Addr(), 3, uintptr(commandBuffer), uintptr(faceMask), uintptr(writeMask))	
	return
}


func vkCmdSetStencilReference(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, reference uint32) {
	syscall.Syscall(procvkCmdSetStencilReference.Addr(), 3, uintptr(commandBuffer), uintptr(faceMask), uintptr(reference))	
	return
}


func vkCmdBindDescriptorSets(commandBuffer VkCommandBuffer, pipelineBindPoint VkPipelineBindPoint, layout VkPipelineLayout, firstSet uint32, descriptorSetCount uint32, pDescriptorSets *VkDescriptorSet, dynamicOffsetCount uint32, pDynamicOffsets *uint32) {
	syscall.Syscall9(procvkCmdBindDescriptorSets.Addr(), 8, uintptr(commandBuffer), uintptr(pipelineBindPoint), uintptr(layout), uintptr(firstSet), uintptr(descriptorSetCount), uintptr(unsafe.Pointer(pDescriptorSets)), uintptr(dynamicOffsetCount), uintptr(unsafe.Pointer(pDynamicOffsets)), 0)	
	return
}


func vkCmdBindIndexBuffer(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, indexType VkIndexType) {
	syscall.Syscall6(procvkCmdBindIndexBuffer.Addr(), 4, uintptr(commandBuffer), uintptr(buffer), uintptr(offset), uintptr(indexType), 0, 0)	
	return
}


func vkCmdBindVertexBuffers(commandBuffer VkCommandBuffer, firstBinding uint32, bindingCount uint32, pBuffers *VkBuffer, pOffsets *VkDeviceSize) {
	syscall.Syscall6(procvkCmdBindVertexBuffers.Addr(), 5, uintptr(commandBuffer), uintptr(firstBinding), uintptr(bindingCount), uintptr(unsafe.Pointer(pBuffers)), uintptr(unsafe.Pointer(pOffsets)), 0)	
	return
}


func vkCmdDraw(commandBuffer VkCommandBuffer, vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	syscall.Syscall6(procvkCmdDraw.Addr(), 5, uintptr(commandBuffer), uintptr(vertexCount), uintptr(instanceCount), uintptr(firstVertex), uintptr(firstInstance), 0)	
	return
}


func vkCmdDrawIndexed(commandBuffer VkCommandBuffer, indexCount uint32, instanceCount uint32, firstIndex uint32, vertexOffset int32, firstInstance uint32) {
	syscall.Syscall6(procvkCmdDrawIndexed.Addr(), 6, uintptr(commandBuffer), uintptr(indexCount), uintptr(instanceCount), uintptr(firstIndex), uintptr(vertexOffset), uintptr(firstInstance))	
	return
}


func vkCmdDrawIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, drawCount uint32, stride uint32) {
	syscall.Syscall6(procvkCmdDrawIndirect.Addr(), 5, uintptr(commandBuffer), uintptr(buffer), uintptr(offset), uintptr(drawCount), uintptr(stride), 0)	
	return
}


func vkCmdDrawIndexedIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, drawCount uint32, stride uint32) {
	syscall.Syscall6(procvkCmdDrawIndexedIndirect.Addr(), 5, uintptr(commandBuffer), uintptr(buffer), uintptr(offset), uintptr(drawCount), uintptr(stride), 0)	
	return
}


func vkCmdDispatch(commandBuffer VkCommandBuffer, groupCountX uint32, groupCountY uint32, groupCountZ uint32) {
	syscall.Syscall6(procvkCmdDispatch.Addr(), 4, uintptr(commandBuffer), uintptr(groupCountX), uintptr(groupCountY), uintptr(groupCountZ), 0, 0)	
	return
}


func vkCmdDispatchIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize) {
	syscall.Syscall(procvkCmdDispatchIndirect.Addr(), 3, uintptr(commandBuffer), uintptr(buffer), uintptr(offset))	
	return
}


func vkCmdCopyBuffer(commandBuffer VkCommandBuffer, srcBuffer VkBuffer, dstBuffer VkBuffer, regionCount uint32, pRegions *VkBufferCopy) {
	syscall.Syscall6(procvkCmdCopyBuffer.Addr(), 5, uintptr(commandBuffer), uintptr(srcBuffer), uintptr(dstBuffer), uintptr(regionCount), uintptr(unsafe.Pointer(pRegions)), 0)	
	return
}


func vkCmdCopyImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount uint32, pRegions *VkImageCopy) {
	syscall.Syscall9(procvkCmdCopyImage.Addr(), 7, uintptr(commandBuffer), uintptr(srcImage), uintptr(srcImageLayout), uintptr(dstImage), uintptr(dstImageLayout), uintptr(regionCount), uintptr(unsafe.Pointer(pRegions)), 0, 0)	
	return
}


func vkCmdBlitImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount uint32, pRegions *VkImageBlit, filter VkFilter) {
	syscall.Syscall9(procvkCmdBlitImage.Addr(), 8, uintptr(commandBuffer), uintptr(srcImage), uintptr(srcImageLayout), uintptr(dstImage), uintptr(dstImageLayout), uintptr(regionCount), uintptr(unsafe.Pointer(pRegions)), uintptr(filter), 0)	
	return
}


func vkCmdCopyBufferToImage(commandBuffer VkCommandBuffer, srcBuffer VkBuffer, dstImage VkImage, dstImageLayout VkImageLayout, regionCount uint32, pRegions *VkBufferImageCopy) {
	syscall.Syscall6(procvkCmdCopyBufferToImage.Addr(), 6, uintptr(commandBuffer), uintptr(srcBuffer), uintptr(dstImage), uintptr(dstImageLayout), uintptr(regionCount), uintptr(unsafe.Pointer(pRegions)))	
	return
}


func vkCmdCopyImageToBuffer(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstBuffer VkBuffer, regionCount uint32, pRegions *VkBufferImageCopy) {
	syscall.Syscall6(procvkCmdCopyImageToBuffer.Addr(), 6, uintptr(commandBuffer), uintptr(srcImage), uintptr(srcImageLayout), uintptr(dstBuffer), uintptr(regionCount), uintptr(unsafe.Pointer(pRegions)))	
	return
}


func vkCmdUpdateBuffer(commandBuffer VkCommandBuffer, dstBuffer VkBuffer, dstOffset VkDeviceSize, dataSize VkDeviceSize, pData uintptr) {
	syscall.Syscall6(procvkCmdUpdateBuffer.Addr(), 5, uintptr(commandBuffer), uintptr(dstBuffer), uintptr(dstOffset), uintptr(dataSize), uintptr(pData), 0)	
	return
}

//transfer support is only available when VK_KHR_maintenance1 is enabled, as documented in valid usage language in the specification
func vkCmdFillBuffer(commandBuffer VkCommandBuffer, dstBuffer VkBuffer, dstOffset VkDeviceSize, size VkDeviceSize, data uint32) {
	syscall.Syscall6(procvkCmdFillBuffer.Addr(), 5, uintptr(commandBuffer), uintptr(dstBuffer), uintptr(dstOffset), uintptr(size), uintptr(data), 0)	
	return
}


func vkCmdClearColorImage(commandBuffer VkCommandBuffer, image VkImage, imageLayout VkImageLayout, pColor *VkClearColorValue, rangeCount uint32, pRanges *VkImageSubresourceRange) {
	syscall.Syscall6(procvkCmdClearColorImage.Addr(), 6, uintptr(commandBuffer), uintptr(image), uintptr(imageLayout), uintptr(unsafe.Pointer(pColor)), uintptr(rangeCount), uintptr(unsafe.Pointer(pRanges)))	
	return
}


func vkCmdClearDepthStencilImage(commandBuffer VkCommandBuffer, image VkImage, imageLayout VkImageLayout, pDepthStencil *VkClearDepthStencilValue, rangeCount uint32, pRanges *VkImageSubresourceRange) {
	syscall.Syscall6(procvkCmdClearDepthStencilImage.Addr(), 6, uintptr(commandBuffer), uintptr(image), uintptr(imageLayout), uintptr(unsafe.Pointer(pDepthStencil)), uintptr(rangeCount), uintptr(unsafe.Pointer(pRanges)))	
	return
}


func vkCmdClearAttachments(commandBuffer VkCommandBuffer, attachmentCount uint32, pAttachments *VkClearAttachment, rectCount uint32, pRects *VkClearRect) {
	syscall.Syscall6(procvkCmdClearAttachments.Addr(), 5, uintptr(commandBuffer), uintptr(attachmentCount), uintptr(unsafe.Pointer(pAttachments)), uintptr(rectCount), uintptr(unsafe.Pointer(pRects)), 0)	
	return
}


func vkCmdResolveImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount uint32, pRegions *VkImageResolve) {
	syscall.Syscall9(procvkCmdResolveImage.Addr(), 7, uintptr(commandBuffer), uintptr(srcImage), uintptr(srcImageLayout), uintptr(dstImage), uintptr(dstImageLayout), uintptr(regionCount), uintptr(unsafe.Pointer(pRegions)), 0, 0)	
	return
}


func vkCmdSetEvent(commandBuffer VkCommandBuffer, event VkEvent, stageMask VkPipelineStageFlags) {
	syscall.Syscall(procvkCmdSetEvent.Addr(), 3, uintptr(commandBuffer), uintptr(event), uintptr(stageMask))	
	return
}


func vkCmdResetEvent(commandBuffer VkCommandBuffer, event VkEvent, stageMask VkPipelineStageFlags) {
	syscall.Syscall(procvkCmdResetEvent.Addr(), 3, uintptr(commandBuffer), uintptr(event), uintptr(stageMask))	
	return
}


func vkCmdWaitEvents(commandBuffer VkCommandBuffer, eventCount uint32, pEvents *VkEvent, srcStageMask VkPipelineStageFlags, dstStageMask VkPipelineStageFlags, memoryBarrierCount uint32, pMemoryBarriers *VkMemoryBarrier, bufferMemoryBarrierCount uint32, pBufferMemoryBarriers *VkBufferMemoryBarrier, imageMemoryBarrierCount uint32, pImageMemoryBarriers *VkImageMemoryBarrier) {
	syscall.Syscall12(procvkCmdWaitEvents.Addr(), 11, uintptr(commandBuffer), uintptr(eventCount), uintptr(unsafe.Pointer(pEvents)), uintptr(srcStageMask), uintptr(dstStageMask), uintptr(memoryBarrierCount), uintptr(unsafe.Pointer(pMemoryBarriers)), uintptr(bufferMemoryBarrierCount), uintptr(unsafe.Pointer(pBufferMemoryBarriers)), uintptr(imageMemoryBarrierCount), uintptr(unsafe.Pointer(pImageMemoryBarriers)), 0)	
	return
}


func vkCmdPipelineBarrier(commandBuffer VkCommandBuffer, srcStageMask VkPipelineStageFlags, dstStageMask VkPipelineStageFlags, dependencyFlags VkDependencyFlags, memoryBarrierCount uint32, pMemoryBarriers *VkMemoryBarrier, bufferMemoryBarrierCount uint32, pBufferMemoryBarriers *VkBufferMemoryBarrier, imageMemoryBarrierCount uint32, pImageMemoryBarriers *VkImageMemoryBarrier) {
	syscall.Syscall12(procvkCmdPipelineBarrier.Addr(), 10, uintptr(commandBuffer), uintptr(srcStageMask), uintptr(dstStageMask), uintptr(dependencyFlags), uintptr(memoryBarrierCount), uintptr(unsafe.Pointer(pMemoryBarriers)), uintptr(bufferMemoryBarrierCount), uintptr(unsafe.Pointer(pBufferMemoryBarriers)), uintptr(imageMemoryBarrierCount), uintptr(unsafe.Pointer(pImageMemoryBarriers)), 0, 0)	
	return
}


func vkCmdBeginQuery(commandBuffer VkCommandBuffer, queryPool VkQueryPool, query uint32, flags VkQueryControlFlags) {
	syscall.Syscall6(procvkCmdBeginQuery.Addr(), 4, uintptr(commandBuffer), uintptr(queryPool), uintptr(query), uintptr(flags), 0, 0)	
	return
}


func vkCmdEndQuery(commandBuffer VkCommandBuffer, queryPool VkQueryPool, query uint32) {
	syscall.Syscall(procvkCmdEndQuery.Addr(), 3, uintptr(commandBuffer), uintptr(queryPool), uintptr(query))	
	return
}


func vkCmdBeginConditionalRenderingEXT(commandBuffer VkCommandBuffer, pConditionalRenderingBegin *VkConditionalRenderingBeginInfoEXT) {
	syscall.Syscall(procvkCmdBeginConditionalRenderingEXT.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pConditionalRenderingBegin)), 0)	
	return
}


func vkCmdEndConditionalRenderingEXT(commandBuffer VkCommandBuffer) {
	syscall.Syscall(procvkCmdEndConditionalRenderingEXT.Addr(), 1, uintptr(commandBuffer), 0, 0)	
	return
}


func vkCmdResetQueryPool(commandBuffer VkCommandBuffer, queryPool VkQueryPool, firstQuery uint32, queryCount uint32) {
	syscall.Syscall6(procvkCmdResetQueryPool.Addr(), 4, uintptr(commandBuffer), uintptr(queryPool), uintptr(firstQuery), uintptr(queryCount), 0, 0)	
	return
}


func vkCmdWriteTimestamp(commandBuffer VkCommandBuffer, pipelineStage VkPipelineStageFlagBits, queryPool VkQueryPool, query uint32) {
	syscall.Syscall6(procvkCmdWriteTimestamp.Addr(), 4, uintptr(commandBuffer), uintptr(pipelineStage), uintptr(queryPool), uintptr(query), 0, 0)	
	return
}


func vkCmdCopyQueryPoolResults(commandBuffer VkCommandBuffer, queryPool VkQueryPool, firstQuery uint32, queryCount uint32, dstBuffer VkBuffer, dstOffset VkDeviceSize, stride VkDeviceSize, flags VkQueryResultFlags) {
	syscall.Syscall9(procvkCmdCopyQueryPoolResults.Addr(), 8, uintptr(commandBuffer), uintptr(queryPool), uintptr(firstQuery), uintptr(queryCount), uintptr(dstBuffer), uintptr(dstOffset), uintptr(stride), uintptr(flags), 0)	
	return
}


func vkCmdPushConstants(commandBuffer VkCommandBuffer, layout VkPipelineLayout, stageFlags VkShaderStageFlags, offset uint32, size uint32, pValues uintptr) {
	syscall.Syscall6(procvkCmdPushConstants.Addr(), 6, uintptr(commandBuffer), uintptr(layout), uintptr(stageFlags), uintptr(offset), uintptr(size), uintptr(pValues))	
	return
}


func vkCmdBeginRenderPass(commandBuffer VkCommandBuffer, pRenderPassBegin *VkRenderPassBeginInfo, contents VkSubpassContents) {
	syscall.Syscall(procvkCmdBeginRenderPass.Addr(), 3, uintptr(commandBuffer), uintptr(unsafe.Pointer(pRenderPassBegin)), uintptr(contents))	
	return
}


func vkCmdNextSubpass(commandBuffer VkCommandBuffer, contents VkSubpassContents) {
	syscall.Syscall(procvkCmdNextSubpass.Addr(), 2, uintptr(commandBuffer), uintptr(contents), 0)	
	return
}


func vkCmdEndRenderPass(commandBuffer VkCommandBuffer) {
	syscall.Syscall(procvkCmdEndRenderPass.Addr(), 1, uintptr(commandBuffer), 0, 0)	
	return
}


func vkCmdExecuteCommands(commandBuffer VkCommandBuffer, commandBufferCount uint32, pCommandBuffers *VkCommandBuffer) {
	syscall.Syscall(procvkCmdExecuteCommands.Addr(), 3, uintptr(commandBuffer), uintptr(commandBufferCount), uintptr(unsafe.Pointer(pCommandBuffers)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
//On Success: VK_SUCCESS
func vkCreateAndroidSurfaceKHR(instance VkInstance, pCreateInfo *VkAndroidSurfaceCreateInfoKHR, pAllocator *VkAllocationCallbacks, pSurface *VkSurfaceKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateAndroidSurfaceKHR.Addr(), 4, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPhysicalDeviceDisplayPropertiesKHR(physicalDevice VkPhysicalDevice, pPropertyCount *uint32, pProperties *VkDisplayPropertiesKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetPhysicalDeviceDisplayPropertiesKHR.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPhysicalDeviceDisplayPlanePropertiesKHR(physicalDevice VkPhysicalDevice, pPropertyCount *uint32, pProperties *VkDisplayPlanePropertiesKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetPhysicalDeviceDisplayPlanePropertiesKHR.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetDisplayPlaneSupportedDisplaysKHR(physicalDevice VkPhysicalDevice, planeIndex uint32, pDisplayCount *uint32, pDisplays *VkDisplayKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetDisplayPlaneSupportedDisplaysKHR.Addr(), 4, uintptr(physicalDevice), uintptr(planeIndex), uintptr(unsafe.Pointer(pDisplayCount)), uintptr(unsafe.Pointer(pDisplays)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetDisplayModePropertiesKHR(physicalDevice VkPhysicalDevice, display VkDisplayKHR, pPropertyCount *uint32, pProperties *VkDisplayModePropertiesKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetDisplayModePropertiesKHR.Addr(), 4, uintptr(physicalDevice), uintptr(display), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INITIALIZATION_FAILED
//On Success: VK_SUCCESS
func vkCreateDisplayModeKHR(physicalDevice VkPhysicalDevice, display VkDisplayKHR, pCreateInfo *VkDisplayModeCreateInfoKHR, pAllocator *VkAllocationCallbacks, pMode *VkDisplayModeKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateDisplayModeKHR.Addr(), 5, uintptr(physicalDevice), uintptr(display), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pMode)), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkGetDisplayPlaneCapabilitiesKHR(physicalDevice VkPhysicalDevice, mode VkDisplayModeKHR, planeIndex uint32, pCapabilities *VkDisplayPlaneCapabilitiesKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetDisplayPlaneCapabilitiesKHR.Addr(), 4, uintptr(physicalDevice), uintptr(mode), uintptr(planeIndex), uintptr(unsafe.Pointer(pCapabilities)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateDisplayPlaneSurfaceKHR(instance VkInstance, pCreateInfo *VkDisplaySurfaceCreateInfoKHR, pAllocator *VkAllocationCallbacks, pSurface *VkSurfaceKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateDisplayPlaneSurfaceKHR.Addr(), 4, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INCOMPATIBLE_DISPLAY_KHR, VK_ERROR_DEVICE_LOST, VK_ERROR_SURFACE_LOST_KHR
//On Success: VK_SUCCESS
func vkCreateSharedSwapchainsKHR(device VkDevice, swapchainCount uint32, pCreateInfos *VkSwapchainCreateInfoKHR, pAllocator *VkAllocationCallbacks, pSwapchains *VkSwapchainKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateSharedSwapchainsKHR.Addr(), 5, uintptr(device), uintptr(swapchainCount), uintptr(unsafe.Pointer(pCreateInfos)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSwapchains)), 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroySurfaceKHR(instance VkInstance, surface VkSurfaceKHR, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroySurfaceKHR.Addr(), 3, uintptr(instance), uintptr(surface), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_SURFACE_LOST_KHR
//On Success: VK_SUCCESS
func vkGetPhysicalDeviceSurfaceSupportKHR(physicalDevice VkPhysicalDevice, queueFamilyIndex uint32, surface VkSurfaceKHR, pSupported *VkBool32) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetPhysicalDeviceSurfaceSupportKHR.Addr(), 4, uintptr(physicalDevice), uintptr(queueFamilyIndex), uintptr(surface), uintptr(unsafe.Pointer(pSupported)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_SURFACE_LOST_KHR
//On Success: VK_SUCCESS
func vkGetPhysicalDeviceSurfaceCapabilitiesKHR(physicalDevice VkPhysicalDevice, surface VkSurfaceKHR, pSurfaceCapabilities *VkSurfaceCapabilitiesKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetPhysicalDeviceSurfaceCapabilitiesKHR.Addr(), 3, uintptr(physicalDevice), uintptr(surface), uintptr(unsafe.Pointer(pSurfaceCapabilities)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_SURFACE_LOST_KHR
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPhysicalDeviceSurfaceFormatsKHR(physicalDevice VkPhysicalDevice, surface VkSurfaceKHR, pSurfaceFormatCount *uint32, pSurfaceFormats *VkSurfaceFormatKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetPhysicalDeviceSurfaceFormatsKHR.Addr(), 4, uintptr(physicalDevice), uintptr(surface), uintptr(unsafe.Pointer(pSurfaceFormatCount)), uintptr(unsafe.Pointer(pSurfaceFormats)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_SURFACE_LOST_KHR
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPhysicalDeviceSurfacePresentModesKHR(physicalDevice VkPhysicalDevice, surface VkSurfaceKHR, pPresentModeCount *uint32, pPresentModes *VkPresentModeKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetPhysicalDeviceSurfacePresentModesKHR.Addr(), 4, uintptr(physicalDevice), uintptr(surface), uintptr(unsafe.Pointer(pPresentModeCount)), uintptr(unsafe.Pointer(pPresentModes)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_DEVICE_LOST, VK_ERROR_SURFACE_LOST_KHR, VK_ERROR_NATIVE_WINDOW_IN_USE_KHR, VK_ERROR_INITIALIZATION_FAILED
//On Success: VK_SUCCESS
func vkCreateSwapchainKHR(device VkDevice, pCreateInfo *VkSwapchainCreateInfoKHR, pAllocator *VkAllocationCallbacks, pSwapchain *VkSwapchainKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateSwapchainKHR.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSwapchain)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroySwapchainKHR(device VkDevice, swapchain VkSwapchainKHR, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroySwapchainKHR.Addr(), 3, uintptr(device), uintptr(swapchain), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetSwapchainImagesKHR(device VkDevice, swapchain VkSwapchainKHR, pSwapchainImageCount *uint32, pSwapchainImages *VkImage) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetSwapchainImagesKHR.Addr(), 4, uintptr(device), uintptr(swapchain), uintptr(unsafe.Pointer(pSwapchainImageCount)), uintptr(unsafe.Pointer(pSwapchainImages)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_DEVICE_LOST, VK_ERROR_OUT_OF_DATE_KHR, VK_ERROR_SURFACE_LOST_KHR, VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT
//On Success: VK_SUCCESS, VK_TIMEOUT, VK_NOT_READY, VK_SUBOPTIMAL_KHR
func vkAcquireNextImageKHR(device VkDevice, swapchain VkSwapchainKHR, timeout uint64, semaphore VkSemaphore, fence VkFence, pImageIndex *uint32) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkAcquireNextImageKHR.Addr(), 6, uintptr(device), uintptr(swapchain), uintptr(timeout), uintptr(semaphore), uintptr(fence), uintptr(unsafe.Pointer(pImageIndex)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_DEVICE_LOST, VK_ERROR_OUT_OF_DATE_KHR, VK_ERROR_SURFACE_LOST_KHR, VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT
//On Success: VK_SUCCESS, VK_SUBOPTIMAL_KHR
func vkQueuePresentKHR(queue VkQueue, pPresentInfo *VkPresentInfoKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkQueuePresentKHR.Addr(), 2, uintptr(queue), uintptr(unsafe.Pointer(pPresentInfo)), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
//On Success: VK_SUCCESS
func vkCreateViSurfaceNN(instance VkInstance, pCreateInfo *VkViSurfaceCreateInfoNN, pAllocator *VkAllocationCallbacks, pSurface *VkSurfaceKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateViSurfaceNN.Addr(), 4, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateWaylandSurfaceKHR(instance VkInstance, pCreateInfo *VkWaylandSurfaceCreateInfoKHR, pAllocator *VkAllocationCallbacks, pSurface *VkSurfaceKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateWaylandSurfaceKHR.Addr(), 4, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkGetPhysicalDeviceWaylandPresentationSupportKHR(physicalDevice VkPhysicalDevice, queueFamilyIndex uint32, display *wl_display) (ret0 VkBool32) {
	r0, _, _ := syscall.Syscall(procvkGetPhysicalDeviceWaylandPresentationSupportKHR.Addr(), 3, uintptr(physicalDevice), uintptr(queueFamilyIndex), uintptr(unsafe.Pointer(display)))	
	ret0 = VkBool32(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateWin32SurfaceKHR(instance VkInstance, pCreateInfo *VkWin32SurfaceCreateInfoKHR, pAllocator *VkAllocationCallbacks, pSurface *VkSurfaceKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateWin32SurfaceKHR.Addr(), 4, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkGetPhysicalDeviceWin32PresentationSupportKHR(physicalDevice VkPhysicalDevice, queueFamilyIndex uint32) (ret0 VkBool32) {
	r0, _, _ := syscall.Syscall(procvkGetPhysicalDeviceWin32PresentationSupportKHR.Addr(), 2, uintptr(physicalDevice), uintptr(queueFamilyIndex), 0)	
	ret0 = VkBool32(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateXlibSurfaceKHR(instance VkInstance, pCreateInfo *VkXlibSurfaceCreateInfoKHR, pAllocator *VkAllocationCallbacks, pSurface *VkSurfaceKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateXlibSurfaceKHR.Addr(), 4, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkGetPhysicalDeviceXlibPresentationSupportKHR(physicalDevice VkPhysicalDevice, queueFamilyIndex uint32, dpy *Display, visualID VisualID) (ret0 VkBool32) {
	r0, _, _ := syscall.Syscall6(procvkGetPhysicalDeviceXlibPresentationSupportKHR.Addr(), 4, uintptr(physicalDevice), uintptr(queueFamilyIndex), uintptr(unsafe.Pointer(dpy)), uintptr(visualID), 0, 0)	
	ret0 = VkBool32(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateXcbSurfaceKHR(instance VkInstance, pCreateInfo *VkXcbSurfaceCreateInfoKHR, pAllocator *VkAllocationCallbacks, pSurface *VkSurfaceKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateXcbSurfaceKHR.Addr(), 4, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkGetPhysicalDeviceXcbPresentationSupportKHR(physicalDevice VkPhysicalDevice, queueFamilyIndex uint32, connection *xcb_connection_t, visual_id xcb_visualid_t) (ret0 VkBool32) {
	r0, _, _ := syscall.Syscall6(procvkGetPhysicalDeviceXcbPresentationSupportKHR.Addr(), 4, uintptr(physicalDevice), uintptr(queueFamilyIndex), uintptr(unsafe.Pointer(connection)), uintptr(visual_id), 0, 0)	
	ret0 = VkBool32(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateDirectFBSurfaceEXT(instance VkInstance, pCreateInfo *VkDirectFBSurfaceCreateInfoEXT, pAllocator *VkAllocationCallbacks, pSurface *VkSurfaceKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateDirectFBSurfaceEXT.Addr(), 4, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkGetPhysicalDeviceDirectFBPresentationSupportEXT(physicalDevice VkPhysicalDevice, queueFamilyIndex uint32, dfb *IDirectFB) (ret0 VkBool32) {
	r0, _, _ := syscall.Syscall(procvkGetPhysicalDeviceDirectFBPresentationSupportEXT.Addr(), 3, uintptr(physicalDevice), uintptr(queueFamilyIndex), uintptr(unsafe.Pointer(dfb)))	
	ret0 = VkBool32(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateImagePipeSurfaceFUCHSIA(instance VkInstance, pCreateInfo *VkImagePipeSurfaceCreateInfoFUCHSIA, pAllocator *VkAllocationCallbacks, pSurface *VkSurfaceKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateImagePipeSurfaceFUCHSIA.Addr(), 4, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
//On Success: VK_SUCCESS
func vkCreateStreamDescriptorSurfaceGGP(instance VkInstance, pCreateInfo *VkStreamDescriptorSurfaceCreateInfoGGP, pAllocator *VkAllocationCallbacks, pSurface *VkSurfaceKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateStreamDescriptorSurfaceGGP.Addr(), 4, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkCreateDebugReportCallbackEXT(instance VkInstance, pCreateInfo *VkDebugReportCallbackCreateInfoEXT, pAllocator *VkAllocationCallbacks, pCallback *VkDebugReportCallbackEXT) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateDebugReportCallbackEXT.Addr(), 4, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pCallback)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyDebugReportCallbackEXT(instance VkInstance, callback VkDebugReportCallbackEXT, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyDebugReportCallbackEXT.Addr(), 3, uintptr(instance), uintptr(callback), uintptr(unsafe.Pointer(pAllocator)))	
	return
}


func vkDebugReportMessageEXT(instance VkInstance, flags VkDebugReportFlagsEXT, objectType VkDebugReportObjectTypeEXT, object uint64, location uint, messageCode int32, pLayerPrefix *byte, pMessage *byte) {
	syscall.Syscall9(procvkDebugReportMessageEXT.Addr(), 8, uintptr(instance), uintptr(flags), uintptr(objectType), uintptr(object), uintptr(location), uintptr(messageCode), uintptr(unsafe.Pointer(pLayerPrefix)), uintptr(unsafe.Pointer(pMessage)), 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkDebugMarkerSetObjectNameEXT(device VkDevice, pNameInfo *VkDebugMarkerObjectNameInfoEXT) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkDebugMarkerSetObjectNameEXT.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pNameInfo)), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkDebugMarkerSetObjectTagEXT(device VkDevice, pTagInfo *VkDebugMarkerObjectTagInfoEXT) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkDebugMarkerSetObjectTagEXT.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pTagInfo)), 0)	
	ret0 = VkResult(r0)
	return
}


func vkCmdDebugMarkerBeginEXT(commandBuffer VkCommandBuffer, pMarkerInfo *VkDebugMarkerMarkerInfoEXT) {
	syscall.Syscall(procvkCmdDebugMarkerBeginEXT.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pMarkerInfo)), 0)	
	return
}


func vkCmdDebugMarkerEndEXT(commandBuffer VkCommandBuffer) {
	syscall.Syscall(procvkCmdDebugMarkerEndEXT.Addr(), 1, uintptr(commandBuffer), 0, 0)	
	return
}


func vkCmdDebugMarkerInsertEXT(commandBuffer VkCommandBuffer, pMarkerInfo *VkDebugMarkerMarkerInfoEXT) {
	syscall.Syscall(procvkCmdDebugMarkerInsertEXT.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pMarkerInfo)), 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_FORMAT_NOT_SUPPORTED
//On Success: VK_SUCCESS
func vkGetPhysicalDeviceExternalImageFormatPropertiesNV(physicalDevice VkPhysicalDevice, format VkFormat, type0 VkImageType, tiling VkImageTiling, usage VkImageUsageFlags, flags VkImageCreateFlags, externalHandleType VkExternalMemoryHandleTypeFlagsNV, pExternalImageFormatProperties *VkExternalImageFormatPropertiesNV) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall9(procvkGetPhysicalDeviceExternalImageFormatPropertiesNV.Addr(), 8, uintptr(physicalDevice), uintptr(format), uintptr(type0), uintptr(tiling), uintptr(usage), uintptr(flags), uintptr(externalHandleType), uintptr(unsafe.Pointer(pExternalImageFormatProperties)), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkGetMemoryWin32HandleNV(device VkDevice, memory VkDeviceMemory, handleType VkExternalMemoryHandleTypeFlagsNV, pHandle *HANDLE) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetMemoryWin32HandleNV.Addr(), 4, uintptr(device), uintptr(memory), uintptr(handleType), uintptr(unsafe.Pointer(pHandle)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkCmdExecuteGeneratedCommandsNV(commandBuffer VkCommandBuffer, isPreprocessed VkBool32, pGeneratedCommandsInfo *VkGeneratedCommandsInfoNV) {
	syscall.Syscall(procvkCmdExecuteGeneratedCommandsNV.Addr(), 3, uintptr(commandBuffer), uintptr(isPreprocessed), uintptr(unsafe.Pointer(pGeneratedCommandsInfo)))	
	return
}


func vkCmdPreprocessGeneratedCommandsNV(commandBuffer VkCommandBuffer, pGeneratedCommandsInfo *VkGeneratedCommandsInfoNV) {
	syscall.Syscall(procvkCmdPreprocessGeneratedCommandsNV.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pGeneratedCommandsInfo)), 0)	
	return
}


func vkCmdBindPipelineShaderGroupNV(commandBuffer VkCommandBuffer, pipelineBindPoint VkPipelineBindPoint, pipeline VkPipeline, groupIndex uint32) {
	syscall.Syscall6(procvkCmdBindPipelineShaderGroupNV.Addr(), 4, uintptr(commandBuffer), uintptr(pipelineBindPoint), uintptr(pipeline), uintptr(groupIndex), 0, 0)	
	return
}


func vkGetGeneratedCommandsMemoryRequirementsNV(device VkDevice, pInfo *VkGeneratedCommandsMemoryRequirementsInfoNV, pMemoryRequirements *VkMemoryRequirements2) {
	syscall.Syscall(procvkGetGeneratedCommandsMemoryRequirementsNV.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pInfo)), uintptr(unsafe.Pointer(pMemoryRequirements)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateIndirectCommandsLayoutNV(device VkDevice, pCreateInfo *VkIndirectCommandsLayoutCreateInfoNV, pAllocator *VkAllocationCallbacks, pIndirectCommandsLayout *VkIndirectCommandsLayoutNV) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateIndirectCommandsLayoutNV.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pIndirectCommandsLayout)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyIndirectCommandsLayoutNV(device VkDevice, indirectCommandsLayout VkIndirectCommandsLayoutNV, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyIndirectCommandsLayoutNV.Addr(), 3, uintptr(device), uintptr(indirectCommandsLayout), uintptr(unsafe.Pointer(pAllocator)))	
	return
}


func vkGetPhysicalDeviceFeatures2(physicalDevice VkPhysicalDevice, pFeatures *VkPhysicalDeviceFeatures2) {
	syscall.Syscall(procvkGetPhysicalDeviceFeatures2.Addr(), 2, uintptr(physicalDevice), uintptr(unsafe.Pointer(pFeatures)), 0)	
	return
}


func vkGetPhysicalDeviceFeatures2KHR() {
	syscall.Syscall(procvkGetPhysicalDeviceFeatures2KHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkGetPhysicalDeviceProperties2(physicalDevice VkPhysicalDevice, pProperties *VkPhysicalDeviceProperties2) {
	syscall.Syscall(procvkGetPhysicalDeviceProperties2.Addr(), 2, uintptr(physicalDevice), uintptr(unsafe.Pointer(pProperties)), 0)	
	return
}


func vkGetPhysicalDeviceProperties2KHR() {
	syscall.Syscall(procvkGetPhysicalDeviceProperties2KHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkGetPhysicalDeviceFormatProperties2(physicalDevice VkPhysicalDevice, format VkFormat, pFormatProperties *VkFormatProperties2) {
	syscall.Syscall(procvkGetPhysicalDeviceFormatProperties2.Addr(), 3, uintptr(physicalDevice), uintptr(format), uintptr(unsafe.Pointer(pFormatProperties)))	
	return
}


func vkGetPhysicalDeviceFormatProperties2KHR() {
	syscall.Syscall(procvkGetPhysicalDeviceFormatProperties2KHR.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_FORMAT_NOT_SUPPORTED
//On Success: VK_SUCCESS
func vkGetPhysicalDeviceImageFormatProperties2(physicalDevice VkPhysicalDevice, pImageFormatInfo *VkPhysicalDeviceImageFormatInfo2, pImageFormatProperties *VkImageFormatProperties2) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetPhysicalDeviceImageFormatProperties2.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pImageFormatInfo)), uintptr(unsafe.Pointer(pImageFormatProperties)))	
	ret0 = VkResult(r0)
	return
}


func vkGetPhysicalDeviceImageFormatProperties2KHR() {
	syscall.Syscall(procvkGetPhysicalDeviceImageFormatProperties2KHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkGetPhysicalDeviceQueueFamilyProperties2(physicalDevice VkPhysicalDevice, pQueueFamilyPropertyCount *uint32, pQueueFamilyProperties *VkQueueFamilyProperties2) {
	syscall.Syscall(procvkGetPhysicalDeviceQueueFamilyProperties2.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pQueueFamilyPropertyCount)), uintptr(unsafe.Pointer(pQueueFamilyProperties)))	
	return
}


func vkGetPhysicalDeviceQueueFamilyProperties2KHR() {
	syscall.Syscall(procvkGetPhysicalDeviceQueueFamilyProperties2KHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkGetPhysicalDeviceMemoryProperties2(physicalDevice VkPhysicalDevice, pMemoryProperties *VkPhysicalDeviceMemoryProperties2) {
	syscall.Syscall(procvkGetPhysicalDeviceMemoryProperties2.Addr(), 2, uintptr(physicalDevice), uintptr(unsafe.Pointer(pMemoryProperties)), 0)	
	return
}


func vkGetPhysicalDeviceMemoryProperties2KHR() {
	syscall.Syscall(procvkGetPhysicalDeviceMemoryProperties2KHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkGetPhysicalDeviceSparseImageFormatProperties2(physicalDevice VkPhysicalDevice, pFormatInfo *VkPhysicalDeviceSparseImageFormatInfo2, pPropertyCount *uint32, pProperties *VkSparseImageFormatProperties2) {
	syscall.Syscall6(procvkGetPhysicalDeviceSparseImageFormatProperties2.Addr(), 4, uintptr(physicalDevice), uintptr(unsafe.Pointer(pFormatInfo)), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)), 0, 0)	
	return
}


func vkGetPhysicalDeviceSparseImageFormatProperties2KHR() {
	syscall.Syscall(procvkGetPhysicalDeviceSparseImageFormatProperties2KHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkCmdPushDescriptorSetKHR(commandBuffer VkCommandBuffer, pipelineBindPoint VkPipelineBindPoint, layout VkPipelineLayout, set uint32, descriptorWriteCount uint32, pDescriptorWrites *VkWriteDescriptorSet) {
	syscall.Syscall6(procvkCmdPushDescriptorSetKHR.Addr(), 6, uintptr(commandBuffer), uintptr(pipelineBindPoint), uintptr(layout), uintptr(set), uintptr(descriptorWriteCount), uintptr(unsafe.Pointer(pDescriptorWrites)))	
	return
}


func vkTrimCommandPool(device VkDevice, commandPool VkCommandPool, flags VkCommandPoolTrimFlags) {
	syscall.Syscall(procvkTrimCommandPool.Addr(), 3, uintptr(device), uintptr(commandPool), uintptr(flags))	
	return
}


func vkTrimCommandPoolKHR() {
	syscall.Syscall(procvkTrimCommandPoolKHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkGetPhysicalDeviceExternalBufferProperties(physicalDevice VkPhysicalDevice, pExternalBufferInfo *VkPhysicalDeviceExternalBufferInfo, pExternalBufferProperties *VkExternalBufferProperties) {
	syscall.Syscall(procvkGetPhysicalDeviceExternalBufferProperties.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pExternalBufferInfo)), uintptr(unsafe.Pointer(pExternalBufferProperties)))	
	return
}


func vkGetPhysicalDeviceExternalBufferPropertiesKHR() {
	syscall.Syscall(procvkGetPhysicalDeviceExternalBufferPropertiesKHR.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkGetMemoryWin32HandleKHR(device VkDevice, pGetWin32HandleInfo *VkMemoryGetWin32HandleInfoKHR, pHandle *HANDLE) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetMemoryWin32HandleKHR.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pGetWin32HandleInfo)), uintptr(unsafe.Pointer(pHandle)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_INVALID_EXTERNAL_HANDLE
//On Success: VK_SUCCESS
func vkGetMemoryWin32HandlePropertiesKHR(device VkDevice, handleType VkExternalMemoryHandleTypeFlagBits, handle HANDLE, pMemoryWin32HandleProperties *VkMemoryWin32HandlePropertiesKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetMemoryWin32HandlePropertiesKHR.Addr(), 4, uintptr(device), uintptr(handleType), uintptr(handle), uintptr(unsafe.Pointer(pMemoryWin32HandleProperties)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkGetMemoryFdKHR(device VkDevice, pGetFdInfo *VkMemoryGetFdInfoKHR, pFd *int32) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetMemoryFdKHR.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pGetFdInfo)), uintptr(unsafe.Pointer(pFd)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_INVALID_EXTERNAL_HANDLE
//On Success: VK_SUCCESS
func vkGetMemoryFdPropertiesKHR(device VkDevice, handleType VkExternalMemoryHandleTypeFlagBits, fd int32, pMemoryFdProperties *VkMemoryFdPropertiesKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetMemoryFdPropertiesKHR.Addr(), 4, uintptr(device), uintptr(handleType), uintptr(fd), uintptr(unsafe.Pointer(pMemoryFdProperties)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkGetPhysicalDeviceExternalSemaphoreProperties(physicalDevice VkPhysicalDevice, pExternalSemaphoreInfo *VkPhysicalDeviceExternalSemaphoreInfo, pExternalSemaphoreProperties *VkExternalSemaphoreProperties) {
	syscall.Syscall(procvkGetPhysicalDeviceExternalSemaphoreProperties.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pExternalSemaphoreInfo)), uintptr(unsafe.Pointer(pExternalSemaphoreProperties)))	
	return
}


func vkGetPhysicalDeviceExternalSemaphorePropertiesKHR() {
	syscall.Syscall(procvkGetPhysicalDeviceExternalSemaphorePropertiesKHR.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkGetSemaphoreWin32HandleKHR(device VkDevice, pGetWin32HandleInfo *VkSemaphoreGetWin32HandleInfoKHR, pHandle *HANDLE) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetSemaphoreWin32HandleKHR.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pGetWin32HandleInfo)), uintptr(unsafe.Pointer(pHandle)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_INVALID_EXTERNAL_HANDLE
//On Success: VK_SUCCESS
func vkImportSemaphoreWin32HandleKHR(device VkDevice, pImportSemaphoreWin32HandleInfo *VkImportSemaphoreWin32HandleInfoKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkImportSemaphoreWin32HandleKHR.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pImportSemaphoreWin32HandleInfo)), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkGetSemaphoreFdKHR(device VkDevice, pGetFdInfo *VkSemaphoreGetFdInfoKHR, pFd *int32) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetSemaphoreFdKHR.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pGetFdInfo)), uintptr(unsafe.Pointer(pFd)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_INVALID_EXTERNAL_HANDLE
//On Success: VK_SUCCESS
func vkImportSemaphoreFdKHR(device VkDevice, pImportSemaphoreFdInfo *VkImportSemaphoreFdInfoKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkImportSemaphoreFdKHR.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pImportSemaphoreFdInfo)), 0)	
	ret0 = VkResult(r0)
	return
}


func vkGetPhysicalDeviceExternalFenceProperties(physicalDevice VkPhysicalDevice, pExternalFenceInfo *VkPhysicalDeviceExternalFenceInfo, pExternalFenceProperties *VkExternalFenceProperties) {
	syscall.Syscall(procvkGetPhysicalDeviceExternalFenceProperties.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pExternalFenceInfo)), uintptr(unsafe.Pointer(pExternalFenceProperties)))	
	return
}


func vkGetPhysicalDeviceExternalFencePropertiesKHR() {
	syscall.Syscall(procvkGetPhysicalDeviceExternalFencePropertiesKHR.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkGetFenceWin32HandleKHR(device VkDevice, pGetWin32HandleInfo *VkFenceGetWin32HandleInfoKHR, pHandle *HANDLE) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetFenceWin32HandleKHR.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pGetWin32HandleInfo)), uintptr(unsafe.Pointer(pHandle)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_INVALID_EXTERNAL_HANDLE
//On Success: VK_SUCCESS
func vkImportFenceWin32HandleKHR(device VkDevice, pImportFenceWin32HandleInfo *VkImportFenceWin32HandleInfoKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkImportFenceWin32HandleKHR.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pImportFenceWin32HandleInfo)), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkGetFenceFdKHR(device VkDevice, pGetFdInfo *VkFenceGetFdInfoKHR, pFd *int32) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetFenceFdKHR.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pGetFdInfo)), uintptr(unsafe.Pointer(pFd)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_INVALID_EXTERNAL_HANDLE
//On Success: VK_SUCCESS
func vkImportFenceFdKHR(device VkDevice, pImportFenceFdInfo *VkImportFenceFdInfoKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkImportFenceFdKHR.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pImportFenceFdInfo)), 0)	
	ret0 = VkResult(r0)
	return
}

//On Success: VK_SUCCESS
func vkReleaseDisplayEXT(physicalDevice VkPhysicalDevice, display VkDisplayKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkReleaseDisplayEXT.Addr(), 2, uintptr(physicalDevice), uintptr(display), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_INITIALIZATION_FAILED
//On Success: VK_SUCCESS
func vkAcquireXlibDisplayEXT(physicalDevice VkPhysicalDevice, dpy *Display, display VkDisplayKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkAcquireXlibDisplayEXT.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(dpy)), uintptr(display))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkGetRandROutputDisplayEXT(physicalDevice VkPhysicalDevice, dpy *Display, rrOutput RROutput, pDisplay *VkDisplayKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetRandROutputDisplayEXT.Addr(), 4, uintptr(physicalDevice), uintptr(unsafe.Pointer(dpy)), uintptr(rrOutput), uintptr(unsafe.Pointer(pDisplay)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkDisplayPowerControlEXT(device VkDevice, display VkDisplayKHR, pDisplayPowerInfo *VkDisplayPowerInfoEXT) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkDisplayPowerControlEXT.Addr(), 3, uintptr(device), uintptr(display), uintptr(unsafe.Pointer(pDisplayPowerInfo)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkRegisterDeviceEventEXT(device VkDevice, pDeviceEventInfo *VkDeviceEventInfoEXT, pAllocator *VkAllocationCallbacks, pFence *VkFence) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkRegisterDeviceEventEXT.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pDeviceEventInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pFence)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkRegisterDisplayEventEXT(device VkDevice, display VkDisplayKHR, pDisplayEventInfo *VkDisplayEventInfoEXT, pAllocator *VkAllocationCallbacks, pFence *VkFence) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkRegisterDisplayEventEXT.Addr(), 5, uintptr(device), uintptr(display), uintptr(unsafe.Pointer(pDisplayEventInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pFence)), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_DEVICE_LOST, VK_ERROR_OUT_OF_DATE_KHR
//On Success: VK_SUCCESS
func vkGetSwapchainCounterEXT(device VkDevice, swapchain VkSwapchainKHR, counter VkSurfaceCounterFlagBitsEXT, pCounterValue *uint64) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetSwapchainCounterEXT.Addr(), 4, uintptr(device), uintptr(swapchain), uintptr(counter), uintptr(unsafe.Pointer(pCounterValue)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_SURFACE_LOST_KHR
//On Success: VK_SUCCESS
func vkGetPhysicalDeviceSurfaceCapabilities2EXT(physicalDevice VkPhysicalDevice, surface VkSurfaceKHR, pSurfaceCapabilities *VkSurfaceCapabilities2EXT) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetPhysicalDeviceSurfaceCapabilities2EXT.Addr(), 3, uintptr(physicalDevice), uintptr(surface), uintptr(unsafe.Pointer(pSurfaceCapabilities)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INITIALIZATION_FAILED
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkEnumeratePhysicalDeviceGroups(instance VkInstance, pPhysicalDeviceGroupCount *uint32, pPhysicalDeviceGroupProperties *VkPhysicalDeviceGroupProperties) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkEnumeratePhysicalDeviceGroups.Addr(), 3, uintptr(instance), uintptr(unsafe.Pointer(pPhysicalDeviceGroupCount)), uintptr(unsafe.Pointer(pPhysicalDeviceGroupProperties)))	
	ret0 = VkResult(r0)
	return
}


func vkEnumeratePhysicalDeviceGroupsKHR() {
	syscall.Syscall(procvkEnumeratePhysicalDeviceGroupsKHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkGetDeviceGroupPeerMemoryFeatures(device VkDevice, heapIndex uint32, localDeviceIndex uint32, remoteDeviceIndex uint32, pPeerMemoryFeatures *VkPeerMemoryFeatureFlags) {
	syscall.Syscall6(procvkGetDeviceGroupPeerMemoryFeatures.Addr(), 5, uintptr(device), uintptr(heapIndex), uintptr(localDeviceIndex), uintptr(remoteDeviceIndex), uintptr(unsafe.Pointer(pPeerMemoryFeatures)), 0)	
	return
}


func vkGetDeviceGroupPeerMemoryFeaturesKHR() {
	syscall.Syscall(procvkGetDeviceGroupPeerMemoryFeaturesKHR.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR
//On Success: VK_SUCCESS
func vkBindBufferMemory2(device VkDevice, bindInfoCount uint32, pBindInfos *VkBindBufferMemoryInfo) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkBindBufferMemory2.Addr(), 3, uintptr(device), uintptr(bindInfoCount), uintptr(unsafe.Pointer(pBindInfos)))	
	ret0 = VkResult(r0)
	return
}


func vkBindBufferMemory2KHR() {
	syscall.Syscall(procvkBindBufferMemory2KHR.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkBindImageMemory2(device VkDevice, bindInfoCount uint32, pBindInfos *VkBindImageMemoryInfo) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkBindImageMemory2.Addr(), 3, uintptr(device), uintptr(bindInfoCount), uintptr(unsafe.Pointer(pBindInfos)))	
	ret0 = VkResult(r0)
	return
}


func vkBindImageMemory2KHR() {
	syscall.Syscall(procvkBindImageMemory2KHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkCmdSetDeviceMask(commandBuffer VkCommandBuffer, deviceMask uint32) {
	syscall.Syscall(procvkCmdSetDeviceMask.Addr(), 2, uintptr(commandBuffer), uintptr(deviceMask), 0)	
	return
}


func vkCmdSetDeviceMaskKHR() {
	syscall.Syscall(procvkCmdSetDeviceMaskKHR.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkGetDeviceGroupPresentCapabilitiesKHR(device VkDevice, pDeviceGroupPresentCapabilities *VkDeviceGroupPresentCapabilitiesKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetDeviceGroupPresentCapabilitiesKHR.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pDeviceGroupPresentCapabilities)), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_SURFACE_LOST_KHR
//On Success: VK_SUCCESS
func vkGetDeviceGroupSurfacePresentModesKHR(device VkDevice, surface VkSurfaceKHR, pModes *VkDeviceGroupPresentModeFlagsKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetDeviceGroupSurfacePresentModesKHR.Addr(), 3, uintptr(device), uintptr(surface), uintptr(unsafe.Pointer(pModes)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_DEVICE_LOST, VK_ERROR_OUT_OF_DATE_KHR, VK_ERROR_SURFACE_LOST_KHR, VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT
//On Success: VK_SUCCESS, VK_TIMEOUT, VK_NOT_READY, VK_SUBOPTIMAL_KHR
func vkAcquireNextImage2KHR(device VkDevice, pAcquireInfo *VkAcquireNextImageInfoKHR, pImageIndex *uint32) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkAcquireNextImage2KHR.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pAcquireInfo)), uintptr(unsafe.Pointer(pImageIndex)))	
	ret0 = VkResult(r0)
	return
}


func vkCmdDispatchBase(commandBuffer VkCommandBuffer, baseGroupX uint32, baseGroupY uint32, baseGroupZ uint32, groupCountX uint32, groupCountY uint32, groupCountZ uint32) {
	syscall.Syscall9(procvkCmdDispatchBase.Addr(), 7, uintptr(commandBuffer), uintptr(baseGroupX), uintptr(baseGroupY), uintptr(baseGroupZ), uintptr(groupCountX), uintptr(groupCountY), uintptr(groupCountZ), 0, 0)	
	return
}


func vkCmdDispatchBaseKHR() {
	syscall.Syscall(procvkCmdDispatchBaseKHR.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPhysicalDevicePresentRectanglesKHR(physicalDevice VkPhysicalDevice, surface VkSurfaceKHR, pRectCount *uint32, pRects *VkRect2D) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetPhysicalDevicePresentRectanglesKHR.Addr(), 4, uintptr(physicalDevice), uintptr(surface), uintptr(unsafe.Pointer(pRectCount)), uintptr(unsafe.Pointer(pRects)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateDescriptorUpdateTemplate(device VkDevice, pCreateInfo *VkDescriptorUpdateTemplateCreateInfo, pAllocator *VkAllocationCallbacks, pDescriptorUpdateTemplate *VkDescriptorUpdateTemplate) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateDescriptorUpdateTemplate.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pDescriptorUpdateTemplate)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkCreateDescriptorUpdateTemplateKHR() {
	syscall.Syscall(procvkCreateDescriptorUpdateTemplateKHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkDestroyDescriptorUpdateTemplate(device VkDevice, descriptorUpdateTemplate VkDescriptorUpdateTemplate, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyDescriptorUpdateTemplate.Addr(), 3, uintptr(device), uintptr(descriptorUpdateTemplate), uintptr(unsafe.Pointer(pAllocator)))	
	return
}


func vkDestroyDescriptorUpdateTemplateKHR() {
	syscall.Syscall(procvkDestroyDescriptorUpdateTemplateKHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkUpdateDescriptorSetWithTemplate(device VkDevice, descriptorSet VkDescriptorSet, descriptorUpdateTemplate VkDescriptorUpdateTemplate, pData uintptr) {
	syscall.Syscall6(procvkUpdateDescriptorSetWithTemplate.Addr(), 4, uintptr(device), uintptr(descriptorSet), uintptr(descriptorUpdateTemplate), uintptr(pData), 0, 0)	
	return
}


func vkUpdateDescriptorSetWithTemplateKHR() {
	syscall.Syscall(procvkUpdateDescriptorSetWithTemplateKHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkCmdPushDescriptorSetWithTemplateKHR(commandBuffer VkCommandBuffer, descriptorUpdateTemplate VkDescriptorUpdateTemplate, layout VkPipelineLayout, set uint32, pData uintptr) {
	syscall.Syscall6(procvkCmdPushDescriptorSetWithTemplateKHR.Addr(), 5, uintptr(commandBuffer), uintptr(descriptorUpdateTemplate), uintptr(layout), uintptr(set), uintptr(pData), 0)	
	return
}


func vkSetHdrMetadataEXT(device VkDevice, swapchainCount uint32, pSwapchains *VkSwapchainKHR, pMetadata *VkHdrMetadataEXT) {
	syscall.Syscall6(procvkSetHdrMetadataEXT.Addr(), 4, uintptr(device), uintptr(swapchainCount), uintptr(unsafe.Pointer(pSwapchains)), uintptr(unsafe.Pointer(pMetadata)), 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_DEVICE_LOST, VK_ERROR_OUT_OF_DATE_KHR, VK_ERROR_SURFACE_LOST_KHR, VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT
//On Success: VK_SUCCESS, VK_SUBOPTIMAL_KHR
func vkGetSwapchainStatusKHR(device VkDevice, swapchain VkSwapchainKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetSwapchainStatusKHR.Addr(), 2, uintptr(device), uintptr(swapchain), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_DEVICE_LOST, VK_ERROR_SURFACE_LOST_KHR
//On Success: VK_SUCCESS
func vkGetRefreshCycleDurationGOOGLE(device VkDevice, swapchain VkSwapchainKHR, pDisplayTimingProperties *VkRefreshCycleDurationGOOGLE) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetRefreshCycleDurationGOOGLE.Addr(), 3, uintptr(device), uintptr(swapchain), uintptr(unsafe.Pointer(pDisplayTimingProperties)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_DEVICE_LOST, VK_ERROR_OUT_OF_DATE_KHR, VK_ERROR_SURFACE_LOST_KHR
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPastPresentationTimingGOOGLE(device VkDevice, swapchain VkSwapchainKHR, pPresentationTimingCount *uint32, pPresentationTimings *VkPastPresentationTimingGOOGLE) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetPastPresentationTimingGOOGLE.Addr(), 4, uintptr(device), uintptr(swapchain), uintptr(unsafe.Pointer(pPresentationTimingCount)), uintptr(unsafe.Pointer(pPresentationTimings)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
//On Success: VK_SUCCESS
func vkCreateIOSSurfaceMVK(instance VkInstance, pCreateInfo *VkIOSSurfaceCreateInfoMVK, pAllocator *VkAllocationCallbacks, pSurface *VkSurfaceKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateIOSSurfaceMVK.Addr(), 4, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
//On Success: VK_SUCCESS
func vkCreateMacOSSurfaceMVK(instance VkInstance, pCreateInfo *VkMacOSSurfaceCreateInfoMVK, pAllocator *VkAllocationCallbacks, pSurface *VkSurfaceKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateMacOSSurfaceMVK.Addr(), 4, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
//On Success: VK_SUCCESS
func vkCreateMetalSurfaceEXT(instance VkInstance, pCreateInfo *VkMetalSurfaceCreateInfoEXT, pAllocator *VkAllocationCallbacks, pSurface *VkSurfaceKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateMetalSurfaceEXT.Addr(), 4, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkCmdSetViewportWScalingNV(commandBuffer VkCommandBuffer, firstViewport uint32, viewportCount uint32, pViewportWScalings *VkViewportWScalingNV) {
	syscall.Syscall6(procvkCmdSetViewportWScalingNV.Addr(), 4, uintptr(commandBuffer), uintptr(firstViewport), uintptr(viewportCount), uintptr(unsafe.Pointer(pViewportWScalings)), 0, 0)	
	return
}


func vkCmdSetDiscardRectangleEXT(commandBuffer VkCommandBuffer, firstDiscardRectangle uint32, discardRectangleCount uint32, pDiscardRectangles *VkRect2D) {
	syscall.Syscall6(procvkCmdSetDiscardRectangleEXT.Addr(), 4, uintptr(commandBuffer), uintptr(firstDiscardRectangle), uintptr(discardRectangleCount), uintptr(unsafe.Pointer(pDiscardRectangles)), 0, 0)	
	return
}


func vkCmdSetSampleLocationsEXT(commandBuffer VkCommandBuffer, pSampleLocationsInfo *VkSampleLocationsInfoEXT) {
	syscall.Syscall(procvkCmdSetSampleLocationsEXT.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pSampleLocationsInfo)), 0)	
	return
}


func vkGetPhysicalDeviceMultisamplePropertiesEXT(physicalDevice VkPhysicalDevice, samples VkSampleCountFlagBits, pMultisampleProperties *VkMultisamplePropertiesEXT) {
	syscall.Syscall(procvkGetPhysicalDeviceMultisamplePropertiesEXT.Addr(), 3, uintptr(physicalDevice), uintptr(samples), uintptr(unsafe.Pointer(pMultisampleProperties)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_SURFACE_LOST_KHR
//On Success: VK_SUCCESS
func vkGetPhysicalDeviceSurfaceCapabilities2KHR(physicalDevice VkPhysicalDevice, pSurfaceInfo *VkPhysicalDeviceSurfaceInfo2KHR, pSurfaceCapabilities *VkSurfaceCapabilities2KHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetPhysicalDeviceSurfaceCapabilities2KHR.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pSurfaceInfo)), uintptr(unsafe.Pointer(pSurfaceCapabilities)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_SURFACE_LOST_KHR
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPhysicalDeviceSurfaceFormats2KHR(physicalDevice VkPhysicalDevice, pSurfaceInfo *VkPhysicalDeviceSurfaceInfo2KHR, pSurfaceFormatCount *uint32, pSurfaceFormats *VkSurfaceFormat2KHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetPhysicalDeviceSurfaceFormats2KHR.Addr(), 4, uintptr(physicalDevice), uintptr(unsafe.Pointer(pSurfaceInfo)), uintptr(unsafe.Pointer(pSurfaceFormatCount)), uintptr(unsafe.Pointer(pSurfaceFormats)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPhysicalDeviceDisplayProperties2KHR(physicalDevice VkPhysicalDevice, pPropertyCount *uint32, pProperties *VkDisplayProperties2KHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetPhysicalDeviceDisplayProperties2KHR.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPhysicalDeviceDisplayPlaneProperties2KHR(physicalDevice VkPhysicalDevice, pPropertyCount *uint32, pProperties *VkDisplayPlaneProperties2KHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetPhysicalDeviceDisplayPlaneProperties2KHR.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetDisplayModeProperties2KHR(physicalDevice VkPhysicalDevice, display VkDisplayKHR, pPropertyCount *uint32, pProperties *VkDisplayModeProperties2KHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetDisplayModeProperties2KHR.Addr(), 4, uintptr(physicalDevice), uintptr(display), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkGetDisplayPlaneCapabilities2KHR(physicalDevice VkPhysicalDevice, pDisplayPlaneInfo *VkDisplayPlaneInfo2KHR, pCapabilities *VkDisplayPlaneCapabilities2KHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetDisplayPlaneCapabilities2KHR.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pDisplayPlaneInfo)), uintptr(unsafe.Pointer(pCapabilities)))	
	ret0 = VkResult(r0)
	return
}


func vkGetBufferMemoryRequirements2(device VkDevice, pInfo *VkBufferMemoryRequirementsInfo2, pMemoryRequirements *VkMemoryRequirements2) {
	syscall.Syscall(procvkGetBufferMemoryRequirements2.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pInfo)), uintptr(unsafe.Pointer(pMemoryRequirements)))	
	return
}


func vkGetBufferMemoryRequirements2KHR() {
	syscall.Syscall(procvkGetBufferMemoryRequirements2KHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkGetImageMemoryRequirements2(device VkDevice, pInfo *VkImageMemoryRequirementsInfo2, pMemoryRequirements *VkMemoryRequirements2) {
	syscall.Syscall(procvkGetImageMemoryRequirements2.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pInfo)), uintptr(unsafe.Pointer(pMemoryRequirements)))	
	return
}


func vkGetImageMemoryRequirements2KHR() {
	syscall.Syscall(procvkGetImageMemoryRequirements2KHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkGetImageSparseMemoryRequirements2(device VkDevice, pInfo *VkImageSparseMemoryRequirementsInfo2, pSparseMemoryRequirementCount *uint32, pSparseMemoryRequirements *VkSparseImageMemoryRequirements2) {
	syscall.Syscall6(procvkGetImageSparseMemoryRequirements2.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pInfo)), uintptr(unsafe.Pointer(pSparseMemoryRequirementCount)), uintptr(unsafe.Pointer(pSparseMemoryRequirements)), 0, 0)	
	return
}


func vkGetImageSparseMemoryRequirements2KHR() {
	syscall.Syscall(procvkGetImageSparseMemoryRequirements2KHR.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateSamplerYcbcrConversion(device VkDevice, pCreateInfo *VkSamplerYcbcrConversionCreateInfo, pAllocator *VkAllocationCallbacks, pYcbcrConversion *VkSamplerYcbcrConversion) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateSamplerYcbcrConversion.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pYcbcrConversion)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkCreateSamplerYcbcrConversionKHR() {
	syscall.Syscall(procvkCreateSamplerYcbcrConversionKHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkDestroySamplerYcbcrConversion(device VkDevice, ycbcrConversion VkSamplerYcbcrConversion, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroySamplerYcbcrConversion.Addr(), 3, uintptr(device), uintptr(ycbcrConversion), uintptr(unsafe.Pointer(pAllocator)))	
	return
}


func vkDestroySamplerYcbcrConversionKHR() {
	syscall.Syscall(procvkDestroySamplerYcbcrConversionKHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkGetDeviceQueue2(device VkDevice, pQueueInfo *VkDeviceQueueInfo2, pQueue *VkQueue) {
	syscall.Syscall(procvkGetDeviceQueue2.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pQueueInfo)), uintptr(unsafe.Pointer(pQueue)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkCreateValidationCacheEXT(device VkDevice, pCreateInfo *VkValidationCacheCreateInfoEXT, pAllocator *VkAllocationCallbacks, pValidationCache *VkValidationCacheEXT) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateValidationCacheEXT.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pValidationCache)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyValidationCacheEXT(device VkDevice, validationCache VkValidationCacheEXT, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyValidationCacheEXT.Addr(), 3, uintptr(device), uintptr(validationCache), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetValidationCacheDataEXT(device VkDevice, validationCache VkValidationCacheEXT, pDataSize *uint, pData uintptr) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetValidationCacheDataEXT.Addr(), 4, uintptr(device), uintptr(validationCache), uintptr(unsafe.Pointer(pDataSize)), uintptr(pData), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkMergeValidationCachesEXT(device VkDevice, dstCache VkValidationCacheEXT, srcCacheCount uint32, pSrcCaches *VkValidationCacheEXT) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkMergeValidationCachesEXT.Addr(), 4, uintptr(device), uintptr(dstCache), uintptr(srcCacheCount), uintptr(unsafe.Pointer(pSrcCaches)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkGetDescriptorSetLayoutSupport(device VkDevice, pCreateInfo *VkDescriptorSetLayoutCreateInfo, pSupport *VkDescriptorSetLayoutSupport) {
	syscall.Syscall(procvkGetDescriptorSetLayoutSupport.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pSupport)))	
	return
}


func vkGetDescriptorSetLayoutSupportKHR() {
	syscall.Syscall(procvkGetDescriptorSetLayoutSupportKHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkGetSwapchainGrallocUsageANDROID(device VkDevice, format VkFormat, imageUsage VkImageUsageFlags, grallocUsage *int32) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetSwapchainGrallocUsageANDROID.Addr(), 4, uintptr(device), uintptr(format), uintptr(imageUsage), uintptr(unsafe.Pointer(grallocUsage)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkGetSwapchainGrallocUsage2ANDROID(device VkDevice, format VkFormat, imageUsage VkImageUsageFlags, swapchainImageUsage VkSwapchainImageUsageFlagsANDROID, grallocConsumerUsage *uint64, grallocProducerUsage *uint64) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetSwapchainGrallocUsage2ANDROID.Addr(), 6, uintptr(device), uintptr(format), uintptr(imageUsage), uintptr(swapchainImageUsage), uintptr(unsafe.Pointer(grallocConsumerUsage)), uintptr(unsafe.Pointer(grallocProducerUsage)))	
	ret0 = VkResult(r0)
	return
}


func vkAcquireImageANDROID(device VkDevice, image VkImage, nativeFenceFd int32, semaphore VkSemaphore, fence VkFence) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkAcquireImageANDROID.Addr(), 5, uintptr(device), uintptr(image), uintptr(nativeFenceFd), uintptr(semaphore), uintptr(fence), 0)	
	ret0 = VkResult(r0)
	return
}


func vkQueueSignalReleaseImageANDROID(queue VkQueue, waitSemaphoreCount uint32, pWaitSemaphores *VkSemaphore, image VkImage, pNativeFenceFd *int32) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkQueueSignalReleaseImageANDROID.Addr(), 5, uintptr(queue), uintptr(waitSemaphoreCount), uintptr(unsafe.Pointer(pWaitSemaphores)), uintptr(image), uintptr(unsafe.Pointer(pNativeFenceFd)), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_FEATURE_NOT_PRESENT, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetShaderInfoAMD(device VkDevice, pipeline VkPipeline, shaderStage VkShaderStageFlagBits, infoType VkShaderInfoTypeAMD, pInfoSize *uint, pInfo uintptr) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetShaderInfoAMD.Addr(), 6, uintptr(device), uintptr(pipeline), uintptr(shaderStage), uintptr(infoType), uintptr(unsafe.Pointer(pInfoSize)), uintptr(pInfo))	
	ret0 = VkResult(r0)
	return
}


func vkSetLocalDimmingAMD(device VkDevice, swapChain VkSwapchainKHR, localDimmingEnable VkBool32) {
	syscall.Syscall(procvkSetLocalDimmingAMD.Addr(), 3, uintptr(device), uintptr(swapChain), uintptr(localDimmingEnable))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPhysicalDeviceCalibrateableTimeDomainsEXT(physicalDevice VkPhysicalDevice, pTimeDomainCount *uint32, pTimeDomains *VkTimeDomainEXT) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetPhysicalDeviceCalibrateableTimeDomainsEXT.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pTimeDomainCount)), uintptr(unsafe.Pointer(pTimeDomains)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkGetCalibratedTimestampsEXT(device VkDevice, timestampCount uint32, pTimestampInfos *VkCalibratedTimestampInfoEXT, pTimestamps *uint64, pMaxDeviation *uint64) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetCalibratedTimestampsEXT.Addr(), 5, uintptr(device), uintptr(timestampCount), uintptr(unsafe.Pointer(pTimestampInfos)), uintptr(unsafe.Pointer(pTimestamps)), uintptr(unsafe.Pointer(pMaxDeviation)), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkSetDebugUtilsObjectNameEXT(device VkDevice, pNameInfo *VkDebugUtilsObjectNameInfoEXT) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkSetDebugUtilsObjectNameEXT.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pNameInfo)), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkSetDebugUtilsObjectTagEXT(device VkDevice, pTagInfo *VkDebugUtilsObjectTagInfoEXT) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkSetDebugUtilsObjectTagEXT.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pTagInfo)), 0)	
	ret0 = VkResult(r0)
	return
}


func vkQueueBeginDebugUtilsLabelEXT(queue VkQueue, pLabelInfo *VkDebugUtilsLabelEXT) {
	syscall.Syscall(procvkQueueBeginDebugUtilsLabelEXT.Addr(), 2, uintptr(queue), uintptr(unsafe.Pointer(pLabelInfo)), 0)	
	return
}


func vkQueueEndDebugUtilsLabelEXT(queue VkQueue) {
	syscall.Syscall(procvkQueueEndDebugUtilsLabelEXT.Addr(), 1, uintptr(queue), 0, 0)	
	return
}


func vkQueueInsertDebugUtilsLabelEXT(queue VkQueue, pLabelInfo *VkDebugUtilsLabelEXT) {
	syscall.Syscall(procvkQueueInsertDebugUtilsLabelEXT.Addr(), 2, uintptr(queue), uintptr(unsafe.Pointer(pLabelInfo)), 0)	
	return
}


func vkCmdBeginDebugUtilsLabelEXT(commandBuffer VkCommandBuffer, pLabelInfo *VkDebugUtilsLabelEXT) {
	syscall.Syscall(procvkCmdBeginDebugUtilsLabelEXT.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pLabelInfo)), 0)	
	return
}


func vkCmdEndDebugUtilsLabelEXT(commandBuffer VkCommandBuffer) {
	syscall.Syscall(procvkCmdEndDebugUtilsLabelEXT.Addr(), 1, uintptr(commandBuffer), 0, 0)	
	return
}


func vkCmdInsertDebugUtilsLabelEXT(commandBuffer VkCommandBuffer, pLabelInfo *VkDebugUtilsLabelEXT) {
	syscall.Syscall(procvkCmdInsertDebugUtilsLabelEXT.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pLabelInfo)), 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkCreateDebugUtilsMessengerEXT(instance VkInstance, pCreateInfo *VkDebugUtilsMessengerCreateInfoEXT, pAllocator *VkAllocationCallbacks, pMessenger *VkDebugUtilsMessengerEXT) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateDebugUtilsMessengerEXT.Addr(), 4, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pMessenger)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyDebugUtilsMessengerEXT(instance VkInstance, messenger VkDebugUtilsMessengerEXT, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyDebugUtilsMessengerEXT.Addr(), 3, uintptr(instance), uintptr(messenger), uintptr(unsafe.Pointer(pAllocator)))	
	return
}


func vkSubmitDebugUtilsMessageEXT(instance VkInstance, messageSeverity VkDebugUtilsMessageSeverityFlagBitsEXT, messageTypes VkDebugUtilsMessageTypeFlagsEXT, pCallbackData *VkDebugUtilsMessengerCallbackDataEXT) {
	syscall.Syscall6(procvkSubmitDebugUtilsMessageEXT.Addr(), 4, uintptr(instance), uintptr(messageSeverity), uintptr(messageTypes), uintptr(unsafe.Pointer(pCallbackData)), 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_INVALID_EXTERNAL_HANDLE
//On Success: VK_SUCCESS
func vkGetMemoryHostPointerPropertiesEXT(device VkDevice, handleType VkExternalMemoryHandleTypeFlagBits, pHostPointer uintptr, pMemoryHostPointerProperties *VkMemoryHostPointerPropertiesEXT) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetMemoryHostPointerPropertiesEXT.Addr(), 4, uintptr(device), uintptr(handleType), uintptr(pHostPointer), uintptr(unsafe.Pointer(pMemoryHostPointerProperties)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkCmdWriteBufferMarkerAMD(commandBuffer VkCommandBuffer, pipelineStage VkPipelineStageFlagBits, dstBuffer VkBuffer, dstOffset VkDeviceSize, marker uint32) {
	syscall.Syscall6(procvkCmdWriteBufferMarkerAMD.Addr(), 5, uintptr(commandBuffer), uintptr(pipelineStage), uintptr(dstBuffer), uintptr(dstOffset), uintptr(marker), 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateRenderPass2(device VkDevice, pCreateInfo *VkRenderPassCreateInfo2, pAllocator *VkAllocationCallbacks, pRenderPass *VkRenderPass) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateRenderPass2.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pRenderPass)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkCreateRenderPass2KHR() {
	syscall.Syscall(procvkCreateRenderPass2KHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkCmdBeginRenderPass2(commandBuffer VkCommandBuffer, pRenderPassBegin *VkRenderPassBeginInfo, pSubpassBeginInfo *VkSubpassBeginInfo) {
	syscall.Syscall(procvkCmdBeginRenderPass2.Addr(), 3, uintptr(commandBuffer), uintptr(unsafe.Pointer(pRenderPassBegin)), uintptr(unsafe.Pointer(pSubpassBeginInfo)))	
	return
}


func vkCmdBeginRenderPass2KHR() {
	syscall.Syscall(procvkCmdBeginRenderPass2KHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkCmdNextSubpass2(commandBuffer VkCommandBuffer, pSubpassBeginInfo *VkSubpassBeginInfo, pSubpassEndInfo *VkSubpassEndInfo) {
	syscall.Syscall(procvkCmdNextSubpass2.Addr(), 3, uintptr(commandBuffer), uintptr(unsafe.Pointer(pSubpassBeginInfo)), uintptr(unsafe.Pointer(pSubpassEndInfo)))	
	return
}


func vkCmdNextSubpass2KHR() {
	syscall.Syscall(procvkCmdNextSubpass2KHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkCmdEndRenderPass2(commandBuffer VkCommandBuffer, pSubpassEndInfo *VkSubpassEndInfo) {
	syscall.Syscall(procvkCmdEndRenderPass2.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pSubpassEndInfo)), 0)	
	return
}


func vkCmdEndRenderPass2KHR() {
	syscall.Syscall(procvkCmdEndRenderPass2KHR.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_DEVICE_LOST
//On Success: VK_SUCCESS
func vkGetSemaphoreCounterValue(device VkDevice, semaphore VkSemaphore, pValue *uint64) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetSemaphoreCounterValue.Addr(), 3, uintptr(device), uintptr(semaphore), uintptr(unsafe.Pointer(pValue)))	
	ret0 = VkResult(r0)
	return
}


func vkGetSemaphoreCounterValueKHR() {
	syscall.Syscall(procvkGetSemaphoreCounterValueKHR.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_DEVICE_LOST
//On Success: VK_SUCCESS, VK_TIMEOUT
func vkWaitSemaphores(device VkDevice, pWaitInfo *VkSemaphoreWaitInfo, timeout uint64) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkWaitSemaphores.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pWaitInfo)), uintptr(timeout))	
	ret0 = VkResult(r0)
	return
}


func vkWaitSemaphoresKHR() {
	syscall.Syscall(procvkWaitSemaphoresKHR.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkSignalSemaphore(device VkDevice, pSignalInfo *VkSemaphoreSignalInfo) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkSignalSemaphore.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pSignalInfo)), 0)	
	ret0 = VkResult(r0)
	return
}


func vkSignalSemaphoreKHR() {
	syscall.Syscall(procvkSignalSemaphoreKHR.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_INVALID_EXTERNAL_HANDLE_KHR
//On Success: VK_SUCCESS
func vkGetAndroidHardwareBufferPropertiesANDROID(device VkDevice, buffer *AHardwareBuffer, pProperties *VkAndroidHardwareBufferPropertiesANDROID) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetAndroidHardwareBufferPropertiesANDROID.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(pProperties)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkGetMemoryAndroidHardwareBufferANDROID(device VkDevice, pInfo *VkMemoryGetAndroidHardwareBufferInfoANDROID, pBuffer **AHardwareBuffer) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetMemoryAndroidHardwareBufferANDROID.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pInfo)), uintptr(unsafe.Pointer(pBuffer)))	
	ret0 = VkResult(r0)
	return
}


func vkCmdDrawIndirectCount(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, countBuffer VkBuffer, countBufferOffset VkDeviceSize, maxDrawCount uint32, stride uint32) {
	syscall.Syscall9(procvkCmdDrawIndirectCount.Addr(), 7, uintptr(commandBuffer), uintptr(buffer), uintptr(offset), uintptr(countBuffer), uintptr(countBufferOffset), uintptr(maxDrawCount), uintptr(stride), 0, 0)	
	return
}


func vkCmdDrawIndirectCountKHR() {
	syscall.Syscall(procvkCmdDrawIndirectCountKHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkCmdDrawIndirectCountAMD() {
	syscall.Syscall(procvkCmdDrawIndirectCountAMD.Addr(), 0, 0, 0, 0)	
	return
}


func vkCmdDrawIndexedIndirectCount(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, countBuffer VkBuffer, countBufferOffset VkDeviceSize, maxDrawCount uint32, stride uint32) {
	syscall.Syscall9(procvkCmdDrawIndexedIndirectCount.Addr(), 7, uintptr(commandBuffer), uintptr(buffer), uintptr(offset), uintptr(countBuffer), uintptr(countBufferOffset), uintptr(maxDrawCount), uintptr(stride), 0, 0)	
	return
}


func vkCmdDrawIndexedIndirectCountKHR() {
	syscall.Syscall(procvkCmdDrawIndexedIndirectCountKHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkCmdDrawIndexedIndirectCountAMD() {
	syscall.Syscall(procvkCmdDrawIndexedIndirectCountAMD.Addr(), 0, 0, 0, 0)	
	return
}


func vkCmdSetCheckpointNV(commandBuffer VkCommandBuffer, pCheckpointMarker uintptr) {
	syscall.Syscall(procvkCmdSetCheckpointNV.Addr(), 2, uintptr(commandBuffer), uintptr(pCheckpointMarker), 0)	
	return
}


func vkGetQueueCheckpointDataNV(queue VkQueue, pCheckpointDataCount *uint32, pCheckpointData *VkCheckpointDataNV) {
	syscall.Syscall(procvkGetQueueCheckpointDataNV.Addr(), 3, uintptr(queue), uintptr(unsafe.Pointer(pCheckpointDataCount)), uintptr(unsafe.Pointer(pCheckpointData)))	
	return
}


func vkCmdBindTransformFeedbackBuffersEXT(commandBuffer VkCommandBuffer, firstBinding uint32, bindingCount uint32, pBuffers *VkBuffer, pOffsets *VkDeviceSize, pSizes *VkDeviceSize) {
	syscall.Syscall6(procvkCmdBindTransformFeedbackBuffersEXT.Addr(), 6, uintptr(commandBuffer), uintptr(firstBinding), uintptr(bindingCount), uintptr(unsafe.Pointer(pBuffers)), uintptr(unsafe.Pointer(pOffsets)), uintptr(unsafe.Pointer(pSizes)))	
	return
}


func vkCmdBeginTransformFeedbackEXT(commandBuffer VkCommandBuffer, firstCounterBuffer uint32, counterBufferCount uint32, pCounterBuffers *VkBuffer, pCounterBufferOffsets *VkDeviceSize) {
	syscall.Syscall6(procvkCmdBeginTransformFeedbackEXT.Addr(), 5, uintptr(commandBuffer), uintptr(firstCounterBuffer), uintptr(counterBufferCount), uintptr(unsafe.Pointer(pCounterBuffers)), uintptr(unsafe.Pointer(pCounterBufferOffsets)), 0)	
	return
}


func vkCmdEndTransformFeedbackEXT(commandBuffer VkCommandBuffer, firstCounterBuffer uint32, counterBufferCount uint32, pCounterBuffers *VkBuffer, pCounterBufferOffsets *VkDeviceSize) {
	syscall.Syscall6(procvkCmdEndTransformFeedbackEXT.Addr(), 5, uintptr(commandBuffer), uintptr(firstCounterBuffer), uintptr(counterBufferCount), uintptr(unsafe.Pointer(pCounterBuffers)), uintptr(unsafe.Pointer(pCounterBufferOffsets)), 0)	
	return
}


func vkCmdBeginQueryIndexedEXT(commandBuffer VkCommandBuffer, queryPool VkQueryPool, query uint32, flags VkQueryControlFlags, index uint32) {
	syscall.Syscall6(procvkCmdBeginQueryIndexedEXT.Addr(), 5, uintptr(commandBuffer), uintptr(queryPool), uintptr(query), uintptr(flags), uintptr(index), 0)	
	return
}


func vkCmdEndQueryIndexedEXT(commandBuffer VkCommandBuffer, queryPool VkQueryPool, query uint32, index uint32) {
	syscall.Syscall6(procvkCmdEndQueryIndexedEXT.Addr(), 4, uintptr(commandBuffer), uintptr(queryPool), uintptr(query), uintptr(index), 0, 0)	
	return
}


func vkCmdDrawIndirectByteCountEXT(commandBuffer VkCommandBuffer, instanceCount uint32, firstInstance uint32, counterBuffer VkBuffer, counterBufferOffset VkDeviceSize, counterOffset uint32, vertexStride uint32) {
	syscall.Syscall9(procvkCmdDrawIndirectByteCountEXT.Addr(), 7, uintptr(commandBuffer), uintptr(instanceCount), uintptr(firstInstance), uintptr(counterBuffer), uintptr(counterBufferOffset), uintptr(counterOffset), uintptr(vertexStride), 0, 0)	
	return
}


func vkCmdSetExclusiveScissorNV(commandBuffer VkCommandBuffer, firstExclusiveScissor uint32, exclusiveScissorCount uint32, pExclusiveScissors *VkRect2D) {
	syscall.Syscall6(procvkCmdSetExclusiveScissorNV.Addr(), 4, uintptr(commandBuffer), uintptr(firstExclusiveScissor), uintptr(exclusiveScissorCount), uintptr(unsafe.Pointer(pExclusiveScissors)), 0, 0)	
	return
}


func vkCmdBindShadingRateImageNV(commandBuffer VkCommandBuffer, imageView VkImageView, imageLayout VkImageLayout) {
	syscall.Syscall(procvkCmdBindShadingRateImageNV.Addr(), 3, uintptr(commandBuffer), uintptr(imageView), uintptr(imageLayout))	
	return
}


func vkCmdSetViewportShadingRatePaletteNV(commandBuffer VkCommandBuffer, firstViewport uint32, viewportCount uint32, pShadingRatePalettes *VkShadingRatePaletteNV) {
	syscall.Syscall6(procvkCmdSetViewportShadingRatePaletteNV.Addr(), 4, uintptr(commandBuffer), uintptr(firstViewport), uintptr(viewportCount), uintptr(unsafe.Pointer(pShadingRatePalettes)), 0, 0)	
	return
}


func vkCmdSetCoarseSampleOrderNV(commandBuffer VkCommandBuffer, sampleOrderType VkCoarseSampleOrderTypeNV, customSampleOrderCount uint32, pCustomSampleOrders *VkCoarseSampleOrderCustomNV) {
	syscall.Syscall6(procvkCmdSetCoarseSampleOrderNV.Addr(), 4, uintptr(commandBuffer), uintptr(sampleOrderType), uintptr(customSampleOrderCount), uintptr(unsafe.Pointer(pCustomSampleOrders)), 0, 0)	
	return
}


func vkCmdDrawMeshTasksNV(commandBuffer VkCommandBuffer, taskCount uint32, firstTask uint32) {
	syscall.Syscall(procvkCmdDrawMeshTasksNV.Addr(), 3, uintptr(commandBuffer), uintptr(taskCount), uintptr(firstTask))	
	return
}


func vkCmdDrawMeshTasksIndirectNV(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, drawCount uint32, stride uint32) {
	syscall.Syscall6(procvkCmdDrawMeshTasksIndirectNV.Addr(), 5, uintptr(commandBuffer), uintptr(buffer), uintptr(offset), uintptr(drawCount), uintptr(stride), 0)	
	return
}


func vkCmdDrawMeshTasksIndirectCountNV(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, countBuffer VkBuffer, countBufferOffset VkDeviceSize, maxDrawCount uint32, stride uint32) {
	syscall.Syscall9(procvkCmdDrawMeshTasksIndirectCountNV.Addr(), 7, uintptr(commandBuffer), uintptr(buffer), uintptr(offset), uintptr(countBuffer), uintptr(countBufferOffset), uintptr(maxDrawCount), uintptr(stride), 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCompileDeferredNV(device VkDevice, pipeline VkPipeline, shader uint32) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkCompileDeferredNV.Addr(), 3, uintptr(device), uintptr(pipeline), uintptr(shader))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkCreateAccelerationStructureNV(device VkDevice, pCreateInfo *VkAccelerationStructureCreateInfoNV, pAllocator *VkAllocationCallbacks, pAccelerationStructure *VkAccelerationStructureNV) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateAccelerationStructureNV.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pAccelerationStructure)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyAccelerationStructureKHR(device VkDevice, accelerationStructure VkAccelerationStructureKHR, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyAccelerationStructureKHR.Addr(), 3, uintptr(device), uintptr(accelerationStructure), uintptr(unsafe.Pointer(pAllocator)))	
	return
}


func vkDestroyAccelerationStructureNV() {
	syscall.Syscall(procvkDestroyAccelerationStructureNV.Addr(), 0, 0, 0, 0)	
	return
}


func vkGetAccelerationStructureMemoryRequirementsKHR(device VkDevice, pInfo *VkAccelerationStructureMemoryRequirementsInfoKHR, pMemoryRequirements *VkMemoryRequirements2) {
	syscall.Syscall(procvkGetAccelerationStructureMemoryRequirementsKHR.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pInfo)), uintptr(unsafe.Pointer(pMemoryRequirements)))	
	return
}


func vkGetAccelerationStructureMemoryRequirementsNV(device VkDevice, pInfo *VkAccelerationStructureMemoryRequirementsInfoNV, pMemoryRequirements *VkMemoryRequirements2KHR) {
	syscall.Syscall(procvkGetAccelerationStructureMemoryRequirementsNV.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pInfo)), uintptr(unsafe.Pointer(pMemoryRequirements)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkBindAccelerationStructureMemoryKHR(device VkDevice, bindInfoCount uint32, pBindInfos *VkBindAccelerationStructureMemoryInfoKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkBindAccelerationStructureMemoryKHR.Addr(), 3, uintptr(device), uintptr(bindInfoCount), uintptr(unsafe.Pointer(pBindInfos)))	
	ret0 = VkResult(r0)
	return
}


func vkBindAccelerationStructureMemoryNV() {
	syscall.Syscall(procvkBindAccelerationStructureMemoryNV.Addr(), 0, 0, 0, 0)	
	return
}


func vkCmdCopyAccelerationStructureNV(commandBuffer VkCommandBuffer, dst VkAccelerationStructureKHR, src VkAccelerationStructureKHR, mode VkCopyAccelerationStructureModeKHR) {
	syscall.Syscall6(procvkCmdCopyAccelerationStructureNV.Addr(), 4, uintptr(commandBuffer), uintptr(dst), uintptr(src), uintptr(mode), 0, 0)	
	return
}


func vkCmdCopyAccelerationStructureKHR(commandBuffer VkCommandBuffer, pInfo *VkCopyAccelerationStructureInfoKHR) {
	syscall.Syscall(procvkCmdCopyAccelerationStructureKHR.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pInfo)), 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_OPERATION_DEFERRED_KHR, VK_OPERATION_NOT_DEFERRED_KHR
func vkCopyAccelerationStructureKHR(device VkDevice, pInfo *VkCopyAccelerationStructureInfoKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkCopyAccelerationStructureKHR.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pInfo)), 0)	
	ret0 = VkResult(r0)
	return
}


func vkCmdCopyAccelerationStructureToMemoryKHR(commandBuffer VkCommandBuffer, pInfo *VkCopyAccelerationStructureToMemoryInfoKHR) {
	syscall.Syscall(procvkCmdCopyAccelerationStructureToMemoryKHR.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pInfo)), 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_OPERATION_DEFERRED_KHR, VK_OPERATION_NOT_DEFERRED_KHR
func vkCopyAccelerationStructureToMemoryKHR(device VkDevice, pInfo *VkCopyAccelerationStructureToMemoryInfoKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkCopyAccelerationStructureToMemoryKHR.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pInfo)), 0)	
	ret0 = VkResult(r0)
	return
}


func vkCmdCopyMemoryToAccelerationStructureKHR(commandBuffer VkCommandBuffer, pInfo *VkCopyMemoryToAccelerationStructureInfoKHR) {
	syscall.Syscall(procvkCmdCopyMemoryToAccelerationStructureKHR.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pInfo)), 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_OPERATION_DEFERRED_KHR, VK_OPERATION_NOT_DEFERRED_KHR
func vkCopyMemoryToAccelerationStructureKHR(device VkDevice, pInfo *VkCopyMemoryToAccelerationStructureInfoKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkCopyMemoryToAccelerationStructureKHR.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pInfo)), 0)	
	ret0 = VkResult(r0)
	return
}


func vkCmdWriteAccelerationStructuresPropertiesKHR(commandBuffer VkCommandBuffer, accelerationStructureCount uint32, pAccelerationStructures *VkAccelerationStructureKHR, queryType VkQueryType, queryPool VkQueryPool, firstQuery uint32) {
	syscall.Syscall6(procvkCmdWriteAccelerationStructuresPropertiesKHR.Addr(), 6, uintptr(commandBuffer), uintptr(accelerationStructureCount), uintptr(unsafe.Pointer(pAccelerationStructures)), uintptr(queryType), uintptr(queryPool), uintptr(firstQuery))	
	return
}


func vkCmdWriteAccelerationStructuresPropertiesNV() {
	syscall.Syscall(procvkCmdWriteAccelerationStructuresPropertiesNV.Addr(), 0, 0, 0, 0)	
	return
}


func vkCmdBuildAccelerationStructureNV(commandBuffer VkCommandBuffer, pInfo *VkAccelerationStructureInfoNV, instanceData VkBuffer, instanceOffset VkDeviceSize, update VkBool32, dst VkAccelerationStructureKHR, src VkAccelerationStructureKHR, scratch VkBuffer, scratchOffset VkDeviceSize) {
	syscall.Syscall9(procvkCmdBuildAccelerationStructureNV.Addr(), 9, uintptr(commandBuffer), uintptr(unsafe.Pointer(pInfo)), uintptr(instanceData), uintptr(instanceOffset), uintptr(update), uintptr(dst), uintptr(src), uintptr(scratch), uintptr(scratchOffset))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkWriteAccelerationStructuresPropertiesKHR(device VkDevice, accelerationStructureCount uint32, pAccelerationStructures *VkAccelerationStructureKHR, queryType VkQueryType, dataSize uint, pData uintptr, stride uint) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall9(procvkWriteAccelerationStructuresPropertiesKHR.Addr(), 7, uintptr(device), uintptr(accelerationStructureCount), uintptr(unsafe.Pointer(pAccelerationStructures)), uintptr(queryType), uintptr(dataSize), uintptr(pData), uintptr(stride), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkCmdTraceRaysKHR(commandBuffer VkCommandBuffer, pRaygenShaderBindingTable *VkStridedBufferRegionKHR, pMissShaderBindingTable *VkStridedBufferRegionKHR, pHitShaderBindingTable *VkStridedBufferRegionKHR, pCallableShaderBindingTable *VkStridedBufferRegionKHR, width uint32, height uint32, depth uint32) {
	syscall.Syscall9(procvkCmdTraceRaysKHR.Addr(), 8, uintptr(commandBuffer), uintptr(unsafe.Pointer(pRaygenShaderBindingTable)), uintptr(unsafe.Pointer(pMissShaderBindingTable)), uintptr(unsafe.Pointer(pHitShaderBindingTable)), uintptr(unsafe.Pointer(pCallableShaderBindingTable)), uintptr(width), uintptr(height), uintptr(depth), 0)	
	return
}


func vkCmdTraceRaysNV(commandBuffer VkCommandBuffer, raygenShaderBindingTableBuffer VkBuffer, raygenShaderBindingOffset VkDeviceSize, missShaderBindingTableBuffer VkBuffer, missShaderBindingOffset VkDeviceSize, missShaderBindingStride VkDeviceSize, hitShaderBindingTableBuffer VkBuffer, hitShaderBindingOffset VkDeviceSize, hitShaderBindingStride VkDeviceSize, callableShaderBindingTableBuffer VkBuffer, callableShaderBindingOffset VkDeviceSize, callableShaderBindingStride VkDeviceSize, width uint32, height uint32, depth uint32) {
	syscall.Syscall15(procvkCmdTraceRaysNV.Addr(), 15, uintptr(commandBuffer), uintptr(raygenShaderBindingTableBuffer), uintptr(raygenShaderBindingOffset), uintptr(missShaderBindingTableBuffer), uintptr(missShaderBindingOffset), uintptr(missShaderBindingStride), uintptr(hitShaderBindingTableBuffer), uintptr(hitShaderBindingOffset), uintptr(hitShaderBindingStride), uintptr(callableShaderBindingTableBuffer), uintptr(callableShaderBindingOffset), uintptr(callableShaderBindingStride), uintptr(width), uintptr(height), uintptr(depth))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkGetRayTracingShaderGroupHandlesKHR(device VkDevice, pipeline VkPipeline, firstGroup uint32, groupCount uint32, dataSize uint, pData uintptr) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetRayTracingShaderGroupHandlesKHR.Addr(), 6, uintptr(device), uintptr(pipeline), uintptr(firstGroup), uintptr(groupCount), uintptr(dataSize), uintptr(pData))	
	ret0 = VkResult(r0)
	return
}


func vkGetRayTracingShaderGroupHandlesNV() {
	syscall.Syscall(procvkGetRayTracingShaderGroupHandlesNV.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkGetRayTracingCaptureReplayShaderGroupHandlesKHR(device VkDevice, pipeline VkPipeline, firstGroup uint32, groupCount uint32, dataSize uint, pData uintptr) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetRayTracingCaptureReplayShaderGroupHandlesKHR.Addr(), 6, uintptr(device), uintptr(pipeline), uintptr(firstGroup), uintptr(groupCount), uintptr(dataSize), uintptr(pData))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkGetAccelerationStructureHandleNV(device VkDevice, accelerationStructure VkAccelerationStructureKHR, dataSize uint, pData uintptr) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetAccelerationStructureHandleNV.Addr(), 4, uintptr(device), uintptr(accelerationStructure), uintptr(dataSize), uintptr(pData), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INVALID_SHADER_NV
//On Success: VK_SUCCESS, VK_PIPELINE_COMPILE_REQUIRED_EXT
func vkCreateRayTracingPipelinesNV(device VkDevice, pipelineCache VkPipelineCache, createInfoCount uint32, pCreateInfos *VkRayTracingPipelineCreateInfoNV, pAllocator *VkAllocationCallbacks, pPipelines *VkPipeline) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateRayTracingPipelinesNV.Addr(), 6, uintptr(device), uintptr(pipelineCache), uintptr(createInfoCount), uintptr(unsafe.Pointer(pCreateInfos)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pPipelines)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS
//On Success: VK_SUCCESS, VK_OPERATION_DEFERRED_KHR, VK_OPERATION_NOT_DEFERRED_KHR, VK_PIPELINE_COMPILE_REQUIRED_EXT
func vkCreateRayTracingPipelinesKHR(device VkDevice, pipelineCache VkPipelineCache, createInfoCount uint32, pCreateInfos *VkRayTracingPipelineCreateInfoKHR, pAllocator *VkAllocationCallbacks, pPipelines *VkPipeline) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateRayTracingPipelinesKHR.Addr(), 6, uintptr(device), uintptr(pipelineCache), uintptr(createInfoCount), uintptr(unsafe.Pointer(pCreateInfos)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pPipelines)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPhysicalDeviceCooperativeMatrixPropertiesNV(physicalDevice VkPhysicalDevice, pPropertyCount *uint32, pProperties *VkCooperativeMatrixPropertiesNV) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetPhysicalDeviceCooperativeMatrixPropertiesNV.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))	
	ret0 = VkResult(r0)
	return
}


func vkCmdTraceRaysIndirectKHR(commandBuffer VkCommandBuffer, pRaygenShaderBindingTable *VkStridedBufferRegionKHR, pMissShaderBindingTable *VkStridedBufferRegionKHR, pHitShaderBindingTable *VkStridedBufferRegionKHR, pCallableShaderBindingTable *VkStridedBufferRegionKHR, buffer VkBuffer, offset VkDeviceSize) {
	syscall.Syscall9(procvkCmdTraceRaysIndirectKHR.Addr(), 7, uintptr(commandBuffer), uintptr(unsafe.Pointer(pRaygenShaderBindingTable)), uintptr(unsafe.Pointer(pMissShaderBindingTable)), uintptr(unsafe.Pointer(pHitShaderBindingTable)), uintptr(unsafe.Pointer(pCallableShaderBindingTable)), uintptr(buffer), uintptr(offset), 0, 0)	
	return
}

//Errors: VK_ERROR_INCOMPATIBLE_VERSION_KHR
//On Success: VK_SUCCESS
func vkGetDeviceAccelerationStructureCompatibilityKHR(device VkDevice, version *VkAccelerationStructureVersionKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetDeviceAccelerationStructureCompatibilityKHR.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(version)), 0)	
	ret0 = VkResult(r0)
	return
}


func vkGetImageViewHandleNVX(device VkDevice, pInfo *VkImageViewHandleInfoNVX) (ret0 uint32) {
	r0, _, _ := syscall.Syscall(procvkGetImageViewHandleNVX.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pInfo)), 0)	
	ret0 = uint32(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_UNKNOWN
//On Success: VK_SUCCESS
func vkGetImageViewAddressNVX(device VkDevice, imageView VkImageView, pProperties *VkImageViewAddressPropertiesNVX) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetImageViewAddressNVX.Addr(), 3, uintptr(device), uintptr(imageView), uintptr(unsafe.Pointer(pProperties)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_SURFACE_LOST_KHR
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPhysicalDeviceSurfacePresentModes2EXT(physicalDevice VkPhysicalDevice, pSurfaceInfo *VkPhysicalDeviceSurfaceInfo2KHR, pPresentModeCount *uint32, pPresentModes *VkPresentModeKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetPhysicalDeviceSurfacePresentModes2EXT.Addr(), 4, uintptr(physicalDevice), uintptr(unsafe.Pointer(pSurfaceInfo)), uintptr(unsafe.Pointer(pPresentModeCount)), uintptr(unsafe.Pointer(pPresentModes)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_SURFACE_LOST_KHR
//On Success: VK_SUCCESS
func vkGetDeviceGroupSurfacePresentModes2EXT(device VkDevice, pSurfaceInfo *VkPhysicalDeviceSurfaceInfo2KHR, pModes *VkDeviceGroupPresentModeFlagsKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetDeviceGroupSurfacePresentModes2EXT.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pSurfaceInfo)), uintptr(unsafe.Pointer(pModes)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INITIALIZATION_FAILED, VK_ERROR_SURFACE_LOST_KHR
//On Success: VK_SUCCESS
func vkAcquireFullScreenExclusiveModeEXT(device VkDevice, swapchain VkSwapchainKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkAcquireFullScreenExclusiveModeEXT.Addr(), 2, uintptr(device), uintptr(swapchain), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_SURFACE_LOST_KHR
//On Success: VK_SUCCESS
func vkReleaseFullScreenExclusiveModeEXT(device VkDevice, swapchain VkSwapchainKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkReleaseFullScreenExclusiveModeEXT.Addr(), 2, uintptr(device), uintptr(swapchain), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY, VK_ERROR_INITIALIZATION_FAILED
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR(physicalDevice VkPhysicalDevice, queueFamilyIndex uint32, pCounterCount *uint32, pCounters *VkPerformanceCounterKHR, pCounterDescriptions *VkPerformanceCounterDescriptionKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR.Addr(), 5, uintptr(physicalDevice), uintptr(queueFamilyIndex), uintptr(unsafe.Pointer(pCounterCount)), uintptr(unsafe.Pointer(pCounters)), uintptr(unsafe.Pointer(pCounterDescriptions)), 0)	
	ret0 = VkResult(r0)
	return
}


func vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR(physicalDevice VkPhysicalDevice, pPerformanceQueryCreateInfo *VkQueryPoolPerformanceCreateInfoKHR, pNumPasses *uint32) {
	syscall.Syscall(procvkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pPerformanceQueryCreateInfo)), uintptr(unsafe.Pointer(pNumPasses)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_TIMEOUT
//On Success: VK_SUCCESS
func vkAcquireProfilingLockKHR(device VkDevice, pInfo *VkAcquireProfilingLockInfoKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkAcquireProfilingLockKHR.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pInfo)), 0)	
	ret0 = VkResult(r0)
	return
}


func vkReleaseProfilingLockKHR(device VkDevice) {
	syscall.Syscall(procvkReleaseProfilingLockKHR.Addr(), 1, uintptr(device), 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkGetImageDrmFormatModifierPropertiesEXT(device VkDevice, image VkImage, pProperties *VkImageDrmFormatModifierPropertiesEXT) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetImageDrmFormatModifierPropertiesEXT.Addr(), 3, uintptr(device), uintptr(image), uintptr(unsafe.Pointer(pProperties)))	
	ret0 = VkResult(r0)
	return
}


func vkGetBufferOpaqueCaptureAddress(device VkDevice, pInfo *VkBufferDeviceAddressInfo) (ret0 uint64) {
	r0, _, _ := syscall.Syscall(procvkGetBufferOpaqueCaptureAddress.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pInfo)), 0)	
	ret0 = uint64(r0)
	return
}


func vkGetBufferOpaqueCaptureAddressKHR() {
	syscall.Syscall(procvkGetBufferOpaqueCaptureAddressKHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkGetBufferDeviceAddress(device VkDevice, pInfo *VkBufferDeviceAddressInfo) (ret0 VkDeviceAddress) {
	r0, _, _ := syscall.Syscall(procvkGetBufferDeviceAddress.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pInfo)), 0)	
	ret0 = VkDeviceAddress(r0)
	return
}


func vkGetBufferDeviceAddressKHR() {
	syscall.Syscall(procvkGetBufferDeviceAddressKHR.Addr(), 0, 0, 0, 0)	
	return
}


func vkGetBufferDeviceAddressEXT() {
	syscall.Syscall(procvkGetBufferDeviceAddressEXT.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS
func vkCreateHeadlessSurfaceEXT(instance VkInstance, pCreateInfo *VkHeadlessSurfaceCreateInfoEXT, pAllocator *VkAllocationCallbacks, pSurface *VkSurfaceKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateHeadlessSurfaceEXT.Addr(), 4, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV(physicalDevice VkPhysicalDevice, pCombinationCount *uint32, pCombinations *VkFramebufferMixedSamplesCombinationNV) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pCombinationCount)), uintptr(unsafe.Pointer(pCombinations)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkInitializePerformanceApiINTEL(device VkDevice, pInitializeInfo *VkInitializePerformanceApiInfoINTEL) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkInitializePerformanceApiINTEL.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pInitializeInfo)), 0)	
	ret0 = VkResult(r0)
	return
}


func vkUninitializePerformanceApiINTEL(device VkDevice) {
	syscall.Syscall(procvkUninitializePerformanceApiINTEL.Addr(), 1, uintptr(device), 0, 0)	
	return
}

//Errors: VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkCmdSetPerformanceMarkerINTEL(commandBuffer VkCommandBuffer, pMarkerInfo *VkPerformanceMarkerInfoINTEL) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkCmdSetPerformanceMarkerINTEL.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pMarkerInfo)), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkCmdSetPerformanceStreamMarkerINTEL(commandBuffer VkCommandBuffer, pMarkerInfo *VkPerformanceStreamMarkerInfoINTEL) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkCmdSetPerformanceStreamMarkerINTEL.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pMarkerInfo)), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkCmdSetPerformanceOverrideINTEL(commandBuffer VkCommandBuffer, pOverrideInfo *VkPerformanceOverrideInfoINTEL) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkCmdSetPerformanceOverrideINTEL.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pOverrideInfo)), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkAcquirePerformanceConfigurationINTEL(device VkDevice, pAcquireInfo *VkPerformanceConfigurationAcquireInfoINTEL, pConfiguration *VkPerformanceConfigurationINTEL) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkAcquirePerformanceConfigurationINTEL.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pAcquireInfo)), uintptr(unsafe.Pointer(pConfiguration)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkReleasePerformanceConfigurationINTEL(device VkDevice, configuration VkPerformanceConfigurationINTEL) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkReleasePerformanceConfigurationINTEL.Addr(), 2, uintptr(device), uintptr(configuration), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkQueueSetPerformanceConfigurationINTEL(queue VkQueue, configuration VkPerformanceConfigurationINTEL) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkQueueSetPerformanceConfigurationINTEL.Addr(), 2, uintptr(queue), uintptr(configuration), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_TOO_MANY_OBJECTS, VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkGetPerformanceParameterINTEL(device VkDevice, parameter VkPerformanceParameterTypeINTEL, pValue *VkPerformanceValueINTEL) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetPerformanceParameterINTEL.Addr(), 3, uintptr(device), uintptr(parameter), uintptr(unsafe.Pointer(pValue)))	
	ret0 = VkResult(r0)
	return
}


func vkGetDeviceMemoryOpaqueCaptureAddress(device VkDevice, pInfo *VkDeviceMemoryOpaqueCaptureAddressInfo) (ret0 uint64) {
	r0, _, _ := syscall.Syscall(procvkGetDeviceMemoryOpaqueCaptureAddress.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pInfo)), 0)	
	ret0 = uint64(r0)
	return
}


func vkGetDeviceMemoryOpaqueCaptureAddressKHR() {
	syscall.Syscall(procvkGetDeviceMemoryOpaqueCaptureAddressKHR.Addr(), 0, 0, 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPipelineExecutablePropertiesKHR(device VkDevice, pPipelineInfo *VkPipelineInfoKHR, pExecutableCount *uint32, pProperties *VkPipelineExecutablePropertiesKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetPipelineExecutablePropertiesKHR.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pPipelineInfo)), uintptr(unsafe.Pointer(pExecutableCount)), uintptr(unsafe.Pointer(pProperties)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPipelineExecutableStatisticsKHR(device VkDevice, pExecutableInfo *VkPipelineExecutableInfoKHR, pStatisticCount *uint32, pStatistics *VkPipelineExecutableStatisticKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetPipelineExecutableStatisticsKHR.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pExecutableInfo)), uintptr(unsafe.Pointer(pStatisticCount)), uintptr(unsafe.Pointer(pStatistics)), 0, 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPipelineExecutableInternalRepresentationsKHR(device VkDevice, pExecutableInfo *VkPipelineExecutableInfoKHR, pInternalRepresentationCount *uint32, pInternalRepresentations *VkPipelineExecutableInternalRepresentationKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkGetPipelineExecutableInternalRepresentationsKHR.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pExecutableInfo)), uintptr(unsafe.Pointer(pInternalRepresentationCount)), uintptr(unsafe.Pointer(pInternalRepresentations)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkCmdSetLineStippleEXT(commandBuffer VkCommandBuffer, lineStippleFactor uint32, lineStipplePattern uint16) {
	syscall.Syscall(procvkCmdSetLineStippleEXT.Addr(), 3, uintptr(commandBuffer), uintptr(lineStippleFactor), uintptr(lineStipplePattern))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPhysicalDeviceToolPropertiesEXT(physicalDevice VkPhysicalDevice, pToolCount *uint32, pToolProperties *VkPhysicalDeviceToolPropertiesEXT) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetPhysicalDeviceToolPropertiesEXT.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pToolCount)), uintptr(unsafe.Pointer(pToolProperties)))	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR
//On Success: VK_SUCCESS
func vkCreateAccelerationStructureKHR(device VkDevice, pCreateInfo *VkAccelerationStructureCreateInfoKHR, pAllocator *VkAllocationCallbacks, pAccelerationStructure *VkAccelerationStructureKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreateAccelerationStructureKHR.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pAccelerationStructure)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkCmdBuildAccelerationStructureKHR(commandBuffer VkCommandBuffer, infoCount uint32, pInfos *VkAccelerationStructureBuildGeometryInfoKHR, ppOffsetInfos **VkAccelerationStructureBuildOffsetInfoKHR) {
	syscall.Syscall6(procvkCmdBuildAccelerationStructureKHR.Addr(), 4, uintptr(commandBuffer), uintptr(infoCount), uintptr(unsafe.Pointer(pInfos)), uintptr(unsafe.Pointer(ppOffsetInfos)), 0, 0)	
	return
}


func vkCmdBuildAccelerationStructureIndirectKHR(commandBuffer VkCommandBuffer, pInfo *VkAccelerationStructureBuildGeometryInfoKHR, indirectBuffer VkBuffer, indirectOffset VkDeviceSize, indirectStride uint32) {
	syscall.Syscall6(procvkCmdBuildAccelerationStructureIndirectKHR.Addr(), 5, uintptr(commandBuffer), uintptr(unsafe.Pointer(pInfo)), uintptr(indirectBuffer), uintptr(indirectOffset), uintptr(indirectStride), 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_OPERATION_DEFERRED_KHR, VK_OPERATION_NOT_DEFERRED_KHR
func vkBuildAccelerationStructureKHR(device VkDevice, infoCount uint32, pInfos *VkAccelerationStructureBuildGeometryInfoKHR, ppOffsetInfos **VkAccelerationStructureBuildOffsetInfoKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkBuildAccelerationStructureKHR.Addr(), 4, uintptr(device), uintptr(infoCount), uintptr(unsafe.Pointer(pInfos)), uintptr(unsafe.Pointer(ppOffsetInfos)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkGetAccelerationStructureDeviceAddressKHR(device VkDevice, pInfo *VkAccelerationStructureDeviceAddressInfoKHR) (ret0 VkDeviceAddress) {
	r0, _, _ := syscall.Syscall(procvkGetAccelerationStructureDeviceAddressKHR.Addr(), 2, uintptr(device), uintptr(unsafe.Pointer(pInfo)), 0)	
	ret0 = VkDeviceAddress(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkCreateDeferredOperationKHR(device VkDevice, pAllocator *VkAllocationCallbacks, pDeferredOperation *VkDeferredOperationKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkCreateDeferredOperationKHR.Addr(), 3, uintptr(device), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pDeferredOperation)))	
	ret0 = VkResult(r0)
	return
}


func vkDestroyDeferredOperationKHR(device VkDevice, operation VkDeferredOperationKHR, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyDeferredOperationKHR.Addr(), 3, uintptr(device), uintptr(operation), uintptr(unsafe.Pointer(pAllocator)))	
	return
}


func vkGetDeferredOperationMaxConcurrencyKHR(device VkDevice, operation VkDeferredOperationKHR) (ret0 uint32) {
	r0, _, _ := syscall.Syscall(procvkGetDeferredOperationMaxConcurrencyKHR.Addr(), 2, uintptr(device), uintptr(operation), 0)	
	ret0 = uint32(r0)
	return
}

//On Success: VK_SUCCESS, VK_NOT_READY
func vkGetDeferredOperationResultKHR(device VkDevice, operation VkDeferredOperationKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetDeferredOperationResultKHR.Addr(), 2, uintptr(device), uintptr(operation), 0)	
	ret0 = VkResult(r0)
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY, VK_ERROR_OUT_OF_DEVICE_MEMORY
//On Success: VK_SUCCESS, VK_THREAD_DONE_KHR, VK_THREAD_IDLE_KHR
func vkDeferredOperationJoinKHR(device VkDevice, operation VkDeferredOperationKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkDeferredOperationJoinKHR.Addr(), 2, uintptr(device), uintptr(operation), 0)	
	ret0 = VkResult(r0)
	return
}


func vkCmdSetCullModeEXT(commandBuffer VkCommandBuffer, cullMode VkCullModeFlags) {
	syscall.Syscall(procvkCmdSetCullModeEXT.Addr(), 2, uintptr(commandBuffer), uintptr(cullMode), 0)	
	return
}


func vkCmdSetFrontFaceEXT(commandBuffer VkCommandBuffer, frontFace VkFrontFace) {
	syscall.Syscall(procvkCmdSetFrontFaceEXT.Addr(), 2, uintptr(commandBuffer), uintptr(frontFace), 0)	
	return
}


func vkCmdSetPrimitiveTopologyEXT(commandBuffer VkCommandBuffer, primitiveTopology VkPrimitiveTopology) {
	syscall.Syscall(procvkCmdSetPrimitiveTopologyEXT.Addr(), 2, uintptr(commandBuffer), uintptr(primitiveTopology), 0)	
	return
}


func vkCmdSetViewportWithCountEXT(commandBuffer VkCommandBuffer, viewportCount uint32, pViewports *VkViewport) {
	syscall.Syscall(procvkCmdSetViewportWithCountEXT.Addr(), 3, uintptr(commandBuffer), uintptr(viewportCount), uintptr(unsafe.Pointer(pViewports)))	
	return
}


func vkCmdSetScissorWithCountEXT(commandBuffer VkCommandBuffer, scissorCount uint32, pScissors *VkRect2D) {
	syscall.Syscall(procvkCmdSetScissorWithCountEXT.Addr(), 3, uintptr(commandBuffer), uintptr(scissorCount), uintptr(unsafe.Pointer(pScissors)))	
	return
}


func vkCmdBindVertexBuffers2EXT(commandBuffer VkCommandBuffer, firstBinding uint32, bindingCount uint32, pBuffers *VkBuffer, pOffsets *VkDeviceSize, pSizes *VkDeviceSize, pStrides *VkDeviceSize) {
	syscall.Syscall9(procvkCmdBindVertexBuffers2EXT.Addr(), 7, uintptr(commandBuffer), uintptr(firstBinding), uintptr(bindingCount), uintptr(unsafe.Pointer(pBuffers)), uintptr(unsafe.Pointer(pOffsets)), uintptr(unsafe.Pointer(pSizes)), uintptr(unsafe.Pointer(pStrides)), 0, 0)	
	return
}


func vkCmdSetDepthTestEnableEXT(commandBuffer VkCommandBuffer, depthTestEnable VkBool32) {
	syscall.Syscall(procvkCmdSetDepthTestEnableEXT.Addr(), 2, uintptr(commandBuffer), uintptr(depthTestEnable), 0)	
	return
}


func vkCmdSetDepthWriteEnableEXT(commandBuffer VkCommandBuffer, depthWriteEnable VkBool32) {
	syscall.Syscall(procvkCmdSetDepthWriteEnableEXT.Addr(), 2, uintptr(commandBuffer), uintptr(depthWriteEnable), 0)	
	return
}


func vkCmdSetDepthCompareOpEXT(commandBuffer VkCommandBuffer, depthCompareOp VkCompareOp) {
	syscall.Syscall(procvkCmdSetDepthCompareOpEXT.Addr(), 2, uintptr(commandBuffer), uintptr(depthCompareOp), 0)	
	return
}


func vkCmdSetDepthBoundsTestEnableEXT(commandBuffer VkCommandBuffer, depthBoundsTestEnable VkBool32) {
	syscall.Syscall(procvkCmdSetDepthBoundsTestEnableEXT.Addr(), 2, uintptr(commandBuffer), uintptr(depthBoundsTestEnable), 0)	
	return
}


func vkCmdSetStencilTestEnableEXT(commandBuffer VkCommandBuffer, stencilTestEnable VkBool32) {
	syscall.Syscall(procvkCmdSetStencilTestEnableEXT.Addr(), 2, uintptr(commandBuffer), uintptr(stencilTestEnable), 0)	
	return
}


func vkCmdSetStencilOpEXT(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, failOp VkStencilOp, passOp VkStencilOp, depthFailOp VkStencilOp, compareOp VkCompareOp) {
	syscall.Syscall6(procvkCmdSetStencilOpEXT.Addr(), 6, uintptr(commandBuffer), uintptr(faceMask), uintptr(failOp), uintptr(passOp), uintptr(depthFailOp), uintptr(compareOp))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkCreatePrivateDataSlotEXT(device VkDevice, pCreateInfo *VkPrivateDataSlotCreateInfoEXT, pAllocator *VkAllocationCallbacks, pPrivateDataSlot *VkPrivateDataSlotEXT) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkCreatePrivateDataSlotEXT.Addr(), 4, uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pPrivateDataSlot)), 0, 0)	
	ret0 = VkResult(r0)
	return
}


func vkDestroyPrivateDataSlotEXT(device VkDevice, privateDataSlot VkPrivateDataSlotEXT, pAllocator *VkAllocationCallbacks) {
	syscall.Syscall(procvkDestroyPrivateDataSlotEXT.Addr(), 3, uintptr(device), uintptr(privateDataSlot), uintptr(unsafe.Pointer(pAllocator)))	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS
func vkSetPrivateDataEXT(device VkDevice, objectType VkObjectType, objectHandle uint64, privateDataSlot VkPrivateDataSlotEXT, data uint64) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall6(procvkSetPrivateDataEXT.Addr(), 5, uintptr(device), uintptr(objectType), uintptr(objectHandle), uintptr(privateDataSlot), uintptr(data), 0)	
	ret0 = VkResult(r0)
	return
}


func vkGetPrivateDataEXT(device VkDevice, objectType VkObjectType, objectHandle uint64, privateDataSlot VkPrivateDataSlotEXT, pData *uint64) {
	syscall.Syscall6(procvkGetPrivateDataEXT.Addr(), 5, uintptr(device), uintptr(objectType), uintptr(objectHandle), uintptr(privateDataSlot), uintptr(unsafe.Pointer(pData)), 0)	
	return
}


func vkCmdCopyBuffer2KHR(commandBuffer VkCommandBuffer, pCopyBufferInfo *VkCopyBufferInfo2KHR) {
	syscall.Syscall(procvkCmdCopyBuffer2KHR.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pCopyBufferInfo)), 0)	
	return
}


func vkCmdCopyImage2KHR(commandBuffer VkCommandBuffer, pCopyImageInfo *VkCopyImageInfo2KHR) {
	syscall.Syscall(procvkCmdCopyImage2KHR.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pCopyImageInfo)), 0)	
	return
}


func vkCmdBlitImage2KHR(commandBuffer VkCommandBuffer, pBlitImageInfo *VkBlitImageInfo2KHR) {
	syscall.Syscall(procvkCmdBlitImage2KHR.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pBlitImageInfo)), 0)	
	return
}


func vkCmdCopyBufferToImage2KHR(commandBuffer VkCommandBuffer, pCopyBufferToImageInfo *VkCopyBufferToImageInfo2KHR) {
	syscall.Syscall(procvkCmdCopyBufferToImage2KHR.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pCopyBufferToImageInfo)), 0)	
	return
}


func vkCmdCopyImageToBuffer2KHR(commandBuffer VkCommandBuffer, pCopyImageToBufferInfo *VkCopyImageToBufferInfo2KHR) {
	syscall.Syscall(procvkCmdCopyImageToBuffer2KHR.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pCopyImageToBufferInfo)), 0)	
	return
}


func vkCmdResolveImage2KHR(commandBuffer VkCommandBuffer, pResolveImageInfo *VkResolveImageInfo2KHR) {
	syscall.Syscall(procvkCmdResolveImage2KHR.Addr(), 2, uintptr(commandBuffer), uintptr(unsafe.Pointer(pResolveImageInfo)), 0)	
	return
}


func vkCmdSetFragmentShadingRateKHR(commandBuffer VkCommandBuffer, pFragmentSize *VkExtent2D, combinerOps2 VkFragmentShadingRateCombinerOpKHR, combinerOps1 VkFragmentShadingRateCombinerOpKHR) {
	syscall.Syscall6(procvkCmdSetFragmentShadingRateKHR.Addr(), 4, uintptr(commandBuffer), uintptr(unsafe.Pointer(pFragmentSize)), uintptr(combinerOps2), uintptr(combinerOps1), 0, 0)	
	return
}

//Errors: VK_ERROR_OUT_OF_HOST_MEMORY
//On Success: VK_SUCCESS, VK_INCOMPLETE
func vkGetPhysicalDeviceFragmentShadingRatesKHR(physicalDevice VkPhysicalDevice, pFragmentShadingRateCount *uint32, pFragmentShadingRates *VkPhysicalDeviceFragmentShadingRateKHR) (ret0 VkResult) {
	r0, _, _ := syscall.Syscall(procvkGetPhysicalDeviceFragmentShadingRatesKHR.Addr(), 3, uintptr(physicalDevice), uintptr(unsafe.Pointer(pFragmentShadingRateCount)), uintptr(unsafe.Pointer(pFragmentShadingRates)))	
	ret0 = VkResult(r0)
	return
}

