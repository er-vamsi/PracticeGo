package main

import(
	"fmt"
	"encoding/json"
	"net/http"
	"log"
)

type helooWorldResponse struct{
	Message string `json:"message"`
}

func main(){
	port := 8080
	http.HandleFunc("/helloworld",helloWroldHandler)
	log.Printf("Server srtarting on port: %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port), nil))	
}

func helloWroldHandler(w http.ResponseWriter, r *http.Request){
	response := &helooWorldResponse{Message: "Hello World!"}
	data, err := json.Marshal(response)

	if err != nil {
		panic("Oooops!")
	}
	fmt.Fprint(w, string(data))
}