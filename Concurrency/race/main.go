package main
import ("fmt"
		"sync"
		"time"
		"math/rand")

var wg sync.WaitGroup
var counter int

func main(){
	wg.Add(2)
	go increment("Foo: ")
	go increment("Bar: ")
	wg.Wait()
	fmt.Println("Final Count: ", counter)
}

func increment(s string){
	for i :=0; i<=20; i++{
		c := counter
		c++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = c
		fmt.Println(s, i,"Counter: ",counter)	
	}
	wg.Done()
}