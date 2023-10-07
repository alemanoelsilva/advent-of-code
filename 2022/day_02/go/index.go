package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	WON          = 6
	DRAW         = 3
	LOST         = 0
	ROCK_PLAY    = 1
	PAPER_PLAY   = 2
	SCISSOR_PLAY = 3
)

func getRoundPoint(roundPlay string) int {
	switch roundPlay {
	// Elf plays Rock | I play Rock
	case "A X":
		return DRAW
		// Elf plays Rock | I play Paper
	case "A Y":
		return WON
		// Elf plays Rock | I play Scissor
	case "A Z":
		return LOST
		// Elf plays Paper | I play Rock
	case "B X":
		return LOST
		// Elf plays Paper | I play Paper
	case "B Y":
		return DRAW
		// Elf plays Paper | I play Scissor
	case "B Z":
		return WON
		// Elf plays Scissor | I play Rock
	case "C X":
		return WON
		// Elf plays Scissor | I play Paper
	case "C Y":
		return LOST
		// Elf plays Scissor | I play Scissor
	case "C Z":
		return DRAW
	default:
		return 0
	}
}

func getMyPlayPoints(roundPlay string) int {
	// return [elfPlay, myPlay]
	plays := strings.Split(roundPlay, " ")

	myPlay := plays[1]

	switch myPlay {
	// I play Rock
	case "X":
		return ROCK_PLAY
		// I play Paper
	case "Y":
		return PAPER_PLAY
		// I play Scissor
	case "Z":
		return SCISSOR_PLAY
	default:
		return 0
	}
}

func getTotalPoints(contentInString string) int {
	rounds := strings.Split(contentInString, "\n")

	totalPoints := 0

	for _, round := range rounds {
		totalPoints = totalPoints + getRoundPoint(round) + getMyPlayPoints(round)
	}

	return totalPoints
}

func main() {
	start := time.Now()
	// filename := "../sample.txt"
	filename := "../input.txt"

	contentInByte, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	totalPoints := getTotalPoints(string(contentInByte))
	fmt.Printf("The total points  %v.\n", totalPoints)

	timeElapsed := time.Since(start)
	fmt.Printf("MY  SOLUTION IN GO %v\n", timeElapsed)
}
