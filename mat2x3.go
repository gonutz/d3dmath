package d3dmath

import "fmt"

// Mat2x3 is a 2x3 matrix of float32s in row-major order. It represents a
// homogeneous 3x3 matrix where the last line is 0,0,1 implicitly.
type Mat2x3 [6]float32

// Add returns the sum of m + n.
func (m Mat2x3) Add(n Mat2x3) (sum Mat2x3) {
	for i := range sum {
		sum[i] = m[i] + n[i]
	}
	return
}

// Sub returns the difference of m - n.
func (m Mat2x3) Sub(n Mat2x3) (diff Mat2x3) {
	for i := range diff {
		diff[i] = m[i] - n[i]
	}
	return
}

// Mul returns the product of m * n.
func (m Mat2x3) Mul(n Mat2x3) Mat2x3 {
	return Mat2x3{
		m[0]*n[0] + m[1]*n[3],
		m[0]*n[1] + m[1]*n[4],
		m[0]*n[2] + m[1]*n[5] + m[2],

		m[3]*n[0] + m[4]*n[3],
		m[3]*n[1] + m[4]*n[4],
		m[3]*n[2] + m[4]*n[5] + m[5],
	}
}

// Identity2x3 returns the 2 by 3 homogeneous identity matrix.
func Identity2x3() Mat2x3 {
	return Mat2x3{
		1, 0, 0,
		0, 1, 0,
	}
}

// Mul2x3 returns the product of the given matrices.
func Mul2x3(m0 Mat2x3, m ...Mat2x3) Mat2x3 {
	if len(m) == 0 {
		return m0
	}
	return m0.Mul(Mul2x3(m[0], m[1:]...))
}

// ToMat3 returns the 3 by 3 representation of m.
func (m Mat2x3) ToMat3() Mat3 {
	return Mat3{
		m[0], m[1], m[2],
		m[3], m[4], m[5],
		0, 0, 1,
	}
}

func (m Mat2x3) String() string {
	return fmt.Sprintf(`%.2f %.2f %.2f
%.2f %.2f %.2f`, m[0], m[1], m[2], m[3], m[4], m[5])
}
