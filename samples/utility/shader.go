package utility

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"unsafe"

	"go.cld.moe/vk_google/vulkan"
)

type BindingInfo struct {
	Set            uint32
	Binding        uint32
	DescriptorType vulkan.DescriptorType
	Count          uint32
	StageFlags     vulkan.ShaderStageFlags
}

type ShaderPipeline struct {
	Ctx                 *VulkanContext
	Pipeline            vulkan.Pipeline
	PipelineLayout      vulkan.PipelineLayout
	DescriptorSetLayout vulkan.DescriptorSetLayout
	DescriptorPool      vulkan.DescriptorPool
	DescriptorSet       vulkan.DescriptorSet
	BindPoint           vulkan.PipelineBindPoint
	Bindings            []BindingInfo
}

// CreateShaderModule creates a VkShaderModule from SPIR-V byte slice.
func CreateShaderModule(ctx *VulkanContext, code []byte) (vulkan.ShaderModule, error) {
	if len(code) == 0 {
		return 0, fmt.Errorf("SPIR-V code is empty")
	}

	u32Slice := make([]uint32, (len(code)+3)/4)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&u32Slice[0])), len(code)), code)

	createInfo := vulkan.NewShaderModuleCreateInfo()
	createInfo.Code = u32Slice

	module, res := vulkan.CreateShaderModule(ctx.Device, createInfo, nil)
	if res != vulkan.SUCCESS {
		return 0, fmt.Errorf("vkCreateShaderModule failed: %v", res)
	}
	return module, nil
}

// CompileGLSLToSPIRV invokes glslangValidator or glslc if available to compile GLSL source into SPIR-V bytecode.
func CompileGLSLToSPIRV(glslSource string, stage string) ([]byte, error) {
	compilerPath, err := exec.LookPath("glslangValidator")
	if err != nil {
		compilerPath, err = exec.LookPath("glslc")
		if err != nil {
			return nil, fmt.Errorf("neither glslangValidator nor glslc found on PATH: %w", err)
		}
	}

	tmpDir, err := os.MkdirTemp("", "vk_shader_*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tmpDir)

	srcFile := filepath.Join(tmpDir, "shader."+stage)
	spvFile := filepath.Join(tmpDir, "shader.spv")

	if err := os.WriteFile(srcFile, []byte(glslSource), 0644); err != nil {
		return nil, err
	}

	var cmd *exec.Cmd
	if filepath.Base(compilerPath) == "glslangValidator" {
		cmd = exec.Command(compilerPath, "-V", srcFile, "-o", spvFile)
	} else {
		cmd = exec.Command(compilerPath, "-fshader-stage="+stage, srcFile, "-o", spvFile)
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("shader compilation failed: %v, stderr: %s", err, stderr.String())
	}

	return os.ReadFile(spvFile)
}

// ReflectSPIRV inspects SPIR-V bytecode to extract descriptor bindings and types.
func ReflectSPIRV(code []byte, stage vulkan.ShaderStageFlags) ([]BindingInfo, error) {
	if len(code) < 20 {
		return nil, fmt.Errorf("invalid SPIR-V bytecode length")
	}
	words := make([]uint32, len(code)/4)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&words[0])), len(code)), code)
	if words[0] != 0x07230203 {
		return nil, fmt.Errorf("invalid SPIR-V magic 0x%x", words[0])
	}

	decorations := make(map[uint32]map[uint32]uint32)
	types := make(map[uint32]uint32)
	arrayCounts := make(map[uint32]uint32)
	constants := make(map[uint32]uint32)
	varPointers := make(map[uint32]uint32)
	varStorage := make(map[uint32]uint32)

	idx := 5
	for idx < len(words) {
		word := words[idx]
		op := word & 0xFFFF
		length := word >> 16
		if length == 0 || idx+int(length) > len(words) {
			break
		}
		inst := words[idx : idx+int(length)]

		switch op {
		case 71: // OpDecorate
			target := inst[1]
			dec := inst[2]
			val := uint32(0)
			if len(inst) > 3 {
				val = inst[3]
			}
			if decorations[target] == nil {
				decorations[target] = make(map[uint32]uint32)
			}
			decorations[target][dec] = val

		case 43: // OpConstant
			if len(inst) >= 4 {
				constants[inst[2]] = inst[3]
			}

		case 28: // OpTypeArray
			if len(inst) >= 4 {
				if lenVal, ok := constants[inst[3]]; ok {
					arrayCounts[inst[1]] = lenVal
				}
			}

		case 32: // OpTypePointer
			if len(inst) >= 4 {
				types[inst[1]] = inst[3]
			}

		case 59: // OpVariable
			if len(inst) >= 4 {
				varPointers[inst[2]] = inst[1]
				varStorage[inst[2]] = inst[3]
			}
		}
		idx += int(length)
	}

	var bindings []BindingInfo
	for varID, decs := range decorations {
		setNum, hasSet := decs[34]         // DecorationDescriptorSet
		bindingNum, hasBinding := decs[33] // DecorationBinding
		if !hasSet || !hasBinding {
			continue
		}

		storageClass := varStorage[varID]
		typeID := varPointers[varID]
		underlyingType := types[typeID]

		count := uint32(1)
		if arrLen, isArr := arrayCounts[underlyingType]; isArr {
			count = arrLen
		}

		descType := vulkan.DESCRIPTOR_TYPE_STORAGE_BUFFER
		switch storageClass {
		case 0: // UniformConstant
			descType = vulkan.DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER
		case 2: // Uniform
			if _, isBufferBlock := decorations[underlyingType][3]; isBufferBlock {
				descType = vulkan.DESCRIPTOR_TYPE_STORAGE_BUFFER
			} else {
				descType = vulkan.DESCRIPTOR_TYPE_UNIFORM_BUFFER
			}
		case 12: // StorageBuffer
			descType = vulkan.DESCRIPTOR_TYPE_STORAGE_BUFFER
		}

		bindings = append(bindings, BindingInfo{
			Set:            setNum,
			Binding:        bindingNum,
			DescriptorType: descType,
			Count:          count,
			StageFlags:     stage,
		})
	}

	sort.Slice(bindings, func(i, j int) bool {
		return bindings[i].Binding < bindings[j].Binding
	})

	return bindings, nil
}

