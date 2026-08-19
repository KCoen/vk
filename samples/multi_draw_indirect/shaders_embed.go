package multi_draw_indirect

import _ "embed"

//go:embed shaders/vert.spv
var VertShaderSPV []byte

//go:embed shaders/frag.spv
var FragShaderSPV []byte

//go:embed shaders/cull.spv
var CullShaderSPV []byte
