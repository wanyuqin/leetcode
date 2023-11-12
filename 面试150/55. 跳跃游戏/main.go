package main

func main() {
	println(canJump([]int{2, 0}))
}

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	// 每次都判断覆盖范围是否可以到达终点，并且取最大的覆盖范围即可
	cover := 0
	for i := 0; i <= cover; i++ {
		cover = max(i+nums[i], cover)
		if cover >= len(nums) {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
