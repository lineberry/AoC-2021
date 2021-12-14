package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	caveMap, err := readLines("input-med.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(caveMap)
}

func readLines(path string) (map[string][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	caveMap := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		startAndEnd := strings.Split(scanner.Text(), "-")
		start := startAndEnd[0]
		end := startAndEnd[1]

		_, startExists := caveMap[start]
		_, endExists := caveMap[end]
		if !startExists && end != "start" {
			caveMap[start] = []string{end}
		} else if end != "start" {
			caveMap[start] = append(caveMap[start], end)
		}

		if !endExists && start != "start" {
			caveMap[end] = []string{start}
		} else if start != "start" {
			caveMap[end] = append(caveMap[end], start)
		}
	}
	//Remove any entries where there is only one small destination
	for key := range caveMap {
		if len(caveMap[key]) == 1 && caveMap[key][0] == strings.ToLower(caveMap[key][0]) && caveMap[key][0] != "end" {
			delete(caveMap, key)
		}
	}
	return caveMap, scanner.Err()
}

func Part1(caveMap map[string][]string) {
	fmt.Println(caveMap)
}
