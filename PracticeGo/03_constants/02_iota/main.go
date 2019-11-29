package main

import "fmt"

const A = iota
const B = iota

const (
	C = iota
	D = iota
)

//continuous increment of iota
const(
	E = iota
	F
	G
	H = "Vamsi"
	I
	J = iota
)

func main(){
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Println(I)
	fmt.Println(J)
}