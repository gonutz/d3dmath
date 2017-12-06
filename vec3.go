package d3dmath

import (
	"fmt"
	"math"
)

type Vec3 [3]float32

func (v Vec3) Add(w Vec3) Vec3 {
	return Vec3{v[0] + w[0], v[1] + w[1], v[2] + w[2]}
}

func (v Vec3) Sub(w Vec3) Vec3 {
	return Vec3{v[0] - w[0], v[1] - w[1], v[2] - w[2]}
}

func (v Vec3) Dot(w Vec3) float32 {
	return v[0]*w[0] + v[1]*w[1] + v[2]*w[2]
}

func (v Vec3) Cross(w Vec3) Vec3 {
	return Vec3{
		v[1]*w[2] - v[2]*w[1],
		v[2]*w[0] - v[0]*w[2],
		v[0]*w[1] - v[1]*w[0],
	}
}

func (v Vec3) MulScalar(s float32) Vec3 {
	return Vec3{v[0] * s, v[1] * s, v[2] * s}
}

//          | m0 m1 m2
//          | m3 m4 m5
//    *     | m6 m7 m8
// ---------+---------
// v0 v1 v2 | ...
func (v Vec3) MulMat(m Mat3) Vec3 {
	return Vec3{
		v[0]*m[0] + v[1]*m[3] + v[2]*m[6],
		v[0]*m[1] + v[1]*m[4] + v[2]*m[7],
		v[0]*m[2] + v[1]*m[5] + v[2]*m[8],
	}
}

func (v Vec3) SquareNorm() float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

func (v Vec3) Norm() float32 {
	return float32(math.Sqrt(float64(v.SquareNorm())))
}

func (v Vec3) Normalized() Vec3 {
	f := 1.0 / v.Norm()
	return Vec3{f * v[0], f * v[1], f * v[2]}
}

func (v Vec3) Homgeneous() Vec4 {
	return Vec4{v[0], v[1], v[2], 1}
}

func (v Vec3) String() string {
	return fmt.Sprintf("(%.2f %.2f %.2f)", v[0], v[1], v[2])
}

func AddVec3(v0 Vec3, v ...Vec3) Vec3 {
	if len(v) == 0 {
		return v0
	}
	return v0.Add(AddVec3(v[0], v[1:]...))
}