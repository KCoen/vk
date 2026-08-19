package utility

import (
	"math"
)

// Vec2 represents a 2D float32 vector.
type Vec2 struct {
	X, Y float32
}

// Vec3 represents a 3D float32 vector.
type Vec3 struct {
	X, Y, Z float32
}

// Vec4 represents a 4D float32 vector.
type Vec4 struct {
	X, Y, Z, W float32
}

// Mat4 represents a 4x4 column-major matrix matching GLSL/Vulkan layout.
type Mat4 [16]float32

// Mat4Identity returns an identity matrix.
func Mat4Identity() Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Mat4Multiply multiplies two 4x4 matrices: a * b.
func Mat4Multiply(a, b Mat4) Mat4 {
	var out Mat4
	for col := 0; col < 4; col++ {
		for row := 0; row < 4; row++ {
			var sum float32
			for k := 0; k < 4; k++ {
				sum += a[k*4+row] * b[col*4+k]
			}
			out[col*4+row] = sum
		}
	}
	return out
}

// Mat4LookAt creates a view matrix looking from 'eye' to 'center' with up vector 'up'.
func Mat4LookAt(eye, center, up Vec3) Mat4 {
	f := Vec3Normalize(Vec3Sub(center, eye))
	s := Vec3Normalize(Vec3Cross(f, up))
	u := Vec3Cross(s, f)

	return Mat4{
		s.X, u.X, -f.X, 0,
		s.Y, u.Y, -f.Y, 0,
		s.Z, u.Z, -f.Z, 0,
		-Vec3Dot(s, eye), -Vec3Dot(u, eye), Vec3Dot(f, eye), 1,
	}
}

// Mat4Perspective creates a perspective projection matrix for Vulkan (depth [0, 1] and flipped Y).
func Mat4Perspective(fovyRad, aspect, zNear, zFar float32) Mat4 {
	tanHalfFovy := float32(math.Tan(float64(fovyRad) / 2.0))

	var m Mat4
	m[0] = 1.0 / (aspect * tanHalfFovy)
	m[5] = -1.0 / tanHalfFovy // Vulkan Y-flip
	m[10] = zFar / (zNear - zFar)
	m[11] = -1.0
	m[14] = (zFar * zNear) / (zNear - zFar)
	return m
}

// Vec3 operations
func Vec3Add(a, b Vec3) Vec3 { return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func Vec3Sub(a, b Vec3) Vec3 { return Vec3{a.X - b.X, a.Y - b.Y, a.Z - b.Z} }
func Vec3Scale(v Vec3, s float32) Vec3 { return Vec3{v.X * s, v.Y * s, v.Z * s} }
func Vec3Dot(a, b Vec3) float32 { return a.X*b.X + a.Y*b.Y + a.Z*b.Z }
func Vec3Cross(a, b Vec3) Vec3 {
	return Vec3{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}
func Vec3Length2(v Vec3) float32 { return Vec3Dot(v, v) }
func Vec3Length(v Vec3) float32  { return float32(math.Sqrt(float64(Vec3Length2(v)))) }
func Vec3Normalize(v Vec3) Vec3 {
	l := Vec3Length(v)
	if l == 0 {
		return Vec3{}
	}
	return Vec3Scale(v, 1.0/l)
}

// BoundingSphere represents a sphere defined by center and radius.
type BoundingSphere struct {
	Center Vec3
	Radius float32
}

// ComputeBoundingSphere computes an enclosing bounding sphere for a list of points.
func ComputeBoundingSphere(pts []Vec3) BoundingSphere {
	if len(pts) == 0 {
		return BoundingSphere{}
	}
	var center Vec3
	for _, p := range pts {
		center = Vec3Add(center, p)
	}
	center = Vec3Scale(center, 1.0/float32(len(pts)))

	var maxDist2 float32
	for _, p := range pts {
		d2 := Vec3Length2(Vec3Sub(p, center))
		if d2 > maxDist2 {
			maxDist2 = d2
		}
	}
	radius := float32(math.Sqrt(float64(maxDist2))) * 1.001
	return BoundingSphere{
		Center: center,
		Radius: radius,
	}
}

// VisibilityTester extracts viewing frustum planes from world-view-projection matrix and tests sphere visibility.
type VisibilityTester struct {
	Planes [6]Vec4
}

// NewVisibilityTester extracts and normalizes the 6 frustum planes from a projection*view matrix.
func NewVisibilityTester(mat Mat4) VisibilityTester {
	var vt VisibilityTester
	planeIdx := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			sign := float32(-1.0)
			if j > 0 {
				sign = 1.0
			}
			var plane Vec4
			plane.X = mat[3] + sign*mat[i]
			plane.Y = mat[7] + sign*mat[4+i]
			plane.Z = mat[11] + sign*mat[8+i]
			plane.W = mat[15] + sign*mat[12+i]

			// Normalize plane
			len3 := float32(math.Sqrt(float64(plane.X*plane.X + plane.Y*plane.Y + plane.Z*plane.Z)))
			if len3 > 0 {
				plane.X /= len3
				plane.Y /= len3
				plane.Z /= len3
				plane.W /= len3
			}
			vt.Planes[planeIdx] = plane
			planeIdx++
		}
	}
	return vt
}

// IsVisible checks if a bounding sphere is within or intersects the frustum.
func (vt *VisibilityTester) IsVisible(origin Vec3, radius float32) bool {
	// Planes to test: Left, Right, Near, Far (indices 0, 1, 4, 5)
	testPlanes := []int{0, 1, 4, 5}
	for _, idx := range testPlanes {
		p := vt.Planes[idx]
		dist := p.X*origin.X + p.Y*origin.Y + p.Z*origin.Z + p.W
		if dist+radius < 0 {
			return false
		}
	}
	return true
}
