package docparser

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// DocRefpage contains parsed documentation for a Vulkan refpage.
type DocRefpage struct {
	Name        string
	ShortDesc   string
	Description string
	Params      map[string]string
	Members     map[string]string
	Type        string
}

// DocIndex stores all parsed refpages indexed by name.
type DocIndex struct {
	Refpages      map[string]*DocRefpage
	ExtensionDocs map[string]string // extName -> cleaned plain text overview
}

// EnsureVulkanDocs checks if docsDir exists; if not, clones KhronosGroup/Vulkan-Docs.
func EnsureVulkanDocs(docsDir string) error {
	if _, err := os.Stat(docsDir); err == nil {
		return nil
	}

	fmt.Printf("Cloning https://github.com/KhronosGroup/Vulkan-Docs to %s...\n", docsDir)
	cmd := exec.Command("git", "clone", "--depth", "1", "https://github.com/KhronosGroup/Vulkan-Docs.git", docsDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ParseVulkanDocs traverses docsDir and parses all AsciiDoc refpage blocks.
func ParseVulkanDocs(docsDir string) (*DocIndex, error) {
	idx := &DocIndex{
		Refpages:      make(map[string]*DocRefpage),
		ExtensionDocs: make(map[string]string),
	}

	if err := EnsureVulkanDocs(docsDir); err != nil {
		return nil, fmt.Errorf("failed to ensure Vulkan-Docs: %w", err)
	}

	openPattern := regexp.MustCompile(`\[open,\s*refpage=[\x27\x22]([^\x27\x22]+)[\x27\x22](?:,\s*desc=[\x27\x22]([^\x27\x22]*)[\x27\x22])?(?:,\s*type=[\x27\x22]([^\x27\x22]*)[\x27\x22])?.*?\]`)

	err := filepath.Walk(docsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() || !strings.HasSuffix(path, ".adoc") {
			return nil
		}

		contentBytes, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		content := string(contentBytes)

		parseRefpageBlocks(content, openPattern, idx)
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Parse extension overview .adoc files in appendices/
	appendicesDir := filepath.Join(docsDir, "appendices")
	if entries, err := os.ReadDir(appendicesDir); err == nil {
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasPrefix(entry.Name(), "VK_") && strings.HasSuffix(entry.Name(), ".adoc") {
				extName := strings.TrimSuffix(entry.Name(), ".adoc")
				fullPath := filepath.Join(appendicesDir, entry.Name())
				if contentBytes, err := os.ReadFile(fullPath); err == nil {
					cleaned := CleanExtensionAdoc(string(contentBytes))
					if cleaned != "" {
						idx.ExtensionDocs[extName] = cleaned
					}
				}
			}
		}
	}

	return idx, nil
}

func parseRefpageBlocks(content string, openPattern *regexp.Regexp, idx *DocIndex) {
	locs := openPattern.FindAllStringSubmatchIndex(content, -1)
	if len(locs) == 0 {
		return
	}

	for i, loc := range locs {
		nameMatch := content[loc[2]:loc[3]]
		descMatch := ""
		if loc[4] != -1 && loc[5] != -1 {
			descMatch = content[loc[4]:loc[5]]
		}
		typeMatch := ""
		if loc[6] != -1 && loc[7] != -1 {
			typeMatch = content[loc[6]:loc[7]]
		}

		startPos := loc[1]
		endPos := len(content)
		if i+1 < len(locs) {
			endPos = locs[i+1][0]
		}

		blockText := content[startPos:endPos]
		// Find block boundaries delimited by "--"
		dashIdx := strings.Index(blockText, "--")
		if dashIdx != -1 {
			rest := blockText[dashIdx+2:]
			endDash := strings.Index(rest, "--")
			if endDash != -1 {
				blockText = rest[:endDash]
			} else {
				blockText = rest
			}
		}

		refpage := parseSingleRefpageBlock(nameMatch, descMatch, typeMatch, blockText)
		idx.Refpages[nameMatch] = refpage
	}
}

var (
	noteBlockRegex  = regexp.MustCompile(`(?s)\[(NOTE|WARNING|IMPORTANT|TIP|CAUTION)\]\s*={4,}.*?={4,}`)
	validUsageRegex = regexp.MustCompile(`(?s)\.(?:Valid Usage|Host Synchronization|Command Properties).*`)
	sidebarBlockReg = regexp.MustCompile(`(?s)\*{4,}.*?\*{4,}`)
	commentRegex    = regexp.MustCompile(`(?m)//[^\n]*$`)
	bulletSplitReg  = regexp.MustCompile(`(?m)\n\s*\*\s+`)
)

func parseSingleRefpageBlock(name, shortDesc, typeKind, block string) *DocRefpage {
	ref := &DocRefpage{
		Name:      name,
		ShortDesc: cleanAsciiDoc(shortDesc),
		Type:      typeKind,
		Params:    make(map[string]string),
		Members:   make(map[string]string),
	}

	// 1. Remove NOTE/WARNING/IMPORTANT blocks, Valid Usage sections, sidebar blocks, and comments
	cleanedBlock := noteBlockRegex.ReplaceAllString(block, "")
	cleanedBlock = validUsageRegex.ReplaceAllString(cleanedBlock, "")
	cleanedBlock = sidebarBlockReg.ReplaceAllString(cleanedBlock, "")
	cleanedBlock = commentRegex.ReplaceAllString(cleanedBlock, "")

	// 2. Extract bullet points (parameters, members, features)
	items := bulletSplitReg.Split(cleanedBlock, -1)
	if len(items) > 1 {
		blankLineRegex := regexp.MustCompile(`(?m)\n\s*\n`)
		for _, item := range items[1:] {
			// Take only the text before any blank line (which marks the end of this list item)
			bulletText := blankLineRegex.Split(item, 2)[0]
			cleanedItem := cleanAsciiDoc(bulletText)
			words := strings.Fields(cleanedItem)
			if len(words) == 0 {
				continue
			}

			rawField := words[0]
			fieldName := strings.TrimSuffix(rawField, ":")
			fieldDesc := strings.TrimSpace(strings.TrimPrefix(cleanedItem, rawField))
			fieldDesc = strings.TrimPrefix(fieldDesc, ":")
			fieldDesc = strings.TrimSpace(fieldDesc)

			if fieldName != "" && fieldDesc != "" {
				ref.Params[fieldName] = fieldDesc
				ref.Members[fieldName] = fieldDesc
			}
		}
	}

	// 3. Extract introductory description from preamble
	preamble := items[0]
	lines := strings.Split(preamble, "\n")
	var descLines []string
	for _, l := range lines {
		trimmed := strings.TrimSpace(l)
		if trimmed == "" || strings.HasPrefix(trimmed, "include::") || strings.HasPrefix(trimmed, ":") {
			continue
		}
		if strings.HasPrefix(trimmed, "ifdef::") || strings.HasPrefix(trimmed, "ifndef::") || strings.HasPrefix(trimmed, "endif::") {
			continue
		}
		descLines = append(descLines, trimmed)
	}

	if len(descLines) > 0 {
		ref.Description = cleanAsciiDoc(strings.Join(descLines, " "))
	}

	return ref
}

var (
	latexmathReg  = regexp.MustCompile(`latexmath:\[(.*?)\]`)
	eqReg         = regexp.MustCompile(`\[eq\]#(.*?)#`)
	anchorReg     = regexp.MustCompile(`\[\[[^\]]+\]\]`)
	attrDefReg    = regexp.MustCompile(`(?m)^:[a-zA-Z0-9_\-]+:[^\n]*$`)
	macroRegex    = regexp.MustCompile(`(?:pname|fname|flink|sname|slink|elink|ename|tlink|tname|dlink|dname|bname|code):([a-zA-Z0-9_:]+)`)
	specWordReg   = regexp.MustCompile(`\b(must|should|may|can|cannot|undefined):`)
	xrefLabelReg  = regexp.MustCompile(`<<[^,>]+,\s*([^>]+)>>`)
	xrefReg       = regexp.MustCompile(`<<([^>]+)>>`)
	includeReg    = regexp.MustCompile(`(?m)^include::[^\n]*$`)
	ifdefReg      = regexp.MustCompile(`(?m)^(?:ifdef|ifndef|endif)::[^\n]*$`)
	backtickReg   = regexp.MustCompile("`([^`]+)`")
	noteHeaderReg = regexp.MustCompile(`\[(NOTE|WARNING|IMPORTANT|TIP|CAUTION)\]`)
	multiSpaceReg = regexp.MustCompile(`\s+`)
	spacePunctReg = regexp.MustCompile(`\s+([,.;:])`)
)

func cleanAsciiDoc(text string) string {
	text = noteBlockRegex.ReplaceAllString(text, "")
	text = latexmathReg.ReplaceAllString(text, "$1")
	text = eqReg.ReplaceAllString(text, "$1")
	text = anchorReg.ReplaceAllString(text, "")
	text = attrDefReg.ReplaceAllString(text, "")
	text = macroRegex.ReplaceAllString(text, "$1")
	text = specWordReg.ReplaceAllString(text, "$1")
	text = xrefLabelReg.ReplaceAllString(text, "$1")
	text = xrefReg.ReplaceAllString(text, "$1")
	text = includeReg.ReplaceAllString(text, "")
	text = ifdefReg.ReplaceAllString(text, "")
	text = backtickReg.ReplaceAllString(text, "$1")
	text = noteHeaderReg.ReplaceAllString(text, "")
	text = strings.ReplaceAll(text, "{generated}", "")
	text = strings.ReplaceAll(text, "{anchor-prefix}", "")
	text = multiSpaceReg.ReplaceAllString(text, " ")
	text = spacePunctReg.ReplaceAllString(text, "$1")
	return strings.TrimSpace(text)
}

// Find retrieves documentation for the given symbol name.
func (idx *DocIndex) Find(name string) *DocRefpage {
	if idx == nil || idx.Refpages == nil {
		return nil
	}
	return idx.Refpages[name]
}

var (
	adocHeadingReg   = regexp.MustCompile(`^=+\s+(.*)`)
	adocDefItemReg   = regexp.MustCompile(`^\*([^*]+)\*::`)
	adocDefItemReg2  = regexp.MustCompile(`^\*([^*]+)\*:`)
	adocMacroReg     = regexp.MustCompile(`(?:pname|fname|flink|sname|slink|elink|ename|tlink|tname|dlink|dname|bname|code|apiext|text|etext|ptext|dtext|vname|vlink):([a-zA-Z0-9_:]+)`)
	adocXrefLabelReg = regexp.MustCompile(`<<[^,>]+,\s*([^>]+)>>`)
	adocXrefReg      = regexp.MustCompile(`<<([^>]+)>>`)
	adocEqReg        = regexp.MustCompile(`(?:\[eq\])?#([^#]+)#`)
	adocBacktickReg  = regexp.MustCompile("`([^`]+)`")
	adocPlusReg      = regexp.MustCompile(`\+([^+]+)\+`)
	adocBoldReg      = regexp.MustCompile(`\*([^*]+)\*`)
	adocItalicReg    = regexp.MustCompile(`_([^_]+)_`)
)

// CleanExtensionAdoc strips AsciiDoc syntax from an extension overview file content and returns clean plain text.
func CleanExtensionAdoc(content string) string {
	lines := strings.Split(content, "\n")
	var resultLines []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "//") {
			continue
		}
		if strings.HasPrefix(trimmed, "include::") {
			continue
		}
		if strings.HasPrefix(trimmed, "ifdef::") || strings.HasPrefix(trimmed, "ifndef::") || strings.HasPrefix(trimmed, "endif::") || strings.HasPrefix(trimmed, "ifeval::") {
			continue
		}
		if trimmed == "====" || trimmed == "----" || trimmed == "...." || trimmed == "****" || trimmed == "--" {
			continue
		}
		if strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]") {
			continue
		}

		if m := adocHeadingReg.FindStringSubmatch(trimmed); len(m) > 1 {
			title := strings.TrimSpace(m[1])
			if title != "" {
				if !strings.HasSuffix(title, ":") {
					title += ":"
				}
				resultLines = append(resultLines, title)
			}
			continue
		}

		if m := adocDefItemReg.FindStringSubmatch(trimmed); len(m) > 1 {
			line = m[1] + ":"
		} else if m := adocDefItemReg2.FindStringSubmatch(trimmed); len(m) > 1 {
			line = m[1] + ":"
		}

		cleaned := cleanInlineAdoc(line)

		if strings.HasPrefix(strings.TrimSpace(cleaned), "*") {
			indent := ""
			for _, r := range cleaned {
				if r == ' ' || r == '\t' {
					indent += string(r)
				} else {
					break
				}
			}
			rest := strings.TrimLeft(cleaned, " \t")
			if strings.HasPrefix(rest, "*** ") {
				cleaned = indent + "    - " + strings.TrimPrefix(rest, "*** ")
			} else if strings.HasPrefix(rest, "** ") {
				cleaned = indent + "  - " + strings.TrimPrefix(rest, "** ")
			} else if strings.HasPrefix(rest, "* ") {
				cleaned = indent + "- " + strings.TrimPrefix(rest, "* ")
			}
		}

		resultLines = append(resultLines, cleaned)
	}

	var finalLines []string
	prevEmpty := true
	for _, l := range resultLines {
		isEmpty := strings.TrimSpace(l) == ""
		if isEmpty && prevEmpty {
			continue
		}
		finalLines = append(finalLines, l)
		prevEmpty = isEmpty
	}

	return strings.TrimSpace(strings.Join(finalLines, "\n"))
}

func cleanInlineAdoc(text string) string {
	text = adocMacroReg.ReplaceAllString(text, "$1")
	text = adocXrefLabelReg.ReplaceAllString(text, "$1")
	text = adocXrefReg.ReplaceAllString(text, "$1")
	text = adocEqReg.ReplaceAllString(text, "$1")
	text = adocBacktickReg.ReplaceAllString(text, "$1")
	text = adocPlusReg.ReplaceAllString(text, "$1")
	text = adocBoldReg.ReplaceAllString(text, "$1")
	text = adocItalicReg.ReplaceAllString(text, "$1")
	text = strings.ReplaceAll(text, "{generated}", "")
	text = strings.ReplaceAll(text, "{refprefix}", "")
	text = strings.ReplaceAll(text, "{anchor-prefix}", "")
	return text
}
