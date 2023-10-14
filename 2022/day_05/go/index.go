package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type position struct {
	start int
	end   int
}

type TempStack struct {
	name          int
	cratePosition position
}

type Stack struct {
	name   int
	crates []string
}

func (s Stack) GetCrates(stackName int) []string {
	if s.name == stackName {
		return s.crates
	}

	return []string{}
}

func (s *Stack) AddCrate(crate string) {
	s.crates = append(s.crates, crate)
	fmt.Printf("crates added |%v| to |%v|\n", crate, s)
}

func (s *Stack) AddCrateInit(crate string) {
	s.crates = append([]string{crate}, s.crates...)
	fmt.Printf("crates added |%v| to |%v|\n", crate, s)
}

func (s *Stack) RemoveCrate(quantity int) []string {
	fmt.Printf("stack |%v|\n", s)
	fmt.Printf("stack quantity |%v|\n", quantity)

	var cratesToRemove []string

	// when moving all crates, the order needs to remain
	fmt.Printf("INVERT? len |%v| --- quantity |%v|\n", len(s.crates), quantity)
	if quantity > 1 {
		cratesToRemove = s.crates[0:quantity]
		s.crates = append(s.crates[quantity:])

		fmt.Printf("crates to remove |%v|\n", cratesToRemove)
	} else {
		for i := 0; i < quantity; i++ {
			cratesToRemove = append(cratesToRemove, s.crates[0])
			s.crates = append(s.crates[1:])
		}

		fmt.Printf("crates to remove |%v|\n", cratesToRemove)
	}

	return cratesToRemove

}

func getFileContent() []string {
	// filename := "../sample.txt"
	filename := "../input.txt"

	contentInByte, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(contentInByte), "\n")
}

func buildCrateStacks(content []string) []Stack {
	var tempStack []TempStack
	var stacks []Stack

	// build tempo stack
	for _, stackIndexName := range content {
		_, err := strconv.Atoi(stackIndexName[1:2])

		// means that i'm in the index row
		if err == nil {
			// filter numeric value
			var idx []int
			stackArr := strings.Split(stackIndexName, "")
			for _, stackArrValue := range stackArr {
				stackArrValue, err := strconv.Atoi(stackArrValue)
				if err == nil {
					idx = append(idx, stackArrValue)
				}
			}

			for i, name := range idx {
				if i == 0 {
					tempStack = append(tempStack, TempStack{
						name:          name,
						cratePosition: position{start: 1, end: 2},
					})
				} else {
					tempStack = append(tempStack, TempStack{
						name:          name,
						cratePosition: position{start: (i * 4) + 1, end: (i * 4) + 2},
					})
				}

			}
			break
		}

	}

	// initialize stacks
	for _, ts := range tempStack {
		stacks = append(stacks, Stack{
			name:   ts.name,
			crates: []string{},
		})
	}

	// build stacks
	for _, stackIndexName := range content {
		// when reach the stack index, break
		if stackIndexName[1:2] == "1" {
			break
		}

		for _, ts := range tempStack {

			stackName := ts.name - 1

			crate := stackIndexName[ts.cratePosition.start:ts.cratePosition.end]

			if crate != " " {
				stacks[stackName].AddCrate(crate)
			}
		}
	}

	fmt.Printf("Hey stacks |%v|\n", stacks)
	return stacks
}

func getMovement(row string) (int, int, int) {
	fmt.Printf("row bkp |%v|\n", row)
	const (
		MOVE = "move"
		FROM = "from"
		TO   = "to"
	)

	var quantity, origin, destiny int
	str := row
	str = strings.Replace(str, MOVE, "", 2)
	str = strings.Replace(str, FROM, "", 2)
	str = strings.Replace(str, TO, "", 2)

	isQuantity := true
	var isOrigin, isDestiny bool

	for _, actionValue := range strings.Split(str, " ") {
		if actionValue != " " {
			value, err := strconv.Atoi(actionValue)
			if value == 0 {
				continue
			}
			fmt.Printf("row value |%v|\n", value)

			if err == nil {
				if isQuantity {
					quantity = int(value)
					isQuantity = false
					isOrigin = true
					continue
				}

				if isOrigin {
					origin = value - 1 // array starts on ZERO
					isOrigin = false
					isDestiny = true
					continue
				}

				if isDestiny {
					destiny = value - 1 // array starts on ZERO
				}
			}
		}
	}

	fmt.Printf("quantity |%v| --- origin |%v| --- destiny |%v|\n", quantity, origin, destiny)
	return quantity, origin, destiny
}

func reverse(str []string) []string {
	n := len(str)
	for i := 0; i < n/2; i++ {
		str[i], str[n-i-1] = str[n-i-1], str[i]
	}
	return str
}

func moveStacks() string {
	rows := getFileContent()
	stacks := buildCrateStacks(rows)

	var result []string

	for _, row := range rows {
		fmt.Printf("row |%v|\n", row)
		if len(row) > 0 && row[0:1] == "m" {
			quantity, origin, destiny := getMovement(row)
			cratesToMove := stacks[origin].RemoveCrate(quantity)

			fmt.Printf("INVERT? len |%v| --- quantity |%v|\n", len(stacks[origin].crates), quantity)
			if quantity > 1 {
				fmt.Println("inverted ")
				for _, crate := range reverse(cratesToMove) {
					stacks[destiny].AddCrateInit(crate)
				}
			} else {
				fmt.Println("normal ")
				for _, crate := range cratesToMove {
					stacks[destiny].AddCrateInit(crate)
				}
			}
		}
	}

	fmt.Printf("FINAL stacks |%v|\n", stacks)

	for _, stack := range stacks {
		result = append(result, stack.crates[0])
	}

	return strings.Join(result, "")
}

func main() {
	start := time.Now()

	// code

	result := moveStacks()

	// end code

	fmt.Printf("After the rearrangement procedure completes, what crate ends up on top of each stack:  |%v|.\n", result)

	timeElapsed := time.Since(start)
	fmt.Printf("MY  SOLUTION IN GO %v\n", timeElapsed)

}
