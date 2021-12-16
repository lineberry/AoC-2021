package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cave struct {
	Name      string
	IsSmall   bool
	Neighbors map[string]*Cave
}

func NewCave(name string) *Cave {
	isSmall := strings.ToLower(name) == name

	return &Cave{
		Name:      name,
		IsSmall:   isSmall,
		Neighbors: map[string]*Cave{},
	}
}

type CaveMap struct {
	Caves map[string]*Cave
}

func NewCaveMap() *CaveMap {
	return &CaveMap{
		Caves: map[string]*Cave{},
	}
}

func (m *CaveMap) AddCave(caveName string) {
	cave := m.Caves[caveName]
	if cave == nil {
		c := NewCave(caveName)
		m.Caves[caveName] = c
	}
}

func (m *CaveMap) AddEdge(caveName1, caveName2 string) {
	cave1 := m.Caves[caveName1]
	cave2 := m.Caves[caveName2]

	if cave1 == nil || cave2 == nil {
		panic("not all caves exist")
	}

	if _, ok := cave1.Neighbors[caveName2]; ok {
		return
	}

	if caveName1 == "start" || caveName2 == "end" {
		cave1.Neighbors[caveName2] = cave2
	} else if caveName2 == "start" || caveName1 == "end" {
		cave2.Neighbors[caveName1] = cave1
	} else {
		cave1.Neighbors[caveName2] = cave2
		cave2.Neighbors[caveName1] = cave1
	}
}

func main() {
	caveMap, err := readLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(caveMap)
	Part2(caveMap)
}

func readLines(path string) (*CaveMap, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	caveMap := NewCaveMap()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		startAndEnd := strings.Split(scanner.Text(), "-")
		start := startAndEnd[0]
		end := startAndEnd[1]

		caveMap.AddCave(start)
		caveMap.AddCave(end)
		caveMap.AddEdge(start, end)
	}

	return caveMap, scanner.Err()
}

func Part1(caveMap *CaveMap) {
	isVisited := make(map[string]bool)
	pathList := []string{"start"}
	pathCount := new(int)
	PrintAllPaths("start", "end", caveMap, isVisited, pathList, pathCount)
	fmt.Println(*pathCount)
}

func Part2(caveMap *CaveMap) {
	visitCount := make(map[string]int)
	pathList := []string{"start"}
	pathCount := new(int)
	PrintAllPathsPartTwo("start", "end", caveMap, visitCount, pathList, false, pathCount)
	fmt.Println(*pathCount)
}

func PrintAllPaths(start string, end string, caveMap *CaveMap, isVisited map[string]bool, localPathList []string, pathCount *int) {
	if start == end {
		(*pathCount)++
		//fmt.Println(localPathList)
		//fmt.Println(*pathCount)
		//fmt.Println(isVisited)
		return
	}
	if caveMap.Caves[start].IsSmall {
		isVisited[start] = true
	}

	for neighborName := range caveMap.Caves[start].Neighbors {
		if !isVisited[neighborName] {
			localPathList = append(localPathList, neighborName)
			PrintAllPaths(neighborName, end, caveMap, isVisited, localPathList, pathCount)
			localPathList = localPathList[:len(localPathList)-1]
		}

	}

	if caveMap.Caves[start].IsSmall {
		isVisited[start] = false
	}
}

func PrintAllPathsPartTwo(start string, end string, caveMap *CaveMap, visitCount map[string]int, localPathList []string, isTriggered bool, pathCount *int) {
	if start == end {
		(*pathCount)++
		return
	}
	if caveMap.Caves[start].IsSmall {
		visitCount[start]++
		if visitCount[start] == 2 {
			isTriggered = true
		}
	}

	for neighborName := range caveMap.Caves[start].Neighbors {
		if !caveMap.Caves[neighborName].IsSmall || (visitCount[neighborName] < 2 && !isTriggered) || (visitCount[neighborName] < 1 && isTriggered) {
			localPathList = append(localPathList, neighborName)
			PrintAllPathsPartTwo(neighborName, end, caveMap, visitCount, localPathList, isTriggered, pathCount)
			localPathList = localPathList[:len(localPathList)-1]
		}

	}

	if caveMap.Caves[start].IsSmall {
		visitCount[start]--
	}
}
