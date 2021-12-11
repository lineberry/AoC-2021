package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cast"
)

type InAndOut struct {
	Input  string
	Output string
}

func main() {
	lines, err := ReadLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(lines)
	Part2(lines)
}

func ReadLines(path string) ([]InAndOut, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//Change this for real input
	lines := make([]InAndOut, 200)
	scanner := bufio.NewScanner(file)
	currentPosition := 0
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " | ")

		lines[currentPosition] = InAndOut{Input: splitLine[0], Output: splitLine[1]}

		currentPosition++
	}
	return lines, scanner.Err()
}

func Part1(lines []InAndOut) {
	uniqueDigitCount := 0

	for _, line := range lines {
		digits := strings.Fields(line.Output)
		for _, d := range digits {
			dLength := len(d)
			if dLength == 2 || dLength == 4 || dLength == 3 || dLength == 7 { //1, 4, 7 and 8
				uniqueDigitCount++
			}
		}
	}

	fmt.Println(uniqueDigitCount)
}

func GetLetterCounts(input string) map[rune]int {
	letterCount := map[rune]int{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
		'e': 0,
		'f': 0,
		'g': 0,
	}
	for _, letter := range input {
		if letter == ' ' {
			continue
		}
		letterCount[letter] += 1
	}
	return letterCount
}

func GetLetterDifference(longerString string, shorterString string) rune {
	for _, r := range longerString {
		if !strings.ContainsRune(shorterString, r) {
			return r
		}
	}
	return 'x'
}

func GetMiddleMapping(fourString string, wireMap map[string]rune) rune {
	for _, c := range fourString {
		if c == wireMap["top-left"] || c == wireMap["top-right"] || c == wireMap["bottom-right"] {
			continue
		}
		return c
	}
	return 'x'
}

func TrimLeadingZeros(input string) string {
	for i, r := range input {
		if r == '0' {
			continue
		} else {
			return input[i:4]
		}
	}
	return input
}

func Part2(lines []InAndOut) {
	totalSumOfOutputStrings := 0
	for _, line := range lines {
		wireMap := map[string]rune{
			"top":          'x',
			"top-left":     'x',
			"top-right":    'x',
			"middle":       'x',
			"bottom-left":  'x',
			"bottom-right": 'x',
			"bottom":       'x',
		}
		digits := strings.Fields(line.Input)
		var oneString string
		var fourString string
		var sevenString string
		for _, d := range digits {
			dLength := len(d)
			if dLength == 2 {
				oneString = d
			}
			if dLength == 4 {
				fourString = d
			}
			if dLength == 3 {
				sevenString = d
			}
		}

		wireMap["top"] = GetLetterDifference(sevenString, oneString)
		letterCount := GetLetterCounts(line.Input)
		var sevenLetters string

		for key, value := range letterCount {
			if value == 6 {
				wireMap["top-left"] = key
			}
			if value == 9 {
				wireMap["bottom-right"] = key
			}
			if value == 4 {
				wireMap["bottom-left"] = key
			}
			if value == 8 && key != wireMap["top"] {
				wireMap["top-right"] = key
			}
			if value == 7 {
				sevenLetters += string(key)
			}
		}
		middleMapping := GetMiddleMapping(fourString, wireMap)
		wireMap["middle"] = middleMapping
		wireMap["bottom"] = GetLetterDifference(sevenLetters, string(middleMapping))

		//fmt.Println(wireMap)

		outputStrings := strings.Fields(line.Output)
		var finalOutputForLine string
		for _, output := range outputStrings {
			if len(output) == 6 && strings.ContainsRune(output, wireMap["top-right"]) && strings.ContainsRune(output, wireMap["bottom-left"]) {
				finalOutputForLine += "0"
			}
			if len(output) == 2 {
				finalOutputForLine += "1"
			}
			if len(output) == 5 && !strings.ContainsRune(output, wireMap["bottom-right"]) {
				finalOutputForLine += "2"
			}
			if len(output) == 5 && strings.ContainsRune(output, wireMap["top-right"]) && strings.ContainsRune(output, wireMap["bottom-right"]) {
				finalOutputForLine += "3"
			}
			if len(output) == 4 {
				finalOutputForLine += "4"
			}
			if len(output) == 5 && strings.ContainsRune(output, wireMap["top-left"]) {
				finalOutputForLine += "5"
			}
			if len(output) == 6 && !strings.ContainsRune(output, wireMap["top-right"]) {
				finalOutputForLine += "6"
			}
			if len(output) == 3 {
				finalOutputForLine += "7"
			}
			if len(output) == 7 {
				finalOutputForLine += "8"
			}
			if len(output) == 6 && strings.ContainsRune(output, wireMap["top-right"]) && strings.ContainsRune(output, wireMap["middle"]) {
				finalOutputForLine += "9"
			}
		}
		trimmedNum := TrimLeadingZeros(finalOutputForLine)
		totalSumOfOutputStrings += cast.ToInt(trimmedNum)
	}
	fmt.Println(totalSumOfOutputStrings)
}
