package d3dmath

import "fmt"

// Mat3 is a 3x3 matrix of float32s in row-major order.
type Mat3 [9]float32

func (m Mat3) Add(n Mat3) (sum Mat3) {
	for i := range sum {
		sum[i] = m[i] + n[i]
	}
	return
}

func (m Mat3) Sub(n Mat3) (diff Mat3) {
	for i := range diff {
		diff[i] = m[i] - n[i]
	}
	return
}

//          | n0             n1             n2
//          | n3             n4             n5
//    *     | n6             n7             n8
// ---------+---------------------------------------------
// m0 m1 m2 | m0n0+m1n3+m2n6 m0n1+m1n4+m2n7 m0n2+m1n5+m2n8
// m3 m4 m5 | m3n0+m4n3+m5n6 m3n1+m4n4+m5n7 m3n2+m4n5+m5n8
// m6 m7 m8 | m6n0+m7n3+m8n6 m6n1+m7n4+m8n7 m6n2+m7n5+m8n8
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

func Identity3() Mat3 {
	return Mat3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

func Mul3(m0 Mat3, m ...Mat3) Mat3 {
	if len(m) == 0 {
		return m0
	}
	return m0.Mul(Mul3(m[0], m[1:]...))
}

func (m Mat3) Homgeneous() Mat4 {
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
