package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter Meters: ")
	fmt.Scan(&meters)
	resultYards := meters * metersToYards
	fmt.Println("Result: ", resultYards)
}
