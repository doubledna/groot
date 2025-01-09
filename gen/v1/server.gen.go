// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package v1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all task
	// (GET /api/v1/task)
	GetTask(c *gin.Context)
	// Create a new task
	// (POST /api/v1/task)
	CreateTask(c *gin.Context)
	// Create a new task
	// (PUT /api/v1/task)
	UpdateTask(c *gin.Context)
	// Returns a task by mode
	// (GET /api/v1/task/mode/{mode})
	GetTaskByMode(c *gin.Context, mode string)
	// Delete a task by name
	// (DELETE /api/v1/task/{name})
	DeleteTask(c *gin.Context, name string)
	// Returns a task by name
	// (GET /api/v1/task/{name})
	GetTaskByName(c *gin.Context, name string)
	// Returns all task type
	// (GET /api/v1/task_type)
	GetTaskType(c *gin.Context)
	// Create a new task type
	// (POST /api/v1/task_type)
	CreateTaskType(c *gin.Context)
	// Delete a task type by name
	// (DELETE /api/v1/task_type/{name})
	DeleteTaskType(c *gin.Context, name string)
	// Returns a task type by name
	// (GET /api/v1/task_type/{name})
	GetTaskTypeByName(c *gin.Context, name string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetTask operation middleware
func (siw *ServerInterfaceWrapper) GetTask(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetTask(c)
}

// CreateTask operation middleware
func (siw *ServerInterfaceWrapper) CreateTask(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateTask(c)
}

// UpdateTask operation middleware
func (siw *ServerInterfaceWrapper) UpdateTask(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateTask(c)
}

// GetTaskByMode operation middleware
func (siw *ServerInterfaceWrapper) GetTaskByMode(c *gin.Context) {

	var err error

	// ------------- Path parameter "mode" -------------
	var mode string

	err = runtime.BindStyledParameterWithOptions("simple", "mode", c.Param("mode"), &mode, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter mode: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetTaskByMode(c, mode)
}

// DeleteTask operation middleware
func (siw *ServerInterfaceWrapper) DeleteTask(c *gin.Context) {

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithOptions("simple", "name", c.Param("name"), &name, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter name: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteTask(c, name)
}

// GetTaskByName operation middleware
func (siw *ServerInterfaceWrapper) GetTaskByName(c *gin.Context) {

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithOptions("simple", "name", c.Param("name"), &name, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter name: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetTaskByName(c, name)
}

// GetTaskType operation middleware
func (siw *ServerInterfaceWrapper) GetTaskType(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetTaskType(c)
}

// CreateTaskType operation middleware
func (siw *ServerInterfaceWrapper) CreateTaskType(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateTaskType(c)
}

// DeleteTaskType operation middleware
func (siw *ServerInterfaceWrapper) DeleteTaskType(c *gin.Context) {

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithOptions("simple", "name", c.Param("name"), &name, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter name: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteTaskType(c, name)
}

// GetTaskTypeByName operation middleware
func (siw *ServerInterfaceWrapper) GetTaskTypeByName(c *gin.Context) {

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithOptions("simple", "name", c.Param("name"), &name, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter name: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetTaskTypeByName(c, name)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/api/v1/task", wrapper.GetTask)
	router.POST(options.BaseURL+"/api/v1/task", wrapper.CreateTask)
	router.PUT(options.BaseURL+"/api/v1/task", wrapper.UpdateTask)
	router.GET(options.BaseURL+"/api/v1/task/mode/:mode", wrapper.GetTaskByMode)
	router.DELETE(options.BaseURL+"/api/v1/task/:name", wrapper.DeleteTask)
	router.GET(options.BaseURL+"/api/v1/task/:name", wrapper.GetTaskByName)
	router.GET(options.BaseURL+"/api/v1/task_type", wrapper.GetTaskType)
	router.POST(options.BaseURL+"/api/v1/task_type", wrapper.CreateTaskType)
	router.DELETE(options.BaseURL+"/api/v1/task_type/:name", wrapper.DeleteTaskType)
	router.GET(options.BaseURL+"/api/v1/task_type/:name", wrapper.GetTaskTypeByName)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xYXW/bNhf+KwTf96IdNMtph63w1dp1K3LRZGjSqyIYGOrYYkORDEnZMQz994EfsmSL",
	"jpM0HeItCBDbIs/hOc95zge1wlRWSgoQ1uDJChtaQkX819+1ltp9IZyfTvHkywr/X8MUT/D/8k4ojxL5",
	"JzBKCgN/SF0Ri5tshZWWCrRl4PUVxBL/CYZqpiyTAk+wBltrgcAdhliBXhQwJTW3iBn049FLnOFpUDjB",
	"TNiff8IZtksF4SfMQOOmybCG65ppKPDkSzjnYr1NXn4FanFz0WT4A9hzYq7+KafcIuLMWJxhZqHyErcd",
	"541r1qYTrcnyAf6d+5Un7KM38Bv8PIFFG8dNU6kGYuGtHZpriblCYRlZVkHHI2M1EzNnDtVSnCmgQ2m3",
	"gowCmqGp1AhuSKU4oB/Wf1Kj12OTUgpzEAl7jJUakLcKboDW7jHSYGpuk2qumCiGWiqir5AtwUnKWlNA",
	"ThAR4zWn9FSygB3ouCX0QgoKzh0FmsmC0ZcpLYJUu7T4pYSIIksuScIHRTSpwII2qA0/WpQgkK6FYGKG",
	"yE5nAl5DlYsSbAnaI+ONWhCDTE0pGDOtOV9GzKHotF5KyYEIT8teDm152Cr0MFuJFiWjZXfOJXApZskA",
	"1qq4jZpheQc1tzLDc6FnZoxHDG6Pxx3qa6xaQmZdrvRsG6bbOtlaQDYT7iG09D/uySovn6ZWGhy/N+XO",
	"VrkblpBkgugohWiAeG9TyjC0/XNTU+x0Imhwz1JZCsaQ2W12tDuSWTEFDYImxGdsDqiUC0fdkoiCQxum",
	"oNXVNm9frTmCGwtaEM6Xe0GPqHRGdTa0QKRicb9m3Fb9YYdiCQ5+Fuy6BjdSyOk6QR8wTrBiRxO6f6ft",
	"Z9I3ONHmzyN54rY5OgbyC0uozwqoCOOuOStmgVS/mgWZzUCPmGyza4LPwjP09s9jdA6kcqVEO6HSWmUm",
	"ed4TarIt3z5oKS06VSCc+OvR2DdYNmW0TQvOKAjjAY4Hfjw+H5whFYhQZUZSz/IolFfMU8yCrszp9Az0",
	"nLmMSJiW+z25A5BZ3vfLm4gzPAdtgtFHo/Fo7PS6U4lieIJf+0eu1NrSBzIniuXzo9xGds8gUfQ/+fHJ",
	"IMJ5y0xHBu/6ceHgibNqts55r/zVeNwGKs4VRCkeMcu/Gqe8HeP3jWLtEZ4CiaLbHhxi58fyRzs73C4S",
	"J9cCbhRQC0UoRZ7Gpq4qopcp3NxoIU0C4d/CsEeQgEUa4rBjjfJ1Dca+k8Xy0ZxcV6wdAFuJSFG4D5fb",
	"fiTE/Zy1uobmmQAbBBiG1TGgfhABPvu555kAh06AJtsourmbg/OV+9/sL8Bxcl+iOD0ny/C75cew2t1W",
	"fLvfmnvdDcp16hjbKVhauotxuOfYsmud8bDNUGc96LYHrotnGqQbwWb8BlxYOcCbEH8ONjEUv/fPe4ri",
	"PWOTCGFXrBW3suCEVHdlQTzpMFgQ8CueEgnSoWuyu6Z8MtLrlD8Jq//JYB9Gyodwb6X8Xzbeyu40eLe3",
	"qSQH4uuV7x2FcCPcNSMsFRzEKB6QvPs8nka+G8rX4H+3uWw/7k9kQDtMguyIeDJf792nvdf7m3Uk0T1q",
	"eAz7gRXy/Qx58u17I6J37OG3sqAHzUOa+b+VCAfQ2jep4LeCnqeD9h7mwKWqQFgUdg1eER69+mU0Ho1H",
	"R5M34zfjWHtwc9H8HQAA//+a/yuA/x4AAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
