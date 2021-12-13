package main

import (
	f "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f.Println("Hello from Advent of Code 2021 - Day 3")
	f.Println("-- First Part --")
	FirstPart()
	f.Println("-- Second Part --")
	SecondPart()
}

func FirstPart() {
	input := readInput("./input.txt")
	f.Printf("Number of records: %d\n", len(input))

	gamaBin, epsilonBin := "", ""
	for i, l := 0, len(input[0]); i < l; i++ {
		counter := make([]int, 2)

		for j, k := 0, len(input); j < k; j++ {
			bit, err := strconv.Atoi(string([]rune(input[j])[i]))
			if err != nil {
				f.Printf("error during parsing string to int :: %s\n", err)
				os.Exit(1)
			}

			counter[bit]++
		}

		if counter[0] > counter[1] {
			gamaBin += "0"
			epsilonBin += "1"
		} else {
			gamaBin += "1"
			epsilonBin += "0"
		}
	}

	gama, err := strconv.ParseInt(gamaBin, 2, 32)
	if err != nil {
		f.Printf("error during parsing binary to int :: %s\n", err)
	}
	epsilon, err := strconv.ParseInt(epsilonBin, 2, 32)
	if err != nil {
		f.Printf("error during parsing binary to int :: %s\n", err)
	}

	f.Printf("Gama rate: %d\n", gama)
	f.Printf("Epsilon rate: %d\n", epsilon)
	f.Printf("Power consuption: %d\n", (gama * epsilon))
}

func SecondPart() {
	input := readInput("./input.txt")
	oxygenRating := getOxygenRating(input)
	co2Rating := getCO2Rating(input)

	f.Printf("Oxygen rating: %d\n", oxygenRating)
	f.Printf("CO2 rating: %d\n", co2Rating)
	f.Printf("Life support rating: %d\n", (oxygenRating * co2Rating))
}

func getOxygenRating(input []string) (int) {
	for i := 0; i < len(input[0]); i++ {
		ones := make([]string, 0)
		zeros := make([]string, 0)

		for j, l := 0, len(input); j < l; j++ {
			bit := string([]rune(input[j])[i])
			if bit == "0" {
				zeros = append(zeros, input[j])
			} else {
				ones = append(ones, input[j])
			}
		}

		if len(zeros) > len(ones) {
			input = zeros
		} else {
			input = ones
		}

		if len(input) == 1 {
			rate, err := strconv.ParseInt(input[0], 2, 32)
			if err != nil {
				f.Printf("error during parsing bin to int :: %s\n", err)
				os.Exit(1)
			}
			return int(rate)
		}
	}

	return 0
}

func getCO2Rating(input []string) (int) {
	for i := 0; i < len(input[0]); i++ {
		ones := make([]string, 0)
		zeros := make([]string, 0)

		// f.Println(input)
		for j, l := 0, len(input); j < l; j++ {
			bit := string([]rune(input[j])[i])
			if bit == "0" {
				zeros = append(zeros, input[j])
			} else {
				ones = append(ones, input[j])
			}
		}

		if len(ones) < len(zeros) {
			input = ones
		} else {
			input = zeros
		}

		if len(input) == 1 {
			rate, err := strconv.ParseInt(input[0], 2, 32)
			if err != nil {
				f.Printf("error during parsing bin to int :: %s\n", err)
				os.Exit(1)
			}
			return int(rate)
		}
	}

	return 0
}

func readInput(filename string) ([]string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		f.Printf("error during reading file :: %s\n", err)
		os.Exit(1)
	}
	return strings.Split(string(bytes), "\n")
}