Ai did most the work.

I couldn't find vulkan bindings for go that where recent and din't rely on CGO

its pretty much like the Cpp Api but it has some improvements to make it a bit nicer to work with    
* in most places pointers and counters are replaced with slices. 
* PNext is typechecked when possible, 
* SType is automatically initialized, 
* it has comments extracted from the documentation.
* out pointers are turned into return values
* there is a cool multi draw indirect GPU culling sample with imgui.

You probably want to use `vulkan` instead of `vulkanbase` or `vulkansc`, all extensions also have their own package.

To update when they release vulkan 1.5, update the Vulkan-Docs submodule and run `go run ./cmd/vk-gen` it will regenerate everything

---

## Loading Commands

Vulkan functions are dynamically loaded in stages matching the Vulkan driver and loader lifecycle:

### 1. Global Commands (`vulkan.Init()`)
Calling `vulkan.Init()` loads the platform's Vulkan dynamic library (`libvulkan.so.1`, `vulkan-1.dll`, `libvulkan.1.dylib`, etc.) and automatically performs a tailcall to `vulkan.InitCommands(0, 0)`.
- Resolves global entry points like `vulkan.CreateInstance`, `vulkan.EnumerateInstanceExtensionProperties`, and `vulkan.EnumerateInstanceLayerProperties`.
- You do **not** need to call `vulkan.InitCommands(0, 0)` manually.

```go
if err := vulkan.Init(); err != nil {
    log.Fatalf("failed to initialize Vulkan loader: %v", err)
}
```

### 2. Instance Commands (`vulkan.InitCommands(instance, 0)`)
After creating a `vulkan.Instance`, call `vulkan.InitCommands(instance, 0)`:
- Resolves instance-level commands like `vulkan.EnumeratePhysicalDevices`, `vulkan.GetPhysicalDeviceProperties`, `vulkan.CreateDevice`, etc.

```go
instance, res := vulkan.CreateInstance(&createInfo, nil)
if res != vulkan.SUCCESS {
    log.Fatalf("failed to create instance: %v", res)
}
defer vulkan.DestroyInstance(instance, nil)

vulkan.InitCommands(instance, 0)
```

### 3. Device Commands (`vulkan.InitCommands(instance, device)`)
After creating a `vulkan.Device`, call `vulkan.InitCommands(instance, device)`:
- Resolves device-level commands directly through `vkGetDeviceProcAddr` (e.g. `vulkan.AllocateMemory`, `vulkan.CmdDraw`, `vulkan.QueueSubmit`), bypassing instance-level dispatch overhead for high performance.

```go
device, res := vulkan.CreateDevice(physicalDevice, &deviceCreateInfo, nil)
if res != vulkan.SUCCESS {
    log.Fatalf("failed to create device: %v", res)
}
defer vulkan.DestroyDevice(device, nil)

vulkan.InitCommands(instance, device)
```

### 4. Multi-Device / Multi-GPU Support (`*Commands`)
`InitCommands` returns a `*vulkan.Commands` dispatch table:
- In multi-GPU / multi-device setups, store the returned `*Commands` pointer and call methods on it directly (e.g., `cmds1.CmdDraw(...)`, `cmds2.CmdDraw(...)`) to avoid mutating package-level global state across devices.

```go
cmds1 := vulkan.InitCommands(instance, device1)
cmds2 := vulkan.InitCommands(instance, device2)

cmds1.CmdDraw(cmdBuf1, 3, 1, 0, 0)
cmds2.CmdDraw(cmdBuf2, 6, 1, 0, 0)
```

### 5. Loading Extensions (`<ext>.Init`)
Extension packages provide an `Init(instance, device) *Commands` function to resolve extension-specific procedures and return an extension command table:

```go
import "github.com/KCoen/vk/extensions/khr_swapchain"

swapchainCmds := khr_swapchain.Init(instance, device)
swapchain, res := khr_swapchain.CreateSwapchainKHR(device, &swapchainInfo, nil)
```
