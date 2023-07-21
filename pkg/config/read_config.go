package config

import (
	"fmt"
	"log"

	types "github.com/Degoke/dekube-core/pkg/types"
)

var validPullPolicyOptions = map[string]bool{
	"Always":       true,
	"IfNotPresent": true,
	"Never":        true,
}

// ReadConfig constitutes config from env variables
type ReadConfig struct {
}

// Read fetches config from environmental variables.
func (ReadConfig) Read(hasEnv types.HasEnv) (BootstrapConfig, error) {
	cfg := BootstrapConfig{}

	dekubeConfig, err := types.ReadConfig{}.Read(hasEnv)
	if err != nil {
		return cfg, err
	}

	cfg.DekubeConfig = *dekubeConfig

	httpProbe := types.ParseBoolValue(hasEnv.Getenv("http_probe"), false)
	setNonRootUser := types.ParseBoolValue(hasEnv.Getenv("set_nonroot_user"), false)

	imagePullPolicy := types.ParseString(hasEnv.Getenv("image_pull_policy"), "Always")

	if !validPullPolicyOptions[imagePullPolicy] {
		return cfg, fmt.Errorf("invalid image_pull_policy configured: %s", imagePullPolicy)
	}

	cfg.DefaultAppNamespace = types.ParseString(hasEnv.Getenv("app_namespace"), "dekube-app")

	cfg.HTTPProbe = httpProbe
	cfg.SetNonRootUser = setNonRootUser

	cfg.ImagePullPolicy = imagePullPolicy

	return cfg, nil
}

// BootstrapConfig contains the server configuration values as well as default
// Function configuration parameters that are passed to the function factory.
type BootstrapConfig struct {
	// HTTPProbe when set to true switches readiness and liveness probe to
	// access /_/health over HTTP instead of accessing /tmp/.lock.
	HTTPProbe bool

	// SetNonRootUser determines if the Function is deployed with a overridden
	// non-root user id.  Currently this is preconfigured to the uid 12000.
	SetNonRootUser bool

	// ImagePullPolicy controls the ImagePullPolicy set on the Function Deployment.
	ImagePullPolicy string

	// DefaultFunctionNamespace defines which namespace in which Functions are deployed.
	// Value is set via the function_namespace environment variable. If the
	// variable is not set, it is set to "default".
	DefaultAppNamespace string

	DekubeConfig types.DekubeConfig
}

// Fprint pretty-prints the config with the stdlib logger. One line per config value.
// When the verbose flag is set to false, it prints the same output as prior to
// the 0.12.0 release.
func (c BootstrapConfig) Fprint(verbose bool) {
	log.Printf("HTTP Read Timeout: %s\n", c.DekubeConfig.GetReadTimeout())
	log.Printf("HTTP Write Timeout: %s\n", c.DekubeConfig.WriteTimeout)
	log.Printf("ImagePullPolicy: %s\n", c.ImagePullPolicy)
	log.Printf("DefaultFunctionNamespace: %s\n", c.DefaultAppNamespace)

	if verbose {
		log.Printf("MaxIdleConns: %d\n", c.DekubeConfig.MaxIdleConns)
		log.Printf("MaxIdleConnsPerHost: %d\n", c.DekubeConfig.MaxIdleConnsPerHost)
		log.Printf("HTTPProbe: %v\n", c.HTTPProbe)
		log.Printf("SetNonRootUser: %v\n", c.SetNonRootUser)
	}
}
