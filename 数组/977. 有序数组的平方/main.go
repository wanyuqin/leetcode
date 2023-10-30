package main

import "fmt"

func main() {
	res := sortedSquares([]int{-4, -1, 0, 3, 10})
	fmt.Println(res)
	res = sortedSquares_2([]int{-7, -3, 2, 3, 11})
	fmt.Println(res)

}

func sortedSquares(nums []int) []int {
	i := 0
	j := len(nums) - 1
	k := len(nums) - 1
	res := make([]int, len(nums))
	for k >= 0 {

		if nums[i]*nums[i] < nums[j]*nums[j] {
			res[k] = nums[j] * nums[j]
			j--
			k--
			continue
		}
		if nums[i]*nums[i] > nums[j]*nums[j] {
			res[k] = nums[i] * nums[i]
			i++
			k--
		}

		if nums[i]*nums[i] == nums[j]*nums[j] {
			if i != j {
				res[k] = nums[i] * nums[i]
				i++
				k--
				res[k] = nums[j] * nums[j]
				k--
				j--
			} else {
				res[k] = nums[j] * nums[j]
				k--
			}

		}

	}
	return res

}

func sortedSquares_2(nums []int) []int {
	n := len(nums)
	i, j, k := 0, n-1, n-1

	res := make([]int, len(nums))
	for i <= j {
		lr, rr := nums[i]*nums[i], nums[j]*nums[j]
		if lr < rr {
			res[k] = rr
			j--
		} else {
			res[k] = lr
			i++
		}
		k--

	}
	return res

}
