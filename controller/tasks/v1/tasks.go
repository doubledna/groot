package v1

import (
	"github.com/gin-gonic/gin"
	genv1 "groot/gen/v1"
	"groot/internal/models/tasks"
	taskrepo "groot/internal/repository/tasks"
	"groot/internal/response"
	"net/http"
	"time"
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
	err := c.Bind(&newTask)
	if err != nil {
		taskError(c, response.CodeTaskCreatePostDataFormatError.WithError(err))
		return
	}
	// check whether it is a valid task type
	var task tasks.Task
	taskType := newTask.TaskType
	check := CheckTaskTypeValid(taskType)
	if check {
		// todo: need check some key is it effective
		task.Kind = newTask.Kind
		task.TaskType = taskType
		task.Name = newTask.Name
		task.Mode = newTask.Mode
		task.CronSpec = newTask.CronSpec
		task.Payload = newTask.Payload
		task.Result = newTask.Result
		task.Event = newTask.Event
		task.CreatedAt = getTime(newTask.CreateAt)
		task.UpdatedAt = getTime(newTask.UpdateAt)
		if task.Kind == "task" && task.Name != "" && task.Mode != "" && task.CronSpec != "" &&
			task.Payload != "" {
			result, err := taskrepo.CreateTask(task)
			if err != nil {
				taskError(c, response.CodeTasKCreateFailed.WithError(err))
				return
			}
			getTask(c, result, *response.CodeSuccess)
			return
		} else {
			taskError(c, *response.CodeTaskCreatePostDataIsNull)
			return
		}
	} else {
		taskError(c, *response.CodeTaskCreateTaskTypeInvalid)
		return
	}
}

func (t *TaskStore) UpdateTask(c *gin.Context) {
	var updateTask genv1.NewTask
	err := c.Bind(&updateTask)
	if err != nil {
		taskError(c, response.CodeTaskCreatePostDataFormatError.WithError(err))
		return
	}
	var task tasks.Task
	taskType := updateTask.TaskType
	check := CheckTaskTypeValid(taskType)
	if check == false {
		taskError(c, *response.CodeTaskCreateTaskTypeInvalid)
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
			taskError(c, response.CodeTasKUpdateFailed.WithError(err))
			return
		}
		getTask(c, result, *response.CodeSuccess)
		return
	} else {
		taskError(c, *response.CodeTaskUpdatePutDataIsNull)
		return
	}
}

func (t *TaskStore) GetTask(c *gin.Context) {
	result, err := taskrepo.ListTask()
	if err != nil {
		taskError(c, response.CodeTaskQueryFailed.WithError(err))
		return
	}
	getTask(c, result, *response.CodeSuccess)
}

func (t *TaskStore) GetTaskByName(c *gin.Context, name string) {
	result, err := taskrepo.GetTask(name)
	if err != nil {
		taskError(c, response.CodeTaskQueryFailed.WithError(err))
		return
	}
	if len(result) == 0 {
		taskError(c, *response.CodeTaskQueryDataIsNull)
		return
	}
	getTask(c, result, *response.CodeSuccess)
}

func (t *TaskStore) GetTaskByMode(c *gin.Context, name string) {
	// todo: add check mode is it effective
	result, err := taskrepo.GetTaskByMode(name)
	if err != nil {
		taskError(c, response.CodeTaskQueryFailed.WithError(err))
		return
	}
	if len(result) == 0 {
		taskError(c, *response.CodeTaskQueryDataIsNull)
		return
	}
	getTask(c, result, *response.CodeSuccess)
}

func (t *TaskStore) DeleteTask(c *gin.Context, name string) {
	result, err := taskrepo.DeleteTask(name)
	if err != nil {
		taskError(c, response.CodeTaskDeleteFailed.WithError(err))
		return
	}
	getTask(c, result, *response.CodeSuccess)
}
