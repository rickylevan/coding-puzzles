package lib

// 3.1

// small for testing ease. caller must not push >= 3
// items to A, B, or C
const arrLen = 9

type triStack struct {
	arr       [arrLen]int
	stackIdxs [3]int
}

func NewTriStack() *triStack {
	var ts triStack
	ts.stackIdxs[0] = 0 * arrLen / 3
	ts.stackIdxs[1] = 1 * arrLen / 3
	ts.stackIdxs[2] = 2 * arrLen / 3
	return &ts
}

// stack selects which stack to augment
func (ts *triStack) push(x int, stack int) {
	ts.arr[ts.stackIdxs[stack]] = x
	ts.stackIdxs[stack]++
}

func (ts *triStack) pop(stack int) int {
	val := ts.arr[ts.stackIdxs[stack]]
	ts.stackIdxs[stack]--
	return val
}

// 3.2

// Stack with O(1) min() operation too

// first let's make a stack
type stack struct {
	data []int
	next int
}

func NewStack() *stack {
	var s stack
	s.data = make([]int, 0)
	return &s
}

func (s *stack) push(val int) {
	s.data = append(s.data, val)
	s.next++
}

func (s *stack) pop() *int {
	if len(s.data) == 0 {
		return nil
	} else {
		out := s.data[s.next-1]
		s.next--
		return &out
	}
}

func (s *stack) peek() *int {
	if len(s.data) == 0 {
		return nil
	} else {
		return &s.data[s.next-1]
	}
}

// then use two stacks for data and min values
type minStack struct {
	data *stack
	mins *stack
}

func NewMinStack() *minStack {
	var ms minStack
	ms.data = NewStack()
	ms.mins = NewStack()
	return &ms
}

func (ms *minStack) push(val int) {
	ms.data.push(val)
	if ms.mins.peek() == nil {
		ms.mins.push(val)
	} else if val <= *ms.mins.peek() {
		ms.mins.push(val)
	}
}

func (ms *minStack) pop() *int {
	val := ms.data.pop()
	if min := ms.mins.peek(); min != nil && *min == *val {
		ms.mins.pop()
	}
	return val
}

func (ms *minStack) min() *int {
	return ms.mins.peek()
}
