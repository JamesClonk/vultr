package cmd

import (
	"fmt"
	"log"

	"github.com/jawher/mow.cli"
)

func reservedIpAttach(cmd *cli.Cmd) {
	ip_address := cmd.StringArg("ip_address", "", "ip_address to attach")
	attach_subid := cmd.StringArg("attach_SUBID", "", "subid to attach ip")
	cmd.Action = func() {
		err := GetClient().AttachReservedIp(*ip_address, *attach_subid)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Attach IP to SUBID\n\n")
		lengths := []int{40, 10}
		tabsPrint(Columns{"ip_address", "attached_SUBID"}, lengths)
		tabsPrint(Columns{*ip_address, *attach_subid}, lengths)
		tabsFlush()
	}
}

func reservedIpConvert(cmd *cli.Cmd) {
  cmd.Spec = "SUBID IPADDRESS"
	subid := cmd.StringArg("SUBID", "", "subid convert to reverse")
	ip_address := cmd.StringArg("IPADDRESS", "", "ipaddress to convert")
	cmd.Action = func() {
    // fmt.Println("meno-0")
    // fmt.Printf("meno-3 %v %v\n", *subid, *ip_address)
		osubid, err := GetClient().ConvertReservedIp(*subid, *ip_address)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("SUBIID to Attach IP\n\n")
		lengths := []int{10, 40, 10}
		tabsPrint(Columns{"SUBID", "ip_address", "attach_subid"}, lengths)
		tabsPrint(Columns{*subid, *ip_address, osubid.SUBID}, lengths)
		tabsFlush()
	}
}

func reservedIpCreate(cmd *cli.Cmd, ip_type string) {
	dcid := cmd.StringArg("DCID", "", "DCID Datacenter id")
	cmd.Action = func() {
		subid, err := GetClient().CreateReservedIp(*dcid, ip_type)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Reserved IP created\n\n")
		lengths := []int{10, 4}
		tabsPrint(Columns{"SUBID", "TYPE"}, lengths)
		tabsPrint(Columns{subid, ip_type}, lengths)
		tabsFlush()
	}
}

func reservedIpDestroy(cmd *cli.Cmd) {
	subid := cmd.StringArg("SUBID", "", "SUBID of the ip")
	cmd.Action = func() {
		err := GetClient().DestroyReservedIp(*subid)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Destroyed IP\n\n")
		lengths := []int{10}
		tabsPrint(Columns{"SUBID"}, lengths)
		tabsPrint(Columns{*subid}, lengths)
		tabsFlush()
	}
}

func reservedIpDetach(cmd *cli.Cmd) {
	ip_address := cmd.StringArg("ip_address", "", "ip_address to attach")
	detach_subid := cmd.StringArg("detach_SUBID", "", "subid to detach ip")
	cmd.Action = func() {
		err := GetClient().DetachReservedIp(*ip_address, *detach_subid)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Detach IP to SUBID\n\n")
		lengths := []int{40, 10}
		tabsPrint(Columns{"ip_address", "detached_SUBID"}, lengths)
		tabsPrint(Columns{*ip_address, *detach_subid}, lengths)
		tabsFlush()
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
		lengths := []int{10, 4, 4, 32, 4, 10, 10}
		tabsPrint(Columns{"SUBID", "DCID", "ip_type", "subnet", "prefix", "label", "attached"}, lengths)
		for _, ip := range ips {
			tabsPrint(Columns{ip.SUBID, ip.DCID, ip.Ip_type, ip.Subnet, ip.Subnet_size, ip.Label, ip.Attached_SUBID}, lengths)
		}
		tabsFlush()
	}
}
