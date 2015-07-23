package cmd

import (
	"fmt"
	"log"

	"github.com/jawher/mow.cli"
)

func dnsDomainList(cmd *cli.Cmd) {
	cmd.Action = func() {
		dnsdomains, err := GetClient().GetDnsDomains()
		if err != nil {
			log.Fatal(err)
		}

		lengths := []int{40, 24}
		tabsPrint(Columns{"DOMAIN", "DATE"}, lengths)
		for _, dnsdomain := range dnsdomains {
			tabsPrint(Columns{dnsdomain.Domain, dnsdomain.Created}, lengths)
		}
		tabsFlush()
	}
}

func dnsDomainCreate(cmd *cli.Cmd) {
	cmd.Spec = "-d -s"
	domain := cmd.StringOpt("d domain", "", "dns domain name")
	serverip := cmd.StringOpt("s serverip", "", "dns domain ip")

	cmd.Action = func() {
		err := GetClient().CreateDnsDomain(*domain, *serverip)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("DnsDomain created\n")
	}
}

func dnsDomainDelete(cmd *cli.Cmd) {
	cmd.Spec = "-d"
	domain := cmd.StringOpt("d domain", "", "dns domain name")
	cmd.Action = func() {
		if err := GetClient().DeleteDnsDomain(*domain); err != nil {
			log.Fatal(err)
		}
		fmt.Println("DnsDomain deleted")
	}
}

func dnsRecordList(cmd *cli.Cmd) {
	cmd.Spec = "-d"
	domain := cmd.StringOpt("d domain", "", "dns domain name")

	cmd.Action = func() {
		dnsrecords, err := GetClient().GetDnsRecords(*domain)
		if err != nil {
			log.Fatal(err)
		}

		lengths := []int{10, 10, 15, 50, 10}
		tabsPrint(Columns{"RECORDID", "TYPE", "NAME", "DATA", "PRIORITY"}, lengths)
		for _, dnsrecord := range dnsrecords {
			tabsPrint(Columns{dnsrecord.RecordID, dnsrecord.Type, dnsrecord.Name, dnsrecord.Data, dnsrecord.Priority}, lengths)
		}
		tabsFlush()
	}
}

func dnsRecordCreate(cmd *cli.Cmd) {
	cmd.Spec = "-d -n -t -D [OPTIONS]"

	domain := cmd.StringOpt("d domain", "", "dns domain name")
	name := cmd.StringOpt("n name", "", "dns record name")
	rtype := cmd.StringOpt("t type", "", "dns record type")
	data := cmd.StringOpt("D data", "", "dns record data")

	// options
	priority := cmd.IntOpt("priority", 0, "dns record priority")
	ttl := cmd.IntOpt("ttl", 300, "dns record priority")

	cmd.Action = func() {
		err := GetClient().CreateDnsRecord(*domain, *name, *rtype, *data, *priority, *ttl)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("DnsRecord created\n")
	}
}

func dnsRecordDelete(cmd *cli.Cmd) {
	cmd.Spec = "-d -r"

	domain := cmd.StringOpt("d domain", "", "dns domain name")
	record := cmd.IntOpt("r record", 0, "RECORDID of a dns record to delete")

	cmd.Action = func() {
		if err := GetClient().DeleteDnsRecord(*domain, *record); err != nil {
			log.Fatal(err)
		}
		fmt.Println("DnsRecord deleted")
	}
}
