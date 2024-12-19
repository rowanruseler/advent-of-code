package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input/puzzle_part1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	xmas := "XMAS"
	xmasLength := len(xmas)
	rows := len(grid)
	cols := len(grid[0])

	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}

	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				found := true
				for k := 0; k < xmasLength; k++ {
					nr := r + dx*k
					nc := c + dy*k
					if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != xmas[k] {
						found = false
						break
					}
				}
				if found {
					count++
				}
			}
		}
	}
	fmt.Println("XMAS Count", count)
}
