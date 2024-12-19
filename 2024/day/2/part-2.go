package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func isSafeRow(row []int) bool {
	isIncreasing := true
	isDecreasing := true
	for i := 1; i < len(row); i++ {
		diff := row[i] - row[i-1]
		if diff < 1 || diff > 3 {
			isIncreasing = false
		}
		if diff > -1 || diff < -3 {
			isDecreasing = false
		}
	}
	return isIncreasing || isDecreasing
}

func problemDamper(row []int) bool {
	for i := 0; i < len(row); i++ {
		newRow := append([]int{}, row[:i]...)
		newRow = append(newRow, row[i+1:]...)

		if isSafeRow(newRow) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input/puzzle_part2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		row := make([]int, len(parts))
		for i, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Error parsing %d:%v\n", lineNumber, err)
				return
			}
			row[i] = num
		}

		if isSafeRow(row) {
			fmt.Printf("%d Safe: %v\n", lineNumber, row)
			safeCount++
		} else if problemDamper(row) {
			fmt.Printf("%d Safe: %v\n", lineNumber, row)
			safeCount++
		} else {
			fmt.Printf("%d Unsafe: %v\n", lineNumber, row)
		}
		lineNumber++
	}

	fmt.Printf("Reports safe: %v\n", safeCount)

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
