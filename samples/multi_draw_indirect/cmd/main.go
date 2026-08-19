package main

import (
	"fmt"

	"runtime"

	"github.com/AllenDang/cimgui-go/imgui"
	"go.cld.moe/vk_google/samples/multi_draw_indirect"
	"go.cld.moe/vk_google/vulkan"
)

func init() {
	// GLFW requires all calls to happen on the main OS thread.
	runtime.LockOSThread()
}

func main() {
	fmt.Println("=== Vulkan Go Port: Multi-Draw Indirect Sample with cimgui-go ===")

	sample := &multi_draw_indirect.MultiDrawIndirectSample{
		RenderMode: multi_draw_indirect.RenderModeGPU,
		EnableMDI:  true,
		FreezeCull: false,
	}

	sample.OnReady = func() {
		if err := sample.Prepare(); err != nil {
			sample.Destroy()
			return
		}
	}

	sample.OnPreRender = func(cmd vulkan.CommandBuffer) {
		sample.UpdateSceneUniform()
		if !sample.FreezeCull {
			if sample.RenderMode == multi_draw_indirect.RenderModeGPU {
				sample.GPUCull(cmd)
			} else {
				sample.CPUCull()
			}
		}
	}

	sample.OnRender = func(cmd vulkan.CommandBuffer) {
		sample.RecordFrameCommands(cmd)
	}

	sample.OnImGui = func() {
		// Emit Dear ImGui Demo Window
		imgui.ShowDemoWindow()

		// Emit Sample Controls Window
		if imgui.Begin("Sample Controls") {
			imgui.Checkbox("Enable Multi-Draw Indirect", &sample.EnableMDI)

			isGPU := (sample.RenderMode == multi_draw_indirect.RenderModeGPU)
			if imgui.Checkbox("Use GPU Compute Culling", &isGPU) {
				if isGPU {
					sample.RenderMode = multi_draw_indirect.RenderModeGPU
				} else {
					sample.RenderMode = multi_draw_indirect.RenderModeCPU
				}
			}
			imgui.Checkbox("Freeze Frustum Culling", &sample.FreezeCull)

			imgui.Separator()

			cubeCount := int32(len(sample.Models))
			if imgui.SliderInt("Cube Count", &cubeCount, 1, 350000) {
				if int(cubeCount) != len(sample.Models) && cubeCount > 0 {
					if err := sample.RebuildScene(int(cubeCount)); err != nil {
						fmt.Printf("RebuildScene error: %v\n", err)
					}
				}
			}

			imgui.Separator()

			imgui.Text(fmt.Sprintf("Total Cubes: %d", len(sample.Models)))
			imgui.Text(fmt.Sprintf("Camera Pos: (%.1f, %.1f, %.1f)", sample.Camera.Eye.X, sample.Camera.Eye.Y, sample.Camera.Eye.Z))

			if sample.RenderMode == multi_draw_indirect.RenderModeGPU {
				imgui.Text("Mode: GPU (Compute Frustum Culling)")
			} else {
				visibleCPU := sample.CPUCull()
				imgui.Text(fmt.Sprintf("Visible (CPU): %d / %d", visibleCPU, len(sample.Models)))
			}
			imgui.End()
		}
	}

	// Print instruction tips
	fmt.Println("\nControls:")
	fmt.Println(" - Drag with Left Mouse Button to look around")
	fmt.Println(" - Press W, A, S, D to fly around the scene")
	fmt.Println(" - Interact with Dear ImGui windows and controls")
	fmt.Println(" - Close the window to exit")

	sample.BaseDemo.Init()

	fmt.Println("=== Demo Over Exit ===")
}
