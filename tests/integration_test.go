package tests

import (
	"testing"

	"go.cld.moe/vk_google/extensions/khr_swapchain"
	"go.cld.moe/vk_google/vulkan"
	"go.cld.moe/vk_google/vulkanbase"
	"go.cld.moe/vk_google/vulkansc"
)

func TestApiBranchesAndExtensions(t *testing.T) {
	// Test Extension constants
	if khr_swapchain.ExtensionName != "VK_KHR_swapchain" {
		t.Errorf("expected VK_KHR_swapchain, got %s", khr_swapchain.ExtensionName)
	}
	if khr_swapchain.ExtensionNumber != 2 {
		t.Errorf("expected extension number 2, got %d", khr_swapchain.ExtensionNumber)
	}

	// Verify API branches exports
	_ = vulkan.Init
	_ = vulkanbase.Init
	_ = vulkansc.Init
	_ = khr_swapchain.Init

	// Create test struct
	info := vulkan.NewInstanceCreateInfo()
	if info.SType != vulkan.STRUCTURE_TYPE_INSTANCE_CREATE_INFO {
		t.Errorf("expected SType STRUCTURE_TYPE_INSTANCE_CREATE_INFO, got %v", info.SType)
	}
}
