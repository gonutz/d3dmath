package d3dmath

import (
	"math"
	"testing"
)

func TestRotationConversionFactors(t *testing.T) {
	checkFloat(t, 0.5*TurnsToRad, math.Pi)
	checkFloat(t, math.Pi*RadToTurns, 0.5)
	checkFloat(t, math.Pi*RadToDeg, 180)
	checkFloat(t, 180*DegToRad, math.Pi)
	checkFloat(t, 0.5*TurnsToDeg, 180)
	checkFloat(t, 180*DegToTurns, 0.5)
}

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

func TestVec2MulMat(t *testing.T) {
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
	v := Vec2{3, 4}.Homogeneous()
	checkFloats(t, v[:], 3, 4, 1)
}

func TestVec2String(t *testing.T) {
	v := Vec2{3, 4}
	checkString(t, v.String(), "(3.00 4.00)")
}

func TestAddVec2(t *testing.T) {
	a := Vec2{1, 2}
	b := Vec2{3, 2}
	c := Vec2{4, 7}
	sum := AddVec2(a, b, c)
	checkFloats(t, sum[:], 8, 11)
}

func TestVec3Negate(t *testing.T) {
	v := Vec3{2, 3, -5}.Negate()
	checkFloats(t, v[:], -2, -3, 5)
}

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
	v := Vec3{3, 4, 5}.Homogeneous()
	checkFloats(t, v[:], 3, 4, 5, 1)
}

func TestVec3DropZ(t *testing.T) {
	v := Vec3{1, 2, 3}.DropZ()
	checkFloats(t, v[:], 1, 2)
}

func TestVec3ByZ(t *testing.T) {
	v := Vec3{1, 2, 3}.ByZ()
	checkFloats(t, v[:], 1.0/3, 2.0/3)
}

func TestVec3String(t *testing.T) {
	v := Vec3{3, 4, 5}
	checkString(t, v.String(), "(3.00 4.00 5.00)")
}

func TestAddVec3(t *testing.T) {
	a := Vec3{1, 2, 3}
	b := Vec3{-3, 2, 5}
	c := Vec3{9, 6, 7}
	sum := AddVec3(a, b, c)
	checkFloats(t, sum[:], 7, 10, 15)
}

func TestVec4Negate(t *testing.T) {
	v := Vec4{-2, 3, -5, 9}.Negate()
	checkFloats(t, v[:], 2, -3, 5, -9)
}

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

func TestVec4SquareNorm(t *testing.T) {
	v := Vec4{2, 3, 4, 5}.SquareNorm()
	checkFloat(t, v, 54)
}

func TestVec4Norm(t *testing.T) {
	v := Vec4{2, 3, 4, 5}.Norm()
	checkFloat(t, v, float32(math.Sqrt(54)))
}

func TestVec4Normalized(t *testing.T) {
	v := Vec4{2, 3, 4, 5}.Normalized()
	f := 1.0 / float32(math.Sqrt(54))
	checkFloats(t, v[:], 2*f, 3*f, 4*f, 5*f)
}

func TestVec4DropW(t *testing.T) {
	v := Vec4{2, 3, 4, 5}.DropW()
	checkFloats(t, v[:], 2, 3, 4)
}

func TestVec4ByW(t *testing.T) {
	v := Vec4{2, 3, 4, 5}.ByW()
	checkFloats(t, v[:], 2.0/5, 3.0/5, 4.0/5)
}

func TestVec4String(t *testing.T) {
	v := Vec4{3, 4, 5, -1}
	checkString(t, v.String(), "(3.00 4.00 5.00 -1.00)")
}

func TestAddVec4(t *testing.T) {
	a := Vec4{1, 2, 3, 4}
	b := Vec4{4, 8, 5, 9}
	c := Vec4{7, 3, 6, 2}
	sum := AddVec4(a, b, c)
	checkFloats(t, sum[:], 12, 13, 14, 15)
}

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

func TestMulMat2(t *testing.T) {
	a := Mat2{
		1, 2,
		3, 4,
	}
	b := Mat2{
		4, 3,
		2, 1,
	}
	c := Mat2{
		3, 5,
		4, 2,
	}
	prod := Mul2(a, b, c)
	checkFloats(t, prod[:],
		44, 50,
		112, 126,
	)
}

