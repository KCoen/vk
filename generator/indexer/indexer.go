package indexer

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"go.cld.moe/vk_google/generator/docparser"
	"go.cld.moe/vk_google/generator/parser"
)

// BuildIndex indexes the parsed XML registry into a semantic database.
func BuildIndex(reg *parser.Registry, docIdx *docparser.DocIndex) (*Index, error) {
	idx := &Index{
		Platforms:    make(map[string]parser.Platform),
		Tags:         make(map[string]parser.Tag),
		Constants:    make(map[string]ConstantInfo),
		Handles:      make(map[string]HandleInfo),
		BaseTypes:    make(map[string]string),
		EnumGroups:   make(map[string]*EnumGroup),
		Bitmasks:     make(map[string]BitmaskInfo),
		Structs:      make(map[string]*StructInfo),
		TypeAliases:  make(map[string]TypeAliasInfo),
		FuncPointers:    make(map[string]string),
		FuncPointerDefs: make(map[string]*FuncPointerInfo),
		Commands:        make(map[string]*CommandInfo),
		ApiBranches: make(map[string]*ApiBranchInfo),
		Versions:    make(map[string]*VersionInfo),
		Extensions:  make(map[string]*ExtensionInfo),
		Formats:     make(map[string]FormatInfo),
		SyncStages:  make(map[string]SyncStageInfo),
		SyncAccesses: make(map[string]SyncAccessInfo),
		SpirvExts:   reg.SpirvExtensions.Extensions,
		SpirvCaps:   reg.SpirvCapabilities.Capabilities,
		StructExtensions: make(map[string][]string),
		DocIndex:    docIdx,
	}

	for _, bName := range []string{"vulkan", "vulkanbase", "vulkansc"} {
		title := "Vulkan Core API"
		if bName == "vulkanbase" {
			title = "Vulkan Base API"
		} else if bName == "vulkansc" {
			title = "Vulkan Safety Critical (SC) API"
		}
		idx.ApiBranches[bName] = &ApiBranchInfo{
			Name:            bName,
			PkgName:         bName,
			Title:           title,
			Handles:         make(map[string]HandleInfo),
			Bitmasks:        make(map[string]BitmaskInfo),
			Structs:         make(map[string]*StructInfo),
			EnumGroups:      make(map[string]*EnumGroup),
			Constants:       make(map[string]ConstantInfo),
			TypeAliases:     make(map[string]TypeAliasInfo),
			FuncPointerDefs: make(map[string]*FuncPointerInfo),
		}
	}

	// 1. Index Tags
	for _, t := range reg.Tags.Tags {
		idx.Tags[t.Name] = t
	}

	// 2. Index Types (Basetypes, Handles, Bitmasks, Structs, Unions, Aliases, FuncPointers)
	for _, t := range reg.Types.Types {
		name := t.Name
		if name == "" {
			name = t.NameTag
		}
		if name == "" && t.Category == "funcpointer" {
			name = t.Proto.Name
			if name == "" {
				name = extractTagContent(t.InnerXML, "name")
			}
		}
		if name == "" {
			continue
		}

		// Check for alias
		if t.Alias != "" {
			alias := TypeAliasInfo{
				Name:     name,
				GoName:   CleanTypeName(name),
				AliasOf:  t.Alias,
				GoAlias:  CleanTypeName(t.Alias),
				Category: t.Category,
				Comment:  t.Comment,
				Api:      t.Api,
			}
			idx.TypeAliases[name] = alias
			for bName, branch := range idx.ApiBranches {
				if ApiMatches(t.Api, bName) {
					branch.TypeAliases[name] = alias
				}
			}
			continue
		}

		switch t.Category {
		case "basetype":
			idx.BaseTypes[name] = t.TypeTag

		case "handle":
			kind := HandleDispatchable
			if strings.Contains(t.InnerXML, "VK_DEFINE_NON_DISPATCHABLE_HANDLE") {
				kind = HandleNonDispatchable
			}
			h := HandleInfo{
				Name:        name,
				GoName:      CleanTypeName(name),
				Parent:      t.Parent,
				ObjTypeEnum: t.ObjTypeEnum,
				Kind:        kind,
				Comment:     t.Comment,
				Api:         t.Api,
			}
			idx.Handles[name] = h
			for bName, branch := range idx.ApiBranches {
				if ApiMatches(t.Api, bName) {
					branch.Handles[name] = h
				}
			}

		case "bitmask":
			bitwidth := 32
			if t.TypeTag == "VkFlags64" || strings.Contains(t.InnerXML, "VkFlags64") {
				bitwidth = 64
			}
			bitsType := t.Requires
			if bitsType == "" {
				bitsType = t.Bitvalues
			}
			bm := BitmaskInfo{
				Name:     name,
				GoName:   CleanTypeName(name),
				BitsType: bitsType,
				Bitwidth: bitwidth,
				Comment:  t.Comment,
				Api:      t.Api,
			}
			idx.Bitmasks[name] = bm
			for bName, branch := range idx.ApiBranches {
				if ApiMatches(t.Api, bName) {
					branch.Bitmasks[name] = bm
				}
			}

		case "struct", "union":
			sInfo := parseStructInfo(t)
			idx.Structs[name] = sInfo
			if t.Api != "" {
				for bName, branch := range idx.ApiBranches {
					if ApiMatches(t.Api, bName) {
						branch.Structs[name] = sInfo
					}
				}
			}

		case "funcpointer":
			fpInfo := parseFuncPointerInfo(t)
			idx.FuncPointerDefs[fpInfo.Name] = fpInfo
			idx.FuncPointers[fpInfo.Name] = fpInfo.Name
			for bName, branch := range idx.ApiBranches {
				if ApiMatches(t.Api, bName) {
					branch.FuncPointerDefs[fpInfo.Name] = fpInfo
				}
			}

		case "enum":
			// Listed in <types>, defined in <enums> or handled as alias
		}
	}

	// 3. Index Constants
	for _, eg := range reg.Enums {
		if eg.Name == "API Constants" {
			for _, e := range eg.Enums {
				val := normalizeConstantValue(e.Value)
				c := ConstantInfo{
					Name:    e.Name,
					GoName:  CleanConstantName(e.Name),
					Value:   val,
					Comment: e.Comment,
					Type:    e.Type,
					Api:     e.Api,
				}
				idx.Constants[e.Name] = c
				for bName, branch := range idx.ApiBranches {
					if ApiMatches(e.Api, bName) {
						branch.Constants[e.Name] = c
					}
				}
			}
		}
	}

	// 4. Index Enums and Bitmask Groups
	for _, eg := range reg.Enums {
		if eg.Name == "API Constants" {
			continue
		}

		bitwidth := 32
		if eg.Bitwidth == "64" {
			bitwidth = 64
		}

		group := &EnumGroup{
			Name:     eg.Name,
			GoName:   CleanTypeName(eg.Name),
			Type:     eg.Type,
			Bitwidth: bitwidth,
			Comment:  eg.Comment,
			Api:      eg.Api,
		}

		for _, e := range eg.Enums {
			valStr, valInt := resolveEnumValue(e)
			group.Values = append(group.Values, EnumValue{
				Name:        e.Name,
				GoName:      CleanEnumName(e.Name),
				ValueStr:    valStr,
				ValueInt:    valInt,
				Comment:     e.Comment,
				IsAlias:     e.Alias != "",
				AliasOf:     CleanEnumName(e.Alias),
				ExtendsEnum: e.Extends,
				Api:         e.Api,
			})
		}
		idx.EnumGroups[eg.Name] = group

		for bName, branch := range idx.ApiBranches {
			if !ApiMatches(eg.Api, bName) {
				continue
			}
			bGroup := &EnumGroup{
				Name:     eg.Name,
				GoName:   CleanTypeName(eg.Name),
				Type:     eg.Type,
				Bitwidth: bitwidth,
				Comment:  eg.Comment,
				Api:      eg.Api,
			}
			for _, e := range eg.Enums {
				if !ApiMatches(e.Api, bName) {
					continue
				}
				valStr, valInt := resolveEnumValue(e)
				bGroup.Values = append(bGroup.Values, EnumValue{
					Name:        e.Name,
					GoName:      CleanEnumName(e.Name),
					ValueStr:    valStr,
					ValueInt:    valInt,
					Comment:     e.Comment,
					IsAlias:     e.Alias != "",
					AliasOf:     CleanEnumName(e.Alias),
					ExtendsEnum: e.Extends,
					Api:         e.Api,
				})
			}
			branch.EnumGroups[eg.Name] = bGroup
		}
	}

	// 4. Index Extension Enums & Features enum extensions
	// Process Features first
	for _, f := range reg.Features {
		fApis := strings.Split(f.Api, ",")
		for _, req := range f.Requires {
			for _, e := range req.Enums {
				extNum := 0
				if e.ExtNumber != "" {
					extNum, _ = strconv.Atoi(e.ExtNumber)
				}
				var valStr string
				var valInt int64
				var cInfo *ConstantInfo
				if e.Extends != "" {
					valStr, valInt = calculateExtensionEnumValue(e, extNum)
					group := idx.EnumGroups[e.Extends]
					if group == nil {
						group = &EnumGroup{
							Name:   e.Extends,
							GoName: CleanTypeName(e.Extends),
							Type:   "enum",
						}
						idx.EnumGroups[e.Extends] = group
					}
					group.Values = append(group.Values, EnumValue{
						Name:        e.Name,
						GoName:      CleanEnumName(e.Name),
						ValueStr:    valStr,
						ValueInt:    valInt,
						Comment:     e.Comment,
						IsAlias:     e.Alias != "",
						AliasOf:     CleanEnumName(e.Alias),
						ExtendsEnum: e.Extends,
					})
				} else if e.Value != "" {
					c := ConstantInfo{
						Name:    e.Name,
						GoName:  CleanConstantName(e.Name),
						Value:   normalizeConstantValue(e.Value),
						Comment: e.Comment,
					}
					idx.Constants[e.Name] = c
					cInfo = &c
				}

				for _, a := range fApis {
					a = strings.TrimSpace(a)
					if branch, ok := idx.ApiBranches[a]; ok {
						if ApiMatches(req.Api, a) && ApiMatches(e.Api, a) {
							addEnumExtensionToBranch(branch, e, valStr, valInt, cInfo)
						}
					}
				}
			}
		}
	}

	// Process Extensions enums
	for _, ext := range reg.Extensions.Extensions {
		extNum, _ := strconv.Atoi(ext.Number)
		for _, req := range ext.Requires {
			for _, e := range req.Enums {
				currentExtNum := extNum
				if e.ExtNumber != "" {
					currentExtNum, _ = strconv.Atoi(e.ExtNumber)
				}
				var valStr string
				var valInt int64
				var cInfo *ConstantInfo
				if e.Extends != "" {
					valStr, valInt = calculateExtensionEnumValue(e, currentExtNum)
					group := idx.EnumGroups[e.Extends]
					if group == nil {
						group = &EnumGroup{
							Name:   e.Extends,
							GoName: CleanTypeName(e.Extends),
							Type:   "enum",
						}
						idx.EnumGroups[e.Extends] = group
					}
					group.Values = append(group.Values, EnumValue{
						Name:        e.Name,
						GoName:      CleanEnumName(e.Name),
						ValueStr:    valStr,
						ValueInt:    valInt,
						Comment:     e.Comment,
						IsAlias:     e.Alias != "",
						AliasOf:     CleanEnumName(e.Alias),
						ExtendsEnum: e.Extends,
					})
				} else if e.Value != "" {
					c := ConstantInfo{
						Name:    e.Name,
						GoName:  CleanConstantName(e.Name),
						Value:   normalizeConstantValue(e.Value),
						Comment: e.Comment,
					}
					idx.Constants[e.Name] = c
					cInfo = &c
				}

				for bName, branch := range idx.ApiBranches {
					if ApiMatches(ext.Supported, bName) && ApiMatches(req.Api, bName) && ApiMatches(e.Api, bName) {
						addEnumExtensionToBranch(branch, e, valStr, valInt, cInfo)
					}
				}
			}
		}
	}

	// 5. Index Commands
	for _, cmd := range reg.Commands.Commands {
		if cmd.Name != "" && cmd.Alias != "" {
			// Command alias
			idx.Commands[cmd.Name] = &CommandInfo{
				Name:    cmd.Name,
				GoName:  CleanCmdName(cmd.Name),
				IsAlias: true,
				AliasOf: CleanCmdName(cmd.Alias),
			}
			if cmd.Api != "" {
				for _, a := range strings.Split(cmd.Api, ",") {
					a = strings.TrimSpace(a)
					if branch, ok := idx.ApiBranches[a]; ok {
						branch.Commands = append(branch.Commands, cmd.Name)
					}
				}
			}
			continue
		}

		proto := cmd.Proto
		if proto.Name == "" {
			continue
		}

		cinfo := parseCommandInfo(cmd)
		idx.Commands[proto.Name] = cinfo

		if cmd.Api != "" {
			for _, a := range strings.Split(cmd.Api, ",") {
				a = strings.TrimSpace(a)
				if branch, ok := idx.ApiBranches[a]; ok {
					branch.Commands = append(branch.Commands, proto.Name)
				}
			}
		}
	}

	// 6. Index Features into API Branches (vulkan, vulkanbase, vulkansc) and Core Versions
	versionNums := []string{"1.0", "1.1", "1.2", "1.3", "1.4"}
	for _, v := range versionNums {
		pkgName := "v" + strings.ReplaceAll(v, ".", "_")
		vinfo := &VersionInfo{
			Number:  v,
			PkgName: pkgName,
		}
		idx.Versions[v] = vinfo
	}

	for _, f := range reg.Features {
		apis := strings.Split(f.Api, ",")
		for _, a := range apis {
			a = strings.TrimSpace(a)
			branch := idx.ApiBranches[a]
			if branch != nil {
				for _, req := range f.Requires {
					if req.Api != "" && !ApiMatches(req.Api, a) {
						continue
					}
					for _, t := range req.Types {
						if sInfo, ok := idx.Structs[t.Name]; ok {
							branch.Structs[t.Name] = sInfo
						}
						if hInfo, ok := idx.Handles[t.Name]; ok {
							branch.Handles[t.Name] = hInfo
						}
						if bInfo, ok := idx.Bitmasks[t.Name]; ok {
							branch.Bitmasks[t.Name] = bInfo
						}
						if aInfo, ok := idx.TypeAliases[t.Name]; ok {
							branch.TypeAliases[t.Name] = aInfo
						}
						if fpInfo, ok := idx.FuncPointerDefs[t.Name]; ok {
							branch.FuncPointerDefs[t.Name] = fpInfo
						}
					}
					for _, c := range req.Commands {
						if c.Api != "" && !ApiMatches(c.Api, a) {
							continue
						}
						branch.Commands = append(branch.Commands, c.Name)
					}
				}
			}
		}

		if strings.Contains(f.Api, "vulkan") {
			vinfo := idx.Versions[f.Number]
			if vinfo != nil {
				for _, req := range f.Requires {
					for _, t := range req.Types {
						vinfo.Types = append(vinfo.Types, t.Name)
					}
					for _, e := range req.Enums {
						vinfo.Enums = append(vinfo.Enums, e.Name)
					}
					for _, c := range req.Commands {
						vinfo.Commands = append(vinfo.Commands, c.Name)
					}
				}
			}
		}
	}

	// 7. Index Extension types into API Branches
	for _, ext := range reg.Extensions.Extensions {
		if ext.Supported == "" || ext.Supported == "disabled" {
			continue
		}
		for bName, branch := range idx.ApiBranches {
			if ApiMatches(ext.Supported, bName) {
				for _, req := range ext.Requires {
					if ApiMatches(req.Api, bName) {
						for _, t := range req.Types {
							if sInfo, ok := idx.Structs[t.Name]; ok {
								branch.Structs[t.Name] = sInfo
							}
							if hInfo, ok := idx.Handles[t.Name]; ok {
								branch.Handles[t.Name] = hInfo
							}
							if bInfo, ok := idx.Bitmasks[t.Name]; ok {
								branch.Bitmasks[t.Name] = bInfo
							}
							if aInfo, ok := idx.TypeAliases[t.Name]; ok {
								branch.TypeAliases[t.Name] = aInfo
							}
							if fpInfo, ok := idx.FuncPointerDefs[t.Name]; ok {
								branch.FuncPointerDefs[t.Name] = fpInfo
							}
						}
					}
				}
			}
		}
	}

	// Transitive closure for structs in each ApiBranch
	for _, branch := range idx.ApiBranches {
		for {
			added := false
			for _, s := range branch.Structs {
				for _, m := range s.Members {
					mTypeName := m.GoHighType
					mTypeName = strings.TrimPrefix(mTypeName, "[]")
					mTypeName = strings.TrimPrefix(mTypeName, "*")
					vkName := "Vk" + mTypeName
					if nestedS, ok := idx.Structs[vkName]; ok {
						if _, exists := branch.Structs[vkName]; !exists {
							branch.Structs[vkName] = nestedS
							added = true
						}
					}
					if nestedS, ok := idx.Structs[m.Name]; ok {
						if _, exists := branch.Structs[m.Name]; !exists {
							branch.Structs[m.Name] = nestedS
							added = true
						}
					}
				}
			}
			if !added {
				break
			}
		}
	}

	// Deduplicate ApiBranch commands
	for _, branch := range idx.ApiBranches {
		var uniqueCmds []string
		seen := make(map[string]bool)
		for _, c := range branch.Commands {
			if !seen[c] {
				seen[c] = true
				uniqueCmds = append(uniqueCmds, c)
			}
		}
		branch.Commands = uniqueCmds
	}

	// 7. Attach AsciiDoc Documentation from DocIndex
	if docIdx != nil {
		for _, cmd := range idx.Commands {
			if ref := docIdx.Find(cmd.Name); ref != nil {
				cmd.ShortDesc = ref.ShortDesc
				cmd.DocDesc = ref.Description
				cmd.ParamDocs = ref.Params
				for i := range cmd.Params {
					if pDoc, ok := ref.Params[cmd.Params[i].Name]; ok {
						cmd.Params[i].DocDesc = pDoc
					}
				}
			}
		}

		for _, s := range idx.Structs {
			if ref := docIdx.Find(s.Name); ref != nil {
				s.ShortDesc = ref.ShortDesc
				s.DocDesc = ref.Description
				for i := range s.Members {
					if mDoc, ok := ref.Members[s.Members[i].Name]; ok {
						s.Members[i].DocDesc = mDoc
					}
				}
			}
		}

		for name, h := range idx.Handles {
			if ref := docIdx.Find(name); ref != nil {
				h.ShortDesc = ref.ShortDesc
				h.DocDesc = ref.Description
				idx.Handles[name] = h
			}
		}

		for _, eg := range idx.EnumGroups {
			if ref := docIdx.Find(eg.Name); ref != nil {
				eg.ShortDesc = ref.ShortDesc
				eg.DocDesc = ref.Description
			}
		}

		for name, bm := range idx.Bitmasks {
			if ref := docIdx.Find(name); ref != nil {
				bm.ShortDesc = ref.ShortDesc
				bm.DocDesc = ref.Description
				idx.Bitmasks[name] = bm
			}
		}
	}

	// Index struct extension relationships
	for _, s := range idx.Structs {
		for _, parent := range s.StructExtends {
			parent = strings.TrimSpace(parent)
			if parent != "" {
				idx.StructExtensions[parent] = append(idx.StructExtensions[parent], s.GoName)
			}
		}
	}
	for parent := range idx.StructExtensions {
		sort.Strings(idx.StructExtensions[parent])
	}

	// 7. Index Extensions
	for _, ext := range reg.Extensions.Extensions {
		extNum, _ := strconv.Atoi(ext.Number)
		extInfo := &ExtensionInfo{
			Name:         ext.Name,
			GoName:       CleanExtGoName(ext.Name),
			PkgName:      CleanExtPkgName(ext.Name),
			Number:       extNum,
			Type:         ext.Type,
			Author:       ext.Author,
			Contact:      ext.Contact,
			Supported:    ext.Supported,
			PromotedTo:   ext.PromotedTo,
			DeprecatedBy: ext.DeprecatedBy,
			Depends:      ext.Depends,
			Platform:     ext.Platform,
			Comment:      ext.Comment,
		}

		for _, req := range ext.Requires {
			for _, t := range req.Types {
				extInfo.Types = append(extInfo.Types, t.Name)
			}
			for _, e := range req.Enums {
				extInfo.Enums = append(extInfo.Enums, e.Name)
			}
			for _, c := range req.Commands {
				extInfo.Commands = append(extInfo.Commands, c.Name)
			}
		}
		idx.Extensions[ext.Name] = extInfo
	}

	// 8. Index Formats
	for _, fmtElem := range reg.Formats.Formats {
		idx.Formats[fmtElem.Name] = FormatInfo{
			Name:           fmtElem.Name,
			GoName:         CleanEnumName(fmtElem.Name),
			Class:          fmtElem.Class,
			BlockSize:      fmtElem.BlockSize,
			TexelsPerBlock: fmtElem.TexelsPerBlock,
			Packed:         fmtElem.Packed,
			Components:     fmtElem.Components,
			Planes:         fmtElem.Planes,
		}
	}

	// 9. Index Sync stages & accesses
	for _, s := range reg.Sync.SyncStages {
		idx.SyncStages[s.Name] = SyncStageInfo{
			Name:    s.Name,
			GoName:  CleanEnumName(s.Name),
			Alias:   s.Alias,
			Support: s.Support,
		}
	}
	for _, a := range reg.Sync.SyncAccesses {
		idx.SyncAccesses[a.Name] = SyncAccessInfo{
			Name:    a.Name,
			GoName:  CleanEnumName(a.Name),
			Alias:   a.Alias,
			Support: a.Support,
		}
	}

	return idx, nil
}

