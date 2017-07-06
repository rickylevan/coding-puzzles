package lib

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
