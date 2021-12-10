package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Queue []string

func (q *Queue) Push(x string) {
	*q = append(*q, x)
}

func (q *Queue) Pop() string {
	h := *q
	var el string
	l := len(h)
	el, *q = h[l-1], h[0:l-1]
	return el
}

func NewQueue() *Queue {
	return &Queue{}
}

func main() {
	lines, err := readLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	part1(lines)
	part2(lines)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//Change this for real input
	lines := make([]string, 90)
	scanner := bufio.NewScanner(file)
	currentLineIndex := 0

	for scanner.Scan() {
		lines[currentLineIndex] = scanner.Text()
		currentLineIndex++
	}
	return lines, scanner.Err()
}

func getCorruptionPoints(received string) int {
	rv := 0
	switch received {
	case ")":
		rv = 3
	case "]":
		rv = 57
	case "}":
		rv = 1197
	case ">":
		rv = 25137
	}
	return rv
}

func doesOpenMatchClose(open string, close string) bool {
	if (open == "(" && close == ")") || (open == "[" && close == "]") || (open == "{" && close == "}") || (open == "<" && close == ">") {
		return true
	}
	return false
}

func part1(lines []string) {
	totalPoints := 0
	for _, line := range lines {
		totalPoints += getCorruptionPointsForLine(line)
	}
	fmt.Println(totalPoints)
}

func getCorruptionPointsForLine(line string) int {
	q := NewQueue()
	for _, c := range line {
		readString := string(c)
		if readString == "(" || readString == "[" || readString == "{" || readString == "<" {
			//fmt.Println("Pushing", charString)
			q.Push(readString)
		} else {
			poppedString := q.Pop()
			//fmt.Println("Popped", poppedString)
			if !doesOpenMatchClose(poppedString, readString) {
				return getCorruptionPoints(readString)
			}
		}
	}
	return 0
}

func getMiddleScore(scores []int) int {
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func getIncompleteLineScore(completionString []string) int {
	totalScore := 0
	for _, char := range completionString {
		charPoints := 0
		charString := char

		switch charString {
		case ")":
			charPoints = 1
		case "]":
			charPoints = 2
		case "}":
			charPoints = 3
		case ">":
			charPoints = 4
		}

		totalScore = totalScore*5 + charPoints
	}
	return totalScore
}

func part2(lines []string) {
	var lineScores []int
	for _, line := range lines {
		if getCorruptionPointsForLine(line) > 0 {
			continue
		}
		q := NewQueue()
		for _, c := range line {
			readString := string(c)
			if readString == "(" || readString == "[" || readString == "{" || readString == "<" {
				q.Push(readString)
			} else {
				q.Pop()
			}
		}
		//fmt.Println("Remaining open chars:", *q)
		completeQ := NewQueue()
		originalQLen := len(*q)
		for i := 0; i < originalQLen; i++ {
			switch q.Pop() {
			case "(":
				completeQ.Push(")")
			case "[":
				completeQ.Push("]")
			case "{":
				completeQ.Push("}")
			case "<":
				completeQ.Push(">")
			}
		}
		lineScores = append(lineScores, getIncompleteLineScore(*completeQ))
	}
	fmt.Println(getMiddleScore(lineScores))
}
