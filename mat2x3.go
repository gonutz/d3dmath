package d3dmath

// Mat2x3 is a 2x3 matrix of float32s in row-major order. It represents a
// homogeneous 3x3 matrix where the last line is 0,0,1 implicitly.
type Mat2x3 [6]float32

func (m Mat2x3) Add(n Mat2x3) (sum Mat2x3) {
	for i := range sum {
		sum[i] = m[i] + n[i]
	}
	return
}

func (m Mat2x3) Sub(n Mat2x3) (diff Mat2x3) {
	for i := range diff {
		diff[i] = m[i] - n[i]
	}
	return
}

//          | n0        n1        n2
//          | n3        n4        n5
//    *     |  0         0         1
// ---------+---------------------------------------------
// m0 m1 m2 | m0n0+m1n3 m0n1+m1n4 m0n2+m1n5+m2
// m3 m4 m5 | m3n0+m4n3 m3n1+m4n4 m3n2+m4n5+m5
//  0  0  1 |  0        0         1
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

func Identity2x3() Mat2x3 {
	return Mat2x3{
		1, 0, 0,
		0, 1, 0,
	}
}

func Mul2x3(m0 Mat2x3, m ...Mat2x3) Mat2x3 {
	if len(m) == 0 {
		return m0
	}
	return m0.Mul(Mul2x3(m[0], m[1:]...))
}

func (m Mat2x3) ToMat3() Mat3 {
	return Mat3{
		m[0], m[1], m[2],
		m[3], m[4], m[5],
		0, 0, 1,
	}
}
