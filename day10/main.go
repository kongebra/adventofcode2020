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
	// partOne()
	partTwo()
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

type node struct {
	Value int
	Children []*node
}

func newNode(value int) *node {
	return &node{
		Value: value,
		Children: make([]*node, 0),
	}
}

func partTwo() {
	set := make(map[int]*node, 0)

	for _, value := range input {
		// not in set
		if _, ok := set[value]; !ok {
			set[value] = newNode(value)
		}
	}
	
	// make all Children
	for k, v := range set {
		// get node if exists
		for i := 1; i <= 3; i++ {
			if n, ok := set[k+i]; ok {
				v.Children = append(v.Children, n)
			}
		}
	}

	min := utils.FindMin(input)
	root, ok := set[min]
	if !ok {
		panic("panic at the disco")
	}

	// count := 0
	// for _, child := range root.Children {
	// 	if (len(child.Children) == 0) {
	// 		count++
	// 	}
	// }

	paths := 1
	search(root, func(n int) {
		fmt.Println(n)
	})
	
	fmt.Println("What is the total number of distinct ways you can arrange the adapters to connect the charging outlet to your device?", paths)
}

var (
	visited map[int]bool = make(map[int]bool)
)

func search(n *node, cb func(int)) {
	visited[n.Value] = true
	cb(n.Value)

	for _, v := range n.Children {
		if visited[v.Value] {
			continue
		}

		search(v, cb)
	}
}