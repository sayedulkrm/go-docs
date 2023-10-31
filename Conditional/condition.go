package main

import "fmt"

func main() {

	// control   ==> conditional || loop control

	// conditionalIF()

	// evenOdd()

	// lgNum()

	switchcase()

}

func conditionalIF() {

	number := 1

	if number > 0 {

		fmt.Printf("Positive")
	} else if number < 0 {

		fmt.Printf("Negative")
	} else {

		fmt.Printf("Zero")
	}

}

func evenOdd() {
	var number int

	fmt.Printf("Enter any number: ")
	fmt.Scan(&number)

	fmt.Printf("The number is %v , And it is ", number)

	if number > 0 {

		fmt.Printf("Positive")
	} else if number < 0 {

		fmt.Printf("Negative")
	} else {

		fmt.Printf("Zero")
	}

	fmt.Println(" And your Number is")

	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

}

func lgNum() {

	var number1, number2 int

	fmt.Println("Enter 2 Number")
	fmt.Scan(&number1, &number2)

	if number1 > number2 {
		fmt.Printf("Large Number is = %v \n", number1)
	} else if number2 > number1 {
		fmt.Printf("Large Number is = %v \n", number2)
	} else {
		fmt.Printf("Both are equal")

	}

}

func switchcase() {
	// digit spelling

	var digit int

	fmt.Printf("Enter a digit from 0 - 9")

	fmt.Scan(&digit)

	switch digit {
	case 0:
		fmt.Printf("Zero \n")

	case 1:
		fmt.Printf("One \n")

	case 2:
		fmt.Printf("Two \n")

	case 3:
		fmt.Printf("Three \n")

	case 4, 5, 6, 7:
		fmt.Printf("Thats big number \n")

	default:
		fmt.Printf("Unknown Digits \n")

	}

}
