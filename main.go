package main

import (
	"flag"
	"log"
	"time"

	"github.com/Degoke/dekube-core/pkg/signals"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"

	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	clientset "github.com/Degoke/dekube-core/pkg/client/clientset/versioned"
	informers "github.com/Degoke/dekube-core/pkg/client/informers/externalversions"
	"github.com/Degoke/dekube-core/pkg/config"
	dekubeController "github.com/Degoke/dekube-core/pkg/controller"
	"github.com/Degoke/dekube-core/pkg/k8s"
	"github.com/Degoke/dekube-core/pkg/types"

	"github.com/Degoke/dekube-core/pkg/handlers"

	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	masterURL  string
	kubeconfig string
	verbose    bool
)

type ServerSetup struct {
	appFactory     k8s.AppFactory
}

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	// set up signals so we handle the shutdown signal gracefully
	ctx := signals.SetupSignalHandler()
	logger := klog.FromContext(ctx)

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		logger.Error(err, "Error building kubeconfig")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		logger.Error(err, "Error building kubernetes clientset")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	dekubeClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		logger.Error(err, "Error building kubernetes clientset")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	readConfig := config.ReadConfig{}
	osEnv := types.OsEnv{}
	config, err := readConfig.Read(osEnv)

	if err != nil {
		log.Fatalf("Error reading config: %s", err.Error())
	}

	config.Fprint(verbose)

	deployConfig := k8s.DeploymentConfig{
		RuntimeHTTPPort: 8080,
		HTTPProbe:       config.HTTPProbe,
		ReadinessProbe: &k8s.ProbeConfig{
			InitialDelaySeconds: int32(2),
			TimeoutSeconds:      int32(1),
			PeriodSeconds:       int32(2),
		},
		LivenessProbe: &k8s.ProbeConfig{
			InitialDelaySeconds: int32(2),
			TimeoutSeconds:      int32(1),
			PeriodSeconds:       int32(2),
		},
		ImagePullPolicy:   config.ImagePullPolicy,
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	dekubeInformerFactory := informers.NewSharedInformerFactory(dekubeClient, time.Second*30)

	controller := dekubeController.NewController(ctx, kubeClient, dekubeClient,
		kubeInformerFactory.Apps().V1().Deployments(),
		dekubeInformerFactory.Dekube().V1().Apps())

	// notice that there is no need to run Start methods in a separate goroutine. (i.e. go kubeInformerFactory.Start(ctx.done())
	// Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	kubeInformerFactory.Start(ctx.Done())
	dekubeInformerFactory.Start(ctx.Done())

	factory := k8s.NewAppFactory(dekubeClient, deployConfig)

	setup := ServerSetup{
		appFactory: factory,
	}

	go runServer(&setup)

	if err = controller.Run(ctx, 2); err != nil {
		logger.Error(err, "Error running controller")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.BoolVar(&verbose, "verbose", false, "Print verbose config information")
}

func runServer(s *ServerSetup) {
	factory := s.appFactory

	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time": time.Now().Format(time.RFC3339),
		})
	})

	

	r.GET("/dekube/v1/apps", handlers.MakeReaderHandler(factory))
	r.POST("/dekube/v1/apps", handlers.MakeDeployHandler(factory))
	r.Run(":8080")
}