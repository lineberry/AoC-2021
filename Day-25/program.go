package main

import (
	"bufio"
	"fmt"
	"os"
)

var seafloorHeight int
var seafloorWidth int

func main() {
	//seafloorHeight, seafloorWidth = 9, 10 //Test size
	seafloorHeight, seafloorWidth = 137, 139 //Real size

	seaFloor, err := readLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(seaFloor)
}

func readLines(path string) ([][]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	seaFloor := make([][]rune, seafloorHeight)
	rowCounter := 0
	for scanner.Scan() {
		row := scanner.Text()
		seaFloor[rowCounter] = make([]rune, seafloorWidth)
		for i, r := range row {
			seaFloor[rowCounter][i] = r
		}

		rowCounter++
	}
	return seaFloor, err
}

func Step(seaFloor *[][]rune) bool {
	didMoveEast := MoveEast(seaFloor)
	didMoveSouth := MoveSouth(seaFloor)

	return didMoveEast || didMoveSouth
}

func MoveEast(seaFloor *[][]rune) bool {
	didMove := false
	for y := 0; y < seafloorHeight; y++ {
		furthestWestValue := (*seaFloor)[y][0]
		for x := 0; x < seafloorWidth; x++ {
			if (*seaFloor)[y][x] == '>' {
				if x < seafloorWidth-1 {
					if (*seaFloor)[y][x+1] == '.' {
						(*seaFloor)[y][x] = '.'
						(*seaFloor)[y][x+1] = '>'
						x++
						didMove = true
					}
				} else if furthestWestValue == '.' {
					(*seaFloor)[y][0] = '>'
					(*seaFloor)[y][x] = '.'
					didMove = true
				}
			}
		}
	}
	return didMove
}

func MoveSouth(seaFloor *[][]rune) bool {
	didMove := false
	for x := 0; x < seafloorWidth; x++ {
		furthestNorthValue := (*seaFloor)[0][x]
		for y := 0; y < seafloorHeight; y++ {
			if (*seaFloor)[y][x] == 'v' {
				if y < seafloorHeight-1 {
					if (*seaFloor)[y+1][x] == '.' {
						(*seaFloor)[y][x] = '.'
						(*seaFloor)[y+1][x] = 'v'
						y++
						didMove = true
					}
				} else if furthestNorthValue == '.' {
					(*seaFloor)[0][x] = 'v'
					(*seaFloor)[y][x] = '.'
					didMove = true
				}
			}
		}
	}
	return didMove
}

func PrintSeaFloor(seaFloor *[][]rune) {
	for y := 0; y < seafloorHeight; y++ {
		for x := 0; x < seafloorWidth; x++ {
			fmt.Print(string((*seaFloor)[y][x]))
		}
		fmt.Println()
	}
}

func Part1(seaFloor [][]rune) {
	stepCount := 0
	for {
		didMove := Step(&seaFloor)
		stepCount++
		fmt.Println(stepCount)
		if !didMove {
			fmt.Println(stepCount)
			break
		}
	}
}
