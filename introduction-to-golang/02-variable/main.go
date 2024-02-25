package main

import "fmt"

const LoginToken int = 12345


func main() {
	var username string = "metythorn"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// default values and some aliases
	var anotherVariable2 float64
	fmt.Println(anotherVariable2)
	fmt.Printf("Variable is of type: %T \n", anotherVariable2)


	// default values and some aliases
	var anotherVariable3 string
	fmt.Println(anotherVariable3)
	fmt.Printf("Variable is of type: %T \n", anotherVariable3)

	// implicit type
	var website = "sigmoidx.com"
	fmt.Println(website)

	// no var style
	numberOfUser := "i love u"
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)



}