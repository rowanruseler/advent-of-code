package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input/puzzle_part2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	re := regexp.MustCompile(`(mul\((\d+),(\d+)\))|(do\(\))|(don\'t\(\))`)

	mul := true
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			if match == "do()" {
				mul = true
			} else if match == "don't()" {
				mul = false
			} else if strings.HasPrefix(match, "mul(") && mul {
				inside := match[4 : len(match)-1]
				parts := strings.Split(inside, ",")
                                x, _ := strconv.Atoi(parts[0])
                                y, _ := strconv.Atoi(parts[1])
                                total += x * y
			}
		}
	}

	fmt.Println("Instructions produces:", total)
}
