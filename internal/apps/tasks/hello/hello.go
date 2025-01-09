package hello

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"groot/internal/zlog"
)

const (
	TypeIamGroot = "hello"
)

// IamGrootTaskPayload : A list of IamGroot task payload
type IamGrootTaskPayload struct {
	UserID int
	Name   string
}

// HandleIamGrootTask specific execution method
func HandleIamGrootTask(ctx context.Context, t *asynq.Task) error {
	var htp IamGrootTaskPayload
	if err := json.Unmarshal(t.Payload(), &htp); err != nil {
		return err
	}
	zlog.Info("[*]", zap.Int("ID", htp.UserID), zap.String("Message", htp.Name))
	return nil
}

// NewIamGrootTask create a task using task type and payload
//func NewIamGrootTask(id int, name string) (*asynq.Task, error) {
//	payload, err := json.Marshal(IamGrootTaskPayload{UserID: id, Name: name})
//	if err != nil {
//		return nil, err
//	}
//	return asynq.NewTask(TypeIamGroot, payload), nil
//}

// CronIamGrootTask encapsulate tasks
//func CronIamGrootTask(id int, name string) *asynq.Task {
//	taskIamGroot, err := NewIamGrootTask(id, name)
//	if err != nil {
//		log.Println(err)
//	}
//	return taskIamGroot
//}
