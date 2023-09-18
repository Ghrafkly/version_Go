package main

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"sync"
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

func printSyncMap(m *sync.Map) {
	printMap := make(map[int]int)
	m.Range(func(key, value interface{}) bool {
		printMap[key.(int)] = value.(int)
		return true
	})

	//for i := 101; i < 1000; i++ {
	//	fmt.Println(i, printMap[i])
	//}
}
