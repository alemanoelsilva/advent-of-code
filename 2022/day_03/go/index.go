package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func getItemValue(item string) int {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	value := 0
	for index, word := range alphabet {
		if item == string(word) {
			value = index + 1
		}
	}

	return value
}

func calculateCommonItem(items string) string {
	itemsLen := len(items)

	firstRucksack := items[0 : itemsLen/2]
	secondRucksack := items[itemsLen/2 : itemsLen]

	var commonItem string

	for i := 0; i < itemsLen/2; i++ {
		for j := 0; j < itemsLen/2; j++ {
			if firstRucksack[i] == secondRucksack[j] {
				commonItem = string(firstRucksack[i])
				break
			}
		}
	}

	return commonItem
}

func getItemsArraySortedByLen(items []string) (string, string, string) {
	itemsSorted := sort.StringSlice(items)

	firstRucksack := itemsSorted[0]
	secondRucksack := itemsSorted[1]
	thirdRucksack := itemsSorted[2]

	return firstRucksack, secondRucksack, thirdRucksack
}

func calculateCommonItemPerGroup(items ...string) string {
	firstRucksack, secondRucksack, thirdRucksack := getItemsArraySortedByLen(items)

	var commonItem string

	for i := 0; i < len(firstRucksack); i++ {
		for j := 0; j < len(secondRucksack); j++ {
			for k := 0; k < len(thirdRucksack); k++ {
				if firstRucksack[i] == secondRucksack[j] && secondRucksack[j] == thirdRucksack[k] {
					commonItem = string(firstRucksack[i])
					break
				}
			}
		}
	}

	return commonItem
}

func main() {
	start := time.Now()
	// filename := "../sample.txt"
	filename := "../input.txt"

	contentInByte, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	rucksacks := strings.Split(string(contentInByte), "\n")

	total := 0
	totalPerGroup := 0
	for index, rucksack := range rucksacks {
		groupMod := (index + 1) % 3
		if groupMod == 0 && index > 0 {
			groupItem := calculateCommonItemPerGroup(rucksacks[index-2], rucksacks[index-1], rucksacks[index])
			totalPerGroup += getItemValue(groupItem)
		}

		item := calculateCommonItem(rucksack)
		total += getItemValue(item)
	}

	fmt.Printf("The sum of the priorities of those item types is:  %v.\n", total)
	fmt.Printf("The sum of the group priorities of those item types is:  %v.\n", totalPerGroup)

	timeElapsed := time.Since(start)
	fmt.Printf("MY  SOLUTION IN GO %v\n", timeElapsed)
}
