package lib

import (
	"fmt"
	"strconv"
)

type node struct {
	val  int
	next *node
}

func (l *node) String() string {
	var out []rune
	out = append(out, []rune("&[ ")...)
	out = append(out, []rune(strconv.Itoa(l.val)+" ")...)
	for l.next != nil {
		l = l.next
		out = append(out, []rune(strconv.Itoa(l.val)+" ")...)
	}
	out = append(out, ']')

	return string(out)
}

func (l *node) Print() {
	fmt.Println(l.String())
}

// 2.1

// remove duplicates from a linked list.
// The FOLLOWUP is a question, not a command haha. The answer is
// to roll through the thing in n^2 time, each time trimming all
// the dups then bumping forward the head, rinse and repeat

func RemoveDups(l *node) {
	if l == nil || l.next == nil {
		return
	}

	a := l
	b := l.next

	seen := make(map[int]struct{})
	seen[a.val] = struct{}{}

	for b.next != nil {
		cur := b.val
		if _, ok := seen[b.val]; ok {
			b = b.next
			a.next = b
		} else {
			b = b.next
			a = a.next
		}
		seen[cur] = struct{}{}
	}
}
