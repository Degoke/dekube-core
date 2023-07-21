package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// App is a specification for a App resource
type App struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppSpec   `json:"spec"`

	Status AppStatus `json:"status"`
}

// AppSpec is the spec for a App resource
type AppSpec struct {
	Name string `json:"name"`
	Image string `json:"image"`
	Replicas       *int32 `json:"replicas"`
	Annotations *map[string]string `json:"annotations,omitempty"`
	Environment *map[string]string `json:"environment,omitempty"`
	Secrets []string `json:"secrets,omitempty"`
	Labels *map[string]string `json:"labels,omitempty"`
	Limits *AppResources `json:"limits,omitempty"`
	Requests *AppResources `json:"requests,omitempty"`
	Domain string `json:"domain"`
	Port string `json:"port"`
}

// AppStatus is the status for a App resource
type AppStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// AppResources is used to set CPU and memory limits and requests
type AppResources struct {
	Memory string `json:"memory,omitempty"`
	CPU    string `json:"cpu,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of App resources
type AppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []App `json:"items"`
}