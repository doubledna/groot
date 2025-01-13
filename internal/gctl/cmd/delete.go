package cmd

import (
	"groot/internal/zlog"
	"net/http"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  `Delete a task from the command line or from a file`,
	Run: func(cmd *cobra.Command, args []string) {
	  deleteTask("x", "y")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteTask(serverAddress string, taskName string) {
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