// PrepareComputeShader reflects the SPIR-V compute bytecode and sets up DescriptorSets and Compute Pipeline.
func PrepareComputeShader(ctx *VulkanContext, computeSPV []byte) (*ShaderPipeline, error) {
	bindings, err := ReflectSPIRV(computeSPV, vulkan.SHADER_STAGE_COMPUTE_BIT)
	if err != nil {
		return nil, fmt.Errorf("reflect compute shader failed: %w", err)
	}

	sp := &ShaderPipeline{
		Ctx:       ctx,
		BindPoint: vulkan.PIPELINE_BIND_POINT_COMPUTE,
		Bindings:  bindings,
	}

	if err := sp.buildLayoutsAndPool(); err != nil {
		return nil, err
	}

	cullModule, err := CreateShaderModule(ctx, computeSPV)
	if err != nil {
		sp.Destroy()
		return nil, fmt.Errorf("create compute shader module failed: %w", err)
	}
	defer vulkan.DestroyShaderModule(ctx.Device, cullModule, nil)

	computePipeInfo := vulkan.NewComputePipelineCreateInfo()
	computePipeInfo.Stage = vulkan.PipelineShaderStageCreateInfo{
		SType:  vulkan.STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO,
		Stage:  vulkan.SHADER_STAGE_COMPUTE_BIT,
		Module: cullModule,
		Name:   "main",
	}
	computePipeInfo.Layout = sp.PipelineLayout

	cPipe, res := vulkan.CreateComputePipelines(ctx.Device, 0, []vulkan.ComputePipelineCreateInfo{*computePipeInfo}, nil)
	if res != vulkan.SUCCESS || cPipe == 0 {
		sp.Destroy()
		return nil, fmt.Errorf("vkCreateComputePipelines failed: %v", res)
	}
	sp.Pipeline = cPipe

	return sp, nil
}

// GraphicsPipelineConfig holds parameters for creating a graphics pipeline.
type GraphicsPipelineConfig struct {
	VertSPV          []byte
	FragSPV          []byte
	VertexBindings   []vulkan.VertexInputBindingDescription
	VertexAttributes []vulkan.VertexInputAttributeDescription
	Topology         vulkan.PrimitiveTopology
	ColorFormat      vulkan.Format
	DepthFormat      vulkan.Format
	Width            uint32
	Height           uint32
}

