package cmd

import (
	"fmt"
	"log"

	"github.com/jawher/mow.cli"
)

func blockStorageCreate(cmd *cli.Cmd) {
	cmd.Spec = "-n -r -s"

	name := cmd.StringOpt("n name", "", "Name/label of new block storage")
	regionID := cmd.IntOpt("r region", 1, "Region (DCID)")
	size := cmd.IntOpt("s size", 10, "Size in GB")

	cmd.Action = func() {
		storage, err := GetClient().CreateBlockStorage(*name, *regionID, *size)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Block storage created\n\n")
		lengths := []int{12, 32, 8, 8}
		tabsPrint(Columns{"SUBID", "NAME", "SIZE_GB", "DCID"}, lengths)
		tabsPrint(Columns{storage.ID, storage.Name, storage.SizeGB, storage.RegionID}, lengths)
		tabsFlush()
	}
}

func blockStorageResize(cmd *cli.Cmd) {
	cmd.Spec = "SUBID SIZE_GB"

	id := cmd.StringArg("SUBID", "", "SUBID of block storage to resize (see <storage list>)")
	size := cmd.IntArg("SIZE_GB", 0, "New size in GB")

	cmd.Action = func() {
		if err := GetClient().ResizeBlockStorage(*id, *size); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Block storage resized")
	}
}

func blockStorageLabel(cmd *cli.Cmd) {
	cmd.Spec = "SUBID NAME"

	id := cmd.StringArg("SUBID", "", "SUBID of block storage to rename (see <storage list>)")
	name := cmd.StringArg("NAME", "", "New name/label of block storage")

	cmd.Action = func() {
		if err := GetClient().LabelBlockStorage(*id, *name); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Block storage renamed")
	}
}

func blockStorageDelete(cmd *cli.Cmd) {
	id := cmd.StringArg("SUBID", "", "SUBID of block storage to delete (see <storage list>)")
	cmd.Action = func() {
		if err := GetClient().DeleteBlockStorage(*id); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Block storage deleted")
	}
}

func blockStorageList(cmd *cli.Cmd) {
	cmd.Action = func() {
		storages, err := GetClient().GetBlockStorages()
		if err != nil {
			log.Fatal(err)
		}

		if len(storages) == 0 {
			fmt.Println()
			return
		}

		lengths := []int{12, 32, 16, 8, 8, 8, 12, 24}
		tabsPrint(Columns{
			"SUBID", "NAME", "STATUS", "SIZE_GB", "COST",
			"DCID", "ATTACHED_TO", "CREATED_DATE",
		}, lengths)
		for _, storage := range storages {
			tabsPrint(Columns{
				storage.ID, storage.Name, storage.Status, storage.SizeGB, storage.Cost,
				storage.RegionID, storage.AttachedTo, storage.Created,
			}, lengths)
		}
		tabsFlush()
	}
}
