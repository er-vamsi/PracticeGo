package main

import (
	"log"
	"net/http"

	"github.com/er-vamsi/PracticeGo/TunaGoApp/web/controller"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/home", controller.HomePageHandler).Methods("GET")

	r.HandleFunc("/register", controller.RegisterPageHandler).Methods("GET")
	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("web/templates/styles/"))))
	r.HandleFunc("/register", controller.RegisterHandler).Methods("POST")

	r.HandleFunc("/login", controller.LoginPageHandler).Methods("GET")
	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("web/templates/styles/"))))
	r.HandleFunc("/login", controller.LoginHandler).Methods("POST")

	r.HandleFunc("/profile", controller.ProfileHandler).Methods("GET")

	log.Print("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
