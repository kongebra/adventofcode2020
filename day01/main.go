package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Metode:

// Ta index 0 (eksempel 1721)
// Trekk den fra 2020 (reset = 299)
// Søk etter rest

// Hvis ikke finnes gå til index 2, gjør det samme, gjør det samme


func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}

	task01(input)

	task02(input)
}

func getInput() ([]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	var result []int

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		result = append(result, x)
	}

	return result, nil
}

func task01(input []int) {
	for i := 0; i < len(input); i++ {
		firstValue := input[i]
		restValue := 2020 - firstValue

		for j := i + 1; j < len(input); j++ {
			secondValue := input[j]

			if (restValue == secondValue) {
				product := firstValue * secondValue
				fmt.Printf("[%d, %d], product = %d\n", firstValue, secondValue, product)
				return
			}
		}
	}
}

func task02(input []int) {
	for i := 0; i < len(input); i++ {
		firstValue := input[i]

		for j := i + 1; j < len(input); j++ {
			secondValue := input[j]
			restValue := 2020 - firstValue - secondValue

			for k := j +1; k < len(input); k++ {
				thirdValue := input[k]

				if (thirdValue == restValue) {
					product := firstValue * secondValue * thirdValue
					fmt.Printf("[%d, %d, %d], product = %d\n", firstValue, secondValue, thirdValue, product)
					return
				}
			}
		}
	}
}