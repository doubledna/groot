package tasks

import (
	"groot/internal"
	"groot/internal/models/tasks"
)

// CreateTaskType create task type
func CreateTaskType(tt tasks.TaskType) ([]tasks.TaskType, error) {
	var taskTypes []tasks.TaskType
	result := internal.DB.Create(&tt)
	if result.Error != nil {
		return nil, result.Error
	}
	taskTypes = append(taskTypes, tt)
	return taskTypes, nil
}

// ListTaskType list all task type
func ListTaskType() ([]tasks.TaskType, error) {
	var taskTypes []tasks.TaskType
	result := internal.DB.Find(&taskTypes)
	if result.Error != nil {
		return nil, result.Error
	}
	return taskTypes, nil
}

// GetTaskType query a task type by name
func GetTaskType(name string) ([]tasks.TaskType, error) {
	var taskTypes []tasks.TaskType
	result := internal.DB.Where("name = ?", name).Find(&taskTypes)
	if result.Error != nil {
		return nil, result.Error
	}
	return taskTypes, nil
}

// DeleteTaskType delete task type by name
func DeleteTaskType(name string) ([]tasks.TaskType, error) {
	var taskTypes []tasks.TaskType
	var taskType tasks.TaskType
	result := internal.DB.Where("name = ?", name).Delete(&taskType)
	if result.Error != nil {
		return nil, result.Error
	}
	taskTypes = append(taskTypes, taskType)
	return taskTypes, nil
}
