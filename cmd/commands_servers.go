package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	vultr "github.com/JamesClonk/vultr/lib"
	"github.com/jawher/mow.cli"
)

func serversCreate(cmd *cli.Cmd) {
	cmd.Spec = "-n -r -p -o [OPTIONS]"

	name := cmd.StringOpt("n name", "", "Name of new virtual machine")
	regionID := cmd.IntOpt("r region", 1, "Region (DCID)")
	planID := cmd.IntOpt("p plan", 29, "Plan (VPSPLANID)")
	osID := cmd.IntOpt("o os", 160, "Operating system (OSID)")

	// options
	ipxe := cmd.StringOpt("ipxe", "", "Chainload the specified URL on bootup, via iPXE, for custom OS")
	iso := cmd.IntOpt("iso", 0, "ISOID of a specific ISO to mount during the deployment, for custom OS")
	script := cmd.IntOpt("s script", 0, "SCRIPTID of a startup script to execute on boot (see <scripts>)")
	userDataFile := cmd.StringOpt("user-data", "", "Path to file with user-data")
	snapshot := cmd.StringOpt("snapshot", "", "SNAPSHOTID (see <snapshots>) to restore for the initial installation")
	sshkey := cmd.StringOpt("k sshkey", "", "SSHKEYID (see <sshkeys>) of SSH key to apply to this server on install")
	ipv6 := cmd.BoolOpt("ipv6", false, "Assign an IPv6 subnet to this virtual machine (where available)")
	privateNetworking := cmd.BoolOpt("private-networking", false, "Add private networking support for this virtual machine")
	autoBackups := cmd.BoolOpt("autobackups", false, "Enable automatic backups for this virtual machine")
	notifyActivate := cmd.BoolOpt("notify-activate", true, "Send an activation email when the server is ready")

	cmd.Action = func() {
		options := &vultr.ServerOptions{
			IPXEChainURL:         *ipxe,
			ISO:                  *iso,
			Script:               *script,
			Snapshot:             *snapshot,
			SSHKey:               *sshkey,
			IPV6:                 *ipv6,
			PrivateNetworking:    *privateNetworking,
			AutoBackups:          *autoBackups,
			DontNotifyOnActivate: !*notifyActivate,
		}
		if *userDataFile != "" {
			data, err := ioutil.ReadFile(*userDataFile)
			if err != nil {
				log.Fatal(err)
			}
			options.UserData = string(data)
		}

		server, err := GetClient().CreateServer(*name, *regionID, *planID, *osID, options)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Virtual machine created\n\n")
		lengths := []int{12, 32, 8, 12, 8}
		tabsPrint(Columns{"SUBID", "NAME", "DCID", "VPSPLANID", "OSID"}, lengths)
		tabsPrint(Columns{server.ID, server.Name, server.RegionID, server.PlanID, *osID}, lengths)
		tabsFlush()
	}
}

