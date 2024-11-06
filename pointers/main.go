package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	var p Person = Person{Name: "Alice", Age: 30}
	fmt.Println(p.Name, p.Age)
	p.Age = 50
	fmt.Println(p.Name, p.Age)

}
