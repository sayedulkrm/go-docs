package main

import "fmt"

// variables naming rules
// 1. letters, numbers, _ , and underscore
// 2. cannot use Keywords as variable names
// 3. cannot have spaces
// 4. cannot start with number
// 5. case sensitive. EG: "Name" and "name" are different
// 6. cannot be empty

func main() {

	// var declarations
	var name string
	var age int
	var height float32

	// var initialization
	name = "Jack"
	age = 25
	height = 5.11112222333344445555

	fmt.Println(name, "is a student")
	fmt.Println(name, "'s height is", float32(height))
	fmt.Println(name, "'s age is", age)
	fmt.Println(name, "'s name is \"Jack\"")
}
