package utility

import (
	"math"
)

// FlyCamera is a first-person fly camera driven by yaw/pitch angles.
// It handles mouse-look drag and WASD keyboard movement.
type FlyCamera struct {
	Eye   Vec3
	Yaw   float64 // radians, horizontal angle
	Pitch float64 // radians, vertical angle, clamped to ±PitchLimit

	// Movement speed in world-units per second.
	Speed float32
	// Mouse drag sensitivity in radians per pixel.
	Sensitivity float64
	// Maximum absolute pitch in radians (default π/2 − ε).
	PitchLimit float64

	// internal drag state
	isDragging bool
	firstMouse bool
	lastX      float64
	lastY      float64
	Fovy       float64
	Aspect     float64
}

func (c *FlyCamera) View() Mat4 {
	return Mat4LookAt(c.Eye, c.Center(), Vec3{X: 0, Y: 1, Z: 0})
}

func (c *FlyCamera) Proj() Mat4 {
	return Mat4Perspective(float32(c.Fovy), float32(c.Aspect), 0.1, 256.0)
}

func (c *FlyCamera) ProjView() Mat4 {
	return Mat4Multiply(c.Proj(), c.View())
}

func (c *FlyCamera) Matrices() (view, proj, projView Mat4) {
	view = c.View()
	proj = c.Proj()
	projView = Mat4Multiply(proj, view)
	return view, proj, projView
}

// NewFlyCamera creates a FlyCamera at the given eye position looking toward center.
// Yaw and pitch are derived from the initial look direction.
func NewFlyCamera(eye, center Vec3, screen Vec2) FlyCamera {
	d := Vec3Normalize(Vec3Sub(center, eye))
	pitch := math.Asin(float64(d.Y))
	yaw := math.Atan2(float64(d.Z), float64(d.X))
	return FlyCamera{
		Eye:         eye,
		Yaw:         yaw,
		Pitch:       pitch,
		Speed:       15.0,
		Sensitivity: 0.003,
		PitchLimit:  1.5,
		firstMouse:  true,
		Fovy:        (60.0 * math.Pi / 180.0),
		Aspect:      float64(screen.X / screen.Y),
	}
}

// Forward returns the unit vector in the direction the camera is looking.
func (c *FlyCamera) Forward() Vec3 {
	return Vec3Normalize(Vec3{
		X: float32(math.Cos(c.Pitch) * math.Cos(c.Yaw)),
		Y: float32(math.Sin(c.Pitch)),
		Z: float32(math.Cos(c.Pitch) * math.Sin(c.Yaw)),
	})
}

// Right returns the unit vector to the camera's right.
func (c *FlyCamera) Right() Vec3 {
	return Vec3Normalize(Vec3Cross(c.Forward(), Vec3{X: 0, Y: 1, Z: 0}))
}

// Center returns the look-at target point (eye + forward).
func (c *FlyCamera) Center() Vec3 {
	return Vec3Add(c.Eye, c.Forward())
}

// BeginDrag starts a mouse-drag rotation session.
func (c *FlyCamera) BeginDrag() {
	c.isDragging = true
	c.firstMouse = true
}

// EndDrag stops the mouse-drag rotation session.
func (c *FlyCamera) EndDrag() {
	c.isDragging = false
}

// IsDragging reports whether a drag rotation is in progress.
func (c *FlyCamera) IsDragging() bool { return c.isDragging }

// MouseMove updates yaw/pitch from a new cursor position.
// Call this from a cursor-position callback whenever IsDragging() is true.
func (c *FlyCamera) MouseMove(xpos, ypos float64) {
	if c.firstMouse {
		c.lastX = xpos
		c.lastY = ypos
		c.firstMouse = false
		return
	}
	dx := xpos - c.lastX
	dy := ypos - c.lastY
	c.lastX = xpos
	c.lastY = ypos

	c.Yaw += dx * c.Sensitivity
	c.Pitch -= dy * c.Sensitivity
	if c.Pitch > c.PitchLimit {
		c.Pitch = c.PitchLimit
	}
	if c.Pitch < -c.PitchLimit {
		c.Pitch = -c.PitchLimit
	}
}

// MoveForward moves the camera forward (positive) or backward (negative) by speed*dt.
func (c *FlyCamera) MoveForward(dt float32) {
	c.Eye = Vec3Add(c.Eye, Vec3Scale(c.Forward(), c.Speed*dt))
}

// MoveBack moves the camera backward by speed*dt.
func (c *FlyCamera) MoveBack(dt float32) {
	c.Eye = Vec3Sub(c.Eye, Vec3Scale(c.Forward(), c.Speed*dt))
}

// MoveLeft strafes the camera left by speed*dt.
func (c *FlyCamera) MoveLeft(dt float32) {
	c.Eye = Vec3Sub(c.Eye, Vec3Scale(c.Right(), c.Speed*dt))
}

// MoveRight strafes the camera right by speed*dt.
func (c *FlyCamera) MoveRight(dt float32) {
	c.Eye = Vec3Add(c.Eye, Vec3Scale(c.Right(), c.Speed*dt))
}
