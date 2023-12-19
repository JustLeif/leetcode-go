package leetcode

func containsDuplicate(nums []int) bool {

	numMap := make(map[int]int)
	for _, num := range nums {
		if numMap[num] > 0 {
			return true
		}
		numMap[num] += 1
	}

	return false
}
