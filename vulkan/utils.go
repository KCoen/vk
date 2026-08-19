package vulkan

import (
	"bytes"
	"unsafe"
)

// StringToNullTerminated converts a Go string into a pointer to a null-terminated C string.
func StringToNullTerminated(s string) *byte {
	if s == "" {
		return nil
	}
	b := make([]byte, len(s)+1)
	copy(b, s)
	b[len(s)] = 0
	return &b[0]
}

// NullTerminatedToString reads a null-terminated C string into a Go string.
func NullTerminatedToString(ptr *byte) string {
	if ptr == nil {
		return ""
	}
	var length int
	curr := ptr
	for *curr != 0 {
		length++
		curr = (*byte)(unsafe.Add(unsafe.Pointer(curr), 1))
	}
	return string(unsafe.Slice(ptr, length))
}

// StringSliceToNullTerminated converts a Go []string into a slice of null-terminated C string pointers (**byte).
func StringSliceToNullTerminated(slice []string) **byte {
	if len(slice) == 0 {
		return nil
	}
	ptrs := make([]*byte, len(slice))
	for i, s := range slice {
		ptrs[i] = StringToNullTerminated(s)
	}
	return &ptrs[0]
}

// SliceData returns a pointer to the first element of slice, or nil if empty.
func SliceData[T any](s []T) *T {
	if len(s) == 0 {
		return nil
	}
	return &s[0]
}

// ByteSliceToString converts a fixed null-terminated byte slice/array to Go string.
func ByteSliceToString(b []byte) string {
	if n := bytes.IndexByte(b, 0); n >= 0 {
		return string(b[:n])
	}
	return string(b)
}

// VulkanStruct is implemented by all high-level Vulkan structures that can be converted to C-ABI layout.
type VulkanStruct interface {
	rawPointer() unsafe.Pointer
}

// ExtractRawPointer converts a high-level Vulkan struct, raw struct pointer, or unsafe.Pointer to an unsafe.Pointer for C-ABI pNext chains.
func ExtractRawPointer(v any) unsafe.Pointer {
	if v == nil {
		return nil
	}
	if vs, ok := v.(VulkanStruct); ok {
		return vs.rawPointer()
	}
	switch val := v.(type) {
	case unsafe.Pointer:
		return val
	case uintptr:
		return unsafe.Pointer(val)
	}
	return nil
}

// ExtractPointer extracts an unsafe.Pointer from any pointer or struct pointer.
func ExtractPointer(v any) unsafe.Pointer {
	return ExtractRawPointer(v)
}
