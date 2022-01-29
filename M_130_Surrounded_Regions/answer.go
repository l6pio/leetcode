package M_130_Surrounded_Regions

func Solve(board [][]byte, y int, x int, maxY int, maxX int) {
	if y < 0 || y > maxY || x < 0 || x > maxX {
		return
	}

	if board[y][x] == 'X' || board[y][x] == '#' {
		return
	}

	board[y][x] = '#'
	Solve(board, y, x-1, maxY, maxX)
	Solve(board, y, x+1, maxY, maxX)
	Solve(board, y-1, x, maxY, maxX)
	Solve(board, y+1, x, maxY, maxX)
}

func solve(board [][]byte) {
	rows := len(board)
	cols := len(board[0])

	for y := 0; y < rows; y++ {
		Solve(board, y, 0, rows-1, cols-1)
		Solve(board, y, cols-1, rows-1, cols-1)
	}

	for x := 1; x < cols-1; x++ {
		Solve(board, 0, x, rows-1, cols-1)
		Solve(board, rows-1, x, rows-1, cols-1)
	}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if board[y][x] == 'O' {
				board[y][x] = 'X'
			} else if board[y][x] == '#' {
				board[y][x] = 'O'
			}
		}
	}
}
