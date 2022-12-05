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

	defer func() {
		readFile.Close()
	}()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	overlappingSections := 0

	for fileScanner.Scan() {
		sections := fileScanner.Text()
		if areOverlapping(sections) {
			fmt.Println("Overlapping sections:", sections)
			overlappingSections++
		}
	}

	fmt.Println("Total number of overlapping sections", overlappingSections)
}

func areOverlapping(pairedSections string) bool {
	sections := splitSections(pairedSections)

	firstPairOpen := sections[0]
	firstPairClose := sections[1]
	secondPairOpen := sections[2]
	secondPairClose := sections[3]

	if firstPairOpen <= secondPairOpen {
		return true
	}

	if firstPairClose >= secondPairClose {
		return true
	}

	if secondPairOpen <= firstPairOpen {
		return true
	}

	if secondPairClose >= firstPairClose {
		return true
	}

	return false
}

func splitSections(pairedSections string) []int {
	splitedSections := strings.FieldsFunc(pairedSections, func(r rune) bool {
		return r == '-' || r == ','
	})
	sections := make([]int, 0)
	for _, char := range splitedSections {
		intValue, err := strconv.Atoi(string(char))
		if err != nil {
			panic("not numeric value in sections")
		}
		sections = append(sections, intValue)
	}
	if len(sections) == 0 {
		panic("no sections")
	}
	return sections
}
