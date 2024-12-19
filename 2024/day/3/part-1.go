package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input/puzzle_part1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	re := regexp.MustCompile(`(mul\((\d+),(\d+)\))|(do\(\))|(don\'t\(\))`)
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
		        x, _ := strconv.Atoi(match[2])
		        y, _ := strconv.Atoi(match[3])

		        total += x * y
		}
	}

	fmt.Println("Instructions produces:", total)
}
