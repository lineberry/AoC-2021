package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	err := readLines("input-test.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1()
}

func readLines(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

	}
	return err
}

func Part1() {

}
