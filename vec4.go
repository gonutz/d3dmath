package d3dmath

import (
	"fmt"
	"math"
)

// Vec4 is a 4-element row vector. Elements are called x, y, z, w in the docs.
type Vec4 [4]float32

// Negate returns a vector with all elements of v negated.
func (v Vec4) Negate() Vec4 {
	return Vec4{-v[0], -v[1], -v[2], -v[3]}
}

// Add returns the sum of v + w.
func (v Vec4) Add(w Vec4) Vec4 {
	return Vec4{v[0] + w[0], v[1] + w[1], v[2] + w[2], v[3] + w[3]}
}

// Sub returns the difference of v - w.
func (v Vec4) Sub(w Vec4) Vec4 {
	return Vec4{v[0] - w[0], v[1] - w[1], v[2] - w[2], v[3] - w[3]}
}

// Dot returns the dot-product of v and w.
func (v Vec4) Dot(w Vec4) float32 {
	return v[0]*w[0] + v[1]*w[1] + v[2]*w[2] + v[3]*w[3]
}

// MulScalar returns a vector with all elements of v scaled by s.
func (v Vec4) MulScalar(s float32) Vec4 {
	return Vec4{v[0] * s, v[1] * s, v[2] * s, v[3] * s}
}

// MulMat returns the product of row vector v and matrix m.
func (v Vec4) MulMat(m Mat4) Vec4 {
	return Vec4{
		v[0]*m[0] + v[1]*m[4] + v[2]*m[8] + v[3]*m[12],
		v[0]*m[1] + v[1]*m[5] + v[2]*m[9] + v[3]*m[13],
		v[0]*m[2] + v[1]*m[6] + v[2]*m[10] + v[3]*m[14],
		v[0]*m[3] + v[1]*m[7] + v[2]*m[11] + v[3]*m[15],
	}
}

// SquareNorm returns the square of the length of v.
func (v Vec4) SquareNorm() float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3]
}

// Norm returns the length of v.
func (v Vec4) Norm() float32 {
	return float32(math.Sqrt(float64(v.SquareNorm())))
}

// Normalized returns a copy of v with elements normalized so the returned
// vector has length 1.
func (v Vec4) Normalized() Vec4 {
	f := 1.0 / v.Norm()
	return Vec4{f * v[0], f * v[1], f * v[2], f * v[3]}
}

// DropW returns a 3-element vector where x, y and z are the same as in v.
// This can be useful when going back from a homogeneous 4-element vector with w
// == 1, down one dimension to a 3-element vector.
// If w != 1 then use ByW() to divide by w instead.
func (v Vec4) DropW() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

// ByW returns a 3-element vector where x, y and z are the same as in v but
// divided by w. This can be useful when going back from a homogeneous 4-element
// vector with w != 1, down one dimension to a 3-element vector.
func (v Vec4) ByW() Vec3 {
	f := 1.0 / v[3]
	return Vec3{f * v[0], f * v[1], f * v[2]}
}

func (v Vec4) String() string {
	return fmt.Sprintf("(%.2f %.2f %.2f %.2f)", v[0], v[1], v[2], v[3])
}

// AddVec4 returns the sum of all given vectors.
func AddVec4(v0 Vec4, v ...Vec4) Vec4 {
	if len(v) == 0 {
		return v0
	}
	return v0.Add(AddVec4(v[0], v[1:]...))
}
