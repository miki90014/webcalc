package health

import (
	"errors"
	"net/http"

	"github.com/Icikowski/kubeprobes"
	"github.com/rs/zerolog/log"
)

//Live represents a kubernetees probe of lievness probe.
var Live = kubeprobes.NewStatefulProbe()

//Ready represents a kubernetees probe of lievness probe.
var Ready = kubeprobes.NewStatefulProbe()

//Liveness represents a lievness status of probe.
func Liveness(w http.ResponseWriter, r *http.Request) {
	err := Live.GetProbeFunction()
	if err != nil {
		http.Error(w, "503 Service Unavailable", http.StatusServiceUnavailable)
		log.Error().Err(errors.New("503")).Msgf("Service Unavailable: %s", err)
	} else {
		(w).WriteHeader(http.StatusOK)
	}
}

//Readiness represents a readiness status of probe.
func Readiness(w http.ResponseWriter, r *http.Request) {
	err := Ready.GetProbeFunction()
	if err != nil {
		http.Error(w, "503 Service Unavailable", http.StatusServiceUnavailable)
		log.Error().Err(errors.New("503")).Msgf("Service Unavailable: %s", err)
	} else {
		(w).WriteHeader(http.StatusOK)
	}
}
