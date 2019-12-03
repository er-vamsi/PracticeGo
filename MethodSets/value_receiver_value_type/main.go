package main

import ("fmt"
		"math"
)

type circle struct{
	radious float64
}

func (c circle) area() float64{
	return math.Pi * c.radious * c.radious
}

type Shape interface{
	area() float64
}

func info(s Shape){
	fmt.Println(s.area())
}

func main(){
	c := circle{10}
	info(c)
}