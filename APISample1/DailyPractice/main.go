package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type validationContextKey string

type HelloWorldRequest struct {
	Name string `json:"name"`
}

type HelloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 9090

	handler := newValidationHandler(newHelloWroldHandler())

	http.Handle("/helloworld", handler)

	log.Printf("Server starting at %v port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var req HelloWorldRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), validationContextKey("name"), req.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(rw, r)
}

type helloWroldHandler struct{}

func newHelloWroldHandler() http.Handler {
	return helloWroldHandler{}
}

func (h helloWroldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validationContextKey("name")).(string)
	resp := HelloWorldResponse{Message: "Hi " + name}
	encoder := json.NewEncoder(rw)

	encoder.Encode(resp)
}
