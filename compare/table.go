package compare

import (
	"github.com/sethigeet/table-diff/database"
)

// CompareTables compares table like structures of the Table type
// returns the rows that are not same across the two tables.
// The returned array is an array(all the pairs of different rows) of
// arrays(pair of different rows) of arrays(each row)
//
//    // For Example:
//    diffRows, err := CompareTables(table1, table2)
func CompareTables(table1, table2 database.Table, keyColIds, ignoreColIds []uint64) [][]database.Row {
	var diffRows [][]database.Row

	for _, row1 := range table1 {
		idx, found := findInOtherTable(row1, table2, keyColIds)

		// Row is not found in the other table
		if !found {
			diffRows = append(diffRows, []database.Row{row1, nil})
			continue
		}

		// Row was found in the other table
		row2 := table2[idx]

		isSame := CompareRows(row1, row2, append(keyColIds, ignoreColIds...))
		if !isSame {
			diffRows = append(diffRows, []database.Row{row1, row2})
		}

		// Remove the row from the table2 as it is also present in table1
		table2[idx] = nil
	}

	for _, row := range table2 {
		if row != nil {
			diffRows = append(diffRows, []database.Row{nil, row})
		}
	}

	return diffRows
}

// findInOtherTable tries to find the same row in the other table and return the
// index of the row found in the other table if it was found. It also returns a
// bool stating whether the row was found in the other table or not.
func findInOtherTable(rowToFind database.Row, table database.Table, keyColIds []uint64) (uint64, bool) {
	// TODO: Implement an efficient search algorithm here!

	var rowId uint64
	var isSame bool
	for _, keyColId := range keyColIds {
		valToCompare := rowToFind[keyColId]
		for i, row := range table {
			if row != nil && *(row[keyColId]).(*interface{}) == *(valToCompare).(*interface{}) {
				rowId = uint64(i)
				isSame = true
				break
			} else {
				isSame = false
			}
		}
	}

	return rowId, isSame
}
