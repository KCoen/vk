package vk

//----------------------------------------
//--Build-in
//----------------------------------------
type unk = int
type external_type = unk
type funcptr = uintptr
type handle = uintptr

//----------------------------------------
//--Macros
//----------------------------------------
func VK_MAKE_VERSION(major, minor, patch int) uint32 {
	return ((((uint32)(major)) << 22) | (((uint32)(minor)) << 12) | ((uint32)(patch)))
}
