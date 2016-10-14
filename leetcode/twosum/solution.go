package twosum

func twoSum(nums []int, target int) []int {
	result := []int{}
	match := make(map[int]int)
	for i, v := range nums {
		if pos, ok := match[v]; ok {
			result = append(result, pos)
			result = append(result, i)
		} else {
			match[target-v] = i
		}
	}
	return result
}
