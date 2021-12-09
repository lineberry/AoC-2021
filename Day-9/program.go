package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	outputStrings, err := ReadLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(outputStrings)
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	outputStrings := make([]string, 200)
	scanner := bufio.NewScanner(file)
	currentPosition := 0
	for scanner.Scan() {
		// Convert to int
		outputString := strings.Split(scanner.Text(), " | ")[1]

		outputStrings[currentPosition] = outputString

		currentPosition++
	}
	return outputStrings, scanner.Err()
}

func Part1(outputStrings []string) {
	uniqueDigitCount := 0

	for _, outputString := range outputStrings {
		digits := strings.Fields(outputString)
		for _, d := range digits {
			dLength := len(d)
			if dLength == 2 || dLength == 4 || dLength == 3 || dLength == 7 { //1, 4, 7 and 8
				uniqueDigitCount++
			}
		}
	}

	fmt.Println(uniqueDigitCount)
}
