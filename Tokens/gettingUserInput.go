package main

import (
	"fmt"
	"time"
)

func main() {
	joss()

	// var fullname, country string

	// var num1, num2 int

	// Scan, Scanln, and Scanf are used to take input from the user
	// fmt.Print("Enter your fullname: ")
	// fmt.Scan(&fullname)
	// fmt.Print("Enter your country: ")
	// fmt.Scan(&country)

	// fmt.Print("Enter 2 numbers: ")
	// fmt.Scan(&num1, &num2)

	// fmt.Scanf("%v %v", &num1, &num2)

	// fmt.Println("Sum of the numbers is: ", num1+num2)

	// fmt.Println("Hello, " + fullname + " from " + country)

}

func joss() {
	start := time.Now()

	for i := 0; i <= 10000; i++ {
		fmt.Printf("\rCounting: %d", i)
	}

	elapsed := time.Since(start)
	seconds := elapsed.Seconds()

	fmt.Printf("\nCounted to one million in %.2f seconds\n", seconds)

	fmt.Print("The percentage is 100%")
}
