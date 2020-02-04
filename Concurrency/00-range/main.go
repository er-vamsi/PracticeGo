package main

import ("fmt"
		"sync"
)

func main(){
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		for i:=0; i<10; i++{
			c <- i
		}
		wg.Done()
		
	}()

	go func(){
		wg.Wait()
		close(c)
	}()

	//close(c)
	for n := range(c){
		fmt.Println(n)
	}
}