// PrepareGraphicsShader reflects vertex & fragment shaders and sets up DescriptorSets and Graphics Pipeline.
func PrepareGraphicsShader(ctx *VulkanContext, cfg GraphicsPipelineConfig) (*ShaderPipeline, error) {
	vertBindings, err := ReflectSPIRV(cfg.VertSPV, vulkan.SHADER_STAGE_VERTEX_BIT)
	if err != nil {
		return nil, fmt.Errorf("reflect vert shader failed: %w", err)
	}
	fragBindings, err := ReflectSPIRV(cfg.FragSPV, vulkan.SHADER_STAGE_FRAGMENT_BIT)
	if err != nil {
		return nil, fmt.Errorf("reflect frag shader failed: %w", err)
	}

	mergedMap := make(map[uint32]BindingInfo)
	for _, b := range vertBindings {
		mergedMap[b.Binding] = b
	}
	for _, b := range fragBindings {
		if existing, ok := mergedMap[b.Binding]; ok {
			existing.StageFlags |= b.StageFlags
			mergedMap[b.Binding] = existing
		} else {
			mergedMap[b.Binding] = b
		}
	}

	mergedBindings := make([]BindingInfo, 0, len(mergedMap))
	for _, b := range mergedMap {
		mergedBindings = append(mergedBindings, b)
	}
	sort.Slice(mergedBindings, func(i, j int) bool {
		return mergedBindings[i].Binding < mergedBindings[j].Binding
	})

	sp := &ShaderPipeline{
		Ctx:       ctx,
		BindPoint: vulkan.PIPELINE_BIND_POINT_GRAPHICS,
		Bindings:  mergedBindings,
	}

	if err := sp.buildLayoutsAndPool(); err != nil {
		return nil, err
	}

	vertModule, err := CreateShaderModule(ctx, cfg.VertSPV)
	if err != nil {
		sp.Destroy()
		return nil, fmt.Errorf("create vert shader module failed: %w", err)
	}
	defer vulkan.DestroyShaderModule(ctx.Device, vertModule, nil)

	fragModule, err := CreateShaderModule(ctx, cfg.FragSPV)
	if err != nil {
		sp.Destroy()
		return nil, fmt.Errorf("create frag shader module failed: %w", err)
	}
	defer vulkan.DestroyShaderModule(ctx.Device, fragModule, nil)

	stages := []vulkan.PipelineShaderStageCreateInfo{
		{SType: vulkan.STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO, Stage: vulkan.SHADER_STAGE_VERTEX_BIT, Module: vertModule, Name: "main"},
		{SType: vulkan.STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO, Stage: vulkan.SHADER_STAGE_FRAGMENT_BIT, Module: fragModule, Name: "main"},
	}

	vertexInputInfo := vulkan.NewPipelineVertexInputStateCreateInfo()
	vertexInputInfo.VertexBindingDescriptions = cfg.VertexBindings
	vertexInputInfo.VertexAttributeDescriptions = cfg.VertexAttributes

	inputAssembly := vulkan.NewPipelineInputAssemblyStateCreateInfo()
	inputAssembly.Topology = cfg.Topology
	if inputAssembly.Topology == 0 {
		inputAssembly.Topology = vulkan.PRIMITIVE_TOPOLOGY_TRIANGLE_LIST
	}

	viewportState := vulkan.NewPipelineViewportStateCreateInfo()
	viewportState.Viewports = []vulkan.Viewport{{X: 0, Y: 0, Width: float32(cfg.Width), Height: float32(cfg.Height), MinDepth: 0.0, MaxDepth: 1.0}}
	viewportState.Scissors = []vulkan.Rect2D{{Offset: vulkan.Offset2D{X: 0, Y: 0}, Extent: vulkan.Extent2D{Width: cfg.Width, Height: cfg.Height}}}

	rasterizer := vulkan.NewPipelineRasterizationStateCreateInfo()
	rasterizer.PolygonMode = vulkan.POLYGON_MODE_FILL
	rasterizer.CullMode = vulkan.CULL_MODE_BACK_BIT
	rasterizer.FrontFace = vulkan.FRONT_FACE_COUNTER_CLOCKWISE
	rasterizer.LineWidth = 1.0

	multisample := vulkan.NewPipelineMultisampleStateCreateInfo()
	multisample.RasterizationSamples = vulkan.SAMPLE_COUNT_1_BIT

	depthStencil := vulkan.NewPipelineDepthStencilStateCreateInfo()
	depthStencil.DepthTestEnable = vulkan.True
	depthStencil.DepthWriteEnable = vulkan.True
	depthStencil.DepthCompareOp = vulkan.COMPARE_OP_LESS

	colorBlendAttachment := vulkan.PipelineColorBlendAttachmentState{
		ColorWriteMask: vulkan.COLOR_COMPONENT_R_BIT | vulkan.COLOR_COMPONENT_G_BIT | vulkan.COLOR_COMPONENT_B_BIT | vulkan.COLOR_COMPONENT_A_BIT,
		BlendEnable:    vulkan.False,
	}
	colorBlending := vulkan.NewPipelineColorBlendStateCreateInfo()
	colorBlending.Attachments = []vulkan.PipelineColorBlendAttachmentState{colorBlendAttachment}

	colorFmt := cfg.ColorFormat
	if colorFmt == 0 {
		colorFmt = vulkan.FORMAT_B8G8R8A8_UNORM
	}
	depthFmt := cfg.DepthFormat
	if depthFmt == 0 {
		depthFmt = vulkan.FORMAT_D32_SFLOAT
	}

	dynRendInfo := vulkan.NewPipelineRenderingCreateInfo()
	dynRendInfo.ColorAttachmentFormats = []vulkan.Format{colorFmt}
	dynRendInfo.DepthAttachmentFormat = depthFmt

	pipelineInfo := vulkan.NewGraphicsPipelineCreateInfo()
	pipelineInfo.Stages = stages
	pipelineInfo.VertexInputState = vertexInputInfo
	pipelineInfo.InputAssemblyState = inputAssembly
	pipelineInfo.ViewportState = viewportState
	pipelineInfo.RasterizationState = rasterizer
	pipelineInfo.MultisampleState = multisample
	pipelineInfo.DepthStencilState = depthStencil
	pipelineInfo.ColorBlendState = colorBlending
	pipelineInfo.Layout = sp.PipelineLayout
	pipelineInfo.SetNext(dynRendInfo)

	pipe, res := vulkan.CreateGraphicsPipelines(ctx.Device, 0, []vulkan.GraphicsPipelineCreateInfo{*pipelineInfo}, nil)
	if res != vulkan.SUCCESS || pipe == 0 {
		sp.Destroy()
		return nil, fmt.Errorf("vkCreateGraphicsPipelines failed: %v", res)
	}
	sp.Pipeline = pipe

	return sp, nil
}

