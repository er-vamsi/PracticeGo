package main

import ("fmt"
		"sync"
		)

func main(){
	c := make(chan int)
	var wg sync.WaitGroup
	//wg.Add(1)

	go func(){
		wg.Add(1)
		for i := 0; i <= 9; i++{
			c <- i
		}
		wg.Done()
	}()

		go func(){
		wg.Add(1)
		for i := 0; i <= 9; i++{
			c <- i
		}
		wg.Done()
	}()

	go func(){
		wg.Wait()
		close(c)

	}()
	
	for i := range(c){
		fmt.Println(i)
	}

}