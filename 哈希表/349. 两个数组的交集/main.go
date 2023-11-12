package main

import "fmt"

func main() {
	ints := intersection([]int{1, 2, 2, 1}, []int{2, 2})
	fmt.Println(ints)
}

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	tmp := make([]int, 1001)
	res := make([]int, 0)
	for i := range nums1 {
		m[nums1[i]]++
	}

	for i := range nums2 {
		if _, ok := m[nums2[i]]; ok {
			//存在
			if tmp[nums2[i]] == 0 {
				tmp[nums2[i]]++
			}

		}
	}

	for i := range tmp {
		if tmp[i] > 0 {
			res = append(res, i)
		}
	}

	return res
}
