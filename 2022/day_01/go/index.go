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

func getMaxElfCalories(filename string) int {
	contentInByte, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	elvesCaloriesArr := strings.Split(string(contentInByte), "\n\n")

	var resultList []int
	var sum int = 0

	start := time.Now()
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

	timeElapsed := time.Since(start)

	fmt.Printf("MY  SOLUTION IN GO %v\n", timeElapsed)

	return resultList[0]
}

func main() {
	// filename := "../sample.txt"
	filename := "../input.txt"
	result := getMaxElfCalories(filename)
	fmt.Printf("The Elf carrying the most Calories has %v Calories.", result)
}
