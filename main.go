package main

import (
	"log"
	"net/http"
	"sync"

	"konta.monika/webcalc/calc"
	"konta.monika/webcalc/health"

	"github.com/Icikowski/kubeprobes"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {

}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router2 := mux.NewRouter().StrictSlash(true)

	kp := kubeprobes.New(
		kubeprobes.WithLivenessProbes(health.Live.GetProbeFunction()),
		kubeprobes.WithReadinessProbes(health.Ready.GetProbeFunction()),
	)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/sum/{a}/{b}", calc.Sum).Methods("GET")
	router.HandleFunc("/diff/{a}/{b}", calc.Diff).Methods("GET")
	router.HandleFunc("/mul/{a}/{b}", calc.Mul).Methods("GET")
	router.HandleFunc("/div/{a}/{b}", calc.Div).Methods("GET")
	router.HandleFunc("/factorial/{a}", calc.Fac).Methods("GET")

	router2.HandleFunc("/live", health.Liveness).Methods("GET")
	router2.HandleFunc("/ready", health.Readiness).Methods("GET")

	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		log.Fatal(http.ListenAndServe(":8080", router))
		wg.Done()
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":8081", kp))
		log.Fatal(http.ListenAndServe(":8081", router2))
		wg.Done()
	}()

	go func() {
		for {
			health.CheckServerStatus()
		}
	}()

	health.Live.MarkAsUp()
	health.Ready.MarkAsUp()
	wg.Wait()

}

func main() {
	handleRequest()
}
