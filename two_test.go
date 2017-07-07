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

func Test22RB(t *testing.T) {
	var rb ringBuf
	rb.New(uint(3))
	rb.Push(1)
	rb.Push(2)
	rb.Push(3)
	rb.Push(4)
	rb.Push(5)
	rb.Pop()

	if !reflect.DeepEqual(rb.data, []int{4, 5, 3}) {
		t.Fail()
	}
	if rb.cur != 1 {
		t.Fail()
	}
}

func Test22(t *testing.T) {
	l := &node{1, &node{2, &node{3, &node{4, nil}}}}
	var k uint = 2
	if KthFromLast(l, k) != 3 {
		t.Fail()
	}
}
