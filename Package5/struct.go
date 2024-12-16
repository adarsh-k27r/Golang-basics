package main

import "fmt"

// struct data structure in Golang : similar to class in C++

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
	rent
}

func (e gasEngine) milesLeft() uint8 { // A method, similar to function but tightly coupled to struct data type.
	return e.gallons * e.mpg
}

type owner struct {
	name string
}

type rent struct {
	money uint8
}

// Interfaces are like Templates in C++ : Makes our function more generic for other types.
type engine interface {
	milesLeft() uint8 // method signature.
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("\nYou can make it there")
	} else {
		fmt.Println("\nNeed to fuel up first")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15, ownerInfo: owner{"Adarsh"}, rent: rent{45}}
	// {25, 15, owner{"Adarsh"}, rent{45}} will set it in specified order as well.
	myEngine.mpg = 15
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name, myEngine.money)

	// anonymous struct are not reuseable.
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}

	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	// calling struct method
	fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())

	canMakeIt(myEngine, 20)

}
