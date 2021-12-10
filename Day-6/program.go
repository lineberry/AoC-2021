package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cast"
)

func main() {
	ages, err := ReadLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(ages)
	Part2(ages)
}

func ReadLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ages []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert to int
		stringAges := strings.Split(scanner.Text(), ",")
		for _, a := range stringAges {
			ages = append(ages, cast.ToInt(a))
		}
	}
	return ages, scanner.Err()
}

func Part1(ages []int) {
	fishCount := Simulate(ages, 1, 80)
	fmt.Println(fishCount)
}

func Part2(ages []int) {
	fishCount := FastSimulate(ages, 256)
	fmt.Println(fishCount)
}

func GetLowestAge(ages []int) int {
	lowestSeen := 8
	for _, age := range ages {
		if age < lowestSeen {
			lowestSeen = age
		}
	}
	return lowestSeen
}

func Simulate(ages []int, currentDay int, targetDay int) int {
	var rv []int

	for i := 0; i < len(ages); i++ {
		if ages[i] == 0 {
			rv = append(rv, 6)
			rv = append(rv, 8)
		} else {
			rv = append(rv, ages[i]-1)
		}
	}

	if currentDay >= targetDay {
		return len(rv)
	}
	currentDay += 1
	return Simulate(rv, currentDay, targetDay)
}

func GetSumOfSlice(listOfInts []int) int {
	rv := 0
	for _, n := range listOfInts {
		rv += n
	}
	return rv
}

func FastSimulate(ages []int, targetDay int) int {
	fishCounts := make([]int, 9)
	for i := 0; i < len(ages); i++ {
		fishCounts[ages[i]] += 1
	}

	for i := 0; i < targetDay; i++ {
		originalSeven := fishCounts[7]
		originalZero := fishCounts[0]
		fishCounts[7] = fishCounts[8]
		fishCounts[8] = fishCounts[0] //Spawn new fish
		fishCounts[0] = fishCounts[1]
		fishCounts[1] = fishCounts[2]
		fishCounts[2] = fishCounts[3]
		fishCounts[3] = fishCounts[4]
		fishCounts[4] = fishCounts[5]
		fishCounts[5] = fishCounts[6]
		fishCounts[6] = originalSeven + originalZero
	}

	return GetSumOfSlice(fishCounts)
}
