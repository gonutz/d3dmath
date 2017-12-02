package d3dmath

import "testing"

func TestMat2Add(t *testing.T) {
	m := Mat2{
		1, 2,
		3, 4,
	}.Add(Mat2{
		3, 2,
		5, 6,
	})
	checkFloats(t, m[:],
		4, 4,
		8, 10,
	)
}

func TestMat2Sub(t *testing.T) {
	m := Mat2{
		1, 2,
		3, 4,
	}.Sub(Mat2{
		3, 2,
		1, 6,
	})
	checkFloats(t, m[:],
		-2, 0,
		2, -2,
	)
}

func TestMat2MulMat2(t *testing.T) {
	m := Mat2{
		1, 2,
		3, 4,
	}.Mul(Mat2{
		3, 2,
		1, 6,
	})
	checkFloats(t, m[:],
		5, 14,
		13, 30,
	)
}

func TestIdentity2(t *testing.T) {
	m := Identity2()
	checkFloats(t, m[:],
		1, 0,
		0, 1,
	)
}

func TestMat2Homgeneous(t *testing.T) {
	m := Mat2{
		1, 2,
		3, 4,
	}.Homgeneous()
	checkFloats(t, m[:],
		1, 2, 0,
		3, 4, 0,
		0, 0, 1,
	)
}

func TestMat2String(t *testing.T) {
	m := Mat2{
		1, 2,
		3, 4,
	}
	checkString(t, m.String(), "1.00 2.00\n3.00 4.00")
}
