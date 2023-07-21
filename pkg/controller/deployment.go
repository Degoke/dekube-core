package controller

import (
	"encoding/json"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	dekubev1 "github.com/Degoke/dekube-core/pkg/apis/dekube/v1"
	"github.com/google/go-cmp/cmp"
	glog "k8s.io/klog"
)

const (
	annotationAppSpec = "com.dekube.app.spec"
)

// newDeployment creates a new Deployment for a App resource. It also sets
// the appropriate OwnerReferences on the resource so handleObject can discover
// the App resource that 'owns' it.
func newDeployment(app *dekubev1.App) *appsv1.Deployment {
	labels := map[string]string{
		"app":        app.Spec.Name,
		"controller": app.Name,
	}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Spec.Name,
			Namespace: app.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(app, dekubev1.SchemeGroupVersion.WithKind("App")),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: app.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  app.Spec.Name,
							Image: app.Spec.Image,
						},
					},
				},
			},
		},
	}
}

func makeEnvVars(app *dekubev1.App) []corev1.EnvVar {
	envVars := []corev1.EnvVar{}

	if app.Spec.Environment != nil {
		for k, v := range *app.Spec.Environment {
			envVars = append(envVars, corev1.EnvVar{
				Name:  k,
				Value: v,
			})
		}
	}

	return envVars
}

func makeLabels(app *dekubev1.App) map[string]string {
	labels := map[string]string{
		"dekube_app": app.Spec.Name,
		"app":           app.Spec.Name,
		"controller":    app.Name,
	}
	if app.Spec.Labels != nil {
		for k, v := range *app.Spec.Labels {
			labels[k] = v
		}
	}

	return labels
}

func makeAnnotations(app *dekubev1.App) map[string]string {
	annotations := make(map[string]string)

	// disable scraping since the watchdog doesn't expose a metrics endpoint
	annotations["prometheus.io.scrape"] = "false"

	// copy function annotations
	if app.Spec.Annotations != nil {
		for k, v := range *app.Spec.Annotations {
			annotations[k] = v
		}
	}

	// save function spec in deployment annotations
	// used to detect changes in function spec
	specJSON, err := json.Marshal(app.Spec)
	if err != nil {
		glog.Errorf("Failed to marshal function spec: %s", err.Error())
		return annotations
	}

	annotations[annotationAppSpec] = string(specJSON)
	return annotations
}

func deploymentNeedsUpdate(app *dekubev1.App, deployment *appsv1.Deployment) bool {
	prevAppSpecJson := deployment.ObjectMeta.Annotations[annotationAppSpec]
	if prevAppSpecJson == "" {
		// is a new deployment or is an old deployment that is missing the annotation
		return true
	}

	prevAppSpec := &dekubev1.AppSpec{}
	err := json.Unmarshal([]byte(prevAppSpecJson), prevAppSpec)
	if err != nil {
		glog.Errorf("Failed to parse previous function spec: %s", err.Error())
		return true
	}
	prevApp := dekubev1.App{
		Spec: *prevAppSpec,
	}

	if diff := cmp.Diff(prevApp.Spec, app.Spec); diff != "" {
		glog.V(2).Infof("Change detected for %s diff\n%s", app.Name, diff)
		return true
	} else {
		glog.V(3).Infof("No changes detected for %s", app.Name)
	}

	return false
}