package health

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/Icikowski/kubeprobes"
	"github.com/rs/zerolog/log"
)

//Live represents a kubernetees probe of lievness probe.
var Live = kubeprobes.NewStatefulProbe()

//Ready represents a kubernetees probe of lievness probe.
var Ready = kubeprobes.NewStatefulProbe()

//A Liveness represents a lievness status of probe.
func Liveness(w http.ResponseWriter, r *http.Request) {
	err := Live.GetProbeFunction()
	if err != nil {
		http.Error(w, "503 Service Unavailable", http.StatusServiceUnavailable)
		log.Error().Err(errors.New("503")).Msgf("Service Unavailable: %s", err)
	} else {
		(w).WriteHeader(http.StatusOK)
	}
}

//A Readiness represents a readiness status of probe.
func Readiness(w http.ResponseWriter, r *http.Request) {
	err := Ready.GetProbeFunction()
	if err != nil {
		http.Error(w, "503 Service Unavailable", http.StatusServiceUnavailable)
		log.Error().Err(errors.New("503")).Msgf("Service Unavailable: %s", err)
	} else {
		(w).WriteHeader(http.StatusOK)
	}
}

//A CheckServer check server status and returns 500 error when it's down.
func CheckServerStatus() {
	time.Sleep(5 * time.Second)
	_, err := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Error().Err(errors.New("500")).Msgf("Internal Server Error: %s", err)
	} else {
		log.Info().Msgf("Status: %s", "200 Status OK")
	}

}
