package health

import (
	"errors"
	"net/http"

	"github.com/Icikowski/kubeprobes"
	"github.com/rs/zerolog/log"
)

var Live = kubeprobes.NewStatefulProbe()
var Ready = kubeprobes.NewStatefulProbe()

func Liveness(w http.ResponseWriter, r *http.Request) {
	err := Live.GetProbeFunction()
	if err != nil {
		http.Error(w, "503 Service Unavailable", http.StatusServiceUnavailable)
		log.Error().Err(errors.New("503")).Msgf("Service Unavailable: %s", err)
	} else {
		(w).WriteHeader(http.StatusOK)
	}
}

func Readiness(w http.ResponseWriter, r *http.Request) {
	err := Ready.GetProbeFunction()
	if err != nil {
		http.Error(w, "503 Service Unavailable", http.StatusServiceUnavailable)
		log.Error().Err(errors.New("503")).Msgf("Service Unavailable: %s", err)
	} else {
		(w).WriteHeader(http.StatusOK)
	}
}
