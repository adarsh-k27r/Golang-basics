package main

import "fmt"

func main() {
	// default value when memory location is not specified is `nil`
	var p *int32 = new(int32) // default value is zero here (of int32)
	var i int32
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)
	*p = 10                                          // assigning value to the pointer directly.
	fmt.Printf("\nThe value p points to is: %v", *p) // Use `%p` as a marker for memory location.

	p = &i // Assigning address of i to pointer variable - p. Now both point at the same memory location.
	*p = 16
	fmt.Printf("\nThe value p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)

	// slices use pointers under the hood when we copy them directly using assignment operator.
	// Slices are exception to the non-pointer variable in this context.

}