func parseStructInfo(t parser.Type) *StructInfo {
	sInfo := &StructInfo{
		Name:          t.Name,
		GoName:        CleanTypeName(t.Name),
		Category:      t.Category,
		Comment:       t.Comment,
		ReturnedOnly:  t.ReturnedOnly == "true",
		IsAlias:       t.Alias != "",
		AliasOf:       CleanTypeName(t.Alias),
		Api:           t.Api,
	}

	if t.StructExtends != "" {
		sInfo.StructExtends = strings.Split(t.StructExtends, ",")
	}

	// First pass: extract all members
	seenMemberNames := make(map[string]bool)
	for _, m := range t.Members {
		if seenMemberNames[m.Name] {
			continue
		}
		seenMemberNames[m.Name] = true

		comment := m.CommentAttr
		if comment == "" {
			comment = m.CommentTag
		}

		cType, isPointer, isDoublePointer, isConst, arrayDims := ParseFullCType(m.InnerXML, m.Type, m.Name)
		goRawType := MapCTypeToGo(cType)
		goHighType := goRawType

		isString := (m.Type == "char" && isPointer && !isDoublePointer)
		isStringSlice := (m.Type == "char" && isDoublePointer)

		if isString {
			goHighType = "string"
		} else if isStringSlice {
			goHighType = "[]string"
		}

		isFuncPointer := strings.HasPrefix(m.Type, "PFN_")
		if isFuncPointer {
			goHighType = m.Type
			goRawType = "uintptr"
		}

		if m.Name == "pNext" {
			goHighType = "any"
			goRawType = "unsafe.Pointer"
		}

		member := StructMember{
			Name:            m.Name,
			GoName:          CleanMemberName(m.Name),
			CType:           cType,
			GoRawType:       goRawType,
			GoHighType:      goHighType,
			IsPointer:       isPointer,
			IsDoublePointer: isDoublePointer,
			IsConst:         isConst,
			IsString:        isString,
			IsStringSlice:   isStringSlice,
			IsFuncPointer:   isFuncPointer,
			ArrayDimensions: arrayDims,
			LenAttr:         m.Len,
			DefaultValue:    m.Values,
			GoDefaultValue:  CleanEnumName(m.Values),
			Comment:         comment,
			Optional:        strings.Contains(m.Optional, "true"),
			Api:             m.Api,
		}

		if m.Name == "sType" {
			sInfo.HasSType = true
			if m.Values != "" {
				sInfo.STypeDefault = m.Values
				sInfo.GoSTypeDefault = CleanEnumName(m.Values)
			}
		}
		if m.Name == "pNext" {
			sInfo.HasNext = true
		}

		sInfo.Members = append(sInfo.Members, member)
	}

	// Deduplicate member GoNames
	seenGoNames := make(map[string]int)
	for i := range sInfo.Members {
		m := &sInfo.Members[i]
		if seenGoNames[m.GoName] > 0 {
			if strings.HasPrefix(m.Name, "pp") && len(m.Name) > 2 {
				m.GoName = "Pp" + strings.ToUpper(m.Name[2:3]) + m.Name[3:]
			} else if strings.HasPrefix(m.Name, "p") && len(m.Name) > 1 {
				m.GoName = "P" + strings.ToUpper(m.Name[1:2]) + m.Name[2:]
			} else {
				m.GoName = fmt.Sprintf("%s%d", m.GoName, seenGoNames[m.GoName]+1)
			}
		}
		seenGoNames[m.GoName]++
	}

	// Second pass: identify count+pointer member pairs to convert to slices in high-level struct
	countMemberMap := make(map[string]int) // name -> index
	for i, m := range sInfo.Members {
		if !m.IsPointer && (strings.HasSuffix(m.Name, "Count") || strings.HasSuffix(m.Name, "Size") || m.Name == "count") {
			countMemberMap[m.Name] = i
		}
	}

	for i := range sInfo.Members {
		m := &sInfo.Members[i]
		if (m.IsPointer || m.IsStringSlice) && !m.IsString && m.Name != "pNext" {
			// Check if there is a matching count member
			for countName := range countMemberMap {
				// e.g. descriptorCount -> pDescriptors or pBindings or pImageInfo
				// e.g. enabledExtensionCount -> ppEnabledExtensionNames
				// e.g. waitSemaphoreCount -> pWaitSemaphores
				// Check prefix match or len match
				if (m.LenAttr != "" && strings.Contains(m.LenAttr, countName)) || isCountMatch(countName, m.Name) {
					if !m.IsStringSlice {
						m.IsSlice = true
						// Transform high-level type to []T
						elemType := strings.TrimPrefix(m.GoRawType, "*")
						m.GoHighType = "[]" + elemType
					}
					m.SliceCountMember = countName
					break
				}
			}
		}
	}

	return sInfo
}

