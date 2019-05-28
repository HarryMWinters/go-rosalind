package dna

// Given a string of letters respresenting nucleotides,
// return the counts of A, C, G, T in that string.

import (
	"fmt"
	"strings"

	rosalib "github.com/HarryMWinters/rosalind/lib"
)

// DNA Counts the numbers of each nucleotide in a sequence.
func DNA() {
	data := rosalib.NaiveOpen("rosalind_dna (1).txt")

	var numA, numC, numG, numT int
	sequence := strings.ToUpper(string(data))
	numA = strings.Count(sequence, "A")
	numC = strings.Count(sequence, "C")
	numG = strings.Count(sequence, "G")
	numT = strings.Count(sequence, "T")
	fmt.Printf("%d %d %d %d\n", numA, numC, numG, numT)
}
