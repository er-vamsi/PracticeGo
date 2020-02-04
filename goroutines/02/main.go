package main

import "fmt"
import "time"

func main(){
	var msg = "Hello"
	go func(){
		fmt.Println(msg)
	}()
	msg = "Goodbye!!"  //Race condition
	time.Sleep(100 * time.Microsecond)
}