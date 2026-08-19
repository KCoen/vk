package emitter

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"

	"go.cld.moe/vk_google/generator/indexer"
)

// GenerateAll generates all Go packages from the indexed Vulkan registry and writes them to outDir.
func GenerateAll(idx *indexer.Index, outDir string) error {
	// 1. Remove legacy vk directory if present
	_ = os.RemoveAll(filepath.Join(outDir, "vk"))

	// 2. Generate API branch packages (vulkan, vulkanbase, vulkansc)
	for _, branch := range idx.ApiBranches {
		branchDir := filepath.Join(outDir, branch.PkgName)
		if err := os.MkdirAll(branchDir, 0755); err != nil {
			return fmt.Errorf("failed to create API branch directory %s: %w", branchDir, err)
		}
		_ = os.Remove(filepath.Join(branchDir, "commands.go"))
		_ = os.Remove(filepath.Join(branchDir, "doc.go"))

		bFiles, err := EmitApiBranchPackage(branch, idx)
		if err != nil {
			return fmt.Errorf("failed to emit API branch package %s: %w", branch.PkgName, err)
		}
		for filename, content := range bFiles {
			if err := writeFile(filepath.Join(branchDir, filename), content); err != nil {
				return err
			}
		}
	}

	// 3. Generate extension packages
	extensionsDir := filepath.Join(outDir, "extensions")
	if err := os.MkdirAll(extensionsDir, 0755); err != nil {
		return fmt.Errorf("failed to create extensions directory: %w", err)
	}

	for _, ext := range idx.Extensions {
		// Only generate supported extensions for vulkan
		if !strings.Contains(ext.Supported, "vulkan") && ext.Supported != "" {
			continue
		}

		extDir := filepath.Join(extensionsDir, ext.PkgName)
		if err := os.MkdirAll(extDir, 0755); err != nil {
			return fmt.Errorf("failed to create extension directory %s: %w", extDir, err)
		}
		_ = os.Remove(filepath.Join(extDir, "doc.go"))
		_ = os.Remove(filepath.Join(extDir, "extension.go"))
		_ = os.Remove(filepath.Join(extDir, "commands.go"))

		extFiles, err := EmitExtensionPackage(ext, idx)
		if err != nil {
			return fmt.Errorf("failed to emit extension package %s: %w", ext.Name, err)
		}
		for filename, content := range extFiles {
			if err := writeFile(filepath.Join(extDir, filename), content); err != nil {
				return err
			}
		}
	}

	return nil
}

func isValidBranchType(typeName string, branch *indexer.ApiBranchInfo, idx *indexer.Index) bool {
	if typeName == "" || typeName == "void" || typeName == "char" || typeName == "int" || typeName == "float" || typeName == "double" {
		return true
	}
	t := strings.TrimPrefix(typeName, "[]")
	t = strings.TrimPrefix(t, "*")
	t = strings.TrimPrefix(t, "*")
	switch t {
	case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "byte", "string", "uintptr", "unsafe.Pointer", "any", "Bool32", "DeviceSize", "DeviceAddress", "Flags", "Flags64", "SampleMask":
		return true
	}
	vkName := t
	if strings.HasPrefix(t, "Raw") {
		vkName = "Vk" + strings.TrimPrefix(t, "Raw")
	}
	if !strings.HasPrefix(vkName, "Vk") && !strings.HasPrefix(vkName, "PFN_vk") {
		vkName = "Vk" + vkName
	}
	if _, ok := branch.Structs[vkName]; ok {
		return true
	}
	if _, ok := branch.Handles[vkName]; ok {
		return true
	}
	if _, ok := branch.Bitmasks[vkName]; ok {
		return true
	}
	if _, ok := branch.EnumGroups[vkName]; ok {
		return true
	}
	if _, ok := branch.TypeAliases[vkName]; ok {
		return true
	}
	if _, ok := branch.FuncPointerDefs[vkName]; ok {
		return true
	}
	if _, ok := branch.Structs[t]; ok {
		return true
	}
	if _, ok := branch.Handles[t]; ok {
		return true
	}
	if _, ok := branch.Bitmasks[t]; ok {
		return true
	}
	if _, ok := branch.EnumGroups[t]; ok {
		return true
	}
	if _, ok := branch.TypeAliases[t]; ok {
		return true
	}
	if _, ok := branch.FuncPointerDefs[t]; ok {
		return true
	}
	if _, ok := idx.BaseTypes[vkName]; ok {
		return true
	}
	if _, ok := idx.BaseTypes[t]; ok {
		return true
	}
	return false
}

func writeFile(path string, content string) error {
	formatted, err := format.Source([]byte(content))
	if err != nil {
		// Write raw content to aid debugging if formatting fails
		_ = os.WriteFile(path, []byte(content), 0644)
		return fmt.Errorf("go/format error on %s: %w", path, err)
	}
	return os.WriteFile(path, formatted, 0644)
}
