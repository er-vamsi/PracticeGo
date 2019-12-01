//This is belongs to Methods

package main

import "fmt"

type SPrint string

func (s SPrint) greet(st string){
	fmt.Println(st)
}

type SPrint1 string

func (s1 SPrint1) greet(st1, st2 string) string{
	return fmt.Sprint(st1,st2)
}
func main(){
	var s SPrint
	s.greet("Thota")
	var s1 SPrint1
	n := s1.greet("Venkata","Thota")
	fmt.Println(n)

}