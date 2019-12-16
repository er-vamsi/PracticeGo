package main
import ("fmt"
		"sync"
		"time"
		"math/rand")

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main(){
	wg.Add(2)
	go increment("Foo: ")
	go increment("Bar: ")
	wg.Wait()
}

func increment(s string){
	for i:=0; i<20;i++{
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter: ",counter)
		mutex.Unlock()		
	}
	wg.Done()
}