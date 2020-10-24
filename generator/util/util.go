package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

// Helper functions
func Clamp(a, mi, ma int) int {
	return Min(Max(a, mi), ma)
}
func Max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Select_one(vls ...string) (r string) {
	for _, v := range vls {
		if v != "" {
			if r != "" {
				panic(fmt.Sprintf("Expected one but got %v", vls))
			}
			r = v
		}
	}
	if r == "" {
		panic(fmt.Sprintf("Expected one but got none"))
	}
	return
}

// Is any value in a equal to any value in vals
func Is_any(a string, vals ...string) bool {
	for _, v := range vals {
		if a == v {
			return true
		}
	}
	return false
}

// Is any value in a equal to any value in vals
func Is_any_byte(a byte, vals ...byte) bool {
	for _, v := range vals {
		if a == v {
			return true
		}
	}
	return false
}
func Select_either(vls ...string) (r string) {
	for _, v := range vls {
		if v != "" {
			return v
		}
	}
	panic(fmt.Sprintf("Expected one but got none"))
	return
}

func Vardump(i interface{}) (ret string) {
	write := func(format string, vals ...interface{}) {
		ret += fmt.Sprintf(format, vals...)
	}
	indent := func(depth int) string {
		s := ""
		for i := 0; i < depth; i++ {
			s += "\t"
		}
		return s
	}
	var recurse func(depth int, name string, original reflect.Value)
	recurse = func(depth int, name string, original reflect.Value) {
		switch original.Kind() {
		case reflect.Ptr:
			ov := original.Elem()
			if !ov.IsValid() {
				return
			}
			recurse(depth+1, name, ov)
		case reflect.Interface:
			originalValue := original.Elem()
			recurse(depth+1, "??", originalValue)
		case reflect.Struct:
			for i := 0; i < original.NumField(); i++ {
				recurse(depth+1, original.Type().Field(i).Name, original.Field(i))
			}
		case reflect.Slice:
			fallthrough
		case reflect.Array:
			for i := 0; i < original.Len(); i++ {
				recurse(depth+1, fmt.Sprintf("%d", i), original.Index(i))
			}
		case reflect.Map:
			for _, key := range original.MapKeys() {
				recurse(depth+1, key.String(), original.MapIndex(key))
			}
		default:
			write("%s%s:%#v\n", indent(depth), name, original)
		}
	}
	v := reflect.ValueOf(i)
	write("%s\n", reflect.TypeOf(i).Name())
	recurse(0, reflect.TypeOf(i).Name(), v)
	return
}
func Vardump_json(i interface{}) string {
	var a bytes.Buffer
	e := json.NewEncoder(&a)
	e.SetEscapeHTML(false)
	e.SetIndent("", "\t")
	er := e.Encode(i)
	if er != nil {
		panic(er)
	}
	return string(a.Bytes())
}
