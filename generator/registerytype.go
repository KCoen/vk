package main

import (
	"fmt"
	"regexp"
	"strings"
)

type VkTypeInfo struct { //@TODO Some of the members are public just for vardump reason, clean them up
	Simpletype string // Simple type as defined by the XML
	Fulltype   string // The full C type
	GoType     string // Go Should use this type
	GoName     string // Go should use this name
	Name       string
	Comment    string
	GoComment  string
	Isconst    bool
	Istypedef  bool              // XML defined the type using a typedef
	Ispointer  int               // If pointer of any kind
	IsCallback bool              // If typedef of a PFN_ function
	Scopes     map[string]string // All paths in the object @TODO Add attributes
	BitOffset  int
	Xml        string
	FuncName   string
	Size       int

	ReturnType *VkTypeInfo // @TODO
	Arguments  []VkTypeInfo
}

var regexp_func *regexp.Regexp

func init() {
	regexp_func = regexp.MustCompile(
		`\(VKAPI_PTR\s+\*(PFN_[A-z0-9]*)\)` + // Name
			`\(\s*(.*)\s*\)`) // Argument list
}

func parseTypeFromXmlString(_xml string) VkTypeInfo {
	var c VkTypeInfo
	c.Scopes = map[string]string{}
	c.Xml = _xml

	is_simple_scope := func(path string) bool {
		for _, v := range c.Scopes {
			if strings.HasPrefix(v, path) {
				return false
			}
		}
		return true
	}

	callback := func(full_path string, char byte) {
		c.Scopes[full_path] += string(char)

		// Replace /foo/bar/bracket/fox/cuba
		// with /foo/bar/cuba
		// Eg bracket forces nested elements to be demoted to text
		var path = full_path
		i := strings.Index(full_path, "/BRACKET")
		if i > 0 {
			path = full_path[:i]
		}

		depth := strings.Count(path, "/")

		if strings.HasSuffix(path, "/type") {
			c.Simpletype += string(char)
			c.Fulltype += string(char)
		} else if is_any(path, "/member", "/param") || strings.HasSuffix(path, "/type") {
			c.Fulltype += string(char)
		} else if depth < 3 && strings.HasSuffix(path, "/name") {
			if strings.Contains(c.Fulltype, "typedef") {
				c.Fulltype += string(char)
			}
		}

		if depth < 3 && strings.HasSuffix(path, "/name") {
			c.Name += string(char)
		}
		if depth < 3 && strings.HasSuffix(path, "/comment") {
			c.Comment += string(char)
		}
	}

	var linear_xml_parse func(path string, d string) (i int)
	linear_xml_parse = func(path string, d string) (i int) {
		for ; i < len(d); i++ {
			if d[i] == '<' {
				name, tag, closing := readtag(d[i:])
				i += len(tag)
				if closing {
					return
				}
				if d[i-1] != '/' { // Short tag notation
					i++ // Go Into Contents
					i += linear_xml_parse(path+"/"+name, d[i:])
				}
			} else {
				// Add Pseudo XML path for <type>asdf</type>(<enum>stuff</enum) xml layouts parsing
				// @TODO maybe we can just do this with the xml lib
				if is_any_byte(d[i], '[', '(') {
					path += "/BRACKET"
				}
				if !(is_any_byte(d[i], '\t', '\n', '\r', ';') ||
					(d[i] == ' ' && d[i-1] == ' ')) {
					callback(path, d[i])
				}
				if is_any_byte(d[i], ']', ')') {
					path = strings.TrimSuffix(path, "/BRACKET")
				}
			}
		}
		return
	}

	linear_xml_parse("", _xml)

	// Assume that when /type contains non-whitespace characters and /type/type also contains characters the simple type is /type/type
	_, ok := c.Scopes["/type/type"]
	if ok {
		if len(c.Scopes["/type"]) > 0 && is_simple_scope("/type/type") {
			c.Simpletype = c.Scopes["/type/type"]
		}
	}

	// Parse left to right deconstructing the C type and constructing the go type
	{
		gotype := ""
		ctype := c.Fulltype
		for len(ctype) > 0 {
			ctype = strings.Trim(ctype, "\t\n\r ")
			switch {
			case strings.HasPrefix(ctype, "const"):
				c.Isconst = true
				ctype = strings.TrimPrefix(ctype, "const")
			case strings.HasPrefix(ctype, "typedef"):
				c.Istypedef = true
				ctype = strings.TrimPrefix(ctype, "typedef")
			case strings.HasPrefix(ctype, "*"):
				gotype = "*" + gotype
				gotype = ctype_to_go(gotype)
				c.Ispointer++
				ctype = strings.TrimPrefix(ctype, "*")
			case c.Simpletype != "" && strings.HasPrefix(ctype, c.Simpletype):
				gotype += c.Simpletype
				gotype = ctype_to_go(gotype)
				ctype = strings.TrimPrefix(ctype, c.Simpletype)
			case strings.HasPrefix(ctype, "struct"):
				ctype = strings.TrimPrefix(ctype, "struct")
			case c.Name != "" && strings.HasPrefix(ctype, c.Name):
				ctype = strings.TrimPrefix(ctype, c.Name)
			default:
				named_size := ""

				c0, _ := sscanf(ctype, "[%d]", &c.Size)
				c1, _ := sscanf(ctype, "[%s]", &named_size)
				c2, _ := sscanf(ctype, ":%d", &c.BitOffset)

				if c0+c1+c2 == 0 {
					println("Din't know what to do with: " + ctype)
				}
				if c0+c1 > 0 {
					gotype = ctype + gotype
				}

				ctype = ""
			}
		}

		// Reconstruct Go Type
		c.GoType = gotype
		c.GoName = renameKeywords(c.Name)
		c.GoComment = c.Comment
		if c.BitOffset != 0 {
			c.GoType = strings.Replace(c.GoType, ctype_to_go(c.Simpletype), "struct{}", -1) // @TODO Make sure the padding is corrected on these structs
			c.GoComment = fmt.Sprintf("BITFIELD: %d %s", c.BitOffset, c.GoComment)
		}
	}

	// Parse function type definitions @TODO
	s := regexp_func.FindAllStringSubmatch(c.Fulltype, 5)
	if len(s) > 0 {
		c.FuncName = s[0][1]
		//ArgListString := s[0][2]
	}
	return c
}

