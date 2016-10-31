package cliapp

import (
	"strconv"

	"github.com/bpicode/fritzctl/fatals"
	"github.com/bpicode/fritzctl/fritz"
	"github.com/bpicode/fritzctl/logger"
	"github.com/mitchellh/cli"
)

type temperatureCommand struct {
}

func (cmd *temperatureCommand) Help() string {
	return "Set the temperature of a HKR device. Example usage: fritzctl temperature 21.0 MY_HKR_DEVICE"
}

func (cmd *temperatureCommand) Synopsis() string {
	return "Set the temperature of a HKR device."
}

func (cmd *temperatureCommand) Run(args []string) int {
	temp, errorParse := strconv.ParseFloat(args[0], 64)
	fatals.AssertNoError(errorParse, "Cannot parse temperature value:", errorParse)
	f := fritz.UsingClient(clientLogin())
	res, err := f.Temperature(args[1], temp)
	fatals.AssertNoError(err, "Unable to set temperature:", err)
	logger.Info("Success! FRITZ!Box answered: " + res)
	return 0
}

func temperature() (cli.Command, error) {
	p := temperatureCommand{}
	return &p, nil
}
