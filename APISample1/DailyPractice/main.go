package main

import(
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type helloWroldRequest struct{
	Name string
}

type helloWorldResponse struct{
	Msg string `json:"msg"`
}

func main(){
	port := 8080
	http.HandleFunc("/helloworld",helloWorldHandler)
	log.Printf("Server starting on port:%v",port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request){
	var request helloWroldRequest	
	decoder := json.NewDecoder(r.Body)
	derr := decoder.Decode(&request)
	if derr != nil {
		panic("Error in Decoding!!")
	}
	resp := helloWorldResponse{Msg:"Hello World! Welcome " + request.Name}
	encode := json.NewEncoder(w)
	encode.Encode(resp)
}

