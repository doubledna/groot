package router

import (
	"github.com/gin-gonic/gin"
	"groot/controller"
)

func healthCheckRegister(router *gin.RouterGroup) {
	router.GET("", controller.HealthCheck)
}
