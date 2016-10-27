package search_insert_position

func searchInsert(nums []int, target int) int {
	begin := 0
	end := len(nums) - 1
	for begin <= end {
		mid := (begin + end) >> 1
		if nums[mid] < target {
			begin = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			return mid
		}
	}
	return end + 1
}
