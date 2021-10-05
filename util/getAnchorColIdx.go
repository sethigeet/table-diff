package util

import (
	"fmt"

	"github.com/sethigeet/table-diff/config"
)

// GetKeyColId gets the index of the column that is being used as the key so
// that the index can later be used for indexing rows.
func GetKeyColId(keyCol string, cols []string) (uint64, bool) {
	for i, col := range cols {
		if keyCol == col {
			return uint64(i), true
		}
	}

	return 0, false
}

// GetKeyColIds gets the indexes of all the columns that are being used as the
// keys in the configuration so that the indexes can later be used for indexing
// rows.
func GetKeyColIds(cols []string) ([]uint64, error) {
	keyColIds := make([]uint64, len(config.Config.DiffAlgorithm.KeyColumns))

	for i, keyCol := range config.Config.DiffAlgorithm.KeyColumns {
		idx, found := GetKeyColId(keyCol, cols)
		if !found {
			return nil, fmt.Errorf(
				"the anchor column specified in the config(%s) does not exist in the table",
				keyCol,
			)
		}

		keyColIds[i] = idx
	}

	return keyColIds, nil
}
