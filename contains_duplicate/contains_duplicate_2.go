package contains_duplicate

// 219. Contains Duplicate II https://leetcode.com/problems/contains-duplicate-ii/

// Given an integer array nums and an integer k,
// return true if there are two distinct indices i and j
// in the array such that nums[i] == nums[j] and abs(i - j) <= k.

func containsNearbyDuplicate2(nums []int, k int) bool {
	window := make(map[int]struct{})

	if len(nums) <= k {
		for i := 0; i < len(nums); i++ {
			window[nums[i]] = struct{}{}
		}
		return len(window) != len(nums)
	}

	right := 0
	// take first window
	for ; right <= k; right++ {
		window[nums[right]] = struct{}{}
	}

	// for unique - elements in window must be more than step
	// e.g. len(window) = 3 if k = 2  has not duplicates
	if len(window) <= k {
		return true
	}

	for ; right < len(nums); right++ {
		delete(window, nums[right-k-1])
		window[nums[right]] = struct{}{}

		if len(window) <= k {
			return true
		}
	}

	return false
}
