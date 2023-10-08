package main

import "fmt"

func main() {

	var fullname, country string

	var num1, num2 int

	// Scan, Scanln, and Scanf are used to take input from the user
	// fmt.Print("Enter your fullname: ")
	// fmt.Scan(&fullname)
	// fmt.Print("Enter your country: ")
	// fmt.Scan(&country)

	fmt.Print("Enter 2 numbers: ")
	// fmt.Scan(&num1, &num2)

	fmt.Scanf("%v %v", &num1, &num2)

	fmt.Println("Sum of the numbers is: ", num1+num2)

	fmt.Println("Hello, " + fullname + " from " + country)

}
