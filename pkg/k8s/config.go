package k8s

type ProbeConfig struct {
	InitialDelaySeconds int32
	TimeoutSeconds      int32
	PeriodSeconds       int32
}

type DeploymentConfig struct {
	RuntimeHTTPPort int32
	HTTPProbe       bool
	ReadinessProbe  *ProbeConfig
	LivenessProbe   *ProbeConfig
	ImagePullPolicy string
}