package multi_draw_indirect

import (
	"fmt"
	"math"
	"unsafe"

	"go.cld.moe/vk_google/samples/base"
	"go.cld.moe/vk_google/samples/utility"
	"go.cld.moe/vk_google/vulkan"
)

type RenderMode int

const (
	RenderModeCPU RenderMode = iota
	RenderModeGPU
)

type Vertex struct {
	Position [3]float32
	UV       [2]float32
}

type GpuModelInformation struct {
	CenterX      float32
	CenterY      float32
	CenterZ      float32
	Radius       float32
	TextureIndex uint32
	FirstIndex   uint32
	IndexCount   uint32
	Pad          uint32
}

type SceneUniform struct {
	View       utility.Mat4
	Proj       utility.Mat4
	ProjView   utility.Mat4
	ModelCount uint32
	Pad        [3]uint32
}

type SceneModel struct {
	Vertices           []Vertex
	Indices            []uint16
	VertexBufferOffset uint64
	IndexBufferOffset  uint64
	TextureIndex       uint32
	BoundingSphere     utility.BoundingSphere
}

type MultiDrawIndirectSample struct {
	base.BaseDemo

	RenderMode RenderMode
	EnableMDI  bool
	FreezeCull bool

	Models      []SceneModel
	CPUCommands []vulkan.DrawIndexedIndirectCommand

	VertexBuffer       *utility.Buffer
	IndexBuffer        *utility.Buffer
	ModelInfoBuffer    *utility.Buffer
	SceneUniformBuffer *utility.Buffer
	IndirectCallBuffer *utility.Buffer
	CpuStagingBuffer   *utility.Buffer

	Textures []*utility.Texture

	RenderShader *utility.ShaderPipeline
	CullShader   *utility.ShaderPipeline
}

func (s *MultiDrawIndirectSample) Prepare() error {
	s.generateScene(64)

	if err := s.createTextures(); err != nil {
		return fmt.Errorf("createTextures failed: %w", err)
	}

	if err := s.initBuffers(); err != nil {
		return fmt.Errorf("initBuffers failed: %w", err)
	}

	cullShader, err := utility.PrepareComputeShader(s.Ctx, CullShaderSPV)
	if err != nil {
		return fmt.Errorf("PrepareComputeShader failed: %w", err)
	}
	s.CullShader = cullShader
	s.CullShader.BindResources(s.ModelInfoBuffer, s.SceneUniformBuffer, s.IndirectCallBuffer)

	colorFmt := s.ColorFormat
	if colorFmt == 0 {
		colorFmt = vulkan.FORMAT_B8G8R8A8_UNORM
	}

	renderShader, err := utility.PrepareGraphicsShader(s.Ctx, utility.GraphicsPipelineConfig{
		VertSPV: VertShaderSPV,
		FragSPV: FragShaderSPV,
		VertexBindings: []vulkan.VertexInputBindingDescription{
			{Binding: 0, Stride: uint32(unsafe.Sizeof(Vertex{})), InputRate: vulkan.VERTEX_INPUT_RATE_VERTEX},
			{Binding: 1, Stride: uint32(unsafe.Sizeof(GpuModelInformation{})), InputRate: vulkan.VERTEX_INPUT_RATE_INSTANCE},
		},
		VertexAttributes: []vulkan.VertexInputAttributeDescription{
			{Location: 0, Binding: 0, Format: vulkan.FORMAT_R32G32B32_SFLOAT, Offset: uint32(unsafe.Offsetof(Vertex{}.Position))},
			{Location: 1, Binding: 0, Format: vulkan.FORMAT_R32G32_SFLOAT, Offset: uint32(unsafe.Offsetof(Vertex{}.UV))},
			{Location: 2, Binding: 1, Format: vulkan.FORMAT_R32G32B32_SFLOAT, Offset: uint32(unsafe.Offsetof(GpuModelInformation{}.CenterX))},
			{Location: 3, Binding: 1, Format: vulkan.FORMAT_R32_SFLOAT, Offset: uint32(unsafe.Offsetof(GpuModelInformation{}.Radius))},
			{Location: 4, Binding: 1, Format: vulkan.FORMAT_R32_UINT, Offset: uint32(unsafe.Offsetof(GpuModelInformation{}.TextureIndex))},
		},
		ColorFormat: colorFmt,
		DepthFormat: vulkan.FORMAT_D32_SFLOAT,
		Width:       s.Width,
		Height:      s.Height,
	})
	if err != nil {
		return fmt.Errorf("PrepareGraphicsShader failed: %w", err)
	}
	s.RenderShader = renderShader
	s.RenderShader.BindResources(s.Textures, s.SceneUniformBuffer)

	s.UpdateSceneUniform()
	s.CPUCull()

	return nil
}

