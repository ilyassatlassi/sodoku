package main

import (
	"fmt"
	// "os"
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

func IsUnique(str string) bool {
	s := []rune(str)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] && s[i] != '.' && i != j {
				return (false)
			}
		}
	}
	return (true)
}

func CheckArgs(Args []string) bool {
	if len(Args) != 10 {
		return (false)
	}
	for i := 1; i < len(Args); i++ {
		if len(Args[i]) != 9 || !IsValid(Args[i]) || !IsUnique(Args[i]) {
			return (false)
		}
	}
	return (true)
}

func CreateSodoku(Args []string) [][]rune {
	var board [][]rune
	for i := 1; i < len(Args); i++ {
		board = append(board, []rune(Args[i]))
	}
	return (board)
}

func ValidCol(board [][]rune, col int, c rune) bool {
	for i := 0; i < 9; i++ {
		if c == board[i][col] {
			return (false)
		}
	}
	return (true)
}

func ValidRow(board [][]rune, row int, c rune) bool {
	for j := 0; j < 9; j++ {
		if c == board[row][j] {
			return (false)
		}
	}
	return (true)
}

func ValidSubGrid(board [][]rune, row int, col int, c rune) bool {
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
	if ValidRow(board, row, c) && ValidCol(board, col, c) && ValidSubGrid(board, row, col, c) {
		return (true)
	}
	return (false)
}

func SolveGame(board [][]rune, row, col int) bool {
	if row == 9 {
		return (true)
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
					return (true)
				}
				board[row][col] = '.'
			}
		}
	}
	return (false)
}

func PrintResults(board [][]rune, vertical, horizontal string, border bool) {
	if border {
		fmt.Println(horizontal)
	}
	drawer := 1
	ver := 1
	for i := 0; i < len(board); i++ {
		if border {
			fmt.Print(vertical)
			ver = 1
		}
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(string(board[i][j]))
			if border && ver%3 == 0 {
				fmt.Print(vertical)
			}
			if border {
				if j < len(board[i])-1 && ver%3 != 0 {
					fmt.Print(" ")
				}
			} else {
				if j < len(board[i])-1 {
					fmt.Print(" ")
				}
			}
			ver++
		}
		if border && drawer%3 == 0 {
			fmt.Print("\n" + horizontal)
		}
		fmt.Println()
		drawer++
	}
}

func main(){
	arguments := []string{"name",".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}
	if CheckArgs(arguments) {
		board := CreateSodoku(arguments)
		border := false
		// Ver │ ║
		vertical := "│"
		// Hor ════════════════════ / ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰
		horizontal := "☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰"
		if SolveGame(board, 0, 0) {
			PrintResults(board, vertical, horizontal, border)
		} else {
			fmt.Println("Error")
		}
	} else {
		fmt.Println("Error")
	}
}
