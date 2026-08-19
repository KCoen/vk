package base

import (
	"fmt"
	"log"

	"time"
	"unsafe"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/go-gl/glfw/v3.3/glfw"
	"go.cld.moe/vk_google/extensions/khr_surface"
	"go.cld.moe/vk_google/extensions/khr_swapchain"
	"go.cld.moe/vk_google/samples/utility"
	"go.cld.moe/vk_google/samples/utility/cimgui_vulkan"
	"go.cld.moe/vk_google/vulkan"
)

type Vec3 = utility.Vec3

var fatalf = log.Fatalf

const maxFramesInFlight = 2

type BaseDemo struct {
	Window *glfw.Window
	Ctx    *utility.VulkanContext
	Width  uint32
	Height uint32

	// Render Targets
	// ColorImage/ColorView/ColorMemory are the offscreen render target used by RenderFrame().
	// For the interactive path, the swapchain image is passed directly to RecordFrameCommands.
	ColorFormat vulkan.Format
	ColorImage  vulkan.Image
	ColorMemory vulkan.DeviceMemory
	ColorView   vulkan.ImageView
	DepthImage  vulkan.Image
	DepthMemory vulkan.DeviceMemory
	DepthView   vulkan.ImageView

	Camera utility.FlyCamera

	OnPreRender func(cmd vulkan.CommandBuffer)
	OnRender    func(cmd vulkan.CommandBuffer)
	OnReady     func()
	OnImGui     func()
}

// CreateDepthImage allocates the depth buffer and pre-transitions it to
// VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL so it is ready for dynamic rendering.
func (s *BaseDemo) CreateDepthImage() error {
	depthImgInfo := vulkan.NewImageCreateInfo()
	depthImgInfo.ImageType = vulkan.IMAGE_TYPE_2D
	depthImgInfo.Extent = vulkan.Extent3D{Width: s.Width, Height: s.Height, Depth: 1}
	depthImgInfo.MipLevels = 1
	depthImgInfo.ArrayLayers = 1
	depthImgInfo.Format = vulkan.FORMAT_D32_SFLOAT
	depthImgInfo.Tiling = vulkan.IMAGE_TILING_OPTIMAL
	depthImgInfo.InitialLayout = vulkan.IMAGE_LAYOUT_UNDEFINED
	depthImgInfo.Usage = vulkan.IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT
	depthImgInfo.Samples = vulkan.SAMPLE_COUNT_1_BIT

	var res vulkan.Result
	s.DepthImage, res = vulkan.CreateImage(s.Ctx.Device, depthImgInfo, nil)
	if res != vulkan.SUCCESS {
		return fmt.Errorf("failed to create depth image: %v", res)
	}

	dMemReqs := vulkan.GetImageMemoryRequirements(s.Ctx.Device, s.DepthImage)
	dMemIdx, err := s.Ctx.FindMemoryType(dMemReqs.MemoryTypeBits, vulkan.MEMORY_PROPERTY_DEVICE_LOCAL_BIT)
	if err != nil {
		return err
	}
	dAlloc := vulkan.NewMemoryAllocateInfo()
	dAlloc.AllocationSize = dMemReqs.Size
	dAlloc.MemoryTypeIndex = dMemIdx
	s.DepthMemory, res = vulkan.AllocateMemory(s.Ctx.Device, dAlloc, nil)
	if res != vulkan.SUCCESS {
		return fmt.Errorf("failed to allocate depth memory: %v", res)
	}
	vulkan.BindImageMemory(s.Ctx.Device, s.DepthImage, s.DepthMemory, 0)

	dViewInfo := vulkan.NewImageViewCreateInfo()
	dViewInfo.Image = s.DepthImage
	dViewInfo.ViewType = vulkan.IMAGE_VIEW_TYPE_2D
	dViewInfo.Format = vulkan.FORMAT_D32_SFLOAT
	dViewInfo.SubresourceRange = vulkan.ImageSubresourceRange{
		AspectMask: vulkan.IMAGE_ASPECT_DEPTH_BIT,
		LevelCount: 1,
		LayerCount: 1,
	}
	s.DepthView, res = vulkan.CreateImageView(s.Ctx.Device, dViewInfo, nil)
	if res != vulkan.SUCCESS {
		return fmt.Errorf("failed to create depth image view: %v", res)
	}

	// Pre-transition depth image to DEPTH_ATTACHMENT_OPTIMAL once at init.
	return s.Ctx.ExecuteOneTimeCommands(func(cmd vulkan.CommandBuffer) {
		barrier := vulkan.NewImageMemoryBarrier()
		barrier.OldLayout = vulkan.IMAGE_LAYOUT_UNDEFINED
		barrier.NewLayout = vulkan.IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL
		barrier.SrcQueueFamilyIndex = vulkan.QUEUE_FAMILY_IGNORED
		barrier.DstQueueFamilyIndex = vulkan.QUEUE_FAMILY_IGNORED
		barrier.Image = s.DepthImage
		barrier.SubresourceRange = vulkan.ImageSubresourceRange{
			AspectMask: vulkan.IMAGE_ASPECT_DEPTH_BIT,
			LevelCount: 1,
			LayerCount: 1,
		}
		barrier.SrcAccessMask = 0
		barrier.DstAccessMask = vulkan.ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT | vulkan.ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT
		vulkan.CmdPipelineBarrier(
			cmd,
			vulkan.PIPELINE_STAGE_TOP_OF_PIPE_BIT,
			vulkan.PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT,
			0, nil, nil,
			[]vulkan.ImageMemoryBarrier{*barrier},
		)
	})
}

