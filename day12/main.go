package main

import (
	"fmt"
	"strconv"

	"github.com/kongebra/adventofcode2020/utils"
)


const (
	filename = "input.txt"

	north   = 'N'
	south   = 'S'
	east    = 'E'
	west    = 'W'
	left    = 'L'
	right   = 'R'
	forward = 'F'
)

var (
	input []string
)

func init() {
	var err error
	input, err = utils.ReadFileToStrings(filename)
	if err != nil {
		panic(err)
	}
}

func main() {
	partOne()
}

func partOne() {
	facing := east
	eastPosition := 0
	northPosition := 0

	for _, instruction := range input {
		action := rune(instruction[0])
		value, err := strconv.Atoi(instruction[1:])
		if err != nil {
			panic(err)
		}

		f, e, n := move(facing, action, value)
		facing = f
		eastPosition += e
		northPosition += n

		// fmt.Printf("%s\tfacing: %c\teast: %d\tnorth: %d\n", instruction, facing, eastPosition, northPosition)
	}

	eastPosition = utils.AbsInt(eastPosition)
	northPosition = utils.AbsInt(northPosition)

	sum := eastPosition + northPosition

	fmt.Println("What is the Manhattan distance between that location and the ship's starting position?", sum)
}

func move(facing rune, action rune, value int) (rune, int, int) {
	switch action {
	case north:
		return facing, 0, value

	case south:
		return facing, 0, -value

	case east:
		return facing, value, 0

	case west:
		return facing, -value, 0

	case left:
		return turnLeft(facing, value), 0, 0

	case right:
		return turnRight(facing, value), 0, 0

	case forward:
		e, n := moveForward(facing, value)
		return facing, e, n

	default:
		fmt.Println("ERROR, UNKNOWN ACTION:", action)
		break
	}

	return ' ', 0, 0
}

func turnLeft(facing rune, value int) rune {
	steps := value / 90
	
	result := facing

	for i := 0; i < steps; i++ {
		switch result {
		case north:
			result = west
			break

		case south:
			result = east
			break

		case east:
			result = north
			break

		case west:
			result = south
			break
		}
	}

	return result
}

func turnRight(facing rune, value int) rune {
	steps := value / 90
	
	result := facing

	for i := 0; i < steps; i++ {
		switch result {
		case north:
			result = east
			break

		case south:
			result = west
			break

		case east:
			result = south
			break

		case west:
			result = north
			break
		}
	}

	return result
}

func moveForward(facing rune, value int) (int, int) {
	switch facing {
	case north:
		return 0, value

	case south:
		return 0, -value

	case east:
		return value, 0

	case west:
		return -value, 0

	default:
		return 0, 0
	}
}