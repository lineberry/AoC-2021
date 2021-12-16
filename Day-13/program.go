package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cast"
)

type Coordinate struct {
	x, y int
}

type FoldInstruction struct {
	foldIndex int
	foldAxis  string
}

func main() {
	grid, coords, instructions, err := readLines("input-test.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(&grid, coords, instructions)
}

func readLines(path string) ([][]string, []Coordinate, []FoldInstruction, error) {
	gridHeight := 15
	gridWidth := 11
	readInstructions := false

	grid := make([][]string, gridHeight)
	var foldInstructions []FoldInstruction
	var coords []Coordinate
	for y := 0; y < len(grid); y++ {
		grid[y] = make([]string, gridWidth)
		for x := 0; x < gridWidth; x++ {
			grid[y][x] = " "
		}
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		scannerText := scanner.Text()
		if len(scannerText) == 0 {
			readInstructions = true
			continue
		}

		if !readInstructions {
			xAndY := strings.Split(scannerText, ",")
			x := cast.ToInt(xAndY[0])
			y := cast.ToInt(xAndY[1])
			coords = append(coords, Coordinate{x: x, y: y})
			grid[y][x] = "█"
		} else {
			foldFields := strings.Fields(scannerText)
			axisAndIndex := strings.Split(foldFields[2], "=")
			foldAxis := axisAndIndex[0]
			foldIndex := cast.ToInt(axisAndIndex[1])
			foldInstructions = append(foldInstructions, FoldInstruction{foldAxis: foldAxis, foldIndex: foldIndex})
		}
	}

	return grid, coords, foldInstructions, scanner.Err()
}

func GetCountOfDots(grid *[][]string) int {
	derefGrid := *grid
	count := 0
	for y := 0; y < len(derefGrid); y++ {
		for x := 0; x < len(derefGrid[y]); x++ {
			if derefGrid[y][x] == "█" {
				count++
			}
		}
	}
	return count
}

func PrintGrid(grid *[][]string) {
	for y := 0; y < len(*grid); y++ {
		fmt.Println((*grid)[y])
	}
}

func GetFoldedCoordinate(coord Coordinate, instruction FoldInstruction) Coordinate {
	yNew := instruction.foldIndex - (coord.y - instruction.foldIndex)
	xNew := instruction.foldIndex - (coord.x - instruction.foldIndex)
	rv := Coordinate{x: xNew, y: yNew}
	switch instruction.foldAxis {
	case "y":
		rv.x = coord.x
	case "x":
		rv.y = coord.y
	}
	return rv
}

func Fold(grid *[][]string, coords []Coordinate, instruction FoldInstruction) {
	fmt.Println("Fold")
}

func Part1(grid *[][]string, coords []Coordinate, instructions []FoldInstruction) {
	PrintGrid(grid)
	for _, i := range instructions {
		Fold(grid, coords, i)
	}
	fmt.Println(GetCountOfDots(grid))
}
