package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	crates := make([]string, 0)
	procedures := make([]string, 0)
	feedCratesInput := true

	for fileScanner.Scan() {
		input := fileScanner.Text()

		if feedCratesInput {
			if input == "" {
				feedCratesInput = false
				continue
			}
			crates = append(crates, input)
			continue
		}

		procedures = append(procedures, input)
	}

	// fmt.Println(crates)
	// fmt.Println(procedures)
	stacks := mapCratesToStack(crates)
	fmt.Println(stacks)

	for _, p := range procedures {
		executeProcedure(p, stacks)
	}

	fmt.Println(stacks)

}

func executeProcedure(procedure string, stack map[string][]string) {
	slicedProcedure := strings.Split(procedure, " ")

	quantity, err := strconv.Atoi(slicedProcedure[1])
	if err != nil {
		panic(err)
	}

	origin := slicedProcedure[3]
	destiny := slicedProcedure[5]

	for i := 1; i < quantity; i++ {
		c := stack[origin][len(stack[origin])-1]
		stack[origin] = stack[origin][:len(stack[origin])-1]

		stack[destiny] = append(stack[destiny], c)
	}

}

func parseCratesString(s string) []string {
	var slicedString = strings.Split(s, "")
	var chunkSize = 4
	var chunks [][]string
	for i := 0; i < len(slicedString); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slicedString capacity
		if end > len(slicedString) {
			end = len(slicedString)
		}

		chunks = append(chunks, slicedString[i:end])
	}

	var result []string

	for _, c := range chunks {
		if c[1] != " " {
			result = append(result, c[1])
		}
	}

	return result
}

func reverseList(list []string) []string {
	var reversedList []string
	for i := len(list) - 1; i >= 0; i-- {
		reversedList = append(reversedList, list[i])
	}
	return reversedList
}

func mapCratesToStack(crates []string) map[string][]string {
	cratesStacks := make(map[string][]string)

	newKey := ""
	for _, c := range crates[len(crates)-1] {
		if c == blankSpaceRune {
			if newKey != "" {
				cratesStacks[newKey] = make([]string, 0)
			}
			newKey = ""
			continue
		}
		newKey += string(c)
	}

	crates = crates[:len(crates)-1]

	for _, crate := range crates {
		for i, s := range parseCratesString(crate) {
			cratesStacks[strconv.Itoa(i+1)] = append(cratesStacks[strconv.Itoa(i+1)], s)
		}
		// fmt.Println(i, parseCratesString(crate))
	}

	for k, v := range cratesStacks {
		cratesStacks[k] = reverseList(v)
	}

	return cratesStacks
}
