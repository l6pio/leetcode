package M_96_Unique_Binary_Search_Trees

func numTrees(n int) int {
	c := make([]int, n+1)
	c[0] = 1
	c[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			c[i] += c[j-1] * c[i-j]
		}
	}
	return c[n]
}
