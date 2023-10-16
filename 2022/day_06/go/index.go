package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func getFileContent() string {
	// filename := "../sample.txt"
	filename := "../input.txt"

	contentInByte, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(contentInByte)
}

func countStartOfPacketMarker(input string) int {
	total := 0
	start := input[0:3]
	datastream := input[3:]

	for idx, data := range strings.Split(datastream, "") {
		startArr := strings.Split(start, "")
		startArr = append(startArr, data)

		fourDigitsStr := strings.Join(startArr, "")
		if startArr[0] != startArr[1] && startArr[0] != startArr[2] &&
			startArr[0] != startArr[3] && startArr[1] != startArr[2] &&
			startArr[1] != startArr[3] && startArr[2] != startArr[3] {
			total = idx + 4
			break
		}

		start = fourDigitsStr[1:]

	}

	return total
}

func countStartOfMessageMarker(input string) int {
	total := 0
	start := input[0:13]
	datastream := input[13:]

	fmt.Printf("input |%v|\n", input)
	fmt.Printf("start |%v|\n", start)
	fmt.Printf("datastream |%v|\n", datastream)

	for idx, data := range strings.Split(datastream, "") {
		startArr := strings.Split(start, "")
		startArr = append(startArr, data)

		fourteenDigitsStr := strings.Join(startArr, "")
		// fmt.Printf("%v --> startArr |%v|\n", idx, startArr)

		isStartOfMessage := true

		// going throughout each 14 positions
		for i, startStr := range startArr {
			// for each iteration, check all next positions
			for j := i + 1; j < len(startArr); j++ {
				if startStr == startArr[j] {
					isStartOfMessage = false
				}
			}
		}

		if isStartOfMessage {
			total = idx + 14
			break
		}

		start = fourteenDigitsStr[1:]

	}

	return total
}

func main() {
	start := time.Now()
	// code

	input := getFileContent()
	totalPacket := countStartOfPacketMarker(input)
	totalMessage := countStartOfMessageMarker(input)

	// end code

	fmt.Printf("How many characters need to be processed before the first start-of-packet  marker is detected:  |%v|.\n", totalPacket)
	fmt.Printf("How many characters need to be processed before the first start-of-message marker is detected:  |%v|.\n", totalMessage)

	timeElapsed := time.Since(start)
	fmt.Printf("MY  SOLUTION IN GO %v\n", timeElapsed)
}
