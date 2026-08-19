package emitter

import "fmt"

func EmitBranchErrors(pkgName string) string {
	return fmt.Sprintf(`package %s

// Error implements the Go error interface for Result.
func (r Result) Error() string {
	return r.String()
}

// Success returns true if the result is >= SUCCESS (meaning not an error).
func (r Result) Success() bool {
	return r >= SUCCESS
}

// ToError returns nil if the result is SUCCESS, otherwise returns the Result as error.
func (r Result) ToError() error {
	if r == SUCCESS {
		return nil
	}
	return r
}
`, pkgName)
}

func EmitBranchLoader(pkgName string) string {
	return fmt.Sprintf(`package %s

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

var (
	vulkanLib                uintptr
	vkGetInstanceProcAddrPtr uintptr
	vkGetDeviceProcAddrPtr   uintptr
	customProcLoader         func(name string) uintptr
	loaderMu                 sync.RWMutex
)

// Init initializes the Vulkan library by dynamically loading the Vulkan loader without CGO.
func Init() error {
	loaderMu.Lock()
	defer loaderMu.Unlock()

	if vulkanLib != 0 || customProcLoader != nil {
		return nil
	}

	libNames := getVulkanLibraryNames()
	var lastErr error
	for _, name := range libNames {
		h, err := purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err == nil && h != 0 {
			vulkanLib = h
			break
		}
		lastErr = err
	}

	if vulkanLib == 0 {
		return fmt.Errorf("failed to load Vulkan library (%%v): %%w", libNames, lastErr)
	}

	addr, err := purego.Dlsym(vulkanLib, "vkGetInstanceProcAddr")
	if err != nil || addr == 0 {
		return errors.New("failed to resolve vkGetInstanceProcAddr")
	}
	vkGetInstanceProcAddrPtr = addr
	return nil
}

// InitWithLoader allows using a custom proc address loader function (such as from GLFW or SDL).
func InitWithLoader(loader func(name string) uintptr) error {
	if loader == nil {
		return errors.New("loader cannot be nil")
	}
	loaderMu.Lock()
	defer loaderMu.Unlock()
	customProcLoader = loader
	addr := loader("vkGetInstanceProcAddr")
	if addr != 0 {
		vkGetInstanceProcAddrPtr = addr
	}
	return nil
}

// GetInstanceProcAddr resolves a Vulkan procedure address for an instance.
func GetInstanceProcAddr(instance Instance, name string) uintptr {
	loaderMu.RLock()
	custom := customProcLoader
	gpa := vkGetInstanceProcAddrPtr
	lib := vulkanLib
	loaderMu.RUnlock()

	if custom != nil {
		if addr := custom(name); addr != 0 {
			return addr
		}
	}

	if gpa != 0 {
		cName := StringToNullTerminated(name)
		r1, _, _ := purego.SyscallN(gpa, uintptr(instance), uintptr(unsafe.Pointer(cName)))
		if r1 != 0 {
			return r1
		}
	}

	if lib != 0 {
		if addr, err := purego.Dlsym(lib, name); err == nil && addr != 0 {
			return addr
		}
	}
	return 0
}

// GetDeviceProcAddr resolves a Vulkan procedure address for a device.
func GetDeviceProcAddr(device Device, name string) uintptr {
	loaderMu.RLock()
	gpa := vkGetDeviceProcAddrPtr
	loaderMu.RUnlock()

	if gpa == 0 {
		gpa = GetInstanceProcAddr(0, "vkGetDeviceProcAddr")
		if gpa != 0 {
			loaderMu.Lock()
			vkGetDeviceProcAddrPtr = gpa
			loaderMu.Unlock()
		}
	}

	if gpa != 0 {
		cName := StringToNullTerminated(name)
		r1, _, _ := purego.SyscallN(gpa, uintptr(device), uintptr(unsafe.Pointer(cName)))
		if r1 != 0 {
			return r1
		}
	}
	return GetInstanceProcAddr(0, name)
}

// CallSyscall executes a function pointer via Purego SyscallN without CGO.
func CallSyscall(fn uintptr, args ...uintptr) (uintptr, uintptr, error) {
	if fn == 0 {
		return 0, 0, errors.New("vulkan function pointer is nil (procedure not loaded or supported)")
	}
	r1, r2, _ := purego.SyscallN(fn, args...)
	return r1, r2, nil
}

func getVulkanLibraryNames() []string {
	switch runtime.GOOS {
	case "windows":
		return []string{"vulkan-1.dll"}
	case "darwin":
		return []string{"libMoltenVK.dylib", "libvulkan.1.dylib", "libvulkan.dylib"}
	case "android":
		return []string{"libvulkan.so"}
	default:
		return []string{"libvulkan.so.1", "libvulkan.so"}
	}
}
`, pkgName)
}

func EmitBranchUtils(pkgName string) string {
	return fmt.Sprintf(`package %s

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
`, pkgName)
}
