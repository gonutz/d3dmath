package d3dmath

import (
	"fmt"
	"math"
)

// Mat4 is a 4x4 matrix in row-major order of floats
type Mat4 [16]float32

func (m Mat4) Add(n Mat4) (sum Mat4) {
	for i := range sum {
		sum[i] = m[i] + n[i]
	}
	return
}

func (m Mat4) Sub(n Mat4) (diff Mat4) {
	for i := range diff {
		diff[i] = m[i] - n[i]
	}
	return
}

//                 | n0  n1  n2  n3
//                 | n4  n5  n6  n7
//             *   | n8  n9  n10 n11
//                 | n12 n13 n14 n15
// ----------------+---------------------------------------------
// m0  m1  m2  m3  |
// m4  m5  m6  m7  | ...
// m8  m9  m10 m11 |
// m12 m13 m14 m15 |
func (m Mat4) Mul(n Mat4) Mat4 {
	return Mat4{
		m[0]*n[0] + m[1]*n[4] + m[2]*n[8] + m[3]*n[12],
		m[0]*n[1] + m[1]*n[5] + m[2]*n[9] + m[3]*n[13],
		m[0]*n[2] + m[1]*n[6] + m[2]*n[10] + m[3]*n[14],
		m[0]*n[3] + m[1]*n[7] + m[2]*n[11] + m[3]*n[15],

		m[4]*n[0] + m[5]*n[4] + m[6]*n[8] + m[7]*n[12],
		m[4]*n[1] + m[5]*n[5] + m[6]*n[9] + m[7]*n[13],
		m[4]*n[2] + m[5]*n[6] + m[6]*n[10] + m[7]*n[14],
		m[4]*n[3] + m[5]*n[7] + m[6]*n[11] + m[7]*n[15],

		m[8]*n[0] + m[9]*n[4] + m[10]*n[8] + m[11]*n[12],
		m[8]*n[1] + m[9]*n[5] + m[10]*n[9] + m[11]*n[13],
		m[8]*n[2] + m[9]*n[6] + m[10]*n[10] + m[11]*n[14],
		m[8]*n[3] + m[9]*n[7] + m[10]*n[11] + m[11]*n[15],

		m[12]*n[0] + m[13]*n[4] + m[14]*n[8] + m[15]*n[12],
		m[12]*n[1] + m[13]*n[5] + m[14]*n[9] + m[15]*n[13],
		m[12]*n[2] + m[13]*n[6] + m[14]*n[10] + m[15]*n[14],
		m[12]*n[3] + m[13]*n[7] + m[14]*n[11] + m[15]*n[15],
	}
}

func Mul4(m0 Mat4, m ...Mat4) Mat4 {
	if len(m) == 0 {
		return m0
	}
	return m0.Mul(Mul4(m[0], m[1:]...))
}

func (m Mat4) Transposed() Mat4 {
	return Mat4{
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15],
	}
}

func Identity4() Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func Translate(dx, dy, dz float32) Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		dx, dy, dz, 1,
	}
}

func TranslateV(v Vec3) Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		v[0], v[1], v[2], 1,
	}
}

func Scale(dx, dy, dz float32) Mat4 {
	return Mat4{
		dx, 0, 0, 0,
		0, dy, 0, 0,
		0, 0, dz, 0,
		0, 0, 0, 1,
	}
}

func ScaleV(v Vec3) Mat4 {
	return Mat4{
		v[0], 0, 0, 0,
		0, v[1], 0, 0,
		0, 0, v[2], 0,
		0, 0, 0, 1,
	}
}

func RotateX(radians float32) Mat4 {
	s, c := math.Sincos(float64(radians))
	sin, cos := float32(s), float32(c)
	return Mat4{
		1, 0, 0, 0,
		0, cos, sin, 0,
		0, -sin, cos, 0,
		0, 0, 0, 1,
	}
}

func RotateY(radians float32) Mat4 {
	s, c := math.Sincos(float64(radians))
	sin, cos := float32(s), float32(c)
	return Mat4{
		cos, 0, -sin, 0,
		0, 1, 0, 0,
		sin, 0, cos, 0,
		0, 0, 0, 1,
	}
}

func RotateZ(radians float32) Mat4 {
	s, c := math.Sincos(float64(radians))
	sin, cos := float32(s), float32(c)
	return Mat4{
		cos, sin, 0, 0,
		-sin, cos, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func RotateAbout(v Vec3, radians float32) Mat4 {
	sqLen := v.SquareNorm()
	if sqLen < 0.99999 || sqLen > 1.00001 {
		v = v.Normalized()
	}
	s, c := math.Sincos(float64(radians))
	sin, cos := float32(s), float32(c)
	return Mat4{
		cos + v[0]*v[0]*(1-cos), v[0]*v[1]*(1-cos) + v[2]*sin, v[0]*v[2]*(1-cos) - v[1]*sin, 0,
		v[1]*v[0]*(1-cos) - v[2]*sin, cos + v[1]*v[1]*(1-cos), v[1]*v[2]*(1-cos) + v[0]*sin, 0,
		v[2]*v[0]*(1-cos) + v[1]*sin, v[2]*v[1]*(1-cos) - v[0]*sin, cos + v[2]*v[2]*(1-cos), 0,
		0, 0, 0, 1,
	}
}

func Ortho(left, right, bottom, top, near, far float32) Mat4 {
	return Mat4{
		2 / (right - left), 0, 0, (right + left) / (left - right),
		0, 2 / (top - bottom), 0, (top + bottom) / (bottom - top),
		0, 0, 2 / (far - near), (far + near) / (near - far),
		0, 0, 0, 1,
	}
}

func Perspective(fovRadians, aspect, near, far float32) Mat4 {
	f := 1 / float32(math.Tan(float64(fovRadians)/2))
	dz := far - near
	return Mat4{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, far / dz, 1,
		0, 0, -near * far / dz, 0,
	}
}

func LookAt(pos, target, up Vec3) Mat4 {
	z := target.Sub(pos).Normalized()
	x := up.Cross(z).Normalized()
	y := z.Cross(x)
	return Mat4{
		x[0], y[0], z[0], 0,
		x[1], y[1], z[1], 0,
		x[2], y[2], z[2], 0,
		-x.Dot(pos), -y.Dot(pos), -z.Dot(pos), 1,
	}
}

func (m Mat4) String() string {
	return fmt.Sprintf(`%.2f %.2f %.2f %.2f
%.2f %.2f %.2f %.2f
%.2f %.2f %.2f %.2f
%.2f %.2f %.2f %.2f`, m[0], m[1], m[2], m[3], m[4], m[5], m[6], m[7], m[8],
		m[9], m[10], m[11], m[12], m[13], m[14], m[15])
}
