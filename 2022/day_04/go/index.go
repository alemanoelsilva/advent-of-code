package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	FIRST_NUMBER  = 0
	SECOND_NUMBER = 1
)

func getInteger(s string) int {
	n, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println("String to Int ERROR")
		os.Exit((1))
	}

	return n
}

func getOverlapPathsByPairs(paths []string) int {
	total := 0

	for _, path := range paths {
		// path: 14-28,13-28 | 72-81,82-91 | 59-61,15-60
		elves := strings.Split(path, ",")

		// elves: ["14-28","13-28"] | ["72-81","82-91"] | ["59-61","15-60"]
		elfA := strings.Split(elves[0], "-")
		elfB := strings.Split(elves[1], "-")

		// elfA: ["14","28"] | ["72","81"] | ["59","61"]
		// elfB: ["13","28"] | ["82","91"] | ["15","60"]
		var intElfA []int
		var intElfB []int

		for _, elf := range elfA {
			intElfA = append(intElfA, getInteger(elf)) // must to convert to int
		}
		for _, elf := range elfB {
			intElfB = append(intElfB, getInteger(elf))
		}

		if (intElfA[FIRST_NUMBER] >= intElfB[FIRST_NUMBER] && intElfA[SECOND_NUMBER] <= intElfB[SECOND_NUMBER]) ||
			(intElfA[FIRST_NUMBER] <= intElfB[FIRST_NUMBER] && intElfA[SECOND_NUMBER] >= intElfB[SECOND_NUMBER]) {
			total++
		}
	}

	return total
}

func getOverlapPathsInTotal(paths []string) int {
	total := 0

	for _, path := range paths {
		elves := strings.Split(path, ",")

		elfA := strings.Split(elves[0], "-")
		elfB := strings.Split(elves[1], "-")

		var intElfA []int
		var intElfB []int

		for _, elf := range elfA {
			intElfA = append(intElfA, getInteger(elf))
		}
		for _, elf := range elfB {
			intElfB = append(intElfB, getInteger(elf))
		}

		if intElfA[FIRST_NUMBER] >= intElfB[FIRST_NUMBER] && intElfA[FIRST_NUMBER] <= intElfB[SECOND_NUMBER] ||
			intElfA[SECOND_NUMBER] >= intElfB[FIRST_NUMBER] && intElfA[SECOND_NUMBER] <= intElfB[SECOND_NUMBER] ||
			intElfB[FIRST_NUMBER] >= intElfA[FIRST_NUMBER] && intElfB[FIRST_NUMBER] <= intElfA[SECOND_NUMBER] ||
			intElfB[SECOND_NUMBER] >= intElfA[FIRST_NUMBER] && intElfB[SECOND_NUMBER] <= intElfA[SECOND_NUMBER] {
			total++
		}
	}

	return total
}

func main() {
	start := time.Now()
	// filename := "../sample.txt"
	filename := "../input.txt"

	contentInByte, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	paths := strings.Split(string(contentInByte), "\n")

	totalOverlapPathsByPairs := getOverlapPathsByPairs(paths)
	totalOverlapPathsInTotal := getOverlapPathsInTotal(paths)

	fmt.Printf("Assignment pairs does one range fully contain the other is:  %v .\n", totalOverlapPathsByPairs)
	fmt.Printf("Assignment overlap in total is:  %v .\n", totalOverlapPathsInTotal)

	timeElapsed := time.Since(start)
	fmt.Printf("MY  SOLUTION IN GO %v\n", timeElapsed)
}
