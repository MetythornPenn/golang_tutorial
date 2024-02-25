package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our app")
	fmt.Println("Please rate our app between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	// // convert input to dtype from string to integer
	// int_input, err := strconv.ParseUint(input, 1, 32)
	// final_input := int(int_input)

	// if err != nil {
	// 	fmt.Println("Invalid input")
	// 	return
	// }
	
	// if final_input > 5 {
	// 	fmt.Println("Please input number less or equal 5")
	// 	return
	// } else if final_input < 1 {
	// 	fmt.Println("Please input number greater than or equal 1")
	// 	return
	// }
	fmt.Println("Thanks for rating: ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adding 1 to your rating: ", numRating + 1)
	}

}