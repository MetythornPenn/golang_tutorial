package main

import "fmt"



func main(){
	fmt.Println("Hello World")

	var fruitList [4]string

	fruitList[0] = "Mango"
	fruitList[1] = "Banana"
	fruitList[2] = "Orange"
	fruitList[3] = "Grapes"
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var fruitList2 = [4]string{"Mango","Banana","Orange","Grapes"}
	fmt.Println(fruitList2)
	fmt.Println(len(fruitList2))

}