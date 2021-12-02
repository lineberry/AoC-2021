package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cast"
)

type instruction struct {
	direction string
	distance  int
}

func main() {
	part1()
	part2()
}

func readLines(path string) ([]instruction, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]instruction, 1000)
	scanner := bufio.NewScanner(file)
	currentLineIndex := 0

	for scanner.Scan() {
		//split on space
		lineSlice := strings.Fields(scanner.Text())

		if err != nil {
			return nil, err
		}

		lines[currentLineIndex] = instruction{lineSlice[0], cast.ToInt(lineSlice[1])}
		currentLineIndex++
	}
	return lines, scanner.Err()
}

func part1() {
	instructions, err := readLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	xPos, zPos := 0, 0

	for _, movement := range instructions {
		switch movement.direction {
		case "forward":
			xPos += movement.distance
		case "up":
			zPos -= movement.distance
		case "down":
			zPos += movement.distance
		}
	}

	fmt.Println(xPos * zPos)
}

func part2() {
	instructions, err := readLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	xPos, zPos, aim := 0, 0, 0

	for _, movement := range instructions {
		switch movement.direction {
		case "forward":
			xPos += movement.distance
			zPos += aim * movement.distance
		case "up":
			aim -= movement.distance
		case "down":
			aim += movement.distance
		}
	}

	fmt.Println(xPos * zPos)
}
