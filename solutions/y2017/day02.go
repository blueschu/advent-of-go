package y2017

import (
	"bytes"
	"fmt"
	"math"
	"strconv"

	"github.com/blueschu/advent-of-go/advent"
)

func SolveDay02(input []byte) (advent.Solution, error) {
	const RowLen = 16

	lines := bytes.Split(bytes.TrimSpace(input), []byte("\n"))

	spreadsheet, err := parseSpreadsheet(lines, RowLen)
	if err != nil {
		return advent.Solution{}, err
	}

	var (
		lineChecksum         = 0
		divisibilityChecksum = 0
	)

	for _, line := range spreadsheet {
		lineChecksum += computeLineChecksum(line)
		divisibilityChecksum += computeDivisibilityChecksum(line)
	}

	return advent.SolutionFromInts(lineChecksum, divisibilityChecksum), nil
}

func parseSpreadsheet(lines [][]byte, rowLen int) ([][]int, error) {
	var valueLines = make([][]int, len(lines))
	// Preallocate slice to hold spreadsheet values
	var values = make([]int, len(lines)*rowLen)

	for i, line := range lines {
		cells := bytes.Split(line, []byte("\t"))

		if len(cells) != rowLen {
			return nil, fmt.Errorf("bad input - row %d of spreadsheet contained %d entires, expected %d", i, len(cells), rowLen)
		}

		// Write spreadsheet values to the next rowLen indices
		for j, digits := range cells {
			value, err := strconv.Atoi(string(digits))
			if err != nil {
				return nil, fmt.Errorf("bad input - %v is no a number", digits)
			}
			values[j] = value
		}
		valueLines[i] = values[:rowLen]
		values = values[rowLen:]
	}
	return valueLines, nil
}

func computeLineChecksum(line []int) int {
	var (
		min = math.MaxUint32
		max = 0
	)
	for _, b := range line {
		if b < min {
			min = b
		}
		if b > max {
			max = b
		}
	}
	return max - min
}

func computeDivisibilityChecksum(line []int) (checksum int) {
	for i, finger := range line[:len(line)-1] {
		for _, b := range line[i+1:] {
			if b%finger == 0 {
				checksum += b / finger
			} else if finger%b == 0 {
				checksum += finger / b
			}
		}
	}
	return
}
