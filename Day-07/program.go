package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cast"
)

func main() {
	crabPositions, err := ReadLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(crabPositions)
}

func ReadLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var positions []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert to int
		stringPositions := strings.Split(scanner.Text(), ",")
		for _, p := range stringPositions {
			positions = append(positions, cast.ToInt(p))
		}
	}
	return positions, scanner.Err()
}

func Part1(crabPositions []int) {
	lowestFuelUsage, lowestPosition := 9999999999999, 0
	possiblePositions := MakePossiblePositions()

	for _, pp := range possiblePositions {
		totalFuelUsed := 0
		for _, cp := range crabPositions {
			totalFuelUsed += CalcFuelUsagePartTwo(cp, pp)
		}
		if totalFuelUsed < lowestFuelUsage {
			lowestFuelUsage = totalFuelUsed
			lowestPosition = pp
		}
	}
	fmt.Println("Lowest fuel used is ", lowestFuelUsage, " when all crabs move to ", lowestPosition)
}

func MakePossiblePositions() []int {
	rv := make([]int, 1925)
	for i := 0; i <= 1924; i++ {
		rv[i] = i
	}
	return rv
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func CalcFuelUsage(crabPosition int, positionToMoveTo int) int {
	return Abs(crabPosition - positionToMoveTo)
}

func CalcFuelUsagePartTwo(crabPosition int, positionToMoveTo int) int {
	fuelUsed := 0
	positionMoves := Abs(crabPosition - positionToMoveTo)
	for i := 0; i <= positionMoves; i++ {
		fuelUsed += i
	}
	return fuelUsed
}
