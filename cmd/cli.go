package cmd

import "github.com/jawher/mow.cli"

var apiKey *string

type CLI struct {
	*cli.Cli
}

func NewCLI() *CLI {
	c := &CLI{cli.App("vultr", "A Vultr CLI")}

	apiKey = c.String(cli.StringOpt{
		Name:   "k api-key",
		Desc:   "Vultr API-Key",
		EnvVar: "VULTR_KEY",
	})

	// TODO: read apiKey (and other stuff like default OS image, default region, etc..) from ~/.vultr.json file if present
	// TODO: command line arguments like --api-key take precendence over values from ~/.vultr.json though!
	// TODO: add "vultr auth" command, which prompts for apiKey, then stores it into ~/.vultr.json

	return c
}
