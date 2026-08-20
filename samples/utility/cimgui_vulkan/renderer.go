package cimgui_vulkan

import (
	"fmt"
	"math"
	"unsafe"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/KCoen/vk/samples/utility"
	"github.com/KCoen/vk/vulkan"
)

// Renderer manages Vulkan rendering resources for cimgui-go draw data.
type Renderer struct {
	Ctx                 *utility.VulkanContext
	DescriptorPool      vulkan.DescriptorPool
	DescriptorSetLayout vulkan.DescriptorSetLayout
	DescriptorSet       vulkan.DescriptorSet
	PipelineLayout      vulkan.PipelineLayout
	Pipeline            vulkan.Pipeline
	FontSampler         vulkan.Sampler
	FontImage           vulkan.Image
	FontMemory          vulkan.DeviceMemory
	FontView            vulkan.ImageView

	VertexBuffer     *utility.Buffer
	IndexBuffer      *utility.Buffer
	VertexBufferSize int
	IndexBufferSize  int
}

// NewRenderer creates and initializes a Vulkan renderer for cimgui-go.
// colorFormat must match the swapchain format used during rendering.
func NewRenderer(ctx *utility.VulkanContext, colorFormat vulkan.Format) (*Renderer, error) {
	r := &Renderer{
		Ctx: ctx,
	}

	if err := r.initPipeline(colorFormat); err != nil {
		r.Destroy()
		return nil, fmt.Errorf("failed to init pipeline: %w", err)
	}

	return r, nil
}

