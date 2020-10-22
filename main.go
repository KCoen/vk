package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)
const CONFIG_DIR = "vk"
const CONFIG_PACKAGE = "vk"
const CONFIG_REGISTERY_FILE = "vk.xml"

var _ = strconv.AppendBool

var sprintf = fmt.Sprintf
var printf = fmt.Printf
var textcontext struct {
	at interface{}
}


func base_type(in string) string {
	switch in {
	case "int32_t": fallthrough
	case "int64_t": fallthrough
	case "uint64_t": fallthrough
	case "uint32_t": fallthrough
	case "int16_t":
		return string(in[:len(in)-2])
	case "char": fallthrough
	case "uint8_t":
		return "byte"
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "size_t":
		return "uint"
	case "*void":
		return "uintptr"
	default:
		return in
	}
}

func valid_lh(in string) string {
	switch in {
	case "range":
		fallthrough
	case "type":
		return in + "0"
	default:
		return in
	}
}
func textpanic(v interface{}) {
	fmt.Printf("in %v\n", textcontext.at)
	panic(v)
}
func select_one(vls ...string) (r string) {
	for _, v := range vls {
		if v != "" {
			if r != "" {
				textpanic(fmt.Sprintf("Expected one but got %v", vls))
			}
			r = v
		}
	}
	if r == "" {
		textpanic(fmt.Sprintf("Expected one but got none"))
	}
	return
}
func select_either(vls ...string) (r string) {
	for _, v := range vls {
		if v != "" {
			return v;
		}
	}
	textpanic(fmt.Sprintf("Expected one but got none"))
	return
}

func stubfunction() string {
	return ``
	return `{
	return;
}`;
}

