package main

import (
	"fmt"
	"io"
	"os"
)

var instructionValues map[rune]int = map[rune]int{
	'(': 1,
	')': -1,
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}
	// Part 1

	floor := findFloor(input)
	fmt.Printf("floor number is: %d", floor)

	// Part 2
	instructionNumber := findFirstBasementEntry(input)
	fmt.Printf("first basement entry was on instruction number: %d", instructionNumber)
}

func findFloor(instructions string) int {
	f := 0
	for _, r := range instructions {
		f += instructionValues[r]
	}
	return f
}

func readInput(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	bytes, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func findFirstBasementEntry(instructions string) int {
	var floor int
	for i, r := range instructions {
		floor += instructionValues[r]
		if floor == -1 {
			return (i + 1)
		}
	}
	return -1
}