func (sp *ShaderPipeline) buildLayoutsAndPool() error {
	layoutBindings := make([]vulkan.DescriptorSetLayoutBinding, len(sp.Bindings))
	poolSizesMap := make(map[vulkan.DescriptorType]uint32)

	for i, b := range sp.Bindings {
		layoutBindings[i] = vulkan.DescriptorSetLayoutBinding{
			Binding:         b.Binding,
			DescriptorType:  b.DescriptorType,
			DescriptorCount: b.Count,
			StageFlags:      b.StageFlags,
		}
		poolSizesMap[b.DescriptorType] += b.Count
	}

	layoutInfo := vulkan.NewDescriptorSetLayoutCreateInfo()
	layoutInfo.Bindings = layoutBindings
	dsl, res := vulkan.CreateDescriptorSetLayout(sp.Ctx.Device, layoutInfo, nil)
	if res != vulkan.SUCCESS {
		return fmt.Errorf("vkCreateDescriptorSetLayout failed: %v", res)
	}
	sp.DescriptorSetLayout = dsl

	pipeLayoutInfo := vulkan.NewPipelineLayoutCreateInfo()
	pipeLayoutInfo.SetLayouts = []vulkan.DescriptorSetLayout{dsl}
	pl, res := vulkan.CreatePipelineLayout(sp.Ctx.Device, pipeLayoutInfo, nil)
	if res != vulkan.SUCCESS {
		return fmt.Errorf("vkCreatePipelineLayout failed: %v", res)
	}
	sp.PipelineLayout = pl

	var poolSizes []vulkan.DescriptorPoolSize
	for dt, count := range poolSizesMap {
		poolSizes = append(poolSizes, vulkan.DescriptorPoolSize{Type: dt, DescriptorCount: count * 2})
	}

	poolInfo := vulkan.NewDescriptorPoolCreateInfo()
	poolInfo.PoolSizes = poolSizes
	poolInfo.MaxSets = 2

	pool, res := vulkan.CreateDescriptorPool(sp.Ctx.Device, poolInfo, nil)
	if res != vulkan.SUCCESS {
		return fmt.Errorf("vkCreateDescriptorPool failed: %v", res)
	}
	sp.DescriptorPool = pool

	allocInfo := vulkan.NewDescriptorSetAllocateInfo()
	allocInfo.DescriptorPool = pool
	allocInfo.SetLayouts = []vulkan.DescriptorSetLayout{dsl}

	set, res := vulkan.AllocateDescriptorSets(sp.Ctx.Device, allocInfo)
	if res != vulkan.SUCCESS || set == 0 {
		return fmt.Errorf("allocate descriptor set failed: %v", res)
	}
	sp.DescriptorSet = set

	return nil
}

