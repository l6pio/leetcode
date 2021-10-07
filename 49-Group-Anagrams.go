package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		val := string(bytes)
		arr, ok := m[val]
		if ok {
			m[val] = append(arr, str)
		} else {
			m[val] = []string{str}
		}
	}

	var ret [][]string
	for idx := range m {
		ret = append(ret, m[idx])
	}
	return ret
}

func main() {
	ret := groupAnagrams([]string{"bat", "bbdd"})
	fmt.Printf("%v", ret)
}
