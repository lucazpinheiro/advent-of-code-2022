package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	finalCratesOrder := ""
	cratesInput := make([]string, 0)
	parseProcedures := false
	crates := make(map[string][]string, 0)

	for fileScanner.Scan() {
		input := fileScanner.Text()
		if input == "" && !parseProcedures {
			crates = handleCrates(cratesInput)
			parseProcedures = true
			continue
		}
		cratesInput = append(cratesInput, input)
		if parseProcedures {
			executeProcedure(input, crates)
		}
	}

	fmt.Println(crates)
	expectedFinalCratesOrder := "CMZ"

	if finalCratesOrder == expectedFinalCratesOrder {
		fmt.Println("passing")
	}
}

func executeProcedure(rawProcedure string, crates map[string][]string) {
	procedure := strings.Split(rawProcedure, " ")
	quantity, err := strconv.Atoi(procedure[1])
	if err != nil {
		panic(err)
	}
	origin := procedure[3]
	destiny := procedure[5]

	for i := 1; i <= quantity; i++ {
		crateOrigin := crates[origin]
		lastElement := crateOrigin[len(crateOrigin)-1]
		crates[origin] = crateOrigin[:len(crateOrigin)-1] // pop

		crates[destiny] = append(crates[destiny], lastElement)
	}
}

func parseCratesIdentifiers(s string) []string {
	cratesIdentifiers := make([]string, 0)
	acc := ""
	for _, v := range s {
		if string(v) == " " {
			if acc != "" {
				cratesIdentifiers = append(cratesIdentifiers, acc)
				acc = ""
			}
			continue
		}
		acc += string(v)
	}
	return cratesIdentifiers
}

func handleCrates(cratesInput []string) map[string][]string {
	identifiersLine := cratesInput[len(cratesInput)-1]
	identifiers := parseCratesIdentifiers(identifiersLine)

	crates := make(map[string][]string, 0)

	for _, k := range identifiers {
		crates[k] = []string{}
	}

	for _, line := range cratesInput[:len(cratesInput)-1] {
		currentCrane := 1
		acc := 0
		for _, char := range line {
			if string(char) == " " {
				acc++
				continue
			}
			if string(char) == "[" || string(char) == "]" {
				acc++
				continue
			}
			if acc > 2 {
				currentCrane += 1
			}
			crate := crates[fmt.Sprint(currentCrane)]
			crate = append(crate, string(char))
			crates[fmt.Sprint(currentCrane)] = crate
			acc = 0
		}
	}

	return crates
}
