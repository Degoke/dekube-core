package controller

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	dekubev1 "github.com/Degoke/dekube-core/pkg/apis/dekube/v1"

)

// makeResources creates deployment resource limits and requests requirements from App specs
func makeResources(app *dekubev1.App) (*corev1.ResourceRequirements, error) {
	resources := &corev1.ResourceRequirements{
		Limits:   corev1.ResourceList{},
		Requests: corev1.ResourceList{},
	}

	// Set Memory limits
	if app.Spec.Limits != nil && len(app.Spec.Limits.Memory) > 0 {
		qty, err := resource.ParseQuantity(app.Spec.Limits.Memory)
		if err != nil {
			return resources, err
		}
		resources.Limits[corev1.ResourceMemory] = qty
	}
	if app.Spec.Requests != nil && len(app.Spec.Requests.Memory) > 0 {
		qty, err := resource.ParseQuantity(app.Spec.Requests.Memory)
		if err != nil {
			return resources, err
		}
		resources.Requests[corev1.ResourceMemory] = qty
	}

	// Set CPU limits
	if app.Spec.Limits != nil && len(app.Spec.Limits.CPU) > 0 {
		qty, err := resource.ParseQuantity(app.Spec.Limits.CPU)
		if err != nil {
			return resources, err
		}
		resources.Limits[corev1.ResourceCPU] = qty
	}
	if app.Spec.Requests != nil && len(app.Spec.Requests.CPU) > 0 {
		qty, err := resource.ParseQuantity(app.Spec.Requests.CPU)
		if err != nil {
			return resources, err
		}
		resources.Requests[corev1.ResourceCPU] = qty
	}

	return resources, nil
}
