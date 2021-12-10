package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part2()
}

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert to int
		i, err := strconv.Atoi(scanner.Text())

		if err != nil {
			return nil, err
		}

		lines = append(lines, i)
	}
	return lines, scanner.Err()
}

func part1() {
	inputSlice, err := readLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	prev, increaseCount := 0, 0

	for _, value := range inputSlice {
		if prev == 0 {
			prev = value
			continue
		}
		if value > prev {
			increaseCount++
		}
		prev = value
	}

	fmt.Println(increaseCount)
}

func part2() {
	inputSlice, err := readLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	prev, increaseCount := 0, 0

	for index := range inputSlice {
		slidingWindow := inputSlice[index : index+3]
		slidingSum := 0

		if slidingWindow[len(slidingWindow)-1] == 0 {
			continue
		}

		for _, slidingValue := range slidingWindow {
			slidingSum += slidingValue
		}

		if index > 0 && slidingSum > prev {
			increaseCount++
		}
		prev = slidingSum
	}

	fmt.Println(increaseCount)
}
