package rosalib

import "io/ioutil"

// CheckError Checks if e is an error and raises it if e != nil
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// NaiveOpen Loads whole file into memory and return byte array.
// Second value is error. Problematic if file is large since tbe whole
// file is stored in memory at once.
func NaiveOpen(filepath string) []byte {
	data, err := ioutil.ReadFile(filepath)
	checkError(err)
	return data
}

// TODO: Write less naive open
