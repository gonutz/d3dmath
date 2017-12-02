package d3dmath

import "testing"

func TestVec4Add(t *testing.T) {
	v := Vec4{2, 3, -5, 9}.Add(Vec4{5, 7, 1, -1})
	checkFloats(t, v[:], 7, 10, -4, 8)
}

func TestVec4Sub(t *testing.T) {
	v := Vec4{2, 3, -1, 10}.Sub(Vec4{5, 1, 4, 5})
	checkFloats(t, v[:], -3, 2, -5, 5)
}

func TestVec4Dot(t *testing.T) {
	v := Vec4{2, 3, 4, 3}.Dot(Vec4{5, 1, 2, 2})
	checkFloat(t, v, 27)
}

func TestVec4MulScalar(t *testing.T) {
	v := Vec4{2, 3, 4, 5}.MulScalar(2)
	checkFloats(t, v[:], 4, 6, 8, 10)
}

func TestVec4MulMat4(t *testing.T) {
	v := Vec4{2, 3, 4, 5}.MulMat(Mat4{
		2, 4, 3, 2,
		5, 3, 1, 3,
		7, 8, 2, 6,
		3, 2, 5, 3,
	})
	checkFloats(t, v[:], 4+15+28+15, 8+9+32+10, 6+3+8+25, 4+9+24+15)
}

func TestVec4String(t *testing.T) {
	v := Vec4{3, 4, 5, -1}
	checkString(t, v.String(), "(3.00 4.00 5.00 -1.00)")
}
