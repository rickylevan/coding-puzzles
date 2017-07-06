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
