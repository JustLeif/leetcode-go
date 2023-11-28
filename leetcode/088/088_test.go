package leetcodego

import (
	"testing"
)

func Test088(t *testing.T) {
	testarr := []int{2, 3, 6, 10, 0, 0, 0, 0}
	Solution088(testarr, 4, []int{4, 7, 3, 6}, 4)
	// This if statement check if the two arrays are equal
	if !compareSlices(testarr, []int{2, 3, 3, 4, 6, 6, 7, 10}) {
		t.Fail()
	}
}

func compareSlices(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
