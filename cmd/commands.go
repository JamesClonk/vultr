package cmd

import (
	"fmt"
	"runtime"

	vultr "github.com/JamesClonk/vultr/lib"
	"github.com/jawher/mow.cli"
)

func (c *CLI) RegisterCommands() {
	// sshkeys
	c.Command("sshkey", "control SSH public keys on your Vultr account", func(cmd *cli.Cmd) {
		cmd.Command("create", "upload and add new SSH public key to your Vultr account", sshKeysCreate)
		cmd.Command("update", "update an existing SSH public key in your Vultr account", sshKeysUpdate)
		cmd.Command("delete", "remove an existing SSH public key from your Vultr account", sshKeysDelete)
		cmd.Command("list", "list all existing SSH public keys from your Vultr account", sshKeysList)
	})

	// version
	c.Command("version", "vultr CLI version", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			fmt.Printf("Client version: %s\n", vultr.Version)
			fmt.Printf("Vultr API version: %s\n", vultr.APIVersion)
			fmt.Printf("Vultr API endpoint: %s\n", vultr.DefaultEndpoint)
			fmt.Printf("OS/Arch (client): %s/%s\n", runtime.GOOS, runtime.GOARCH)
			fmt.Printf("Go version: %s\n", runtime.Version())
		}
	})
}

// for debugging..
func printAPIKey(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Println(*apiKey)
	}
}
