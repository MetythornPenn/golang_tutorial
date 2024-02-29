package main

import "fmt"



func main(){
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	fmt.Println("\nFor loop")
	for d:=0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println("\nUsing range")
	for i := range days {
		fmt.Println(days[i])
	}

	// using range with index
	for i, day := range days {
		fmt.Printf("Index is %v and value is %v\n", i, day)
	}

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 5 {
			break
		}

		fmt.Println("Value is less than 10")
		rougeValue++
	}

}