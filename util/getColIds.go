package util

import (
	"fmt"
)

// GetColId gets the index of the column that is being used as the key so
// that the index can later be used for indexing rows.
func GetColId(configCol string, cols []string) (uint64, bool) {
	for i, col := range cols {
		if configCol == col {
			return uint64(i), true
		}
	}

	return 0, false
}

// GetColIds gets the indexes of all the columns that are being used as the
// keys in the configuration so that the indexes can later be used for indexing
// rows.
func GetColIds(configCols, cols []string) ([]uint64, error) {
	colIds := make([]uint64, len(configCols))

	for i, configCol := range configCols {
		idx, found := GetColId(configCol, cols)
		if !found {
			return nil, fmt.Errorf(
				"the column specified in the config(%s) does not exist in the table",
				configCol,
			)
		}

		colIds[i] = idx
	}

	return colIds, nil
}
