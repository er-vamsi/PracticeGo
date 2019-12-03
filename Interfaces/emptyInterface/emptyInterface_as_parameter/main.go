package main

import "fmt"

type animal struct{
	sound string
}

type cat struct{
	animal
	friendly bool
}

type dog struct{
	animal
	friendly bool
}

type human string

func specs(a interface{}){
	fmt.Println(a)
}

func specs1(a struct{}){
	fmt.Println(a)
}

func main(){
	cat1 := cat{animal{"Meow"}, true}
	cat2 := cat{animal{"LoudMeow"}, false}
	dog := dog{animal{"Barking"}, true}
	h1 := human("Thota")
	
	specs(cat1)  // specs function accepts any type argument due to empty interface
	specs(cat2)
	specs(dog)
	specs(h1)

	//specs1(cat1)  // Returns error of missmatch types between cat and struct

}