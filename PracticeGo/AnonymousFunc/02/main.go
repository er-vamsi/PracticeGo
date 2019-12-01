package main

import "fmt"

func main(){
	increment := func (a,b int) int{
		return a+b
	}

	fmt.Println(increment(2,3))
}