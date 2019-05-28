package main

import (
	"fmt"
	"io/ioutil"
)

// Given a DNA string return the reverse complement

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func stringReverse(s string) string {
	reversed := []rune(s)
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return string(reversed)
}

func main() {
	data, err := ioutil.ReadFile("rosalind_revc (2).txt")
	checkError(err)
	var reverseComplement string
	for i := 0; i < len(data); i++ {
		switch string(data[i]) {
		case "A":
			reverseComplement += "T"
		case "T":
			reverseComplement += "A"
		case "C":
			reverseComplement += "G"
		case "G":
			reverseComplement += "C"
		case "\n":
		default:
			panic("Unsupported character detected.")
		}
	}
	reverseComplement = stringReverse(reverseComplement)
	fmt.Printf("%v", reverseComplement)
}
