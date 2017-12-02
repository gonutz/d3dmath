package d3dmath

import "fmt"

type Vec4 [4]float32

func (v Vec4) Add(w Vec4) Vec4 {
	return Vec4{v[0] + w[0], v[1] + w[1], v[2] + w[2], v[3] + w[3]}
}

func (v Vec4) Sub(w Vec4) Vec4 {
	return Vec4{v[0] - w[0], v[1] - w[1], v[2] - w[2], v[3] - w[3]}
}

func (v Vec4) Dot(w Vec4) float32 {
	return v[0]*w[0] + v[1]*w[1] + v[2]*w[2] + v[3]*w[3]
}

func (v Vec4) MulScalar(s float32) Vec4 {
	return Vec4{v[0] * s, v[1] * s, v[2] * s, v[3] * s}
}

//             | m0  m1  m2  m3
//             | m4  m5  m6  m7
//         *   | m8  m9  m10 m11
//             | m12 m13 m14 m15
// ------------+----------------
// v0 v1 v2 v3 | ...
func (v Vec4) MulMat(m Mat4) Vec4 {
	return Vec4{
		v[0]*m[0] + v[1]*m[4] + v[2]*m[8] + v[3]*m[12],
		v[0]*m[1] + v[1]*m[5] + v[2]*m[9] + v[3]*m[13],
		v[0]*m[2] + v[1]*m[6] + v[2]*m[10] + v[3]*m[14],
		v[0]*m[3] + v[1]*m[7] + v[2]*m[11] + v[3]*m[15],
	}
}

func (v Vec4) DropW() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

func (v Vec4) String() string {
	return fmt.Sprintf("(%.2f %.2f %.2f %.2f)", v[0], v[1], v[2], v[3])
}
