package main

import (
	"fmt"
)

func main() {
	fmt.Println("The game is Tic Tac Toe. Here is the game board:")
	printMockBoard()
	fmt.Println("Standard rules apply.")
	fmt.Println("In order to select a space, type the number corresponding to the space you to select according to the mock board above and press enter.")

	var moves [9]int
	var rows [3]int
	var columns [3]int
	var choice int
	currPlayer := 1
	winner := 0
	for i := 0; i < 9; i++ {
		if currPlayer == 1 {
			fmt.Print("Player 1 Move: ")
		} else {
			fmt.Print("Player 2 Move: ")
		}
		_, err := fmt.Scanf("%d", &choice)

		if err != nil {
			fmt.Println("Please enter an integer between 1 and 9.")
		}

		moves[choice] = currPlayer
		printBoard(moves)

		row := determineRow(choice)
		rows[row] += currPlayer

		column := determineColumn(choice)
		columns[column] += currPlayer

		if rows[row] == 3 {
			winner = 1
			break
		}

		if rows[row] == -3 {
			winner = -1
			break
		}

		if columns[column] == 3 {
			winner = 1
			break
		}

		if columns[column] == -3 {
			winner = -1
			break
		}

		if currPlayer == 1 {
			currPlayer = -1
		} else {
			currPlayer = 1
		}
	}

	if winner == 1 {
		fmt.Println("Player 1 Wins!")
	} else if winner == -1 {
		fmt.Println("Player 2 Wins!")
	} else {
		fmt.Println("Uh Oh, It's a Draw!")
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
