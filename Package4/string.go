package main

import (
	"fmt"
	"strings"
)

// strings are immutable in GO. (i.e We can't modify them once created)

func main() {
	// GO uses `utf-8` to store strings / character.
	var myString = "résumé"
	var indexed = myString[1]

	fmt.Printf("%v, %T\n", indexed, indexed) // %T : For the data type
	// It outputs utf-8 encoding for the character. That too only first byte value in all the cases including
	// multi-byte character like é .

	for i, v := range myString { // i - starting index of the corrosponding character ; v - complete utf-8 encoding ;
		fmt.Println(i, v)
	}

	// In the above ways, we can deal only with the underlying utf-8 encoding of array of Bytes.

	fmt.Printf("The length of myString is %v\n", len(myString))

	// Let's look at []rune which directly deals with characters original utf-8 encoding.
	// runes are just alias for int32 character.

	var myString2 = []rune("résumé")
	var indexed2 = myString2[1]

	fmt.Printf("%v, %T\n", indexed2, indexed2)
	// It outputs original utf-8 encoding for the character index-wise.

	for i, v := range myString2 { // Here we get continuous index.
		fmt.Println(i, v)
	}

	fmt.Printf("The length of myString2 is %v\n", len(myString2))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i] // We are actually creating completely new string everytime due to it's immutable nature in GO.
	}
	fmt.Printf("\n%v", catStr)

	// Instead we can use strings package of GO to avoid the inefficiency.

	var stringBuilder strings.Builder
	for i := range strSlice {
		stringBuilder.WriteString(strSlice[i]) // Here We are not creating new string everytime.
	}

	var catStr2 = stringBuilder.String()
	fmt.Printf("\n%v", catStr2)

}
