package main

import "fmt"
import "strconv"
import "os"


func fizzBuzzSwitch(lowerBound int, upperBound int) {
	for i := lowerBound; i <= upperBound; i++ {
		switch {
			case (i % 3 == 0 && i % 5 == 0): fmt.Println("FizzBuzz")
			case (i % 3 == 0): fmt.Println("Fizz")
			case (i % 5 == 0): fmt.Println("Buzz")
			default: fmt.Println(i);
		} 
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: fizzBuzzSwitch.go [lower bound] [upper bound]")
		os.Exit(1)
	}
	lowerBound, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Exit(2)
	}
	upperBound, err := strconv.Atoi(os.Args[2])
	if err != nil {
		os.Exit(2)
	}
	if (lowerBound > upperBound) {
		fmt.Println("lower bound must not be greater than upper bound")
		os.Exit(3)
	}
	fizzBuzzSwitch(lowerBound, upperBound)
}

