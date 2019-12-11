package main

import "fmt"

func main(){
	var greetings = map[int]string{
		0:	"Hi",
		1:	"Hello",
		2:	"Bonjour",
		3:	"Namaskar",
	}

	//delete(greetings, 2)

	if val, exists := greetings[2]; exists{
		fmt.Println("Value Exist")
		fmt.Println("Value:",val)
		fmt.Println("Exist:",exists)
	}else{
		fmt.Println("Value doesn't Exist")
		fmt.Println("Value:",val)
		fmt.Println("Exist:",exists)		
	}

}