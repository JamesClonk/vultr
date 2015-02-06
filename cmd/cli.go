package cmd

import (
	"fmt"

	"github.com/jawher/mow.cli"
)

var apiKey *string

type App struct {
	*cli.Cli
}

func NewApp() *App {
	app := &App{cli.App("vultr", "A Vultr CLI")}

	apiKey = app.String(cli.StringOpt{
		Name: "k api-key",
		//Value:  "xyz1234567890abc",
		Desc:   "Vultr API-Key",
		EnvVar: "VULTR_KEY",
	})

	return app
}

func (app *App) RegisterCommands() {
	// sshkeys
	app.Command("sshkey", "control SSH public keys on Vultr account", func(cmd *cli.Cmd) {
		cmd.Command("create", "upload and add new SSH public key to Vultr account", printApiKey)
		cmd.Command("update", "update an existing SSH public key in your Vultr account", printApiKey)
		cmd.Command("delete", "remove an existing SSH public key from your Vultr account", printApiKey)
		cmd.Command("list", "list all SSH public keys in Vultr account", printApiKey)
	})
}

func printApiKey(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Println(*apiKey)
	}
}
