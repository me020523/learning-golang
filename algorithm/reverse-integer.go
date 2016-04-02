package algorithm
import (
	"math"
)
func ReverseInteger(x int) int {

	digitals := []int{}

	for tmp := x; tmp != 0; tmp /= 10 {
		digitals = append(digitals, tmp % 10)
	}

	var sum int = 0
	size := len(digitals)
	for i := size - 1; i >= 0; i-- {
		sum = sum + (digitals[i] * int(math.Pow(float64(10), float64(size - 1 - i))))
	}
	if sum >= 2147483647 || sum <= -2147483648{
		return 0
	}
	return sum
}
