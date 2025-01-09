package internal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"groot/internal/apps"
	"groot/internal/config"
)

var DB *gorm.DB
var err error

// NewDB : init database
func NewDB() (*gorm.DB, error) {
	user := config.GetString("mysql.username")
	password := config.GetString("mysql.password")
	addr := config.GetString("mysql.address")
	port := config.GetString("mysql.port")
	database := config.GetString("mysql.database")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user,
		password, addr, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "groot_",
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, err
	}

	if db.Error != nil {
		return nil, err
	}
	return db, nil
}

// Init : initialize resources
func Init() error {
	// DB init
	DB, err = NewDB()
	if err != nil {
		return err
	}

	err = apps.MigrateTaskTable(DB)
	if err != nil {
		return err
	}

	err = apps.MigrateTaskTypeTable(DB)
	if err != nil {
		return err
	}
	return nil
}
