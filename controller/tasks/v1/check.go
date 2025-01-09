package v1

import (
	taskrepo "groot/internal/repository/tasks"
	"groot/internal/zlog"
)

func CheckTaskTypeValid(taskType string) bool {
	result, err := taskrepo.ListTaskType()
	if err != nil {
		zlog.Errorf("list task type error", err)
		return false
	}
	for _, types := range result {
		if taskType == types.Name {
			return true
		}
	}
	return false
}
