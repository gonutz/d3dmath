package d3dmath

import (
	"math"
	"testing"
)

func TestVec3Add(t *testing.T) {
	v := Vec3{2, 3, -5}.Add(Vec3{5, 7, 1})
	checkFloats(t, v[:], 7, 10, -4)
}

func TestVec3Sub(t *testing.T) {
	v := Vec3{2, 3, -1}.Sub(Vec3{5, 1, 4})
	checkFloats(t, v[:], -3, 2, -5)
}

func TestVec3Dot(t *testing.T) {
	v := Vec3{2, 3, 4}.Dot(Vec3{5, 1, 2})
	checkFloat(t, v, 21)
}

func TestVec3Cross(t *testing.T) {
	v := Vec3{2, 3, 4}.Cross(Vec3{5, 1, 7})
	checkFloats(t, v[:], 17, 6, -13)
}

func TestVec3MulScalar(t *testing.T) {
	v := Vec3{2, 3, 4}.MulScalar(2)
	checkFloats(t, v[:], 4, 6, 8)
}

func TestVec3MulMat3(t *testing.T) {
	v := Vec3{2, 3, 4}.MulMat(Mat3{
		2, 4, 3,
		5, 3, 1,
		7, 8, 2,
	})
	checkFloats(t, v[:], 4+15+28, 8+9+32, 6+3+8)
}

func TestVec3SquareNorm(t *testing.T) {
	v := Vec3{2, 3, 4}.SquareNorm()
	checkFloat(t, v, 29)
}

func TestVec3Norm(t *testing.T) {
	v := Vec3{2, 3, 4}.Norm()
	checkFloat(t, v, float32(math.Sqrt(29)))
}

func TestVec3Normalized(t *testing.T) {
	v := Vec3{2, 3, 4}.Normalized()
	f := 1.0 / float32(math.Sqrt(29))
	checkFloats(t, v[:], 2*f, 3*f, 4*f)
}

func TestVec3Homogeneous(t *testing.T) {
	v := Vec3{3, 4, 5}.Homgeneous()
	checkFloats(t, v[:], 3, 4, 5, 1)
}

func TestVec3String(t *testing.T) {
	v := Vec3{3, 4, 5}
	checkString(t, v.String(), "(3.00 4.00 5.00)")
}
