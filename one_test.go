package lib

import (
	"bytes"
	"testing"
)

func Test11A(t *testing.T) {
	if !AllUniqueFree("okay") {
		t.Fail()
	}
	if AllUniqueFree("yessir") {
		t.Fail()
	}
}

func Test11B(t *testing.T) {
	if !AllUniqueSimple("okay") {
		t.Fail()
	}
	if AllUniqueSimple("yessir") {
		t.Fail()
	}
}

func Test12(t *testing.T) {
	if !bytes.Equal(ReverseStyleC([]byte{'Y', 'E', 'S', 0}),
		[]byte{'S', 'E', 'Y', 0}) {
		t.Fail()
	}
}

func Test13(t *testing.T) {
	if ArePermutations("Okay", "Sure") {
		t.Fail()
	}
	if !ArePermutations("Yessir", "eYriss") {
		t.Fail()
	}
}

func Test14(t *testing.T) {
	if SillyReplace("okay man") != "okay%20man" {
		t.Fail()
	}
}

func Test15(t *testing.T) {
	if Compress("aaabcc") != "a3b1c2" {
		t.Error(Compress("aaabcc"))
	}
}

func Test17(t *testing.T) {
	m := [][]int{{3, 4, 2}, {0, 1, 5}, {2, 8, 9}}
	n := [][]int{{0, 4, 2}, {0, 0, 0}, {0, 8, 9}}

	ZeroCrosses(m)
	// too lazy to write deep equal
	if m[0][0] != n[0][0] {
		t.Fail()
	}

	if m[1][2] != n[1][2] {
		t.Fail()
	}

}
