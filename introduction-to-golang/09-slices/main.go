package main

import "fmt"



func main(){
	fmt.Println("Hello World")

	var slice = []string{"apple", "banana", "orange"}
	fmt.Println(slice)

	slice = append(slice, "mango", "grape")

	fmt.Println(slice[1:])

	hightScore := make([]int, 4)
	hightScore[0] = 1
	hightScore[1] = 2
	hightScore[2] = 3
	hightScore[3] = 4
	hightScore = append(hightScore, 5)
	fmt.Println(hightScore)

	// remove value from slice base on index 
	hightScore = append(hightScore[:2], hightScore[3:]...)

	fmt.Println(hightScore)


}