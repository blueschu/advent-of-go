package main

import (
	"fmt"
	"os"

	"github.com/blueschu/advent-of-go/advent"
)

func main() {
	puzzleSelection, err := advent.ParsePuzzleFromArgs(os.Args[1:])

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		printUsage(os.Args[0])
		os.Exit(1)
	}

	var solver = advent.PuzzleSolver{}

	result, err := solver.SolvePuzzle(puzzleSelection)

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Part 1: %v\nPart 2: %v\n", result.Part1, result.Part2)
}

func printUsage(prog string) {
	fmt.Printf("Usage: %v year day [input]\n", prog)
}
