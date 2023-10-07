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
	var name string = "Jack"
	// if we use a variable like the above, we dont need to add the type declaration.

	age := 25
	// If we give := we	don't need to add "var" declaration. Its shortcut

	var height float32

	// var initialization
	height = 5.11112222333344445555

	const COUNTRY = "Bangladesh"
	// const Value cant be changeded but var value can be changed

	fmt.Println(name, "is a student")
	fmt.Println(name, "'s height is", float32(height))
	fmt.Println(name, "'s age is", age)
	fmt.Println(name, "'s name is \"Jack\"")

	// Formatting the print

	fmt.Printf("%v is a boy \n", name)
	fmt.Printf("%v is his age \n ", age)
	fmt.Printf("%v is from %v \n ", name, COUNTRY)

}
