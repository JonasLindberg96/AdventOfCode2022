package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func day1Cal() {
	fmt.Println("Advent Of Code 2022, Day 1: Calorie Counting")
	readFile, err := os.Open("lib/kcal_day1.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	// Scan file, append combined calories per elf to map
	elvesMap := make(map[int]int)
	maxCalories := 0
	i := 1
	calorieCount := 0

	for fileScanner.Scan() {
		if len(fileScanner.Text()) != 0 {
			cal, _ := strconv.Atoi(fileScanner.Text())
			calorieCount = calorieCount + cal
			if calorieCount > maxCalories {
				//whichElf = i
				maxCalories = calorieCount
			}
		} else {
			elvesMap[i] = calorieCount
			i = i + 1
			calorieCount = 0
		}
	}
	readFile.Close()

	// Sort the map based on calories (value)
	keys := make([]int, 0, len(elvesMap))

	for key := range elvesMap {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return elvesMap[keys[i]] > elvesMap[keys[j]]
	})

	totalCal := 0
	for l := 0; l < 3; l++ {
		fmt.Println("Elf ", l, ":", elvesMap[keys[l]])
		totalCal = totalCal + elvesMap[keys[l]]
	}
	fmt.Println("Total cals of top three elves: ", totalCal)
}

func main() {
	day1Cal()
}
