package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	template, rules, err := readLines("input-test.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(template, rules)
}

func readLines(path string) (string, map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "error", nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	readRules := false
	template := ""
	rules := make(map[string]string)

	for scanner.Scan() {
		scannerText := scanner.Text()
		if len(scannerText) == 0 {
			readRules = true
			continue
		}

		if !readRules {
			template = scannerText
		} else {
			pair := strings.Split(scannerText, " -> ")
			rules[pair[0]] = pair[1]
		}
	}
	return template, rules, err
}

func Part1(template string, rules map[string]string) {
	updatedPolymer := template
	for i := 0; i < 40; i++ {
		fmt.Println("Step", i+1)
		updatedPolymer = Step(updatedPolymer, rules)
		//fmt.Println(updatedPolymer)
	}
	letterCounts := SumElementCounts(updatedPolymer)
	fmt.Println("Difference between highest and lowest is", GetMaxElementCount(letterCounts)-GetMinElementCount(letterCounts))
}

func SumElementCounts(polymer string) map[string]int {
	rv := make(map[string]int)
	for _, r := range polymer {
		rv[string(r)]++
	}
	return rv
}

func GetMinElementCount(letterCount map[string]int) int {
	lowestSeen := 9999
	for _, v := range letterCount {
		if v < lowestSeen {
			lowestSeen = v
		}
	}
	return lowestSeen
}

func GetMaxElementCount(letterCount map[string]int) int {
	highestSeen := 0
	for _, v := range letterCount {
		if v > highestSeen {
			highestSeen = v
		}
	}
	return highestSeen
}

func JoinPolymerSlices(polymerSlices []string) string {
	rv := ""
	for i, s := range polymerSlices {
		if i == 0 {
			rv += s
		} else {
			rv += s[1:3]
		}
	}
	return rv
}

func Step(template string, rules map[string]string) string {
	stringSlices := make([]string, len(template)-1)
	for i := 0; i < len(stringSlices); i++ {
		stringSlices[i] = template[i : i+2]
	}
	for i := 0; i < len(stringSlices); i++ {
		for key, value := range rules {
			if stringSlices[i] == key {
				stringSlices[i] = string(stringSlices[i][0]) + value + string(stringSlices[i][1])
				break
			}
		}

	}
	return JoinPolymerSlices(stringSlices)
}
