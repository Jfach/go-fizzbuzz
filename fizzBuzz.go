package main

import "fmt"
import "os"
import "strconv"


func fizzBuzz(lowerBound int, upperBound int) {
	for i := lowerBound; i <= upperBound; i++ {
		if (i % 3 == 0) && (i % 5 == 0) {
			fmt.Println("FizzBuzz")
		} else if (i % 3 == 0) {
			fmt.Println("Fizz")
		} else if (i % 5 == 0) {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}


func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: fizzBuzz.go [lower bound] [upper bound]")
		os.Exit(1)
	}
	lowerBound, err := strconv.Atoi(os.Args[1])  
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	upperBound, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if (lowerBound > upperBound) {
		fmt.Println("lower bound must not be greater than upper bound")
		os.Exit(3)
	}
	fizzBuzz(lowerBound, upperBound)
}