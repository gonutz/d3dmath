package d3dmath

import (
	"fmt"
	"math"
)

// Vec2 is a 2-element row vector. Elements are called x, y in the docs.
type Vec2 [2]float32

// Negate returns a vector with all elements of v negated.
func (v Vec2) Negate() Vec2 {
	return Vec2{-v[0], -v[1]}
}

// Add returns the sum of v + w.
func (v Vec2) Add(w Vec2) Vec2 {
	return Vec2{v[0] + w[0], v[1] + w[1]}
}

// Sub returns the difference of v - w.
func (v Vec2) Sub(w Vec2) Vec2 {
	return Vec2{v[0] - w[0], v[1] - w[1]}
}

// Dot returns the dot-product of v and w.
func (v Vec2) Dot(w Vec2) float32 {
	return v[0]*w[0] + v[1]*w[1]
}

// MulScalar returns a vector with all elements of v scaled by s.
func (v Vec2) MulScalar(s float32) Vec2 {
	return Vec2{v[0] * s, v[1] * s}
}

// MulMat returns the product of row vector v and matrix m.
func (v Vec2) MulMat(m Mat2) Vec2 {
	return Vec2{
		v[0]*m[0] + v[1]*m[2],
		v[0]*m[1] + v[1]*m[3],
	}
}

// SquareNorm returns the square of the length of v.
func (v Vec2) SquareNorm() float32 {
	return v[0]*v[0] + v[1]*v[1]
}

// Norm returns the length of v.
func (v Vec2) Norm() float32 {
	return float32(math.Hypot(float64(v[0]), float64(v[1])))
}

// Normalized returns a copy of v with elements normalized so the returned
// vector has length 1.
func (v Vec2) Normalized() Vec2 {
	f := 1.0 / v.Norm()
	return Vec2{f * v[0], f * v[1]}
}

// Homogeneous returns a 3-element vector where x and y are the same as in v and
// z is 1.
func (v Vec2) Homogeneous() Vec3 {
	return Vec3{v[0], v[1], 1}
}

func (v Vec2) String() string {
	return fmt.Sprintf("(%.2f %.2f)", v[0], v[1])
}

// AddVec2 returns the sum of all given vectors.
func AddVec2(v0 Vec2, v ...Vec2) Vec2 {
	if len(v) == 0 {
		return v0
	}
	return v0.Add(AddVec2(v[0], v[1:]...))
}
