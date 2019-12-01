package main

import "fmt"

func greeter() func() string{
	return func() string{
		return "Hello..World!!!"
	}
}

func main(){
	greetings := greeter()
	fmt.Println(greetings())
	fmt.Printf("%T \n",greetings)
}