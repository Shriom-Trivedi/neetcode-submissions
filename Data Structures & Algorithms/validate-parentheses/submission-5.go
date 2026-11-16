func isValid(s string) bool {
    config := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	stack := make([]string, 0)
	strArr := []rune(s)

	for _, char := range strArr {
		fmt.Println(stack)

		if len(stack) > 0 && string(char) == config[stack[len(stack) - 1]] {
			stack = stack[:len(stack) - 1]
		} else {
			stack = append(stack, string(char))
		}
	}

	fmt.Println(stack)
	return len(stack) == 0
}
