package main

import "fmt"
import "sync"

var wg = sync.WaitGroup{}

func main(){
	var msg = "Hello"
	wg.Add(1)
	go func(msg string){
		fmt.Println(msg)
		wg.Done()
	}(msg)
	wg.Wait()
}