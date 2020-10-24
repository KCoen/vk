package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"go/format"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"cld.moe/vk/generator/util"
)

const CONFIG_DIR = ".."
const CONFIG_PACKAGE = "vk"
const CONFIG_REGISTERY_FILE = "vk.xml"

var sprintf = fmt.Sprintf
var printf = fmt.Printf
var sscanf = fmt.Sscanf
var select_one = util.Select_one
var select_either = util.Select_either
var is_any_byte = util.Is_any_byte
var is_any = util.Is_any

func main() {
	// Read File
	d, e := ioutil.ReadFile(CONFIG_REGISTERY_FILE)
	if e != nil {
		panic(e)
	}

	var reg Registry
	e = xml.Unmarshal(d, &reg)
	if e != nil {
		panic(e)
	}

	var f bytes.Buffer
	// Helper functions
	wf := func(s string, vals ...interface{}) {
		f.Write([]byte(sprintf(s, vals...)))
	}
	wl := func(ss ...string) {
		for _, s := range ss {
			f.Write([]byte(s))
			if len(s) < 1 || s[len(s)-1] != '\n' {
				f.Write([]byte{'\n'})
			}
		}
	}
	art := func(s string) {
		wl("//----------------------------------------")
		wf("//--%s\n", s)
		wl("//----------------------------------------\n")
	}

	wl("//Generated\n")
	wf("package %s\n\n", CONFIG_PACKAGE)

	art("Types")

	for _, t := range reg.Types.Type {
		name := select_one(t.Name, t.AttrName)

		switch t.Category {
		case "basetype":
			if len(t.Type) > 0 {
				wf("type %s = %s\n", name, ctype_to_go(t.Type[0]))
			} else {
				wf("type %s = unk\n", name)
			}
		case "funcpointer":
			wf("type %s = funcptr\n", name)
		case "handle":
			wf("type %s = uint64\n", name)
		case "":
			if t.Requires != "" {
				if is_buildin_type(name) == false {
					wf("type %s = external_type\n", name) // Externally defined type
				}
			}
		case "bitmask":
			if name == "int" {
				continue
			}

			vktype := parseTypeFromXmlString(fmt.Sprintf("<type>%s</type>", t.Xml))
			if vktype.Istypedef {
				wf("type %s = %s\n", name, vktype.GoType)
				if t.Requires != "" {
					wf("type %s = %s\n", t.Requires, vktype.GoType)
				}
			} else {
				wf("type %s = uint32\n", name)
			}
		case "union":
			wf("type %s = struct{uintptr}\n", name) //@TODO
		case "enum":
			if t.Alias != "" {
				wf("type %s = %s\n", name, t.Alias)
			}
		case "struct":
			wf("type %s struct {\n", name)
			for _, m := range t.Member {
				t := parseTypeFromXmlString(fmt.Sprintf("<member>%s</member>", m.Xml))
				wf("\t%s %s", t.GoName, t.GoType)
				if t.GoComment != "" {
					wf(" //%s", t.GoComment)
				}
				wl("")
			}
			wl("}\n")
		case "include": // Nop
		case "define": //@Todo
		default:
			printf("Unknown type category %s\n", t.Category)
		}
	}

	art("Enums")

	for _, e := range reg.Enums {
		bitwidth := 32 //??
		if e.Type == "bitmask" {
			bitwidth = 64 ///??
		}

		if e.Name == "API Constants" {
			e.Name = ""
		} else {
			if e.Type != "bitmask" {
				wf("type %s = %s\n", e.Name, bitwidth_to_type(bitwidth))
			}
		}
		for _, ev := range e.Enum {
			lh := ev.Name
			rh := select_one(fixNumberLiterals(ev.Value), ev.Alias, formatBitPos(ev.Bitpos))
			wf("const %s %s = %s\n", lh, e.Name, rh)
		}
	}

	art("Functions")

	funlist := []*Fn{}
	for _, c := range reg.Commands.Command {
		name := renameKeywords(select_one(c.Name, c.Proto.Name))

		f := Fn{
			Name:           name,
			PrintTrace:     false,
			dllname:        "vulkan-1.dll",
			dllfuncname:    name,
			Params:         make([]*Param, 0),
			comment:        c.Comment,
			successcodes:   strings.Split(c.Successcodes, ","),
			renderpass:     strings.Split(c.Renderpass, ","),
			pipeline:       strings.Split(c.Pipeline, ","),
			errorcodes:     strings.Split(c.Errorcodes, ","),
			cmdbufferlevel: strings.Split(c.Cmdbufferlevel, ","),
		}

		if c.Proto.Type != "" && c.Proto.Type != "void" {
			f.Rets = &Rets{
				Type: ctype_to_go(c.Proto.Type),
				Name: "ret0",
			}
		} else {
			f.Rets = &Rets{}
		}
		for i, p := range c.Param {
			vktype := parseTypeFromXmlString(sprintf("<param>%s</param>", string(p.Xml)))

			// If its sized array as argument,we unroll it
			if vktype.Size > 0 && vktype.Ispointer == 0 {
				for index := vktype.Size; index > 0; index-- {
					f.Params = append(f.Params, &Param{
						Name:      sprintf("%s%d", vktype.GoName, index),
						Type:      ctype_to_go(vktype.Simpletype),
						fn:        &f,
						tmpVarIdx: i,
					})
				}
			} else {
				f.Params = append(f.Params, &Param{
					Name:      vktype.GoName,
					Type:      vktype.GoType,
					fn:        &f,
					tmpVarIdx: i,
				})
			}
		}
		funlist = append(funlist, &f)
	}

	{
		formatedbytes, err := format.Source(f.Bytes())
		if err != nil {
			panic(err)
		}
		// Open Output
		file, e := os.OpenFile(CONFIG_DIR+"/vulkan.go", os.O_WRONLY|os.O_TRUNC, 0664)
		if e != nil {
			panic(e)
		}
		defer file.Close()

		file.Write(formatedbytes)
	}

	{
		src := Source{funlist}
		var w bytes.Buffer
		w.Write([]byte("//Generated\n\n"))
		w.Write([]byte("package " + CONFIG_PACKAGE + "\n"))

		var err error
		funcMap := template.FuncMap{
			"newlazydll": func(dll string) string {
				arg := "\"" + dll + ".dll\""
				return "windows.NewLazySystemDLL(" + arg + ")"
			},
		}
		t := template.Must(template.New("main").Funcs(funcMap).Parse(tsrc))
		err = t.Execute(&w, src)
		if err != nil {
			panic(err)
		}

		formatedbytes, err := format.Source(w.Bytes())
		if err != nil {
			panic(err)
		}

		file, e := os.OpenFile(CONFIG_DIR+"/vulkan_windows.go", os.O_TRUNC|os.O_WRONLY, 0664)
		if e != nil {
			panic(e)
		}
		defer file.Close()
		file.Write(formatedbytes)
	}
}

