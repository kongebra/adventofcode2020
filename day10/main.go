package main

import (
	"fmt"

	"github.com/kongebra/adventofcode2020/utils"
)


const (
	filename = "input.txt"
)

var (
	input []int
)

func init() {
	var err error
	input, err = utils.ReadFileToInts(filename)
	if err != nil {
		panic(err)
	}
}

func main() {
	partOne()
}

func partOne() {
	max := utils.FindMax(input)

	current := 0
	target := max + 3

	count := 0
	maxCount := 1_000_000

	oneJolts := 0
	threeJolts := 0

	for {
		searchList := make([]int, 0)

		for _, value := range input {
			if value <= current + 3 && value > current {
				searchList = append(searchList, value)
			}
		}

		if len(searchList) > 0 {
			searchMin := utils.FindMin(searchList)

			if current + 1 == searchMin {
				oneJolts++
			}

			if current + 3 == searchMin {
				threeJolts++
			}

			current = searchMin
		}

		if current + 3 >= target {
			if current + 1 == target {
				oneJolts++
			}

			if current + 3 == target {
				threeJolts++
			}

			break
		}

		// Safty measure
		count++
		if count >= maxCount {
			break
		}
	}

	fmt.Println("What is the number of 1-jolt differences multiplied by the number of 3-jolt differences?", oneJolts * threeJolts)
}