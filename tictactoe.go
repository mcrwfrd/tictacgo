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

func (b board) printMap() {
	var spaces = b.dimension * b.dimension
	for i := 0; i < spaces; i++ {
		if ((i+1)%b.dimension == 0) {
			fmt.Printf(" %d \n", i)
		} else {
			fmt.Printf(" %d |", i)
		}
	}
}

type game struct {
	player1 player
	player2 player
	board board
	moves int
	rowScore []int
	columnScore []int
	currPlayer player
}

func (g *game) toggleCurrPlayer() {
	if (g.currPlayer == g.player1) {
		fmt.Println("here")
		g.currPlayer = g.player2
	} else {
		g.currPlayer = g.player1
	}
}

func main() {
	const dim = 3
	const numMoves = dim * dim

	g := new(game)
	g.player1 = player{name: "Player 1", marker: 1}
	g.player2 = player{name: "Player 2", marker: -1}
	g.board = board{dimension: 3}
	g.moves = g.board.dimension * g.board.dimension
	g.rowScore = []int{0, 0, 0}
	g.columnScore = []int{0, 0, 0}
	g.currPlayer = g.player1

	fmt.Println("The game is Tic Tac Toe. Here is the game board:")
	g.board.printMap()
	fmt.Println("Standard rules apply.")
	fmt.Println("In order to select a space, type the number corresponding to the space you to select according to the mock board above and press enter.")

	var moves [9]int
	var rows [3]int
	var columns [3]int
	var choice int

	var winner player
	for i := 0; i < g.moves; i++ {
		fmt.Print(g.currPlayer.name)
		_, err := fmt.Scanf("%d", &choice)

		if err != nil {
			fmt.Println("Please enter an integer between 1 and 9.")
		}

		moves[choice] = g.currPlayer.marker
		printBoard(moves)

		row := determineRow(choice)
		rows[row] += g.currPlayer.marker

		column := determineColumn(choice)
		columns[column] += g.currPlayer.marker

		if rows[row] == 3 {
			winner = g.player1
			break
		}

		if rows[row] == -3 {
			winner = g.player2
			break
		}

		if columns[column] == 3 {
			winner = g.player1
			break
		}

		if columns[column] == -3 {
			winner = g.player2
			break
		}

		g.toggleCurrPlayer()
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

