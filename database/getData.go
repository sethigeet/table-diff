package database

import (
	"fmt"
	"strings"

	"github.com/sethigeet/table-diff/config"
)

// GetTableData gets the all the data from the specified table name using the
// already initialized databse connection and returns it in an array of the
// appropriate size.
//
//    // For Example:
//    data := GetTableData("my_table")
func GetTableData(tableName string) ([]string, Table, error) {
	// Construct the sql query
	query := fmt.Sprintf(
		"SELECT * FROM %s ORDER BY %s;",
		tableName,
		strings.Join(config.Config.DiffAlgorithm.KeyColumns, ","),
	)

	// Execute the sql query
	res, err := DB.Query(query)
	if err != nil {
		return nil, nil, fmt.Errorf("errors while getting the data: \n%s", err)
	}
	defer res.Close()

	var table Table

	// Get the names of all the columns that were returned
	cols, err := res.Columns()
	if err != nil {
		return nil, nil, err
	}
	// Interate through all the rows
	for res.Next() {
		// Make temporary array to store the current row
		temp := make([]interface{}, len(cols))
		// Fill the array with empty values
		for i := range temp {
			temp[i] = new(interface{})
		}
		// Copy the row into our table
		if err := res.Scan(temp...); err != nil {
			return nil, nil, err
		}
		// Add the temporary array to the table
		table = append(table, temp)
	}

	return cols, table, nil
}
