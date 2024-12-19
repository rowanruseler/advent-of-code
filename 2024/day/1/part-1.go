package main

import (
        "bufio"
        "fmt"
        "math"
        "os"
        "sort"
        "strconv"
        "strings"
)

func main() {
        file, err := os.Open("input/puzzle_part1.txt")
        if err != nil {
                fmt.Println("Error opening file:", err)
                return
        }
        defer file.Close()
        
        var left, right []int
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

                left = append(left, leftNum)
                right = append(right, rightNum)
        }

        if err := scanner.Err(); err != nil {
                fmt.Println("Error reading file:", err)
                return
        }

        sort.Ints(left)
        sort.Ints(right)

        sum := 0
        distances := make([]int, len(left))
        for i := range left {
                distances[i] = int(math.Abs(float64(left[i] - right[i])))
                sum += distances[i]
        }

        fmt.Println("Sorted left numbers:", left)
        fmt.Println("Sorted right numbers:", right)
        fmt.Println("Distances:", distances)
        fmt.Println("Sum of distances:", sum)
}
