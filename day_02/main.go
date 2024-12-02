package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func diff(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func directionChange(levels []int) bool {
	prevDiff := levels[1] - levels[0]

	for i := 1; i < len(levels)-1; i++ {
		currDiff := levels[i+1] - levels[i]
		if prevDiff*currDiff < 0 {
			return true
		}
		prevDiff = currDiff
	}

	return false
}

func readFileContent(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: {err}")
	}
	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		levels := strings.Fields(scanner.Text())
		sequence := make([]int, 0, len(levels))
		for _, numStr := range levels {
			num := 0
			fmt.Sscan(numStr, &num)
			sequence = append(sequence, num)
		}

		reports = append(reports, sequence)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func checkSafetyScore(reports [][]int) (int, error) {
	safetyScore := 0
	for _, levels := range reports {
		areLevelsSafe := true

		for j := 0; j < len(levels)-1; j++ {
			difference := diff(levels[j], levels[j+1])
			change := directionChange(levels)
			if difference > 3 || difference < 1 || change == true {
				areLevelsSafe = false
				break
			}
		}
		if areLevelsSafe {
			safetyScore += 1
		}
	}
	return safetyScore, nil

}

func main() {
	listOfScores, err := readFileContent("input.txt")
	if err != nil {
		fmt.Println("Error: {err}")
		os.Exit(1)
	}

	score, err := checkSafetyScore(listOfScores)
	if err != nil {
		fmt.Printf("Error: {err}")
		os.Exit(1)
	}

	fmt.Printf("Number of safe reports: %d\n", score)
}
