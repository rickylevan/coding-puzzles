package lib

import (
	"reflect"
	"testing"
)

func Test31(t *testing.T) {
	ts := NewTriStack()
	ts.push(7, 0)
	ts.push(5, 1)
	ts.push(3, 2)
	ts.push(6, 1)
	ts.pop(1)

	if !reflect.DeepEqual(ts,
		&triStack{arr: [arrLen]int{7, 0, 0, 5, 6, 0, 3, 0, 0},
			stackIdxs: [3]int{1, 4, 7}}) {
		t.Fail()
	}
}
