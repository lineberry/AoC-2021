package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cast"
)

type RebootStep struct {
	TurnOn bool
	xMin   int
	xMax   int
	yMin   int
	yMax   int
	zMin   int
	zMax   int
}

func main() {
	rebootSteps, err := readLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(rebootSteps)
}

func readLines(path string) ([]RebootStep, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var rebootSteps []RebootStep

	for scanner.Scan() {
		lineFields := strings.Fields(scanner.Text())
		instruction := lineFields[0]
		coordinates := lineFields[1]
		splitCoords := strings.Split(coordinates, ",")
		xCoords := strings.Split(splitCoords[0][2:], "..")
		yCoords := strings.Split(splitCoords[1][2:], "..")
		zCoords := strings.Split(splitCoords[2][2:], "..")

		turnOn := instruction == "on"

		xMin := cast.ToInt(xCoords[0]) + 50
		xMax := cast.ToInt(xCoords[1]) + 50
		yMin := cast.ToInt(yCoords[0]) + 50
		yMax := cast.ToInt(yCoords[1]) + 50
		zMin := cast.ToInt(zCoords[0]) + 50
		zMax := cast.ToInt(zCoords[1]) + 50

		if xMin > xMax {
			originalXMin := xMin
			xMin = xMax
			xMax = originalXMin
		}

		if yMin > yMax {
			originalYMin := yMin
			yMin = yMax
			yMax = originalYMin
		}

		if zMin > zMax {
			originalZMin := zMin
			zMin = zMax
			zMax = originalZMin
		}

		if xMin < 0 || xMax > 100 || yMin < 0 || yMax > 100 || zMin < 0 || zMax > 100 {
			continue
		}

		rebootSteps = append(rebootSteps, RebootStep{TurnOn: turnOn,
			xMin: xMin,
			xMax: xMax,
			yMin: yMin,
			yMax: yMax,
			zMin: zMin,
			zMax: zMax,
		})
	}
	return rebootSteps, err
}

func GetCubeOnCount(reactorCore *[][][]bool) int {
	var rv int

	for z := 0; z < len(*reactorCore); z++ {
		for y := 0; y < len((*reactorCore)[z]); y++ {
			for x := 0; x < len((*reactorCore)[z][y]); x++ {
				if (*reactorCore)[z][y][x] {
					rv++
				}
			}
		}
	}

	return rv
}

func Part1(rebootSteps []RebootStep) {
	reactorCore := make([][][]bool, 101)
	for z := 0; z < 101; z++ {
		reactorCore[z] = make([][]bool, 101)
		for y := 0; y < 101; y++ {
			reactorCore[z][y] = make([]bool, 101)
		}
	}

	for _, step := range rebootSteps {
		for z := step.zMin; z <= step.zMax; z++ {
			for y := step.yMin; y <= step.yMax; y++ {
				for x := step.xMin; x <= step.xMax; x++ {
					reactorCore[z][y][x] = step.TurnOn
				}
			}
		}
	}

	fmt.Println(GetCubeOnCount(&reactorCore))
}
