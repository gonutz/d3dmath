package d3dmath

import "testing"

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
