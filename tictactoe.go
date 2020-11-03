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
		if (i+1)%b.dimension == 0 {
			fmt.Printf(" %d \n", i)
		} else {
			fmt.Printf(" %d |", i)
		}
	}
}

func (b board) printBoard(m [9]int) {
	fmt.Printf(" %s | %s | %s \n", b.getSymbol(m[0]), b.getSymbol(m[1]), b.getSymbol(m[2]))
	fmt.Printf(" %s | %s | %s \n", b.getSymbol(m[3]), b.getSymbol(m[4]), b.getSymbol(m[5]))
	fmt.Printf(" %s | %s | %s \n", b.getSymbol(m[6]), b.getSymbol(m[7]), b.getSymbol(m[8]))
}

func (b board) getSymbol(i int) string {
	if i == 1 {
		return "X"
	}

	if i == -1 {
		return "O"
	}

	return "-"
}

func (b board) determineRow(i int) int {
	if i <= 2 {
		return 0
	} else if i <= 5 {
		return 1
	} else {
		return 2
	}
}

func (b board) determineColumn(i int) int {
	if i%3 == 0 {
		return 0
	} else if (i+1)%3 == 0 {
		return 2
	} else {
		return 1
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
	winner player
}

func (g *game) getWinner(squareChoice int) (player, bool){
	row := g.board.determineRow(squareChoice)
	g.rowScore[row] += g.currPlayer.marker

	column := g.board.determineColumn(squareChoice)
	g.columnScore[column] += g.currPlayer.marker

	if g.rowScore[row] == 3 {
		g.winner = g.player1
	}

	if g.rowScore[row] == -3 {
		g.winner = g.player2
	}

	if g.columnScore[column] == 3 {
		g.winner = g.player1
	}

	if g.columnScore[column] == -3 {
		g.winner = g.player2
	}

	return player{}, false
}

func (g *game) toggleCurrPlayer() {
	if g.currPlayer == g.player1 {
		g.currPlayer = g.player2
	} else {
		g.currPlayer = g.player1
	}
}

func main() {
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
	fmt.Println("To play, just type the number in the square you want to choose.")

	var moves [9]int
	var choice int

	for (player{}) == g.winner {
		fmt.Printf("%s's turn: ", g.currPlayer.name)
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Please enter an integer between 1 and 9.")
			break
		}

		if moves[choice] != 0 {
			fmt.Println("That square is already taken.")
			break
		}

		moves[choice] = g.currPlayer.marker
		g.board.printBoard(moves)
		g.getWinner(choice)
		g.toggleCurrPlayer()
	}

	if g.winner != (player{}) {
		fmt.Printf("%s is the winner!", g.winner.name)
	} else {
		fmt.Print("Doh! The game was a draw.")
	}
}
