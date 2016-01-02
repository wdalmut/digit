package main

import (
	"os"

	"github.com/wdalmut/digit/command"
	"gopkg.in/wdalmut/cli.v1"
)

func main() {
	c := cli.NewCli("digit", "0.0.1")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"say":     command.SayCommandFactory,
		"convert": command.ConvertCommandFactory,
	}

	exitStatus := c.Run()
	os.Exit(exitStatus)
}
