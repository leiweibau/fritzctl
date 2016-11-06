package cliapp

import (
	"os"

	"github.com/bpicode/fritzctl/meta"
	"github.com/mitchellh/cli"
)

// Create creates a new CLI application.
func Create() *cli.CLI {
	c := cli.NewCLI(meta.ApplicationName, meta.Version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"configure":   configure,
		"list":        list,
		"ping":        ping,
		"switch":      delegating(pairOf("on", switchOnDevice), pairOf("off", switchOffDevice)),
		"toggle":      toggleDevice,
		"temperature": temperature,
	}
	return c
}
