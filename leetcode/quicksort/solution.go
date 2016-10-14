package quicksort

func internalQuickSort(data []int, begin int, end int) {
	if end-begin+1 <= 1 {
		return
	}
	left, right := begin, end
	pivot := data[left]
	for left < right {
		for ; left < right && data[right] >= pivot; right-- {
		}
		if left < right {
			data[left] = data[right]
		}
		for ; left < right && data[left] < pivot; left++ {
		}
		if left < right {
			data[right] = data[left]
		}
	}
	data[left] = pivot
	internalQuickSort(data, begin, left-1)
	internalQuickSort(data, left+1, end)
}
func QuickSort(data []int) {
	internalQuickSort(data, 0, len(data)-1)
}
