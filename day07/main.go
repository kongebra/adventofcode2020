package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	Color string
	Children map[string]child
	Parents []*node
}

type child struct {
	Node *node
	Amount int
}

func newNode(color string) *node {
	return &node{
		Color: color,
		Children: make(map[string]child, 0),
		Parents: make([]*node, 0),
	}
}

func makeChild(node *node, amount int) child {
	return child{
		Node: node,
		Amount: amount,
	}
}

const (
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
	set := make(map[string]*node, 0)

	for _, line := range input {
		// split string
		arr := strings.Split(line, " bags contain ")

		// get color of current bag
		color := arr[0]

		// check if it is already made
		if _, ok := set[color]; !ok {
			set[color] = newNode(color)
		}

		// pointer to parent
		parent := set[color]

		// get children string
		childrenString := arr[1]
		
		// clean string
		childrenString = strings.ReplaceAll(childrenString, " bags", "")
		childrenString = strings.ReplaceAll(childrenString, " bag", "")
		childrenString = strings.ReplaceAll(childrenString, ".", "")
		
		// split children-string
		childrenArr := strings.Split(childrenString, ", ")

		// loop over
		for _, childStr := range childrenArr {
			// contain no other bags
			if childStr == "no other" {
				break
			}

			// split string
			split := strings.Split(childStr, " ")
			
			// amount of childs
			childAmount, err := strconv.Atoi(split[0])
			if err != nil {
				panic(err)
			}

			// color of the child
			childColor := strings.Join(split[1:], " ")

			// check if it is already made
			if _, ok := set[childColor]; !ok {
				set[childColor] = newNode(childColor)
			}

			// pointer to child
			child := set[childColor]

			// set parent
			child.Parents = append(child.Parents, parent)

			// set children
			parent.Children[childColor] = makeChild(child, childAmount)
		}
	}

	sg := set["shiny gold"]
	
	parents := partOne(sg)
	children := partTwo(sg)


	fmt.Println("Part One", len(parents))
	fmt.Println("Part Two", children)
}

func partOne(child *node) map[string]*node {
	result := make(map[string]*node, 0)

	for _, parent := range child.Parents {
		if _, ok := result[parent.Color]; !ok {
			result[parent.Color] = parent

			if (len(parent.Parents) > 0) {
				grandParents := partOne(parent)
				for grandParentColor, grandParent := range grandParents {
					result[grandParentColor] = grandParent
				}
			}
		}
	}

	return result
}

func partTwo(parent *node) int {
	count := 0

	for _, child := range parent.Children {
		amount := partTwoHelper(child.Amount, child.Node)

		count += amount
	}
	
	return count
}

func partTwoHelper(multiplier int, child *node) int {
	// counting it self
	count := 1

	for _, grandChild := range child.Children {
		amount := partTwoHelper(grandChild.Amount, grandChild.Node)

		count += amount
	}

	return multiplier * count
}