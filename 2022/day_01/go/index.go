package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func getMaxElfCalories(contentInString string) (int, int) {

	var resultList []int
	var sum int = 0

	elvesCaloriesArr := strings.Split(contentInString, "\n\n")
	for _, elfCaloriesInString := range elvesCaloriesArr {
		elfCaloriesArr := strings.Split(elfCaloriesInString, "\n")

		for elfCalorieIndex, stringSubValue := range elfCaloriesArr {
			intValue, _ := strconv.Atoi(stringSubValue)

			sum = sum + intValue

			if len(elfCaloriesArr) == elfCalorieIndex+1 {
				resultList = append(resultList, sum)
				sum = 0
			}
		}

	}

	sort.Sort(sort.Reverse(sort.IntSlice(resultList)))

	elfWithMoreCalorie := resultList[0]
	secondElf := resultList[1]
	thirdElf := resultList[2]

	threeElvesWithMoreCalories := elfWithMoreCalorie + secondElf + thirdElf

	return elfWithMoreCalorie, threeElvesWithMoreCalories
}

func main() {
	start := time.Now()
	// filename := "../sample.txt"
	filename := "../input.txt"

	contentInByte, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	elfWithMoreCalorie, threeElvesWithMoreCalories := getMaxElfCalories(string(contentInByte))
	fmt.Printf("The Elf carrying the most Calories has %v Calories.\n", elfWithMoreCalorie)
	fmt.Printf("The top 3 Elves are carrying %v Calories.\n", threeElvesWithMoreCalories)

	timeElapsed := time.Since(start)
	fmt.Printf("MY  SOLUTION IN GO %v\n", timeElapsed)
}
