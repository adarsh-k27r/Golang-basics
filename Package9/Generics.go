package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))

	var floatSlice = []float32{1, 2, 3}
	fmt.Println(sumSlice[float32](floatSlice))

}

// Generics are similar to Templates in C++
func sumSlice[T int | float32](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
