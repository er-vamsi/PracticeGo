package main

import "fmt"

func main(){
	data := []float64{10,20,30,40}
	average(data...)
}

func average(n ...float64){
	total := 0.0
	for _,i := range(n){
		total += i
	}
	fmt.Println("Average: ", total/float64(len(n)))
}