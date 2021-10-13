package main

func searchNearest(arr []int, target int, left int, right int, accurate bool) int {
	if left > right {
		if accurate {
			return -1
		} else {
			return right
		}
	}

	middle := (left + right) / 2
	if arr[middle] == target {
		return middle
	} else if arr[middle] < target {
		return searchNearest(arr, target, middle+1, right, accurate)
	} else {
		return searchNearest(arr, target, left, middle-1, accurate)
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	leftNums := make([]int, m)
	for i := 0; i < m; i++ {
		leftNums[i] = matrix[i][0]
	}

	row := searchNearest(leftNums, target, 0, m-1, false)
	if row < 0 {
		return false
	}

	idx := searchNearest(matrix[row], target, 0, n-1, true)
	return idx > -1
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	ret := searchMatrix(matrix, 0)
	println(ret)
}
