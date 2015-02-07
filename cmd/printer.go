package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

var tw = new(tabwriter.Writer)

func init() {
	tw.Init(os.Stdout, 0, 8, 2, '\t', 0)
}

type Columns []interface{}

func printTabbedLine(values Columns, lengths []int) {
	if len(values) != len(lengths) {
		log.Fatalf("Internal error! Mismatch during tabbed line print. Values: %d, Lengths: %d\n", len(values), len(lengths))
	}

	for i, value := range values {
		format := "\t%s"
		if i == 0 {
			format = "%s"
		}
		fmt.Fprintf(tw, format, max(fmt.Sprintf("%v", value), lengths[i]))
	}
	fmt.Fprintf(tw, "\n")
}

func tabsFlush() {
	tw.Flush()
}

func max(input string, maxLength int) string {
	if len(input) > maxLength {
		input = input[:maxLength-2] + ".."
	}
	return input
}
