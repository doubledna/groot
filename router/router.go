package router

import (
	"context"
	"fmt"
	tasksv1 "groot/controller/tasks/v1"
	genv1 "groot/gen/v1"
	"groot/internal"
	"groot/internal/config"
	"net/http"
	"os"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	"go.uber.org/zap"
)

type HTTPServer struct {
	Server *http.Server
}

func NewHTTPServer() (*HTTPServer, error) {
	var s HTTPServer
	r := gin.New()
	logger, _ := zap.NewProduction()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	// init database resource
	if err := internal.Init(); err != nil {
		return nil, err
	}

	healthCheck := r.Group("/healthz")
	healthCheckRegister(healthCheck)

	// openapi
	swagger, err := genv1.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil
	r.Use(middleware.OapiRequestValidator(swagger))

	// task routes
	taskServer := tasksv1.NewTaskStore()
	genv1.RegisterHandlers(r, taskServer)

	s.Server = &http.Server{Addr: config.GetString("web.address"), Handler: r}
	return &s, nil
}

func (s *HTTPServer) Run() error {
	fmt.Printf("starting server %s\n", config.GetString("web.address"))
	return s.Server.ListenAndServe()
}

func (s *HTTPServer) Close() error {
	fmt.Println("shutting down server")
	defer fmt.Println("server exited")
	return s.Server.Shutdown(context.Background())
}
