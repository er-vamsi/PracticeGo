package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type helloWorldResponse struct{
	Message string
}

func main(){
	port := 8080
	http.HandleFunc("/helloworld",helloWorldHandler)
	http.Handle()

	log.Printf("Server starting on port %v\n",port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}
func helloWorldHandler(w http.ResponseWriter, r *http.Request){
	response := &helloWorldResponse{Message: "Hello World!"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooooops!")
	}	
	fmt.Fprint(w, string(data))
}