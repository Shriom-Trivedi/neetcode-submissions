import "slices"

func productExceptSelf(nums []int) []int {
	pre := make([]int, len(nums))

	count_pre := 1
	for idx, n := range nums {
		pre[idx] = count_pre
		count_pre *= n
	}

	post := make([]int, len(nums))

	count_post := 1
	for idx, n := range slices.Backward(nums) {
		post[idx] = count_post
		count_post *= n
	}

	result := make([]int, len(nums))
	for idx, n := range pre {
		result[idx] = n * post[idx]
	}

	return result
}
