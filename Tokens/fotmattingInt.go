package main

import "fmt"

func main() {

	// string formatting
	var name string

	currentName := "John"

	fmt.Printf("%s is your current Name\n", currentName)
	fmt.Printf("%v is your current Name\n", currentName)
	fmt.Printf("%q is your current Name\n", currentName)

	fmt.Printf("Enter your name = ")
	fmt.Scanf("%v", &name)

	fmt.Printf("Hello %v\n", name)

	// formating number
	var Decinum1 int

	fmt.Printf("Enter a Decimal number= ")
	fmt.Scanf("%v", &Decinum1)

	fmt.Printf("Binary number is = %b\n", Decinum1)
	fmt.Printf("Octal number is = %o\n", Decinum1)
	fmt.Printf("Hexadecimal number is = %x\n", Decinum1)

}
