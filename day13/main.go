package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/kongebra/adventofcode2020/utils"
)

type bus struct {
	ID int
	Departures []int
}

const (
	filename = "input.txt"
)

var (
	input []string

	earliestTimestamp string
	busIDs []string
)

func init() {
	var err error
	input, err = utils.ReadFileToStrings(filename)
	if err != nil {
		panic(err)
	}

	earliestTimestamp = input[0]
	busIDs = strings.Split(input[1], ",")
}

func main() {
	partOne()
}

func partOne() {
	earliestDeparture, err := strconv.Atoi(earliestTimestamp)
	if err != nil {
		panic(err)
	}

	busses := make([]bus, 0)

	for _, busID := range busIDs {
		if busID != "x" {
			id, err := strconv.Atoi(busID)
			if err != nil {
				panic(err)
			}

			busDepartures := make([]int, 0)
			for i:= 0; i < earliestDeparture * 2; i += id {
				if (i >= earliestDeparture) {
					busDepartures = append(busDepartures, i)
				}
			}

			bus := bus{
				ID: id,
				Departures: busDepartures,
			}

			busses = append(busses, bus)
		}
	}

	lowestID := math.MinInt32
	lowest := math.MaxInt32

	for _, bus := range busses {
		for _, depature := range bus.Departures {
			if (depature > earliestDeparture) {
				if (depature < lowest) {
					lowestID = bus.ID
					lowest = depature
				}
			}
		}
	}

	result := (lowest - earliestDeparture) * lowestID

	fmt.Println("What is the ID of the earliest bus you can take to the airport multiplied by the number of minutes you'll need to wait for that bus?", result)
}