package main

func main() {
	println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

// 去重复
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	// 去重
	for i := range nums {
		m[nums[i]] = true
	}
	//
	ans := 0

	for key, _ := range m {
		// 若该数前面不存在数，则把该数当作连续序列的第一个数
		if !m[key-1] {
			long := 1
			for m[key+1] {
				key++
				long++
			}
			if long > ans {
				ans = long
			}
		}
	}

	return ans
}
