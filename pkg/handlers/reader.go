package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Degoke/dekube-core/pkg/k8s"
	"github.com/gin-gonic/gin"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MakeReaderHandler(factory k8s.AppFactory) gin.HandlerFunc {
	return func(c *gin.Context) {
		deploy := factory.Client.DekubeV1().Apps("default")

		apps, err := deploy.List(context.TODO(), metav1.ListOptions{})

		if err != nil {
			wrappedErr := fmt.Errorf("unable to fetch apps: %s", err.Error())
			log.Println(wrappedErr)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": wrappedErr.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, apps)
	}
}