package main

func GenerateParenthesis(tmp string, ret *[]string, left int, right int) {
	if left == 0 && right == 0 {
		*ret = append(*ret, tmp)
		return
	}

	if left > 0 {
		GenerateParenthesis(tmp+"(", ret, left-1, right)
	}

	if right > 0 && left < right {
		GenerateParenthesis(tmp+")", ret, left, right-1)
	}
}

func generateParenthesis(n int) []string {
	var ret []string
	GenerateParenthesis("", &ret, n, n)
	return ret
}

//func main() {
//	generateParenthesis(3)
//}