func isCountMatch(countName, ptrName string) bool {
	cStem := strings.TrimSuffix(countName, "Count")
	cStem = strings.TrimSuffix(cStem, "Size")
	pStem := ptrName
	if strings.HasPrefix(pStem, "pp") {
		pStem = pStem[2:]
	} else if strings.HasPrefix(pStem, "p") {
		pStem = pStem[1:]
	}

	if cStem == "" {
		return true
	}
	if strings.EqualFold(cStem, pStem) {
		return true
	}
	if strings.HasPrefix(strings.ToLower(pStem), strings.ToLower(cStem)) {
		return true
	}
	if strings.HasPrefix(strings.ToLower(cStem), strings.ToLower(pStem)) {
		return true
	}
	return false
}

func parseFuncPointerInfo(t parser.Type) *FuncPointerInfo {
	name := t.Name
	if name == "" {
		name = t.NameTag
	}
	if name == "" && t.Proto.Name != "" {
		name = t.Proto.Name
	}
	if name == "" {
		name = extractTagContent(t.InnerXML, "name")
	}

	retCType, isPointer, _, _, _ := ParseFullCType(t.Proto.InnerXML, t.Proto.Type, t.Proto.Name)
	var retGoType string
	if isPointer && retCType == "void*" {
		retGoType = "unsafe.Pointer"
	} else if retCType != "" && retCType != "void" {
		retGoType = MapCTypeToGo(retCType)
	}

	fp := &FuncPointerInfo{
		Name:         name,
		GoName:       name,
		ReturnType:   retCType,
		ReturnGoType: retGoType,
		Comment:      t.Comment,
	}

	for _, p := range t.Params {
		cType, isPointer, _, isConst, _ := ParseFullCType(p.InnerXML, p.Type, p.Name)
		goType := MapCTypeToGo(cType)
		if isPointer && strings.HasPrefix(p.Type, "Vk") {
			structName := strings.TrimPrefix(p.Type, "Vk")
			goType = "*Raw" + structName
		}
		if goType == "" {
			goType = "unsafe.Pointer"
		}
		fp.Params = append(fp.Params, FuncPointerParam{
			Name:      p.Name,
			GoName:    CleanParamName(p.Name),
			Type:      cType,
			GoType:    goType,
			IsPointer: isPointer,
			IsConst:   isConst,
		})
	}

	return fp
}

