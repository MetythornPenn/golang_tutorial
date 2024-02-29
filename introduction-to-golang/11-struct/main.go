package main

import "fmt"

type User struct {
	Name 		string
	Email	 	string
	Status 	bool
	Age 		int
}


func main(){
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent
	xyber := User{"xyber", "xyber@go.dev", true, 20}
	fmt.Println(xyber)
	fmt.Printf("xyber details are: %+v\n", xyber)
	fmt.Printf("Name is %v and email is %v\n", xyber.Name, xyber.Email)

}




