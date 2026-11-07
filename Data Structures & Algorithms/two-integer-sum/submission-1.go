func twoSum(nums []int, target int) []int {
    keys := make(map[int]int)

	// {3: 0, }

	for i, item := range nums {
		comp := target - item
		if idx, found := keys[comp]; found {
			return []int{idx, i}
		}
		keys[item] = i
	}
	return nil
}
