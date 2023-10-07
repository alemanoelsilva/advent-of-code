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

func getRoundChoicePoint(roundPlay string) int {
	switch roundPlay {
	// Elf plays Rock | I need to lose
	case "A X":
		return SCISSOR_PLAY
		// Elf plays Rock | I need to draw
	case "A Y":
		return ROCK_PLAY
		// Elf plays Rock | I need to win
	case "A Z":
		return PAPER_PLAY
		// Elf plays Paper | I need to lose
	case "B X":
		return ROCK_PLAY
		// Elf plays Paper | I need to draw
	case "B Y":
		return PAPER_PLAY
		// Elf plays Paper | I need to win
	case "B Z":
		return SCISSOR_PLAY
		// Elf plays Scissor | I need to lose
	case "C X":
		return PAPER_PLAY
		// Elf plays Scissor | I need to draw
	case "C Y":
		return SCISSOR_PLAY
		// Elf plays Scissor | I need to win
	case "C Z":
		return ROCK_PLAY
	default:
		return 0
	}
}

func getRoundResultPoint(roundPlay string) int {
	plays := strings.Split(roundPlay, " ")

	myPlay := plays[1]

	switch myPlay {
	// I will lose
	case "X":
		return LOST
		// I will draw
	case "Y":
		return DRAW
		// I will win
	case "Z":
		return WON
	default:
		return 0
	}
}

func getTotalPointsOfRoundResult(contentInString string) int {
	rounds := strings.Split(contentInString, "\n")

	totalPoints := 0

	for _, round := range rounds {
		totalPoints = totalPoints + getRoundPoint(round) + getMyPlayPoints(round)
	}

	return totalPoints
}

func getTotalPointsOfChoiceResult(contentInString string) int {
	rounds := strings.Split(contentInString, "\n")

	totalPoints := 0

	for _, round := range rounds {
		choicePoints := getRoundChoicePoint(round)
		roundResultPoints := getRoundResultPoint(round)
		totalPoints = totalPoints + choicePoints + roundResultPoints
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

	totalPointsOfRound := getTotalPointsOfRoundResult(string(contentInByte))
	totalPointsOfChoice := getTotalPointsOfChoiceResult(string(contentInByte))
	fmt.Printf("The total points of round  %v.\n", totalPointsOfRound)
	fmt.Printf("The total points of choice %v.\n", totalPointsOfChoice)

	timeElapsed := time.Since(start)
	fmt.Printf("MY  SOLUTION IN GO %v\n", timeElapsed)
}
