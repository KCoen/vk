package indexer

import (
	"strings"
)

// MapCTypeToGo converts a C Vulkan type (e.g. "uint32_t", "VkInstance", "VkResult") to its Go type.
func MapCTypeToGo(cType string) string {
	clean := strings.TrimSpace(cType)

	// Strip "const " and "struct "
	clean = strings.TrimPrefix(clean, "const ")
	clean = strings.TrimPrefix(clean, "struct ")
	clean = strings.TrimSpace(clean)

	if clean == "" || clean == "void" {
		return ""
	}

	// Count pointers
	pointerStars := ""
	for strings.HasSuffix(clean, "*") {
		pointerStars += "*"
		clean = strings.TrimSuffix(clean, "*")
		clean = strings.TrimSpace(clean)
		clean = strings.TrimPrefix(clean, "const ")
		clean = strings.TrimSpace(clean)
	}

	// Platform types mapping
	switch clean {
	case "void":
		if pointerStars == "*" {
			return "unsafe.Pointer"
		}
		if pointerStars == "**" {
			return "*unsafe.Pointer"
		}
		return "unsafe.Pointer"
	case "char":
		if pointerStars == "*" {
			return "*byte"
		}
		if pointerStars == "**" {
			return "**byte"
		}
		return "byte"
	case "uint8_t":
		return pointerStars + "uint8"
	case "uint16_t":
		return pointerStars + "uint16"
	case "uint32_t":
		return pointerStars + "uint32"
	case "uint64_t":
		return pointerStars + "uint64"
	case "int8_t":
		return pointerStars + "int8"
	case "int16_t":
		return pointerStars + "int16"
	case "int32_t", "int":
		return pointerStars + "int32"
	case "int64_t":
		return pointerStars + "int64"
	case "float":
		return pointerStars + "float32"
	case "double":
		return pointerStars + "float64"
	case "size_t":
		return pointerStars + "uintptr"
	// Platform types
	case "Display", "Window", "VisualID", "RROutput",
		"xcb_connection_t", "xcb_window_t", "xcb_visualid_t",
		"wl_display", "wl_surface",
		"HINSTANCE", "HWND", "HMONITOR", "HANDLE", "SECURITY_ATTRIBUTES", "LPCWSTR",
		"ANativeWindow", "AHardwareBuffer", "OH_NativeBuffer", "OHBufferHandle", "OHNativeWindow",
		"IDirectFB", "IDirectFBSurface", "ubm_device", "ubm_surface",
		"CAMetalLayer", "MTLDevice_id", "MTLCommandQueue_id", "MTLBuffer_id", "MTLTexture_id", "MTLSharedEvent_id", "IOSurfaceRef",
		"NvSciSyncAttrList", "NvSciSyncObj", "NvSciSyncFence", "NvSciBufAttrList", "NvSciBufObj",
		"_screen_context", "_screen_window", "_screen_buffer":
		if pointerStars == "**" {
			return "*uintptr"
		}
		return "uintptr"
	case "DWORD", "zx_handle_t", "GgpStreamDescriptor", "GgpFrameToken":
		return pointerStars + "uint32"
	}

	// Video header types (StdVideo...)
	if strings.HasPrefix(clean, "StdVideo") {
		if pointerStars == "*" {
			return "unsafe.Pointer"
		}
		if pointerStars == "**" {
			return "*unsafe.Pointer"
		}
		return "uintptr"
	}

	// Check if funcpointer
	if strings.HasPrefix(clean, "PFN_") {
		return clean
	}

	return pointerStars + CleanTypeName(clean)
}

// ParseFullCType analyzes a member/param inner text to determine base type, pointer levels, arrays, const.
func ParseFullCType(innerXML, baseType, name string) (cType string, isPointer bool, isDoublePointer bool, isConst bool, arrayDims []string) {
	text := innerXML
	for {
		start := strings.Index(text, "<comment>")
		if start == -1 {
			break
		}
		end := strings.Index(text[start:], "</comment>")
		if end == -1 {
			break
		}
		text = text[:start] + text[start+end+10:]
	}

	var clean strings.Builder
	inTag := false
	for _, r := range text {
		if r == '<' {
			inTag = true
			continue
		}
		if r == '>' {
			inTag = false
			continue
		}
		if !inTag {
			clean.WriteRune(r)
		}
	}

	raw := clean.String()
	raw = strings.TrimSpace(raw)

	isConst = strings.Contains(raw, "const")
	pointerCount := strings.Count(raw, "*")
	if pointerCount >= 2 {
		isDoublePointer = true
		isPointer = true
	} else if pointerCount == 1 {
		isPointer = true
	}

	if idx := strings.Index(raw, "["); idx != -1 {
		rest := raw[idx:]
		for {
			openIdx := strings.Index(rest, "[")
			if openIdx == -1 {
				break
			}
			closeIdx := strings.Index(rest[openIdx:], "]")
			if closeIdx == -1 {
				break
			}
			dim := strings.TrimSpace(rest[openIdx+1 : openIdx+closeIdx])
			if dim != "" {
				arrayDims = append(arrayDims, dim)
			}
			rest = rest[openIdx+closeIdx+1:]
		}
	}

	cType = baseType
	if isDoublePointer {
		cType += "**"
	} else if isPointer {
		cType += "*"
	}

	return cType, isPointer, isDoublePointer, isConst, arrayDims
}
