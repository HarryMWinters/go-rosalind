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
	panic("String reverse not implemented!")
}

func main() {
	data, err := ioutil.ReadFile("rosalind_revc (1).txt")
	checkError(err)
	//sequence := string(data)
	var reverseComplement string
	for i := 0; i < len(data); i++ {
		fmt.Println(string(data[i]))
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
			reverseComplement += "\n"
		default:
			panic("Unsupported character detected.")
		}
	}
	reverseComplement = stringReverse(reverseComplement)
	fmt.Printf("%v", reverseComplement)
}
