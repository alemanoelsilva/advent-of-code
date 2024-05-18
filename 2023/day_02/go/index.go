package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const RED = "red"
const GREEN = "green"
const BLUE = "blue"

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

const BREAK_SETS = ";"
const BREAK_CUBES = ","

const TEST_FILE_NAME = "../sample_part_1.txt"

func getCubeValue(cube, color string) int {
	cubeValue := strings.Replace(cube, color, "", 1)
	cubeNumber := strings.Trim(cubeValue, " ")

	cubeInt, err := strconv.Atoi(cubeNumber)
	if err != nil {
		log.Panic("Cube int is not a number")
	}

	return cubeInt
}

func getTotalGamesId(lines []string) int {
	sum := 0

	// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	for _, line := range lines {
		lineAux := strings.Split(strings.Replace(line, "Game ", "", 1), ":")

		gameIndex, err := strconv.Atoi(lineAux[0])
		if err != nil {
			log.Panic("Game index is not a number")
		}

		isValidSet := true

		sets := strings.Split(lineAux[1], BREAK_SETS)

		// [
		//	"8 green, 6 blue, 20 red";
		//	"5 blue, 4 red, 13 green";
		//	"5 green, 1 red"
		// ]
		for _, set := range sets {
			cubesLine := strings.Split(set, BREAK_CUBES)

			redSum := 0
			greenSum := 0
			blueSum := 0

			//	["8 green", "6 blue", "20 red"]
			for _, cube := range cubesLine {
				if strings.Contains(cube, RED) {
					redSum += getCubeValue(cube, RED)
					if redSum > MAX_RED {
						isValidSet = false
						break
					}
				}

				if strings.Contains(cube, GREEN) {
					greenSum += getCubeValue(cube, GREEN)
					if greenSum > MAX_GREEN {
						isValidSet = false
						break
					}
				}

				if strings.Contains(cube, BLUE) {
					blueSum += getCubeValue(cube, BLUE)
					if blueSum > MAX_BLUE {
						isValidSet = false
						break
					}
				}
			}

			if !isValidSet {
				break
			}
		}

		if isValidSet {
			sum = sum + gameIndex
		}
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

	// filenamePart1 := "../sample_part_1.txt"
	// filenamePart2 := "../sample_part_2.txt"
	// filenamePart1 := "../input_part_1.txt"
	// filenamePart2 := "../input_part_2.txt"

	contentInBytePart1, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// contentInBytePart2, err := os.ReadFile(filenamePart2)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	lines := strings.Split(string(contentInBytePart1), "\n")

	twoDigitNumberSumPart1 := getTotalGamesId(lines)
	// twoDigitNumberSumPart2 := getTwoDigitNumberPart2(string(contentInBytePart2))
	fmt.Printf("What is the sum of the IDs of those games? - part 1: %v .\n", twoDigitNumberSumPart1)
	// fmt.Printf("What is the sum of the IDs of those games? - part 2: %v .\n", twoDigitNumberSumPart2)

	timeElapsed := time.Since(start)
	fmt.Printf("MY  SOLUTION IN GO %v.\n", timeElapsed)
	// fmt.Printf("MY  SOLUTION IN GO %v\n", timeElapsed)
}
