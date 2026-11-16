// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package column

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSortIndexDuplicates keeps every record with the same value in the
// sort index: equal values are ordered by the record.
func TestSortIndexDuplicates(t *testing.T) {
	c := NewCollection()
	c.CreateColumn("col1", ForString())
	assert.NoError(t, c.CreateSortIndex("sortedCol1", "col1"))
	for _, name := range []string{"bob", "alice", "bob", "carter", "alice", "bob"} {
		c.Insert(func(r Row) error {
			r.SetString("col1", name)
			return nil
		})
	}
	sorted := func() (names []string, idx []uint32) {
		c.Query(func(txn *Txn) error {
			col1 := txn.String("col1")
			return txn.Ascend("sortedCol1", func(i uint32) {
				name, _ := col1.Get()
				names = append(names, name)
				idx = append(idx, i)
			})
		})
		return
	}
	names, idx := sorted()
	assert.Equal(t, []string{"alice", "alice", "bob", "bob", "bob", "carter"}, names)
	assert.Equal(t, []uint32{1, 4, 0, 2, 5, 3}, idx)

	// an update moves only the changed record, a delete removes only it
	assert.NoError(t, c.QueryAt(2, func(r Row) error {
		r.SetString("col1", "alice")
		return nil
	}))
	assert.True(t, c.DeleteAt(0))
	names, idx = sorted()
	assert.Equal(t, []string{"alice", "alice", "alice", "bob", "carter"}, names)
	assert.Equal(t, []uint32{1, 2, 4, 5, 3}, idx)
	indexCol, _ := c.cols.Load("sortedCol1")
	assert.Equal(t, 5, indexCol.Column.(*columnSortIndex).btree.Len())
	assert.Equal(t, 5, len(indexCol.Column.(*columnSortIndex).backMap))
}
