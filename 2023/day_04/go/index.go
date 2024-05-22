package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const TEST_FILE_NAME = "../sample_part_1.txt"
const FILE_NAME = "../input_part_1.txt"

func getPoints(winners int) float64 {
	if winners == 0 {
		return 0
	}
	return math.Pow(2, float64(winners-1))
}

func convertStringToInt(str []string) []int {
	var i []int
	for _, value := range str {
		number, _ := strconv.Atoi(value)
		i = append(i, number)
	}

	return i
}

func getTotalCardsPoints(lines []string) float64 {
	sum := 0.0

	for _, line := range lines {
		parts := strings.Split(line, " | ")
		gameStr := strings.Split(parts[0], ": ")[1]
		playStr := strings.Split(strings.Trim(parts[1], " "), " ")
		playInNumbers := convertStringToInt(playStr)

		games := strings.Split(strings.Trim(gameStr, " "), " ")
		gamesInNumbers := convertStringToInt(games)

		var winners []int

		for _, g := range gamesInNumbers {
			for _, p := range playInNumbers {
				if g == p && p > 0 {
					winners = append(winners, g)
				}
			}
		}

		sum = sum + getPoints(len(winners))
		fmt.Printf("Winners: %v | temp points: %v | total points: %v\n", winners, getPoints(len(winners)), sum)
	}

	return sum
}

func main() {
	start := time.Now()

	filename := TEST_FILE_NAME
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		fmt.Println("Using test file")
	}

	contentInBytePart1, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(contentInBytePart1), "\n")

	totalPartNumbers := getTotalCardsPoints(lines)
	// twoDigitNumberSumPart2 := getTotalPowerCube(lines)
	fmt.Printf("How many points are they worth in total? - part 1: %v .\n", totalPartNumbers)
	// fmt.Printf("How many points are they worth in total? - part 2: %v .\n", twoDigitNumberSumPart2)

	timeElapsed := time.Since(start)
	fmt.Printf("MY  SOLUTION IN GO %v.\n", timeElapsed)
}
