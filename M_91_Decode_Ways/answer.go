package M_91_Decode_Ways

func numDecodings(s string) int {
	sLen := len(s)
	count := make([]int, sLen+1)
	count[0] = 1
	for i := 1; i <= sLen; i++ {
		if s[i-1] != '0' {
			count[i] = count[i-1]
		}
		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1]-'0' <= 6)) {
			count[i] += count[i-2]
		}
	}
	return count[sLen]
}
