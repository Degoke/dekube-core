package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/Degoke/dekube-core/pkg/k8s"

	"github.com/Degoke/dekube-core/pkg/types"
	// appsv1 "k8s.io/api/apps/v1"
	// apiv1 "k8s.io/api/core/v1"
	// corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/apimachinery/pkg/util/intstr"

	dekubev1 "github.com/Degoke/dekube-core/pkg/apis/dekube/v1"

	"github.com/gin-gonic/gin"
	"net/http"
)

// initialReplicasCount how many replicas to start of creating for a function
const initialReplicasCount = 1

// MakeDeployHandler creates a handler to create new App in the cluster
func MakeDeployHandler(factory k8s.AppFactory) gin.HandlerFunc {
	// secrets := k8s.NewSecretsClient(factory.Client)

	return func(c *gin.Context) {
		var app types.AppDeployment

		c.Bind(&app)

		// if err := ValidateDeployRequest(&request); err != nil {
		// 	wrappedErr := fmt.Errorf("validation failed: %s", err.Error())
		// 	http.Error(w, wrappedErr.Error(), http.StatusBadRequest)
		// 	return
		// }

		namespace := app.Namespace

		// existingSecrets, err := secrets.GetSecrets(namespace, request.Secrets)
		// if err != nil {
		// 	wrappedErr := fmt.Errorf("unable to fetch secrets: %s", err.Error())
		// 	http.Error(w, wrappedErr.Error(), http.StatusBadRequest)
		// 	return
		// }

		deploymentSpec, specErr := makeDeploymentSpec(app, factory)

		// var profileList []k8s.Profile
		// if request.Annotations != nil {
		// 	profileNamespace := factory.Config.ProfilesNamespace
		// 	profileList, err = factory.GetProfiles(ctx, profileNamespace, *request.Annotations)
		// 	if err != nil {
		// 		wrappedErr := fmt.Errorf("failed create Deployment spec: %s", err.Error())
		// 		log.Println(wrappedErr)
		// 		http.Error(w, wrappedErr.Error(), http.StatusBadRequest)
		// 		return
		// 	}
		// }
		// for _, profile := range profileList {
		// 	factory.ApplyProfile(profile, deploymentSpec)
		// }

		if specErr != nil {
			wrappedErr := fmt.Errorf("failed create Deployment spec: %s", specErr.Error())
			log.Println(wrappedErr)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": wrappedErr.Error(),
			})
			return
		}

		deploy := factory.Client.DekubeV1().Apps(namespace)

		_, err := deploy.Create(context.TODO(), deploymentSpec, metav1.CreateOptions{})
		if err != nil {
			wrappedErr := fmt.Errorf("unable create Deployment: %s", err.Error())
			log.Println(wrappedErr)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": wrappedErr.Error(),
			})
			return
		}

		log.Printf("Deployment created: %s.%s\n", app.Name, namespace)

		// service := factory.Client.CoreV1().Services(namespace)
		// serviceSpec, err := makeServiceSpec(request, factory)
		// if err != nil {
		// 	wrappedErr := fmt.Errorf("failed create Service spec: %s", err.Error())
		// 	log.Println(wrappedErr)
		// 	http.Error(w, wrappedErr.Error(), http.StatusBadRequest)
		// 	return
		// }

		// if _, err = service.Create(context.TODO(), serviceSpec, metav1.CreateOptions{}); err != nil {
		// 	wrappedErr := fmt.Errorf("failed create Service: %s", err.Error())
		// 	log.Println(wrappedErr)
		// 	http.Error(w, wrappedErr.Error(), http.StatusBadRequest)
		// 	return
		// }

		// log.Printf("Service created: %s.%s\n", request.Service, namespace)

		c.JSON(http.StatusCreated, gin.H{
			"message": "created",
			"status":  "success",
		})
	}
}

