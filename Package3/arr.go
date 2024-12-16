package main

import "fmt"

func main() {
	var arr [3]int32 // array declaration
	fmt.Println(arr) // takes up the default value of data type of array if not initialised.

	var arr2 [3]int32 = [3]int32{1, 2, 3} // declaration + initialisation.
	fmt.Println(arr2)                     // print the whole array.
	fmt.Println(arr2[1:3])                // slicing the array [inclusive, exclusive] by index.
	arr2[2] = 45                          // inserting value
	fmt.Println(arr2[2])                  // Accessing value by index

	// Accessing memory address
	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])

	arr3 := [4]int8{1, 2, 3, 4} // short-hand property
	fmt.Println(arr3)

	fmt.Println("Slice data structure:")
	slices()

	fmt.Println("Map data structure:")
	map_ds()

}

func slices() {
	intArr := [...]int32{2, 4, 8} // spread operator [...]
	fmt.Println(intArr)

	// slice declaration + initialisation
	var intSlice []int32 = []int32{3, 6, 9, 12}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	// cap() stands for total capacity.

	intSlice = append(intSlice, 15)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) // appending another slice using spread operator
	fmt.Println(intSlice)

	// slice declaration using make() function
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)
	// make(type, length, capacity)
	// capacity is optional. In case of absence: capacity == length
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice3), cap(intSlice3))

	// slices reallocate whole another array while append() if capacity is exceeded.

	// Array / Slice iteration
	for i, v := range intSlice2 {
		fmt.Printf("The index is %v, and value is %v \n", i, v)
	}

	// while loop
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	// Another approach for while loop in GO
	for {
		if i < 0 {
			break
		}
		fmt.Println(i)
		i = i - 1
	}

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}

func map_ds() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adarsh": 23, "Aman": 19, "Jacob": 27}
	fmt.Println(myMap2)

	var age, ok = myMap2["Arsh"] // It also returns an optional boolean value.
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid name")
	}

	delete(myMap2, "Adarsh") // It deletes by reference so no return value is given.
	fmt.Println(myMap2)

	// Loop in GO
	// Map iteration doesn't preserve order.
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, age: %v\n", name, age)
	}

}
