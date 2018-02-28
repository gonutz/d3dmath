package d3dmath

import "testing"

func TestMat3Add(t *testing.T) {
	m := Mat3{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}.Add(Mat3{
		3, 2, 3,
		5, 6, 4,
		17, 6, 5,
	})
	checkFloats(t, m[:],
		4, 4, 6,
		9, 11, 10,
		24, 14, 14,
	)
}

func TestMat3Sub(t *testing.T) {
	m := Mat3{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}.Sub(Mat3{
		3, 2, 3,
		5, 6, 4,
		17, 6, 5,
	})
	checkFloats(t, m[:],
		-2, 0, 0,
		-1, -1, 2,
		-10, 2, 4,
	)
}

func TestMat3MulMat3(t *testing.T) {
	m := Mat3{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}.Mul(Mat3{
		2, 4, 3,
		5, 6, 7,
		4, 2, 3,
	})
	checkFloats(t, m[:],
		24, 22, 26,
		57, 58, 65,
		90, 94, 104,
	)
}

func TestIdentity3(t *testing.T) {
	m := Identity3()
	checkFloats(t, m[:],
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	)
}

func TestMat3Homgeneous(t *testing.T) {
	m := Mat3{
		1, 2, 5,
		3, 4, 6,
		5, 3, 7,
	}.Homogeneous()
	checkFloats(t, m[:],
		1, 2, 5, 0,
		3, 4, 6, 0,
		5, 3, 7, 0,
		0, 0, 0, 1,
	)
}

func TestMat3String(t *testing.T) {
	m := Mat3{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	checkString(t, m.String(), "1.00 2.00 3.00\n4.00 5.00 6.00\n7.00 8.00 9.00")
}
