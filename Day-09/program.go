package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cast"
)

type Coordinate struct {
	xPos int
	yPos int
}

func main() {
	depthReadings, err := ReadLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(depthReadings)
	Part2(depthReadings)
}

func ReadLines(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//Change this for real vs test input
	//matrixHeight, matrixWidth := 7, 12
	matrixHeight, matrixWidth := 102, 102

	depthReadings := make([][]int, matrixHeight)
	for i := 0; i < len(depthReadings); i++ {
		depthReadings[i] = make([]int, matrixWidth)
	}
	for i := 0; i < len(depthReadings[0]); i++ { //Set first and last lines to -1
		depthReadings[0][i] = 10 //so that we don't have to worry about matrix size
		depthReadings[matrixHeight-1][i] = 10
	}

	scanner := bufio.NewScanner(file)
	currentLine := 1
	for scanner.Scan() {
		depthReadings[currentLine][0] = 10             //Set first and last positions to -1
		depthReadings[currentLine][matrixWidth-1] = 10 //so that we don't have to worry about matrix size
		depthLineString := scanner.Text()
		for i := 0; i < len(depthLineString); i++ {
			depthReadings[currentLine][i+1] = cast.ToInt(string(depthLineString[i]))
		}
		currentLine++
	}
	return depthReadings, scanner.Err()
}

func PrintDepthMap(depthMap [][]int) {
	for i := 0; i < len(depthMap); i++ {
		fmt.Println(depthMap[i])
	}
}

func IsLowestPoint(depthReadings [][]int, xPos int, yPos int) bool {
	depth := depthReadings[yPos][xPos]
	xPlusOne := depthReadings[yPos][xPos+1]
	xMinusOne := depthReadings[yPos][xPos-1]
	yPlusOne := depthReadings[yPos+1][xPos]
	yMinusOne := depthReadings[yPos-1][xPos]

	return depth < xPlusOne && depth < xMinusOne && depth < yPlusOne && depth < yMinusOne
}

func GetRiskLevel(depthReading int) int {
	return depthReading + 1
}

func GetSumRiskLevel(depthReadings []int) int {
	riskSum := 0
	for i := 0; i < len(depthReadings); i++ {
		riskSum += GetRiskLevel(depthReadings[i])
	}
	return riskSum
}

func FindBasinSize(depthReadings [][]int, toBeCheckedCoords []Coordinate, alreadyCheckedCoords map[string]Coordinate) int {
	if len(toBeCheckedCoords) == 0 {
		return len(alreadyCheckedCoords)
	}

	var nextToBeCheckedCoords []Coordinate
	for i := 0; i < len(toBeCheckedCoords); i++ {
		initialCoord := toBeCheckedCoords[i]
		initialValue := depthReadings[initialCoord.yPos][initialCoord.xPos]
		upValue := depthReadings[initialCoord.yPos-1][initialCoord.xPos]
		rightValue := depthReadings[initialCoord.yPos][initialCoord.xPos+1]
		downValue := depthReadings[initialCoord.yPos+1][initialCoord.xPos]
		leftValue := depthReadings[initialCoord.yPos][initialCoord.xPos-1]

		_, upAlreadyExists := alreadyCheckedCoords[cast.ToString(initialCoord.xPos)+cast.ToString(initialCoord.yPos-1)]
		_, rightAlreadyExists := alreadyCheckedCoords[cast.ToString(initialCoord.xPos+1)+cast.ToString(initialCoord.yPos)]
		_, downAlreadyExists := alreadyCheckedCoords[cast.ToString(initialCoord.xPos)+cast.ToString(initialCoord.yPos+1)]
		_, leftAlreadyExists := alreadyCheckedCoords[cast.ToString(initialCoord.xPos-1)+cast.ToString(initialCoord.yPos)]

		//Check up
		if upValue > initialValue && upValue < 9 && !upAlreadyExists {
			nextToBeCheckedCoords = append(nextToBeCheckedCoords, Coordinate{xPos: initialCoord.xPos, yPos: initialCoord.yPos - 1})
		}
		//Check right
		if rightValue > initialValue && rightValue < 9 && !rightAlreadyExists {
			nextToBeCheckedCoords = append(nextToBeCheckedCoords, Coordinate{xPos: initialCoord.xPos + 1, yPos: initialCoord.yPos})
		}
		//Check down
		if downValue > initialValue && downValue < 9 && !downAlreadyExists {
			nextToBeCheckedCoords = append(nextToBeCheckedCoords, Coordinate{xPos: initialCoord.xPos, yPos: initialCoord.yPos + 1})
		}
		//Check left
		if leftValue > initialValue && leftValue < 9 && !leftAlreadyExists {
			nextToBeCheckedCoords = append(nextToBeCheckedCoords, Coordinate{xPos: initialCoord.xPos - 1, yPos: initialCoord.yPos})
		}
		alreadyCheckedCoords[cast.ToString(initialCoord.xPos)+cast.ToString(initialCoord.yPos)] = initialCoord

	}
	return FindBasinSize(depthReadings, nextToBeCheckedCoords, alreadyCheckedCoords)
}

func GetTop(listOfInts []int, topCount int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(listOfInts)))
	return listOfInts[0:topCount]
}

func GetProductOfSlice(listOfInts []int) int {
	rv := 1
	for _, n := range listOfInts {
		rv *= n
	}
	return rv
}

func Part1(depthReadings [][]int) {
	var lowestPoints []int

	for yPos := 1; yPos < len(depthReadings)-1; yPos++ {
		for xPos := 1; xPos < len(depthReadings[yPos])-1; xPos++ {
			if IsLowestPoint(depthReadings, xPos, yPos) {
				lowestPoints = append(lowestPoints, depthReadings[yPos][xPos])
			}
		}
	}
	fmt.Println("Risk Level:", GetSumRiskLevel(lowestPoints))
}

func Part2(depthReadings [][]int) {
	var lowestPoints []Coordinate

	for yPos := 1; yPos < len(depthReadings)-1; yPos++ {
		for xPos := 1; xPos < len(depthReadings[yPos])-1; xPos++ {
			if IsLowestPoint(depthReadings, xPos, yPos) {
				lowestPoints = append(lowestPoints, Coordinate{xPos: xPos, yPos: yPos})
			}
		}
	}

	var basinSizes []int
	for _, point := range lowestPoints {
		toBeCheckedCoords := []Coordinate{point}
		alreadyCheckedCoords := make(map[string]Coordinate)
		basinSize := FindBasinSize(depthReadings, toBeCheckedCoords, alreadyCheckedCoords)
		basinSizes = append(basinSizes, basinSize)
	}

	topThreeBasins := GetTop(basinSizes, 3)

	fmt.Println("Largest basin sizes:", GetProductOfSlice(topThreeBasins))
}
