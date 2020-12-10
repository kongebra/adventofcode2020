package utils

import (
	"bufio"
	"os"
	"strconv"
)

func readFileToStrings(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		result = append(result, line)
	}

	return result, nil
}

func readFileToInts(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		value, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		result = append(result, value)
	}

	return result, nil
}