// BindResources connects buffers and textures to descriptor set bindings automatically.
func (sp *ShaderPipeline) BindResources(resources ...any) {
	var writes []vulkan.WriteDescriptorSet
	usedBindings := make(map[uint32]bool)

	for _, res := range resources {
		switch v := res.(type) {
		case *Buffer:
			if v == nil {
				continue
			}
			var matched *BindingInfo
			for i := range sp.Bindings {
				b := &sp.Bindings[i]
				if !usedBindings[b.Binding] && (b.DescriptorType == vulkan.DESCRIPTOR_TYPE_UNIFORM_BUFFER || b.DescriptorType == vulkan.DESCRIPTOR_TYPE_STORAGE_BUFFER) {
					matched = b
					usedBindings[b.Binding] = true
					break
				}
			}
			if matched != nil {
				bInfo := vulkan.DescriptorBufferInfo{Buffer: v.Handle, Offset: 0, Range: vulkan.DeviceSize(v.Size)}
				writes = append(writes, vulkan.WriteDescriptorSet{
					SType:          vulkan.STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET,
					DstSet:         sp.DescriptorSet,
					DstBinding:     matched.Binding,
					DescriptorType: matched.DescriptorType,
					BufferInfo:     []vulkan.DescriptorBufferInfo{bInfo},
				})
			}

		case []*Texture:
			if len(v) == 0 {
				continue
			}
			var matched *BindingInfo
			for i := range sp.Bindings {
				b := &sp.Bindings[i]
				if !usedBindings[b.Binding] && b.DescriptorType == vulkan.DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER {
					matched = b
					usedBindings[b.Binding] = true
					break
				}
			}
			if matched != nil {
				imageInfos := make([]vulkan.DescriptorImageInfo, len(v))
				for idx, tex := range v {
					imageInfos[idx] = vulkan.DescriptorImageInfo{
						Sampler:     tex.Sampler,
						ImageView:   tex.View,
						ImageLayout: vulkan.IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL,
					}
				}
				writes = append(writes, vulkan.WriteDescriptorSet{
					SType:          vulkan.STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET,
					DstSet:         sp.DescriptorSet,
					DstBinding:     matched.Binding,
					DescriptorType: matched.DescriptorType,
					ImageInfo:      imageInfos,
				})
			}
		}
	}

	if len(writes) > 0 {
		vulkan.UpdateDescriptorSets(sp.Ctx.Device, writes, nil)
	}
}

// Bind binds the pipeline and descriptor set to the command buffer.
func (sp *ShaderPipeline) Bind(cmd vulkan.CommandBuffer) {
	vulkan.CmdBindPipeline(cmd, sp.BindPoint, sp.Pipeline)
	vulkan.CmdBindDescriptorSets(cmd, sp.BindPoint, sp.PipelineLayout, 0, []vulkan.DescriptorSet{sp.DescriptorSet}, nil)
}

// Dispatch binds the compute pipeline and dispatches compute work.
func (sp *ShaderPipeline) Dispatch(cmd vulkan.CommandBuffer, groupCountX, groupCountY, groupCountZ uint32) {
	sp.Bind(cmd)
	vulkan.CmdDispatch(cmd, groupCountX, groupCountY, groupCountZ)
}

// Destroy cleans up allocated Vulkan pipeline and descriptor resources.
func (sp *ShaderPipeline) Destroy() {
	if sp.Ctx == nil || sp.Ctx.Device == 0 {
		return
	}
	if sp.Pipeline != 0 {
		vulkan.DestroyPipeline(sp.Ctx.Device, sp.Pipeline, nil)
		sp.Pipeline = 0
	}
	if sp.PipelineLayout != 0 {
		vulkan.DestroyPipelineLayout(sp.Ctx.Device, sp.PipelineLayout, nil)
		sp.PipelineLayout = 0
	}
	if sp.DescriptorPool != 0 {
		vulkan.DestroyDescriptorPool(sp.Ctx.Device, sp.DescriptorPool, nil)
		sp.DescriptorPool = 0
	}
	if sp.DescriptorSetLayout != 0 {
		vulkan.DestroyDescriptorSetLayout(sp.Ctx.Device, sp.DescriptorSetLayout, nil)
		sp.DescriptorSetLayout = 0
	}
}
