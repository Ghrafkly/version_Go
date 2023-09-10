package main

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func prettyPrint(pp PrettyPrinter) {
	headerFmt := color.New(color.FgGreen, color.Bold, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	headers := make([]interface{}, len(pp.Headers))
	for i, h := range pp.Headers {
		headers[i] = h
	}

	tbl := table.New(headers...)

	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithPadding(5)

	for _, t := range pp.Tracker {
		tbl.AddRow(t.name, t.time, t.count, t.enabled)
	}

	tbl.Print()
}
