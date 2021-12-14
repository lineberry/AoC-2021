package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cast"
)

type Coordinate struct {
	xPos, yPos int
}

func main() {
	octopusMatrix, err := readLines("input-test.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(octopusMatrix)
}

func readLines(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	octopusMatrix := make([][]int, 10)
	scanner := bufio.NewScanner(file)
	currentLineIndex := 0

	for scanner.Scan() {
		octopusMatrix[currentLineIndex] = make([]int, 10)
		lineText := scanner.Text()
		for i, r := range lineText {
			initialEnergy := cast.ToInt(string(r))
			octopusMatrix[currentLineIndex][i] = initialEnergy
		}
		currentLineIndex++
	}
	return octopusMatrix, scanner.Err()
}

func PrintOctopusMatrix(om [][]int) {
	for i := 0; i < len(om); i++ {
		fmt.Println(om[i])
	}
}

func GetCountOfOctopusesThatCanFlash(om [][]int, el int) int {
	rv := 0
	for y, row := range om {
		for x := 0; x < len(row); x++ {
			if om[y][x] >= el {
				rv++
			}
		}
	}
	return rv
}

func GetOctopusNeighbors(xOrigin int, yOrigin int) []Coordinate {
	var rv []Coordinate
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if (xOrigin+x >= 0 && xOrigin+x <= 9) && (yOrigin+y >= 0 && yOrigin+y <= 9) {
				rv = append(rv, Coordinate{xPos: xOrigin + x, yPos: yOrigin + y})
			}
		}
	}
	return rv
}

func Part1(om [][]int) {
	flashCount := 0
	energyFlashThreshold := 10

	for step := 0; step < 999; step++ {
		stepFlashCount := 0
		//Increase Energy By 1
		for y := 0; y < len(om); y++ {
			for x := 0; x < len(om[y]); x++ {
				om[y][x]++
			}
		}
		for { //For as long as there are octopuses to flash in the step
			if GetCountOfOctopusesThatCanFlash(om, energyFlashThreshold) == 0 {
				break
			}
			//Flash octopuses that are at > 9
			for y, row := range om {
				for x := 0; x < len(row); x++ {
					if om[y][x] >= energyFlashThreshold {
						flashCount++
						stepFlashCount++
						om[y][x] = 0
						neighborCoords := GetOctopusNeighbors(x, y)
						for _, nc := range neighborCoords {
							if om[nc.yPos][nc.xPos] != 0 {
								om[nc.yPos][nc.xPos]++
							}
						}
					}
				}
			}
		}
		if stepFlashCount == 100 {
			fmt.Println("Sync on step", step+1)
			break
		}
		//PrintOctopusMatrix(om)
		//fmt.Println()
	}

	fmt.Println("Total flashes:", flashCount)
}
