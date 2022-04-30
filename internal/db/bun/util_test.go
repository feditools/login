package bun

import (
	"fmt"
	"testing"
)

func TestOffset(t *testing.T) {
	tables := []struct {
		index  int
		count  int
		offset int
	}{
		{
			index:  0,
			count:  10,
			offset: 0,
		},
		{
			index:  0,
			count:  20,
			offset: 0,
		},
		{
			index:  0,
			count:  30,
			offset: 0,
		},
		{
			index:  1,
			count:  10,
			offset: 10,
		},
		{
			index:  2,
			count:  10,
			offset: 20,
		},
		{
			index:  3,
			count:  10,
			offset: 30,
		},
		{
			index:  44,
			count:  16,
			offset: 704,
		},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running offset %d %d", i, table.index, table.count)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			o := offset(table.index, table.count)
			if o != table.offset {
				t.Errorf("[%d] wrong offset: got '%d', want '%d'", i, o, table.offset)
			}
		})
	}
}

func BenchmarkOffset(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = offset(152, 30)
	}
}
