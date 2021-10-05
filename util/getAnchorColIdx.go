package util

import (
	"fmt"

	"github.com/sethigeet/table-diff/config"
)

func GetAnchorColIdx(cols []string) (uint64, error) {
	for i, col := range cols {
		if config.Config.DiffAlgorithm.AnchorColumn == col {
			return uint64(i), nil
		}
	}

	return 0, fmt.Errorf(
		"the anchor column specified in the config(%s) does not exist in the table",
		config.Config.DiffAlgorithm.AnchorColumn,
	)
}
