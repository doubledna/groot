package main

import (
	"github.com/hibiken/asynq"
	"groot/internal/config"
	"groot/internal/scheduler"
	"groot/internal/zlog"
	"time"
)

func main() {
	defer zlog.Sync()

	// connect redis
	password := config.GetString("redis.password")
	database := config.GetInt("redis.database")
	addr := config.GetString("redis.address") + ":" + config.GetString("redis.port")
	serverAddress := config.GetString("transport.protocol") + "://" + config.GetString("web.address")

	r := asynq.RedisClientOpt{Addr: addr, Password: password, DB: database}

	// get once task from server
	zlog.Info("start the once-task scheduler")
	client := asynq.NewClient(r)
	go scheduler.OnceTaskManager(client, serverAddress)

	// get periodic task from server
	zlog.Info("start the periodic task scheduler")
	provider := &scheduler.DataBaseBasedConfigProvider{ServerAddress: serverAddress}
	mgr, err := asynq.NewPeriodicTaskManager(
		asynq.PeriodicTaskManagerOpts{
			RedisConnOpt:               r,
			PeriodicTaskConfigProvider: provider,        // this provider object is the interface to your config source
			SyncInterval:               5 * time.Second, // this field specifies how often sync should happen
		})
	if err != nil {
		zlog.Fatalf("periodic scheduler fatal, reason: %s", err)
	}

	if err := mgr.Run(); err != nil {
		zlog.Fatalf("scheduler fatal, reason: %s", err)
	}
}
