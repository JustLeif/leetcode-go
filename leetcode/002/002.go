package leetcodego

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
func Solution002(nums []int, target int) []int {

	// Init map
	valMap := make(map[int]int)
	for idx, val := range nums {
		valMap[val] = idx
	}

	for i, val := range nums {
		idx, ok := valMap[target-val]
		if ok && idx != i {
			return []int{i, idx}
		}
	}

	return []int{}
}
