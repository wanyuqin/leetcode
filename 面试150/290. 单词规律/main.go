package main

import "strings"

func main() {
	println(wordPattern("abc", "dog cat dog"))
}

func wordPattern(pattern string, s string) bool {
	ss := strings.Split(s, " ")
	pm := make(map[byte]string)
	sm := make(map[string]byte)
	if len(pattern) != len(ss) {
		return false
	}

	for i := range pattern {
		if value, ok := pm[pattern[i]]; !ok {
			pm[pattern[i]] = ss[i]

		} else {
			if ss[i] != value {
				return false
			}
		}

		if value, ok := sm[ss[i]]; !ok {
			sm[ss[i]] = pattern[i]

		} else {
			if pattern[i] != value {
				return false
			}
		}

	}

	return true
}
