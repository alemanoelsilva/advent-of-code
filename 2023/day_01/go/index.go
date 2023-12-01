package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

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
		log.Panicf("concat failed %v", err)
	}
	return twoDigitNumber
}

func getTwoDigitNumber(contentInString string) int {
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

func main() {
	start := time.Now()
	// filename := "../sample.txt"
	filename := "../input.txt"

	contentInByte, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	twoDigitNumberSum := getTwoDigitNumber(string(contentInByte))
	fmt.Printf("What is the sum of all of the calibration values %v.\n", twoDigitNumberSum)
	// fmt.Printf("The top 3 Elves are carrying %v Calories.\n", threeElvesWithMoreCalories)

	timeElapsed := time.Since(start)
	fmt.Printf("MY  SOLUTION IN GO %v\n", timeElapsed)
}
