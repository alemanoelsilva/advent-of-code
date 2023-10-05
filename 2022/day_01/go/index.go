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

func getMaxElfCalories(contentInString string) int {

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

	return resultList[0]
}

func main() {
	start := time.Now()
	// filename := "../sample.txt"
	filename := "../input.txt"

	contentInByte, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	result := getMaxElfCalories(string(contentInByte))
	fmt.Printf("The Elf carrying the most Calories has %v Calories.\n", result)

	timeElapsed := time.Since(start)
	fmt.Printf("MY  SOLUTION IN GO %v\n", timeElapsed)
}
