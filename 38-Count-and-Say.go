package main

import "strconv"

// https://leetcode-cn.com/problems/count-and-say/

func countAndSay(n int) string {
	ret := "1"
	for i := 1; i < n; i++ {
		tmp := ""
		count := 1
		left := 0
		for left < len(ret) - 1 {
			if ret[left] == ret[left + 1] {
				count += 1
				left += 1
				continue
			}
			tmp += strconv.Itoa(count) + string(ret[left])
			count = 1
			left += 1
		}
		tmp += strconv.Itoa(count) + string(ret[left])
		ret = tmp
	}
	return ret
}

func main() {
	say := countAndSay(10)
	println(say)
}
