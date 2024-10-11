package main

import "fmt"

func main() {

	// variable define
	var fullName string
	var age int

	// print command for user and assign user input value to variable
	fmt.Print("Enter Your Full Name:-")
	fmt.Scan(&fullName)
	fmt.Print("Input your age:-")
	fmt.Scan(&age)
	fmt.Print("Enter ")
	// print value
	fmt.Printf("My Name is %v\n", fullName)
	fmt.Printf("I am %v years old", age)

}
