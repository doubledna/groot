package v1

import (
	"github.com/gin-gonic/gin"
	genv1 "groot/gen/v1"
	"groot/internal/models/tasks"
	taskrepo "groot/internal/repository/tasks"
	"groot/internal/response"
	"net/http"
)

type TaskStore struct {
}

func NewTaskStore() *TaskStore {
	return &TaskStore{}
}

// Error wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func ErrorFormat(c *gin.Context, response response.Response) {
	taskErr := genv1.Error{
		Code:      int64(response.Code),
		Message:   response.Message,
		Reference: response.Reference,
		Error:     response.Error,
		Data:      int64(-1),
	}
	c.JSON(http.StatusOK, taskErr)
}

// getTaskType structured response format
func getTaskType(c *gin.Context, task []tasks.TaskType, response response.Response) {
	result := make([]genv1.TaskType, 0, 100)
	for _, value := range task {
		result = append(result, genv1.TaskType{
			Id:   int64(value.ID),
			Kind: value.Kind,
			Name: value.Name,
		})
	}
	resp := genv1.GetTaskType{
		Code:      int64(response.Code),
		Message:   response.Message,
		Reference: response.Reference,
		Error:     response.Error,
		Data:      result,
	}
	c.JSON(http.StatusOK, resp)
}

func (t *TaskStore) CreateTaskType(c *gin.Context) {
	var newTaskType genv1.NewTaskType
	err := c.Bind(&newTaskType)
	if err != nil {
		ErrorFormat(c, response.CodeTaskTypeCreatePostDataFormatError.WithError(err))
		return
	}
	var taskType tasks.TaskType
	taskType.Kind = newTaskType.Kind
	taskType.Name = newTaskType.Name
	if taskType.Kind == "task_type" && taskType.Name != "" {
		result, err := taskrepo.CreateTaskType(taskType)
		if err != nil {
			ErrorFormat(c, response.CodeTaskTypeCreateFailed.WithError(err))
			return
		}
		getTaskType(c, result, *response.CodeSuccess)
		return
	} else {
		ErrorFormat(c, *response.CodeTaskTypeCreatePostDataIsNull)
		return
	}
}

func (t *TaskStore) GetTaskType(c *gin.Context) {
	result, err := taskrepo.ListTaskType()
	if err != nil {
		ErrorFormat(c, response.CodeTaskTypeQueryFailed.WithError(err))
		return
	}
	if len(result) == 0 {
		ErrorFormat(c, *response.CodeTaskTypeQueryDataIsNull)
		return
	}
	getTaskType(c, result, *response.CodeSuccess)
}

func (t *TaskStore) GetTaskTypeByName(c *gin.Context, name string) {
	result, err := taskrepo.GetTaskType(name)
	if err != nil {
		ErrorFormat(c, response.CodeTaskTypeQueryFailed.WithError(err))
		return
	}
	if len(result) == 0 {
		ErrorFormat(c, response.CodeTaskTypeQueryDataIsNull.WithError(err))
		return
	}
	getTaskType(c, result, *response.CodeSuccess)
}

func (t *TaskStore) DeleteTaskType(c *gin.Context, name string) {
	result, err := taskrepo.DeleteTaskType(name)
	if err != nil {
		ErrorFormat(c, response.CodeTaskTypeDeleteFailed.WithError(err))
		return
	}
	getTaskType(c, result, *response.CodeSuccess)
}
