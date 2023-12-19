package leetcode

import "testing"

func Test242(t *testing.T) {

	res := isAnagram("value", "eulav")
	if !res {
		t.Error("Expected res to be true, recieved false")
	}

	res2 := isAnagram("vvvvvvvvvv", "value")
	if res2 {
		t.Error("Expected res to be false, recieved true")
	}

}

func Test242Optimal(t *testing.T) {

	res := isAnagramOptimal("value", "eeevv")
	if res {
		t.Error("Expected res to be false, recieved true")
	}

	res2 := isAnagramOptimal("vvvviii", "viviviv")
	if !res2 {
		t.Error("Expected res2 to be true, recieved false")
	}
}
