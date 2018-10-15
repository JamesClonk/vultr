package cmd

import (
	"fmt"
	"log"

	"github.com/jawher/mow.cli"
)

func backupsList(cmd *cli.Cmd) {
	cmd.Spec = "[SUBID] [BACKUPID]"

	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	backupid := cmd.StringArg("BACKUPID", "", "BACKUPID of a virtual machine")

	cmd.Action = func() {
		backups, err := GetClient().GetBackups(*id, *backupid)
		if err != nil {
			log.Fatal(err)
		}

		if len(backups) == 0 {
			fmt.Println()
			return
		}
		lengths := []int{16, 16, 40, 16, 24}
		tabsPrint(columns{"BACKUPID", "DATECREATED", "DESCRIPTION", "SIZE", "STATUS"}, lengths)
		for _, backup := range backups {
			tabsPrint(columns{backup.ID, backup.Created, backup.Description, backup.Size, backup.Status}, lengths)
		}
		tabsFlush()
	}
}
