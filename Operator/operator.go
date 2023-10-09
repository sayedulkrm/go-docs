package main

import "fmt"

// Types of operators
// Arithmetic  -> +, -, *, /, %

// Comparison
// Logical
// Bitwise
// Assignment
// Increment
// others

func main() {

	// arithmetic()
	calculation()

}

func arithmetic() {
	var num1, num2 int

	fmt.Println("Enter number 1: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter number 2: ")
	fmt.Scanln(&num2)

	fmt.Println(" / is ", num1/num2)

	result := num1 % num2
	fmt.Println(" % is ", result)

	// result2 = float32(result)
	result2 := float32(num1) / float32(num2)

	fmt.Printf("Final result is %.5f ", result2)
}

func calculation() {

	// Calculation of area of circle = 3.14 * r * r

	var radius float32

	// radius := 5
	fmt.Println("Enter radius of circle: ")
	fmt.Scanln(&radius)
	pie := 3.14
	area := float32(pie) * float32(radius) * float32(radius)
	fmt.Printf("Area of circle = %.2f \n \n", area)

	// Calculations of traingle = base * height / 2

	base := float32(10)
	height := float32(20)
	area2 := base * height / 2

	fmt.Printf("Area of traingle = %.2f ", area2)

}
