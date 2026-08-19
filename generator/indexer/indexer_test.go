package indexer_test

import (
	"testing"

	"github.com/KCoen/vk/generator/indexer"
	"github.com/KCoen/vk/generator/parser"
)

func TestIndexVkXML(t *testing.T) {
	reg, err := parser.ParseFile("../../Vulkan-Docs/xml/vk.xml")
	if err != nil {
		t.Fatalf("Failed to parse vk.xml: %v", err)
	}

	idx, err := indexer.BuildIndex(reg, nil)
	if err != nil {
		t.Fatalf("Failed to build index: %v", err)
	}

	t.Logf("Indexed %d platforms", len(idx.Platforms))
	t.Logf("Indexed %d tags", len(idx.Tags))
	t.Logf("Indexed %d constants", len(idx.Constants))
	t.Logf("Indexed %d handles", len(idx.Handles))
	t.Logf("Indexed %d enum groups", len(idx.EnumGroups))
	t.Logf("Indexed %d bitmasks", len(idx.Bitmasks))
	t.Logf("Indexed %d structs/unions", len(idx.Structs))
	t.Logf("Indexed %d commands", len(idx.Commands))
	t.Logf("Indexed %d versions", len(idx.Versions))
	t.Logf("Indexed %d extensions", len(idx.Extensions))

	if len(idx.Handles) == 0 {
		t.Errorf("Expected handles to be indexed, got 0")
	}
	if len(idx.Structs) == 0 {
		t.Errorf("Expected structs to be indexed, got 0")
	}
	if len(idx.Commands) == 0 {
		t.Errorf("Expected commands to be indexed, got 0")
	}

	// Verify VkInstanceCreateInfo
	info, ok := idx.Structs["VkInstanceCreateInfo"]
	if !ok {
		t.Fatalf("Expected VkInstanceCreateInfo in structs")
	}
	if !info.HasSType || info.GoSTypeDefault != "STRUCTURE_TYPE_INSTANCE_CREATE_INFO" {
		t.Errorf("Expected SType default STRUCTURE_TYPE_INSTANCE_CREATE_INFO, got %s", info.GoSTypeDefault)
	}

	// Check string slices in InstanceCreateInfo
	var hasExtSlice, hasLayerSlice bool
	for _, m := range info.Members {
		if m.Name == "ppEnabledExtensionNames" && m.IsStringSlice {
			hasExtSlice = true
		}
		if m.Name == "ppEnabledLayerNames" && m.IsStringSlice {
			hasLayerSlice = true
		}
	}
	if !hasExtSlice {
		t.Errorf("Expected ppEnabledExtensionNames to be string slice")
	}
	if !hasLayerSlice {
		t.Errorf("Expected ppEnabledLayerNames to be string slice")
	}
}
