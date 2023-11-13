package main

func main() {
	println(maxArea([]int{1, 1}))
}

// 双指针思路
// 容量最大
// 容量 = 较小值*index
// 为了让容量最大 那么在移动的过程中，index是在减小的，所以在指针移动的时候，移动数字比较小的一遍
func maxArea(height []int) int {
	max := 0
	i := 0
	j := len(height) - 1
	for i < j {
		mu := min(height[i], height[j]) * (j - i)
		if mu > max {
			max = mu
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}

	}
	return max
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
