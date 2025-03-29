package v1

import (
	genv1 "groot/gen/v1"
	"groot/internal/models/tasks"
	taskrepo "groot/internal/repository/tasks"
	"groot/internal/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getTime(timeStr string) time.Time {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Now()
	}
	return parsedTime
}

func getTask(c *gin.Context, task []tasks.Task, response response.Response) {
	result := make([]genv1.Task, 0, 100)
	for _, value := range task {
		result = append(result, genv1.Task{
			Id:       int64(value.ID),
			Kind:     value.Kind,
			TaskType: value.TaskType,
			Name:     value.Name,
			Mode:     value.Mode,
			CronSpec: value.CronSpec,
			Payload:  value.Payload,
			Result:   value.Result,
			Event:    value.Event,
			CreateAt: value.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt: value.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	resp := genv1.GetTask{
		Code:      int64(response.Code),
		Message:   response.Message,
		Reference: response.Reference,
		Error:     response.Error,
		Data:      result,
	}
	c.JSON(http.StatusOK, resp)
}

func (t *TaskStore) CreateTask(c *gin.Context) {
	var newTask genv1.NewTask
	if err := c.Bind(&newTask); err != nil {
		ErrorFormat(c, response.CodeTaskCreatePostDataFormatError.WithError(err))
		return
	}

	taskType := newTask.TaskType
	if !CheckTaskTypeValid(taskType) {
		ErrorFormat(c, *response.CodeTaskCreateTaskTypeInvalid)
		return
	}

	task := tasks.Task{
		Kind:      newTask.Kind,
		TaskType:  taskType,
		Name:      newTask.Name,
		Mode:      newTask.Mode,
		CronSpec:  newTask.CronSpec,
		Payload:   newTask.Payload,
		Result:    newTask.Result,
		Event:     newTask.Event,
		CreatedAt: getTime(newTask.CreateAt),
		UpdatedAt: getTime(newTask.UpdateAt),
	}

	if task.Kind != "task" || task.Name == "" || task.Mode == "" || 
	   task.CronSpec == "" || task.Payload == "" {
		ErrorFormat(c, *response.CodeTaskCreatePostDataIsNull)
		return
	}

	result, err := taskrepo.CreateTask(task)
	if err != nil {
		ErrorFormat(c, response.CodeTasKCreateFailed.WithError(err))
		return
	}

	getTask(c, result, *response.CodeSuccess)
}

func (t *TaskStore) UpdateTask(c *gin.Context) {
	var updateTask genv1.NewTask
	err := c.Bind(&updateTask)
	if err != nil {
		ErrorFormat(c, response.CodeTaskCreatePostDataFormatError.WithError(err))
		return
	}
	var task tasks.Task
	taskType := updateTask.TaskType
	check := CheckTaskTypeValid(taskType)
	if !check {
		ErrorFormat(c, *response.CodeTaskCreateTaskTypeInvalid)
		return
	}
	task.Kind = updateTask.Kind
	task.TaskType = taskType
	task.Name = updateTask.Name
	task.Mode = updateTask.Mode
	task.CronSpec = updateTask.CronSpec
	task.Payload = updateTask.Payload
	task.Result = updateTask.Result
	task.Event = updateTask.Event
	task.CreatedAt = getTime(updateTask.CreateAt)
	task.UpdatedAt = getTime(updateTask.UpdateAt)
	if task.Kind == "task" && task.Name != "" && task.Mode != "" && task.CronSpec != "" &&
		task.Payload != "" {
		result, err := taskrepo.UpdateTask(task)
		if err != nil {
			ErrorFormat(c, response.CodeTasKUpdateFailed.WithError(err))
			return
		}
		getTask(c, result, *response.CodeSuccess)
		return
	} else {
		ErrorFormat(c, *response.CodeTaskUpdatePutDataIsNull)
		return
	}
}

func (t *TaskStore) GetTask(c *gin.Context) {
	result, err := taskrepo.ListTask()
	if err != nil {
		ErrorFormat(c, response.CodeTaskQueryFailed.WithError(err))
		return
	}
	getTask(c, result, *response.CodeSuccess)
}

func (t *TaskStore) GetTaskByName(c *gin.Context, name string) {
	result, err := taskrepo.GetTask(name)
	if err != nil {
		ErrorFormat(c, response.CodeTaskQueryFailed.WithError(err))
		return
	}
	if len(result) == 0 {
		ErrorFormat(c, *response.CodeTaskQueryDataIsNull)
		return
	}
	getTask(c, result, *response.CodeSuccess)
}

func (t *TaskStore) GetTaskByMode(c *gin.Context, name string) {
	// todo: add check mode is it effective
	result, err := taskrepo.GetTaskByMode(name)
	if err != nil {
		ErrorFormat(c, response.CodeTaskQueryFailed.WithError(err))
		return
	}
	if len(result) == 0 {
		ErrorFormat(c, *response.CodeTaskQueryDataIsNull)
		return
	}
	getTask(c, result, *response.CodeSuccess)
}

func (t *TaskStore) DeleteTask(c *gin.Context, name string) {
	result, err := taskrepo.DeleteTask(name)
	if err != nil {
		ErrorFormat(c, response.CodeTaskDeleteFailed.WithError(err))
		return
	}
	getTask(c, result, *response.CodeSuccess)
}
