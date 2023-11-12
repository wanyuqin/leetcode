package main

func main() {

}

func isAnagram(s string, t string) bool {
	m := make(map[byte]int)
	for i := range s {
		if v, ok := m[s[i]]; ok {
			v++
			m[s[i]] = v
			continue
		}

		m[s[i]] = 1
	}

	for i := range t {
		if v, ok := m[t[i]]; ok {
			v--

			m[t[i]] = v
			if v == 0 {
				delete(m, t[i])
			}

		} else {
			return false
		}
	}

	if len(m) == 0 {
		return true
	}
	return false
}
