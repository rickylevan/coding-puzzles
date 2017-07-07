package lib

import (
	"bytes"
	"reflect"
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

func Test16(t *testing.T) {
	M := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	N := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	Rotate90(M)
	if !reflect.DeepEqual(M, N) {
		t.Error("Matrices not deep equal")
	}
}

func Test17(t *testing.T) {
	M := [][]int{{3, 4, 2}, {0, 1, 5}, {2, 8, 9}}
	N := [][]int{{0, 4, 2}, {0, 0, 0}, {0, 8, 9}}

	ZeroCrosses(M)
	if !reflect.DeepEqual(M, N) {
		t.Error("Matrices not deep equal")
	}

}

func Test18Sub(t *testing.T) {
	if !IsSubstring("overbearing", "bear") {
		t.Fail()
	}
	if IsSubstring("overbearing", "fool") {
		t.Fail()
	}
}

func Test18Rot(t *testing.T) {
	if !IsRotation("waterbottle", "erbottlewat") {
		t.Fail()
	}
}
