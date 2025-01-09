package tasks

import (
	"groot/internal"
	"groot/internal/models/tasks"
)

// CreateTask create task method
func CreateTask(task tasks.Task) ([]tasks.Task, error) {
	var taskLists []tasks.Task
	result := internal.DB.Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	taskLists = append(taskLists, task)
	return taskLists, nil
}

// ListTask get all task
func ListTask() ([]tasks.Task, error) {
	var taskLists []tasks.Task
	result := internal.DB.Find(&taskLists)
	if result.Error != nil {
		return nil, result.Error
	}
	return taskLists, nil
}

// GetTask query a task by name
func GetTask(name string) ([]tasks.Task, error) {
	var taskLists []tasks.Task
	result := internal.DB.Where("name = ?", name).Find(&taskLists)
	if result.Error != nil {
		return nil, result.Error
	}
	return taskLists, nil
}

// GetTaskByMode query tasks by mode
func GetTaskByMode(mode string) ([]tasks.Task, error) {
	var taskLists []tasks.Task
	result := internal.DB.Where("mode = ?", mode).Find(&taskLists)
	if result.Error != nil {
		return nil, result.Error
	}
	return taskLists, nil
}

// UpdateTask update a task by name
func UpdateTask(task tasks.Task) ([]tasks.Task, error) {
	name := task.Name
	var taskLists []tasks.Task
	var taskUpdate tasks.Task
	queryResult := internal.DB.Where("name = ?", name).First(&taskUpdate)
	if queryResult.Error != nil {
		return nil, queryResult.Error
	}
	result := internal.DB.Model(&taskUpdate).Updates(task)
	if result.Error != nil {
		return nil, result.Error
	}
	taskLists = append(taskLists, task)
	return taskLists, nil
}

// DeleteTask delete a task
func DeleteTask(name string) ([]tasks.Task, error) {
	var task tasks.Task
	var taskLists []tasks.Task
	result := internal.DB.Where("name = ?", name).Delete(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	taskLists = append(taskLists, task)
	return taskLists, nil
}
