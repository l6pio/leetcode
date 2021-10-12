package main

import "fmt"

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	isLeftMostZero := false
	isTopMostZero := false

	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			isLeftMostZero = true
		}
	}

	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			isTopMostZero = true
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for i := 1; i < n; i++ {
		if matrix[0][i] == 0 {
			for j := 0; j < m; j++ {
				matrix[j][i] = 0
			}
		}
	}

	if isLeftMostZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

	if isTopMostZero {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
}

func main() {
	matrix := [][]int{
		{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5},
	}
	setZeroes(matrix)
	fmt.Printf("%v\n", matrix)
}
