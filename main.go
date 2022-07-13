package main

import (
	"log"
	"net/http"
	"sync"

	"konta.monika/webcalc/calc"
	"konta.monika/webcalc/health"

	"github.com/gorilla/mux"
)

func main() {
	router1 := mux.NewRouter().StrictSlash(true)
	router2 := mux.NewRouter().StrictSlash(true)

	router1.HandleFunc("/sum/{a}/{b}", calc.Sum).Methods("GET")
	router1.HandleFunc("/diff/{a}/{b}", calc.Diff).Methods("GET")
	router1.HandleFunc("/mul/{a}/{b}", calc.Mul).Methods("GET")
	router1.HandleFunc("/div/{a}/{b}", calc.Div).Methods("GET")
	router1.HandleFunc("/factorial/{a}", calc.Fac).Methods("GET")

	router2.HandleFunc("/live", health.Kp.ServeHTTP).Methods("GET")
	router2.HandleFunc("/ready", health.Kp.ServeHTTP).Methods("GET")

	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		log.Fatal(http.ListenAndServe(":8080", router1))
		wg.Done()
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":8081", router2))
		wg.Done()
	}()

	health.Live.MarkAsUp()
	health.Ready.MarkAsUp()
	wg.Wait()
}
