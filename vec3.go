package d3dmath

import "math"

type Vec3 [3]float32

func (v Vec3) Add(w Vec3) (sum Vec3) {
	for i := range sum {
		sum[i] = v[i] + w[i]
	}
	return
}

func (v Vec3) Sub(w Vec3) (diff Vec3) {
	for i := range diff {
		diff[i] = v[i] + w[i]
	}
	return
}

func (v Vec3) Dot(w Vec3) float32 {
	return v[0]*w[0] + v[1]*w[1] + v[2]*w[2]
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

func (v Vec3) Homgeneous() Vec4 {
	return Vec4{v[0], v[1], v[2], 1}
}
