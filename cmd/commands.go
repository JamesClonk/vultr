package cmd

import (
	"fmt"
	"log"
	"runtime"

	vultr "github.com/JamesClonk/vultr/lib"
	"github.com/jawher/mow.cli"
)

func (c *CLI) RegisterCommands() {
	// sshkeys
	c.Command("sshkey", "control SSH public keys on Vultr account", func(cmd *cli.Cmd) {
		cmd.Command("create", "upload and add new SSH public key to your Vultr account", printAPIKey)
		cmd.Command("update", "update an existing SSH public key in your Vultr account", sshKeysUpdate)
		cmd.Command("delete", "remove an existing SSH public key from your Vultr account", printAPIKey)
		cmd.Command("list", "list all SSH public keys in Vultr account", sshKeysList)
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

func sshKeysUpdate(cmd *cli.Cmd) {
	cmd.Spec = "SSHKEYID [-n | --name] [-k | --key]"

	id := cmd.StringArg("SSHKEYID", "", "SSHKEYID of key to update (see <sshkey list>)")
	name := cmd.StringOpt("n name", "", "New name for the SSH key")
	key := cmd.StringOpt("k key", "", "New SSH key contents")

	cmd.Action = func() {
		sshkey := vultr.SSHKey{
			ID:   *id,
			Name: *name,
			Key:  *key,
		}
		if err := GetClient().UpdateSSHKey(sshkey); err != nil {
			log.Fatal(err)
		}
		fmt.Println("SSH key update success!")
	}
}

func sshKeysList(cmd *cli.Cmd) {
	cmd.Spec = "[-f | --full]"

	full := cmd.BoolOpt("f full", false, "Display full length of SSH key")

	cmd.Action = func() {
		keys, err := GetClient().GetSSHKeys()
		if err != nil {
			log.Fatal(err)
		}

		keyLength := 64
		if *full {
			keyLength = 8192
		}
		lengths := []int{24, 32, keyLength}

		printTabbedLine([]string{"SSHKEYID", "NAME", "KEY"}, lengths)
		for _, key := range keys {
			printTabbedLine([]string{key.ID, key.Name, key.Key}, lengths)
		}
		tabsFlush()
	}
}