func parseCommandInfo(cmd parser.Command) *CommandInfo {
	proto := cmd.Proto
	retCType := proto.Type
	retGoType := MapCTypeToGo(retCType)
	if retGoType == "" {
		retGoType = "void"
	}

	cinfo := &CommandInfo{
		Name:         proto.Name,
		GoName:       CleanCmdName(proto.Name),
		ReturnCType:  retCType,
		ReturnGoType: retGoType,
		Queues:       cmd.Queues,
		RenderPass:   cmd.RenderPass,
		CmdBufferLvl: cmd.CmdBufferLvl,
		Comment:      cmd.Comment,
		Api:          cmd.Api,
	}

	if cmd.SuccessCodes != "" {
		cinfo.SuccessCodes = strings.Split(cmd.SuccessCodes, ",")
	}
	if cmd.ErrorCodes != "" {
		cinfo.ErrorCodes = strings.Split(cmd.ErrorCodes, ",")
	}

	// Parse parameters (filtering api and deduplicating names)
	seenParamNames := make(map[string]bool)
	for _, p := range cmd.Params {
		if seenParamNames[p.Name] {
			continue
		}
		seenParamNames[p.Name] = true
		cType, isPointer, isDoublePointer, isConst, _ := ParseFullCType(p.InnerXML, p.Type, p.Name)
		goRawType := MapCTypeToGo(cType)
		goHighType := goRawType

		isString := (p.Type == "char" && isPointer && !isDoublePointer)
		isStringSlice := (p.Type == "char" && isDoublePointer)

		if isString {
			goHighType = "string"
		} else if isStringSlice {
			goHighType = "[]string"
		}

		param := CommandParam{
			Name:            p.Name,
			GoName:          CleanParamName(p.Name),
			CType:           cType,
			GoRawType:       goRawType,
			GoHighType:      goHighType,
			IsPointer:       isPointer,
			IsDoublePointer: isDoublePointer,
			IsConst:         isConst,
			IsString:        isString,
			IsStringSlice:   isStringSlice,
			LenAttr:         p.Len,
			Optional:        strings.Contains(p.Optional, "true"),
			Api:             p.Api,
		}

		cinfo.Params = append(cinfo.Params, param)
	}

	// Classify command as Global, Instance, or Device
	if len(cinfo.Params) > 0 {
		firstParamType := cinfo.Params[0].CType
		switch firstParamType {
		case "VkDevice", "VkQueue", "VkCommandBuffer":
			cinfo.IsDeviceCmd = true
		case "VkInstance", "VkPhysicalDevice":
			cinfo.IsInstanceCmd = true
		default:
			cinfo.IsGlobalCmd = true
		}
	} else {
		cinfo.IsGlobalCmd = true
	}

	// Identify slice params, count params, 2-call queries, output pointers
	analyzeCommandSignatures(cinfo)

	return cinfo
}

