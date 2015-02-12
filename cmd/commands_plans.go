package cmd

import (
	"fmt"
	"log"

	vultr "github.com/JamesClonk/vultr/lib"
	"github.com/jawher/mow.cli"
)

func planList(cmd *cli.Cmd) {
	cmd.Spec = "[-r]"

	id := cmd.IntOpt("r region", 0, "list only available plans for region (DCID)")

	cmd.Action = func() {
		plans, err := GetClient().GetPlans()
		if err != nil {
			log.Fatal(err)
		}

		if *id != 0 {
			ids, err := GetClient().GetAvailablePlansForRegion(*id)
			if err != nil {
				log.Fatal(err)
			}
			if len(ids) == 0 {
				fmt.Println()
				return
			}
			var filteredPlans []vultr.Plan
			for _, plan := range plans {
				for _, id := range ids {
					if id == plan.ID {
						filteredPlans = append(filteredPlans, plan)
					}
				}
			}
			plans = filteredPlans
		}

		if len(plans) == 0 {
			fmt.Println()
			return
		}

		lengths := []int{12, 48, 8, 8, 8, 12, 8}
		tabsPrint(Columns{"VPSPLANID", "NAME", "VCPU", "RAM", "DISK", "BANDWIDTH", "PRICE"}, lengths)
		for _, plan := range plans {
			tabsPrint(Columns{plan.ID, plan.Name, plan.VCpus, plan.RAM, plan.Disk, plan.Bandwidth, plan.Price}, lengths)
		}
		tabsFlush()
	}
}
