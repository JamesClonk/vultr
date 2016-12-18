package cmd

import (
	"fmt"
	"log"

	"github.com/jawher/mow.cli"
)

func reservedIpAttach(cmd *cli.Cmd) {
	cmd.Spec = "SUBID IP_ADDRESS"

	serverId := cmd.StringArg("SUBID", "", "SUBID of virtual machine to attach to (see <servers>)")
	ip := cmd.StringArg("IP_ADDRESS", "", "IP address to attach (see <reservedips>)")

	cmd.Action = func() {
		if err := GetClient().AttachReservedIp(*ip, *serverId); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Reserved IP attached")
	}
}

func reservedIpConvert(cmd *cli.Cmd) {
	cmd.Spec = "SUBID IP_ADDRESS"

	serverId := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	ip := cmd.StringArg("IP_ADDRESS", "", "IP address to convert to reserved IP")

	cmd.Action = func() {
		id, err := GetClient().ConvertReservedIp(*serverId, *ip)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Reserved IP converted\n\n")
		lengths := []int{12, 48, 12}
		tabsPrint(Columns{"ID", "IP_ADDRESS", "ATTACHED_TO"}, lengths)
		tabsPrint(Columns{id, *ip, *serverId}, lengths)
		tabsFlush()
	}
}

func reservedIpCreate(cmd *cli.Cmd) {
	cmd.Spec = "[-r -t]"

	regionID := cmd.IntOpt("r region", 1, "Region (DCID)")
	ipType := cmd.StringOpt("t type", "v4", "Type of new reserved IP (v4 or v6)")

	cmd.Action = func() {
		id, err := GetClient().CreateReservedIp(*regionID, *ipType)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Reserved IP created\n\n")
		lengths := []int{12, 6, 10}
		tabsPrint(Columns{"ID", "TYPE", "DCID"}, lengths)
		tabsPrint(Columns{id, *ipType, *regionID}, lengths)
		tabsFlush()
	}
}

func reservedIpDestroy(cmd *cli.Cmd) {
	cmd.Spec = "SUBID"

	id := cmd.StringArg("SUBID", "", "SUBID of reserved IP (see <reservedips>)")

	cmd.Action = func() {
		if err := GetClient().DestroyReservedIp(*id); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Reserved IP deleted")
	}
}

func reservedIpDetach(cmd *cli.Cmd) {
	cmd.Spec = "SUBID IP_ADDRESS"

	serverId := cmd.StringArg("SUBID", "", "SUBID of virtual machine to detach from (see <servers>)")
	ip := cmd.StringArg("IP_ADDRESS", "", "IP address to detach (see <reservedips>)")

	cmd.Action = func() {
		if err := GetClient().DetachReservedIp(*ip, *serverId); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Reserved IP detached")
	}
}

func reservedIpList(cmd *cli.Cmd) {
	cmd.Action = func() {
		ips, err := GetClient().ListReservedIp()
		if err != nil {
			log.Fatal(err)
		}

		if len(ips) == 0 {
			fmt.Println()
			return
		}

		lengths := []int{12, 8, 8, 48, 6, 32, 12}
		tabsPrint(Columns{"SUBID", "DCID", "IP_TYPE", "SUBNET", "SIZE", "LABEL", "ATTACHED_TO"}, lengths)
		for _, ip := range ips {
			tabsPrint(Columns{
				ip.ID,
				ip.RegionID,
				ip.IPType,
				ip.Subnet,
				ip.SubnetSize,
				ip.Label,
				ip.AttachedTo,
			}, lengths)
		}
		tabsFlush()
	}
}
