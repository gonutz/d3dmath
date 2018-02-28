package d3dmath

import "testing"

func TestVec2Negate(t *testing.T) {
	v := Vec2{2, -3}.Negate()
	checkFloats(t, v[:], -2, 3)
}

func TestVec2Add(t *testing.T) {
	v := Vec2{2, 3}.Add(Vec2{5, 7})
	checkFloats(t, v[:], 7, 10)
}

func TestVec2Sub(t *testing.T) {
	v := Vec2{2, 3}.Sub(Vec2{5, 1})
	checkFloats(t, v[:], -3, 2)
}

func TestVec2Dot(t *testing.T) {
	v := Vec2{2, 3}.Dot(Vec2{5, 1})
	checkFloat(t, v, 13)
}

func TestVec2MulScalar(t *testing.T) {
	v := Vec2{2, 3}.MulScalar(2)
	checkFloats(t, v[:], 4, 6)
}

func TestVec2MulMat2(t *testing.T) {
	v := Vec2{2, 3}.MulMat(Mat2{
		2, 4,
		5, 3,
	})
	checkFloats(t, v[:], 19, 17)
}

func TestVec2SquareNorm(t *testing.T) {
	v := Vec2{2, 3}.SquareNorm()
	checkFloat(t, v, 13)
}

func TestVec2Norm(t *testing.T) {
	v := Vec2{3, 4}.Norm()
	checkFloat(t, v, 5)
}

func TestVec2Normalized(t *testing.T) {
	v := Vec2{3, 4}.Normalized()
	checkFloats(t, v[:], 3.0/5, 4.0/5)
}

func TestVec2Homogeneous(t *testing.T) {
	v := Vec2{3, 4}.Homgeneous()
	checkFloats(t, v[:], 3, 4, 1)
}

func TestVec2String(t *testing.T) {
	v := Vec2{3, 4}
	checkString(t, v.String(), "(3.00 4.00)")
}
