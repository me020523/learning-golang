package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 5}
	target := 9
	result := []int{0, 1}
	ret := twoSum(nums, target)
	if len(ret) == 0 {
		t.FailNow()
	}
	for i := 0; i < 2; i++ {
		if result[i] != ret[i] {
			t.FailNow()
		}
	}
}
