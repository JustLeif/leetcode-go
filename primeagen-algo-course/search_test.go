package leetcodego

import "testing"

func TestLinearSearchOnArray(t *testing.T) {
	arr := []int{0, 3, 1, 2, 3, 4, 5, 6, 5, 4}
	find := 4
	idx, err := LinearSearchOnArray(arr, find)
	if err != nil {
		t.Fatalf("%d value not found in arr", find)
	}
	t.Logf("%d is the index of %d in arr", idx, find)
}

func TestBinarySearch(t *testing.T) {
	arr := []int{0, 2, 5, 5, 4, 4, 4, 6, 5, 5, 5, 4}
	find := 6
	idx, err := BinarySearchOnArray(arr, find)
	if err != nil {
		t.Fatalf("%d value not found in arr", find)
	}
	t.Logf("%d is the index of %d in arr", idx, find)
	t.Log(arr)
}

func TestCrystalBall(t *testing.T) {
	arr := []bool{false, true, true, true}
	idx := CrystalBallSearch(arr)
	if idx != 1 {
		t.Fatalf("Should have gotten idx 1, instead got %d", idx)
	}
	t.Logf("%d is index of breaking point", idx)

	arr2 := []bool{false, false, false, false, false, true}
	idx2 := CrystalBallSearch(arr2)
	if idx2 != 5 {
		t.Fatalf("Should have gotten idx 5, instead got %d", idx2)
	}
	t.Logf("%d is index of breaking point", idx2)

}
