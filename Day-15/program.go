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

	Part1()
}

func readLines(path string) (, error) {
	file, err := os.Open(path)
	if err != nil {
		return "error", nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		
	}
	return template, rules, err
}

func Part1() {

}