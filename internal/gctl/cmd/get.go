package cmd

import (
	"encoding/json"
	"fmt"
	"groot/internal/models/tasks"
	"groot/internal/zlog"
	"io/ioutil"
	"net/http"
	"os"
	"text/tabwriter"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a task",
	Long:  `Get task list or a task`,
	Run: func(cmd *cobra.Command, args []string) {
	  listTask("x")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

type ListResponse struct {
	Code    int          `json:"code"`
	Data    []tasks.Task `json:"data"`
	Error   string       `json:"error"`
	Message string       `json:"message"`
}

func listTask(serverAddress string) {
	url := serverAddress + "/api/v1/task"
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		zlog.Errorf("error while creating request", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		zlog.Errorf("error while send request", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zlog.Errorf("error while read response", err)
		return
	}

	responseBody := ListResponse{}
	jsonErr := json.Unmarshal(body, &responseBody)
	if jsonErr != nil {
		zlog.Errorf("unmarshal response body failed", jsonErr)
		return
	}

	printError := listTaskPrinter(responseBody.Data)
	if printError != nil {
		zlog.Errorf("error", jsonErr)
		return
	}
	return
}

// formatted output
func listTaskPrinter(s []tasks.Task) error {
	w := tabwriter.NewWriter(os.Stdout, 12, 1, 3, ' ', 0)
	fmt.Fprint(w, "TaskType\tName\tMode\tCronSpec\tResult\tExecutionTime\n")
	for _, item := range s {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%t\t%s\n",
			item.TaskType,
			item.Name,
			item.Mode,
			item.CronSpec,
			item.Result,
			item.UpdatedAt)
	}
	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}
