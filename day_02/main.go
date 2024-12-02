package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFileInput(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: {err}")
	}
	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report := strings.Fields(scanner.Text())
		sequence := make([]int, 0, len(report))
		for _, numStr := range report {
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

func main() {

}
