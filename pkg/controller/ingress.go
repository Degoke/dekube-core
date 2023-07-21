package controller

import (
	"strconv"

	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	glog "k8s.io/klog"

	dekubev1 "github.com/Degoke/dekube-core/pkg/apis/dekube/v1"
)

// newIngress creates a new nginx Ingress for an App resource
func newIngress(app *dekubev1.App) *networkv1.Ingress {
	var port int32
	s, err := strconv.Atoi(app.Spec.Port); 
	if err == nil {
		glog.Errorf("Failed to convert port: %s", err.Error())
		port = int32(8080)
	}

	port = int32(s)

	ingressClassName := "dekube-nginx"
	return &networkv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:        app.Spec.Name,
			Namespace:   app.Namespace,
			Annotations: map[string]string{
				"prometheus.io.scrape": "false",
				"kubernetes.io/ingress.class": "nginx",
    			"kubernetes.io/tls-acme": "true",
			},
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(app, dekubev1.SchemeGroupVersion.WithKind("App")),
			},
		},
		Spec: networkv1.IngressSpec{
			IngressClassName: &ingressClassName,
			Rules: []networkv1.IngressRule{
				{
					Host: app.Spec.Domain,
					IngressRuleValue: networkv1.IngressRuleValue{
					HTTP: &networkv1.HTTPIngressRuleValue{
						Paths: []networkv1.HTTPIngressPath{
							{
								Path: "/",
								Backend: networkv1.IngressBackend{
									Service: &networkv1.IngressServiceBackend{
										Name: app.Spec.Name,
										Port: networkv1.ServiceBackendPort{
											Name:   "http",
											Number: port,
										},
									},
								},
							},
						},
					},
				},
				},
			
		},
	},
}
}
