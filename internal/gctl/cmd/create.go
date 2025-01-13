package cmd

import(
	"fmt"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a task",
	Long:  `Create a task from the command line or from a file`,
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("gctl create -f task.yaml or gctl create task --flag")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
