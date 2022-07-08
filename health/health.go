package health

import (
	"net/http"

	"github.com/Icikowski/kubeprobes"
)

func Live(w http.ResponseWriter, r *http.Request) {
	/*live := kubeprobes.NewStatefulProbe()
	kubeprobes.WithLivenessProbes(live.GetProbeFunction())*/
	//live := kubeprobes.NewStatefulProbe()
	//err := kubeprobes.ProbeFunction()
	/*if sq.isAllGreen() {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	//kubeprobes.WithLivenessProbes(live.GetProbeFunction())
	*/
}

func Ready(w http.ResponseWriter, r *http.Request) {
	ready := kubeprobes.NewStatefulProbe()
	kubeprobes.WithReadinessProbes(ready.GetProbeFunction())
}
