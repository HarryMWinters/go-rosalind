package dna

// Given a string of letters respresenting nucleotides,
// return the counts of A, C, G, T in that string.

import (
	"strconv"
	"strings"
)

type algorithm struct{}

// Execute Counts the numbers of each nucleotide in a sequence.
func (a algorithm) Execute(data []byte) ([]byte, error) {
	sequence := strings.ToUpper(string(data))
	var output string
	output += strconv.Itoa(strings.Count(sequence, "A"))
	output += " " + strconv.Itoa(strings.Count(sequence, "C"))
	output += " " + strconv.Itoa(strings.Count(sequence, "G"))
	output += " " + strconv.Itoa(strings.Count(sequence, "T"))
	return []byte(output), nil
}
