package main

import "fmt"

type Point struct {
	X,Y float64
}

func (p Point) IsAbove(y float64) bool{
	return p.Y > y
}

func main(){
	p := Point{20.0,30.0}
	fmt.Println(p)
	fmt.Println(p.Y)
	fmt.Println("Is Point p above the line y = 15?", p.IsAbove(15.0))

}


