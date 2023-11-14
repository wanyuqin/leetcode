package main

import "fmt"

func main() {
	fmt.Println(groupAnagrams([]string{""}))

	//a := [100]int{}
	//b := [100]int{}
	//
	//a[0] = 1
	//b[0] = 1
	//fmt.Println('a' - 97)
}

// FIXME 可以做优化 key还是表示字母出现的次数 value可以直接表示为对应的分组结果 make(map[[26]int][]string)
func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0, 0)
	if len(strs) == 1 {
		r := [][]string{strs}
		res = append(res, r...)
	}
	hm := make(map[[26]int]int)
	count := 0
	// 创建hash 表
	for i := range strs {
		h := makeStrHash(strs[i])
		// 存放到map 并且记录下标
		if _, ok := hm[h]; !ok {
			hm[h] = count
			count++
		}
	}
	res = make([][]string, count)

	for i := range res {
		res[i] = make([]string, 0)
	}
	// hm 可以定位返回值有几个组
	for i := range strs {
		h := makeStrHash(strs[i])
		index := hm[h]
		group := res[index]
		group = append(group, strs[i])
		res[index] = group
	}

	return res
}

func makeStrHash(s string) [26]int {
	h := [26]int{}
	// 每个字母都成立一个hash表
	for j := range s {
		h[s[j]-'a']++
	}
	return h
}
