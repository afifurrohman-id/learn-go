package main

import "fmt"

func main() {
	// message string not constant

	var message string
	// can reassign
	message += "Hello"
	fmt.Println(message)

	// * another shortcut 
	salary := 1000000 
	fmt.Print(salary)


	// !constant cannot reassign
	const name = "\nAfif"

	// name = "Michael" // ! Error: Cannot reassign constant variable

	fmt.Print(name)
}