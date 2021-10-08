package main

// https://leetcode-cn.com/problems/unique-paths/

func uniquePaths(m int, n int) int {
	c := make([][]int, m)
	for i := 0; i < m; i++ {
		c[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		c[i][0] = 1
	}

	for j := 0; j < n; j++ {
		c[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			c[i][j] = c[i-1][j] + c[i][j-1]
		}
	}
	return c[m-1][n-1]
}

func main() {
	ret := uniquePaths(3, 3)
	println(ret)
}
