package tests

import (
	"testing"

	"github.com/KCoen/vk/extensions/khr_swapchain"
	"github.com/KCoen/vk/vulkan"
	"github.com/KCoen/vk/vulkanbase"
	"github.com/KCoen/vk/vulkansc"
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

	cmds := vulkan.InitCommands(0, 0)
	if cmds == nil {
		t.Fatal("expected non-nil *vulkan.Commands from InitCommands")
	}

	extCmds := khr_swapchain.Init(0, 0)
	if extCmds == nil {
		t.Fatal("expected non-nil *khr_swapchain.Commands from extension Init")
	}

	// Create test struct
	info := vulkan.NewInstanceCreateInfo()
	if info.SType != vulkan.STRUCTURE_TYPE_INSTANCE_CREATE_INFO {
		t.Errorf("expected SType STRUCTURE_TYPE_INSTANCE_CREATE_INFO, got %v", info.SType)
	}
}
