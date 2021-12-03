package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	report, err := readLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	part1(report)
	part2(report)
}

func readLines(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([][]string, 1000)
	for i := range lines {
		lines[i] = make([]string, 12)
	}

	scanner := bufio.NewScanner(file)
	currentLineIndex := 0

	for scanner.Scan() {
		reportLine := scanner.Text()

		for i, c := range reportLine {
			lines[currentLineIndex][i] = string(c)
		}
		currentLineIndex++
	}
	return lines, scanner.Err()
}

func bitwiseInvert(binaryString string) string {
	var rv string
	for _, c := range binaryString {
		switch string(c) {
		case "0":
			rv += "1"
		case "1":
			rv += "0"
		}
	}

	return rv
}

func convertBinaryStringToInt(binaryString string) (int64, error) {
	rv, err := strconv.ParseInt(binaryString, 2, 64)
	return rv, err
}

func getMostCommonBit(report [][]string, columnIndex int) string {
	colSum := 0

	for rowNum := 0; rowNum < len(report); rowNum++ {
		switch report[rowNum][columnIndex] {
		case "0":
			colSum--
		case "1":
			colSum++
		}
	}

	if colSum < 0 {
		return "0"
	}
	if colSum == 0 {
		return "1"
	}

	return "1"
}

func getLeastCommonBit(report [][]string, columnIndex int) string {
	colSum := 0

	for rowNum := 0; rowNum < len(report); rowNum++ {
		switch report[rowNum][columnIndex] {
		case "0":
			colSum--
		case "1":
			colSum++
		}
	}

	if colSum < 0 {
		return "1"
	}
	if colSum == 0 {
		return "0"
	}

	return "0"
}

func filterReport(report [][]string, filterChar string, filterIndex int) [][]string {
	var rv [][]string

	for _, line := range report {
		if line[filterIndex] == filterChar {
			rv = append(rv, line)
		}
	}

	return rv
}

func getOxygenRating(report [][]string, filterIndex int) int64 {
	if len(report) == 1 {
		rv, err := convertBinaryStringToInt(strings.Join(report[0], ""))

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		return rv
	}

	mostCommonBit := getMostCommonBit(report, filterIndex)
	filteredReport := filterReport(report, mostCommonBit, filterIndex)

	return getOxygenRating(filteredReport, filterIndex+1)
}

func getScrubberRating(report [][]string, filterIndex int) int64 {
	if len(report) == 1 {
		rv, err := convertBinaryStringToInt(strings.Join(report[0], ""))

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		return rv
	}

	leastCommonBit := getLeastCommonBit(report, filterIndex)
	filteredReport := filterReport(report, leastCommonBit, filterIndex)

	return getScrubberRating(filteredReport, filterIndex+1)
}

func part1(report [][]string) {
	var reportOutput string

	for colNum := 0; colNum < 12; colNum++ {
		mostCommon := getMostCommonBit(report, colNum)

		reportOutput += mostCommon
	}

	gammaRate, err := convertBinaryStringToInt(reportOutput)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	epsilonRate, err := convertBinaryStringToInt(bitwiseInvert(reportOutput))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println(gammaRate * epsilonRate)
}

func part2(report [][]string) {
	oxygenRating := getOxygenRating(report, 0)
	scrubberRating := getScrubberRating(report, 0)

	println(oxygenRating * scrubberRating)
}
