package lib

import "strconv"

// 1.1

// free solution
func AllUniqueFree(s string) bool {
	m := make(map[rune]struct{})
	for _, c := range s {
		if _, ok := m[c]; ok {
			return false
		}
		m[c] = struct{}{}
	}
	return true
}

// no extra data structure
func AllUniqueSimple(s string) bool {
	for i, v1 := range s {
		for _, v2 := range s[i+1:] {
			if v1 == v2 {
				return false
			}
		}
	}
	return true
}

// 1.2

// reverse a null-terminated list.
// len(chars) must be >= 1
func ReverseStyleC(chars []byte) (out []byte) {

	idx := 0
	for chars[idx] != byte(0) {
		idx++
	}

	idx--

	for idx >= 0 {
		out = append(out, chars[idx])
		idx--
	}

	// put in a fresh null terminator
	out = append(out, byte(0))

	return out
}

// 1.3

// determine if two strings are permutations of
// each other

func ArePermutations(s, t string) bool {

	deepEqual := func(m1, m2 map[rune]int) bool {
		for k1, c1 := range m1 {
			if c2, ok := m2[k1]; !ok {
				return false
			} else if c1 != c2 {
				return false
			}
		}
		return true
	}

	scoop := func(s string) map[rune]int {
		m := make(map[rune]int)
		for _, r := range s {
			m[r]++
		}
		return m
	}

	return deepEqual(scoop(s), scoop(t))
}

// 1.4

// replace all spaces with '%20'
// ("true" len is hard to program)
func SillyReplace(s string) string {
	var out []rune
	for _, c := range s {
		if c == ' ' {
			out = append(out, []rune("%20")...)
		} else {
			out = append(out, c)
		}
	}

	return string(out)
}

// 1.5

// compress like so: aaabcc -> a3b1c2
// ignoring the return original string if smaller,
// that's weird and inelegant

func Compress(s string) string {
	if len(s) == 0 {
		return ""
	}

	serialize := func(cur rune, count int) []rune {
		tmp := string(cur) + strconv.Itoa(count)
		return []rune(tmp)
	}

	var out []rune

	cur := rune(s[0])
	count := 1
	for _, c := range s[1:] {
		if c == cur {
			count++
		} else {
			out = append(out, serialize(cur, count)...)
			count = 1
		}
		cur = c
	}

	out = append(out, serialize(cur, count)...)

	return string(out)
}

// 1.6

// rotate an NxN matrix, where each pixel in the image is 4
// bytes, by 90 degrees. Can I do it in place? Surely the
// answer is Yes. Just using simple ints WLOG.

func Rotate90(M [][]int) {
	if len(M) <= 1 {
		return
	}

	plen := len(M) - 1 // partition length of 4 non-overlaps

	ax, ay := 0, 0
	bx, by := 0, plen
	cx, cy := plen, plen
	dx, dy := plen, 0

	var buf [4]int

	for i := 0; i < plen; i++ {
		// save the values
		buf[0] = M[ax][ay]
		buf[1] = M[bx][by]
		buf[2] = M[cx][cy]
		buf[3] = M[dx][dy]

		// then do the swaps
		M[ax][ay] = buf[3]
		M[bx][by] = buf[0]
		M[cx][cy] = buf[1]
		M[dx][dy] = buf[2]

		ay++
		bx++
		cy--
		dx--
	}
}

// 1.7

// given an MxN matrix, zero out the entire row and column
// of every entry with value 0

// will act in place
func ZeroCrosses(mat [][]int) {

	blankRows := make(map[int]struct{})
	blankCols := make(map[int]struct{})

	for i, row := range mat {
		for j, val := range row {
			if val == 0 {
				blankRows[i] = struct{}{}
				blankCols[j] = struct{}{}
			}
		}
	}

	for row := range blankRows {
		for j := range mat[row] {
			mat[row][j] = 0
		}
	}

	for col := range blankCols {
		for row := range mat {
			mat[row][col] = 0
		}
	}
}

// 1.8

// first need to write an isSubstring method

func IsSubstring(s, sub string) bool {
	if len(sub) > len(s) {
		return false
	}
	if s[:len(sub)] == sub {
		return true
	} else {
		return IsSubstring(s[1:], sub)
	}
}

func IsRotation(a, b string) bool {
	return IsSubstring(a+a, b)
}