func cleanvalue(in string) string {
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

func parsebitpos(in string) string {
	if in != "" {
		in = strings.Replace(in, "f", "", 0)
		in = strings.Replace(in, "d", "", 0)
		return "0x1 << " + in
	}
	return ""
}

func parse_basetype(m kronostype) string {
	if len(m.Type) > 0 {
		return fmt.Sprintf("type %s = %s\n", m.Name, base_type(m.Type[0]))
	} else {
		return fmt.Sprintf("type %s = unk\n", m.Name)
	}
}

func bitwidth_to_type(width int) string {
	switch width {
	case 32:
		return "int32";
	case 64:
		return "int64";
	case 16:
		return "int16";
	default:
		panic(width);
	}
}
func main() {
	d, _ := ioutil.ReadFile(CONFIG_REGISTERY_FILE)
	var reg Registry
	xml.Unmarshal(d, &reg)

	f, e := os.OpenFile(CONFIG_DIR + "/vulkan.go", os.O_TRUNC|os.O_CREATE, 0664)
	if e != nil {
		panic(e)
	}

	wf := func(s string, vals ... interface{}) {
		f.Write([]byte(sprintf(s, vals...)))
	}
	wl := func(ss ... string) {
		for _, s := range(ss) {
			f.Write([]byte(s))
			if s[len(s)-1] != '\n' {
				f.Write([]byte{'\n'})	
			}
		}
	}
	art := func(s string) {
		wl("//----------------------------------------")
		wf("//--%s\n", s)
		wl("//----------------------------------------\n")
	}
	doc_attr := func(ss ...string) {
		if len(ss) % 2 != 0 {
			panic(ss);
		}
		for i := 0;i < len(ss);i += 2{
			if ss[i+1] != "" {
				wf("//%s: %s\n", ss[i], ss[i+1])
			}
		}
	}
	doc_arguments := func(ss ...string) {
		if len(ss) % 2 != 0 {
			panic(ss);
		}
		for i := 0;i < len(ss);i += 2{
			if ss[i+1] != "" {
				wf("/*[%s]=%s*/", ss[i], ss[i+1])
			}
		}
	}

	wf("package %s\n\n", CONFIG_PACKAGE)
	art("Build-in")
	wl("type unk = int")
	wl("type external_type = unk")
	wl("type funcptr = uintptr")
	wl("type handle = uintptr")

	art("Types")

	for _, t := range reg.Types.Type {
		name := select_one(t.Name, t.AttrName)

		switch t.Category {
		case "basetype":
			if len(t.Type) > 0 {
				wf("type %s = %s\n", name, base_type(t.Type[0]))
			} else {
				wf("type %s = unk\n", name)
			}
		case "funcpointer":
			wf("type %s = funcptr\n", name)
		case "handle":
			wf("type %s = uint64\n", name)
		case "":
			if t.Requires != "" {
				wf("type %s = external_type\n", name) // Externally defined type
				continue;
			}
			fallthrough
		case "bitmask":
			if name == "int" {
				continue;
			}
			wf("type %s = uint32\n", name)
		case "union":
			wf("type %s = struct{uintptr}\n", name)
		case "enum":
			if t.Alias != "" {
				wf("type %s = %s\n", name, t.Alias)
			}
		case "struct":
			wf("type %s struct {\n", name)
			for _, m := range t.Member {
				// Fix *
				type_name := base_type(m.Type)
				for i := strings.Count(m.Text, "*"); i > 0; i-- {
					type_name = "*" + type_name
				}
				type_name = base_type(type_name)
				if strings.Count(m.Text, "[") != 0 && strings.Count(m.Text, "]") != 0 {
					if m.Enum != "" {
						type_name = "[" +m.Enum + "]" + type_name	
					}
				}
				wf("\t%s %s\n", valid_lh(m.Name), type_name)
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
		bitwidth := 32
		if e.Type == "bitmask" {
			bitwidth = 64;
		}
		if e.Name == "API Constants" {
			e.Name = ""
		} else {
			fmt.Fprintf(f, "type %s = %s\n", e.Name, bitwidth_to_type(bitwidth))
		}
		for _, ev := range e.Enum {
			textcontext.at = ev
			lh := ev.Name
			rh := select_one(cleanvalue(ev.Value), ev.Alias, parsebitpos(ev.Bitpos))
			fmt.Fprintf(f, "const %s %s = %s\n", lh, e.Name, rh)
		}
	}

	art("Functions")

	funlist := []*Fn{}

	for _,c := range(reg.Commands.Command) {
		doc_attr("Errors", c.Errorcodes, "On success", c.Successcodes, "queues", c.Queues, "renderpass", c.Renderpass, "pipeline", c.Pipeline)
		name := valid_lh(select_one(c.Name, c.Proto.Name))
		//wf("//golinkname: %s\n", name)
		f := Fn{
			Name: name,
			PrintTrace: false,
			dllname: "vulkan-1.dll",
			dllfuncname: name,
			Params: make([]*Param, len(c.Param)),
		}
		if c.Proto.Type != "" && c.Proto.Type != "void" {
			f.Rets = &Rets{
				Type: c.Proto.Type,
				Name: "ret0",
			}	
		} else {
			f.Rets = &Rets{}	
		}
		

		//wf("func %s(", name)
		for i, p :=range(c.Param) {
			param_type_name := base_type(p.Type)
			for i := strings.Count(p.Text, "*"); i > 0; i-- {
				param_type_name = "*" + param_type_name
			}
			param_type_name = base_type(param_type_name)
			param_name := valid_lh(p.Name)

			f.Params[i] = &Param{
				Name: param_name,
				Type: param_type_name,
				fn: &f,
				tmpVarIdx: i,
			}

			if i > 0 {
				//wf(",")	
			}			
			//wf("%s %s", valid_lh(p.Name), param_type_name)
			doc_arguments("optional", p.Optional, "len", p.Len)
		}
		if c.Proto.Type != "" && c.Proto.Type != "void" {
			//wf(") (ret %s) %s\n", c.Proto.Type, stubfunction());
		} else {
			//wf(") %s \n", stubfunction());
		}
		funlist = append(funlist, &f)
	}

	GenCalls(funlist)
}

type kronostype struct {
	AttrName      string `xml:"name,attr"`
	Category      string `xml:"category,attr"`
	Requires      string `xml:"requires,attr"`
	Alias         string `xml:"alias,attr"`
	Parent        string `xml:"parent,attr"`
	Returnedonly  string `xml:"returnedonly,attr"`
	AttrComment   string `xml:"comment,attr"`
	Structextends string `xml:"structextends,attr"`
	Inner         string `xml:",innerxml"`
	Text          string `xml:",chardata"`
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
			Text          string `xml:",chardata"`
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
				Text          string `xml:",chardata"`
				Optional       string `xml:"optional,attr"`
				Externsync     string `xml:"externsync,attr"`
				Len            string `xml:"len,attr"`
				Noautovalidity string `xml:"noautovalidity,attr"`
				Type           string `xml:"type"`
				Name           string `xml:"name"`
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
