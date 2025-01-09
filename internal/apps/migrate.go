package apps

import (
	"fmt"
	"gorm.io/gorm"
	"groot/internal/models/tasks"
)

func MigrateTaskTable(db *gorm.DB) error {
	err := db.AutoMigrate(&tasks.Task{})
	if err != nil {
		fmt.Printf("auto migrate task table failed %s", err)
		return err
	}
	return nil
}

func MigrateTaskTypeTable(db *gorm.DB) error {
	err := db.AutoMigrate(&tasks.TaskType{})
	if err != nil {
		fmt.Printf("auto migrate task type table failed %s", err)
		return err
	}
	return nil
}
