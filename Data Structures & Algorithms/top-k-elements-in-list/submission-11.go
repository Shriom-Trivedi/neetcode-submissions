func topKFrequent(nums []int, k int) []int {
	// Solving this with the help of bucket sort.
	hash := make(map[int]int)

	for _, n := range nums {
		hash[n]++
	}

	buckets := make([][]int, len(nums)+1)

	for num, count := range hash {
		buckets[count] = append(buckets[count], num)
	}

	results := make([]int, 0, k)

	for i := len(buckets) - 1; i >= 0; i-- {
		for _, num := range buckets[i] {
			results = append(results, num)
		}

		if len(results) == k {
			return results
		}
	}
	return results
}
