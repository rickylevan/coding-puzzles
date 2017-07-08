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

type ringBuf struct {
	cur  int // where to read or
	data []int
}

func (rb *ringBuf) New(size uint) {
	rb.cur = 0
	rb.data = make([]int, size)
}

func (rb *ringBuf) Push(val int) {
	rb.data[rb.cur] = val
	rb.cur = (rb.cur + 1) % len(rb.data)
}

func (rb *ringBuf) Pop() int {
	rb.cur = (rb.cur + (len(rb.data) - 1)) % len(rb.data)
	return rb.data[rb.cur]
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

// 2.2

// get the k-to-last element of a linked list
// of course, in the cute way, not a cheap way
func KthFromLast(l *node, k uint) int {
	var rb ringBuf
	rb.New(k)
	// assuming k >= 1, len(l) >= k
	for l.next != nil {
		rb.Push(l.val)
		l = l.next
	}
	// one final push once next is nil
	rb.Push(l.val)

	var out int
	for i := 0; i < int(k); i++ {
		out = rb.Pop()
	}

	return out
}

// 2.3

// receive a middle node only and cut it out
func Snip(mid *node) {
	*mid = *mid.next
}

// 2.4

// Partition a linked list around a value x, such that all nodes
// less than x come before all nodes greater than or equal to x

// the head might move, so we must return a new head
func Partition(list *node, x int) *node {

	// freeze head once we have first value <= x
	freezeHead := false

	head := list
	cur := list
	end := list

	count := 0
	for end.next != nil {
		count++
		end = end.next
	}

	var prev *node

	for i := 0; i < count; i++ {
		if cur.val < x {
			freezeHead = true
		} else {
			// skip over old value now
			if prev != nil {
				prev.next = cur.next
			}
			// write a new object
			end.next = &node{val: cur.val, next: nil}
			end = end.next
		}

		prev = cur
		cur = cur.next
		if !freezeHead {
			head = head.next
		}
	}

	return head
}

// 2.5

// Little endian linked list numbers
func LinkedSum(list *node) int {
	var LSH func(*node, int) int
	LSH = func(list *node, mult int) int {
		if list == nil {
			return 0
		} else {
			return list.val*mult + LSH(list.next, mult*10)
		}
	}

	return LSH(list, 1)
}

// 2.6

// Identify the first point at which a corrupt linked list loops
func FindLoop(list *node) *node {
	m := make(map[*node]struct{})
	// assume that there MUST be a loop
	for {
		if _, ok := m[list]; ok {
			return list
		} else {
			m[list] = struct{}{}
			list = list.next
		}
	}
}

// 2.7

// easy unsafe stack
type lilStack struct {
	data []int
}

func (stk *lilStack) push(val int) {
	stk.data = append(stk.data, val)
}

func (stk *lilStack) pop() int {
	out := stk.data[len(stk.data)-1]
	stk.data = stk.data[:len(stk.data)-1]
	return out
}

// Check if a linked list is a palindrome, in a non cheap way
// So let's
func IsPalin(list *node) bool {

	count := 0
	end := list
	for end != nil {
		end = end.next
		count++
	}

	ls := lilStack{}
	var i int
	for i = 0; i < count/2; i++ {
		ls.push(list.val)
		list = list.next
	}

	// skip middle value if odd length
	if count%2 != 0 {
		i++
		list = list.next
	}

	for i = i; i < count; i++ {
		if list.val != ls.pop() {
			return false
		}
		list = list.next
	}

	return true
}
