package algorithm

import (
	"math"
)

func countDigital(x int) int {
	number := 0
	if x == 0 {
		return 1
	}
	for tmp := x; tmp != 0; tmp /= 10 {
		number++
	}
	return number
}
func getLastNthDigital(x, n int) int {
	v1 := x / int(math.Pow10(n-1))
	v2 := v1 % 10
	return v2
}
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	digitalNum := countDigital(x)

	if digitalNum == 1 {
		return true
	}

	halfNum := digitalNum >> 1
	ret := true
	for i := 0; i < halfNum; i++ {
		lastNth := getLastNthDigital(x, i+1)
		nth := getLastNthDigital(x, digitalNum-i)
		if lastNth != nth {
			ret = false
			break
		}
	}
	return ret
}
