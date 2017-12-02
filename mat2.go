package d3dmath

import "fmt"

// Mat2 is a 2x2 matrix of float32s in row-major order.
type Mat2 [4]float32

func (m Mat2) Add(n Mat2) Mat2 {
	return Mat2{
		m[0] + n[0], m[1] + n[1],
		m[2] + n[2], m[3] + n[3],
	}
}

func (m Mat2) Sub(n Mat2) Mat2 {
	return Mat2{
		m[0] - n[0], m[1] - n[1],
		m[2] - n[2], m[3] - n[3],
	}
}

//       | n0        n1
//    *  | n2        n3
// ------+--------------------
// m0 m1 | m0n0+m1n2 m0n1+m1n3
// m2 m3 | m2n0+m3n2 m2n1+m3n3
func (m Mat2) Mul(n Mat2) Mat2 {
	return Mat2{
		m[0]*n[0] + m[1]*n[2],
		m[0]*n[1] + m[1]*n[3],

		m[2]*n[0] + m[3]*n[2],
		m[2]*n[1] + m[3]*n[3],
	}
}

func Identity2() Mat2 {
	return Mat2{
		1, 0,
		0, 1,
	}
}

func Mul2(m0 Mat2, m ...Mat2) Mat2 {
	if len(m) == 0 {
		return m0
	}
	return m0.Mul(Mul2(m[0], m[1:]...))
}

func (m Mat2) Homgeneous() Mat3 {
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
