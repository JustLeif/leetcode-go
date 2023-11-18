package leetcodego

import (
	"errors"
	"slices"
)

// O(n)
func LinearSearchOnArray(arr []int, val int) (int, error) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return i, nil
		}
	}
	return 0, errors.New("Val not found in arr!")
}

// O(log n)
// Always ask yourself, is the set ordered? If so we can take advantage of this.
func BinarySearchOnArray(arr []int, val int) (int, error) {
	slices.Sort(arr)

	high := len(arr)
	low := 0
	var mid int

	for high > low {
		mid = low + (high-low)/2

		if arr[mid] == val {
			return mid, nil
		} else if arr[mid] > val {
			high = mid
		} else if arr[mid] < val {
			low = mid + 1
		}
	}

	return 0, errors.New("Val not found in arr")
}
