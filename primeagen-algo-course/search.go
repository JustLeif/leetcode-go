package leetcodego

import (
	"errors"
	"math"
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

// Runtime of O(sqrt(n))
// Assuming an arr of booleans, each position representing a floor of a building, figure out in an optimized way how to determine at which floor a crystal ball with break if you drop it from that floor (You have two crystal balls to drop)
func CrystalBallSearch(arr []bool) int {
	sqrtJump := int(math.Sqrt(float64(len(arr))))
	var i int
	for i = sqrtJump; i < len(arr); i += sqrtJump {
		if arr[i] {
			break
		}
	}
	i -= sqrtJump
	for j := 0; j <= sqrtJump && i < len(arr); i, j = i+1, j+1 {

		if arr[i] {
			return i
		}
	}
	return -1
}