func (r *Renderer) updateTexture(tex *imgui.TextureData) error {
	width := uint32(tex.Width())
	height := uint32(tex.Height())
	if width == 0 || height == 0 || tex.Pixels() == 0 {
		return nil
	}

	pixels := unsafe.Slice((*byte)(unsafe.Pointer(tex.Pixels())), tex.SizeInBytes())

	if r.FontImage != 0 {
		vulkan.DestroyImageView(r.Ctx.Device, r.FontView, nil)
		vulkan.DestroyImage(r.Ctx.Device, r.FontImage, nil)
		vulkan.FreeMemory(r.Ctx.Device, r.FontMemory, nil)
		r.FontView = 0
		r.FontImage = 0
		r.FontMemory = 0
	}

	// Create staging buffer
	stagingBuffer, err := utility.CreateBuffer(
		r.Ctx,
		uint64(len(pixels)),
		vulkan.BUFFER_USAGE_TRANSFER_SRC_BIT,
		vulkan.MEMORY_PROPERTY_HOST_VISIBLE_BIT|vulkan.MEMORY_PROPERTY_HOST_COHERENT_BIT,
	)
	if err != nil {
		return fmt.Errorf("failed to create staging buffer: %w", err)
	}
	defer stagingBuffer.Destroy(r.Ctx)

	copy(unsafe.Slice((*byte)(stagingBuffer.Mapped), len(pixels)), pixels)

	// Create Vulkan Image
	imageInfo := vulkan.NewImageCreateInfo()
	imageInfo.ImageType = vulkan.IMAGE_TYPE_2D
	imageInfo.Format = vulkan.FORMAT_R8G8B8A8_UNORM
	imageInfo.Extent = vulkan.Extent3D{Width: width, Height: height, Depth: 1}
	imageInfo.MipLevels = 1
	imageInfo.ArrayLayers = 1
	imageInfo.Samples = vulkan.SAMPLE_COUNT_1_BIT
	imageInfo.Tiling = vulkan.IMAGE_TILING_OPTIMAL
	imageInfo.Usage = vulkan.IMAGE_USAGE_SAMPLED_BIT | vulkan.IMAGE_USAGE_TRANSFER_DST_BIT
	imageInfo.SharingMode = vulkan.SHARING_MODE_EXCLUSIVE
	imageInfo.InitialLayout = vulkan.IMAGE_LAYOUT_UNDEFINED

	var res vulkan.Result
	r.FontImage, res = vulkan.CreateImage(r.Ctx.Device, imageInfo, nil)
	if res != vulkan.SUCCESS {
		return fmt.Errorf("failed to create font image: %v", res)
	}

	memReqs := vulkan.GetImageMemoryRequirements(r.Ctx.Device, r.FontImage)
	memType, err := r.Ctx.FindMemoryType(memReqs.MemoryTypeBits, vulkan.MEMORY_PROPERTY_DEVICE_LOCAL_BIT)
	if err != nil {
		return fmt.Errorf("failed to find memory type for font image: %w", err)
	}

	allocInfo := vulkan.NewMemoryAllocateInfo()
	allocInfo.AllocationSize = memReqs.Size
	allocInfo.MemoryTypeIndex = memType

	r.FontMemory, res = vulkan.AllocateMemory(r.Ctx.Device, allocInfo, nil)
	if res != vulkan.SUCCESS {
		return fmt.Errorf("failed to allocate font memory: %v", res)
	}

	vulkan.BindImageMemory(r.Ctx.Device, r.FontImage, r.FontMemory, 0)

	// Transition image and copy buffer
	err = r.Ctx.ExecuteOneTimeCommands(func(cmd vulkan.CommandBuffer) {
		barrier := vulkan.NewImageMemoryBarrier()
		barrier.OldLayout = vulkan.IMAGE_LAYOUT_UNDEFINED
		barrier.NewLayout = vulkan.IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL
		barrier.SrcQueueFamilyIndex = vulkan.QUEUE_FAMILY_IGNORED
		barrier.DstQueueFamilyIndex = vulkan.QUEUE_FAMILY_IGNORED
		barrier.Image = r.FontImage
		barrier.SubresourceRange = vulkan.ImageSubresourceRange{
			AspectMask: vulkan.IMAGE_ASPECT_COLOR_BIT,
			LevelCount: 1,
			LayerCount: 1,
		}
		barrier.SrcAccessMask = 0
		barrier.DstAccessMask = vulkan.ACCESS_TRANSFER_WRITE_BIT

		vulkan.CmdPipelineBarrier(
			cmd,
			vulkan.PIPELINE_STAGE_TOP_OF_PIPE_BIT,
			vulkan.PIPELINE_STAGE_TRANSFER_BIT,
			0,
			nil, nil,
			[]vulkan.ImageMemoryBarrier{*barrier},
		)

		region := vulkan.BufferImageCopy{
			ImageSubresource: vulkan.ImageSubresourceLayers{
				AspectMask: vulkan.IMAGE_ASPECT_COLOR_BIT,
				LayerCount: 1,
			},
			ImageExtent: vulkan.Extent3D{Width: width, Height: height, Depth: 1},
		}

		vulkan.CmdCopyBufferToImage(
			cmd,
			stagingBuffer.Handle,
			r.FontImage,
			vulkan.IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL,
			[]vulkan.BufferImageCopy{region},
		)

		barrier.OldLayout = vulkan.IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL
		barrier.NewLayout = vulkan.IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL
		barrier.SrcAccessMask = vulkan.ACCESS_TRANSFER_WRITE_BIT
		barrier.DstAccessMask = vulkan.ACCESS_SHADER_READ_BIT

		vulkan.CmdPipelineBarrier(
			cmd,
			vulkan.PIPELINE_STAGE_TRANSFER_BIT,
			vulkan.PIPELINE_STAGE_FRAGMENT_SHADER_BIT,
			0,
			nil, nil,
			[]vulkan.ImageMemoryBarrier{*barrier},
		)
	})
	if err != nil {
		return fmt.Errorf("failed to copy font texture: %w", err)
	}

	// Create Image View
	viewInfo := vulkan.NewImageViewCreateInfo()
	viewInfo.Image = r.FontImage
	viewInfo.ViewType = vulkan.IMAGE_VIEW_TYPE_2D
	viewInfo.Format = vulkan.FORMAT_R8G8B8A8_UNORM
	viewInfo.SubresourceRange = vulkan.ImageSubresourceRange{
		AspectMask: vulkan.IMAGE_ASPECT_COLOR_BIT,
		LevelCount: 1,
		LayerCount: 1,
	}

	r.FontView, res = vulkan.CreateImageView(r.Ctx.Device, viewInfo, nil)
	if res != vulkan.SUCCESS {
		return fmt.Errorf("failed to create font image view: %v", res)
	}

	// Update Descriptor Set
	imageDescInfo := vulkan.DescriptorImageInfo{
		Sampler:     r.FontSampler,
		ImageView:   r.FontView,
		ImageLayout: vulkan.IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL,
	}
	write := vulkan.NewWriteDescriptorSet()
	write.DstSet = r.DescriptorSet
	write.DstBinding = 0
	write.DescriptorType = vulkan.DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER
	write.ImageInfo = []vulkan.DescriptorImageInfo{imageDescInfo}

	vulkan.UpdateDescriptorSets(r.Ctx.Device, []vulkan.WriteDescriptorSet{*write}, nil)
	tex.SetStatus(imgui.TextureStatusOK)

	return nil
}