func (s *MultiDrawIndirectSample) generateScene(numObjects int) {
	gridDim := int(math.Ceil(math.Pow(float64(numObjects), 1.0/3.0)))
	spacing := float32(2.5)

	cubeVertices := []Vertex{
		{Position: [3]float32{-0.5, -0.5, 0.5}, UV: [2]float32{0, 0}},
		{Position: [3]float32{0.5, -0.5, 0.5}, UV: [2]float32{1, 0}},
		{Position: [3]float32{0.5, 0.5, 0.5}, UV: [2]float32{1, 1}},
		{Position: [3]float32{-0.5, 0.5, 0.5}, UV: [2]float32{0, 1}},
		{Position: [3]float32{0.5, -0.5, -0.5}, UV: [2]float32{0, 0}},
		{Position: [3]float32{-0.5, -0.5, -0.5}, UV: [2]float32{1, 0}},
		{Position: [3]float32{-0.5, 0.5, -0.5}, UV: [2]float32{1, 1}},
		{Position: [3]float32{0.5, 0.5, -0.5}, UV: [2]float32{0, 1}},
		{Position: [3]float32{-0.5, -0.5, -0.5}, UV: [2]float32{0, 0}},
		{Position: [3]float32{-0.5, -0.5, 0.5}, UV: [2]float32{1, 0}},
		{Position: [3]float32{-0.5, 0.5, 0.5}, UV: [2]float32{1, 1}},
		{Position: [3]float32{-0.5, 0.5, -0.5}, UV: [2]float32{0, 1}},
		{Position: [3]float32{0.5, -0.5, 0.5}, UV: [2]float32{0, 0}},
		{Position: [3]float32{0.5, -0.5, -0.5}, UV: [2]float32{1, 0}},
		{Position: [3]float32{0.5, 0.5, -0.5}, UV: [2]float32{1, 1}},
		{Position: [3]float32{0.5, 0.5, 0.5}, UV: [2]float32{0, 1}},
		{Position: [3]float32{-0.5, 0.5, 0.5}, UV: [2]float32{0, 0}},
		{Position: [3]float32{0.5, 0.5, 0.5}, UV: [2]float32{1, 0}},
		{Position: [3]float32{0.5, 0.5, -0.5}, UV: [2]float32{1, 1}},
		{Position: [3]float32{-0.5, 0.5, -0.5}, UV: [2]float32{0, 1}},
		{Position: [3]float32{-0.5, -0.5, -0.5}, UV: [2]float32{0, 0}},
		{Position: [3]float32{0.5, -0.5, -0.5}, UV: [2]float32{1, 0}},
		{Position: [3]float32{0.5, -0.5, 0.5}, UV: [2]float32{1, 1}},
		{Position: [3]float32{-0.5, -0.5, 0.5}, UV: [2]float32{0, 1}},
	}

	cubeIndices := []uint16{
		0, 1, 2, 2, 3, 0,
		4, 5, 6, 6, 7, 4,
		8, 9, 10, 10, 11, 8,
		12, 13, 14, 14, 15, 12,
		16, 17, 18, 18, 19, 16,
		20, 21, 22, 22, 23, 20,
	}

	s.Models = make([]SceneModel, 0, numObjects)
	count := 0

	for x := 0; x < gridDim && count < numObjects; x++ {
		for y := 0; y < gridDim && count < numObjects; y++ {
			for z := 0; z < gridDim && count < numObjects; z++ {
				offsetX := (float32(x) - float32(gridDim)/2.0) * spacing
				offsetY := (float32(y) - float32(gridDim)/2.0) * spacing
				offsetZ := (float32(z) - float32(gridDim)/2.0) * spacing

				verts := make([]Vertex, len(cubeVertices))
				pts := make([]utility.Vec3, len(cubeVertices))

				for i, v := range cubeVertices {
					px, py, pz := v.Position[0]+offsetX, v.Position[1]+offsetY, v.Position[2]+offsetZ
					verts[i] = Vertex{Position: [3]float32{px, py, pz}, UV: v.UV}
					pts[i] = utility.Vec3{X: px, Y: py, Z: pz}
				}

				indices := make([]uint16, len(cubeIndices))
				copy(indices, cubeIndices)

				s.Models = append(s.Models, SceneModel{
					Vertices:       verts,
					Indices:        indices,
					TextureIndex:   uint32(count % 16),
					BoundingSphere: utility.ComputeBoundingSphere(pts),
				})
				count++
			}
		}
	}
}

