package util

import (
	"fmt"

	"github.com/sethigeet/table-diff/database"
)

func PrintRow(row database.Row) {
	for _, cell := range row {
		fmt.Printf("%v, ", *(cell).(*interface{}))
	}
}

func PrintTable(table database.Table) {
	for _, row := range table {
		PrintRow(row)
		fmt.Println()
	}
}
