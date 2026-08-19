package indexer

import (
	"strings"
	"unicode"
)

// CleanTypeName strips "Vk" or "VK_" prefix from type names.
func CleanTypeName(name string) string {
	if strings.HasPrefix(name, "Vk") {
		return name[2:]
	}
	if strings.HasPrefix(name, "VK_") {
		return name[3:]
	}
	return name
}

// CleanCmdName strips "vk" or "vkCmd" prefix from command names.
func CleanCmdName(name string) string {
	if strings.HasPrefix(name, "vk") {
		return name[2:]
	}
	return name
}

// CleanEnumName strips "VK_" prefix from enum values.
func CleanEnumName(name string) string {
	if strings.HasPrefix(name, "VK_") {
		return name[3:]
	}
	return name
}

// CleanConstantName strips "VK_" prefix from constant names.
func CleanConstantName(name string) string {
	if strings.HasPrefix(name, "VK_") {
		return name[3:]
	}
	return name
}

// CleanExtPkgName converts "VK_KHR_swapchain" to "khr_swapchain".
func CleanExtPkgName(name string) string {
	lower := strings.ToLower(name)
	if strings.HasPrefix(lower, "vk_") {
		return lower[3:]
	}
	return lower
}

// CleanExtGoName converts "VK_KHR_swapchain" to "KhrSwapchain".
func CleanExtGoName(name string) string {
	parts := strings.Split(name, "_")
	var result strings.Builder
	for _, p := range parts {
		if strings.EqualFold(p, "VK") {
			continue
		}
		if len(p) > 0 {
			result.WriteString(strings.ToUpper(p[:1]) + strings.ToLower(p[1:]))
		}
	}
	return result.String()
}

// CleanMemberName converts C struct member name (e.g. pApplicationInfo, sType, ppEnabledExtensionNames)
// to an exported Go field name (e.g. ApplicationInfo, SType, EnabledExtensionNames).
func CleanMemberName(name string) string {
	if name == "" {
		return ""
	}

	// Special cases
	if name == "sType" {
		return "SType"
	}
	if name == "pNext" {
		return "Next"
	}

	trimmed := name
	// Strip pointer prefixes: p, pp, pfn
	if strings.HasPrefix(trimmed, "pfn") && len(trimmed) > 3 && unicode.IsUpper(rune(trimmed[3])) {
		trimmed = trimmed[3:]
	} else if strings.HasPrefix(trimmed, "pp") && len(trimmed) > 2 && unicode.IsUpper(rune(trimmed[2])) {
		trimmed = trimmed[2:]
	} else if strings.HasPrefix(trimmed, "p") && len(trimmed) > 1 && unicode.IsUpper(rune(trimmed[1])) {
		trimmed = trimmed[1:]
	}

	// Capitalize first letter
	runes := []rune(trimmed)
	runes[0] = unicode.ToUpper(runes[0])
	res := string(runes)

	// Capitalize common acronyms
	res = fixCommonAcronyms(res)
	return res
}

// CleanParamName converts C command param name (e.g. pCreateInfo, pPipelines) to Go param name.
func CleanParamName(name string) string {
	if name == "" {
		return ""
	}

	trimmed := name
	if strings.HasPrefix(trimmed, "pfn") && len(trimmed) > 3 && unicode.IsUpper(rune(trimmed[3])) {
		trimmed = trimmed[3:]
	} else if strings.HasPrefix(trimmed, "pp") && len(trimmed) > 2 && unicode.IsUpper(rune(trimmed[2])) {
		trimmed = trimmed[2:]
	} else if strings.HasPrefix(trimmed, "p") && len(trimmed) > 1 && unicode.IsUpper(rune(trimmed[1])) {
		trimmed = trimmed[1:]
	}

	runes := []rune(trimmed)
	runes[0] = unicode.ToLower(runes[0])
	res := string(runes)

	// Avoid Go keywords
	switch res {
	case "type":
		return "typ"
	case "select":
		return "sel"
	case "map":
		return "mapping"
	case "range":
		return "rng"
	case "default":
		return "def"
	case "func":
		return "fn"
	case "package":
		return "pkg"
	case "var":
		return "v"
	case "const":
		return "c"
	}

	return res
}

func fixCommonAcronyms(name string) string {
	acronyms := []string{"UUID", "RGBA", "BGRA", "RGB", "BGR", "ID", "2D", "3D", "1D", "URL", "URI", "CPU", "GPU", "API", "OS", "FD"}
	for _, acr := range acronyms {
		// Look for lowercase/mixed case versions like "Uuid" -> "UUID", "Id" -> "ID", "Api" -> "API"
		title := strings.ToUpper(acr[:1]) + strings.ToLower(acr[1:])
		if strings.HasSuffix(name, title) {
			name = name[:len(name)-len(title)] + acr
		}
	}
	return name
}
