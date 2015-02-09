package cmd

import (
	"fmt"
	"log"
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
	c.Command("ssh", "ssh into a virtual machine", unimplemented)

	// servers
	c.Command("server", "modify virtual machines", func(cmd *cli.Cmd) {
		cmd.Command("create", "create a new virtual machine", serversCreate)
		cmd.Command("rename", "rename a virtual machine", serversRename)
		cmd.Command("start", "start a virtual machine (restart if already running)", serversStart)
		cmd.Command("halt", "halt a virtual machine (hard power off)", serversHalt)
		cmd.Command("reboot", "reboot a virtual machine (hard reboot)", serversReboot)
		cmd.Command("reinstall", "reinstall OS on a virtual machine (all data will be lost)", unimplemented)
		cmd.Command("change-os", "change OS on a virtual machine (all data will be lost)", unimplemented)
		cmd.Command("delete", "delete a virtual machine", serversDelete)
		cmd.Command("bandwidth", "list bandwidth used by a virtual machine", unimplemented)
		cmd.Command("list", "list all active or pending virtual machines on current account", serversList)
		cmd.Command("show", "list detailed information of a virtual machine", serversShow)
	})
	c.Command("servers", "list all active or pending virtual machines on current account", serversList)

	// snapshots
	c.Command("snapshot", "modify snapshots", func(cmd *cli.Cmd) {
		cmd.Command("create", "create a snapshot from an existing virtual machine", unimplemented)
		cmd.Command("delete", "delete a snapshot", unimplemented)
		cmd.Command("list", "list all snapshots on current account", unimplemented)
	})
	c.Command("snapshots", "list all snapshots on current account", unimplemented)

	// startup scripts
	c.Command("script", "modify startup scripts", func(cmd *cli.Cmd) {
		cmd.Command("create", "create a new startup script", unimplemented)
		cmd.Command("update", "update an existing startup script", unimplemented)
		cmd.Command("delete", "remove an existing startup script", unimplemented)
		cmd.Command("list", "list all startup scripts on current account", unimplemented)
	})
	c.Command("scripts", "list all startup scripts on current account", unimplemented)

	// version
	c.Command("version", "vultr CLI version", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			lengths := []int{24, 48}
			printTabbedLine(Columns{"Client version:", vultr.Version}, lengths)
			printTabbedLine(Columns{"Vultr API endpoint:", vultr.DefaultEndpoint}, lengths)
			printTabbedLine(Columns{"Vultr API version:", vultr.APIVersion}, lengths)
			printTabbedLine(Columns{"OS/Arch (client):", fmt.Sprintf("%v/%v", runtime.GOOS, runtime.GOARCH)}, lengths)
			printTabbedLine(Columns{"Go version:", runtime.Version()}, lengths)
			tabsFlush()
		}
	})
}

func unimplemented(cmd *cli.Cmd) {
	cmd.Action = func() {
		log.Fatal("Command not yet implemented!")
	}
}
