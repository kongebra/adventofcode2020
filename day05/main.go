package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}

	_ = input

	fmt.Println(getColumn("BFFFBBFRRR"))
	fmt.Println(getColumn("FFFBBBFRRR"))
	fmt.Println(getColumn("BBFFBBFRLL"))
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

func getRow(input string) int {
	rowSection := input[:7]
	instructions := strings.Split(rowSection, "")

	row := 0
	front := 0
	back := 127

	for _, instruction := range instructions {
		// fmt.Printf("BEFORE:\t\tinstruction: %s\t\tfront: %d\tback: %d\n", instruction, front, back)

		if (instruction == "F") {
			back = (front + back) / 2
		} else if (instruction == "B") {
			front = (front + back + 1) / 2
		} else {
			fmt.Println("ERROR: Unknown instruction:", instruction)
		}

		// fmt.Printf("AFTER:\t\tinstruction: %s\t\tfront: %d\tback: %d\n\n", instruction, front, back)
	}

	row = front

	return row
}

func getColumn(input string) int {
	columnSection := input[7:]
	instructions := strings.Split(columnSection, "")

	column := 0
	left := 0
	right := 7

	for _, instruction := range instructions {
		if (instruction == "R") {
			left = (left + right + 1) / 2
		} else if (instruction == "L") {
			right = (left + right) / 2
		} else {
			fmt.Println("ERROR: Unknown instruction:", instruction)
		}
	}

	column = left

	return column
}
