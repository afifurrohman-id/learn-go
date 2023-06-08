package main

import "fmt"

func main() {
	//! number (integer / float)
	const (
		salary int = 100000
		count float64 = 123810.1201 
	)
	
		fmt.Print("Integer: ",salary)
		fmt.Print("\nFloat: ",count)
	
	// ! string Sting must be double quote
	const (
	// * Use \n to add enter
	message = "Hello Go!"
	//* String escape character (\) before "
	quote = "Everyday i like a \"tea\"\n"
	)
	
	fmt.Println("\nString: ", message)
	fmt.Print("\nSay: " ,quote)

	// !boolean (true or false)
	const isMarried bool = false

	fmt.Print(isMarried)

}