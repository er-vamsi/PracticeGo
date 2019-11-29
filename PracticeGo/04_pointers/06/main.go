package main

import "fmt"

func zero(z *int) {
	fmt.Println(z) // Prints address of z
	*z = 0         // Changes value of z
}

func main() {
	x := 10
	fmt.Println(x)
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x)
}
