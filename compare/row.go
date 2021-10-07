package compare

import (
	"sync"

	"github.com/sethigeet/table-diff/database"
)

var cache sync.Map

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
	_, val := cache.Load(item)
	if val {
		return true
	}

	// check whether the item exists in the arr or not
	for _, val := range arr {
		if val == item {
			cache.Store(item, true)
			return true
		}
	}

	cache.Store(item, false)
	return false
}
