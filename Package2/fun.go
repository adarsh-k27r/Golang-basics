package main

import (
	"errors"
	"fmt"
)

func main() { // curly braces shall begin just from here only; not from the next line
	var s1 string = "Hello Adarsh!"
	printString(s1)

	var numerator int = 12
	var denominator int = 0
	var result, remainder, err = div(numerator, denominator)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of integer division is %v", result)
	} else {
		fmt.Printf("The result of integer division is %v with %v as remainder", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("Division is exact")
	case 1, 2:
		fmt.Println("Division is closer")
	default:
		fmt.Println("Division not exact")

	}

}

func printString(s string) {
	fmt.Println(s)
}

func div(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("can't divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
