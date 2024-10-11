package main

import "fmt"

func main() {
	// * int number formating
	var intNum = 6

	//* int number to binary
	fmt.Printf("Binary Number= %b\n", intNum)
	//* int number to octal
	fmt.Printf("Octal Number= %o\n", intNum)
	//* int number to Hexa
	fmt.Printf("Hexa Number= %x\n", intNum)

	// * floating number formating ex 1
	var num = 3.1416
	fmt.Printf("%f\n", num)

	// * floating number formating ex 2
	fmt.Printf("%.2f\n", num)

	// * string formating ex 1
	var name = "Tarek Siddique "
	fmt.Printf("%s\n", name)

	// * string fromating ex 2
	fmt.Printf("%q\n", name)

}
