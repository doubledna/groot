package internal

import (
	"groot/internal/zlog"
	"net/http"
)

func DeleteTask(serverAddress string, taskName string) {
	url := serverAddress + "/api/v1/task/" + taskName

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		zlog.Errorf("create request failed", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		zlog.Errorf("send request failed", err)
		return
	}

	defer resp.Body.Close()

	zlog.Info("delete task successfully")
}
