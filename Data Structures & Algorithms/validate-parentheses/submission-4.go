func isValid(s string) bool {
    config := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	stack := make([]string, 0)
	strArr := []rune(s)

	stack = append(stack, "NIL")

	for _, char := range strArr {
		fmt.Println(stack)

		if string(char) == config[stack[len(stack) - 1]] {
			stack = stack[:len(stack) - 1]
		} else {
			stack = append(stack, string(char))
		}
	}

	fmt.Println(stack)
	return len(stack) == 1 && stack[0] == "NIL"
}
