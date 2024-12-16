package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello World!")

	// Variable declaration and initialisation.
	var intNum int16 = 32767
	// intNum += 1 // overflow
	fmt.Println(intNum)

	var fnum float64 = 24367.48434
	fmt.Println("Float number is:", fnum)

	const text string = "text" + " is" + " here" // string concatenation using (+) operator.
	fmt.Println(text)

	const text2 string = `heyy 
	there` // multiple line using backtick.
	fmt.Println(text2)

	var boolNum bool = false
	fmt.Println("Bool is:", boolNum)

	const cnum uint8 = 58 // We can't modify const once initialised unlike variables.
	// Also you can't just declare const, but also have to initialise it.
	fmt.Println("unsigned int is:", cnum)

	// int types : int, int8, int16, int32, int64 ; where suffix number is bit memory size of the int.
	// unsigned int types : uint, uint8, uint16, uint32, uint64
	// float types : float32, float64
	// bool : false, true
	// string

	// Strongly typed: so typecasting is needed at times.
	var a uint = 12
	var b uint32 = 45
	var result uint32 = uint32(a) + b // typecasting
	fmt.Println("result is", result)

	// operators : +, -, *, /, % etc.
	// int division results into integer.

	var text3 = "text3"                        // inferred data type string.
	fmt.Println(utf8.RuneCountInString(text3)) // returns the length of the string.
	fmt.Println(len(text3))                    // returns the number of bytes taken by the string.

	// `rune` is a datatype in itself which represents a character.
	var ch rune = 'a'
	fmt.Println(ch)

	// If variables are only declared but not initialised then they take up their default value.
	// All int, uint, rune and float default value is 0.
	// String default value is ' ' (i.e empty string)
	// Bool default value is ' false '

	var def string
	fmt.Println("Default value of string is:", def)

	// we can even remove ` var ` keyword and use its shorthand property (:=)
	v1 := 1
	fmt.Println(v1)

	// Initialise multiple variables.
	v2, v3 := 2, 3
	fmt.Println(v2, v3)

}
