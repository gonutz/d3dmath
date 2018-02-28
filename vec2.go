package d3dmath

import (
	"fmt"
	"math"
)

type Vec2 [2]float32

func (v Vec2) Negate() Vec2 {
	return Vec2{-v[0], -v[1]}
}

func (v Vec2) Add(w Vec2) Vec2 {
	return Vec2{v[0] + w[0], v[1] + w[1]}
}

func (v Vec2) Sub(w Vec2) Vec2 {
	return Vec2{v[0] - w[0], v[1] - w[1]}
}

func (v Vec2) Dot(w Vec2) float32 {
	return v[0]*w[0] + v[1]*w[1]
}

func (v Vec2) MulScalar(s float32) Vec2 {
	return Vec2{v[0] * s, v[1] * s}
}

//       | m0        m1
//    *  | m2        m3
// ------+--------------------
// v0 v1 | v0m0+v1m2 v0m1+v1m3
func (v Vec2) MulMat(m Mat2) Vec2 {
	return Vec2{
		v[0]*m[0] + v[1]*m[2],
		v[0]*m[1] + v[1]*m[3],
	}
}

func (v Vec2) SquareNorm() float32 {
	return v[0]*v[0] + v[1]*v[1]
}

func (v Vec2) Norm() float32 {
	return float32(math.Hypot(float64(v[0]), float64(v[1])))
}

func (v Vec2) Normalized() Vec2 {
	f := 1.0 / v.Norm()
	return Vec2{f * v[0], f * v[1]}
}

func (v Vec2) Homgeneous() Vec3 {
	return Vec3{v[0], v[1], 1}
}

func (v Vec2) String() string {
	return fmt.Sprintf("(%.2f %.2f)", v[0], v[1])
}
