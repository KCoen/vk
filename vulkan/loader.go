package vulkan

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
		return fmt.Errorf("failed to load Vulkan library (%v): %w", libNames, lastErr)
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