func (s *MultiDrawIndirectSample) createTextures() error {
	s.Textures = make([]*utility.Texture, 16)
	colors := [][4]byte{
		{255, 60, 60, 255}, {60, 255, 60, 255}, {60, 60, 255, 255}, {255, 255, 60, 255},
		{255, 60, 255, 255}, {60, 255, 255, 255}, {255, 128, 0, 255}, {128, 0, 255, 255},
		{0, 255, 128, 255}, {255, 192, 203, 255}, {128, 255, 0, 255}, {0, 128, 255, 255},
		{210, 105, 30, 255}, {192, 192, 192, 255}, {255, 215, 0, 255}, {240, 240, 240, 255},
	}

	for i := 0; i < 16; i++ {
		c := colors[i]
		pixels := make([]byte, 4*4*4)
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				idx := (y*4 + x) * 4
				if (x+y)%2 == 0 {
					pixels[idx], pixels[idx+1], pixels[idx+2], pixels[idx+3] = c[0], c[1], c[2], c[3]
				} else {
					pixels[idx], pixels[idx+1], pixels[idx+2], pixels[idx+3] = c[0]/2, c[1]/2, c[2]/2, c[3]
				}
			}
		}

		tex, err := utility.CreateTexture2D(s.Ctx, 4, 4, vulkan.FORMAT_R8G8B8A8_UNORM, pixels)
		if err != nil {
			return fmt.Errorf("failed to create texture %d: %w", i, err)
		}
		s.Textures[i] = tex
	}
	return nil
}

