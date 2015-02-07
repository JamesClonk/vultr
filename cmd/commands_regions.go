package cmd

import (
	"fmt"
	"log"

	"github.com/jawher/mow.cli"
)

func regionList(cmd *cli.Cmd) {
	cmd.Action = func() {
		regions, err := GetClient().GetRegions()
		if err != nil {
			log.Fatal(err)
		}

		if len(regions) == 0 {
			fmt.Println()
			return
		}

		lengths := []int{8, 48, 24, 8, 8}
		printTabbedLine(Columns{"DCID", "NAME", "CONTINENT", "COUNTRY", "STATE"}, lengths)
		for _, region := range regions {
			printTabbedLine(Columns{region.ID, region.Name, region.Continent, region.Country, region.State}, lengths)
		}
		tabsFlush()
	}
}
