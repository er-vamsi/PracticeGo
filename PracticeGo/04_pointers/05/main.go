package main

import (
	"flag"
	"fmt"
)

func main() {
	s := "Vamsi"
	_ = s
	csvFilename := flag.String("csv", "problems.csv", "A csv file in format of 'question, answer'")

	fmt.Println(csvFilename)
	fmt.Println(*csvFilename)
}