func analyzeCommandSignatures(cinfo *CommandInfo) {
	// Identify 2-call query (e.g. vkEnumeratePhysicalDevices(instance, pCount *uint32, pItems *VkPhysicalDevice))
	// Look for count pointer (uint32_t* / size_t*) followed by array pointer with len pointing to count
	for i := 0; i < len(cinfo.Params); i++ {
		p := &cinfo.Params[i]
		if p.IsPointer && !p.IsConst && (p.CType == "uint32_t*" || p.CType == "size_t*") && i+1 < len(cinfo.Params) {
			nextP := &cinfo.Params[i+1]
			if nextP.IsPointer && (nextP.LenAttr == p.Name || strings.HasPrefix(nextP.LenAttr, p.Name)) {
				// Found 2-call query!
				cinfo.IsOutputQuery = true
				cinfo.OutputCountParam = p
				cinfo.OutputQueryParam = nextP
				p.IsCountParam = true
				nextP.IsSlice = true
				break
			}
		}
	}

	// Identify slice + count pairs in inputs (e.g. uint32_t count, const VkBuffer* pBuffers)
	for i := 0; i < len(cinfo.Params); i++ {
		p := &cinfo.Params[i]
		if !p.IsPointer && (strings.HasSuffix(p.Name, "Count") || strings.HasSuffix(p.Name, "Size") || p.Name == "count") {
			// Search for following pointer param that references this count
			for j := i + 1; j < len(cinfo.Params); j++ {
				sliceP := &cinfo.Params[j]
				if sliceP.IsPointer && !sliceP.IsString && (sliceP.LenAttr == p.Name || isCountMatch(p.Name, sliceP.Name)) {
					p.IsCountParam = true
					p.PairedSliceParam = sliceP.Name
					sliceP.IsSlice = true
					elemType := strings.TrimPrefix(sliceP.GoRawType, "*")
					sliceP.GoHighType = "[]" + elemType
					break
				}
			}
		}
	}

	// Single output pointers (e.g. VkInstance* pInstance, VkDevice* pDevice, VkQueue* pQueue, VkPhysicalDeviceMemoryProperties* pMemoryProperties)
	if !cinfo.IsOutputQuery {
		for i := len(cinfo.Params) - 1; i >= 0; i-- {
			p := &cinfo.Params[i]
			if p.IsPointer && !p.IsConst && !p.IsSlice && !p.IsString && !p.IsCountParam && p.Name != "pAllocator" {
				p.IsOutputSingle = true
				break
			}
		}
	}
}

