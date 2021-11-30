package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())

		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println(i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
