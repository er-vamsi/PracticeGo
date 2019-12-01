package main

import "fmt"

func main(){
	for i:=250;i<=350;i++{
		fmt.Println(i,"-",string(i),"-",[]byte(string(i)))
	}
}