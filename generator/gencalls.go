package main

import (
	"fmt"
	"strconv"
	"strings"

	"cld.moe/vk/generator/util"
)

// This File is partially based on https://github.com/golang/sys/blob/master/windows/syscall.go
// Aim to make it less generic and more specific so the code becomes less unwieldy for later ports to other OS's that are not windows or solaris

type Param struct {
	Name      string
	Type      string
	fn        *Fn
	tmpVarIdx int
}
type Fn struct {
	Name           string
	Params         []*Param
	Rets           *Rets
	PrintTrace     bool
	dllname        string
	dllfuncname    string
	comment        string
	renderpass     []string
	cmdbufferlevel []string
	pipeline       []string
	successcodes   []string
	errorcodes     []string
	// TODO: get rid of this field and just use parameter index instead
	curTmpVarIdx int // insure tmp variables have uniq names
}

// Rets describes function return parameters.
type Rets struct {
	Name         string
	Type         string
	ReturnsError bool
}
type Source struct {
	Funcs []*Fn
}

//join concatenates parameters ps into a string with sep separator.
// Each parameter is converted into string by applying fn to it
// before conversion.
func join(ps []*Param, fn func(*Param) string, sep string) string {
	if len(ps) == 0 {
		return ""
	}
	a := make([]string, 0)
	for _, p := range ps {
		a = append(a, fn(p))
	}
	return strings.Join(a, sep)
}

// tmpVar returns temp variable name that will be used to represent p during syscall.
func (p *Param) tmpVar() string {
	if p.tmpVarIdx < 0 {
		p.tmpVarIdx = p.fn.curTmpVarIdx
		p.fn.curTmpVarIdx++
	}
	return fmt.Sprintf("_p%d", p.tmpVarIdx)
}

// SyscallArgList returns source code fragments representing p parameter
// in syscall. Slices are translated into 2 syscall parameters: pointer to
// the first element and length.
func (p *Param) SyscallArgList() []string {
	t := p.Type
	var s string
	switch {
	case t == "*bool":
		s = fmt.Sprintf("unsafe.Pointer(&%s)", p.tmpVar())
	case t[0] == '*':
		s = fmt.Sprintf("unsafe.Pointer(%s)", p.Name)
	case t == "bool":
		s = p.tmpVar()
	case t == "Coord":
		// Convert a COORD into a uintptr (by fooling the type system). This code
		// assumes the two SHORTs are correctly laid out; the "cast" to uint32 is
		// just to get a pointer to pass.
		s = fmt.Sprintf("*((*uint32)(unsafe.Pointer(&%s)))", p.Name)
	case strings.HasPrefix(t, "[]"):
		return []string{
			fmt.Sprintf("uintptr(unsafe.Pointer(%s))", p.tmpVar()),
			fmt.Sprintf("uintptr(len(%s))", p.Name),
		}
	default:
		s = p.Name
	}
	return []string{fmt.Sprintf("uintptr(%s)", s)}
}

func hasnonempty(set []string) bool {
	for _, v := range set {
		if v != "" {
			return true
		}
	}
	return false
}

func (f *Fn) Comment() (r string) {
	if f.comment != "" {
		r += "//" + f.comment + "\n"
	}
	if hasnonempty(f.errorcodes) {
		r += "//Errors: " + strings.Join(f.errorcodes, ", ") + "\n"
	}
	if hasnonempty(f.successcodes) {
		r += "//On Success: " + strings.Join(f.successcodes, ", ") + "\n"
	}
	return r[:util.Max(len(r)-1, 0)]
}

// ParamList returns source code for function f parameters.
func (f *Fn) ParamList() string {
	return join(f.Params, func(p *Param) string { return p.Name + " " + p.Type }, ", ")
}

// ParamPrintList returns source code of trace printing part correspondent
// to syscall input parameters.
func (f *Fn) ParamPrintList() string {
	return join(f.Params, func(p *Param) string { return fmt.Sprintf(`"%s=", %s, `, p.Name, p.Name) }, `", ", `)
}

// SyscallParamList returns source code for SyscallX parameters for function f.
func (f *Fn) SyscallParamList() string {
	a := make([]string, 0)
	for _, p := range f.Params {
		a = append(a, p.SyscallArgList()...)
	}
	for len(a) < f.SyscallParamCount() {
		a = append(a, "0")
	}
	return strings.Join(a, ", ")
}

// ParamCount return number of syscall parameters for function f.
func (f *Fn) ParamCount() int {
	n := 0
	for _, p := range f.Params {
		n += len(p.SyscallArgList())
	}
	return n
}

// Syscall determines which SyscallX function to use for function f.
func (f *Fn) Syscall() string {
	c := f.SyscallParamCount()
	if c == 3 {
		return "syscall.Syscall"
	}
	return "syscall.Syscall" + strconv.Itoa(c)
}

// SyscallParamCount determines which version of Syscall/Syscall6/Syscall9/...
// to use. It returns parameter count for correspondent SyscallX function.
func (f *Fn) SyscallParamCount() int {
	n := f.ParamCount()
	switch {
	case n <= 3:
		return 3
	case n <= 6:
		return 6
	case n <= 9:
		return 9
	case n <= 12:
		return 12
	case n <= 15:
		return 15
	default:
		panic("too many arguments to system call")
	}
}

// HelperParamList returns source code for helper function f parameters.
func (f *Fn) HelperParamList() string {
	return join(f.Params, func(p *Param) string { return p.Name + " " + p.Type }, ", ")
}
func (f *Fn) HelperName() string {
	return f.Name
}

// DLLName returns DLL function name for function f.
func (f *Fn) DLLFuncName() string {
	if f.dllfuncname == "" {
		return f.Name
	}
	return f.dllfuncname
}

