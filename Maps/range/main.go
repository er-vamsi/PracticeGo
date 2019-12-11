package main

import "fmt"

func main(){
	greetings := map[int]string{
		1:	"Hi",
		2:	"Hello",
		3:	"Bonjour",
		4:	"Namasthe",
	}

	for key, val := range greetings{
		fmt.Println(key,"---",val)
	}
}