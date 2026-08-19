package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// ParseFile parses a vk.xml file from the given path.
func ParseFile(path string) (*Registry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open vk.xml: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read vk.xml: %w", err)
	}

	return ParseBytes(data)
}

// ParseBytes parses vk.xml content from byte slice.
func ParseBytes(data []byte) (*Registry, error) {
	var reg Registry
	if err := xml.Unmarshal(data, &reg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal vk.xml: %w", err)
	}

	// Post-process types and commands to normalize names and inner text
	for i := range reg.Types.Types {
		t := &reg.Types.Types[i]
		if t.Name == "" && t.NameTag != "" {
			t.Name = t.NameTag
		}
		if t.Name == "" {
			// Extract name from inner text if needed (e.g. typedef ... <name>VkFlags</name>)
			t.Name = extractTagContent(t.InnerXML, "name")
		}
	}

	return &reg, nil
}

func extractTagContent(innerXML, tagName string) string {
	openTag := "<" + tagName + ">"
	closeTag := "</" + tagName + ">"
	start := strings.Index(innerXML, openTag)
	if start == -1 {
		return ""
	}
	start += len(openTag)
	end := strings.Index(innerXML[start:], closeTag)
	if end == -1 {
		return ""
	}
	return strings.TrimSpace(innerXML[start : start+end])
}
