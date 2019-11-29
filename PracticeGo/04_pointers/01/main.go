package main

import "fmt"

func main() {
	a := 45
	fmt.Println("a - ", a)
	fmt.Println("a's address - ", &a)
	fmt.Println("a - ", *(&a))
}
