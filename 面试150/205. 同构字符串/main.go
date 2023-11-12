package main

func main() {
	println(isIsomorphic("bbbaaaba", "aaabbbba"))
}

func isIsomorphic(s string, t string) bool {
	sm := make(map[byte]byte)
	tm := make(map[byte]byte)

	for i := range s {
		if _, ok := sm[s[i]]; !ok {
			sm[s[i]] = t[i]
		}

		if _, ok := tm[t[i]]; !ok {
			tm[t[i]] = s[i]
		}

		if (sm[s[i]] != t[i]) || (tm[t[i]] != s[i]) {
			return false
		}
	}

	return true
}

//
//func isIsomorphic(s string, t string) bool {
//	sb := make([]int, 128)
//	tb := make([]int, 128)
//	replace := make(map[byte]byte)
//	for i := range s {
//		sb[s[i]]++
//	}
//
//	for i := range t {
//		tb[t[i]]++
//	}
//
//	for i := range s {
//		if sb[s[i]] != tb[t[i]] {
//			return false
//		}
//		replace[s[i]] = t[i]
//
//	}
//	builder := strings.Builder{}
//	for i := range s {
//		builder.WriteByte(replace[s[i]])
//	}
//	res := builder.String()
//	if res != t {
//		return false
//	}
//	return true
//}
