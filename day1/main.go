package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var elvesCalories []int

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		readFile.Close()
	}()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	currentElfCalories := 0

	for fileScanner.Scan() {
		elfCalorie := fileScanner.Text()
		if elfCalorie == "" {
			elvesCalories = append(elvesCalories, currentElfCalories)
			currentElfCalories = 0
			continue
		}

		currentElfCalories = addCalories(elfCalorie, currentElfCalories)
	}

	if currentElfCalories > 0 {
		elvesCalories = append(elvesCalories, currentElfCalories)
	}

	// part one solution end here
	// itemValue, itemPosition := findHighestItem(elvesCalories)
	// fmt.Printf("Elf %d has the highest calories count with total of: %d\n", itemPosition+1, itemValue)
	var topThreeCaloriesCount [3]int

	for i := 0; i < 3; i++ {
		itemValue, itemPosition := findHighestItem(elvesCalories)
		fmt.Printf("Calories count with total of: %d\n", itemValue)
		topThreeCaloriesCount[i] = itemValue
		elvesCalories = removeItemByIndex(elvesCalories, itemPosition)
	}

	topThreeSum := 0
	for _, v := range topThreeCaloriesCount {
		topThreeSum += v
	}
	fmt.Printf("Top three count sum: %d\n", topThreeSum)
}

func addCalories(calories string, totalCalories int) int {
	intCalories, err := strconv.Atoi(calories)
	if err != nil {
		panic(err)
	}
	return totalCalories + intCalories
}

func removeItemByIndex(list []int, index int) []int {
	before := list[:index]
	after := list[index+1:]
	before = append(before, after...)
	return before
}

func findHighestItem(list []int) (item, position int) {
	highestItem := 0
	for i := 1; i < len(list); i++ {
		if list[i] > list[highestItem] {
			highestItem = i
		}
	}
	return list[highestItem], highestItem
}
