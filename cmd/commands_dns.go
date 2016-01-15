package cmd

import (
	"fmt"
	"log"

	vultr "github.com/JamesClonk/vultr/lib"
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
	domain := cmd.StringOpt("d domain", "", "DNS domain name")
	serverip := cmd.StringOpt("s serverip", "", "DNS domain ip")

	cmd.Action = func() {
		err := GetClient().CreateDnsDomain(*domain, *serverip)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("DNS domain created")
	}
}

func dnsDomainDelete(cmd *cli.Cmd) {
	cmd.Spec = "-d"
	domain := cmd.StringOpt("d domain", "", "DNS domain name")
	cmd.Action = func() {
		if err := GetClient().DeleteDnsDomain(*domain); err != nil {
			log.Fatal(err)
		}
		fmt.Println("DNS domain deleted")
	}
}

func dnsRecordList(cmd *cli.Cmd) {
	cmd.Spec = "-d"
	domain := cmd.StringOpt("d domain", "", "DNS domain name")

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

	domain := cmd.StringOpt("d domain", "", "DNS domain name")
	name := cmd.StringOpt("n name", "", "DNS record name")
	rtype := cmd.StringOpt("t type", "", "DNS record type")
	data := cmd.StringOpt("D data", "", "DNS record data")

	// options
	priority := cmd.IntOpt("priority", 0, "DNS record priority")
	ttl := cmd.IntOpt("ttl", 300, "DNS record TTL")

	cmd.Action = func() {
		err := GetClient().CreateDnsRecord(*domain, *name, *rtype, *data, *priority, *ttl)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("DNS record created")
	}
}

func dnsRecordUpdate(cmd *cli.Cmd) {
	cmd.Spec = "-d -r [OPTIONS]"

	domain := cmd.StringOpt("d domain", "", "DNS domain name")
	record := cmd.IntOpt("r record", 0, "RECORDID of a DNS record to update")

	// options
	name := cmd.StringOpt("n name", "", "DNS record name")
	data := cmd.StringOpt("D data", "", "DNS record data")
	priority := cmd.IntOpt("priority", 0, "DNS record priority")
	ttl := cmd.IntOpt("ttl", 300, "DNS record TTL")

	cmd.Action = func() {
		dnsrecord := vultr.DnsRecord{
			RecordID: *record,
			Name:     *name,
			Data:     *data,
			Priority: *priority,
			TTL:      *ttl,
		}
		err := GetClient().UpdateDnsRecord(*domain, dnsrecord)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("DNS record updated")
	}
}

func dnsRecordDelete(cmd *cli.Cmd) {
	cmd.Spec = "-d -r"

	domain := cmd.StringOpt("d domain", "", "DNS domain name")
	record := cmd.IntOpt("r record", 0, "RECORDID of a DNS record to delete")

	cmd.Action = func() {
		if err := GetClient().DeleteDnsRecord(*domain, *record); err != nil {
			log.Fatal(err)
		}
		fmt.Println("DNS record deleted")
	}
}
