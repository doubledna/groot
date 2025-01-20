// Package response
/*
Define request/response
The Response code includes the error code, which is designed in a separate package.
*/
package response

import (
	"fmt"
	"reflect"
)

type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Reference string      `json:"reference"`
	Error     string      `json:"error"`
	Data      interface{} `json:"data"`
}

func NewResponse(code int, message, reference string) *Response {
	return &Response{
		Code:      code,
		Message:   message,
		Reference: reference,
		Error:     "",
		Data:      nil,
	}
}

func (res *Response) WithError(err error) Response {
	return Response{
		Code:      res.Code,
		Message:   res.Message,
		Reference: res.Reference,
		Error:     fmt.Sprintf("%s", err),
		Data:      nil,
	}
}

func HealthCheckResponse(data interface{}) Response {
	return Response{
		Code:      CodeSuccess.Code,
		Message:   CodeSuccess.Message,
		Reference: CodeSuccess.Reference,
		Error:     "",
		Data:      ensureResultNotNull(data),
	}
}

// ensureResultNotNull
// Description: check return value, convert empty slices and maps with unallocated memory, prevent being serialized to
// null, cases front-end error, In other cases, the original result will be returned directly.
// param result
// return interface{}
func ensureResultNotNull(result interface{}) interface{} {
	v := reflect.Indirect(reflect.ValueOf(result))

	switch v.Type().Kind() {
	case reflect.Slice:
		if v.IsNil() {
			return []interface{}{}
		}
	case reflect.Map:
		if v.IsNil() {
			return map[string]interface{}{}
		}
	default:
		return result
	}

	return result
}

// 响应码说明 (Response code description)
// 1-3 位: HTTP 状态码 （1-3: HTTP status code）
// 4-5 位: 组件编号 (4-5: Components Number)
// 6-8 位: 组件内部错误码 (6-8: Component internal error code)
var (
	// Component Number: 01, system code
	CodeSuccess      = NewResponse(20001000, "success", "")
	CodeBadRequest   = NewResponse(40001000, "bad request", "")
	CodeUnauthorized = NewResponse(40101000, "unauthorized", "")
	CodeForbidden    = NewResponse(40301000, "forbidden", "")
	CodeNotFound     = NewResponse(40401000, "not found resource", "")
	CodeUnknownError = NewResponse(50001000, "unknown error", "")

	// task_type code
	CodeTaskTypeCreatePostDataFormatError = NewResponse(40001001, "add task type failed，the request parameter format error", "")
	CodeTaskTypeCreateFailed              = NewResponse(20001002, "add task type failed", "")
	CodeTaskTypeCreatePostDataIsNull      = NewResponse(20001003, "add task type failed，the request parameter is incorrect or empty", "")
	CodeTaskTypeQueryFailed               = NewResponse(20001004, "query task type list failed", "")
	CodeTaskTypeQueryDataIsNull           = NewResponse(20001005, "query task type result is empty", "")
	CodeTaskTypeDeleteFailed              = NewResponse(20001006, "delete task type failed", "")

	// task code
	CodeTaskCreatePostDataFormatError = NewResponse(40001007, "add task failed，the request parameter format error", "")
	CodeTasKCreateFailed              = NewResponse(20001008, "add task failed", "")
	CodeTaskCreatePostDataIsNull      = NewResponse(20001009, "add task failed，the request parameter is incorrect or empty", "")
	CodeTaskCreateTaskTypeInvalid     = NewResponse(20001010, "add task failed，the request parameter: task_type is invalid", "")
	CodeTaskQueryFailed               = NewResponse(20001011, "query task list failed", "")
	CodeTaskQueryDataIsNull           = NewResponse(20001012, "query task result is empty", "")
	CodeTaskQueryByModeFailed         = NewResponse(20001013, "query task list failed by mode", "")
	CodeTasKUpdateFailed              = NewResponse(20001014, "update task failed", "")
	CodeTaskUpdatePutDataIsNull       = NewResponse(20001015, "update task failed，the request parameter is incorrect or empty", "")
	CodeTaskDeleteFailed              = NewResponse(20001016, "delete task failed", "")

	// user code
	CodeUserCreatePostDataFormatError = NewResponse(40001017, "add user failed，the request parameter format error", "")
	CodeUserCreateFailed              = NewResponse(20001018, "add user failed", "")
	CodeUserCreatePostDataIsNull      = NewResponse(20001019, "add user failed，the request parameter is incorrect or empty", "")

	// login code
	CodeLoginFailed = NewResponse(20001020, "login failed", "")
)
