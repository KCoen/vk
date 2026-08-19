package emitter

import (
	"fmt"
	"strings"

	"go.cld.moe/vk_google/generator/indexer"
)

// FormatDocComment formats a documentation comment into Go standard doc comments (// ...).
func FormatDocComment(leadText string, comments ...string) string {
	var lines []string

	if leadText != "" {
		lines = append(lines, formatLines(leadText)...)
	}

	for _, c := range comments {
		if c == "" {
			continue
		}
		lines = append(lines, formatLines(c)...)
	}

	if len(lines) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, l := range lines {
		if l == "" {
			sb.WriteString("//\n")
		} else {
			sb.WriteString("// ")
			sb.WriteString(l)
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

// FormatCommandDocComment formats a clean and practical GoDoc comment for a Vulkan command.
func FormatCommandDocComment(c *indexer.CommandInfo) string {
	var sb strings.Builder

	// Lead text / summary
	lead := fmt.Sprintf("%s executes %s.", c.GoName, c.Name)
	if c.ShortDesc != "" {
		lead = fmt.Sprintf("%s - %s (%s).", c.GoName, c.ShortDesc, c.Name)
	}
	sb.WriteString(FormatDocComment(lead))

	// Parameters documentation
	var paramDocLines []string
	for _, p := range c.Params {
		if p.DocDesc != "" {
			paramDocLines = append(paramDocLines, fmt.Sprintf("- %s: %s", p.GoName, p.DocDesc))
		}
	}
	if len(paramDocLines) > 0 {
		var pSb strings.Builder
		pSb.WriteString("Parameters:\n")
		for _, pl := range paramDocLines {
			pSb.WriteString("  " + pl + "\n")
		}
		sb.WriteString(FormatDocComment("", pSb.String()))
	}

	// Success and error codes
	if len(c.SuccessCodes) > 0 {
		sb.WriteString(FormatDocComment(fmt.Sprintf("Success codes: %s", strings.Join(c.SuccessCodes, ", "))))
	}
	if len(c.ErrorCodes) > 0 {
		sb.WriteString(FormatDocComment(fmt.Sprintf("Error codes: %s", strings.Join(c.ErrorCodes, ", "))))
	}

	// Spec link
	sb.WriteString(FormatDocComment(fmt.Sprintf("Documented at: %s", FormatVulkanDocLink(c.Name))))

	return sb.String()
}

// FormatStructDocComment formats a clean GoDoc comment for a Vulkan struct using the Name/desc header.
func FormatStructDocComment(s *indexer.StructInfo) string {
	lead := fmt.Sprintf("%s is %s.", s.GoName, s.Name)
	if s.ShortDesc != "" {
		lead = fmt.Sprintf("%s - %s (%s).", s.GoName, s.ShortDesc, s.Name)
	} else if s.Comment != "" {
		lead = fmt.Sprintf("%s - %s (%s).", s.GoName, s.Comment, s.Name)
	}

	return FormatDocComment(lead, fmt.Sprintf("Documented at: %s", FormatVulkanDocLink(s.Name)))
}

// CleanFieldDocComment returns a single-line clean doc comment for a struct field.
func CleanFieldDocComment(m indexer.StructMember) string {
	desc := m.DocDesc
	if desc == "" {
		desc = m.Comment
	}
	if desc == "" {
		return ""
	}
	desc = strings.ReplaceAll(desc, "\r\n", " ")
	desc = strings.ReplaceAll(desc, "\n", " ")
	desc = strings.ReplaceAll(desc, "\r", " ")
	desc = strings.ReplaceAll(desc, "\t", " ")
	desc = strings.Join(strings.Fields(desc), " ")
	return desc
}

// FormatVulkanDocLink generates a link to Khronos specification documentation.
func FormatVulkanDocLink(symbolName string) string {
	return fmt.Sprintf("https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/%s.html", symbolName)
}

func formatLines(text string) []string {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	rawLines := strings.Split(text, "\n")
	var result []string
	for _, l := range rawLines {
		result = append(result, strings.TrimRight(l, " \t"))
	}
	return result
}
