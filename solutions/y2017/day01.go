package y2017

import (
	"bytes"
	"errors"
	"unicode"

	"github.com/blueschu/advent-of-go/advent"
)

func SolveDay01(input []byte) (advent.Solution, error) {
	trimmed := bytes.TrimSpace(input)
	if !checkAllAsciiDigits(trimmed) {
		return advent.Solution{}, errors.New("non-digit byte in input")
	}
	part1 := sumRepeatedDigits(trimmed)
	part2 := sumCircularPartneredDigits(trimmed)
	return advent.SolutionFromInts(part1, part2), nil
}

// checkAllAsciiDigits checks whether all of the provided bytes are ascii digits
// between '0' and '9'.
func checkAllAsciiDigits(input []byte) bool {
	for _, b := range input {
		if !unicode.IsDigit(rune(b)) {
			return false
		}
	}
	return true
}

// sumRepeatedDigits computes the sum of the digits in the provided byte slice
// that are immediately followed by a matching digit.
// It assumes that all of the bytes in the provided sequence represent valid ascii digits.
func sumRepeatedDigits(input []byte) (sum int) {
	var lastByte = input[0]
	for _, b := range input[1:] {
		if lastByte == b {
			sum += int(b - '0')
		}
		lastByte = b
	}
	if lastByte == input[0] {
		sum += int(lastByte - '0')
	}
	return
}

// sumCircularPartneredDigits computes the sum of the digits in the provided byte slice
// whose "circular partner" (the digit halfway around the sequence) is the same digit.
// It assumes that all of the bytes in the provided sequence represent valid ascii digits.
func sumCircularPartneredDigits(input []byte) (sum int) {
	l := len(input) / 2
	for i, b := range input[:l] {
		if b == input[i+l] {
			sum += int(b - '0')
		}
	}
	sum *= 2
	return
}
