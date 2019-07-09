// Package advent provides basic utilities for running solutions to Advent of Code
// puzzles.
package advent

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// intday is the integral type used to represent the day of a puzzle.
type intday uint8

// intyear is the integral type used to represent the year of a puzzle.
type intyear uint16

// puzzleError is the error type returned when a Puzzle cannot be parsed or
// solved correctly.
type puzzleError string

func (e puzzleError) Error() string {
	return fmt.Sprint("advent: cannot solve puzzle - ", string(e))
}

// A Puzzle represents a distinct daily puzzle released during the Advent of Code,
// along with its associated input.
type Puzzle struct {
	day   intday
	year  intyear
	input []byte
}

// ParsePuzzleFromArgs constructs a Puzzle instance from the provided command line
// arguments.
func ParsePuzzleFromArgs(args []string) (Puzzle, error) {
	if len(args) < 2 || len(args) > 3 {
		return Puzzle{}, puzzleError("invalid number of arguments")
	}

	year, err := strconv.ParseInt(args[0], 10, 16)
	if err != nil {
		return Puzzle{}, puzzleError("puzzle year could not be parsed")
	}

	day, err := strconv.ParseInt(args[1], 10, 8)
	if err != nil {
		return Puzzle{}, puzzleError("puzzle day could not be parsed")
	}
	if day < 1 || day > 25 {
		return Puzzle{}, puzzleError("puzzle day must be in [1,25]")
	}

	var input []byte

	if len(args) > 2 {
		filename := args[2]
		if filename == "--" {
			input, err = ioutil.ReadAll(os.Stdin)
		} else {
			input, err = ioutil.ReadFile(filename)
		}
	} else {
		input, err = readDefaultInputFile(intyear(year), intday(day))
	}

	if err != nil {
		return Puzzle{}, err
	}

	return Puzzle{intday(day), intyear(year), input}, nil
}

// readDefaultInputFile reads the default input file for the puzzle associated with the
// specified event year and day.
func readDefaultInputFile(year intyear, day intday) ([]byte, error) {
	return ioutil.ReadFile(fmt.Sprintf("resources/y%d/day%02d.txt", year, day))
}

// A Solution represents an answer to each part of Advent of Code puzzle.
type Solution struct {
	Part1 string
	Part2 string
}

// SolutionFromInts constructs a Solution from a pair of integers that represent the solutions
// to each part of a puzzle.
func SolutionFromInts(part1, part2 int) Solution {
	return Solution{strconv.Itoa(part1), strconv.Itoa(part2)}
}

// A SolutionSet holds the functions that solve each of the 25 puzzles released during
// an Advent of Code event.
type SolutionSet [25]func([]byte) (Solution, error)

// A PuzzleSolver holds all of the solutions produced for Advent of Code puzzles and will run
// a solver function for a given Puzzle instance.
type PuzzleSolver []struct {
	Year      intyear
	Solutions *SolutionSet
}

// SolvePuzzle runs the solution function associated with the specified Puzzle.
func (solver PuzzleSolver) SolvePuzzle(puzzle Puzzle) (Solution, error) {
	yearIndex := -1
	for i, set := range solver {
		if set.Year == puzzle.year {
			yearIndex = i
		}
	}
	if yearIndex < 0 {
		return Solution{}, puzzleError(fmt.Sprintf("no solutions exist for year %d", puzzle.year))
	}

	f := solver[yearIndex].Solutions[puzzle.day-1]
	if f == nil {
		return Solution{}, puzzleError(fmt.Sprintf("no solutions exist for day %d:%d", puzzle.year, puzzle.day))
	}
	return f(puzzle.input)
}
