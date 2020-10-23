// +build test
package vk

import (
	"bytes"
	"testing"
)

func VK_MAKE_VERSION(major, minor, patch int) uint32 {
    return ((((uint32)(major)) << 22) | (((uint32)(minor)) << 12) | ((uint32)(patch)))
}

func CString(b []byte) []byte {
	return b[:bytes.IndexByte(b, 0)]
}

func TestHelloWorld(tst *testing.T) {
	appInfo := VkApplicationInfo{};
	appInfo.sType = VK_STRUCTURE_TYPE_APPLICATION_INFO;
	appInfo.pApplicationName = &([]byte("Hello Triangle")[0]);
	appInfo.applicationVersion = VK_MAKE_VERSION(1, 0, 0);
	appInfo.pEngineName = &([]byte("None")[0])
	appInfo.engineVersion = VK_MAKE_VERSION(1, 0, 0);
	appInfo.apiVersion = VK_MAKE_VERSION(1, 0, 0);

	createInfo := VkInstanceCreateInfo{};
	createInfo.sType = VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO;
	createInfo.pApplicationInfo = &appInfo;
	createInfo.enabledLayerCount = 0;

	var extensionCount uint32 = 0;
	vkEnumerateInstanceExtensionProperties(nil, &extensionCount, nil);
	tst.Logf("Num extensions: %v", extensionCount)
	ext := make([]VkExtensionProperties,extensionCount)
	vkEnumerateInstanceExtensionProperties(nil, &extensionCount, &ext[0]);
	
	for i := uint32(0);i < extensionCount;i++ {
		tst.Logf("Ext[%d] \tname: %s",i, CString(ext[i].extensionName[:]))	
	}	
	var instance VkInstance
	if e := vkCreateInstance(&createInfo, nil, &instance);e != VK_SUCCESS {
		tst.Error(e)
	}
	tst.Logf("%v", instance)
	vkDestroyInstance(instance, nil);
	tst.Logf("%v", instance)
}