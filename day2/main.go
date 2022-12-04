package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	rockValue     = 1
	paperValue    = 2
	scissorsValue = 3
	winPoints     = 6
	drawPoints    = 3
	losePoints    = 0
)

var possibleMoves = map[string]string{
	"A": "rock",
	"X": "rock",
	"B": "paper",
	"Y": "paper",
	"C": "scissors",
	"Z": "scissors",
}

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		readFile.Close()
	}()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	totalPoints := 0

	for fileScanner.Scan() {
		round := fileScanner.Text()
		moves := strings.Split(round, " ")
		opponentMove := moves[0]
		myMove := moves[1]

		if possibleMoves[myMove] == possibleMoves[opponentMove] {
			totalPoints += pointsByTurn("draw", myMove)
			continue
		} else if possibleMoves[myMove] == "rock" && possibleMoves[opponentMove] == "scissors" {
			totalPoints += pointsByTurn("win", myMove)
			continue
		} else if possibleMoves[myMove] == "paper" && possibleMoves[opponentMove] == "rock" {
			totalPoints += pointsByTurn("win", myMove)
			continue
		} else if possibleMoves[myMove] == "scissors" && possibleMoves[opponentMove] == "paper" {
			totalPoints += pointsByTurn("win", myMove)
			continue
		} else {
			totalPoints += pointsByTurn("loss", myMove)
			continue
		}
	}

	fmt.Println("Total points", totalPoints)
}

func pointsByMove(move string) int {
	points := 0
	switch move {
	case "A", "X":
		points = rockValue
	case "B", "Y":
		points = paperValue
	case "C", "Z":
		points = scissorsValue
	}

	if points == 0 {
		panic(fmt.Sprint("Invalid move", move))
	}

	return points
}

func pointsByTurn(result string, move string) int {
	points := 0

	switch result {
	case "win":
		points += winPoints
	case "draw":
		points += drawPoints
	case "loss":
		points += losePoints
	default:
		panic(fmt.Sprint("Invalid result", result))
	}

	points += pointsByMove(move)
	return points
}

// func removeItemByIndex(list []int, index int) []int {
// 	before := list[:index]
// 	after := list[index+1:]
// 	before = append(before, after...)
// 	return before
// }

// func findHighestItem(list []int) (item, position int) {
// 	highestItem := 0
// 	for i := 1; i < len(list); i++ {
// 		if list[i] > list[highestItem] {
// 			highestItem = i
// 		}
// 	}
// 	return list[highestItem], highestItem
// }
