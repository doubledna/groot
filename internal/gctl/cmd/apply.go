package cmd

import(
	"fmt"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Use a file to create or update a task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("gctl apply -f task.yaml")
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}