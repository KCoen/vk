package cimgui_vulkan

import (
	_ "embed"
)

//go:embed shaders/cimgui_vert.spv
var VertShaderSPIRV []byte

//go:embed shaders/cimgui_frag.spv
var FragShaderSPIRV []byte
