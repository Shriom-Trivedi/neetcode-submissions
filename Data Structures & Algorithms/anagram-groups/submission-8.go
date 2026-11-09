import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {

	hash := make(map[string][]string)

	for _, str := range strs {
		s := []rune(str)
		slices.Sort(s)
		sortedStr := string(s)
		
		if _, found := hash[sortedStr]; found {
			hash[sortedStr] = append(hash[sortedStr], str)
		} else {
			hash[sortedStr] = []string{str}
		}
		
	}

	r := make([][]string, 0, len(hash))

	for _, item := range hash {
		r = append(r, item)
	}
	return r
}
