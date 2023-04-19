package d3dmath

import (
	"fmt"
	"math"
)

// Mat4 is a 4 by 4 matrix of float32s in row-major order.
type Mat4 [16]float32

// Add returns the sum of m + n.
func (m Mat4) Add(n Mat4) (sum Mat4) {
	for i := range sum {
		sum[i] = m[i] + n[i]
	}
	return
}

// Sub returns the difference of m - n.
func (m Mat4) Sub(n Mat4) (diff Mat4) {
	for i := range diff {
		diff[i] = m[i] - n[i]
	}
	return
}

// Mul returns the product of m * n.
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

// Mul4 returns the product of the given matrices.
func Mul4(m0 Mat4, m ...Mat4) Mat4 {
	if len(m) == 0 {
		return m0
	}
	return m0.Mul(Mul4(m[0], m[1:]...))
}

// Transposed returns a transposed copy of m.
func (m Mat4) Transposed() Mat4 {
	return Mat4{
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15],
	}
}

// Identity4 returns the 4 by 4 identity matrix.
func Identity4() Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Translate reutrns 4 by 4 matrix that, when multiplied with a homogeneous
// 4-element 3D vector, moves the vector by the given amounts in x, y and z.
func Translate(dx, dy, dz float32) Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		dx, dy, dz, 1,
	}
}

// TranslateV is the same as Translate, but it takes a Vec3 as its argument
// instead of single x, y, z parameters.
func TranslateV(v Vec3) Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		v[0], v[1], v[2], 1,
	}
}

// Scale reutrns 4 by 4 matrix that, when multiplied with a homogeneous
// 4-element 3D vector, scales the vector by the given factors in x, y and z.
func Scale(dx, dy, dz float32) Mat4 {
	return Mat4{
		dx, 0, 0, 0,
		0, dy, 0, 0,
		0, 0, dz, 0,
		0, 0, 0, 1,
	}
}

// ScaleV is the same as Scale, but it takes a Vec3 as its argument instead of
// single x, y, z parameters.
func ScaleV(v Vec3) Mat4 {
	return Mat4{
		v[0], 0, 0, 0,
		0, v[1], 0, 0,
		0, 0, v[2], 0,
		0, 0, 0, 1,
	}
}

// RotateX reutrns 4 by 4 matrix that, when multiplied with a homogeneous
// 4-element 3D vector, rotates the vector about the x-axis by the given angle
// in radians.
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

// RotateY reutrns 4 by 4 matrix that, when multiplied with a homogeneous
// 4-element 3D vector, rotates the vector about the y-axis by the given angle
// in radians.
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

// RotateZ reutrns 4 by 4 matrix that, when multiplied with a homogeneous
// 4-element 3D vector, rotates the vector about the z-axis by the given angle
// in radians.
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

// RotateAbout reutrns 4 by 4 matrix that, when multiplied with a homogeneous
// 4-element 3D vector, rotates the vector about the given vector v by the given
// angle in radians.
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

// Ortho returns an orthographic projection matrix.
func Ortho(left, right, bottom, top, near, far float32) Mat4 {
	return Mat4{
		2 / (right - left), 0, 0, 0,
		0, 2 / (top - bottom), 0, 0,
		0, 0, 2 / (far - near), 0,
		(right + left) / (left - right), (top + bottom) / (bottom - top), (far + near) / (near - far), 1,
	}
}

// Perspective returns an perspective projection matrix.
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

// LookAt returns a matrix that, when used for the camera, looks at target from
// position pos. Since you can tilt your head in infinite ways looking from one
// point at another, the up vector is used to specify which direction is up.
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

// DecomposeAffineTransform decomposes the given matrix into scale, rotation and
// translation matrices that, when multiplied in that order, produce the
// original matrix. See this forum post for reference:
// https://math.stackexchange.com/questions/237369/given-this-transformation-matrix-how-do-i-decompose-it-into-translation-rotati
func DecomposeAffineTransform(m Mat4) (scale, rotation, translation Mat4) {
	translation = Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		m[12], m[13], m[14], 1,
	}
	sx := Vec3{m[0], m[1], m[2]}.Norm()
	sy := Vec3{m[4], m[5], m[6]}.Norm()
	sz := Vec3{m[8], m[9], m[10]}.Norm()
	scale = Mat4{
		sx, 0, 0, 0,
		0, sy, 0, 0,
		0, 0, sz, 0,
		0, 0, 0, 1,
	}
	fx, fy, fz := 1/sx, 1/sy, 1/sz
	rotation = Mat4{
		fx * m[0], fx * m[1], fx * m[2], 0,
		fy * m[4], fy * m[5], fy * m[6], 0,
		fz * m[8], fz * m[9], fz * m[10], 0,
		0, 0, 0, 1,
	}
	return
}
