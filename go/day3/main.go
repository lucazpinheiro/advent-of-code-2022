package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"unicode"
)

type stringSet map[string]struct{}

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
	group := []string{}

	for fileScanner.Scan() {
		items := fileScanner.Text()
		group = append(group, items)
		if len(group) < 3 {
			continue
		}
		badge, err := findBadgeInGroup(group)
		if err != nil {
			panic(err)
		}

		if priority, ok := alphabet[badge]; ok {
			sumOfPriorities += priority
		}
		group = nil

	}

	fmt.Println(sumOfPriorities)
}

func findBadgeInGroup(group []string) (string, error) {

	firstElfItems := stringSet{}
	secondElfItems := stringSet{}
	thirdElfItems := stringSet{}

	for _, item := range group[0] {
		firstElfItems[string(item)] = struct{}{}
	}

	for _, item := range group[1] {
		secondElfItems[string(item)] = struct{}{}
	}

	for _, item := range group[2] {
		thirdElfItems[string(item)] = struct{}{}
	}

	sets := []stringSet{
		firstElfItems,
		secondElfItems,
		thirdElfItems,
	}

	sort.Slice(sets, func(i, j int) bool {
		return len(sets[i]) > len(sets[j])
	})

	commonItems := []string{}

	for firstSetKey, _ := range sets[0] {
		for secondSetKey, _ := range sets[1] {
			if firstSetKey == secondSetKey {
				commonItems = append(commonItems, firstSetKey)
			}
		}
	}

	for _, key := range commonItems {
		for thirdSetKey, _ := range sets[2] {
			if key == thirdSetKey {
				return key, nil
			}
		}
	}

	return "", errors.New("no badge found")
}
