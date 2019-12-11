package main

import "fmt"

func main(){
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2
	m["k3"] = 3

	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	
	delete(m, "k3")
	fmt.Println(m)

	_, ok := m["k3"]
	fmt.Println(ok)

	n := map[string]int{"n1" : 1, "n2": 2}
	fmt.Println(n)
	
	m["k3"] = 3
	fmt.Println(m)
}