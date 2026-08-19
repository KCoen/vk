package docparser_test

import (
	"testing"

	"github.com/KCoen/vk/generator/docparser"
)

func TestParseVulkanDocs(t *testing.T) {
	docIdx, err := docparser.ParseVulkanDocs("../../Vulkan-Docs")
	if err != nil {
		t.Fatalf("failed to parse Vulkan-Docs: %v", err)
	}

	if len(docIdx.Refpages) == 0 {
		t.Fatalf("expected refpages to be parsed, got 0")
	}

	createInstance := docIdx.Find("vkCreateInstance")
	if createInstance == nil {
		t.Fatalf("expected vkCreateInstance refpage")
	}
	if createInstance.ShortDesc == "" {
		t.Errorf("expected non-empty ShortDesc for vkCreateInstance")
	}
	if len(createInstance.Params) == 0 {
		t.Errorf("expected params for vkCreateInstance")
	}

	instanceInfo := docIdx.Find("VkInstanceCreateInfo")
	if instanceInfo == nil {
		t.Fatalf("expected VkInstanceCreateInfo refpage")
	}
	if len(instanceInfo.Members) == 0 {
		t.Errorf("expected members for VkInstanceCreateInfo")
	}

	t.Logf("Successfully indexed %d refpages from Vulkan-Docs", len(docIdx.Refpages))
	t.Logf("vkCreateInstance short desc: %s", createInstance.ShortDesc)
	t.Logf("vkCreateInstance pCreateInfo param doc: %s", createInstance.Params["pCreateInfo"])
}
