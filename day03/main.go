package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}

	// partOne(input)
	partTwo(input)
}

func getInput() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	var result []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		result = append(result, line)
	}

	return result, nil
}

func partOne(input []string) {
	x := 0
	y := 0
	treeCount := 0

	for y < len(input) {
		char := fmt.Sprintf("%c", input[y][x])

		if char == "#" {
			treeCount++
		}

		x += 3
		x %= len(input[y])
		y++
	}

	fmt.Println("Tree count:", treeCount)
}

func partTwo(input []string) {
	

	patterns := [][]int {
		// right 1, down 1
		{ 1, 1 },
		// right 3, down 1
		{ 3, 1 },
		// right 5, down 1
		{ 5, 1 },
		// right 7, down 1
		{ 7, 1 },
		// right 1, down 2
		{ 1, 2 },
	}

	var counts []int

	for _, pattern := range patterns {
		x := 0
		y := 0
		treeCount := 0					

		for y < len(input) {
			char := fmt.Sprintf("%c", input[y][x])
	
			if char == "#" {
				treeCount++
			}
	
			x += pattern[0]
			x %= len(input[y])
			y += pattern[1]
		}

		counts = append(counts, treeCount)
	}

	result := 1
	for _, count := range counts {
		result *= count
	}

	fmt.Println("Result:", result)
}