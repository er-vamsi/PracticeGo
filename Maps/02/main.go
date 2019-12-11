package main

import "fmt"

func main(){
	var greetings = map[string]string{}

	greetings["Thota"] = "Hello"
	greetings["Vamsi"] = "hi"

	var val, exists = greetings["Thota"]
	

	//fmt.Println(greetings == nil)
	fmt.Println(greetings)
	fmt.Println(val)
	fmt.Println(exists)
}