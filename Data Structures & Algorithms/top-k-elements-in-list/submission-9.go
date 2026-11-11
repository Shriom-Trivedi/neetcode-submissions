func topKFrequent(nums []int, k int) []int {
	hash := make(map[int]int)

	for _, n := range nums {
		if _, found := hash[n]; found {
			hash[n] += 1
		} else {
			hash[n] = 1
		}
	}

	result := make([]int, 0, len(hash))

	for key, _ := range hash {
		result = append(result, key)
	}

	sort.Slice(result, func(x, y int) bool {
		return hash[ result[x]] > hash[result[y] ]
	})

	return result[:k]
}
