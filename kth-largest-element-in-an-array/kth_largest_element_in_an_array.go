package kth_largest_element_in_an_array

func findKthLargest(nums []int, k int) int {
	QuickSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func QuickSort(arr []int, L, R int) {
	if L >= R {
		return
	}
	pivot := arr[L]
	left, right := L, R
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		if left < right {
			arr[left] = arr[right]
		}
		for left < right && arr[left] <= pivot {
			left++
		}
		if left < right {
			arr[right] = arr[left]
		}
		if left == right {
			arr[left] = pivot
		}
	}
	QuickSort(arr, L, right-1)
	QuickSort(arr, right+1, R)
}
