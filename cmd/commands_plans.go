package cmd

import (
	"fmt"
	"log"

	"github.com/jawher/mow.cli"
)

func planList(cmd *cli.Cmd) {
	cmd.Action = func() {
		plans, err := GetClient().GetPlans()
		if err != nil {
			log.Fatal(err)
		}

		if len(plans) == 0 {
			fmt.Println()
			return
		}

		lengths := []int{12, 48, 8, 8, 8, 12, 8}
		printTabbedLine(Columns{"VPSPLANID", "NAME", "VCPU", "RAM", "DISK", "BANDWIDTH", "PRICE"}, lengths)
		for _, plan := range plans {
			printTabbedLine(Columns{plan.ID, plan.Name, plan.VCpus, plan.RAM, plan.Disk, plan.Bandwidth, plan.Price}, lengths)
		}
		tabsFlush()
	}
}
