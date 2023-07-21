package controller

import (
	"strconv"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	glog "k8s.io/klog"

	dekubev1 "github.com/Degoke/dekube-core/pkg/apis/dekube/v1"
)

// newService creates a new NodePort Service for an App resource
func newService(app *dekubev1.App) *corev1.Service {
	var port int32
	s, err := strconv.Atoi(app.Spec.Port); 
	if err == nil {
		glog.Errorf("Failed to convert port: %s", err.Error())
		port = int32(8080)
	}
	port = int32(s)
	
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:        app.Spec.Name,
			Namespace:   app.Namespace,
			Annotations: map[string]string{"prometheus.io.scrape": "false"},
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(app, dekubev1.SchemeGroupVersion.WithKind("App")),
			},
		},
		Spec: corev1.ServiceSpec{
			Type:     corev1.ServiceTypeNodePort,
			Selector: map[string]string{"dekube_app": app.Spec.Name},
			Ports: []corev1.ServicePort{
				{
					Name:     "http",
					Protocol: corev1.ProtocolTCP,
					Port:     port,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: port,
					},
				},
			},
		},
	}
}
