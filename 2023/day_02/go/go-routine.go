package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func getTotalGamesId2(contentInString string) int {
	sum := 0

	lines := strings.Split(contentInString, "\n")

	for _, line := range lines {
		l := strings.Replace(line, "Game ", "", 1)
		lineAux := strings.Split(l, ":")

		gameIndex, err := strconv.Atoi(lineAux[0])
		if err != nil {
			log.Panic("Game index is not a number")
		}

		isValidSet := true

		sets := strings.Split(lineAux[1], ";")

		for _, set := range sets {
			cubesLine := strings.Split(set, ",")

			redSum := 0
			greenSum := 0
			blueSum := 0

			redChan := make(chan int)
			greenChan := make(chan int)
			blueChan := make(chan int)

			go processColor("red", cubesLine, redChan)
			go processColor("green", cubesLine, greenChan)
			go processColor("blue", cubesLine, blueChan)

			for i := 0; i < 3; i++ {
				select {
				case redSum = <-redChan:
					if redSum > 12 {
						isValidSet = false
					}
				case greenSum = <-greenChan:
					if greenSum > 13 {
						isValidSet = false
					}
				case blueSum = <-blueChan:
					if blueSum > 14 {
						isValidSet = false
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

func processColor(color string, cubesLine []string, ch chan int) {
	sum := 0
	for _, cube := range cubesLine {
		if strings.Contains(cube, color) {
			cubeValue := strings.Replace(cube, color, "", 1)
			cubeNumber := strings.TrimSpace(cubeValue)

			cubeInt, err := strconv.Atoi(cubeNumber)
			if err != nil {
				log.Panic("Cube int is not a number")
			}

			sum += cubeInt
		}
	}
	ch <- sum
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

	twoDigitNumberSumPart1 := getTotalGamesId2(string(contentInBytePart1))
	// twoDigitNumberSumPart2 := getTwoDigitNumberPart2(string(contentInBytePart2))
	fmt.Printf("What is the sum of the IDs of those games? - part 1: %v .\n", twoDigitNumberSumPart1)
	// fmt.Printf("What is the sum of the IDs of those games? - part 2: %v .\n", twoDigitNumberSumPart2)

	timeElapsed := time.Since(start)
	fmt.Printf("MY  SOLUTION IN GO %v\n", timeElapsed)
}
