package main

import "fmt"


func main(){
	fmt.Println("Welcome to function in golang")
	greater()
	// greaterTwo()

	result := add(1, 2)
	fmt.Println("Result is: ", result)

	proRes, proMsg := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro message is: ", proMsg)

}



func greater() {
	fmt.Println("hello from golang")
}

// normal func 
func add(x int, y int) int {
	return x + y
}

// pro function
func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Pro fucntion"
}