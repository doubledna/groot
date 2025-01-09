package gctl

import (
	"github.com/urfave/cli/v2"
	"groot/internal/config"
	"groot/internal/gctl/internal"
	"log"
)

var (
	file          string
	task          string
	serverAddress = config.GetString("transport.protocol") + "://" + config.GetString("web.address")
)

var ApplyCommand = &cli.Command{
	Name:      "apply",
	Usage:     "create or update tasks using yaml files",
	ArgsUsage: ``,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "file",
			Aliases:     []string{"f"},
			Value:       "",
			Usage:       `gctl apply --file task.yaml`,
			Destination: &file,
		},
	},
	Action: func(*cli.Context) error {
		Apply(file)
		return nil
	},
}

// Apply : create or update task
func Apply(file string) {
	log.Println(file)
}

var GetCommand = &cli.Command{
	Name:  "get",
	Usage: "get task details",
	//Flags: []cli.Flag{
	//	&cli.StringFlag{
	//		Name:        "task",
	//		Aliases:     []string{"t"},
	//		Value:       "",
	//		Usage:       `gctl get --task {task name}`,
	//		Destination: &task,
	//	},
	//},
	//Action: func(*cli.Context) error {
	//	Gets(task)
	//	return nil
	//},
	Subcommands: []*cli.Command{
		{
			Name:  "tasks",
			Usage: "gctl get tasks",
			Action: func(cCtx *cli.Context) error {
				Lists()
				return nil
			},
		},
	},
}

// Lists : list all tasks
func Lists() {
	internal.ListTask(serverAddress)
	return
}

var DeleteCommand = &cli.Command{
	Name:      "delete",
	Usage:     "delete tasks",
	ArgsUsage: ``,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "task",
			Aliases:     []string{"t"},
			Value:       "",
			Usage:       `gctl delete --task {task name}`,
			Destination: &task,
		},
	},
	Action: func(*cli.Context) error {
		Delete(task)
		return nil
	},
}

// Delete : delete task
func Delete(taskName string) {
	internal.DeleteTask(serverAddress, taskName)
	return
}
