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
	// calculation()

	// assign()
	// unaryOperators()

	// logicalOperators()

	bitwiseOp()
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

func assign() {

	// Assign operations -> =, +=, -=, *=, /=, %=

	x := 4
	y := 2

	x = x + 2

	x += 5 // x = x + 5

	x *= y

	x /= y
	fmt.Printf("x = %v \n", x)

	x %= y

	fmt.Printf("x = %v", x)

}

func unaryOperators() {

	x := "12"

	y := "14"

	fmt.Printf("x = %v", x == y)

	// for x := 100; x > 10; x-- {
	// 	fmt.Printf("x = %v \n", x)
	// }

}

func logicalOperators() {

	// --> &&, ||, !

	x := 4

	y := 2

	z := !(x == y)

	fmt.Printf("z = %v \n", z)

}

func bitwiseOp() {

	// ---> & , | , ^ , << , >>

	x := 18

	y := 17

	z := x & y

	a := x | y

	b := x ^ y

	fmt.Printf(" z = %v \n", z)

	fmt.Printf(" a = %v \n", a)

	fmt.Printf(" b = %v \n", b)

}
