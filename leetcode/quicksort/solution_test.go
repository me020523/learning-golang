package quicksort

import (
	"sort"
	"testing"

	"github.com/me020523/learning-golang/leetcode"
)

func TestQuickSort(t *testing.T) {
	const NUM_OF_DATA = 1000000
	data := utils.GenRandomInt(NUM_OF_DATA)
	target := make([]int, NUM_OF_DATA)
	count := copy(target, data)
	if count < NUM_OF_DATA {
		t.Fatal("Failed to copy array data")
	}
	sort.Ints(target)
	QuickSort(data)
	if !utils.EqualInts(data, target) {
		t.FailNow()
	}
}
