package main

import "fmt"

func main() {
	a := 10

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
}
