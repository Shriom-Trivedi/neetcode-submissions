func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1

	max_area := 0

	for left < right {
		area := (right - left) * min(heights[right], heights[left]) // length * width

		if area > max_area {
			max_area = area
		}

		if heights[right] > heights[left] {
			left++
		} else {
			right--
		}


	}

	return max_area
}
