package main

import "fmt"

type Animal struct{
	voice string
}

type cat struct{
	Animal
	friendly bool
}

type dog struct{
	Animal
	friendly bool
}

func main(){
	cat1 := cat{Animal{"Meow"}, true}
	dog1 := dog{Animal{"Bark"}, true}

	categoery := []interface{}{cat1, dog1}

	fmt.Println(categoery)
}