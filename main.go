package main

import (
	"flag"
	"fmt"

	"github.com/HarryMWinters/rosalind/dna"

	rosalib "github.com/HarryMWinters/rosalind/lib"
)

type algorithm interface {
	// TODO: Convert this to streaming
	Execute(data []byte) ([]byte, error)
}

func parseAlgorythm(algoName string, implementedAlgos map[string]algorithm) (struct dna.algorithm, error) {
	algo, ok := implementedAlgos[algoName]
	if !ok {
		// TODO: Handle this gracefully.
		panic("Unsupported algorithm name.")
	}
	newAlgorithm = algo.algorythm{}
	return newAlgorithm, nil
}

func main() {
	var algoName = flag.String("algo", "default_help", "Specify algorythm code here.")
	var inputFile = flag.String("file", "default_help", "Relative path to inpute file.")
	flag.Parse()

	algos := map[string]algorithm{
		"dna": dna.Execute,
	}
	// parseAlgo seems like overkill to just check for key in map
	algo, err := parseAlgorythm(*algoName, algos)
	file := rosalib.NaiveOpen("rosaline_dna.txt")
	result, err = algo(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
	fmt.Println("algoName has value: ", *algoName)
	fmt.Println("Input file has the value: ", *inputFile)
	fmt.Println("Program execution complete.")
}
