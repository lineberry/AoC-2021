package main

import "fmt"

type GameState struct {
	DieRoll       int
	DieRollCount  int
	Player1Space  int
	Player2Space  int
	TurnNumber    int
	Player1Score  int
	Player2Score  int
	IsPlayer1Turn bool
}

func main() {
	Part1(4, 9)
}

func Part1(player1Start int, player2Start int) {
	currentState := GameState{
		DieRoll:       0,
		DieRollCount:  0,
		Player1Space:  player1Start,
		Player2Space:  player2Start,
		TurnNumber:    1,
		Player1Score:  0,
		Player2Score:  0,
		IsPlayer1Turn: true,
	}

	for {
		if currentState.Player1Score >= 1000 || currentState.Player2Score >= 1000 {
			fmt.Println(GetLosingPlayerScore(&currentState) * currentState.DieRollCount)
			break
		}
		PlayerTurn(&currentState)
	}
}

func GetLosingPlayerScore(currentState *GameState) int {
	if currentState.Player1Score < 1000 {
		return currentState.Player1Score
	}
	return currentState.Player2Score
}

func GetValidDieRoll(attemptedRoll int) int {
	if attemptedRoll == 101 {
		return 1
	} else if attemptedRoll == 102 {
		return 2
	} else if attemptedRoll == 103 {
		return 3
	} else {
		return attemptedRoll
	}
}

func MovePlayer(dieRoll int, currentState *GameState) {
	if currentState.IsPlayer1Turn {
		if (currentState.Player1Space+dieRoll)%10 == 0 {
			currentState.Player1Space = 10
		} else {
			currentState.Player1Space = (currentState.Player1Space + dieRoll) % 10
		}
		currentState.Player1Score += currentState.Player1Space
	} else {
		if (currentState.Player2Space+dieRoll)%10 == 0 {
			currentState.Player2Space = 10
		} else {
			currentState.Player2Space = (currentState.Player2Space + dieRoll) % 10
		}
		currentState.Player2Score += currentState.Player2Space
	}
}

func RollDie(currentState *GameState) int {
	roll1 := GetValidDieRoll(currentState.DieRoll + 1)
	roll2 := GetValidDieRoll(currentState.DieRoll + 2)
	roll3 := GetValidDieRoll(currentState.DieRoll + 3)
	currentState.DieRoll = roll3
	currentState.DieRollCount += 3
	return roll1 + roll2 + roll3
}

func PlayerTurn(currentState *GameState) {
	spacesToMove := RollDie(currentState)
	MovePlayer(spacesToMove, currentState)
	currentState.IsPlayer1Turn = !currentState.IsPlayer1Turn
}
