package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
        "strings"
)

func main() {
        file, err := os.Open("input/puzzle_part2.txt")
        if err != nil {
                fmt.Println("Error opening file:", err)
                return
        }
        defer file.Close()
        
        var numbers [][]int
        scanner := bufio.NewScanner(file)
        
        for scanner.Scan() {
                line := scanner.Text()
                parts := strings.Fields(line)
                if len(parts) != 2 {
                    fmt.Println("Skipping invalid line:", line)
                    continue
                }
        
                leftNum, err0 := strconv.Atoi(parts[0])
                rightNum, err1 := strconv.Atoi(parts[1])
                if err0 != nil || err1 != nil {
                    fmt.Println("Error converting line to integers:", line)
                    return
                }
        
                numbers = append(numbers, []int{leftNum, rightNum})
        }
    
        if err := scanner.Err(); err != nil {
                fmt.Println("Error reading file:", err)
                return
        }
    
        rightFreq := make(map[int]int)
        for _, pair := range numbers {
                rightFreq[pair[1]]++
        }
    
        totalScore := 0
        for _, pair := range numbers {
                leftNum := pair[0]
                if count, exists := rightFreq[leftNum]; exists {
                    totalScore += leftNum * count
                }
        }
    
        fmt.Println("Total similarity score:", totalScore)
}