func is_buildin_type(in string) bool {
	return ctype_to_go(in) != in
}
func ctype_to_go(in string) string {
	switch in {
	case "*void":
		return "uintptr"
	default:
		switch strings.TrimLeft(in, "*") {
		case "int_t":
			fallthrough
		case "int":
			return "int32"
		case "int16_t":
			fallthrough
		case "uint16_t":
			fallthrough
		case "int32_t":
			fallthrough
		case "int64_t":
			fallthrough
		case "uint64_t":
			fallthrough
		case "int8_t":
			fallthrough
		case "uint8_t":
			fallthrough
		case "uint32_t":
			return string(in[:len(in)-2])
		case "char":
			return "byte"
		case "float":
			return "float32"
		case "double":
			return "float64"
		case "size_t":
			return "uint"
		default:
			return in
		}
	}
}

// This function read out a tag eg <foobar>, end_tag is set to 1 in the case of </foobar> or <foobar/>
func readtag(d string) (name string, tag string, end_tag bool) {
	if d[0] != '<' {
		panic(0)
	}

	i := 0
	is_end := d[1] == '/' // Closing tag?
	if is_end {
		for ; d[i] != '>'; i++ {
		} // Walk to end of tag
		return "", d[:i], true
	}

	for ; d[i] != '>'; i++ {
		if d[i] == ' ' && name == "" { // @TODO Unicode& Other whitespace
			name = d[1:i] // Slice tag contents
		}
	}
	if name == "" {
		name = d[1:i] // Slice tag contents
	}
	return name, d[:i], d[i-1] == '/'
}
