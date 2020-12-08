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

	// partOne(input)
	partTwo()
}

func getInput() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	var result []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var group []string

	for scanner.Scan() {
		line := scanner.Text()

		if (line == "") {
			// make group to one string
			var groupString string
			groupString = strings.Join(group, "")

			result = append(result, groupString)

			// reset group slice
			group = make([]string, 0)
		} else {
			group = append(group, line)
		}
	}

	return result, nil
}

func getInputPartTwo() ([][]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	var result [][]string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var group []string

	for scanner.Scan() {
		line := scanner.Text()

		if (line == "") {
			result = append(result, group)

			group = make([]string, 0)
		} else {
			group = append(group, line)
		}
	}

	return result, nil
}

func unique(line string) []string {
	slice := strings.Split(line, "")

	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

func partOne(input []string) {
	count := 0

	for _, value := range input {
		uniqueValues := unique(value)
		uniqueCount := len(uniqueValues)

		count += uniqueCount
	}

	fmt.Println("What is the sum of those counts?", count)
}

func partTwo() {
	input, err := getInputPartTwo()
	if err != nil {
		panic(err)
	}

	count := 0

	// Go through each group
	for _, group := range input {
		// only one in the group
		if len(group) == 1 {
			count += len(group[0])
		} else {
			// combine all answers in the group
			joinedAnswers := strings.Join(group, "")
			uniqueAnswers := unique(joinedAnswers)

			// go through all unqiue answers
			for _, answer := range uniqueAnswers {
				// count amount of answers in joined answers
				answerCount := strings.Count(joinedAnswers, answer)
	
				// Check if it occures more then 1 time
				if answerCount > 1 {
					count++
				}
			}
		}
	}

	// 3832 (too high)
	fmt.Println("What is the sum of those counts?", count)
}