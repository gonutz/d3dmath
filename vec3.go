package d3dmath

import (
	"fmt"
	"math"
)

// Vec3 is a 3-element row vector. Elements are called x, y, z in the docs.
type Vec3 [3]float32

// Negate returns a vector with all elements of v negated.
func (v Vec3) Negate() Vec3 {
	return Vec3{-v[0], -v[1], -v[2]}
}

// Add returns the sum of v + w.
func (v Vec3) Add(w Vec3) Vec3 {
	return Vec3{v[0] + w[0], v[1] + w[1], v[2] + w[2]}
}

// Sub returns the difference of v - w.
func (v Vec3) Sub(w Vec3) Vec3 {
	return Vec3{v[0] - w[0], v[1] - w[1], v[2] - w[2]}
}

// Dot returns the dot-product of v and w.
func (v Vec3) Dot(w Vec3) float32 {
	return v[0]*w[0] + v[1]*w[1] + v[2]*w[2]
}

// Cross returns the cross-product of v and w.
func (v Vec3) Cross(w Vec3) Vec3 {
	return Vec3{
		v[1]*w[2] - v[2]*w[1],
		v[2]*w[0] - v[0]*w[2],
		v[0]*w[1] - v[1]*w[0],
	}
}

// MulScalar returns a vector with all elements of v scaled by s.
func (v Vec3) MulScalar(s float32) Vec3 {
	return Vec3{v[0] * s, v[1] * s, v[2] * s}
}

// MulMat returns the product of row vector v and matrix m.
func (v Vec3) MulMat(m Mat3) Vec3 {
	return Vec3{
		v[0]*m[0] + v[1]*m[3] + v[2]*m[6],
		v[0]*m[1] + v[1]*m[4] + v[2]*m[7],
		v[0]*m[2] + v[1]*m[5] + v[2]*m[8],
	}
}

// SquareNorm returns the square of the length of v.
func (v Vec3) SquareNorm() float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// Norm returns the length of v.
func (v Vec3) Norm() float32 {
	return float32(math.Sqrt(float64(v.SquareNorm())))
}

// Normalized returns a copy of v with elements normalized so the returned
// vector has length 1.
func (v Vec3) Normalized() Vec3 {
	f := 1.0 / v.Norm()
	return Vec3{f * v[0], f * v[1], f * v[2]}
}

// Homogeneous returns a 4-element vector where x, y and z are the same as in v
// and w is 1.
func (v Vec3) Homogeneous() Vec4 {
	return Vec4{v[0], v[1], v[2], 1}
}

// DropZ returns a 2-element vector where x and y are the same as in v.
// This can be useful when going back from a homogeneous 3-element vector with z
// == 1, down one dimension to a 2-element vector.
// If z != 1 then use ByZ() to divide by z instead.
func (v Vec3) DropZ() Vec2 {
	return Vec2{v[0], v[1]}
}

// ByW returns a 2-element vector where x and y are the same as in v but divided
// by z. This can be useful when going back from a homogeneous 3-element vector
// with z != 1, down one dimension to a 2-element vector.
func (v Vec3) ByZ() Vec2 {
	f := 1.0 / v[2]
	return Vec2{f * v[0], f * v[1]}
}

func (v Vec3) String() string {
	return fmt.Sprintf("(%.2f %.2f %.2f)", v[0], v[1], v[2])
}

// AddVec3 returns the sum of all given vectors.
func AddVec3(v0 Vec3, v ...Vec3) Vec3 {
	if len(v) == 0 {
		return v0
	}
	return v0.Add(AddVec3(v[0], v[1:]...))
}
