package main

import (
	"fmt"
	"exercises/helpers"
	"exercises/01_loops_and_functions"
	"exercises/02_slices"
	"exercises/03_maps"
	"exercises/04_fibonacci_closure"
	"exercises/05_stringers"
	"exercises/06_errors"
	"exercises/07_readers"
	"exercises/08_rot13Reader"
)

func main() {

	fmt.Printf("The square root of 2 is %v\n", loops_and_functions.Sqrt(2))

	helpers.ImageViewer(slices.GetImage())

	fmt.Printf("The wordcount of this sentence: %v\n", maps.WordCount("The wordcount of this sentence"))

	fmt.Printf("A list of %v fibonacci numbers: %v\n", 10, fibonacci_closure.GetFibonacciList(10))

	for name, ip := range stringers.Hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	a, b := errors.Sqrt(2)
	fmt.Printf("The square root of 2 is %v, with error: %v\n", a, b)
	a, b = errors.Sqrt(-2)
	fmt.Printf("The square root of -2 is %v, with error: %v\n", a, b)

	fmt.Print("Validating reader: ")
	readers.Validate()

	rot13Reader.Crack()
}