func serversRename(cmd *cli.Cmd) {
	cmd.Spec = "SUBID -n"
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	name := cmd.StringOpt("n name", "", "new name of virtual machine")
	cmd.Action = func() {
		if err := GetClient().RenameServer(*id, *name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Virtual machine renamed to: %v\n", *name)
	}
}

func serversStart(cmd *cli.Cmd) {
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	cmd.Action = func() {
		if err := GetClient().StartServer(*id); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Virtual machine (re)started")
	}
}

func serversHalt(cmd *cli.Cmd) {
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	cmd.Action = func() {
		if err := GetClient().HaltServer(*id); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Virtual machine halted")
	}
}

func serversReboot(cmd *cli.Cmd) {
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	cmd.Action = func() {
		if err := GetClient().RebootServer(*id); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Virtual machine rebooted")
	}
}

func serversReinstall(cmd *cli.Cmd) {
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	cmd.Action = func() {
		if err := GetClient().ReinstallServer(*id); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Virtual machine reinstalled")
	}
}

func serversChangeOS(cmd *cli.Cmd) {
	cmd.Spec = "SUBID -o"
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	osID := cmd.IntOpt("o os", 0, "Operating system (OSID)")
	cmd.Action = func() {
		if err := GetClient().ChangeOSofServer(*id, *osID); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Virtual machine operating system changed to: %v\n", *osID)
	}
}

func serversAttachISO(cmd *cli.Cmd) {
	cmd.Spec = "SUBID -i"
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	isoID := cmd.IntOpt("i iso", 0, "ISOID of ISO image")
	cmd.Action = func() {
		if err := GetClient().AttachISOtoServer(*id, *isoID); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Virtual machine ISO attached, server will reboot: %v\n", *isoID)
	}
}

func serversDetachISO(cmd *cli.Cmd) {
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	cmd.Action = func() {
		if err := GetClient().DetachISOfromServer(*id); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Virtual machine ISO detached, server will reboot\n")
	}
}

func serversStatusISO(cmd *cli.Cmd) {
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	cmd.Action = func() {
		iso, err := GetClient().GetISOStatusofServer(*id)
		if err != nil {
			log.Fatal(err)
		}

		if iso.State == "" {
			fmt.Printf("Could not determine ISO state of virtual machine with SUBID %v!\n", *id)
			return
		}

		lengths := []int{12, 8, 24}
		tabsPrint(Columns{"SUBID", "ISOID", "STATE"}, lengths)
		tabsPrint(Columns{*id, iso.ISOID, iso.State}, lengths)
		tabsFlush()
	}
}

func serversListOS(cmd *cli.Cmd) {
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	cmd.Action = func() {
		os, err := GetClient().ListOSforServer(*id)
		if err != nil {
			log.Fatal(err)
		}

		if len(os) == 0 {
			fmt.Println()
			return
		}

		lengths := []int{8, 32, 8, 16, 8, 12}
		tabsPrint(Columns{"OSID", "NAME", "ARCH", "FAMILY", "WINDOWS", "SURCHARGE"}, lengths)
		for _, os := range os {
			tabsPrint(Columns{os.ID, os.Name, os.Arch, os.Family, os.Windows, os.Surcharge}, lengths)
		}
		tabsFlush()
	}
}

func serversDelete(cmd *cli.Cmd) {
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	cmd.Action = func() {
		if err := GetClient().DeleteServer(*id); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Virtual machine deleted")
	}
}

func serversBandwidth(cmd *cli.Cmd) {
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	cmd.Action = func() {
		bandwidth, err := GetClient().BandwidthOfServer(*id)
		if err != nil {
			log.Fatal(err)
		}

		if len(bandwidth) == 0 {
			fmt.Println()
			return
		}

		lengths := []int{24, 24, 24}
		tabsPrint(Columns{"DATE", "INCOMING", "OUTGOING"}, lengths)
		for _, b := range bandwidth {
			tabsPrint(Columns{b["date"], b["incoming"], b["outgoing"]}, lengths)
		}
		tabsFlush()
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
		tabsPrint(Columns{
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
			tabsPrint(Columns{
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
	cmd.Spec = "SUBID [-f]"

	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
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

		tabsPrint(Columns{"Id (SUBID):", server.ID}, lengths)
		tabsPrint(Columns{"Name:", server.Name}, lengths)
		if len(server.Tag) != 0 {
			tabsPrint(Columns{"Tag:", server.Tag}, lengths)
		}
		tabsPrint(Columns{"Operating system:", server.OS}, lengths)
		tabsPrint(Columns{"Status:", server.Status}, lengths)
		tabsPrint(Columns{"Power status:", server.PowerStatus}, lengths)
		tabsPrint(Columns{"Server state:", server.ServerState}, lengths)
		tabsPrint(Columns{"Location:", server.Location}, lengths)
		tabsPrint(Columns{"Region (DCID):", server.RegionID}, lengths)
		tabsPrint(Columns{"VCPU count:", server.VCpus}, lengths)
		tabsPrint(Columns{"RAM:", server.RAM}, lengths)
		tabsPrint(Columns{"Disk:", server.Disk}, lengths)
		tabsPrint(Columns{"Allowed bandwidth:", server.AllowedBandwidth}, lengths)
		tabsPrint(Columns{"Current bandwidth:", server.CurrentBandwidth}, lengths)
		tabsPrint(Columns{"Cost per month:", server.Cost}, lengths)
		tabsPrint(Columns{"Pending charges:", server.PendingCharges}, lengths)
		tabsPrint(Columns{"Plan (VPSPLANID):", server.PlanID}, lengths)
		tabsPrint(Columns{"IP:", server.MainIP}, lengths)
		tabsPrint(Columns{"Netmask:", server.NetmaskV4}, lengths)
		tabsPrint(Columns{"Gateway:", server.GatewayV4}, lengths)
		if len(server.InternalIP) != 0 {
			tabsPrint(Columns{"Internal IP:", server.InternalIP}, lengths)
		}
		for n, v6network := range server.V6Networks {
			tabsPrint(Columns{fmt.Sprintf("#%d IPv6 IP:", n+1), v6network.MainIP}, lengths)
			tabsPrint(Columns{fmt.Sprintf("#%d IPv6 Network:", n+1), v6network.Network}, lengths)
			tabsPrint(Columns{fmt.Sprintf("#%d IPv6 Network Size:", n+1), v6network.NetworkSize}, lengths)
		}
		tabsPrint(Columns{"Created date:", server.Created}, lengths)
		tabsPrint(Columns{"Default password:", server.DefaultPassword}, lengths)
		tabsPrint(Columns{"Auto backups:", server.AutoBackups}, lengths)
		tabsPrint(Columns{"KVM URL:", server.KVMUrl}, lengths)
		tabsFlush()
	}
}

func ipv4List(cmd *cli.Cmd) {
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	cmd.Action = func() {
		list, err := GetClient().ListIPv4(*id)
		if err != nil {
			log.Fatal(err)
		}

		if len(list) == 0 {
			fmt.Printf("No IPv4 information for virtual machine with SUBID %v found!\n", *id)
			return
		}

		lengths := []int{24, 24, 24, 32, 48}
		tabsPrint(Columns{"IP", "NETMASK", "GATEWAY", "TYPE", "REVERSE DNS"}, lengths)
		for _, ip := range list {
			tabsPrint(Columns{
				ip.IP,
				ip.Netmask,
				ip.Gateway,
				ip.Type,
				ip.ReverseDNS,
			}, lengths)
		}
		tabsFlush()
	}
}

func ipv6List(cmd *cli.Cmd) {
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	cmd.Action = func() {
		list, err := GetClient().ListIPv6(*id)
		if err != nil {
			log.Fatal(err)
		}

		if len(list) == 0 {
			fmt.Printf("No IPv6 information for virtual machine with SUBID %v found!\n", *id)
			return
		}

		lengths := []int{48, 32, 24, 32}
		tabsPrint(Columns{"IP", "NETWORK", "NETWORK SIZE", "TYPE"}, lengths)
		for _, ip := range list {
			tabsPrint(Columns{
				ip.IP,
				ip.Network,
				ip.NetworkSize,
				ip.Type,
			}, lengths)
		}
		tabsFlush()
	}
}

func reverseIpv6List(cmd *cli.Cmd) {
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	cmd.Action = func() {
		list, err := GetClient().ListIPv6ReverseDNS(*id)
		if err != nil {
			log.Fatal(err)
		}

		if len(list) == 0 {
			fmt.Printf("No IPv6 reverse DNS entries for virtual machine with SUBID %v found!\n", *id)
			return
		}

		lengths := []int{48, 64}
		tabsPrint(Columns{"IP", "REVERSE DNS"}, lengths)
		for _, ip := range list {
			tabsPrint(Columns{ip.IP, ip.ReverseDNS}, lengths)
		}
		tabsFlush()
	}
}

func reverseIpv6Delete(cmd *cli.Cmd) {
	cmd.Spec = "SUBID IP"
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	ip := cmd.StringArg("IP", "", "IPv6 of virtual machine (see <list-ipv6>)")

	cmd.Action = func() {
		if err := GetClient().DeleteIPv6ReverseDNS(*id, *ip); err != nil {
			log.Fatal(err)
		}
		fmt.Println("IPv6 reverse DNS deleted")
	}
}

func reverseIpv6Set(cmd *cli.Cmd) {
	cmd.Spec = "SUBID IP DNS"
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	ip := cmd.StringArg("IP", "", "IPv6 of virtual machine (see <list-ipv6>)")
	entry := cmd.StringArg("DNS", "", "Reverse DNS entry")

	cmd.Action = func() {
		if err := GetClient().SetIPv6ReverseDNS(*id, *ip, *entry); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("IPv6 reverse DNS set to: %v\n", *entry)
	}
}

func reverseIpv4Default(cmd *cli.Cmd) {
	cmd.Spec = "SUBID IP"
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	ip := cmd.StringArg("IP", "", "IPv4 of virtual machine (see <list-ipv4>)")

	cmd.Action = func() {
		if err := GetClient().DefaultIPv4ReverseDNS(*id, *ip); err != nil {
			log.Fatal(err)
		}
		fmt.Println("IPv4 reverse DNS set back to original setting")
	}
}

func reverseIpv4Set(cmd *cli.Cmd) {
	cmd.Spec = "SUBID IP DNS"
	id := cmd.StringArg("SUBID", "", "SUBID of virtual machine (see <servers>)")
	ip := cmd.StringArg("IP", "", "IPv4 of virtual machine (see <list-ipv4>)")
	entry := cmd.StringArg("DNS", "", "Reverse DNS entry")

	cmd.Action = func() {
		if err := GetClient().SetIPv4ReverseDNS(*id, *ip, *entry); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("IPv4 reverse DNS set to: %v\n", *entry)
	}
}
