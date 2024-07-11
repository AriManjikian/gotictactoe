package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	board    [9]string
	player   string
	turn     int
	gameOver bool
	winner   string
}

func PrintBoard(game *Game) {
	//flush terminal with ansii escape codes
	fmt.Print("\033[H\033[2J")

	fmt.Printf(" %s|%s|%s\n", game.board[0], game.board[1], game.board[2])
	fmt.Printf("-------\n")
	fmt.Printf(" %s|%s|%s\n", game.board[3], game.board[4], game.board[5])
	fmt.Printf("-------\n")
	fmt.Printf(" %s|%s|%s\n", game.board[6], game.board[7], game.board[8])
}

func InitializeGame(game *Game) {
	game.board = [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
	game.player = "X"
	game.turn = 0
	game.gameOver = false
	game.winner = ""
}

func PlaceMove(game *Game, position int) {

	game.board[position-1] = game.player
	game.turn++
}

func SwitchPlayer(game *Game) {
	if game.player == "X" {
		game.player = "O"
		return
	}

	if game.player == "O" {
		game.player = "X"
		return
	}
}

func AskMove(game *Game) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s's turn. Please provide a position (1-9): ", game.player)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		position, err := strconv.Atoi(input)
		if err != nil || position < 1 || position > 10 || game.board[position-1] != " " {

			fmt.Println("Invalid input, please enter a valid move.")
			continue
		}
		return position
	}
}

func CheckWinner(game *Game) {

	winningCombos := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6},
	}

	for _, combo := range winningCombos {
		if game.board[combo[0]] == game.player && game.board[combo[1]] == game.player && game.board[combo[2]] == game.player {
			game.gameOver = true
			game.winner = game.player
			return
		}
	}
	if game.turn == 9 {
		game.gameOver = true
		return
	}
}

func main() {
	var game Game
	InitializeGame(&game)
	for !game.gameOver {
		PrintBoard(&game)
		position := AskMove(&game)
		PlaceMove(&game, position)
		CheckWinner(&game)
		SwitchPlayer(&game)
	}
	PrintBoard(&game)
	if game.winner == "" {
		fmt.Print("The game is a draw.")
	} else {
		fmt.Printf("%s is the winner.", game.winner)
	}
}
