func isPalindrome(s string) bool {

	s = cleanString(s)
	strArr := []rune(s)

	left := 0
	right := len(strArr) - 1

	for left <= right {
		if(strings.ToLower(string(strArr[left])) != strings.ToLower(string(strArr[right]))) {
			return false
		}
		left++
		right--
	}

	return true
}

func cleanString(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	return re.ReplaceAllString(s, "")
}
