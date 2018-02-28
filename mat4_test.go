package d3dmath

import "testing"

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

func TestDecomposeAffine(t *testing.T) {
	// create a transformation with all components
	trans := Translate(1, -2, 3)
	scale := Scale(2, -3, 4)
	rot := RotateAbout(Vec3{3, -4, 5}, 1)
	m := Mul4(trans, scale, rot)
	// decompose and reconstruct the transformation matrix
	scale2, rot2, trans2 := DecomposeAffineTransform(m)
	m2 := Mul4(scale2, rot2, trans2)
	// the transformations must match
	checkFloats(t, m2[:], m[:]...)
}
