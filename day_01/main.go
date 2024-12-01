package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func diff(a, b int) int {
	// From StackOverflow
	// https://stackoverflow.com/a/59453929
	if a < b {
		return b - a
	}
	return a - b
}

func readFileContents(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var leftNums, rightNums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "   ")
		if len(parts) == 2 {
			leftNum := 0
			rightNum := 0
			fmt.Sscan(parts[0], &leftNum)
			fmt.Sscan(parts[1], &rightNum)

			leftNums = append(leftNums, leftNum)
			rightNums = append(rightNums, rightNum)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return leftNums, rightNums, err
}

func calculateDifference(left, right []int) (int, error) {
	sort.Ints(left)
	sort.Ints(right)

	sum := 0

	for i := 0; i < len(left); i++ {
		sum += diff(left[i], right[i])
	}

	return sum, nil
}

func main() {
	left, right, err := readFileContents("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	sum, err := calculateDifference(left, right)
	if err != nil {
		fmt.Printf("Error: ", err)
		os.Exit(1)
	}

	fmt.Printf("Sum of differences: %d\n", sum)
}
