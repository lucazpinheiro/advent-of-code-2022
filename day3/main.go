package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"unicode"
)

var alphabet = map[string]int{}

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

	priority := 1
	for r := 'a'; r < '{'; r++ {
		alphabet[string(r)] = priority
		alphabet[string(unicode.ToUpper(r))] = priority + 26
		priority++
	}

	sumOfPriorities := 0

	for fileScanner.Scan() {
		items := fileScanner.Text()
		duplicatedItem, err := findDuplicatedItem(items)
		if err != nil {
			panic(err)
		}

		if val, ok := alphabet[duplicatedItem]; ok {
			sumOfPriorities += val
		}

	}

	fmt.Println(sumOfPriorities)
}

func findDuplicatedItem(items string) (string, error) {
	splitPoint := len(items) / 2
	firstHalf := items[:splitPoint]
	secondHalf := items[splitPoint:]

	for _, firstHalfItem := range firstHalf {
		for _, secondHalfItem := range secondHalf {
			if firstHalfItem == secondHalfItem {
				return string(firstHalfItem), nil
			}
		}
	}

	return "", errors.New("No duplicated item")
}
