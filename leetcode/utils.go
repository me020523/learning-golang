package utils

import (
	"math/rand"
	"time"
)

func GenRandomInt(capacity int) []int {
	ret := []int{}
	if capacity <= 0 {
		return ret
	}
	rand.Seed(time.Now().Unix())
	for i := 0; i < capacity; i++ {
		tmp := rand.Intn(capacity)
		ret = append(ret, tmp)
	}
	return ret
}

func EqualInts(left, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i, _ := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
