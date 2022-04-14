package main

import "fmt"

func main() {
	input1 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(input1)
	fmt.Println(input1)

	input2 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLifeInPlace(input2)
	fmt.Println(input2)

}
func gameOfLife(board [][]int) {

	newBoard := make([][]int, len(board))
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			newBoard[row] = append(newBoard[row], board[row][col])
		}

	}
	fmt.Println(newBoard)
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			live := 0
			if row-1 > -1 && col-1 > -1 {
				if newBoard[row-1][col-1] == 1 {
					live++
				}
			}
			if row-1 > -1 {
				if newBoard[row-1][col] == 1 {
					live++
				}
			}
			if row-1 > -1 && col+1 < len(newBoard[row]) {
				if newBoard[row-1][col+1] == 1 {
					live++
				}
			}
			if col-1 > -1 {
				if newBoard[row][col-1] == 1 {
					live++
				}
			}

			if col+1 < len(newBoard[row]) {
				if newBoard[row][col+1] == 1 {
					live++
				}
			}
			if col-1 > -1 && row+1 < len(newBoard) {
				if newBoard[row+1][col-1] == 1 {
					live++
				}
			}

			if row+1 < len(newBoard) {
				if newBoard[row+1][col] == 1 {
					live++
				}
			}
			if col+1 < len(newBoard[row]) && row+1 < len(newBoard) {
				if newBoard[row+1][col+1] == 1 {
					live++
				}
			}
			if newBoard[row][col] == 1 {
				if live < 2 {
					board[row][col] = 0
				} else if live > 3 {
					board[row][col] = 0
				}

			} else {
				if live == 3 {
					board[row][col] = 1
				}
			}
		}
	}

}

func gameOfLifeInPlace(board [][]int) {

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			live := 0
			if row-1 > -1 && col-1 > -1 {
				if board[row-1][col-1] == 1 || board[row-1][col-1] == 3 {
					live++
				}
			}
			if row-1 > -1 {
				if board[row-1][col] == 1 || board[row-1][col] == 3 {
					live++
				}
			}
			if row-1 > -1 && col+1 < len(board[row]) {
				if board[row-1][col+1] == 1 || board[row-1][col+1] == 3 {
					live++
				}
			}
			if col-1 > -1 {
				if board[row][col-1] == 1 || board[row][col-1] == 3 {
					live++
				}
			}

			if col+1 < len(board[row]) {
				if board[row][col+1] == 1 || board[row][col+1] == 3 {
					live++
				}
			}
			if col-1 > -1 && row+1 < len(board) {
				if board[row+1][col-1] == 1 || board[row+1][col-1] == 3 {
					live++
				}
			}

			if row+1 < len(board) {
				if board[row+1][col] == 1 || board[row+1][col] == 3 {
					live++
				}
			}
			if col+1 < len(board[row]) && row+1 < len(board) {
				if board[row+1][col+1] == 1 || board[row+1][col+1] == 3 {
					live++
				}
			}
			if board[row][col] == 1 {
				if live < 2 {
					board[row][col] = 3
				} else if live > 3 {
					board[row][col] = 3
				}

			} else {
				if live == 3 {
					board[row][col] = 2
				}
			}
		}
	}
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] == 2 {
				board[row][col] = 1
			}
			if board[row][col] == 3 {
				board[row][col] = 0
			}

		}
	}

}
