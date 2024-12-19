package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func containsMAS(word string) bool {
	return strings.Contains(word, "M") &&
		strings.Contains(word, "A") &&
		strings.Contains(word, "S") &&
		len(word) == 3
}

func getDiagonals(charIndex int, lineIndex int, crossword [][]string) (string, string) {
	lineBelow := lineIndex + 1
	lineAbove := lineIndex - 1
	charForward := charIndex + 1
	charBackward := charIndex - 1

	leftToRightUp := "A"
	leftToRightDown := "A"

	if lineBelow < len(crossword) && charBackward >= 0 {
		leftToRightUp += crossword[lineBelow][charBackward]
	}
	if lineAbove >= 0 && charForward < len(crossword[lineAbove]) {
		leftToRightUp += crossword[lineAbove][charForward]
	}

	if lineAbove >= 0 && charBackward >= 0 {
		leftToRightDown += crossword[lineAbove][charBackward]
	}
	if lineBelow < len(crossword) && charForward < len(crossword[lineBelow]) {
		leftToRightDown += crossword[lineBelow][charForward]
	}

	return leftToRightUp, leftToRightDown
}

func main() {
	file, err := os.Open("input/puzzle_part2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var crossword [][]string
	for scanner.Scan() {
		line := scanner.Text()
		lineCharts := strings.Split(line, "")
		crossword = append(crossword, lineCharts)
	}

	count := 0
	for lineIndex, line := range crossword {
		for charIndex, letter := range line {
			if letter == "A" {
				leftToRightUp, leftToRightDown := getDiagonals(charIndex, lineIndex, crossword)
				if containsMAS(leftToRightDown) && containsMAS(leftToRightUp) {
					count++
				}

			}
		}
	}

	fmt.Println("X-MAS Count:", count)
}
