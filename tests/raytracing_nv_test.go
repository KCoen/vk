package tests

import (
	"os"
	"strings"
	"testing"
)

func TestCreateRayTracingPipelinesNVDefinition(t *testing.T) {
	packages := []struct {
		name                   string
		path                   string
		expectErrorPoolMemory bool
	}{
		{
			name:                   "vulkan",
			path:                   "../vulkan/commands_gen.go",
			expectErrorPoolMemory: false,
		},
		{
			name:                   "vulkanbase",
			path:                   "../vulkanbase/commands_gen.go",
			expectErrorPoolMemory: false,
		},
		{
			name:                   "vulkansc",
			path:                   "../vulkansc/commands_gen.go",
			expectErrorPoolMemory: true,
		},
	}

	for _, pkg := range packages {
		t.Run(pkg.name, func(t *testing.T) {
			content, err := os.ReadFile(pkg.path)
			if err != nil {
				t.Fatalf("Failed to read %s: %v", pkg.path, err)
			}

			strContent := string(content)

			// 1. Verify definition exists
			funcDef := "func CreateRayTracingPipelinesNV("
			if !strings.Contains(strContent, funcDef) {
				t.Fatalf("Expected definition of %s in %s, but it was missing", funcDef, pkg.path)
			}

			// 2. Extract doc comment block preceding func CreateRayTracingPipelinesNV
			idx := strings.Index(strContent, funcDef)
			commentBlock := strContent[:idx]
			lastCommentIdx := strings.LastIndex(commentBlock, "// CreateRayTracingPipelinesNV")
			if lastCommentIdx == -1 {
				t.Fatalf("Could not find doc comment header for CreateRayTracingPipelinesNV in %s", pkg.path)
			}
			cmdDoc := commentBlock[lastCommentIdx:]

			// 3. Verify Error codes reported in the doc comment
			hasPoolMemory := strings.Contains(cmdDoc, "VK_ERROR_OUT_OF_POOL_MEMORY")
			if pkg.expectErrorPoolMemory && !hasPoolMemory {
				t.Errorf("%s doc comment should report VK_ERROR_OUT_OF_POOL_MEMORY, but it did not:\n%s", pkg.name, cmdDoc)
			} else if !pkg.expectErrorPoolMemory && hasPoolMemory {
				t.Errorf("%s doc comment should NOT report VK_ERROR_OUT_OF_POOL_MEMORY, but it did:\n%s", pkg.name, cmdDoc)
			}
		})
	}
}

func TestEnumerateDeviceLayerPropertiesErrorCodes(t *testing.T) {
	packages := []struct {
		name                 string
		path                 string
		expectErrorHostMemory bool
	}{
		{
			name:                 "vulkan",
			path:                 "../vulkan/commands_gen.go",
			expectErrorHostMemory: true,
		},
		{
			name:                 "vulkanbase",
			path:                 "../vulkanbase/commands_gen.go",
			expectErrorHostMemory: true,
		},
		{
			name:                 "vulkansc",
			path:                 "../vulkansc/commands_gen.go",
			expectErrorHostMemory: false,
		},
	}

	for _, pkg := range packages {
		t.Run(pkg.name, func(t *testing.T) {
			content, err := os.ReadFile(pkg.path)
			if err != nil {
				t.Fatalf("Failed to read %s: %v", pkg.path, err)
			}

			strContent := string(content)
			funcDef := "func EnumerateDeviceLayerProperties("
			if !strings.Contains(strContent, funcDef) {
				t.Fatalf("Expected definition of %s in %s", funcDef, pkg.path)
			}

			idx := strings.Index(strContent, funcDef)
			commentBlock := strContent[:idx]
			lastCommentIdx := strings.LastIndex(commentBlock, "// EnumerateDeviceLayerProperties")
			if lastCommentIdx == -1 {
				t.Fatalf("Could not find doc comment header for EnumerateDeviceLayerProperties in %s", pkg.path)
			}
			cmdDoc := commentBlock[lastCommentIdx:]

			hasOutOfHostMemory := strings.Contains(cmdDoc, "VK_ERROR_OUT_OF_HOST_MEMORY")
			if pkg.expectErrorHostMemory && !hasOutOfHostMemory {
				t.Errorf("%s EnumerateDeviceLayerProperties should report VK_ERROR_OUT_OF_HOST_MEMORY, but it did not:\n%s", pkg.name, cmdDoc)
			} else if !pkg.expectErrorHostMemory && hasOutOfHostMemory {
				t.Errorf("%s EnumerateDeviceLayerProperties should NOT report VK_ERROR_OUT_OF_HOST_MEMORY, but it did:\n%s", pkg.name, cmdDoc)
			}
		})
	}
}
