package lib

import (
	"reflect"
	"testing"
)

func Test21(t *testing.T) {
	l := &node{3, &node{2, &node{3, &node{4, nil}}}}
	s := &node{3, &node{2, &node{4, nil}}}
	RemoveDups(l)

	// ooh deep equal works with pointers. Nice
	if !reflect.DeepEqual(l, s) {
		t.Error("Remove dup failed")
	}

}
