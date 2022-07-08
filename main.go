package main

import (
	"fmt"
	"log"
	"net/http"

	"konta.monika/webcalc/calc"

	"github.com/Icikowski/kubeprobes"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	//http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	//panic("")
	fmt.Fprintf(w, "Endpoint called: homePage()")
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router2 := mux.NewRouter().StrictSlash(true)

	Kb := kubeprobes.New()

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/sum/{a}/{b}", calc.Sum).Methods("GET")
	router.HandleFunc("/diff/{a}/{b}", calc.Diff).Methods("GET")
	router.HandleFunc("/mul/{a}/{b}", calc.Mul).Methods("GET")
	router.HandleFunc("/div/{a}/{b}", calc.Div).Methods("GET")
	router.HandleFunc("/factorial/{a}", calc.Fac).Methods("GET")

	router2.HandleFunc("/live", Kb.ServeHTTP).Methods("GET")
	router2.HandleFunc("/ready", Kb.ServeHTTP).Methods("GET")

	go log.Fatal(http.ListenAndServe(":8080", router))
	log.Fatal(http.ListenAndServe(":8081", router2))
	//log.Fatal(http.ListenAndServe(":8080", router))
	//Log.Fatal()
}

func main() {
	handleRequest()
}
