package cmd

import (
	"fmt"
	"log"

	"github.com/jawher/mow.cli"
)

func accountInfo(cmd *cli.Cmd) {
	cmd.Action = func() {
		info, err := GetClient().GetAccountInfo()
		if err != nil {
			log.Fatal(err)
		}

		lengths := []int{16, 16, 24, 24}
		printTabbedLine([]string{"BALANCE", "PENDING CHARGES", "LAST PAYMENT DATE", "LAST PAYMENT AMOUNT"}, lengths)
		printTabbedLine([]string{
			fmt.Sprintf("%v", info.Balance),
			info.PendingCharges, info.LastPaymentDate, info.LastPaymentAmount}, lengths)
		tabsFlush()
	}
}