func (r *Renderer) initPipeline(colorFormat vulkan.Format) error {
	var res vulkan.Result

	// 1. Create Sampler
	samplerInfo := vulkan.NewSamplerCreateInfo()
	samplerInfo.MagFilter = vulkan.FILTER_LINEAR
	samplerInfo.MinFilter = vulkan.FILTER_LINEAR
	samplerInfo.MipmapMode = vulkan.SAMPLER_MIPMAP_MODE_LINEAR
	samplerInfo.AddressModeU = vulkan.SAMPLER_ADDRESS_MODE_REPEAT
	samplerInfo.AddressModeV = vulkan.SAMPLER_ADDRESS_MODE_REPEAT
	samplerInfo.AddressModeW = vulkan.SAMPLER_ADDRESS_MODE_REPEAT
	samplerInfo.MaxAnisotropy = 1.0

	r.FontSampler, res = vulkan.CreateSampler(r.Ctx.Device, samplerInfo, nil)
	if res != vulkan.SUCCESS {
		return fmt.Errorf("failed to create font sampler: %v", res)
	}

	// 2. Descriptor Pool
	poolSize := vulkan.DescriptorPoolSize{
		Type:            vulkan.DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER,
		DescriptorCount: 1,
	}
	poolInfo := vulkan.NewDescriptorPoolCreateInfo()
	poolInfo.MaxSets = 1
	poolInfo.PoolSizes = []vulkan.DescriptorPoolSize{poolSize}

	r.DescriptorPool, res = vulkan.CreateDescriptorPool(r.Ctx.Device, poolInfo, nil)
	if res != vulkan.SUCCESS {
		return fmt.Errorf("failed to create descriptor pool: %v", res)
	}

	// 3. Descriptor Set Layout
	binding := vulkan.DescriptorSetLayoutBinding{
		Binding:         0,
		DescriptorType:  vulkan.DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER,
		DescriptorCount: 1,
		StageFlags:      vulkan.SHADER_STAGE_FRAGMENT_BIT,
	}
	layoutInfo := vulkan.NewDescriptorSetLayoutCreateInfo()
	layoutInfo.Bindings = []vulkan.DescriptorSetLayoutBinding{binding}

	r.DescriptorSetLayout, res = vulkan.CreateDescriptorSetLayout(r.Ctx.Device, layoutInfo, nil)
	if res != vulkan.SUCCESS {
		return fmt.Errorf("failed to create descriptor set layout: %v", res)
	}

	// 4. Allocate Descriptor Set
	allocInfo := vulkan.NewDescriptorSetAllocateInfo()
	allocInfo.DescriptorPool = r.DescriptorPool
	allocInfo.SetLayouts = []vulkan.DescriptorSetLayout{r.DescriptorSetLayout}

	r.DescriptorSet, res = vulkan.AllocateDescriptorSets(r.Ctx.Device, allocInfo)
	if res != vulkan.SUCCESS {
		return fmt.Errorf("failed to allocate descriptor set: %v", res)
	}

	// 5. Pipeline Layout (Push Constant 16 bytes: vec2 scale, vec2 translate)
	pushConst := vulkan.PushConstantRange{
		StageFlags: vulkan.SHADER_STAGE_VERTEX_BIT,
		Offset:     0,
		Size:       16,
	}
	pipelineLayoutInfo := vulkan.NewPipelineLayoutCreateInfo()
	pipelineLayoutInfo.SetLayouts = []vulkan.DescriptorSetLayout{r.DescriptorSetLayout}
	pipelineLayoutInfo.PushConstantRanges = []vulkan.PushConstantRange{pushConst}

	r.PipelineLayout, res = vulkan.CreatePipelineLayout(r.Ctx.Device, pipelineLayoutInfo, nil)
	if res != vulkan.SUCCESS {
		return fmt.Errorf("failed to create pipeline layout: %v", res)
	}

	// 6. Shader Modules
	vertMod, err := utility.CreateShaderModule(r.Ctx, VertShaderSPIRV)
	if err != nil {
		return fmt.Errorf("failed to create vert shader: %w", err)
	}
	defer vulkan.DestroyShaderModule(r.Ctx.Device, vertMod, nil)

	fragMod, err := utility.CreateShaderModule(r.Ctx, FragShaderSPIRV)
	if err != nil {
		return fmt.Errorf("failed to create frag shader: %w", err)
	}
	defer vulkan.DestroyShaderModule(r.Ctx.Device, fragMod, nil)

	vertStage := vulkan.NewPipelineShaderStageCreateInfo()
	vertStage.Stage = vulkan.SHADER_STAGE_VERTEX_BIT
	vertStage.Module = vertMod
	vertStage.Name = "main"

	fragStage := vulkan.NewPipelineShaderStageCreateInfo()
	fragStage.Stage = vulkan.SHADER_STAGE_FRAGMENT_BIT
	fragStage.Module = fragMod
	fragStage.Name = "main"

	shaderStages := []vulkan.PipelineShaderStageCreateInfo{*vertStage, *fragStage}

	// 7. Vertex Input State
	// ImDrawVert: Pos [2]float32 (0..8), UV [2]float32 (8..16), Col uint32 (16..20) = stride 20
	bindingDesc := vulkan.VertexInputBindingDescription{
		Binding:   0,
		Stride:    20,
		InputRate: vulkan.VERTEX_INPUT_RATE_VERTEX,
	}
	attribDescs := []vulkan.VertexInputAttributeDescription{
		{Location: 0, Binding: 0, Format: vulkan.FORMAT_R32G32_SFLOAT, Offset: 0},
		{Location: 1, Binding: 0, Format: vulkan.FORMAT_R32G32_SFLOAT, Offset: 8},
		{Location: 2, Binding: 0, Format: vulkan.FORMAT_R8G8B8A8_UNORM, Offset: 16},
	}

	vertexInput := vulkan.NewPipelineVertexInputStateCreateInfo()
	vertexInput.VertexBindingDescriptions = []vulkan.VertexInputBindingDescription{bindingDesc}
	vertexInput.VertexAttributeDescriptions = attribDescs

	inputAssembly := vulkan.NewPipelineInputAssemblyStateCreateInfo()
	inputAssembly.Topology = vulkan.PRIMITIVE_TOPOLOGY_TRIANGLE_LIST

	viewportState := vulkan.NewPipelineViewportStateCreateInfo()
	viewportState.Viewports = make([]vulkan.Viewport, 1)
	viewportState.Scissors = make([]vulkan.Rect2D, 1)

	rasterizer := vulkan.NewPipelineRasterizationStateCreateInfo()
	rasterizer.PolygonMode = vulkan.POLYGON_MODE_FILL
	rasterizer.CullMode = vulkan.CULL_MODE_NONE
	rasterizer.FrontFace = vulkan.FRONT_FACE_COUNTER_CLOCKWISE
	rasterizer.LineWidth = 1.0

	multisampling := vulkan.NewPipelineMultisampleStateCreateInfo()
	multisampling.RasterizationSamples = vulkan.SAMPLE_COUNT_1_BIT

	depthStencil := vulkan.NewPipelineDepthStencilStateCreateInfo()
	depthStencil.DepthTestEnable = vulkan.False
	depthStencil.DepthWriteEnable = vulkan.False
	depthStencil.DepthCompareOp = vulkan.COMPARE_OP_ALWAYS

	colorBlendAttachment := vulkan.PipelineColorBlendAttachmentState{
		BlendEnable:         vulkan.True,
		SrcColorBlendFactor: vulkan.BLEND_FACTOR_SRC_ALPHA,
		DstColorBlendFactor: vulkan.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA,
		ColorBlendOp:        vulkan.BLEND_OP_ADD,
		SrcAlphaBlendFactor: vulkan.BLEND_FACTOR_ONE,
		DstAlphaBlendFactor: vulkan.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA,
		AlphaBlendOp:        vulkan.BLEND_OP_ADD,
		ColorWriteMask:      vulkan.COLOR_COMPONENT_R_BIT | vulkan.COLOR_COMPONENT_G_BIT | vulkan.COLOR_COMPONENT_B_BIT | vulkan.COLOR_COMPONENT_A_BIT,
	}

	colorBlending := vulkan.NewPipelineColorBlendStateCreateInfo()
	colorBlending.Attachments = []vulkan.PipelineColorBlendAttachmentState{colorBlendAttachment}

	dynamicStates := []vulkan.DynamicState{
		vulkan.DYNAMIC_STATE_VIEWPORT,
		vulkan.DYNAMIC_STATE_SCISSOR,
	}
	dynamicState := vulkan.NewPipelineDynamicStateCreateInfo()
	dynamicState.DynamicStates = dynamicStates

	// Use dynamic rendering — chain PipelineRenderingCreateInfo instead of render pass handle.
	dynRendInfo := vulkan.NewPipelineRenderingCreateInfo()
	dynRendInfo.ColorAttachmentFormats = []vulkan.Format{colorFormat}
	dynRendInfo.DepthAttachmentFormat = vulkan.FORMAT_D32_SFLOAT

	pipelineInfo := vulkan.NewGraphicsPipelineCreateInfo()
	pipelineInfo.Stages = shaderStages
	pipelineInfo.VertexInputState = vertexInput
	pipelineInfo.InputAssemblyState = inputAssembly
	pipelineInfo.ViewportState = viewportState
	pipelineInfo.RasterizationState = rasterizer
	pipelineInfo.MultisampleState = multisampling
	pipelineInfo.DepthStencilState = depthStencil
	pipelineInfo.ColorBlendState = colorBlending
	pipelineInfo.DynamicState = dynamicState
	pipelineInfo.Layout = r.PipelineLayout
	pipelineInfo.SetNext(dynRendInfo)

	r.Pipeline, res = vulkan.CreateGraphicsPipelines(r.Ctx.Device, 0, []vulkan.GraphicsPipelineCreateInfo{*pipelineInfo}, nil)
	if res != vulkan.SUCCESS {
		return fmt.Errorf("failed to create graphics pipeline: %v", res)
	}

	return nil
}

