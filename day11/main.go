package main

import (
	"fmt"

	"github.com/kongebra/adventofcode2020/utils"
)


const (
	filename = "input.txt"
	floor = '.'
	emptySeat = 'L'
	occupiedSeat = '#'
)

var (
	input []string
	initialSeats [][]rune
)

func init() {
	var err error
	input, err = utils.ReadFileToStrings(filename)
	if err != nil {
		panic(err)
	}

	initialSeats = make([][]rune, 0)
	// Convert to 2D Rune Slice
	for row, seats := range input {
		initialSeats = append(initialSeats, []rune{})

		for _, seat := range seats {
			initialSeats[row] = append(initialSeats[row], seat)
		}
	}
}

func main() {
	partOne()
}

func partOne() {
	seatMap := runRulesPartOne(initialSeats)

	for {
		next := runRulesPartOne(seatMap)

		if !isDifferent(seatMap, next) {
			break
		}

		seatMap = next;
	}
	
	count := countOccupiedSeats(seatMap)
	fmt.Println("How many seats end up occupied?", count)
}

func countOccupiedSeats(seatMap [][]rune) int {
	count := 0

	for _, seats := range seatMap {
		for _, seat := range seats {
			if seat == occupiedSeat {
				count++
			}
		}
	}

	return count
}

func isDifferent(a, b [][]rune) bool {
	if len(a) != len(b) {
		return true
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return true
		}

		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return true
			}
		}
	}

	return false
}

func runRulesPartOne(seatMap [][]rune) [][]rune {
	duplicate := make([][]rune, len(seatMap))
	for i := range seatMap {
		duplicate[i] = make([]rune, len(seatMap[i]))
		copy(duplicate[i], seatMap[i])
	}

	for row, seats := range seatMap {
		for col, seat := range seats {
			adjacent := getAdjacentSeats(seatMap, col, row)

			if seat == emptySeat {
				if !occupiedSeatInAdjacent(adjacent) {
					duplicate[row][col] = occupiedSeat
				}
			} else if seat == occupiedSeat {
				if isFourOrMoreAdjacentOccupied(adjacent) {
					duplicate[row][col] = emptySeat
				}
			}


		}
	}

	return duplicate
}

func getAdjacentSeats(seatMap [][]rune, col, row int) []rune {
	adjacent := make([]rune, 0)

	minRow := 0
	maxRow := 0
	minCol := 0
	maxCol := 0

	if row > 0 {
		minRow = -1
	}

	if col > 0 {
		minCol = -1
	}

	if row < len(input) - 1 {
		maxRow = 1
	}

	if col < len(input[row]) - 1 {
		maxCol = 1
	}

	for r := minRow; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			if r != 0 || c != 0 {
				adjacent = append(adjacent, rune(seatMap[row+r][col+c]))
			}
		}
	}

	return adjacent
}

func occupiedSeatInAdjacent(adjacent []rune) bool {
	for _, seat := range adjacent {
		if seat == occupiedSeat {
			return true
		}
	}

	return false
}

func isFourOrMoreAdjacentOccupied(adjacent []rune) bool {
	count := 0

	for _, seat := range adjacent {
		if seat == occupiedSeat {
			count++
		}
	}

	return count >= 4
}