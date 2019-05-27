package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Replace all T's with U's in a string to represent
// the transcription of DNA to RNA

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// TODO: This loads the entire file into memore which is resource intensive
	// and could break with extremely large files (Genome size and larger)
	data, err := ioutil.ReadFile("rosalind_rna (1).txt")
	checkError(err)
	sequence := strings.ReplaceAll(string(data), "T", "U")
	fmt.Printf("%v", sequence)
}
