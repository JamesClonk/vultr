package cmd

import (
	"fmt"
	"log"

	vultr "github.com/JamesClonk/vultr/lib"
	"github.com/jawher/mow.cli"
)

func sshKeysCreate(cmd *cli.Cmd) {
	cmd.Spec = "-n -k"

	name := cmd.StringOpt("n name", "", "Name of the SSH key")
	sshkey := cmd.StringOpt("k key", "", "SSH public key (in authorized_keys format)")

	cmd.Action = func() {
		key, err := GetClient().CreateSSHKey(*name, *sshkey)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("SSH key created\n")
		lengths := []int{24, 32, 64}
		tabsPrint(Columns{"SSHKEYID", "NAME", "KEY"}, lengths)
		tabsPrint(Columns{key.ID, key.Name, key.Key}, lengths)
		tabsFlush()
	}
}

func sshKeysUpdate(cmd *cli.Cmd) {
	cmd.Spec = "SSHKEYID [-n] [-k]"

	id := cmd.StringArg("SSHKEYID", "", "SSHKEYID of key to update (see <sshkeys>)")
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
		fmt.Println("SSH key updated")
	}
}

func sshKeysDelete(cmd *cli.Cmd) {
	id := cmd.StringArg("SSHKEYID", "", "SSHKEYID of key to delete (see <sshkeys>)")
	cmd.Action = func() {
		if err := GetClient().DeleteSSHKey(*id); err != nil {
			log.Fatal(err)
		}
		fmt.Println("SSH key deleted")
	}
}

func sshKeysList(cmd *cli.Cmd) {
	cmd.Spec = "[-f]"

	full := cmd.BoolOpt("f full", false, "Display full length of SSH key")

	cmd.Action = func() {
		keys, err := GetClient().GetSSHKeys()
		if err != nil {
			log.Fatal(err)
		}

		if len(keys) == 0 {
			fmt.Println()
			return
		}

		keyLength := 64
		if *full {
			keyLength = 8192
		}
		lengths := []int{24, 32, keyLength}

		tabsPrint(Columns{"SSHKEYID", "NAME", "KEY"}, lengths)
		for _, key := range keys {
			tabsPrint(Columns{key.ID, key.Name, key.Key}, lengths)
		}
		tabsFlush()
	}
}
