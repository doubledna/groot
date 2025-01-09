// Package scheduler : manage once task
package scheduler

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	v1 "groot/gen/v1"
	response2 "groot/internal/response"
	"groot/internal/zlog"
	"strconv"
	"time"
)

// OnceTaskManager manage once task
func OnceTaskManager(client *asynq.Client, address string) {
	url := address + "/api/v1/task/mode/once"
	for {
		p := GetOnceTaskFromDataBase(url)
		if len(*p) == 0 {
			//zlog.Info("no once tasks")
			time.Sleep(5 * time.Second)
			continue
		}
		for _, cfg := range *p {
			// filter completed tasks
			if cfg.Result == false {
				taskType := cfg.TaskType
				cronSpec, err := strconv.ParseInt(cfg.CronSpec, 10, 64) // set the waiting time before executing the task
				if err != nil {
					zlog.Errorf("can't convert cronSpec string to int64: %s", err)
					continue
				}
				payload := cfg.Payload
				err = NewOnceTaskManager(taskType, payload, cronSpec, client)
				if err != nil {
					zlog.Errorf("create task failed", err)
					continue
				}

				// set task result to true and update updateAt
				cfg.Result = true
				cfg.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
				err = UpdateTask(cfg, address)
				if err != nil {
					zlog.Errorf("update task failed", err)
					continue
				}
			}
		}
		time.Sleep(5 * time.Second)
	}
}

// NewOnceTaskManager create a new once-task
func NewOnceTaskManager(taskType string, payload string, waitTime int64, client *asynq.Client) error {
	payloads := []byte(payload)
	task := asynq.NewTask(taskType, payloads)
	result, err := client.Enqueue(task, asynq.ProcessIn(time.Duration(waitTime)*time.Second))
	if err != nil {
		zlog.Errorf("failed enqueued task", err)
		return err
	}
	zlog.Info("successfully enqueued once task", zap.Any("task detail", result))
	return nil
}

// GetOnceTaskFromDataBase get once task list from database
func GetOnceTaskFromDataBase(url string) *[]v1.Task {
	body, err := response2.GetRequest(url)
	if err != nil {
		zlog.Errorf("get once task from database failed", err)
		return nil
	}

	var responseBody v1.GetTask
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		//zlog.Errorf("parsing task json failed", err)
		var errorBody v1.Error
		err = json.Unmarshal(body, &errorBody)
		if err != nil {
			zlog.Errorf("parsing once task json failed", err)
			return nil
		}
		return &[]v1.Task{}
	}
	return &responseBody.Data
}

func UpdateTask(task v1.Task, address string) error {
	url := address + "/api/v1/task"
	taskBytes, err := json.Marshal(task)
	if err != nil {
		zlog.Errorf("update task conversion failed", err)
		return err
	}

	err = response2.PutRequest(url, taskBytes)
	if err != nil {
		zlog.Errorf("update task request failed", err)
		return err
	}
	return nil
}
