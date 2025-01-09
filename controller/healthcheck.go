package controller

import (
	"github.com/gin-gonic/gin"
	"groot/internal/response"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, response.HealthCheckResponse("I'm groot"))
}
