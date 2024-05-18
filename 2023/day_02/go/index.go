package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func getTotalGamesId(contentInString string) int {
	sum := 0

	lines := strings.Split(contentInString, "\n")

	// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	for _, line := range lines {
		l := strings.Replace(line, "Game ", "", 1)
		lineAux := strings.Split(l, ":")

		gameIndex, err := strconv.Atoi(lineAux[0])
		if err != nil {
			log.Panic("Game index is not a number")
		}

		isValidSet := true

		sets := strings.Split(lineAux[1], ";")

		// [
		//	"8 green, 6 blue, 20 red";
		//	"5 blue, 4 red, 13 green";
		//	"5 green, 1 red"
		// ]
		for _, set := range sets {
			cubesLine := strings.Split(set, ",")

			redSum := 0
			greenSum := 0
			blueSum := 0

			//	["8 green", "6 blue", "20 red"]
			for _, cube := range cubesLine {
				if strings.Contains(cube, "red") {
					cubeValue := strings.Replace(cube, "red", "", 1)
					cubeNumber := strings.Trim(cubeValue, " ")

					cubeInt, err := strconv.Atoi(cubeNumber)
					if err != nil {
						log.Panic("Cube int is not a number")
					}

					redSum = redSum + cubeInt
					if redSum > 12 {
						isValidSet = false
						break
					}
				}

				if strings.Contains(cube, "green") {
					cubeValue := strings.Replace(cube, "green", "", 1)
					cubeNumber := strings.Trim(cubeValue, " ")

					cubeInt, err := strconv.Atoi(cubeNumber)
					if err != nil {
						log.Panic("Cube int is not a number")
					}

					greenSum = greenSum + cubeInt
					if greenSum > 13 {
						isValidSet = false
						break
					}
				}

				if strings.Contains(cube, "blue") {
					cubeValue := strings.Replace(cube, "blue", "", 1)
					cubeNumber := strings.Trim(cubeValue, " ")

					cubeInt, err := strconv.Atoi(cubeNumber)
					if err != nil {
						log.Panic("Cube int is not a number")
					}

					blueSum = blueSum + cubeInt
					if blueSum > 14 {
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

	filename := "../sample_part_1.txt"
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

	twoDigitNumberSumPart1 := getTotalGamesId(string(contentInBytePart1))
	// twoDigitNumberSumPart2 := getTwoDigitNumberPart2(string(contentInBytePart2))
	fmt.Printf("What is the sum of the IDs of those games? - part 1: %v .\n", twoDigitNumberSumPart1)
	// fmt.Printf("What is the sum of the IDs of those games? - part 2: %v .\n", twoDigitNumberSumPart2)

	timeElapsed := time.Since(start)
	fmt.Printf("MY  SOLUTION IN GO %v\n", timeElapsed)
}
