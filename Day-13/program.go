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
	grid, coords, instructions, err := readLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(&grid, coords, instructions)
}

func readLines(path string) ([][]string, []Coordinate, []FoldInstruction, error) {
	//real values
	gridHeight := 895
	gridWidth := 1311
	//test values
	// gridHeight := 15
	// gridWidth := 11
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

func Fold(grid *[][]string, coords []Coordinate, instruction FoldInstruction) []Coordinate {
	var rv []Coordinate
	for _, c := range coords {
		switch instruction.foldAxis {
		case "x":
			if c.x > instruction.foldIndex {
				rv = append(rv, GetFoldedCoordinate(c, instruction))
			} else {
				rv = append(rv, c)
			}
		case "y":
			if c.y > instruction.foldIndex {
				rv = append(rv, GetFoldedCoordinate(c, instruction))
			} else {
				rv = append(rv, c)
			}
		}
	}
	return rv
}

func NewGridFromCoords(coords []Coordinate) *[][]string {
	maxX, maxY := 0, 0
	for _, c := range coords {
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}
	grid := make([][]string, maxY+1)
	for y := 0; y < len(grid); y++ {
		grid[y] = make([]string, maxX+1)
		for x := 0; x < len(grid[y]); x++ {
			grid[y][x] = " "
		}
	}
	for _, c := range coords {
		grid[c.y][c.x] = "█"
	}
	return &grid
}

func Part1(grid *[][]string, coords []Coordinate, instructions []FoldInstruction) {
	localCoords := coords
	localGrid := grid
	for index, i := range instructions {
		localCoords = Fold(grid, localCoords, i)
		localGrid = NewGridFromCoords(localCoords)
		if index == 0 {
			fmt.Println(GetCountOfDots(localGrid))
		}

	}
	PrintGrid(localGrid)
}