func TestMat2Transposed(t *testing.T) {
	m := Mat2{
		1, 2,
		3, 4,
	}.Transposed()
	checkFloats(t, m[:],
		1, 3,
		2, 4,
	)
}

func TestMat2Homogeneous(t *testing.T) {
	m := Mat2{
		1, 2,
		3, 4,
	}.Homogeneous()
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

func TestMulMat3(t *testing.T) {
	a := Mat3{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	b := Mat3{
		2, 5, 4,
		3, 5, 6,
		7, 8, 5,
	}
	c := Mat3{
		5, 4, 1,
		2, 6, 8,
		9, 7, 3,
	}
	prod := Mul3(a, b, c)
	checkFloats(t, prod[:],
		502, 567, 434,
		1195, 1350, 1037,
		1888, 2133, 1640,
	)
}

func TestMat3Transposed(t *testing.T) {
	m := Mat3{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}.Transposed()
	checkFloats(t, m[:],
		1, 4, 7,
		2, 5, 8,
		3, 6, 9,
	)
}

func TestMat3Homogeneous(t *testing.T) {
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

func TestMat2x3Add(t *testing.T) {
	m := Mat2x3{
		1, 2, 3,
		4, 5, 6,
	}.Add(Mat2x3{
		3, 2, 5,
		6, 4, 7,
	})
	checkFloats(t, m[:],
		4, 4, 8,
		10, 9, 13,
	)
}

func TestMat2x3Sub(t *testing.T) {
	m := Mat2x3{
		1, 2, 3,
		4, 5, 6,
	}.Sub(Mat2x3{
		3, 2, 5,
		6, 4, 7,
	})
	checkFloats(t, m[:],
		-2, 0, -2,
		-2, 1, -1,
	)
}

func TestMat2x3Mul(t *testing.T) {
	m := Mat2x3{
		1, 2, 3,
		4, 5, 6,
	}.Mul(Mat2x3{
		3, 2, 5,
		6, 4, 7,
	})
	checkFloats(t, m[:],
		15, 10, 22,
		42, 28, 61,
	)
}

func TestIdentity2x3(t *testing.T) {
	m := Identity2x3()
	checkFloats(t, m[:],
		1, 0, 0,
		0, 1, 0,
	)
}

func TestMulMat2x3(t *testing.T) {
	a := Mat2x3{
		1, 2, 3,
		4, 5, 6,
	}
	b := Mat2x3{
		2, 5, 4,
		3, 5, 6,
	}
	c := Mat2x3{
		5, 4, 1,
		2, 6, 8,
	}
	prod := Mul2x3(a, b, c)
	checkFloats(t, prod[:],
		70, 122, 147,
		205, 362, 435,
	)
}

func TestMat2x3ToMat3(t *testing.T) {
	m := Mat2x3{
		1, 2, 3,
		4, 5, 6,
	}.ToMat3()
	checkFloats(t, m[:],
		1, 2, 3,
		4, 5, 6,
		0, 0, 1,
	)
}

func TestMat2x3String(t *testing.T) {
	m := Mat2x3{
		1, 2, 3,
		4, 5, 6,
	}
	checkString(t, m.String(), "1.00 2.00 3.00\n4.00 5.00 6.00")
}

func TestMat4Add(t *testing.T) {
	m := Mat4{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}.Add(Mat4{
		5, 3, 6, 8,
		7, 4, 9, 6,
		3, 8, 5, 7,
		2, 4, 9, 5,
	})
	checkFloats(t, m[:],
		6, 5, 9, 12,
		12, 10, 16, 14,
		12, 18, 16, 19,
		15, 18, 24, 21,
	)
}

func TestMat4Sub(t *testing.T) {
	m := Mat4{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}.Sub(Mat4{
		5, 3, 6, 8,
		7, 4, 9, 6,
		3, 8, 5, 7,
		2, 4, 9, 5,
	})
	checkFloats(t, m[:],
		-4, -1, -3, -4,
		-2, 2, -2, 2,
		6, 2, 6, 5,
		11, 10, 6, 11,
	)
}

func TestMat4MulMat4(t *testing.T) {
	a := Mat4{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 1, 2, 3,
		4, 5, 6, 7,
	}
	b := Mat4{
		3, 4, 5, 6,
		7, 8, 9, 8,
		6, 4, 3, 2,
		4, 6, 7, 8,
	}
	m := a.Mul(b)
	checkFloats(t, m[:],
		51, 56, 60, 60,
		131, 144, 156, 156,
		58, 70, 81, 90,
		111, 122, 132, 132,
	)
}

func TestMulMat4(t *testing.T) {
	a := Mat4{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 1, 2, 3,
		4, 5, 6, 7,
	}
	b := Mat4{
		3, 4, 5, 6,
		7, 8, 9, 8,
		6, 4, 3, 2,
		4, 6, 7, 8,
	}
	c := Mat4{
		5, 4, 1, 8,
		2, 6, 8, 9,
		9, 7, 3, 7,
		6, 4, 5, 8,
	}
	prod := Mul4(a, b, c)
	checkFloats(t, prod[:],
		1267, 1200, 979, 1812,
		3283, 3104, 2531, 4684,
		1699, 1579, 1311, 2381,
		2779, 2628, 2143, 3966,
	)
}

func TestMat4Transposed(t *testing.T) {
	m := Mat4{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}.Transposed()
	checkFloats(t, m[:],
		1, 5, 9, 13,
		2, 6, 10, 14,
		3, 7, 11, 15,
		4, 8, 12, 16,
	)
}

func TestIdentity4(t *testing.T) {
	m := Identity4()
	checkFloats(t, m[:],
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}

func TestTranslate(t *testing.T) {
	m := Translate(2, 3, 4)
	checkFloats(t, m[:],
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		2, 3, 4, 1,
	)
}

func TestScale(t *testing.T) {
	m := Scale(2, 3, 4)
	checkFloats(t, m[:],
		2, 0, 0, 0,
		0, 3, 0, 0,
		0, 0, 4, 0,
		0, 0, 0, 1,
	)
}

func TestRotation(t *testing.T) {
	// We rotate vector v around different axes.
	v := Vec4{2, 3, 4, 1}
	check := func(m Mat4, x, y, z float32) {
		have := v.MulMat(m)
		checkFloats(t, have[:], x, y, z, 1)
	}

	x := Vec3{1, 0, 0}
	y := Vec3{0, 1, 0}
	z := Vec3{0, 0, 1}

	check(RotateRightHandX(0.25), 2, 4, -3)
	check(RotateRightHandAbout(x, 0.25), 2, 4, -3)
	check(RotateLeftHandX(0.25), 2, -4, 3)
	check(RotateLeftHandAbout(x, 0.25), 2, -4, 3)
	check(RotateRightHandY(0.25), -4, 3, 2)
	check(RotateRightHandAbout(y, 0.25), -4, 3, 2)
	check(RotateLeftHandY(0.25), 4, 3, -2)
	check(RotateLeftHandAbout(y, 0.25), 4, 3, -2)
	check(RotateRightHandZ(0.25), 3, -2, 4)
	check(RotateRightHandAbout(z, 0.25), 3, -2, 4)
	check(RotateLeftHandAbout(z, 0.25), -3, 2, 4)
}

func TestDecomposeAffine(t *testing.T) {
	// create a transformation with all components
	trans := Translate(1, -2, 3)
	scale := Scale(2, -3, 4)
	rot := RotateRightHandAbout(Vec3{3, -4, 5}, 1)
	m := Mul4(trans, scale, rot)
	// decompose and reconstruct the transformation matrix
	scale2, rot2, trans2 := DecomposeAffineTransform(m)
	m2 := Mul4(scale2, rot2, trans2)
	// the transformations must match
	checkFloats(t, m2[:], m[:]...)
}

func checkString(t *testing.T, have, want string) {
	if have != want {
		t.Errorf("string differs, have %q but want %q", have, want)
	}
}

func checkFloat(t *testing.T, have, want float32) {
	if have != want {
		t.Errorf("float differs, have %f but want %f", have, want)
	}
}

func checkFloats(t *testing.T, have []float32, want ...float32) {
	eq := len(have) == len(want)
	if eq {
		for i := range have {
			if have[i] != want[i] {
				eq = false
			}
		}
	}
	if !eq {
		t.Errorf("floats differ, have\n%v\nbut want\n%v", have, want)
	}
}