func makeDeploymentSpec(request types.AppDeployment, factory k8s.AppFactory) (*dekubev1.App, error) {
	// envVars := buildEnvVars(&request)

	initialReplicas := int32p(initialReplicasCount)
	// labels := map[string]string{
	// 	"faas_function": request.Service,
	// }

	// if request.Labels != nil {
	// 	if min := getMinReplicaCount(*request.Labels); min != nil {
	// 		initialReplicas = min
	// 	}
	// 	for k, v := range *request.Labels {
	// 		labels[k] = v
	// 	}
	// }

	// nodeSelector := createSelector(request.Constraints)

	// resources, err := createResources(request)

	// if err != nil {
	// 	return nil, err
	// }

	// var imagePullPolicy apiv1.PullPolicy
	// switch factory.Config.ImagePullPolicy {
	// case "Never":
	// 	imagePullPolicy = apiv1.PullNever
	// case "IfNotPresent":
	// 	imagePullPolicy = apiv1.PullIfNotPresent
	// default:
	// 	imagePullPolicy = apiv1.PullAlways
	// }

	// annotations, err := buildAnnotations(request)
	// if err != nil {
	// 	return nil, err
	// }

	// probes, err := factory.MakeProbes(request)
	// if err != nil {
	// 	return nil, err
	// }

	deploymentSpec := &dekubev1.App{
		ObjectMeta: metav1.ObjectMeta{
			Name:        request.Name,
			// Annotations: annotations,
			Labels: map[string]string{
				"dekube_app": request.Name,
			},
		},
		Spec: dekubev1.AppSpec{
			Image: request.Image,
			// Env:   envVars,
			Name: request.Name,
			Replicas: initialReplicas,
		},
	}

	// deploymentSpec := &appsv1.Deployment{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name:        request.Name,
	// 		// Annotations: annotations,
	// 		Labels: map[string]string{
	// 			"dekube_app": request.Name,
	// 		},
	// 	},
	// 	Spec: appsv1.DeploymentSpec{
	// 		Selector: &metav1.LabelSelector{
	// 			MatchLabels: map[string]string{
	// 				"dekube_app": request.Name,
	// 			},
	// 		},
	// 		Replicas: initialReplicas,
	// 		Strategy: appsv1.DeploymentStrategy{
	// 			Type: appsv1.RollingUpdateDeploymentStrategyType,
	// 			RollingUpdate: &appsv1.RollingUpdateDeployment{
	// 				MaxUnavailable: &intstr.IntOrString{
	// 					Type:   intstr.Int,
	// 					IntVal: int32(0),
	// 				},
	// 				MaxSurge: &intstr.IntOrString{
	// 					Type:   intstr.Int,
	// 					IntVal: int32(1),
	// 				},
	// 			},
	// 		},
	// 		RevisionHistoryLimit: int32p(10),
	// 		Template: apiv1.PodTemplateSpec{
	// 			ObjectMeta: metav1.ObjectMeta{
	// 				Name:        request.Name,
	// 				Labels:      map[string]string{
	// 					"dekube_app": request.Name,
	// 				},
	// 				// Annotations: annotations,
	// 			},
	// 			Spec: apiv1.PodSpec{
	// 				// NodeSelector: nodeSelector,
	// 				Containers: []apiv1.Container{
	// 					{
	// 						Name:  request.Name,
	// 						Image: request.Image,
	// 						Ports: []apiv1.ContainerPort{
	// 							{
	// 								Name:          "http",
	// 								ContainerPort: factory.Config.RuntimeHTTPPort,
	// 								Protocol:      corev1.ProtocolTCP,
	// 							},
	// 						},
	// 						// Env:             envVars,
	// 						// Resources:       *resources,
	// 						ImagePullPolicy: imagePullPolicy,
	// 						LivenessProbe:   probes.Liveness,
	// 						ReadinessProbe:  probes.Readiness,
	// 						// SecurityContext: &corev1.SecurityContext{
	// 						// 	ReadOnlyRootFilesystem: &request.ReadOnlyRootFilesystem,
	// 						// },
	// 					},
	// 				},
	// 				RestartPolicy: corev1.RestartPolicyAlways,
	// 				DNSPolicy:     corev1.DNSClusterFirst,
	// 			},
	// 		},
	// 	},
	// }

	// factory.ConfigureReadOnlyRootFilesystem(request, deploymentSpec)
	// factory.ConfigureContainerUserID(deploymentSpec)

	// if err := factory.ConfigureSecrets(request, deploymentSpec, existingSecrets); err != nil {
	// 	return nil, err
	// }

	return deploymentSpec, nil
}

func int32p(i int32) *int32 {
	return &i
}