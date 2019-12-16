package main

import ("fmt"
		"sync"
		"time"
		"runtime")

var wg sync.WaitGroup

func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	//fmt.Println("Number of CPUs: ", runtime.NumCPU())
}

func main(){
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo(){
	for i := 0; i <= 10; i++{
		fmt.Println("foo: ",i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar(){
	for i := 0; i <= 10; i++{
		fmt.Println("bar: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}