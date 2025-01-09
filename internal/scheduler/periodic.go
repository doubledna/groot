package scheduler

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	v1 "groot/gen/v1"
	response2 "groot/internal/response"
	"groot/internal/zlog"
	"strings"
)

// DataBaseBasedConfigProvider : the task provider is database
type DataBaseBasedConfigProvider struct {
	ServerAddress string
}

type Config struct {
	Cronspec string
	TaskType string
	Payload  string
}

type PeriodicTaskConfigContainer struct {
	Configs []Config
}

// GetConfigs implement the PeriodicTaskConfigProvider interface method: GetConfigs
func (d *DataBaseBasedConfigProvider) GetConfigs() ([]*asynq.PeriodicTaskConfig, error) {
	// Get periodic task from the database
	url := d.ServerAddress + "/api/v1/task/mode/periodic"
	p := GetPeriodicTaskFromDataBase(url)

	var configs []*asynq.PeriodicTaskConfig
	for _, cfg := range p.Configs {
		payload := []byte(cfg.Payload)
		configs = append(configs, &asynq.PeriodicTaskConfig{Cronspec: cfg.Cronspec, Task: asynq.NewTask(cfg.TaskType, payload)})
	}
	return configs, nil
}

// GetPeriodicTaskFromDataBase Get periodic task list from the database
func GetPeriodicTaskFromDataBase(url string) *PeriodicTaskConfigContainer {
	body, err := response2.GetRequest(url)
	if err != nil {
		zlog.Errorf("get periodic task from database failed", err)
		return nil
	}

	var responseBody v1.GetTask
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		var errorBody v1.Error
		err = json.Unmarshal(body, &errorBody)
		if err != nil {
			zlog.Errorf("parsing periodic task json failed", err)
			return nil
		}
		return &PeriodicTaskConfigContainer{}
	}

	var p PeriodicTaskConfigContainer
	for _, task := range responseBody.Data {
		cronSpec := determineCronSpecFormat(task.CronSpec)
		taskConfig := Config{cronSpec, task.TaskType, task.Payload}
		p.Configs = append(p.Configs, taskConfig)
	}
	return &p
}

func determineCronSpecFormat(cronSpec string) string {
	if strings.Contains(cronSpec, "*") {
		return cronSpec
	} else {
		return "@every " + cronSpec
	}
}
