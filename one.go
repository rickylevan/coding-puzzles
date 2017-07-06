package lib

import "fmt"

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

	fmt.Println(out)

	return out
}
