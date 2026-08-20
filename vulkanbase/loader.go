package vulkanbase

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
	loaderMu                 sync.Mutex
)

// Init initializes the Vulkan library by dynamically loading the Vulkan loader without CGO.
func Init() error {
	loaderMu.Lock()
	defer loaderMu.Unlock()

	if vulkanLib != 0 {
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

	cName := StringToNullTerminated("vkGetDeviceProcAddr")
	r1, _, _ := purego.SyscallN(vkGetInstanceProcAddrPtr, 0, uintptr(unsafe.Pointer(cName)))
	if r1 != 0 {
		vkGetDeviceProcAddrPtr = r1
	} else if devAddr, err := purego.Dlsym(vulkanLib, "vkGetDeviceProcAddr"); err == nil && devAddr != 0 {
		vkGetDeviceProcAddrPtr = devAddr
	}

	InitCommands(0, 0)
	return nil
}

// GetInstanceProcAddr resolves a Vulkan procedure address for an instance (instance may be 0).
func GetInstanceProcAddr(instance Instance, name string) uintptr {
	if vkGetInstanceProcAddrPtr == 0 {
		return 0
	}
	cName := StringToNullTerminated(name)
	r1, _, _ := purego.SyscallN(vkGetInstanceProcAddrPtr, uintptr(instance), uintptr(unsafe.Pointer(cName)))
	return r1
}

// GetDeviceProcAddr resolves a Vulkan procedure address for a device (device may be 0).
func GetDeviceProcAddr(device Device, name string) uintptr {
	if vkGetDeviceProcAddrPtr == 0 {
		return 0
	}
	cName := StringToNullTerminated(name)
	r1, _, _ := purego.SyscallN(vkGetDeviceProcAddrPtr, uintptr(device), uintptr(unsafe.Pointer(cName)))
	return r1
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
