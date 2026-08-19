# Vk - Vulkan Bindings

Windows Golang bindings for Vulkan, that don't use cgo

# Linux
Because library mapping in linux is primarily a usermode venture, and golang (without CGO) can only call kernel functions, impelmenting this requires you to implement the runtime linker in golang, in theory this can be done, but its not done here.
