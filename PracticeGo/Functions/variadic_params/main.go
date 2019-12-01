package main

import "fmt"

func main(){
	av := Average(10,12,11.5,3,17)
	fmt.Println(av)
}

//func Average(val []float64) float64  {}  -- Slice Param
func Average(val ...float64) float64{
	total := 0.0
	for _,v := range(val){
		total += v
	}
	return total/float64(len(val))
}