package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cast"
)

type MapGrid struct {
	Grid [][]int
}

type Line struct {
	x1, y1 int
	x2, y2 int
}

func main() {
	mapGrid := MakeMapGrid()
	lines, err := readLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(lines, &mapGrid)
}

func readLines(path string) ([]Line, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]Line, 500)
	scanner := bufio.NewScanner(file)
	currentLineIndex := 0

	for scanner.Scan() {
		xAndYs := strings.Split(scanner.Text(), " -> ")
		x1y1 := strings.Split(xAndYs[0], ",")
		x2y2 := strings.Split(xAndYs[1], ",")

		lines[currentLineIndex] = Line{x1: cast.ToInt(x1y1[0]), y1: cast.ToInt(x1y1[1]), x2: cast.ToInt(x2y2[0]), y2: cast.ToInt(x2y2[1])}
		currentLineIndex++
	}
	return lines, scanner.Err()
}

func MakeMapGrid() MapGrid {
	var rv MapGrid
	emptyGrid := make([][]int, 1000)
	for i := 0; i < len(emptyGrid); i++ {
		emptyGrid[i] = make([]int, 1000)
	}
	rv.Grid = emptyGrid
	return rv
}

func DrawLine(line Line, mapGrid *MapGrid) {
	slope := GetSlope(line)

	//fmt.Println("Line ", line, " has a slope of ", slope)

	switch slope {
	case "-0":
		if line.y2 > line.y1 {
			for i := 0; i <= line.y2-line.y1; i++ {
				mapGrid.Grid[line.y1+i][line.x1] += 1
			}
		} else {
			for i := 0; i <= line.y1-line.y2; i++ {
				mapGrid.Grid[line.y2+i][line.x1] += 1
			}
		}

	case "0":
		if line.x2 > line.x1 {
			for i := 0; i <= line.x2-line.x1; i++ {
				mapGrid.Grid[line.y1][line.x1+i] += 1
			}
		} else {
			for i := 0; i <= line.x1-line.x2; i++ {
				mapGrid.Grid[line.y1][line.x2+i] += 1
			}
		}

	case "1":
		if line.x2 > line.x1 {
			for i := 0; i <= line.x2-line.x1; i++ {
				mapGrid.Grid[line.y1+i][line.x1+i] += 1
			}
		} else {
			for i := 0; i <= line.x1-line.x2; i++ {
				mapGrid.Grid[line.y2+i][line.x2+i] += 1
			}
		}
	case "-1":
		if line.x2 > line.x1 {
			for i := 0; i <= line.x2-line.x1; i++ {
				mapGrid.Grid[line.y1-i][line.x1+i] += 1
			}
		} else {
			for i := 0; i <= line.x1-line.x2; i++ {
				mapGrid.Grid[line.y2-i][line.x2+i] += 1
			}
		}
	}

	//PrintMapGrid(*mapGrid)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GetSlope(line Line) string {
	if line.x2-line.x1 == 0 { //vertical
		return "-0"
	} else if line.y2-line.y1 == 0 { //horizontal
		return "0"
	} else if (line.y2-line.y1)/(line.x2-line.x1) > 0 {
		return "1"
	} else {
		return "-1"
	}
}

func GetCountOfOverlaps(numToLookFor int, mapGrid *MapGrid) int {
	var rv int
	for i := 0; i < len(mapGrid.Grid); i++ {
		for j := 0; j < len(mapGrid.Grid[i]); j++ {
			if mapGrid.Grid[i][j] >= numToLookFor {
				rv++
			}
		}
	}
	return rv
}

func PrintMapGrid(mapGrid MapGrid) {
	for i := 0; i < len(mapGrid.Grid); i++ {
		fmt.Println(mapGrid.Grid[i])
	}
}

func Part1(lines []Line, mapGrid *MapGrid) {
	for _, line := range lines {
		DrawLine(line, mapGrid)
	}

	//PrintMapGrid(*mapGrid)

	part1Answer := GetCountOfOverlaps(2, mapGrid)
	fmt.Println(part1Answer)
}
