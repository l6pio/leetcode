package main

// https://leetcode-cn.com/problems/unique-paths-ii/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 {
		obstacleGrid[0][0] = 0
	} else {
		obstacleGrid[0][0] = 1
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			obstacleGrid[i][0] = 0
		} else {
			obstacleGrid[i][0] = obstacleGrid[i-1][0]
		}
	}

	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			obstacleGrid[0][i] = 0
		} else {
			obstacleGrid[0][i] = obstacleGrid[0][i-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}
	return obstacleGrid[m-1][n-1]
}

func main() {
	//ret := uniquePathsWithObstacles([][]int{
	//	{0, 0, 0}, {0, 1, 0}, {0, 0, 0},
	//})
	//ret := uniquePathsWithObstacles([][]int{
	//	{0},
	//})
	ret := uniquePathsWithObstacles([][]int{
		{1, 0},
	})
	println(ret)
}
