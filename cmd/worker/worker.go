package main

import (
	"github.com/hibiken/asynq"
	taskAnsible "groot/internal/apps/tasks/ansible"
	taskIamGroot "groot/internal/apps/tasks/hello"
	"groot/internal/config"
	"groot/internal/zlog"
)

func main() {
	defer zlog.Sync()
	zlog.Info("start worker...")

	// connect redis
	password := config.GetString("redis.password")
	database := config.GetInt("redis.database")
	addr := config.GetString("redis.address") + ":" + config.GetString("redis.port")

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: addr, Password: password, DB: database},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"default": 3,
			},
			StrictPriority: true,
		},
	)

	mux := asynq.NewServeMux()
	// handle task (registration task)
	mux.HandleFunc(taskIamGroot.TypeIamGroot, taskIamGroot.HandleIamGrootTask)
	mux.HandleFunc(taskAnsible.TypeAnsible, taskAnsible.HandleAnsibleTask)

	if err := srv.Run(mux); err != nil {
		zlog.Fatalf("worker fatal, ", err)
	}
}
