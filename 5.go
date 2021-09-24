package main

func longestPalindrome(s string) string {
	var m [][]bool
	sLen := len(s)

	m = make([][]bool, sLen)
	for i := 0; i < sLen; i++ {
		m[i] = make([]bool, sLen)
	}

	ret := ""
	maxLen := 0
	for i := 0; i < sLen; i++ {
		for j := 0; j <= i; j++ {
			candidateLen := 0
			if s[j] == s[i] && (j == i || j == i-1 || m[j+1][i-1]) {
				m[j][i] = true
				candidateLen = i - j + 1
			}
			if candidateLen > maxLen {
				maxLen = candidateLen
				ret = s[j : i+1]
			}
		}
	}
	return ret
}

//func main() {
//	longestPalindrome("")
//}
