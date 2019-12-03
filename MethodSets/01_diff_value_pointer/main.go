package main

import "fmt"

var a = 10


func funcA(){
	a = a + 1
	fmt.Println("After add: ",a)
}

func funcB(b *int){
	*b = *b + 1
	fmt.Println("Pointer: ", *b)

}

func main(){
	funcA()
	fmt.Println("Value: ", a)
	funcB(&a)
	fmt.Println("Vlaue2: ", a)

}