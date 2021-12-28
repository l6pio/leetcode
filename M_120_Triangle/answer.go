package M_120_Triangle

func minimumTotal(triangle [][]int) int {
	depth := len(triangle)

	if depth == 1 {
		return triangle[0][0]
	}

	for i := 1; i < depth; i++ {
		levelLen := len(triangle[i])
		for j := 0; j < levelLen; j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else if j == levelLen-1 {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				if triangle[i-1][j-1] < triangle[i-1][j] {
					triangle[i][j] += triangle[i-1][j-1]
				} else {
					triangle[i][j] += triangle[i-1][j]
				}
			}
		}
	}

	min := triangle[depth-1][0]
	for j := 1; j < len(triangle[depth-1]); j++ {
		if triangle[depth-1][j] < min {
			min = triangle[depth-1][j]
		}
	}
	return min
}
