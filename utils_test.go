package gobulkinsert

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_splitObjects(t *testing.T) {
	var objArr []interface{}
	for i := 0; i < 100; i++ {
		objArr = append(objArr, i)
	}

	objSet := splitObjects(objArr, 30)

	assert.Len(t, objSet, 4)
	assert.Len(t, objSet[len(objSet)-1], 10)
}

func Test_sortedKeys(t *testing.T) {
	value := map[string]interface{}{}
	for i := 1; i <= 9; i++ {
		value[strconv.Itoa(i)] = i
	}

	keys := sortedKeys(value)

	assert.Len(t, keys, 9)
	assert.Equal(t, keys[0], "1")
	assert.Equal(t, keys[len(keys)-1], "9")
}
