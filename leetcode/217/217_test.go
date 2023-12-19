package leetcode

import "testing"

func Test217(t *testing.T) {

	res := containsDuplicate([]int{1, 2, 3, 4, 5, 6, 7, 7, 7, 5})
	if !res {
		t.Error("Expected true, recieved false.")
	}

	res2 := containsDuplicate([]int{103, 41, 32, 55, 1325, 439})
	if res2 {
		t.Error("Expected false, recieved true.")
	}

}