// ToParams converts r into slice of *Param.
func (r *Rets) ToParams() []*Param {
	ps := make([]*Param, 0)
	if len(r.Name) > 0 {
		ps = append(ps, &Param{Name: r.Name, Type: r.Type})
	}
	if false {
		ps = append(ps, &Param{Name: "err", Type: "error"})
	}
	return ps
}

// List returns source code of syscall return parameters.
func (r *Rets) List() string {
	s := join(r.ToParams(), func(p *Param) string { return p.Name + " " + p.Type }, ", ")
	if len(s) > 0 {
		s = "(" + s + ")"
	}
	return s
}

// BoolTmpVarCode returns source code for bool temp variable.
func (p *Param) BoolTmpVarCode() string {
	const code = `var %[1]s uint32
		if %[2]s {
			%[1]s = 1
		}`
	return fmt.Sprintf(code, p.tmpVar(), p.Name)
}

// BoolPointerTmpVarCode returns source code for bool temp variable.
func (p *Param) BoolPointerTmpVarCode() string {
	const code = `var %[1]s uint32
		if *%[2]s {
			%[1]s = 1
		}`
	return fmt.Sprintf(code, p.tmpVar(), p.Name)
}

// SliceTmpVarCode returns source code for slice temp variable.
func (p *Param) SliceTmpVarCode() string {
	const code = `var %s *%s
		if len(%s) > 0 {
			%s = &%s[0]
		}`
	tmp := p.tmpVar()
	return fmt.Sprintf(code, tmp, p.Type[2:], p.Name, tmp, p.Name)
}

// TmpVarCode returns source code for temp variable.
func (p *Param) TmpVarCode() string {
	switch {
	case p.Type == "bool":
		return p.BoolTmpVarCode()
	case p.Type == "*bool":
		return p.BoolPointerTmpVarCode()
	case strings.HasPrefix(p.Type, "[]"):
		return p.SliceTmpVarCode()
	default:
		return ""
	}
}

// SetReturnValuesCode returns source code that accepts syscall return values.
func (r *Rets) SetReturnValuesCode() string {
	if r.Name == "" && !r.ReturnsError {
		return ""
	}
	retvar := "r0"
	if r.Name == "" {
		retvar = "r1"
	}
	errvar := "_"
	if r.ReturnsError {
		errvar = "e1"
	}
	return fmt.Sprintf("%s, _, %s := ", retvar, errvar)
}

// TmpVarReadbackCode returns source code for reading back the temp variable into the original variable.
func (p *Param) TmpVarReadbackCode() string {
	switch {
	case p.Type == "*bool":
		return fmt.Sprintf("*%s = %s != 0", p.Name, p.tmpVar())
	default:
		return ""
	}
}
func (r *Rets) useLongHandleErrorCode(retvar string) string {
	const code = `if %s {
		err = errnoErr(e1)
	}`
	cond := retvar + " == 0"
	return fmt.Sprintf(code, cond)
}

// SetErrorCode returns source code that sets return parameters.
func (r *Rets) SetErrorCode() string {
	const code = `if r0 != 0 {
		%s = %sErrno(r0)
	}`
	if len(r.Type) == 0 {
		return ""
	}
	if r.Name == "" && !r.ReturnsError {
		return ""
	}
	if r.Name == "" {
		return r.useLongHandleErrorCode("r1")
	}
	if r.Type == "error" {
		return fmt.Sprintf(code, r.Name, "syscall.")
	}
	s := ""
	switch {
	case len(r.Type) == 0:
		s = fmt.Sprintf("%s = %s(r0)", r.Name, r.Type)
	case r.Type[0] == '*':
		s = fmt.Sprintf("%s = (%s)(unsafe.Pointer(r0))", r.Name, r.Type)
	case r.Type == "bool":
		s = fmt.Sprintf("%s = r0 != 0", r.Name)
	default:
		s = fmt.Sprintf("%s = %s(r0)", r.Name, r.Type)
	}
	if !r.ReturnsError {
		return s
	}
	return s + "\n\t" + r.useLongHandleErrorCode(r.Name)
}

var tsrc string = `
{{define "main"}}

import(
	"unsafe"
	"syscall"
	"golang.org/x/sys/windows"
)
var _ unsafe.Pointer

var(
	vulkan = windows.NewLazySystemDLL("vulkan-1.dll")
{{template "funcnames" .}}
)

{{range .Funcs}}{{template "funcbody" .}}{{end}}
{{end}}

{{define "funcnames"}}{{range .Funcs}}	proc{{.DLLFuncName}} = vulkan.NewProc("{{.DLLFuncName}}")
{{end}}{{end}}

{{define "funcbody"}}
{{.Comment}}
func {{.HelperName}}({{.HelperParamList}}) {{template "results" .}}{
{{template "tmpvars" .}}	{{template "syscall" .}}	{{template "tmpvarsreadback" .}}
{{template "seterror" .}}	return
}
{{end}}


{{define "tmpvars"}}{{range .Params}}{{if .TmpVarCode}}	{{.TmpVarCode}}
{{end}}{{end}}{{end}}

{{define "results"}}{{if .Rets.List}}{{.Rets.List}} {{end}}{{end}}

{{define "syscall"}}{{.Rets.SetReturnValuesCode}}{{.Syscall}}(proc{{.DLLFuncName}}.Addr(), {{.ParamCount}}, {{.SyscallParamList}}){{end}}

{{define "tmpvarsreadback"}}{{range .Params}}{{if .TmpVarReadbackCode}}
{{.TmpVarReadbackCode}}{{end}}{{end}}{{end}}

{{define "seterror"}}{{if .Rets.SetErrorCode}}	{{.Rets.SetErrorCode}}
{{end}}{{end}}
`
