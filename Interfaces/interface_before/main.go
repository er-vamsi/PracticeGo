package main

import "fmt"

type Square struct{
	side float64
}

func (s Square) area() float64{
	return s.side * s.side
}

func main(){
	s1 := Square{20.0}
	fmt.Println(s1.area())
}