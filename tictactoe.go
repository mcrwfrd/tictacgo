package main

import (
	"fmt"
)

type player struct {
	name   string
	marker int
}

type board struct {
	dimension int
}

func main() {
	fmt.Println("The game is Tic Tac Toe. Here is the game board:")
	printMockBoard()
	fmt.Println("Standard rules apply.")
	fmt.Println("In order to select a space, type the number corresponding to the space you to select according to the mock board above and press enter.")

	const dim = 3
	const numMoves = dim * dim
	board := board{dimension: dim}

	fmt.Println(board.dimension)
	fmt.Println(numMoves)
	var moves [numMoves]int
	var rows [dim]int
	var columns [dim]int
	var choice int

	player1 := player{name: "Player 1", marker: 1}
	player2 := player{name: "Player 2", marker: -1}

	currPlayer := player1
	var winner player
	for i := 0; i < 9; i++ {
		fmt.Print(currPlayer.name)
		_, err := fmt.Scanf("%d", &choice)

		if err != nil {
			fmt.Println("Please enter an integer between 1 and 9.")
		}

		moves[choice] = currPlayer.marker
		printBoard(moves)

		row := determineRow(choice)
		rows[row] += currPlayer.marker

		column := determineColumn(choice)
		columns[column] += currPlayer.marker

		if rows[row] == 3 {
			winner = player1
			break
		}

		if rows[row] == -3 {
			winner = player2
			break
		}

		if columns[column] == 3 {
			winner = player1
			break
		}

		if columns[column] == -3 {
			winner = player2
			break
		}

		if currPlayer == player1 {
			currPlayer = player2
		} else {
			currPlayer = player1
		}
	}

	if winner.name != "" {
		fmt.Println(winner.name)
	} else {
		fmt.Println("draw")
	}
}

func determineRow(i int) int {
	if i <= 2 {
		return 0
	} else if i <= 5 {
		return 1
	} else {
		return 2
	}
}

func determineColumn(i int) int {
	if i%3 == 0 {
		return 0
	} else if (i+1)%3 == 0 {
		return 2
	} else {
		return 1
	}
}

func printBoard(m [9]int) {
	fmt.Printf(" %d | %d | %d \n", m[0], m[1], m[2])
	fmt.Printf(" %d | %d | %d \n", m[3], m[4], m[5])
	fmt.Printf(" %d | %d | %d \n", m[6], m[7], m[8])
}

func printMockBoard() {
	fmt.Println(" 0 | 1 | 2 ")
	fmt.Println(" 3 | 4 | 5 ")
	fmt.Println(" 6 | 7 | 8 ")
}
