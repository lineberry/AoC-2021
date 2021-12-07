package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cast"
)

type BingoCard struct {
	Numbers    [][]int
	Marks      [][]bool
	AlreadyWon bool
}

func makeBingoCard() BingoCard {
	var rv BingoCard
	emptyNumbers := make([][]int, 5)
	emptyMarks := make([][]bool, 5)

	for i := range emptyMarks {
		emptyMarks[i] = []bool{false, false, false, false, false}
	}

	rv.Numbers = emptyNumbers
	rv.Marks = emptyMarks
	return rv
}

func markCard(cardPtr *BingoCard, calledNumber int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if cardPtr.Numbers[i][j] == calledNumber {
				cardPtr.Marks[i][j] = true
			}
		}
	}
}

func isCardWinner(cardPtr *BingoCard) bool {
	//Check rows
	for i := 0; i < 5; i++ {
		if cardPtr.Marks[i][0] && cardPtr.Marks[i][1] && cardPtr.Marks[i][2] && cardPtr.Marks[i][3] && cardPtr.Marks[i][4] {
			cardPtr.AlreadyWon = true
			return true
		}
	}

	//Check columns
	for i := 0; i < 5; i++ {
		if cardPtr.Marks[0][i] && cardPtr.Marks[1][i] && cardPtr.Marks[2][i] && cardPtr.Marks[3][i] && cardPtr.Marks[4][i] {
			cardPtr.AlreadyWon = true
			return true
		}
	}

	return false
}

func getSumUnmarked(card BingoCard) int {
	var sum int
	for i := range card.Marks {
		for j := range card.Marks[i] {
			if !card.Marks[i][j] {
				sum += card.Numbers[i][j]
			}
		}
	}
	return sum
}

func main() {
	numbers, cards, err := parseInput("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	gameLoop(numbers, cards)
}

func parseInput(path string) ([]int, []BingoCard, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numberCallList []int
	bingoCard := makeBingoCard()
	var bingoCardList []BingoCard
	currentLineIndex, bingoCardLineNum := 0, 0

	for scanner.Scan() {
		//Load number call list
		if currentLineIndex == 0 {
			numStrSlice := strings.Split(scanner.Text(), ",")
			numberCallList = make([]int, len(numStrSlice))

			for n := range numberCallList {
				numberCallList[n] = cast.ToInt(numStrSlice[n])
			}
			currentLineIndex++
		} else if currentLineIndex == 1 {
			currentLineIndex++
			continue //line 2 is blank
		} else {
			//Parse bingo cards here
			line := scanner.Text()

			if len(line) == 0 { //Blank line between cards.
				bingoCardLineNum = 0
				bingoCard = makeBingoCard()
				continue
			}

			lineSlice := strings.Fields(line)
			bingoRow := make([]int, 5)
			for i := range lineSlice {
				bingoRow[i] = cast.ToInt(lineSlice[i])
			}

			bingoCard.Numbers[bingoCardLineNum] = bingoRow

			if bingoCardLineNum == 4 {
				bingoCardList = append(bingoCardList, bingoCard)
			}
			bingoCardLineNum++
		}
	}

	return numberCallList, bingoCardList, scanner.Err()
}

func gameLoop(numberList []int, bingoCards []BingoCard) {
	//foundWinner := false

	for numIdx := 0; numIdx < len(numberList); numIdx++ {
		for cardIdx := 0; cardIdx < len(bingoCards); cardIdx++ {
			num := numberList[numIdx]
			card := &bingoCards[cardIdx]

			if card.AlreadyWon {
				continue
			}
			markCard(card, num)
			if isCardWinner(card) {
				fmt.Println("Found a winner when calling number: ", num)
				fmt.Println(num * getSumUnmarked(*card))
				//foundWinner = true
				//break
			}
		}
		// if foundWinner {
		// 	break
		// }
	}
}
