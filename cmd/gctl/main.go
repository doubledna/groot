package main

import (
	"github.com/urfave/cli/v2"
	"groot/internal/gctl"
	"groot/internal/zlog"
	"os"
)

const (
	usage = `gctl is a command-line tool`
)

func main() {
	defer zlog.Sync()

	app := cli.NewApp()
	app.Name = "gctl"
	app.Usage = usage

	app.Commands = []*cli.Command{
		gctl.ApplyCommand,
		gctl.GetCommand,
		gctl.DeleteCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		zlog.Fatalf("run exit, err: ", err)
	}
}
