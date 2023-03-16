package main

import (
	"fmt"
	"math/rand"
	"os"
)

var board [3][3]string

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = "-"
		}
	}

	printBoard()

	for {
		var row, col int
		fmt.Print("Enter row (0-2): ")
		fmt.Scanln(&row)
		fmt.Print("Enter column (0-2): ")
		fmt.Scanln(&col)

		if row < 0 || row > 2 || col < 0 || col > 2 {
			fmt.Println("Invalid move, please enter a row and column between 0 and 2.")
			continue
		}
		if board[row][col] != "-" {
			fmt.Println("Invalid move, that cell is already occupied.")
			continue
		}

		board[row][col] = "X"
		printBoard()

		if checkWin("X") {
			fmt.Println("Congratulations! You win!")
			os.Exit(0)
		}

		if checkFull() {
			fmt.Println("It's a tie!")
			os.Exit(0)
		}

		var compRow, compCol int
		for {
			compRow = randInt(0, 2)
			compCol = randInt(0, 2)
			if board[compRow][compCol] == "-" {
				break
			}
		}

		board[compRow][compCol] = "O"
		fmt.Println("Computer's move:")
		printBoard()

		if checkWin("O") {
			fmt.Println("Sorry, you lose!")
			os.Exit(0)
		}
		if checkFull() {
			fmt.Println("It's a tie!")
			os.Exit(0)
		}
	}
}

func printBoard() {
	fmt.Println("   0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s ", board[i][j])
		}
		fmt.Println()
	}
}

func checkWin(player string) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}

func checkFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "-" {
				return false
			}
		}
	}
	return true
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}
