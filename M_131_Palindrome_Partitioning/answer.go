package M_131_Palindrome_Partitioning

func analyze(s string, sLen int) [][]bool {
	ret := make([][]bool, sLen)
	for i := 0; i < sLen; i++ {
		ret[i] = make([]bool, sLen)
		for j := 0; j < sLen; j++ {
			ret[i][j] = false
		}
	}

	for i := sLen - 1; i >= 0; i-- {
		for j := i; j < sLen; j++ {
			if s[i] == s[j] && (j-i < 2 || ret[i+1][j-1]) {
				ret[i][j] = true
			}
		}
	}

	return ret
}

func search(s string, sLen int, idx int, dict [][]bool, tmp *[]string, res *[][]string) {
	if idx == sLen {
		*res = append(*res, *tmp)
		return
	}

	for i := idx; i < sLen; i++ {
		if dict[idx][i] {
			*tmp = append(*tmp, s[idx:i+1])
			search(s, sLen, i+1, dict, tmp, res)
			*tmp = append([]string{}, (*tmp)[:len(*tmp)-1]...)
		}
	}
}

func partition(s string) [][]string {
	tmp := make([]string, 0)
	res := make([][]string, 0)
	sLen := len(s)
	dict := analyze(s, sLen)
	search(s, len(s), 0, dict, &tmp, &res)
	return res
}