func (s *MultiDrawIndirectSample) initBuffers() error {
	var totalVertices []Vertex
	var totalIndices []uint16
	modelInfos := make([]GpuModelInformation, len(s.Models))

	var vertOffset, idxOffset uint64

	for i := range s.Models {
		m := &s.Models[i]
		m.VertexBufferOffset = vertOffset
		m.IndexBufferOffset = idxOffset

		firstIndex := uint32(len(totalIndices))
		indexCount := uint32(len(m.Indices))

		totalVertices = append(totalVertices, m.Vertices...)
		totalIndices = append(totalIndices, m.Indices...)

		modelInfos[i] = GpuModelInformation{
			CenterX:      m.BoundingSphere.Center.X,
			CenterY:      m.BoundingSphere.Center.Y,
			CenterZ:      m.BoundingSphere.Center.Z,
			Radius:       m.BoundingSphere.Radius,
			TextureIndex: m.TextureIndex,
			FirstIndex:   firstIndex,
			IndexCount:   indexCount,
		}

		vertOffset += uint64(len(m.Vertices)) * uint64(unsafe.Sizeof(Vertex{}))
		idxOffset += uint64(len(m.Indices)) * uint64(unsafe.Sizeof(uint16(0)))
	}

	vertBytes := unsafe.Slice((*byte)(unsafe.Pointer(&totalVertices[0])), len(totalVertices)*int(unsafe.Sizeof(Vertex{})))
	vBuf, err := utility.CreateDeviceLocalBufferFromData(s.Ctx, vertBytes, vulkan.BUFFER_USAGE_VERTEX_BUFFER_BIT|vulkan.BUFFER_USAGE_STORAGE_BUFFER_BIT)
	if err != nil {
		return err
	}
	s.VertexBuffer = vBuf

	idxBytes := unsafe.Slice((*byte)(unsafe.Pointer(&totalIndices[0])), len(totalIndices)*int(unsafe.Sizeof(uint16(0))))
	iBuf, err := utility.CreateDeviceLocalBufferFromData(s.Ctx, idxBytes, vulkan.BUFFER_USAGE_INDEX_BUFFER_BIT)
	if err != nil {
		return err
	}
	s.IndexBuffer = iBuf

	modelBytes := unsafe.Slice((*byte)(unsafe.Pointer(&modelInfos[0])), len(modelInfos)*int(unsafe.Sizeof(GpuModelInformation{})))
	mBuf, err := utility.CreateDeviceLocalBufferFromData(s.Ctx, modelBytes, vulkan.BUFFER_USAGE_STORAGE_BUFFER_BIT|vulkan.BUFFER_USAGE_VERTEX_BUFFER_BIT)
	if err != nil {
		return err
	}
	s.ModelInfoBuffer = mBuf

	if s.SceneUniformBuffer == nil {
		uBuf, err := utility.CreateBuffer(s.Ctx, uint64(unsafe.Sizeof(SceneUniform{})), vulkan.BUFFER_USAGE_UNIFORM_BUFFER_BIT, vulkan.MEMORY_PROPERTY_HOST_VISIBLE_BIT|vulkan.MEMORY_PROPERTY_HOST_COHERENT_BIT)
		if err != nil {
			return err
		}
		s.SceneUniformBuffer = uBuf
	}

	indirectSize := uint64(len(s.Models)) * uint64(unsafe.Sizeof(vulkan.DrawIndexedIndirectCommand{}))
	indBuf, err := utility.CreateBuffer(s.Ctx, indirectSize, vulkan.BUFFER_USAGE_INDIRECT_BUFFER_BIT|vulkan.BUFFER_USAGE_STORAGE_BUFFER_BIT|vulkan.BUFFER_USAGE_TRANSFER_DST_BIT|vulkan.BUFFER_USAGE_TRANSFER_SRC_BIT, vulkan.MEMORY_PROPERTY_DEVICE_LOCAL_BIT)
	if err != nil {
		return err
	}
	s.IndirectCallBuffer = indBuf

	cpuStageBuf, err := utility.CreateBuffer(s.Ctx, indirectSize, vulkan.BUFFER_USAGE_TRANSFER_SRC_BIT|vulkan.BUFFER_USAGE_TRANSFER_DST_BIT, vulkan.MEMORY_PROPERTY_HOST_VISIBLE_BIT|vulkan.MEMORY_PROPERTY_HOST_COHERENT_BIT)
	if err != nil {
		return err
	}
	s.CpuStagingBuffer = cpuStageBuf

	s.CPUCommands = make([]vulkan.DrawIndexedIndirectCommand, len(s.Models))
	for i, m := range s.Models {
		s.CPUCommands[i] = vulkan.DrawIndexedIndirectCommand{
			IndexCount:    uint32(len(m.Indices)),
			InstanceCount: 1,
			FirstIndex:    uint32(m.IndexBufferOffset / 2),
			VertexOffset:  int32(m.VertexBufferOffset / uint64(unsafe.Sizeof(Vertex{}))),
			FirstInstance: uint32(i),
		}
	}

	cmdBytes := unsafe.Slice((*byte)(unsafe.Pointer(&s.CPUCommands[0])), int(indirectSize))
	s.CpuStagingBuffer.Upload(s.Ctx, cmdBytes, 0)
	s.CpuStagingBuffer.CopyTo(s.Ctx, s.IndirectCallBuffer, indirectSize)

	return nil
}

