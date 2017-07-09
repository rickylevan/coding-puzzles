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
