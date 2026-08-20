package multi_draw_indirect_test

import (
	"testing"

	"github.com/KCoen/vk/samples/multi_draw_indirect"
	"github.com/KCoen/vk/samples/utility"
	"github.com/KCoen/vk/vulkan"
)

func TestFrustumMath(t *testing.T) {
	eye := utility.Vec3{X: 0, Y: 0, Z: 10}
	center := utility.Vec3{X: 0, Y: 0, Z: 0}
	up := utility.Vec3{X: 0, Y: 1, Z: 0}

	view := utility.Mat4LookAt(eye, center, up)
	proj := utility.Mat4Perspective(float32(60.0*3.14159265/180.0), 1.0, 0.1, 100.0)
	projView := utility.Mat4Multiply(proj, view)

	tester := utility.NewVisibilityTester(projView)

	if !tester.IsVisible(utility.Vec3{X: 0, Y: 0, Z: 0}, 1.0) {
		t.Errorf("expected object in front of camera to be visible")
	}

	if tester.IsVisible(utility.Vec3{X: 0, Y: 0, Z: 20}, 1.0) {
		t.Errorf("expected object behind camera to be culled")
	}

	if tester.IsVisible(utility.Vec3{X: 50, Y: 0, Z: 0}, 1.0) {
		t.Errorf("expected object far to the side to be culled")
	}
}

func TestMultiDrawIndirectSample(t *testing.T) {
	if err := vulkan.Init(); err != nil {
		t.Skipf("Vulkan loader not available on host: %v", err)
	}

	cfg := utility.ContextConfig{
		AppName:          "MDI_UnitTest",
		ApiVersion:       vulkan.API_VERSION_1_3,
		EnableValidation: true,
	}

	ctx, err := utility.NewVulkanContext(cfg)
	if err != nil {
		t.Skipf("Vulkan device not available on host: %v", err)
	}
	defer ctx.Destroy()

	sample := &multi_draw_indirect.MultiDrawIndirectSample{
		RenderMode: multi_draw_indirect.RenderModeGPU,
		EnableMDI:  true,
		FreezeCull: false,
	}
	sample.Ctx = ctx
	sample.Width = 800
	sample.Height = 600
	sample.Camera = utility.NewFlyCamera(
		utility.Vec3{X: 0, Y: 10, Z: 25},
		utility.Vec3{X: 0, Y: 0, Z: 0},
		utility.Vec2{X: 800, Y: 600},
	)

	if err := sample.Prepare(); err != nil {
		t.Fatalf("Prepare failed: %v", err)
	}
	defer sample.Destroy()

	if len(sample.Models) == 0 {
		t.Fatalf("Expected scene models to be generated")
	}

	sample.UpdateSceneUniform()
	visibleCPU := sample.CPUCull()
	if visibleCPU == 0 {
		t.Errorf("Expected visible models from CPU culling, got 0")
	}

	if err := sample.RebuildScene(100); err != nil {
		t.Fatalf("RebuildScene failed: %v", err)
	}
	if len(sample.Models) != 100 {
		t.Fatalf("Expected 100 models, got %d", len(sample.Models))
	}
}