func (s *BaseDemo) Init() {
	// 1. Initialize GLFW
	if err := glfw.Init(); err != nil {
		fatalf("Failed to initialize GLFW: %v", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	width, height := uint32(1280), uint32(720)
	s.Width, s.Height = width, height
	var err error
	s.Window, err = glfw.CreateWindow(int(width), int(height), "Vulkan Go: Multi-Draw Indirect (cimgui-go)", nil, nil)
	if err != nil {
		fatalf("Failed to create GLFW window: %v", err)
	}

	// Retrieve required GLFW instance extensions for surface creation
	glfwExts := s.Window.GetRequiredInstanceExtensions()

	// 2. Initialize Vulkan Context with required extensions
	cfg := utility.ContextConfig{
		AppName:            "MultiDrawIndirectGoSample",
		RequireMDI:         true,
		EnableValidation:   true,
		InstanceExtensions: glfwExts,
		DeviceExtensions:   []string{"VK_KHR_swapchain"},
	}

	ctx, err := utility.NewVulkanContext(cfg)
	if err != nil {
		fatalf("Failed to initialize Vulkan context: %v", err)
	}
	s.Ctx = ctx
	defer ctx.Destroy()

	// Initialize extensions
	khr_surface.Init(ctx.Instance, ctx.Device)
	khr_swapchain.Init(ctx.Instance, ctx.Device)

	fmt.Printf("Device: %s\n", vulkan.ByteSliceToString(ctx.DeviceProperties.DeviceName[:]))
	fmt.Printf("Multi-Draw Indirect Supported: %v\n", ctx.SupportsMDI)
	fmt.Printf("Draw Indirect First Instance Supported: %v\n", ctx.SupportsFirstInstance)

	// 3. Create Surface
	surfacePtr, err := s.Window.CreateWindowSurface((*byte)(unsafe.Pointer(ctx.Instance)), nil)
	if err != nil {
		fatalf("Failed to create window surface: %v", err)
	}
	surface := *(*vulkan.SurfaceKHR)(unsafe.Pointer(surfacePtr))
	defer khr_surface.DestroySurfaceKHR(ctx.Instance, surface, nil)

	// Verify surface support on queue
	var supported vulkan.Bool32
	supported, _ = khr_surface.GetPhysicalDeviceSurfaceSupportKHR(ctx.PhysicalDevice, ctx.GraphicsQueueFamily, surface)
	if supported == vulkan.False {
		fatalf("Queue family does not support presentation to surface")
	}

	// 4. Choose Surface Format and ColorSpace
	formats, res := khr_surface.GetPhysicalDeviceSurfaceFormatsKHR(ctx.PhysicalDevice, surface)
	if res != vulkan.SUCCESS || len(formats) == 0 {
		fatalf("Failed to find supported surface formats: %v", res)
	}
	chosenFormat := vulkan.FORMAT_B8G8R8A8_UNORM
	chosenColorSpace := vulkan.COLOR_SPACE_SRGB_NONLINEAR_KHR
	for _, f := range formats {
		if f.Format == vulkan.FORMAT_B8G8R8A8_UNORM || f.Format == vulkan.FORMAT_R8G8B8A8_UNORM {
			chosenFormat = f.Format
			chosenColorSpace = f.ColorSpace
			break
		}
	}
	s.ColorFormat = chosenFormat

	// 5. Create Swapchain
	caps, res := khr_surface.GetPhysicalDeviceSurfaceCapabilitiesKHR(ctx.PhysicalDevice, surface)
	if res != vulkan.SUCCESS {
		fatalf("Failed to get surface capabilities: %v", res)
	}

	imageCount := caps.MinImageCount + 1
	if caps.MaxImageCount > 0 && imageCount > caps.MaxImageCount {
		imageCount = caps.MaxImageCount
	}

	swapchainInfo := vulkan.NewSwapchainCreateInfoKHR()
	swapchainInfo.Surface = surface
	swapchainInfo.MinImageCount = imageCount
	swapchainInfo.ImageFormat = chosenFormat
	swapchainInfo.ImageColorSpace = chosenColorSpace
	swapchainInfo.ImageExtent = vulkan.Extent2D{Width: width, Height: height}
	swapchainInfo.ImageArrayLayers = 1
	swapchainInfo.ImageUsage = vulkan.IMAGE_USAGE_COLOR_ATTACHMENT_BIT
	swapchainInfo.ImageSharingMode = vulkan.SHARING_MODE_EXCLUSIVE
	swapchainInfo.PreTransform = caps.CurrentTransform
	swapchainInfo.CompositeAlpha = vulkan.COMPOSITE_ALPHA_OPAQUE_BIT_KHR
	swapchainInfo.PresentMode = vulkan.PRESENT_MODE_FIFO_KHR
	swapchainInfo.Clipped = vulkan.True

	swapchain, res := khr_swapchain.CreateSwapchainKHR(ctx.Device, swapchainInfo, nil)
	if res != vulkan.SUCCESS {
		fatalf("Failed to create swapchain: %v", res)
	}
	defer khr_swapchain.DestroySwapchainKHR(ctx.Device, swapchain, nil)

	// Get swapchain images
	images, res := khr_swapchain.GetSwapchainImagesKHR(ctx.Device, swapchain)
	if res != vulkan.SUCCESS {
		fatalf("Failed to get swapchain images: %v", res)
	}

	if err := s.CreateDepthImage(); err != nil {
		fatalf("Failed to create depth image: %v", err)
	}

	// Create swapchain image views
	imageViews := make([]vulkan.ImageView, len(images))
	for i, img := range images {
		viewInfo := vulkan.NewImageViewCreateInfo()
		viewInfo.Image = img
		viewInfo.ViewType = vulkan.IMAGE_VIEW_TYPE_2D
		viewInfo.Format = chosenFormat
		viewInfo.SubresourceRange = vulkan.ImageSubresourceRange{
			AspectMask: vulkan.IMAGE_ASPECT_COLOR_BIT,
			LevelCount: 1,
			LayerCount: 1,
		}
		imageViews[i], res = vulkan.CreateImageView(ctx.Device, viewInfo, nil)
		if res != vulkan.SUCCESS {
			fatalf("Failed to create swapchain image view %d: %v", i, res)
		}
		defer vulkan.DestroyImageView(ctx.Device, imageViews[i], nil)
	}

	// 7. Initialize cimgui-go
	imgui.CreateContext()
	defer imgui.DestroyContext()

	io := imgui.CurrentIO()
	io.SetDisplaySize(imgui.Vec2{X: float32(width), Y: float32(height)})
	io.SetBackendFlags(io.BackendFlags() | imgui.BackendFlagsRendererHasTextures)

	// 8. Create cimgui Vulkan Renderer
	imRenderer, err := cimgui_vulkan.NewRenderer(ctx, chosenFormat)
	if err != nil {
		fatalf("Failed to initialize cimgui Vulkan Renderer: %v", err)
	}
	defer imRenderer.Destroy()

	// 9. (No framebuffers needed — dynamic rendering uses image views directly)

	// 10. Create Sync objects
	imageAvailableSemaphores := make([]vulkan.Semaphore, maxFramesInFlight)
	renderFinishedSemaphores := make([]vulkan.Semaphore, len(images))
	inFlightFences := make([]vulkan.Fence, maxFramesInFlight)

	semInfo := vulkan.NewSemaphoreCreateInfo()
	fenceInfo := vulkan.NewFenceCreateInfo()
	fenceInfo.Flags = vulkan.FENCE_CREATE_SIGNALED_BIT

	for i := 0; i < maxFramesInFlight; i++ {
		imageAvailableSemaphores[i], _ = vulkan.CreateSemaphore(ctx.Device, semInfo, nil)
		inFlightFences[i], _ = vulkan.CreateFence(ctx.Device, fenceInfo, nil)

		defer vulkan.DestroySemaphore(ctx.Device, imageAvailableSemaphores[i], nil)
		defer vulkan.DestroyFence(ctx.Device, inFlightFences[i], nil)
	}

	for i := 0; i < len(images); i++ {
		renderFinishedSemaphores[i], _ = vulkan.CreateSemaphore(ctx.Device, semInfo, nil)
		defer vulkan.DestroySemaphore(ctx.Device, renderFinishedSemaphores[i], nil)
	}

	// Allocate reusable command buffers (one per frame in flight)
	commandBuffers := make([]vulkan.CommandBuffer, maxFramesInFlight)
	allocInfo := vulkan.NewCommandBufferAllocateInfo()
	allocInfo.CommandPool = ctx.CommandPool
	allocInfo.Level = vulkan.COMMAND_BUFFER_LEVEL_PRIMARY
	allocInfo.CommandBufferCount = 1

	for i := 0; i < maxFramesInFlight; i++ {
		commandBuffers[i], res = vulkan.AllocateCommandBuffers(ctx.Device, allocInfo)
		if res != vulkan.SUCCESS {
			fatalf("Failed to allocate command buffer %d: %v", i, res)
		}
	}

	// 11. Set Up Camera and Input Controls
	s.Camera = utility.NewFlyCamera(
		Vec3{X: 0, Y: 10, Z: 25},
		Vec3{X: 0, Y: 0, Z: 0},
		utility.Vec2{X: float32(width), Y: float32(height)},
	)

	// Keep track of keys
	keys := make(map[glfw.Key]bool)
	s.Window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Press {
			keys[key] = true
		} else if action == glfw.Release {
			keys[key] = false
		}
	})

	s.Window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		if button == glfw.MouseButtonLeft {
			if action == glfw.Press {
				io.SetMouseButtonDown(0, true)
				if !io.WantCaptureMouse() {
					s.Camera.BeginDrag()
				}
			} else if action == glfw.Release {
				io.SetMouseButtonDown(0, false)
				s.Camera.EndDrag()
			}
		}
	})

	s.Window.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		io.SetMousePos(imgui.Vec2{X: float32(xpos), Y: float32(ypos)})
		if s.Camera.IsDragging() && !io.WantCaptureMouse() {
			s.Camera.MouseMove(xpos, ypos)
		}
	})

	s.OnReady()

	// Render loop
	lastTime := time.Now()
	currentFrame := 0

	for !s.Window.ShouldClose() {
		glfw.PollEvents()

		// Calculate delta time
		now := time.Now()
		deltaTime := now.Sub(lastTime).Seconds()
		lastTime = now

		// Update ImGui Input IO Delta Time
		io.SetDeltaTime(float32(deltaTime))

		// Start new ImGui frame
		imgui.NewFrame()

		s.OnImGui()

		// Finalize ImGui geometry
		imgui.Render()

		// 12. Update Camera position from inputs
		if !io.WantCaptureMouse() {
			dt := float32(deltaTime)
			if keys[glfw.KeyW] {
				s.Camera.MoveForward(dt)
			}
			if keys[glfw.KeyS] {
				s.Camera.MoveBack(dt)
			}
			if keys[glfw.KeyA] {
				s.Camera.MoveLeft(dt)
			}
			if keys[glfw.KeyD] {
				s.Camera.MoveRight(dt)
			}
		}

		// 13. Render Loop Frame
		// Wait for fence of current frame in flight
		vulkan.WaitForFences(ctx.Device, []vulkan.Fence{inFlightFences[currentFrame]}, vulkan.True, ^uint64(0))
		vulkan.ResetFences(ctx.Device, []vulkan.Fence{inFlightFences[currentFrame]})

		// Acquire next image
		imageIndex, res := khr_swapchain.AcquireNextImageKHR(ctx.Device, swapchain, ^uint64(0), imageAvailableSemaphores[currentFrame], 0)
		if res != vulkan.SUCCESS {
			continue
		}

		cmd := commandBuffers[currentFrame]

		beginInfo := vulkan.NewCommandBufferBeginInfo()
		beginInfo.Flags = vulkan.COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT
		vulkan.BeginCommandBuffer(cmd, beginInfo)

		if s.OnPreRender != nil {
			s.OnPreRender(cmd)
		}

		// 2. Transition swapchain image: UNDEFINED → COLOR_ATTACHMENT_OPTIMAL
		// (skipped for the internal offscreen ColorImage which is already in COLOR_ATTACHMENT_OPTIMAL)
		if images[imageIndex] != s.ColorImage {
			colorBarrier := vulkan.NewImageMemoryBarrier()
			colorBarrier.OldLayout = vulkan.IMAGE_LAYOUT_UNDEFINED
			colorBarrier.NewLayout = vulkan.IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
			colorBarrier.SrcQueueFamilyIndex = vulkan.QUEUE_FAMILY_IGNORED
			colorBarrier.DstQueueFamilyIndex = vulkan.QUEUE_FAMILY_IGNORED
			colorBarrier.Image = images[imageIndex]
			colorBarrier.SrcAccessMask = 0
			colorBarrier.DstAccessMask = vulkan.ACCESS_COLOR_ATTACHMENT_WRITE_BIT
			colorBarrier.SubresourceRange = vulkan.ImageSubresourceRange{
				AspectMask: vulkan.IMAGE_ASPECT_COLOR_BIT,
				LevelCount: 1,
				LayerCount: 1,
			}
			vulkan.CmdPipelineBarrier(
				cmd,
				vulkan.PIPELINE_STAGE_TOP_OF_PIPE_BIT,
				vulkan.PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT,
				0, nil, nil,
				[]vulkan.ImageMemoryBarrier{*colorBarrier},
			)
		}

		// 3. Begin Dynamic Rendering
		var clearColor vulkan.ClearValue
		clearColor.Color = vulkan.ClearColorValue{Float32: [4]float32{0.1, 0.1, 0.15, 1.0}}

		colorAttach := vulkan.NewRenderingAttachmentInfo()
		colorAttach.ImageView = imageViews[imageIndex]
		colorAttach.ImageLayout = vulkan.IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
		colorAttach.LoadOp = vulkan.ATTACHMENT_LOAD_OP_CLEAR
		colorAttach.StoreOp = vulkan.ATTACHMENT_STORE_OP_STORE
		colorAttach.ClearValue = clearColor

		var clearDepth vulkan.ClearValue
		*clearDepth.DepthStencil() = vulkan.ClearDepthStencilValue{Depth: 1.0, Stencil: 0}

		depthAttach := vulkan.NewRenderingAttachmentInfo()
		depthAttach.ImageView = s.DepthView
		depthAttach.ImageLayout = vulkan.IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL
		depthAttach.LoadOp = vulkan.ATTACHMENT_LOAD_OP_CLEAR
		depthAttach.StoreOp = vulkan.ATTACHMENT_STORE_OP_DONT_CARE
		depthAttach.ClearValue = clearDepth

		ri := vulkan.NewRenderingInfo()
		ri.RenderArea = vulkan.Rect2D{
			Offset: vulkan.Offset2D{X: 0, Y: 0},
			Extent: vulkan.Extent2D{Width: width, Height: height},
		}
		ri.LayerCount = 1
		ri.ColorAttachments = []vulkan.RenderingAttachmentInfo{*colorAttach}
		ri.DepthAttachment = depthAttach

		vulkan.CmdBeginRendering(cmd, ri)

		s.OnRender(cmd)

		imRenderer.RenderDrawData(imgui.CurrentDrawData(), cmd)

		vulkan.CmdEndRendering(cmd)

		// 8. Transition swapchain image: COLOR_ATTACHMENT_OPTIMAL → PRESENT_SRC_KHR
		// (skipped for the internal offscreen ColorImage; no presentation required)
		if images[imageIndex] != s.ColorImage {
			presentBarrier := vulkan.NewImageMemoryBarrier()
			presentBarrier.OldLayout = vulkan.IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
			presentBarrier.NewLayout = vulkan.IMAGE_LAYOUT_PRESENT_SRC_KHR
			presentBarrier.SrcQueueFamilyIndex = vulkan.QUEUE_FAMILY_IGNORED
			presentBarrier.DstQueueFamilyIndex = vulkan.QUEUE_FAMILY_IGNORED
			presentBarrier.Image = images[imageIndex]
			presentBarrier.SrcAccessMask = vulkan.ACCESS_COLOR_ATTACHMENT_WRITE_BIT
			presentBarrier.DstAccessMask = 0
			presentBarrier.SubresourceRange = vulkan.ImageSubresourceRange{
				AspectMask: vulkan.IMAGE_ASPECT_COLOR_BIT,
				LevelCount: 1,
				LayerCount: 1,
			}
			vulkan.CmdPipelineBarrier(
				cmd,
				vulkan.PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT,
				vulkan.PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT,
				0, nil, nil,
				[]vulkan.ImageMemoryBarrier{*presentBarrier},
			)
		}

		vulkan.EndCommandBuffer(cmd)

		// Submit command buffer
		submitInfo := vulkan.NewSubmitInfo()
		submitInfo.WaitSemaphores = []vulkan.Semaphore{imageAvailableSemaphores[currentFrame]}
		submitInfo.WaitDstStageMask = []vulkan.PipelineStageFlags{vulkan.PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT}
		submitInfo.CommandBuffers = []vulkan.CommandBuffer{cmd}
		// Signal the semaphore corresponding to the acquired image index
		submitInfo.SignalSemaphores = []vulkan.Semaphore{renderFinishedSemaphores[imageIndex]}
		vulkan.QueueSubmit(ctx.GraphicsQueue, []vulkan.SubmitInfo{*submitInfo}, inFlightFences[currentFrame])

		// Present swapchain image
		presentInfo := vulkan.NewPresentInfoKHR()
		presentInfo.WaitSemaphores = []vulkan.Semaphore{renderFinishedSemaphores[imageIndex]}
		presentInfo.Swapchains = []vulkan.SwapchainKHR{swapchain}
		presentInfo.ImageIndices = []uint32{imageIndex}
		khr_swapchain.QueuePresentKHR(ctx.GraphicsQueue, presentInfo)

		// Advance to next frame in flight
		currentFrame = (currentFrame + 1) % maxFramesInFlight
	}

	vulkan.DeviceWaitIdle(ctx.Device)
	fmt.Println("=== Window closed. Exit ===")
}