// RenderDrawData renders cimgui-go draw data into the provided command buffer.
func (r *Renderer) RenderDrawData(drawData *imgui.DrawData, cmd vulkan.CommandBuffer) {
	if drawData == nil || !drawData.Valid() || drawData.TotalVtxCount() == 0 {
		return
	}

	// Update textures if needed
	texSlice := drawData.Textures().Slice()
	for i := range texSlice {
		tex := &texSlice[i]
		if r.FontImage == 0 || tex.Status() == imgui.TextureStatusWantCreate || tex.Pixels() != 0 {
			if err := r.updateTexture(tex); err != nil {
				fmt.Printf("cimgui_vulkan: failed to update texture: %v\n", err)
			}
		}
	}

	dispSize := drawData.DisplaySize()
	dispPos := drawData.DisplayPos()
	if dispSize.X <= 0 || dispSize.Y <= 0 {
		return
	}

	totalVtx := int(drawData.TotalVtxCount())
	totalIdx := int(drawData.TotalIdxCount())
	vtxBytes := totalVtx * 20
	idxBytes := totalIdx * 2

	// Grow vertex buffer if needed
	if r.VertexBuffer == nil || r.VertexBufferSize < vtxBytes {
		if r.VertexBuffer != nil {
			r.VertexBuffer.Destroy(r.Ctx)
		}
		newSize := vtxBytes + 5000*20
		buf, err := utility.CreateBuffer(
			r.Ctx,
			uint64(newSize),
			vulkan.BUFFER_USAGE_VERTEX_BUFFER_BIT,
			vulkan.MEMORY_PROPERTY_HOST_VISIBLE_BIT|vulkan.MEMORY_PROPERTY_HOST_COHERENT_BIT,
		)
		if err != nil {
			return
		}
		r.VertexBuffer = buf
		r.VertexBufferSize = newSize
	}

	// Grow index buffer if needed
	if r.IndexBuffer == nil || r.IndexBufferSize < idxBytes {
		if r.IndexBuffer != nil {
			r.IndexBuffer.Destroy(r.Ctx)
		}
		newSize := idxBytes + 10000*2
		buf, err := utility.CreateBuffer(
			r.Ctx,
			uint64(newSize),
			vulkan.BUFFER_USAGE_INDEX_BUFFER_BIT,
			vulkan.MEMORY_PROPERTY_HOST_VISIBLE_BIT|vulkan.MEMORY_PROPERTY_HOST_COHERENT_BIT,
		)
		if err != nil {
			return
		}
		r.IndexBuffer = buf
		r.IndexBufferSize = newSize
	}

	vtxDst := unsafe.Slice((*byte)(r.VertexBuffer.Mapped), r.VertexBufferSize)
	idxDst := unsafe.Slice((*byte)(r.IndexBuffer.Mapped), r.IndexBufferSize)

	vtxOffsetBytes := 0
	idxOffsetBytes := 0

	cmdLists := drawData.CommandLists()
	for _, cl := range cmdLists {
		vPtr, vCount := cl.GetVertexBuffer()
		iPtr, iCount := cl.GetIndexBuffer()

		if vCount > 0 && vPtr != nil {
			copy(vtxDst[vtxOffsetBytes:vtxOffsetBytes+vCount], unsafe.Slice((*byte)(vPtr), vCount))
			vtxOffsetBytes += vCount
		}
		if iCount > 0 && iPtr != nil {
			copy(idxDst[idxOffsetBytes:idxOffsetBytes+iCount], unsafe.Slice((*byte)(iPtr), iCount))
			idxOffsetBytes += iCount
		}
	}

	// Bind Pipeline & Descriptors
	vulkan.CmdBindPipeline(cmd, vulkan.PIPELINE_BIND_POINT_GRAPHICS, r.Pipeline)
	vulkan.CmdBindDescriptorSets(cmd, vulkan.PIPELINE_BIND_POINT_GRAPHICS, r.PipelineLayout, 0, []vulkan.DescriptorSet{r.DescriptorSet}, nil)
	offset := vulkan.DeviceSize(0)
	vulkan.CmdBindVertexBuffers(cmd, 0, []vulkan.Buffer{r.VertexBuffer.Handle}, &offset)
	vulkan.CmdBindIndexBuffer(cmd, r.IndexBuffer.Handle, 0, vulkan.INDEX_TYPE_UINT16)

	// Dynamic Viewport
	viewport := vulkan.Viewport{
		X:        0,
		Y:        0,
		Width:    dispSize.X,
		Height:   dispSize.Y,
		MinDepth: 0.0,
		MaxDepth: 1.0,
	}
	vulkan.CmdSetViewport(cmd, 0, []vulkan.Viewport{viewport})

	// Push Constants: scale and translation
	scaleX := 2.0 / dispSize.X
	scaleY := 2.0 / dispSize.Y
	transX := -1.0 - dispPos.X*scaleX
	transY := -1.0 - dispPos.Y*scaleY

	pushConstants := [4]float32{scaleX, scaleY, transX, transY}
	vulkan.CmdPushConstants(
		cmd,
		r.PipelineLayout,
		vulkan.SHADER_STAGE_VERTEX_BIT,
		0,
		16,
		unsafe.Pointer(&pushConstants[0]),
	)

	// Execute Draw Commands
	globalVtxOffset := uint32(0)
	globalIdxOffset := uint32(0)

	for _, cl := range cmdLists {
		cmds := cl.Commands()
		for _, drawCmd := range cmds {
			clip := drawCmd.ClipRect()

			clipMinX := math.Max(0, float64(clip.X-dispPos.X))
			clipMinY := math.Max(0, float64(clip.Y-dispPos.Y))
			clipMaxX := math.Min(float64(dispSize.X), float64(clip.Z-dispPos.X))
			clipMaxY := math.Min(float64(dispSize.Y), float64(clip.W-dispPos.Y))

			if clipMaxX <= clipMinX || clipMaxY <= clipMinY {
				continue
			}

			scissor := vulkan.Rect2D{
				Offset: vulkan.Offset2D{
					X: int32(clipMinX),
					Y: int32(clipMinY),
				},
				Extent: vulkan.Extent2D{
					Width:  uint32(clipMaxX - clipMinX),
					Height: uint32(clipMaxY - clipMinY),
				},
			}
			vulkan.CmdSetScissor(cmd, 0, []vulkan.Rect2D{scissor})

			vulkan.CmdDrawIndexed(
				cmd,
				uint32(drawCmd.ElemCount()),
				1,
				globalIdxOffset+uint32(drawCmd.IdxOffset()),
				int32(globalVtxOffset+drawCmd.VtxOffset()),
				0,
			)
		}
		_, vCount := cl.GetVertexBuffer()
		_, iCount := cl.GetIndexBuffer()
		globalVtxOffset += uint32(vCount / 20)
		globalIdxOffset += uint32(iCount / 2)
	}
}

