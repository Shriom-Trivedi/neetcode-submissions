func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums[:])

	count := 1
	final_count := 0

	for i := 0; i<len(nums) - 1; i++ {
		if nums[i] == nums[i + 1] - 1 {
			count += 1
		} else if (nums[i] == nums[i + 1]) {
			continue
		} else {
			if (count > final_count) {
				final_count = count
			}
			count = 1
		}
	}

	if (count > final_count) {
		final_count = count
	}

	return final_count
}

// [2, 3, 4, 4, 5, 10, 20]