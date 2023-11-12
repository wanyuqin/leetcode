package main

import "math"

func main() {
	containsNearbyDuplicate([]int{1, 0, 1, 1}, 1)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	step := math.MaxInt
	for i := range nums {
		if value, ok := m[nums[i]]; !ok {
			m[nums[i]] = i
			continue
		} else {
			// 存在 比较index
			step = min(abs(value, i), step)
			if step <= k {
				return true
			}
			// 更新最新重复下标
			m[nums[i]] = i
		}
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}
