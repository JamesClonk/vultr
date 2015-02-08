package cmd

import (
	"fmt"
	"log"

	"github.com/jawher/mow.cli"
)

func serversCreate(cmd *cli.Cmd) {
	cmd.Spec = "-n | --name -r | --region -p | --plan -o | --os"

	name := cmd.StringOpt("n name", "", "Name of new virtual machine")
	regionID := cmd.IntOpt("r region", 0, "Region (DCID)")
	planID := cmd.IntOpt("p plan", 0, "Plan (VPSPLANID)")
	osID := cmd.IntOpt("o os", 0, "Operating system (OSID)")

	cmd.Action = func() {
		server, err := GetClient().CreateServer(*name, *regionID, *planID, *osID, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Virtual machine create success!\n")
		lengths := []int{12, 32, 8, 12, 8}
		printTabbedLine(Columns{"SUBID", "NAME", "DCID", "VPSPLANID", "OSID"}, lengths)
		printTabbedLine(Columns{server.ID, server.Name, server.RegionID, server.PlanID, *osID}, lengths)
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
	cmd.Spec = "SUBID [-f | --full]"

	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <server list>)")
	full := cmd.BoolOpt("f full", false, "Display full length of KVM URL")

	cmd.Action = func() {
		server, err := GetClient().GetServer(*id)
		if err != nil {
			log.Fatal(err)
		}

		if server.ID == "" {
			fmt.Printf("No virtual machine with SUBID %v found!\n", *id)
			return
		}

		keyLength := 64
		if *full {
			keyLength = 1024
		}
		lengths := []int{24, keyLength}

		printTabbedLine(Columns{"Id (SUBID):", server.ID}, lengths)
		printTabbedLine(Columns{"Name:", server.Name}, lengths)
		printTabbedLine(Columns{"Operating system:", server.OS}, lengths)
		printTabbedLine(Columns{"Status:", server.Status}, lengths)
		printTabbedLine(Columns{"Power status:", server.PowerStatus}, lengths)
		printTabbedLine(Columns{"Location:", server.Location}, lengths)
		printTabbedLine(Columns{"Region (DCID):", server.RegionID}, lengths)
		printTabbedLine(Columns{"VCPU count:", server.VCpus}, lengths)
		printTabbedLine(Columns{"RAM:", server.RAM}, lengths)
		printTabbedLine(Columns{"Disk:", server.Disk}, lengths)
		printTabbedLine(Columns{"Allowed bandwidth:", server.AllowedBandwidth}, lengths)
		printTabbedLine(Columns{"Current bandwidth:", server.CurrentBandwidth}, lengths)
		printTabbedLine(Columns{"Cost per month:", server.Cost}, lengths)
		printTabbedLine(Columns{"Pending charges:", server.PendingCharges}, lengths)
		printTabbedLine(Columns{"Plan (VPSPLANID):", server.PlanID}, lengths)
		printTabbedLine(Columns{"IP:", server.MainIP}, lengths)
		printTabbedLine(Columns{"Netmask:", server.NetmaskV4}, lengths)
		printTabbedLine(Columns{"Gateway:", server.GatewayV4}, lengths)
		printTabbedLine(Columns{"Internal IP:", server.InternalIP}, lengths)
		printTabbedLine(Columns{"IPv6 IP:", server.MainIPV6}, lengths)
		printTabbedLine(Columns{"IPv6 Network:", server.NetworkV6}, lengths)
		printTabbedLine(Columns{"IPv6 Network Size:", server.NetworkSizeV6}, lengths)
		printTabbedLine(Columns{"Created date:", server.Created}, lengths)
		printTabbedLine(Columns{"Default password:", server.DefaultPassword}, lengths)
		printTabbedLine(Columns{"Auto backups:", server.AutoBackups}, lengths)
		printTabbedLine(Columns{"KVM URL:", server.KVMUrl}, lengths)
		tabsFlush()
	}
}
