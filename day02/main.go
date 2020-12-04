package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type passwordLine struct {
	Min int
	Max int
	Char string
	Password string
}

func (p passwordLine) valid() bool {
	count := strings.Count(p.Password, p.Char)
	
	if count >= p.Min && count <= p.Max {
		return true
	}

	return false
}

func (p passwordLine) validTwo() bool {
	minChar := string(p.Password[p.Min - 1])
	maxChar := string(p.Password[p.Max - 1])

	// kan ikke være like
	if (minChar == maxChar) {
		return false
	}

	// ingen av de er riktig
	if minChar == p.Char || maxChar == p.Char {
		return true
	}

	return false
}

func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}

	validCount := 0

	for _, value := range input {
		if (value.validTwo()) {
			validCount++
		}
	}

	fmt.Printf("Valid count: %d\n", validCount)
}

func getInput() ([]passwordLine, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	var result []passwordLine

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		splitOne := strings.Split(line, ":")
		password := splitOne[1][1:]

		splitTwo := strings.Split(splitOne[0], " ")
		char := splitTwo[1]

		splitThree := strings.Split(splitTwo[0], "-")
		
		min, err := strconv.Atoi(splitThree[0])
		if err != nil {
			return nil, err
		}
		
		max, err := strconv.Atoi(splitThree[1])
		if err != nil {
			return nil, err
		}

		result = append(result, passwordLine{
			Min: min,
			Max: max,
			Char: char,
			Password: password,
		})
	}

	return result, nil
}