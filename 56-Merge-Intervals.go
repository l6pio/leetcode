package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/merge-intervals/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	intervalsLen := len(intervals)
	var ret [][]int
	for i := 0; i < intervalsLen; i++ {
		start := intervals[i][0]
		end := intervals[i][1]

		for i < intervalsLen-1 && intervals[i+1][0] <= end {
			if end < intervals[i+1][1] {
				end = intervals[i+1][1]
			}
			i++
		}

		ret = append(ret, []int{start, end})
	}
	return ret
}

func main() {
	//ret := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	ret := merge([][]int{{1, 4}, {2, 3}})
	fmt.Printf("%+v", ret)
}
