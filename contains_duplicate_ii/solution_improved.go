package contains_duplicate

// dynamic programming like solution
func containsNearbyDuplicate2Imp(nums []int, k int) bool {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		lastIdx, exists := dict[nums[i]]
		if exists && (i-lastIdx) <= k {
			return true
		}

		dict[nums[i]] = i
	}

	return false
}
