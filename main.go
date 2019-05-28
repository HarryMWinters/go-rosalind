package main

import (
	"flag"
	"fmt"
	"github.com/HarryMWinters/rosalind/dna"
)

func parseAlgorythm(algoName string, implementedAlgos map) (func, error){

}

func main() {
	var algoName = flag.String("algo", "default_help", "Specify algorythm code here.")
	var inputFile = flag.String("file", "default_help", "Relative path to inpute file.")
	flag.Parse()

	algos := map[string]interface{} {
		"dna": dna.DNA
	}
	
	fmt.Println("algoName has value: ", *algoName)
	fmt.Println("Input file has the value: ", *inputFile)
	fmt.Println("Program execution complete.")
}
