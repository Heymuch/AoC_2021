package main

import (
	f "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f.Println("Hello from Advent of Code - Day 2")
	f.Println("-- First Part --")
	FirstPart()
	f.Println("-- Second Part --")
	SecondPart()
}

func FirstPart() {
	bytes, err := os.ReadFile("./input.txt")
	if err != nil {
		f.Printf("error during reading file :: %s\n", err)
		os.Exit(1)
	}

	data := strings.Split(string(bytes), "\n")
	f.Printf("Number of records: %d\n", len(data))
	hPos, vPos := 0, 0
	for i, l := 0, len(data); i < l; i++ {
		tokens := strings.Split(data[i], " ")
		instruction := tokens[0]
		value, err := strconv.Atoi(tokens[1])
		if err != nil {
			f.Printf("error during conversion of string to int :: %s\n", err)
			os.Exit(1)
		}

		switch instruction {
		case "forward":
			hPos += value
		case "up":
			vPos -= value
		case "down":
			vPos += value
		}
	}

	f.Printf("H position: %d; V position: %d\n", hPos, vPos)
	f.Printf("My answear: %d\n", (hPos * vPos))
}

func SecondPart() {
	bytes, err := os.ReadFile("./input.txt")
	if err != nil {
		f.Printf("error during reading file :: %s\n", err)
		os.Exit(1)
	}

	data := strings.Split(string(bytes), "\n")
	f.Printf("Number of records: %d\n", len(data))
	hPos, vPos, aim := 0, 0, 0
	for i, l := 0, len(data); i < l; i++ {
		tokens := strings.Split(data[i], " ")
		instruction := tokens[0]
		value, err := strconv.Atoi(tokens[1])
		if err != nil {
			f.Printf("error during conversion of string to int :: %s\n", err)
			os.Exit(1)
		}

		switch instruction {
		case "forward":
			hPos += value
			vPos += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}

	f.Printf("H position: %d; V position: %d\n", hPos, vPos)
	f.Printf("My answear: %d\n", (hPos * vPos))
}