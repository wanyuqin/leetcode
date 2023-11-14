package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(summaryRanges([]int{-1}))
}

// 数组升序
// 元素不重复
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	res := make([]string, 0)
	i := 0
	j := i + 1
	start := strconv.Itoa(nums[0])
	end := ""
	for j <= len(nums)-1 {
		if nums[i]+1 == nums[j] {
			i++
			j++
			continue
		}

		if nums[i]+1 != nums[j] {
			r := ""
			if i == 0 {
				r = strconv.Itoa(nums[i])
			} else if i > 0 && nums[i]-1 != nums[i-1] {
				r = strconv.Itoa(nums[i])
			} else {
				end = strconv.Itoa(nums[i])
				r = fmt.Sprintf("%s->%s", start, end)
			}
			res = append(res, r)
			i = j
			j = i + 1
			start = strconv.Itoa(nums[i])
			end = ""
		}
	}

	// 最后的判断
	if nums[i]-1 != nums[i-1] {
		res = append(res, strconv.Itoa(nums[i]))
	}
	if nums[i]-1 == nums[i-1] {
		end = strconv.Itoa(nums[i])
		res = append(res, fmt.Sprintf("%s->%s", start, end))
	}
	return res
}
