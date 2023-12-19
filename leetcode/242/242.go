package leetcode

import "slices"

// Time complexity O(s + t) Memory Complexity O(s + t)
func isAnagram(s string, t string) bool {

	sMap := make(map[rune]int)
	for _, r := range s {
		sMap[r] += 1
	}

	tMap := make(map[rune]int)
	for _, r := range t {
		tMap[r] += 1
	}

	for k, sNum := range sMap {
		if tMap[k] != sNum {
			return false
		}
	}

	return true
}

// Time Complexity O(s + t) Memory Complexity O(s + t)
func isAnagramOptimal(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sRunes := []rune(s)
	tRunes := []rune(t)

	slices.Sort(sRunes)
	slices.Sort(tRunes)

	for i := 0; i < len(sRunes); i++ {
		if sRunes[i] != tRunes[i] {
			return false
		}
	}

	return true
}
