package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(merge([][]int{{1, 9}, {2, 5}, {19, 20}, {10, 11}, {12, 20}, {0, 3}, {0, 1}, {0, 2}}))
}

// 合并区间
// 比较前一个的最小是不是小于当前的最大，最大就合并
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// 先排序，让相邻的区间尽可能挨在一起
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	// 先定义一个结果
	res := make([][]int, 0, len(intervals))
	// 将intervals的第一个元素给进去
	res = append(res, intervals[0])

	// 迭代interval
	for i := 1; i < len(intervals); i++ {
		// 判断
		if intervals[i][0] <= res[len(res)-1][1] {
			// 合并
			// 确定右边界限
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
			continue
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b

	}
	return a
}

func doMerge(start []int, end []int) ([]int, bool) {
	if start[1] > end[0] {
		return []int{start[0], end[1]}, true
	}
	return nil, false
}
