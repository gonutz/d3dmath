package d3dmath

import "fmt"

// Mat3 is a 3 by 3 matrix of float32s in row-major order.
type Mat3 [9]float32

// Add returns the sum of m + n.
func (m Mat3) Add(n Mat3) (sum Mat3) {
	for i := range sum {
		sum[i] = m[i] + n[i]
	}
	return
}

// Sub returns the difference of m - n.
func (m Mat3) Sub(n Mat3) (diff Mat3) {
	for i := range diff {
		diff[i] = m[i] - n[i]
	}
	return
}

// Mul returns the product of m * n.
func (m Mat3) Mul(n Mat3) Mat3 {
	return Mat3{
		m[0]*n[0] + m[1]*n[3] + m[2]*n[6],
		m[0]*n[1] + m[1]*n[4] + m[2]*n[7],
		m[0]*n[2] + m[1]*n[5] + m[2]*n[8],

		m[3]*n[0] + m[4]*n[3] + m[5]*n[6],
		m[3]*n[1] + m[4]*n[4] + m[5]*n[7],
		m[3]*n[2] + m[4]*n[5] + m[5]*n[8],

		m[6]*n[0] + m[7]*n[3] + m[8]*n[6],
		m[6]*n[1] + m[7]*n[4] + m[8]*n[7],
		m[6]*n[2] + m[7]*n[5] + m[8]*n[8],
	}
}

// Identity3 returns the 3 by 3 identity matrix.
func Identity3() Mat3 {
	return Mat3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

// Mul3 returns the product of the given matrices.
func Mul3(m0 Mat3, m ...Mat3) Mat3 {
	if len(m) == 0 {
		return m0
	}
	return m0.Mul(Mul3(m[0], m[1:]...))
}

// Transposed returns a transposed copy of m.
func (m Mat3) Transposed() Mat3 {
	return Mat3{
		m[0], m[3], m[6],
		m[1], m[4], m[7],
		m[2], m[5], m[8],
	}
}

// Homogeneous returns the homogeneous 4-dimensional equivalent of the
// 3-dimensional matrix.
func (m Mat3) Homogeneous() Mat4 {
	return Mat4{
		m[0], m[1], m[2], 0,
		m[3], m[4], m[5], 0,
		m[6], m[7], m[8], 0,
		0, 0, 0, 1,
	}
}

func (m Mat3) String() string {
	return fmt.Sprintf(`%.2f %.2f %.2f
%.2f %.2f %.2f
%.2f %.2f %.2f`, m[0], m[1], m[2], m[3], m[4], m[5], m[6], m[7], m[8])
}
