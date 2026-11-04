import (
	"slices"
)

func isAnagram(s string, t string) bool {
	str1 := []rune(s)
	str2 := []rune(t)

	slices.Sort(str1)
	slices.Sort(str2)

	if (len(str1) != len(str2)) {
		return false
	}

	for idx, _ := range str1 {
		if(str1[idx] != str2[idx]) {
			return false
		}
	}

	return true
}