func (s *MultiDrawIndirectSample) RebuildScene(numObjects int) error {
	if numObjects <= 0 {
		numObjects = 1
	}

	vulkan.DeviceWaitIdle(s.Ctx.Device)

	if s.VertexBuffer != nil {
		s.VertexBuffer.Destroy(s.Ctx)
		s.VertexBuffer = nil
	}
	if s.IndexBuffer != nil {
		s.IndexBuffer.Destroy(s.Ctx)
		s.IndexBuffer = nil
	}
	if s.ModelInfoBuffer != nil {
		s.ModelInfoBuffer.Destroy(s.Ctx)
		s.ModelInfoBuffer = nil
	}
	if s.IndirectCallBuffer != nil {
		s.IndirectCallBuffer.Destroy(s.Ctx)
		s.IndirectCallBuffer = nil
	}
	if s.CpuStagingBuffer != nil {
		s.CpuStagingBuffer.Destroy(s.Ctx)
		s.CpuStagingBuffer = nil
	}

	s.generateScene(numObjects)
	if err := s.initBuffers(); err != nil {
		return fmt.Errorf("rebuild initBuffers failed: %w", err)
	}

	s.RenderShader.BindResources(s.Textures, s.SceneUniformBuffer)
	s.CullShader.BindResources(s.ModelInfoBuffer, s.SceneUniformBuffer, s.IndirectCallBuffer)

	s.UpdateSceneUniform()
	s.CPUCull()

	return nil
}

func (s *MultiDrawIndirectSample) UpdateSceneUniform() {
	view, proj, projView := s.Camera.Matrices()

	uniform := SceneUniform{
		View:       view,
		Proj:       proj,
		ProjView:   projView,
		ModelCount: uint32(len(s.Models)),
	}

	uniformBytes := unsafe.Slice((*byte)(unsafe.Pointer(&uniform)), int(unsafe.Sizeof(SceneUniform{})))
	s.SceneUniformBuffer.Upload(s.Ctx, uniformBytes, 0)
}

func (s *MultiDrawIndirectSample) CPUCull() int {
	var uniform SceneUniform
	uPtr, _ := s.SceneUniformBuffer.Map(s.Ctx)
	if uPtr != nil {
		uniform = *(*SceneUniform)(uPtr)
	}

	tester := utility.NewVisibilityTester(uniform.ProjView)
	visibleCount := 0

	for i, m := range s.Models {
		isVisible := tester.IsVisible(m.BoundingSphere.Center, m.BoundingSphere.Radius)
		instanceCount := uint32(0)
		if isVisible {
			instanceCount = 1
			visibleCount++
		}

		s.CPUCommands[i] = vulkan.DrawIndexedIndirectCommand{
			IndexCount:    uint32(len(m.Indices)),
			InstanceCount: instanceCount,
			FirstIndex:    uint32(m.IndexBufferOffset / 2),
			VertexOffset:  int32(m.VertexBufferOffset / uint64(unsafe.Sizeof(Vertex{}))),
			FirstInstance: uint32(i),
		}
	}

	utility.UploadSlice(s.Ctx, s.CpuStagingBuffer, s.CPUCommands, 0)
	s.CpuStagingBuffer.CopyTo(s.Ctx, s.IndirectCallBuffer, s.IndirectCallBuffer.Size)

	return visibleCount
}

