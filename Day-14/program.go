package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var InjectionCache map[string]string

func main() {
	template, _, err := readLines("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(template)
	Part2(template)
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

	InjectionCache = make(map[string]string)
	for key, value := range rules {
		InjectionCache[key] = string(key[0]) + value + string(key[1])
	}

	return template, rules, err
}

func GetKeysFromMap(theMap map[string]int) []string {
	keys := make([]string, len(theMap))

	i := 0
	for k := range theMap {
		keys[i] = k
		i++
	}
	return keys
}

func Part1(template string) {
	updatedPolymer := template
	for i := 0; i < 10; i++ {
		fmt.Println("Step", i+1)
		updatedPolymer = Step2(updatedPolymer)
		//fmt.Println(updatedPolymer)
	}
	letterCounts := SumElementCounts(updatedPolymer)
	fmt.Println("Difference between highest and lowest is", GetMaxElementCount(letterCounts)-GetMinElementCount(letterCounts))
}

func Part2(template string) {
	pairCount := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]
		pairCount[pair] += 1
	}

	//fmt.Println(pairCount)
	for i := 0; i < 40; i++ {
		pairCount = Step3(pairCount)
		//fmt.Println(pairCount)
	}

	SumElementsFromPairCounts(pairCount, template)
}

func SumElementsFromPairCounts(pairMap map[string]int, template string) {
	elementSums := map[string]int{
		string(template[0]):               1,
		string(template[len(template)-1]): 1,
	}

	for key, value := range pairMap {
		elementSums[string(key[0])] += value
		elementSums[string(key[1])] += value
	}

	fmt.Println(elementSums)
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
	start := time.Now()
	var sb strings.Builder
	for i, s := range polymerSlices {
		if i == 0 {
			sb.WriteString(s)
		} else {
			sb.WriteString(s[1:3])
		}
	}

	end := time.Now()
	fmt.Println("Took", end.Sub(start), "to join slices")
	return sb.String()
}

func Step(template string) string {
	stringSlices := make([]string, len(template)-1)

	start := time.Now()

	for i := 0; i < len(stringSlices); i++ {
		stringSlices[i] = template[i : i+2]
	}

	t1 := time.Now()

	fmt.Println("Took ", t1.Sub(start), "to create stringslices")

	//fmt.Println(InjectionCache)

	for i := 0; i < len(stringSlices); i++ {
		stringSlices[i] = InjectionCache[stringSlices[i]]
	}

	t2 := time.Now()
	fmt.Println("Took ", t2.Sub(t1), "to inject rules")
	return JoinPolymerSlices(stringSlices)
}

func Step2(template string) string {
	//start := time.Now()
	var sb strings.Builder
	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]
		if i == 0 {
			sb.WriteString(InjectionCache[pair])
		} else {
			sb.WriteString(InjectionCache[pair][1:3])
		}

	}
	//end := time.Now()
	//fmt.Println("Took ", end.Sub(start), "to step.")
	return sb.String()
}

func Step3(pairMap map[string]int) map[string]int {
	newPairMap := make(map[string]int)
	for key := range pairMap {
		transformCount := pairMap[key]
		delete(pairMap, key)
		newPairMap[InjectionCache[key][0:2]] += transformCount
		newPairMap[InjectionCache[key][1:]] += transformCount
	}
	return newPairMap
}