// Destroy cleans up all Vulkan resources used by the renderer.
func (r *Renderer) Destroy() {
	if r.Ctx == nil || r.Ctx.Device == 0 {
		return
	}
	if r.VertexBuffer != nil {
		r.VertexBuffer.Destroy(r.Ctx)
		r.VertexBuffer = nil
	}
	if r.IndexBuffer != nil {
		r.IndexBuffer.Destroy(r.Ctx)
		r.IndexBuffer = nil
	}
	if r.Pipeline != 0 {
		vulkan.DestroyPipeline(r.Ctx.Device, r.Pipeline, nil)
		r.Pipeline = 0
	}
	if r.PipelineLayout != 0 {
		vulkan.DestroyPipelineLayout(r.Ctx.Device, r.PipelineLayout, nil)
		r.PipelineLayout = 0
	}
	if r.DescriptorPool != 0 {
		vulkan.DestroyDescriptorPool(r.Ctx.Device, r.DescriptorPool, nil)
		r.DescriptorPool = 0
	}
	if r.DescriptorSetLayout != 0 {
		vulkan.DestroyDescriptorSetLayout(r.Ctx.Device, r.DescriptorSetLayout, nil)
		r.DescriptorSetLayout = 0
	}
	if r.FontSampler != 0 {
		vulkan.DestroySampler(r.Ctx.Device, r.FontSampler, nil)
		r.FontSampler = 0
	}
	if r.FontView != 0 {
		vulkan.DestroyImageView(r.Ctx.Device, r.FontView, nil)
		r.FontView = 0
	}
	if r.FontImage != 0 {
		vulkan.DestroyImage(r.Ctx.Device, r.FontImage, nil)
		r.FontImage = 0
	}
	if r.FontMemory != 0 {
		vulkan.FreeMemory(r.Ctx.Device, r.FontMemory, nil)
		r.FontMemory = 0
	}
}
