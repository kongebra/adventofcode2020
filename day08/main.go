package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operation string

type instruction struct {
	Operation operation
	Argument int
}

func makeInstruction(line string) instruction {
	instruction := instruction{}

	split := strings.Split(line, " ")

	op := split[0]

	instruction.Operation = operation(op)
	
	prefix := split[1][0]
	argument, err := strconv.Atoi(split[1][1:])
	if err != nil {
		panic(err)
	}

	instruction.Argument = argument

	if prefix == '-' {
		instruction.Argument *= -1
	}

	return instruction
}

const (
	opAcc operation = "acc"
	opJmp operation = "jmp"
	opNop operation = "nop"

	filename = "input.txt"
)

var (
	input []string
)

func init() {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		input = append(input, line)
	}
}

func main() {
	partOne(input)
}

func partOne(input []string) {
	instructions := make([]instruction, 0)

	for _, line := range input {
		instruction := makeInstruction(line)
		instructions = append(instructions, instruction)
	}

	set := make(map[int]bool)

	instruction := 0
	accumulator := 0
	for {
		// we have been here before!
		if _, ok := set[instruction]; ok {
			break
		} else {
			// set value
			set[instruction] = true
		}

		line := instructions[instruction]
		
		switch (line.Operation) {
		case opAcc:
			accumulator += line.Argument
			instruction++
			break

		case opJmp:
			instruction += line.Argument
			break

		case opNop:
			instruction++
			break
		}
	}

	fmt.Println("Accumulator:", accumulator)
}