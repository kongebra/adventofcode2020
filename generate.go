package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	today := time.Now()
	_, _, day := today.Date()
	dir := fmt.Sprintf("day%02d", day)
	test := dir + "/test.txt"
	input := dir + "/input.txt"
	mainFile := dir + "/main.go"

	if notExists(dir) {
		os.Mkdir(dir, os.ModePerm)
	} else {
		fmt.Println("Directory for this day already created!")
	}

	createFile(test)
	createFile(input)
	createFile(mainFile)
}

func createFile(path string) {
	if notExists(path) {
		if err := ioutil.WriteFile(path, nil, os.ModePerm); err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("'%s' already exists\n", path)
	}
}

func notExists(path string) bool {
	_, err := os.Stat(path)

	return os.IsNotExist(err)
}