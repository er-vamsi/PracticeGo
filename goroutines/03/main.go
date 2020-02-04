package main

import "fmt"
import "time"

func main(){
	var msg = "Hello"
	go func(msg string){
		fmt.Println(msg)
	}(msg)
	msg = "Goodbye!!"
	time.Sleep(100 * time.Millisecond)
}