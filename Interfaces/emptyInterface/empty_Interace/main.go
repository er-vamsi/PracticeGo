package main

import "fmt"


	type vehicles interface{}

	type vehicle struct{
		seats int
		maxspeed int
		color string
	}

	type car struct{
		vehicle
		Wheels int
		Doors int		
	}

	type plane struct{
		vehicle
		Jet bool
	}

	type boat struct{
		vehicle
		Length int
	}
	
	func main(){
		maruthi := car{}
		vv := car{}
		nissan := car{}
		boeing747 := plane{}
		boeing787 := plane{}
		airbus := plane{}
		ins := boat{}

		rides := []vehicles{maruthi, vv, nissan, boeing747, boeing787,airbus, ins}

		for _,i := range(rides){
			fmt.Println(i)
		}
	}
