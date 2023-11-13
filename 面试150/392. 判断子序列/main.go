package main

func main() {
	println(isSubsequence("b", "abc"))
}

// s 是否为t的子序列 那么 s.len < t.len
func isSubsequence(s string, t string) bool {
	sp := 0
	tp := 0
	count := 0
	if len(s) == 0 {
		return true
	}
	for tp <= len(t)-1 && sp <= len(s)-1 {
		if s[sp] == t[tp] {
			sp++
			count++
		}
		tp++
	}
	if count == len(s) {
		return true
	}

	return false
}