func renameKeywords(in string) string {
	t := token.Lookup(in)
	if t.IsKeyword() {
		return in + "0"
	}
	return in
}
func fixNumberLiterals(in string) string {
	// Clean .f from floats
	if in != "" {
		var a float64
		c, e := fmt.Sscanf(in, "%ff", &a)
		if e == nil && c == 1 {
			return fmt.Sprintf("%f", a)
		}
		// C++ specific expression @TODO
		if in[0] == '(' {
			return sprintf("\"%s\"", in)
		}
	}
	return in
}
func formatBitPos(in string) string {
	if in != "" {
		return "0x1 << " + in
	}
	return ""
}
func parseBaseType(m kronostype) string {
	if len(m.Type) > 0 {
		return fmt.Sprintf("type %s = %s\n", m.Name, ctype_to_go(m.Type[0]))
	} else {
		return fmt.Sprintf("type %s = unk\n", m.Name)
	}
}

func bitwidth_to_type(width int) string {
	switch width {
	case 32:
		return "int32"
	case 64:
		return "int64"
	case 16:
		return "int16"
	default:
		panic(width)
	}
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

type kronostype struct {
	AttrName       string   `xml:"name,attr"`
	Category       string   `xml:"category,attr"`
	Requires       string   `xml:"requires,attr"`
	Alias          string   `xml:"alias,attr"`
	Parent         string   `xml:"parent,attr"`
	Returnedonly   string   `xml:"returnedonly,attr"`
	AttrComment    string   `xml:"comment,attr"`
	Structextends  string   `xml:"structextends,attr"`
	Xml            string   `xml:",innerxml"`
	Text           string   `xml:",chardata"`
	Allowduplicate string   `xml:"allowduplicate,attr"`
	Name           string   `xml:"name"`
	Type           []string `xml:"type"`
	Member         []member `xml:"member"`
	Comment        []string `xml:"comment"`
}

type member struct {
	Noautovalidity string `xml:"noautovalidity,attr"`
	Values         string `xml:"values,attr"`
	Optional       string `xml:"optional,attr"`
	Len            string `xml:"len,attr"`
	Text           string `xml:",chardata"`
	Altlen         string `xml:"altlen,attr"`
	Externsync     string `xml:"externsync,attr"`
	Selection      string `xml:"selection,attr"`
	Selector       string `xml:"selector,attr"`
	Type           string `xml:"type"`
	Name           string `xml:"name"`
	Enum           string `xml:"enum"`
	Comment        string `xml:"comment"`
	Xml            string `xml:",innerxml"`
}

type Registry struct {
	XMLName   xml.Name `xml:"registry"`
	Comment   []string `xml:"comment"`
	Platforms struct {
		Comment  string `xml:"comment,attr"`
		Platform []struct {
			Name    string `xml:"name,attr"`
			Protect string `xml:"protect,attr"`
			Comment string `xml:"comment,attr"`
		} `xml:"platform"`
	} `xml:"platforms"`
	Tags struct {
		Comment string `xml:"comment,attr"`
		Tag     []struct {
			Name    string `xml:"name,attr"`
			Author  string `xml:"author,attr"`
			Contact string `xml:"contact,attr"`
		} `xml:"tag"`
	} `xml:"tags"`
	Types struct {
		AttrComment string       `xml:"comment,attr"`
		Type        []kronostype `xml:"type"`
		Comment     []string     `xml:"comment"`
	} `xml:"types"`
	Enums []struct {
		Name        string `xml:"name,attr"`
		AttrComment string `xml:"comment,attr"`
		Type        string `xml:"type,attr"`
		Enum        []struct {
			Value   string `xml:"value,attr"`
			Name    string `xml:"name,attr"`
			Alias   string `xml:"alias,attr"`
			Comment string `xml:"comment,attr"`
			Bitpos  string `xml:"bitpos,attr"`
		} `xml:"enum"`
		Comment []string `xml:"comment"`
		Unused  struct {
			Start   string `xml:"start,attr"`
			Comment string `xml:"comment,attr"`
		} `xml:"unused"`
	} `xml:"enums"`
	Commands struct {
		Comment string `xml:"comment,attr"`
		Command []struct {
			Text           string `xml:",chardata"`
			Successcodes   string `xml:"successcodes,attr"`
			Errorcodes     string `xml:"errorcodes,attr"`
			Queues         string `xml:"queues,attr"`
			Name           string `xml:"name,attr"`
			Alias          string `xml:"alias,attr"`
			Renderpass     string `xml:"renderpass,attr"`
			Cmdbufferlevel string `xml:"cmdbufferlevel,attr"`
			Pipeline       string `xml:"pipeline,attr"`
			Comment        string `xml:"comment,attr"`
			Proto          struct {
				Type string `xml:"type"`
				Name string `xml:"name"`
			} `xml:"proto"`
			Param []struct {
				Text           string `xml:",chardata"`
				Optional       string `xml:"optional,attr"`
				Externsync     string `xml:"externsync,attr"`
				Len            string `xml:"len,attr"`
				Noautovalidity string `xml:"noautovalidity,attr"`
				Type           string `xml:"type"`
				Name           string `xml:"name"`
				Xml            string `xml:",innerxml"`
			} `xml:"param"`
			Implicitexternsyncparams struct {
				Param string `xml:"param"`
			} `xml:"implicitexternsyncparams"`
		} `xml:"command"`
	} `xml:"commands"`
	Feature []struct {
		Api     string `xml:"api,attr"`
		Name    string `xml:"name,attr"`
		Number  string `xml:"number,attr"`
		Comment string `xml:"comment,attr"`
		Require []struct {
			AttrComment string `xml:"comment,attr"`
			Type        []struct {
				Name    string `xml:"name,attr"`
				Comment string `xml:"comment,attr"`
			} `xml:"type"`
			Enum []struct {
				Name      string `xml:"name,attr"`
				Extends   string `xml:"extends,attr"`
				Extnumber string `xml:"extnumber,attr"`
				Offset    string `xml:"offset,attr"`
				Bitpos    string `xml:"bitpos,attr"`
				Alias     string `xml:"alias,attr"`
				Comment   string `xml:"comment,attr"`
				Dir       string `xml:"dir,attr"`
				Value     string `xml:"value,attr"`
			} `xml:"enum"`
			Command []struct {
				Name string `xml:"name,attr"`
			} `xml:"command"`
			Comment []string `xml:"comment"`
		} `xml:"require"`
	} `xml:"feature"`
	Extensions struct {
		Comment   string `xml:"comment,attr"`
		Extension []struct {
			Name         string `xml:"name,attr"`
			Number       string `xml:"number,attr"`
			Type         string `xml:"type,attr"`
			Author       string `xml:"author,attr"`
			Contact      string `xml:"contact,attr"`
			Supported    string `xml:"supported,attr"`
			Requires     string `xml:"requires,attr"`
			Platform     string `xml:"platform,attr"`
			Comment      string `xml:"comment,attr"`
			Specialuse   string `xml:"specialuse,attr"`
			Deprecatedby string `xml:"deprecatedby,attr"`
			Promotedto   string `xml:"promotedto,attr"`
			Obsoletedby  string `xml:"obsoletedby,attr"`
			Provisional  string `xml:"provisional,attr"`
			Sortorder    string `xml:"sortorder,attr"`
			RequiresCore string `xml:"requiresCore,attr"`
			Require      []struct {
				Feature   string `xml:"feature,attr"`
				Extension string `xml:"extension,attr"`
				Enum      []struct {
					Value     string `xml:"value,attr"`
					Name      string `xml:"name,attr"`
					Offset    string `xml:"offset,attr"`
					Extends   string `xml:"extends,attr"`
					Dir       string `xml:"dir,attr"`
					Comment   string `xml:"comment,attr"`
					Extnumber string `xml:"extnumber,attr"`
					Bitpos    string `xml:"bitpos,attr"`
					Alias     string `xml:"alias,attr"`
				} `xml:"enum"`
				Type []struct {
					Name string `xml:"name,attr"`
				} `xml:"type"`
				Command []struct {
					Name string `xml:"name,attr"`
				} `xml:"command"`
				Comment string `xml:"comment"`
			} `xml:"require"`
		} `xml:"extension"`
	} `xml:"extensions"`
}
