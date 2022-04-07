package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/sethigeet/table-diff/compare"
	"github.com/sethigeet/table-diff/config"
	"github.com/sethigeet/table-diff/database"
	"github.com/sethigeet/table-diff/util"
)

var (
	diffRowsMutex sync.Mutex
	wg            sync.WaitGroup
)

func main() {
	cols, _, err := database.GetTableData(config.Config.DB.Table1Name, "", "LIMIT 0")
	if err != nil {
		log.Fatalf("errors while reading the table: \n%s", err)
	}

	// Get the ids of the key columns
	keyColIds, err := util.GetColIds(config.Config.DiffAlgorithm.KeyColumns, cols)
	if err != nil {
		log.Fatalf("an error occurred: %s", err)
	}

	// Get the ids of the key columns
	ignoreColIds, err := util.GetColIds(config.Config.DiffAlgorithm.IgnoreColumns, cols)
	if err != nil {
		log.Fatalf("an error occurred: %s", err)
	}

	var diffRows [][]database.Row
	chars := "abcdefghijklmnopqrstuvwxyz"
	wg.Add(len(chars))
	for _, char := range chars {
		go func(startChar rune) {
			res := getDiffRowsByChar(keyColIds, ignoreColIds, startChar)

			diffRowsMutex.Lock()
			defer diffRowsMutex.Unlock()

			diffRows = append(diffRows, res...)

			wg.Done()
		}(char)
	}
	wg.Wait()

	// TODO: Remove this later
	// for _, rowPair := range diffRows {
	// 	for _, row := range rowPair {
	// 		util.PrintRow(row)
	// 		fmt.Printf("  |  ")
	// 	}
	// 	fmt.Println()
	// }

	database.Disconnect()
}

func getDiffRowsByChar(keyColIds, ignoreColIds []uint64, startChar rune) [][]database.Row {
	extraQuery := fmt.Sprintf(
		"WHERE LOWER(%s) LIKE %s%%'",
		config.Config.DiffAlgorithm.KeyColumns[0],
		strconv.QuoteRune(startChar)[:2],
	)
	cols1, table1, err := database.GetTableData(config.Config.DB.Table1Name, extraQuery)
	if err != nil {
		log.Fatalf("errors while reading the table: \n%s", err)
	}

	cols2, table2, err := database.GetTableData(config.Config.DB.Table2Name, extraQuery)
	if err != nil {
		log.Fatalf("errors while reading the table: \n%s", err)
	}

	for i, col1 := range cols1 {
		if cols2[i] != col1 {
			log.Fatalln("errors while reading the table: \nthe columns of both the tables are not the same!")
		}
	}

	return compare.CompareTables(table1, table2, keyColIds, ignoreColIds)
}
