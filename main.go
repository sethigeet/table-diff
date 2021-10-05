package main

import (
	"fmt"
	"log"

	"github.com/sethigeet/table-diff/compare"
	"github.com/sethigeet/table-diff/config"
	"github.com/sethigeet/table-diff/database"
	"github.com/sethigeet/table-diff/util"
)

func main() {
	cols1, table1, err := database.GetTableData(config.Config.DB.Table1Name)
	if err != nil {
		log.Fatalf("errors while reading the table: \n%s", err)
	}

	_, table2, err := database.GetTableData(config.Config.DB.Table2Name)
	if err != nil {
		log.Fatalf("errors while reading the table: \n%s", err)
	}

	// Get the ids of the key columns
	keyColIds, err := util.GetColIds(config.Config.DiffAlgorithm.KeyColumns, cols1)
	if err != nil {
		log.Fatalf("an error occurred: %s", err)
	}

	// Get the ids of the key columns
	ignoreColIds, err := util.GetColIds(config.Config.DiffAlgorithm.IgnoreColumns, cols1)
	if err != nil {
		log.Fatalf("an error occurred: %s", err)
	}

	diffRows := compare.CompareTables(table1, table2, keyColIds, ignoreColIds)

	// TODO: Remove this later
	for _, rowPair := range diffRows {
		for _, row := range rowPair {
			util.PrintRow(row)
			fmt.Printf("  |  ")
		}
		fmt.Println()
	}

	database.Disconnect()
}
