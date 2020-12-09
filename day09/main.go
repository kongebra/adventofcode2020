package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	filename = "input.txt"
)

var (
	input []int
	preamble []int
)

func init() {
	input = make([]int, 0)
	preamble = make([]int, 0)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		input = append(input, num)

		if len(preamble) != 25 {
			preamble = append(preamble, num)
		}		
	}
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	invalid := -1

	for _, value := range input[25:] {
		valid := false

		for _, sub := range preamble {
			rest := value - sub

			if rest < 0 {
				continue
			}

			for _, check := range preamble {
				if check == sub {
					continue
				}

				if rest == check {
					preamble = append(preamble, value)
					valid = true
					break
				}
			}

			if valid {
				break
			}
		}

		if !valid {
			invalid = value
			break
		}
	}

	fmt.Println("What is the first number that does not have this property?", invalid)
}

func partTwo() {

}