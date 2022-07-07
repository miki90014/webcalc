package main

import (
	"fmt"
	"log"
	"net/http"

	"konta.monika/webcalc/calc"

	"github.com/gorilla/mux"
	//"github.com/rs/zerolog"
	//"github.com/rs/zerolog/log"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	//http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	//panic("")
	fmt.Fprintf(w, "Endpoint called: homePage()")
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/sum/{a}/{b}", calc.Sum).Methods("GET")
	router.HandleFunc("/diff/{a}/{b}", calc.Diff).Methods("GET")
	router.HandleFunc("/mul/{a}/{b}", calc.Mul).Methods("GET")
	router.HandleFunc("/div/{a}/{b}", calc.Div).Methods("GET")
	router.HandleFunc("/factorial/{a}", calc.Fac).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
	//log.Fatal(http.ListenAndServe(":8080", router))
	//Log.Fatal()
}

func main() {
	handleRequest()
}
