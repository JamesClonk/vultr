package cmd

import (
	"fmt"
	"log"

	"github.com/jawher/mow.cli"
)

func serversCreate(cmd *cli.Cmd) {
	cmd.Spec = "-n | --name -r | --region -p | --plan -o | --os"

	name := cmd.StringOpt("n name", "", "Name of new virtual machine")
	regionId := cmd.IntOpt("r region", 0, "Region (DCID)")
	planId := cmd.IntOpt("p plan", 0, "Plan (VPSPLANID)")
	osId := cmd.IntOpt("o os", 0, "Operating system (OSID)")

	cmd.Action = func() {
		server, err := GetClient().CreateServer(*name, *regionId, *planId, *osId, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Virtual machine create success!\n")
		lengths := []int{12, 32, 8, 12, 8}
		printTabbedLine(Columns{"SUBID", "NAME", "DCID", "VPSPLANID", "OSID"}, lengths)
		printTabbedLine(Columns{server.ID, server.Name, server.RegionID, server.PlanID, *osId}, lengths)
		tabsFlush()
	}
}

func serversDelete(cmd *cli.Cmd) {
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <server list>)")

	cmd.Action = func() {
		if err := GetClient().DeleteServer(*id); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Virtual machine delete success!")
	}
}

func serversList(cmd *cli.Cmd) {
	cmd.Action = func() {
		servers, err := GetClient().GetServers()
		if err != nil {
			log.Fatal(err)
		}

		if len(servers) == 0 {
			fmt.Println()
			return
		}

		lengths := []int{12, 16, 24, 32, 32, 32, 8, 8, 24, 12, 8}
		printTabbedLine(Columns{
			"SUBID",
			"STATUS",
			"IP",
			"NAME",
			"OS",
			"LOCATION",
			"VCPU",
			"RAM",
			"DISK",
			"BANDWIDTH",
			"COST"}, lengths)
		for _, server := range servers {
			printTabbedLine(Columns{
				server.ID,
				server.Status,
				server.MainIP,
				server.Name,
				server.OS,
				server.Location,
				server.VCpus,
				server.RAM,
				server.Disk,
				server.AllowedBandwidth,
				server.Cost,
			}, lengths)
		}
		tabsFlush()
	}
}

func serversShow(cmd *cli.Cmd) {
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <server list>)")

	cmd.Action = func() {
		server, err := GetClient().GetServer(*id)
		if err != nil {
			log.Fatal(err)
		}

		if server.ID == "" {
			fmt.Printf("No virtual machine with SUBID %v found!\n", *id)
			return
		}

		-- display full information
		lengths := []int{12, 16, 24, 32, 32, 32, 8, 8, 24, 12, 8}
		printTabbedLine(Columns{
			"SUBID",
			"STATUS",
			"IP",
			"NAME",
			"OS",
			"LOCATION",
			"VCPU",
			"RAM",
			"DISK",
			"BANDWIDTH",
			"COST"}, lengths)
		printTabbedLine(Columns{
			server.ID,
			server.Status,
			server.MainIP,
			server.Name,
			server.OS,
			server.Location,
			server.VCpus,
			server.RAM,
			server.Disk,
			server.AllowedBandwidth,
			server.Cost,
		}, lengths)
		tabsFlush()
	}
}
