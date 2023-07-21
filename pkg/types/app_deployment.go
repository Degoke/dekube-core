package types

type AppDeployment struct {
	Name string `json:"name"`
	Image string `json:"image"`
	Namespace string `json:"namespace"`
	EnvVars map[string]string `json:"envVars,omitempty"`
	Secrets []string `json:"secrets,omitempty"`
	Labels *map[string]string `json:"labels,omitempty"`
	Annotations *map[string]string `json:"annotations,omitempty"`
	Limits *AppResources `json:"limits,omitempty"`
	Requests *AppResources `json:"requests,omitempty"`
	Replicas       *int32 `json:"replicas"`
	Domain string `json:"domain"`
	Port string `json:"port"`
}

type AppResources struct {
	Memory string `json:"memory,omitempty"`
	CPU    string `json:"cpu,omitempty"`
}