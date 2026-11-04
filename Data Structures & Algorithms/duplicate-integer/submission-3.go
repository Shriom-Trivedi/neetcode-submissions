func hasDuplicate(nums []int) bool {
    keys := make(map[int]bool)

	for _, item := range nums {
		if(keys[item]) {
			return true
		} else {
			keys[item] = true
		}
	}
	return false
}
