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
	Scopes     map[string]string // All paths in the object @TODO Add attributes
	BitOffset  int
	FuncName   string
	Size       int

	ReturnType *VkTypeInfo // @TODO
	Arguments  []VkTypeInfo
}

var regexp_func *regexp.Regexp

var hasPrefix = strings.HasPrefix
var hasSuffix = strings.HasSuffix
var trimPrefix = strings.TrimPrefix

func init() {
	regexp_func = regexp.MustCompile(
		`\(VKAPI_PTR\s+\*(PFN_[A-z0-9]*)\)` + // Name
			`\(\s*(.*)\s*\)`) // Argument list
}

func parseTypeFromXmlString(xml_input string) (c VkTypeInfo) {
	c.Scopes = map[string]string{}

	is_simple_scope := func(path string) bool {
		for _, v := range c.Scopes {
			if hasPrefix(v, path) {
				return false
			}
		}
		return true
	}

	callback := func(full_path string, char byte) {
		c.Scopes[full_path] += string(char)

		// Replace `/foo/bar/BRACKET/fox/cuba` with `/foo/bar/`
		// Eg bracket forces nested elements to be demoted to text
		var path = full_path
		i := strings.Index(full_path, "/BRACKET")
		if i > 0 {
			path = full_path[:i]
		}

		depth := strings.Count(path, "/")

		if hasSuffix(path, "/type") {
			c.Simpletype += string(char)
			c.Fulltype += string(char)
		} else if is_any(path, "/member", "/param") || hasSuffix(path, "/type") {
			c.Fulltype += string(char)
		} else if depth < 3 && hasSuffix(path, "/name") {
			if strings.Contains(c.Fulltype, "typedef") {
				c.Fulltype += string(char)
			}
		}

		if depth < 3 && hasSuffix(path, "/name") {
			c.Name += string(char)
		}
		if depth < 3 && hasSuffix(path, "/comment") {
			c.Comment += string(char)
		}
	}

	// Basic Recursive decent XML Parser that calls the callback every character
	// This could potentially be replaced with an encoding/xml tokenizer
	var linear_xml_parse func(path string, d string) (i int)
	linear_xml_parse = func(path string, d string) (i int) {
		for ; i < len(d); i++ {
			if d[i] == '<' {
				name, tag, closing := readtag(d[i:])
				i += len(tag)
				if closing {
					return
				}
				if d[i-1] != '/' {
					i++
					i += linear_xml_parse(path+"/"+name, d[i:])
				}
			} else {
				// Add Pseudo XML path for paths that are contained with text brackets
				// Eg <type>asdf(<enum>stuff</enum)</type> becomes type/BRACKET/enum
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

	linear_xml_parse("", xml_input)

	// In the case of <type>struct <type>vkAlignedJelly<type></type> we pick /type/type over /type
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
			case hasPrefix(ctype, "const"):
				c.Isconst = true
				ctype = trimPrefix(ctype, "const")
			case hasPrefix(ctype, "typedef"):
				c.Istypedef = true
				ctype = trimPrefix(ctype, "typedef")
			case hasPrefix(ctype, "*"):
				gotype = "*" + gotype
				gotype = ctype_to_go(gotype)
				c.Ispointer++
				ctype = trimPrefix(ctype, "*")
			case c.Simpletype != "" && hasPrefix(ctype, c.Simpletype):
				gotype += c.Simpletype
				gotype = ctype_to_go(gotype)
				ctype = trimPrefix(ctype, c.Simpletype)
			case hasPrefix(ctype, "struct"):
				ctype = trimPrefix(ctype, "struct")
			case c.Name != "" && hasPrefix(ctype, c.Name):
				ctype = trimPrefix(ctype, c.Name)
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
		c.GoName = makeGoSafeIdentifier(c.Name)
		c.GoComment = c.Comment
		if c.BitOffset != 0 {
			c.GoType = strings.Replace(c.GoType, ctype_to_go(c.Simpletype), "struct{}", -1) // @TODO Make sure the padding is corrected on these structs
			c.GoComment = fmt.Sprintf("BITFIELD: %d %s", c.BitOffset, c.GoComment)
		}
	}

	// Parse functions as types  @TODO
	s := regexp_func.FindAllStringSubmatch(c.Fulltype, 5)
	if len(s) > 0 {
		c.FuncName = s[0][1]
		//ArgListString := s[0][2]
	}
	return c
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
