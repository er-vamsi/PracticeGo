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
	sp1 SPrint

}