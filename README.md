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

You need to load the functions in stages:
* Call `vulkan.Init()` first (it loads the library and global stuff like `CreateInstance`).
* Once you have an instance, call `vulkan.InitCommands(instance, 0)`.
* Once you have a device, call `vulkan.InitCommands(instance, device)` so device calls skip the loader overhead.
* If you use multiple devices, `InitCommands` returns a `*Commands` struct with methods so you don't clobber globals.
* Extensions work the exact same way with `<ext_pkg>.Init(instance, device)`.

To update when they release vulkan 1.5, update the Vulkan-Docs submodule and run `go run ./cmd/vk-gen` it will regenerate everything
