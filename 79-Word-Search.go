package main

import "fmt"

// https://leetcode-cn.com/problems/word-search/

func existString(board [][]byte, word string, y int, x int, rows int, cols int) bool {
	if board[y][x] == word[0] {
		if len(word) == 1 {
			return true
		} else {
			top, down, left, right := false, false, false, false

			tmp := board[y][x]
			board[y][x] = 0
			if y > 0 && board[y-1][x] > 0 {
				top = existString(board, word[1:], y-1, x, rows, cols)
			}
			if y < rows-1 && board[y+1][x] > 0 {
				down = existString(board, word[1:], y+1, x, rows, cols)
			}
			if x > 0 && board[y][x-1] > 0 {
				left = existString(board, word[1:], y, x-1, rows, cols)
			}
			if x < cols-1 && board[y][x+1] > 0 {
				right = existString(board, word[1:], y, x+1, rows, cols)
			}
			board[y][x] = tmp
			return top || down || left || right
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	exists := false
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			exists = exists || existString(board, word, y, x, rows, cols)
		}
	}
	return exists
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	ret := exist(board, "ABCCED")
	//board := [][]byte{
	//	{'A'},
	//}
	//ret := exist(board, "A")
	//board := [][]byte{
	//	{'C', 'A', 'A'},
	//	{'A', 'A', 'A'},
	//	{'B', 'C', 'D'},
	//}
	//ret := exist(board, "AAB")
	println(ret)
	fmt.Printf("%+v", board)
}