func resolveEnumValue(e parser.Enum) (string, int64) {
	if e.Value != "" {
		valStr := e.Value
		valInt, _ := strconv.ParseInt(valStr, 0, 64)
		return valStr, valInt
	}
	if e.Bitpos != "" {
		bit, _ := strconv.Atoi(e.Bitpos)
		valUint := uint64(1) << bit
		valStr := fmt.Sprintf("0x%X", valUint)
		return valStr, int64(valUint)
	}
	if e.Alias != "" {
		return CleanEnumName(e.Alias), 0
	}
	return "0", 0
}

func calculateExtensionEnumValue(e parser.Enum, defaultExtNum int) (string, int64) {
	if e.Value != "" {
		valInt, _ := strconv.ParseInt(e.Value, 0, 64)
		return e.Value, valInt
	}
	if e.Bitpos != "" {
		bit, _ := strconv.Atoi(e.Bitpos)
		valUint := uint64(1) << bit
		valStr := fmt.Sprintf("0x%X", valUint)
		return valStr, int64(valUint)
	}
	if e.Offset != "" {
		extNum := defaultExtNum
		if e.ExtNumber != "" {
			extNum, _ = strconv.Atoi(e.ExtNumber)
		}
		offset, _ := strconv.ParseInt(e.Offset, 10, 64)
		valInt := 1000000000 + int64(extNum-1)*1000 + offset
		if e.Dir == "-" {
			valInt = -valInt
		}
		valStr := fmt.Sprintf("%d", valInt)
		return valStr, valInt
	}
	if e.Alias != "" {
		return CleanEnumName(e.Alias), 0
	}
	return "0", 0
}

