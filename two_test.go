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

func Test23(t *testing.T) {
	l := &node{1, &node{2, &node{3, &node{4, nil}}}}
	s := &node{1, &node{2, &node{4, nil}}}
	Snip(l.next.next)
	if !reflect.DeepEqual(l, s) {
		t.Fail()
	}
}

func Test24(t *testing.T) {
	l := &node{20, &node{5, &node{19, &node{4, nil}}}}
	s := &node{5, &node{4, &node{20, &node{19, nil}}}}
	if !reflect.DeepEqual(Partition(l, 8), s) {
		t.Fail()
	}
}

func Test25(t *testing.T) {
	l := &node{2, &node{5, &node{4, nil}}}
	if LinkedSum(l) != 452 {
		t.Fail()
	}
}

func Test26(t *testing.T) {
	var a, b, c, d node
	a = node{'A', &b}
	b = node{'B', &c}
	c = node{'C', &d}
	d = node{'D', &b}

	if FindLoop(&a) != &b {
		t.Fail()
	}
}

func Test27(t *testing.T) {
	l := &node{20, &node{5, &node{5, &node{20, nil}}}}
	if !IsPalin(l) {
		t.Error("oops")
		t.Fail()
	}
}
