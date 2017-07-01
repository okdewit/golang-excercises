package main

import (
	"fmt"
	"exercises/helpers"
	"exercises/01_loops_and_functions"
	"exercises/02_slices"
	"exercises/03_maps"
	"exercises/04_fibonacci_closure"
)

func main() {
	fmt.Printf("The square root of 2 is %v\n", loops_and_functions.Sqrt(2))
	helpers.ImageViewer(slices.GetImage())
	fmt.Printf("The wordcount of this sentence: %v\n", maps.WordCount("The wordcount of this sentence"))
	fmt.Printf("A list of %v fibonacci numbers: %v\n", 10, fibonacci_closure.GetFibonacciList(10))
}

