package main

import "fmt"

func main() {

	//* type 1

	// var fullName string
	// var country string
	// var age int

	// fullName = "Tarek Siddique Nabil"
	// country = "Bangladesh"
	// age = 30

	//* type 2
	// var fullName string = "Tarek Siddique Nabil"
	// var country string = "Bangladesh"
	// var age int = 30

	//* type 3

	fullName := "Tarek Siddique Nabil"
	country := "Bangladesh"
	age := 30

	//* define variable  type 1
	// fmt.Println("My name is ", fullName)
	// fmt.Println("I am live in", country)
	// fmt.Println("I am", age, "years old")

	// * define variable  type 2

	fmt.Printf("My name is %v\n", fullName)
	/*

		*	here if we use Println our output will be = My name is %v Tarek Siddique Nabil

		?	so if we use %v for print variable we need to use printf insted println ,

		!	but if we use printf than our output will be "My name is Tarek Siddique NabilI am live in Bangladesh" we can't see any line brack thats why we used \n escape sequense for a line break


	*/
	fmt.Println("I am live in", country)
	fmt.Println("I am", age, "years old")

}
