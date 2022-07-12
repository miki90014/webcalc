package health

import (
	"github.com/Icikowski/kubeprobes"
)

//Live represents a kubernetees probe of lievness probe.
var Live = kubeprobes.NewStatefulProbe()

//Ready represents a kubernetees probe of lievness probe.
var Ready = kubeprobes.NewStatefulProbe()

var Kp = kubeprobes.New(
	kubeprobes.WithLivenessProbes(Live.GetProbeFunction()),
	kubeprobes.WithReadinessProbes(Ready.GetProbeFunction()),
)