func (s *MultiDrawIndirectSample) GPUCull(cmd vulkan.CommandBuffer) {
	dispatchX := uint32((len(s.Models) + 63) / 64)
	if dispatchX == 0 {
		dispatchX = 1
	}
	s.CullShader.Dispatch(cmd, dispatchX, 1, 1)

	barrier := vulkan.NewBufferMemoryBarrier()
	barrier.SrcAccessMask = vulkan.ACCESS_SHADER_WRITE_BIT
	barrier.DstAccessMask = vulkan.ACCESS_INDIRECT_COMMAND_READ_BIT
	barrier.SrcQueueFamilyIndex = vulkan.QUEUE_FAMILY_IGNORED
	barrier.DstQueueFamilyIndex = vulkan.QUEUE_FAMILY_IGNORED
	barrier.Buffer = s.IndirectCallBuffer.Handle
	barrier.Offset = 0
	barrier.Size = vulkan.DeviceSize(s.IndirectCallBuffer.Size)

	vulkan.CmdPipelineBarrier(
		cmd,
		vulkan.PIPELINE_STAGE_COMPUTE_SHADER_BIT,
		vulkan.PIPELINE_STAGE_DRAW_INDIRECT_BIT,
		0,
		nil,
		[]vulkan.BufferMemoryBarrier{*barrier},
		nil,
	)
}

func (s *MultiDrawIndirectSample) RecordFrameCommands(cmd vulkan.CommandBuffer) {
	s.RenderShader.Bind(cmd)

	var zeroOffset vulkan.DeviceSize = 0
	vulkan.CmdBindVertexBuffers(cmd, 0, []vulkan.Buffer{s.VertexBuffer.Handle}, &zeroOffset)
	vulkan.CmdBindVertexBuffers(cmd, 1, []vulkan.Buffer{s.ModelInfoBuffer.Handle}, &zeroOffset)
	vulkan.CmdBindIndexBuffer(cmd, s.IndexBuffer.Handle, 0, vulkan.INDEX_TYPE_UINT16)

	stride := uint32(unsafe.Sizeof(vulkan.DrawIndexedIndirectCommand{}))
	drawCount := uint32(len(s.Models))

	if s.EnableMDI && s.Ctx.SupportsMDI {
		vulkan.CmdDrawIndexedIndirect(cmd, s.IndirectCallBuffer.Handle, 0, drawCount, stride)
	} else {
		for i := uint32(0); i < drawCount; i++ {
			offset := vulkan.DeviceSize(uint64(i) * uint64(stride))
			vulkan.CmdDrawIndexedIndirect(cmd, s.IndirectCallBuffer.Handle, offset, 1, stride)
		}
	}
}

func (s *MultiDrawIndirectSample) Destroy() {
	if s.Ctx == nil || s.Ctx.Device == 0 {
		return
	}
	vulkan.DeviceWaitIdle(s.Ctx.Device)

	if s.RenderShader != nil {
		s.RenderShader.Destroy()
	}
	if s.CullShader != nil {
		s.CullShader.Destroy()
	}

	for _, tex := range s.Textures {
		tex.Destroy(s.Ctx)
	}
	s.Textures = nil

	if s.VertexBuffer != nil {
		s.VertexBuffer.Destroy(s.Ctx)
	}
	if s.IndexBuffer != nil {
		s.IndexBuffer.Destroy(s.Ctx)
	}
	if s.ModelInfoBuffer != nil {
		s.ModelInfoBuffer.Destroy(s.Ctx)
	}
	if s.SceneUniformBuffer != nil {
		s.SceneUniformBuffer.Destroy(s.Ctx)
	}
	if s.IndirectCallBuffer != nil {
		s.IndirectCallBuffer.Destroy(s.Ctx)
	}
	if s.CpuStagingBuffer != nil {
		s.CpuStagingBuffer.Destroy(s.Ctx)
	}
}
