package k8s

import (
	dekubeclientset "github.com/Degoke/dekube-core/pkg/client/clientset/versioned"
)

type AppFactory struct {
	Client dekubeclientset.Interface
	Config DeploymentConfig
}

func NewAppFactory(clientset dekubeclientset.Interface, config DeploymentConfig) AppFactory {
	return AppFactory{
		Client: clientset,
		Config: config,
	}
}