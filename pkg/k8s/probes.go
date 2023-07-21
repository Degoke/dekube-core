package k8s

import (
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/Degoke/dekube-core/pkg/types"
)

type AppProbes struct {
	Liveness *corev1.Probe
	Readiness *corev1.Probe
}

func(f *AppFactory) MakeProbes(r types.AppDeployment) (*AppProbes, error) {
	var handler corev1.ProbeHandler

	if f.Config.HTTPProbe{
		handler = corev1.ProbeHandler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: "/_/health",
				Port: intstr.IntOrString{
					Type:   intstr.Int,
					IntVal: int32(f.Config.RuntimeHTTPPort),
				},
			},
		}
	} else {
		path := filepath.Join("/tmp/", ".lock")
		handler = corev1.ProbeHandler{
			Exec: &corev1.ExecAction{
				Command: []string{"cat", path},
			},
		}
	}

	probes := AppProbes{}
	probes.Readiness = &corev1.Probe{
		ProbeHandler: handler,
		InitialDelaySeconds: f.Config.ReadinessProbe.InitialDelaySeconds,
		TimeoutSeconds:      int32(f.Config.ReadinessProbe.TimeoutSeconds),
		PeriodSeconds:       int32(f.Config.ReadinessProbe.PeriodSeconds),
		SuccessThreshold:    1,
		FailureThreshold:    3,
	}

	probes.Liveness = &corev1.Probe{
		ProbeHandler: handler,
		InitialDelaySeconds: f.Config.LivenessProbe.InitialDelaySeconds,
		TimeoutSeconds:      int32(f.Config.LivenessProbe.TimeoutSeconds),
		PeriodSeconds:       int32(f.Config.LivenessProbe.PeriodSeconds),
		SuccessThreshold:    1,
		FailureThreshold:    3,
	}

	return &probes, nil
}