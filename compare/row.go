package compare

import (
	"github.com/sethigeet/table-diff/database"
)

var cache map[uint64]interface{} = map[uint64]interface{}{}

// CompareRows compares two rows and returns a bool stating whether the two rows
// are the same or whether they are differences between them.
func CompareRows(row1, row2 database.Row, ignoreColIds []uint64) bool {
	var isSame bool

	for i, cell1 := range row1 {
		if valExistsInArr(ignoreColIds, uint64(i)) {
			continue
		}

		if *(row2[i]).(*interface{}) == *(cell1).(*interface{}) {
			isSame = true
		} else {
			isSame = false
			break
		}
	}

	return isSame
}

func valExistsInArr(arr []uint64, item uint64) bool {
	// check if the result already exists in the cache
	if cache[item] == true {
		return true
	}

	// check whether the item exists in the arr or not
	for _, val := range arr {
		if val == item {
			cache[item] = true
			return true
		}
	}

	cache[item] = false
	return false
}
