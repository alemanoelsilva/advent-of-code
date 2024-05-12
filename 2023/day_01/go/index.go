package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func openFile() *os.File {
	file, err := os.Create("result_2.txt")
	if err != nil {
		log.Fatal("Error on Open File:", err)
		return nil
	}

	return file
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal("Error on Close File:", err)
	}
}

func writeFile(file *os.File, content string) {
	_, err := file.WriteString(content)
	if err != nil {
		fmt.Println("Error on Write File:", err)
		return
	}
}

func reverse(str []string) []string {
	n := len(str)
	for i := 0; i < n/2; i++ {
		str[i], str[n-i-1] = str[n-i-1], str[i]
	}
	return str
}

func concat(intA int, intB int) int {
	str := fmt.Sprintf("%d%d", intA, intB)

	twoDigitNumber, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Error on Concat Int", err)
	}
	return twoDigitNumber
}

func getTwoDigitNumberPart1(contentInString string) int {
	var sum int = 0

	calibratedValuesArr := strings.Split(contentInString, "\n")

	for _, line := range calibratedValuesArr {
		var resultList []int

		for _, value := range line {
			intValue, err := strconv.Atoi(string(value))
			if err == nil {
				resultList = append(resultList, intValue)
				break
			}
		}

		reversedLine := reverse(strings.Split(line, ""))

		for _, value := range reversedLine {
			intValue, err := strconv.Atoi(string(value))
			if err == nil {
				resultList = append(resultList, intValue)
				break
			}
		}

		sum = sum + concat(resultList[0], resultList[1])

	}

	return sum
}

var numberArr = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func getTwoDigitNumberPart2(contentInString string) int {
	file := openFile()
	defer closeFile(file)

	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var sum int = 0

	calibratedValuesArr := strings.Split(contentInString, "\n")

	for _, line := range calibratedValuesArr {
		var resultList []int
		indexOfNumber := 0
		firstNumber := 0
		secondNumber := 0

		// find first numeric value
		for index, value := range line {
			intValue, err := strconv.Atoi(string(value))
			if err == nil {
				firstNumber = intValue
				indexOfNumber = index
				break
			}
		}

		// find first string number name value
		for _, numberStr := range numberArr {
			indexOfStringNumberName := strings.Index(line, numberStr)

			// if index of "string number name value" is lower than "numeric value", get this number, otherwise, get the numeric one
			if indexOfStringNumberName >= 0 && indexOfStringNumberName <= indexOfNumber {
				firstNumber = numbers[numberStr]
				indexOfNumber = indexOfStringNumberName
			}
		}

		resultList = append(resultList, firstNumber)

		// second number
		indexOfNumber = 0

		reversedLine := reverse(strings.Split(line, ""))

		// find second numeric value
		for _, value := range reversedLine {
			intValue, err := strconv.Atoi(string(value))
			if err == nil {
				secondNumber = intValue
				// string is reversed
				realIndex := strings.LastIndex(line, value)
				indexOfNumber = realIndex
				break
			}
		}

		// find first string number name value
		for _, numberStr := range numberArr {
			indexOfStringNumberName := strings.LastIndex(line, numberStr)

			// if index of "string number name value" is lower than "numeric value", get this number, otherwise, get the numeric one
			if indexOfStringNumberName >= 0 && indexOfStringNumberName >= indexOfNumber {
				secondNumber = numbers[numberStr]
				indexOfNumber = indexOfStringNumberName
			}

		}

		resultList = append(resultList, secondNumber)

		// fmt.Printf("Line: %v -- first: %v -- second: %v\n", line, resultList[0], resultList[1])
		strResult := fmt.Sprintf("line: %v -- first: %v, second: %v\n", line, resultList[0], resultList[1])
		// fmt.Print(strResult)

		writeFile(file, strResult)
		sum = sum + concat(resultList[0], resultList[1])
	}

	return sum
}

func main() {
	start := time.Now()
	// filenamePart1 := "../sample_part_1.txt"
	// filenamePart2 := "../sample_part_2.txt"
	filenamePart1 := "../input_part_1.txt"
	filenamePart2 := "../input_part_2.txt"

	contentInBytePart1, err := os.ReadFile(filenamePart1)

	if err != nil {
		log.Fatal(err)
	}

	contentInBytePart2, err := os.ReadFile(filenamePart2)

	if err != nil {
		log.Fatal(err)
	}

	twoDigitNumberSumPart1 := getTwoDigitNumberPart1(string(contentInBytePart1))
	twoDigitNumberSumPart2 := getTwoDigitNumberPart2(string(contentInBytePart2))
	fmt.Printf("What is the sum of all of the calibration values part 1: %v .\n", twoDigitNumberSumPart1)
	fmt.Printf("What is the sum of all of the calibration values part 2: %v .\n", twoDigitNumberSumPart2)

	timeElapsed := time.Since(start)
	fmt.Printf("MY  SOLUTION IN GO %v\n", timeElapsed)
}
