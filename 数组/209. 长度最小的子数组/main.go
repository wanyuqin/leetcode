package main

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	flag := false
	res := len(nums)
	i := 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			flag = true
			if res >= j-i {
				res = j - i
			}
			sum -= nums[i]
			i++
		}
	}
	if !flag {
		return 0
	}
	return res + 1
}
