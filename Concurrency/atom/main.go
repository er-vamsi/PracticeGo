package main
import ("fmt"
		"time"
		"sync"
		"sync/atomic")

var wg sync.WaitGroup
var counter int64


func main(){
	wg.Add(2)
	go increment("Foo: ")
	go increment("Bar: ")
	wg.Wait()
}

func increment(s string){
	for i:=0;i<5;i++{
		time.Sleep(time.Duration(3 * time.Millisecond))
		atomic.AddInt64(&counter,1)
		fmt.Println(s,counter)

	}
	wg.Done()
}