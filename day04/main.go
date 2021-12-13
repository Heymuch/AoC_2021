package main

import (
	f "fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	drawn []int
	boards []Board
}

type Board struct {
	values [5 * 5]int
	checked [5 * 5]bool
}

func (b *Board) Mark(value int) {
	for i, l := 0, len(b.values); i < l; i++ {
		if b.values[i] == value {
			b.checked[i] = true
		}
	}
}

func (b *Board) Check() (bool) {
	for i := 0; i < 5; i++ {
		bingo := true
		for j := 0; j < 5; j++ {
			if !b.checked[i * 5 + j] {
				bingo = false
			}
		}

		if bingo {
			return true
		}

		bingo = true
		for j := 0; j < 5; j++ {
			if !b.checked[j * 5 + i] {
				bingo = false
			}
		}

		if bingo {
			return true
		}
	}

	return false
}

func (b *Board) SumUnchecked() (int) {
	var sum int = 0

	for i, l := 0, len(b.values); i < l; i++ {
		if !b.checked[i] {
			sum += b.values[i]
		}
	}

	return sum
}

func (g *Game) Mark(value int) {
	for i, l := 0, len(g.boards); i < l; i++ {
		g.boards[i].Mark(value)
	}
}

func (g *Game) Check() (bool, int) {
	for i, l := 0, len(g.boards); i < l; i++ {
		if g.boards[i].Check() {
			return true, i
		}
	}
	return false, 0
}

func readInput(filename string) ([]string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		f.Printf("error during reading file :: %s\n", err)
		os.Exit(1)
	}
	return strings.Split(string(bytes), "\n")
}

func fromText(input []string) (*Game) {
	var game Game;

	game.drawn = drawnFromText(input[0])
	game.boards = boardsFromText(input[2:])

	return &game
}

func drawnFromText(input string) ([]int) {
	tokens := strings.Split(input, ",")
	drawn := make([]int, len(tokens))

	for i, l := 0, len(tokens); i < l; i++ {
		value, err := strconv.Atoi(tokens[i])
		if err != nil {
			f.Printf("error during int parsing :: %s\n", err)
		}
		drawn[i] = value
	}

	return drawn
}

func boardsFromText(input []string) ([]Board) {
	boards := make([]Board, 0)

	for i, l := 0, len(input); i < l; i += 6 {
		var b Board
		for j := 0; j < 5; j++ {
			f.Sscanf(input[i + j], "%2d %2d %2d %2d %2d",
				&b.values[j * 5 + 0],  &b.values[j * 5 + 1],  &b.values[j * 5 + 2],  &b.values[j * 5 + 3],  &b.values[j * 5 + 4])
		}
		boards = append(boards, b)
	}

	return boards
}

func main() {
	f.Println("Hello from Advent of Code Day 4!")

	input := readInput("./input.txt")

	f.Println("-- First Part --")
	FirstPart(input)
	f.Println("-- Second Part --")
	SecondPart(input)
}

func FirstPart(input []string) {
	game := fromText(input)

	for i, l := 0, len(game.drawn); i < l; i++ {
		game.Mark(game.drawn[i])

		bingo, index := game.Check()
		if bingo {
			f.Println("BINGO!")
			f.Printf("Board index %d, Sum of Unchecked: %d\n", index, game.boards[index].SumUnchecked())
			f.Printf("Result: %d\n", (game.boards[index].SumUnchecked() * game.drawn[i]))
			return
		}
	}
}

func SecondPart(input []string) {
	game := fromText(input)

	var lastBoard Board
	var lastValue int

	for i, l := 0, len(game.drawn); i < l; i++ {
		game.Mark(game.drawn[i])
		bingo, index := game.Check()

		if bingo {
			lastBoard = game.boards[index]
			lastValue = game.drawn[i]
			game.boards = append(game.boards[:index], game.boards[index + 1:]...)
		}
	}
	f.Println("BINGO!")
	f.Printf("Sum of Unchecked: %d\n", lastBoard.SumUnchecked())
	f.Printf("Result: %d\n", (lastBoard.SumUnchecked() * lastValue))
}