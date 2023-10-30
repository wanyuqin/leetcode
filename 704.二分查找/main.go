package main

import "fmt"

func main() {
	res := search([]int{-1, 0, 3, 5, 9, 12}, 2)
	fmt.Print(res)
}

// 不重复
// 有序
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left = mid + 1

		}

		if target < nums[mid] {
			right = mid - 1

		}

		if target == nums[mid] {
			return mid
		}
	}
	return -1
}
