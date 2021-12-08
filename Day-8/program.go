package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cast"
)

func main() {
	ages, err := ReadLines("input-test.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(ages)
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
		stringAges := strings.Split(scanner.Text(), ",")
		for _, a := range stringAges {
			positions = append(positions, cast.ToInt(a))
		}
	}
	return positions, scanner.Err()
}

func Part1(ages []int) {
	//fmt.Println("Initial state: ", ages)
	fishCount := Simulate(ages, 1, 80)
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
	// lowestAge := GetLowestAge(ages)
	// stepSize := 1
	// if lowestAge > 1 {
	// 	stepSize = lowestAge
	// }

	for i := 0; i < len(ages); i++ {
		if ages[i] == 0 {
			rv = append(rv, 6)
			rv = append(rv, 8)
		} else {
			rv = append(rv, ages[i]-1)
		}
	}
	//lowestAge = GetLowestAge(rv)
	//fmt.Println("Day ", currentDay, rv, "Lowest Age:", lowestAge)
	if currentDay >= targetDay {
		return len(rv)
	}
	currentDay += 1
	fmt.Println(currentDay)
	return Simulate(rv, currentDay, targetDay)
}
