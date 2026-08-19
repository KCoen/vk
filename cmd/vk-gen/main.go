package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/KCoen/vk/generator/docparser"
	"github.com/KCoen/vk/generator/emitter"
	"github.com/KCoen/vk/generator/indexer"
	"github.com/KCoen/vk/generator/parser"
)

func main() {
	xmlPath := flag.String("xml", "Vulkan-Docs/xml/vk.xml", "Path to Vulkan vk.xml file")
	docsPath := flag.String("docs", "Vulkan-Docs", "Path to Vulkan-Docs repository directory")
	outDir := flag.String("out", ".", "Output directory for generated Go packages")
	flag.Parse()

	fmt.Printf("Starting Vulkan Go Codegen from %s to %s...\n", *xmlPath, *outDir)
	start := time.Now()

	// 1. Parse AsciiDoc documentation from Vulkan-Docs
	fmt.Printf("[1/4] Parsing Vulkan documentation from %s...\n", *docsPath)
	docIdx, err := docparser.ParseVulkanDocs(*docsPath)
	if err != nil {
		fmt.Printf("Warning: Could not parse Vulkan docs: %v (continuing with XML comments only)\n", err)
	} else {
		fmt.Printf("      - Parsed %d refpages with descriptions and parameter docs\n", len(docIdx.Refpages))
	}

	// 2. Parse XML
	fmt.Printf("[2/4] Parsing %s...\n", *xmlPath)
	reg, err := parser.ParseFile(*xmlPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing XML: %v\n", err)
		os.Exit(1)
	}

	// 3. Index semantic registry
	fmt.Printf("[3/4] Indexing Vulkan registry data...\n")
	idx, err := indexer.BuildIndex(reg, docIdx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building index: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("      - Handles: %d\n", len(idx.Handles))
	fmt.Printf("      - Enums: %d\n", len(idx.EnumGroups))
	fmt.Printf("      - Bitmasks: %d\n", len(idx.Bitmasks))
	fmt.Printf("      - Structs/Unions: %d\n", len(idx.Structs))
	fmt.Printf("      - Commands: %d\n", len(idx.Commands))
	fmt.Printf("      - API Branches: %d (vulkan, vulkanbase, vulkansc)\n", len(idx.ApiBranches))
	fmt.Printf("      - Extensions: %d\n", len(idx.Extensions))

	// 4. Emit Go code
	fmt.Printf("[4/4] Emitting Go packages...\n")
	if err := emitter.GenerateAll(idx, *outDir); err != nil {
		fmt.Fprintf(os.Stderr, "Error emitting code: %v\n", err)
		os.Exit(1)
	}

	elapsed := time.Since(start)
	fmt.Printf("Successfully generated all Vulkan packages in %v!\n", elapsed)
}
