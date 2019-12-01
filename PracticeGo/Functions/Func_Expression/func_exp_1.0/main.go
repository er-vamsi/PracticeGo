package main

import "fmt"

func main(){
	greetings := func(){   //Assigning a function to variable is called Function Expression
		fmt.Println("Hello World!!!")
	}
	greetings()
	fmt.Printf("%T \n",greetings)
}