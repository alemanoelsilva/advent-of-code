package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const TEST_FILE_NAME = "../sample_part_1.txt"
const FILE_NAME = "../input_part_1.txt"

var directions = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

const NOT_SYMBOLS = "1234567890."

func isPartNumber(row, col int, grid [][]rune) bool {
	for _, dir := range directions {
		newRow, newCol := row+dir[0], col+dir[1]
		if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[0]) {

			str := string(grid[newRow][newCol])
			if !strings.Contains(NOT_SYMBOLS, str) {
				return true
			}
		}
	}
	return false
}

type Concat struct {
	number       string
	isPartNumber bool
}

func getTotalEngineSchematic(lines []string) int {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	// var sumNumber []int
	sum := 0

	for row, line := range lines {
		concat := Concat{
			number:       "",
			isPartNumber: false,
		}

		for col, value := range strings.Split(line, "") {
			_, err := strconv.Atoi(value)
			if err != nil {
				if concat.number != "" {
					num, err := strconv.Atoi(concat.number)
					if err == nil && concat.isPartNumber {
						// // sumNumber = append(sumNumber, num)
						sum = sum + num
					}
					concat.number = ""
					concat.isPartNumber = false
				}
			} else {
				concat.number = concat.number + value
				if !concat.isPartNumber {
					concat.isPartNumber = isPartNumber(row, col, grid)
				}

				// last col of the line
				if col == len(line)-1 {
					if concat.number != "" {
						num, err := strconv.Atoi(concat.number)
						if err == nil && concat.isPartNumber {
							// // sumNumber = append(sumNumber, num)
							sum = sum + num
						}
						concat.number = ""
						concat.isPartNumber = false
					}
				}
			}
		}
	}

	// fmt.Printf("%v\n", sumNumber)

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

	totalPartNumbers := getTotalEngineSchematic(lines)
	// twoDigitNumberSumPart2 := getTotalPowerCube(lines)
	fmt.Printf("What is the sum of all of the part numbers in the engine schematic? - part 1: %v .\n", totalPartNumbers)
	// fmt.Printf("What is the sum of all of the part numbers in the engine schematic? - part 2: %v .\n", twoDigitNumberSumPart2)

	timeElapsed := time.Since(start)
	fmt.Printf("MY  SOLUTION IN GO %v.\n", timeElapsed)
}
