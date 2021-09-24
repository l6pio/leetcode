package main

func longestCommonPrefix(strs []string) string {
	strsLen := len(strs)

	if strsLen == 0 {
		return ""
	}

	if strsLen == 1 {
		return strs[0]
	}

	lens := make([]int, strsLen)
	for i, v := range strs {
		lens[i] = len(v)
	}

	right := 0
	for right < lens[0] {
		for i := 1; i < strsLen; i++ {
			if right >= lens[i] || strs[i][right] != strs[0][right] {
				return strs[0][:right]
			}
		}
		right += 1
	}
	return strs[0][:right]
}

//func main() {
//	prefix := longestCommonPrefix([]string{"dog"})
//	println(prefix)
//}
