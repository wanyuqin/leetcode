package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	m1 := make(map[int]int)

	res := make([]int, 0, 2)
	for i := range nums {

		reduce := target - nums[i]
		if value, ok := m1[reduce]; ok {
			res = append(res, []int{i, value}...)
			break
		}
		m1[nums[i]] = i
	}
	return res

}
