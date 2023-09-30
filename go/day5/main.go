package main

import (
	"bufio"
	"fmt"
	"os"
)

const blankSpaceRune = 32

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	cargo := make([]string, 0)
	operations := make([]string, 0)

	readingCargo := true

	for fileScanner.Scan() {
		input := fileScanner.Text()
		fmt.Println(input)
		if input == "" {
			readingCargo = false
			continue
		}
		if readingCargo {
			cargo = append(cargo, input)
			parseCargoLine(input)
		} else {
			operations = append(operations, input)
		}
	}

	// fmt.Println(cargo)
	// fmt.Println(operations)
}

func parseCargoLine(s string) {
	fmt.Println([]string{s})
}
