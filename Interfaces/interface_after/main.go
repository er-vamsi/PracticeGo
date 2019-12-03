package main

import "fmt"

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

type Shape interface {
	area() float64
}

func info(s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {
	s := Square{20}
	info(s)
}
