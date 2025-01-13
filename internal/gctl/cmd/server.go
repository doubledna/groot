package cmd

import (
	"os"
	"fmt"
    "github.com/spf13/cobra"

)

var rootCmd = &cobra.Command{
	Use:   "gctl",
	Short: "gctl is a command-line tool for groot",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
	  // Do Stuff Here
	  fmt.Println("gctl -h")
	},
}
  
func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Fprintln(os.Stderr, err)
	  os.Exit(1)
	}
}