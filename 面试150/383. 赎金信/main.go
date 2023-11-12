package main

func main() {
	canConstruct("aa", "aab")
}

// 26个英文字母
func canConstruct(ransomNote string, magazine string) bool {
	bucket := make([]int, 26)
	// 存放字母次数
	for i := range ransomNote {
		bucket[ransomNote[i]-97]++
	}

	for i := range magazine {
		bucket[magazine[i]-97]--
	}

	for i := range bucket {
		if bucket[i] > 0 {
			return false
		}
	}

	return true

}

//
//// 26个英文字母
//func canConstruct(ransomNote string, magazine string) bool {
//	rm := make(map[byte]int)
//	mm := make(map[byte]int)
//
//	for i := range ransomNote {
//		if value, ok := rm[ransomNote[i]]; ok {
//			rm[ransomNote[i]] = value + 1
//			continue
//		}
//		rm[ransomNote[i]] = 1
//
//	}
//
//	for i := range magazine {
//		if value, ok := mm[ransomNote[i]]; ok {
//			mm[ransomNote[i]] = value + 1
//			continue
//		}
//		mm[ransomNote[i]] = 1
//	}
//
//	for k, v := range rm {
//		if value, ok := mm[k]; !ok {
//			return false
//		} else {
//			if value < v {
//				return false
//			}
//		}
//
//	}
//
//	return true
//}
