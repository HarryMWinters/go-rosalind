package main

// Given a string of letters respresenting nucleotides,
// return the counts of A, C, G, T in that string.

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// TODO: This loads the entire file into memore which is resource intensive
	// and could break with extremely large files (Genome size and larger)
	data, err := ioutil.ReadFile("rosalind_dna (1).txt")
	checkError(err)
	var numA, numC, numG, numT int
	sequence := strings.ToUpper(string(data))
	numA = strings.Count(sequence, "A")
	numC = strings.Count(sequence, "C")
	numG = strings.Count(sequence, "G")
	numT = strings.Count(sequence, "T")
	fmt.Printf("%d %d %d %d\n", numA, numC, numG, numT)
}
