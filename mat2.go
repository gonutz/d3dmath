package d3dmath

import "fmt"

// Mat2 is a 2 by 2 matrix of float32s in row-major order.
type Mat2 [4]float32

// Add returns the sum of m + n.
func (m Mat2) Add(n Mat2) Mat2 {
	return Mat2{
		m[0] + n[0], m[1] + n[1],
		m[2] + n[2], m[3] + n[3],
	}
}

// Sub returns the difference of m - n.
func (m Mat2) Sub(n Mat2) Mat2 {
	return Mat2{
		m[0] - n[0], m[1] - n[1],
		m[2] - n[2], m[3] - n[3],
	}
}

// Mul returns the product of m * n.
func (m Mat2) Mul(n Mat2) Mat2 {
	return Mat2{
		m[0]*n[0] + m[1]*n[2],
		m[0]*n[1] + m[1]*n[3],

		m[2]*n[0] + m[3]*n[2],
		m[2]*n[1] + m[3]*n[3],
	}
}

// Identity2 returns the 2 by 2 identity matrix.
func Identity2() Mat2 {
	return Mat2{
		1, 0,
		0, 1,
	}
}

// Mul2 returns the product of the given matrices.
func Mul2(m0 Mat2, m ...Mat2) Mat2 {
	if len(m) == 0 {
		return m0
	}
	return m0.Mul(Mul2(m[0], m[1:]...))
}

// Transposed returns a transposed copy of m.
func (m Mat2) Transposed() Mat2 {
	return Mat2{
		m[0], m[2],
		m[1], m[3],
	}
}

// Homogeneous returns the homogeneous 3-dimensional equivalent of the
// 2-dimensional matrix.
func (m Mat2) Homogeneous() Mat3 {
	return Mat3{
		m[0], m[1], 0,
		m[2], m[3], 0,
		0, 0, 1,
	}
}

func (m Mat2) String() string {
	return fmt.Sprintf(`%.2f %.2f
%.2f %.2f`, m[0], m[1], m[2], m[3])
}
