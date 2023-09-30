package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fmt.Println(solution(readFile))
}

func solution(file *os.File) int {
	var values []int

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	acc := 0
	for fileScanner.Scan() {
		strValue := fileScanner.Text()

		if strValue == "" {
			values = append(values, acc)
			acc = 0
			continue
		}

		value, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			panic(err)
		}

		acc += value
	}

	values = append(values, acc)

	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	return values[0] + values[1] + values[2]
}
