package cmd

import (
	"fmt"
	"runtime"

	vultr "github.com/JamesClonk/vultr/lib"
	"github.com/jawher/mow.cli"
)

func (c *CLI) RegisterCommands() {
	// info
	c.Command("info", "display account information", accountInfo)

	// os
	c.Command("os", "list all available operating systems", osList)

	// iso
	c.Command("iso", "list all ISOs currently available on account", isoList)

	// plans
	c.Command("plans", "list all active plans", planList)

	// regions
	c.Command("regions", "list all active regions", regionList)

	// sshkeys
	c.Command("sshkey", "modify SSH public keys", func(cmd *cli.Cmd) {
		cmd.Command("create", "upload and add new SSH public key", sshKeysCreate)
		cmd.Command("update", "update an existing SSH public key", sshKeysUpdate)
		cmd.Command("delete", "remove an existing SSH public key", sshKeysDelete)
		cmd.Command("list", "list all existing SSH public keys", sshKeysList)
	})
	c.Command("sshkeys", "list all existing SSH public keys", sshKeysList)

	// ssh
	c.Command("ssh", "ssh into a virtual machine", printAPIKey)

	// servers
	c.Command("server", "modify virtual machines", func(cmd *cli.Cmd) {
		cmd.Command("create", "create a new virtual machine", printAPIKey)
		cmd.Command("start", "start a virtual machine (restart if already running)", printAPIKey)
		cmd.Command("halt", "halt a virtual machine (hard power off)", printAPIKey)
		cmd.Command("reboot", "reboot a virtual machine (hard reboot)", printAPIKey)
		cmd.Command("reinstall", "reinstall OS on a virtual machine (all data will be lost)", printAPIKey)
		cmd.Command("change-os", "change OS on a virtual machine (all data will be lost)", printAPIKey)
		cmd.Command("delete", "delete a virtual machine", printAPIKey)
		cmd.Command("bandwidth", "list bandwidth used by a virtual machine", printAPIKey)
		cmd.Command("list", "list all active or pending virtual machines on the current account", printAPIKey)
	})
	c.Command("servers", "list all active or pending virtual machines on the current account", printAPIKey)

	// snapshots
	c.Command("snapshot", "modify snapshots", func(cmd *cli.Cmd) {
		cmd.Command("create", "create a snapshot from an existing virtual machine", printAPIKey)
		cmd.Command("delete", "delete a snapshot", printAPIKey)
		cmd.Command("list", "list all snapshots on the current account", printAPIKey)
	})
	c.Command("snapshots", "list all snapshots on the current account", printAPIKey)

	// startup scripts
	c.Command("script", "modify startup scripts", func(cmd *cli.Cmd) {
		cmd.Command("create", "create a new startup script", printAPIKey)
		cmd.Command("update", "update an existing startup script", printAPIKey)
		cmd.Command("delete", "remove an existing startup script", printAPIKey)
		cmd.Command("list", "list all startup scripts on the current account", printAPIKey)
	})
	c.Command("scripts", "list all startup scripts on the current account", printAPIKey)

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
