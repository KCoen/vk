package main

import (
	"strings"
	"testing"

	"go.cld.moe/vk/generator/util"
)

// Honestly there might just be a cgo api for parsing a fully fledged C definition?
func TestTypeDecode(tst *testing.T) {
	SimpleTest := func(s string) {
		c := parseTypeFromXmlString(s)
		if c.Comment != "" && strings.Count(c.Fulltype, c.Comment) > 0 {
			tst.Errorf("Failed %s Fulltype %s contained Comment %s", s, c.Fulltype, c.Comment)
		}
		tst.Log(util.Vardump(c))
	}

	SimpleTest(`<type category="enum" name="VkDriverIdKHR"                                 alias="VkDriverId"/>`)
	SimpleTest(`<member values="VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES"><type>VkStructureType</type> <name>sType</name></member>`)
	// Expect str0 = VkStructureType
	SimpleTest(`<member><type>char</type>                             <name>driverInfo</name>[<enum>VK_MAX_DRIVER_INFO_SIZE</enum>]</member>`)
	// Expect str1 = char[VK_MAX_DRIVER_INFO_SIZE]
	SimpleTest(`<type category="handle" parent="VkDevice"><type>VK_DEFINE_NON_DISPATCHABLE_HANDLE</type>(<name>VkCommandPool</name>)</type>`)
	// Expect str2 = VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkCommandPool)
	// Str2 Return Type is VK_DEFINE_NON_DISPATCHABLE_HANDLE VkCommandPool
	SimpleTest(`<type category="funcpointer">typedef void* (VKAPI_PTR *<name>PFN_vkReallocationFunction</name>)(
    <type>void</type>*                                       pOriginal,
    <type>void</type>*                                       pUserData,
    <type>size_t</type>                                      size,
    <type>size_t</type>                                      alignment,
    <type>VkSystemAllocationScope</type>                     allocationScope);</type>`)
	// Expect str3 = typedef void* (VKAPI_PTR *PFN_vkReallocationFunction)(void* pUserData, void* pOriginal, size_t size, size_t alignment, VkSystemAllocationScope allocationscope);
	// Str3 Return a struct that can identify arguments (to seperate name from type)
	SimpleTest(`<member><type>VkBool32</type>               <name>robustBufferAccess</name><comment>out of bounds buffer accesses are well defined</comment></member>`)
	SimpleTest(`<type requires="VkBufferCreateFlagBits"           category="bitmask">typedef <type>VkFlags</type> <name>VkBufferCreateFlags</name>;</type>`)
	SimpleTest(`<member noautovalidity="true"><type>VkImageLayout</type>   <name>imageLayout</name><comment>Layout the image is expected to be in when accessed using this descriptor (only used if imageView is not VK_NULL_HANDLE).</comment></member>`)
	SimpleTest(`<member><type>VkDeviceSize</type>           <name>maxResourceSize</name><comment>max size (in bytes) of this resource type</comment></member>`)
	SimpleTest(`<member>struct <type>VkBaseOutStructure</type>* <name>pNext</name></member>`)
	SimpleTest(`<param>const <type>VkInstanceCreateInfo</type>* <name>pCreateInfo</name></param>`)
	SimpleTest(`<type requires="VkShaderStageFlagBits"            category="bitmask">typedef <type>VkFlags</type> <name>VkShaderStageFlags</name>;</type>`)
}
