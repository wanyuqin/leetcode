package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(numbers []int, target int) []int {
	head := 0
	tail := len(numbers) - 1
	for head < tail {
		if numbers[head]+numbers[tail] > target {
			tail--
		}

		if numbers[head]+numbers[tail] < target {
			head++
		}

		if numbers[head]+numbers[tail] == target {
			return []int{head + 1, tail + 1}
		}
	}
	return nil

}

// 输入的数组是非递减 那就是递增
// 使用两个指针
//func twoSum(numbers []int, target int) []int {
//	first := 0
//	second := 1
//	for first < len(numbers)-1 {
//
//		if numbers[first]+numbers[second] == target {
//			return []int{first + 1, second + 1}
//		}
//
//		if numbers[first]+numbers[second] < target {
//			if second == len(numbers)-1 {
//				first++
//				second = first + 1
//				continue
//			}
//
//			second++
//		}
//
//		if numbers[first]+numbers[second] > target {
//			first++
//			second = first + 1
//
//		}
//	}
//
//	return nil
//
//}
