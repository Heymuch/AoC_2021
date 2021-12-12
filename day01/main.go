package main

import (
	f "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	SecondPart()
}

func FirstPart() {
	bytes, err := os.ReadFile("./input.txt")
	if err != nil {
		f.Printf("error during reading file :: %s\n", err)
	}
	data := strings.Split(string(bytes), "\n")
	f.Printf("Number of records: %d\n", len(data))

	var numInc int = 0
	for i, l := 1, len(data); i < l; i++ {
		prev, _ := strconv.Atoi(data[i - 1])
		acc, _ := strconv.Atoi(data[i])
		if prev < acc {
			numInc++
		}
	}
	f.Printf("Number of increases: %d\n", numInc)
}

func SecondPart() {
	bytes, err := os.ReadFile("./input.txt")
	if err != nil {
		f.Printf("error during reading file :: %s\n", err)
	}
	dataStr := strings.Split(string(bytes), "\n")
	data := make([]int, len(dataStr))
	for i, l := 0, len(dataStr); i < l; i++ {
		data[i], _ = strconv.Atoi(dataStr[i])
	}

	f.Printf("Number of records: %d\n", len(data))

	var numInc int = 0
	for i, l := 3, len(data); i < l; i++ {
		prev := data[i - 3] + data[i - 2] + data[i - 1]
		acc := data[i - 2] + data[i - 1] + data[i]
		if prev < acc {
			numInc++
		}
	}
	f.Printf("Number of increases: %d\n", numInc)
}