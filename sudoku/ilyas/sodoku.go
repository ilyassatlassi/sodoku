package main

import (
	"fmt"
	"os"
)

func IsValid(s string) bool {
	i := 0
	for i < len(s) {
		if s[i] >= '1' && s[i] <= '9' || s[i] == '.' {
			i++
		} else {
			return (false)
		}
	}
	return (true)
}

func CheckArgs(Args []string) bool {
	if len(Args) != 10 {
		return (false)
	}
	for i := 1; i < len(Args); i++ {
		if len(Args[i]) != 9 || !IsValid(Args[i]) {
			return (false)
		}
	}
	return (true)
}

func PrintResults(board [][]rune) {
	fmt.Println("☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰")
	drawer := 1
	ver := 1
	for i := 0; i < len(board); i++ {
		fmt.Print("║")
		ver = 1
		for j := 0; j < len(board[i]); j++ {
			print(string(board[i][j]))
			if ver%3 == 0 {
				fmt.Print("║")
			}
			if j < len(board[i])-1 && ver%3 != 0 {
				fmt.Print(" ")
			}
			ver++
		}
		if drawer%3 == 0 {
			fmt.Print("\n☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰")
		}
		fmt.Print("\n")
		drawer++
	}
}

func CreateSodoku(Args []string) [][]rune {
	var board [][]rune
	for i := 1; i < len(Args); i++ {
		board = append(board, []rune(Args[i]))
	}
	return (board)
}

func ValidCol(board [][]rune, row int, col int, c rune) bool {
	for i := 0; i < 9; i++ {
		if c == board[i][col] {
			return (false)
		}
	}
	return (true)
}

func ValidRow(board [][]rune, row int, col int, c rune) bool {
	for j := 0; j < 9; j++ {
		if c == board[row][j] {
			return (false)
		}
	}
	return (true)
}

func ValidSquare(board [][]rune, row int, col int, c rune) bool {
	CurrentCol := col / 3
	CurrentRow := row / 3

	for i := 3 * CurrentRow; i < 3*(CurrentRow+1); i++ {
		for j := 3 * CurrentCol; j < 3*(CurrentCol+1); j++ {
			if c == board[i][j] {
				return (false)
			}
		}
	}
	return (true)
}

func ValidCell(board [][]rune, row int, col int, c rune) bool {
	if ValidRow(board, row, col, c) && ValidCol(board, row, col, c) && ValidSquare(board, row, col, c) {
		return (true)
	}
	return (false)
}

func SolveGame(board [][]rune, row, col int) bool {
	if row == 9 {
		return true
	} else if col == 9 {
		col = 0
		return (SolveGame(board, row+1, col))
	} else if string(board[row][col]) != "." {
		return (SolveGame(board, row, col+1))
	} else {
		for j := 1; j < 10; j++ {
			if ValidCell(board, row, col, rune(j+'0')) {
				board[row][col] = rune(j + '0')
				if SolveGame(board, row, col+1) {
					return true
				}
				board[row][col] = '.'
			}
		}
	}
	return false
}

func main() {
	if CheckArgs(os.Args) {
		board := CreateSodoku(os.Args)
		if SolveGame(board, 0, 0) {
			PrintResults(board)
		}
	} else {
		fmt.Println("Error")
	}
}
