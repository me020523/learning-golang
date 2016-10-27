package search_insert_position

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	cases := [][]int{
		{1, 3, 5, 6, 5, 2},
		{1, 3, 5, 6, 2, 1},
		{1, 3, 5, 6, 7, 4},
	}
	for i, _ := range cases {
		func(i int) {
			text := fmt.Sprintf("case %d(%v):", i, cases[i])
			t.Run(text, func(t1 *testing.T) {
				c := cases[i][0:4]
				target := cases[i][4]
				result := cases[i][5]

				v := searchInsert(c, target)
				if v != result {
					t.Errorf("result %d: %d", i, v)
				}
			})
		}(i)
	}
}
