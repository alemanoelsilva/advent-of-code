package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Function to calculate points for a single card
func calculatePoints(winningNumbers, yourNumbers []int) int {
	winningSet := make(map[int]bool)
	for _, number := range winningNumbers {
		winningSet[number] = true
	}

	matchCount := 0
	for _, number := range yourNumbers {
		if winningSet[number] {
			matchCount++
		}
	}

	if matchCount == 0 {
		return 0
	}

	return 1 << (matchCount - 1) // 2^(matchCount-1)
}

// Function to convert a string of space-separated numbers to a slice of integers
func stringToIntSlice(s string) []int {
	fields := strings.Fields(s)
	numbers := make([]int, len(fields))
	for i, field := range fields {
		number, _ := strconv.Atoi(field)
		numbers[i] = number
	}
	return numbers
}

// Function to process all cards and calculate the total points
func calculateTotalPoints(cards []string) int {
	totalPoints := 0

	for _, card := range cards {
		parts := strings.Split(card, " | ")
		winningNumbers := stringToIntSlice(parts[0])
		yourNumbers := stringToIntSlice(parts[1])

		totalPoints += calculatePoints(winningNumbers, yourNumbers)
	}

	return totalPoints
}
func main() {
	// Example input
	filename := "../input_part_1.txt"
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

	// Calculate the total points
	totalPoints := calculateTotalPoints(lines)

	// Print the result
	fmt.Println("Total Points:", totalPoints)
}
