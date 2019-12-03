package main

import ("fmt"
		"math"
)

type circle struct{
	radius float64
}
func (c *circle) area() float64{
	return math.Pi * c.radius * c.radius
}

type shape interface{
	area() float64
}

func info (s shape){
	fmt.Println(s.area())
}


func main(){
	c := circle{5}
	info(&c)
}