func normalizeConstantValue(val string) string {
	val = strings.TrimSpace(val)
	// Replace C bitwise not expressions
	val = strings.ReplaceAll(val, "(~0ULL)", "0xFFFFFFFFFFFFFFFF")
	val = strings.ReplaceAll(val, "(~0ULL-1)", "0xFFFFFFFFFFFFFFFE")
	val = strings.ReplaceAll(val, "(~0ULL-2)", "0xFFFFFFFFFFFFFFFD")
	val = strings.ReplaceAll(val, "(~0U)", "0xFFFFFFFF")
	val = strings.ReplaceAll(val, "(~1U)", "0xFFFFFFFE")
	val = strings.ReplaceAll(val, "(~2U)", "0xFFFFFFFD")
	val = strings.ReplaceAll(val, "(~3U)", "0xFFFFFFFC")
	val = strings.ReplaceAll(val, "(~0U-1)", "0xFFFFFFFE")
	val = strings.ReplaceAll(val, "(~0U-2)", "0xFFFFFFFD")
	val = strings.ReplaceAll(val, "(~0U-3)", "0xFFFFFFFC")

	// If it's a quoted string, leave it alone
	if strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") {
		return val
	}

	// Strip "ULL", "UL", "U", "F", "f" suffixes
	if strings.HasSuffix(val, "ULL") || strings.HasSuffix(val, "ull") {
		val = val[:len(val)-3]
	} else if strings.HasSuffix(val, "UL") || strings.HasSuffix(val, "ul") {
		val = val[:len(val)-2]
	} else if strings.HasSuffix(val, "U") || strings.HasSuffix(val, "u") {
		val = val[:len(val)-1]
	} else if strings.HasSuffix(val, "F") || strings.HasSuffix(val, "f") {
		val = val[:len(val)-1]
	}

	return val
}

func ApiMatches(apiAttr string, targetApi string) bool {
	if apiAttr == "" {
		return true
	}
	for _, a := range strings.Split(apiAttr, ",") {
		if strings.TrimSpace(a) == targetApi {
			return true
		}
	}
	return false
}

func addEnumExtensionToBranch(branch *ApiBranchInfo, e parser.Enum, valStr string, valInt int64, cInfo *ConstantInfo) {
	if e.Extends != "" {
		bGroup := branch.EnumGroups[e.Extends]
		if bGroup == nil {
			bGroup = &EnumGroup{
				Name:   e.Extends,
				GoName: CleanTypeName(e.Extends),
				Type:   "enum",
			}
			branch.EnumGroups[e.Extends] = bGroup
		}
		bGroup.Values = append(bGroup.Values, EnumValue{
			Name:        e.Name,
			GoName:      CleanEnumName(e.Name),
			ValueStr:    valStr,
			ValueInt:    valInt,
			Comment:     e.Comment,
			IsAlias:     e.Alias != "",
			AliasOf:     CleanEnumName(e.Alias),
			ExtendsEnum: e.Extends,
			Api:         e.Api,
		})
	} else if cInfo != nil {
		branch.Constants[e.Name] = *cInfo
	}
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
