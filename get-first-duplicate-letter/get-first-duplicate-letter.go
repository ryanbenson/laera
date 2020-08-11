package getfirstduplicateletter

import "strings"

func getFirstDuplicateLetter_Map(word string) string {
	letters := strings.Split(word, "")

	letterMap := make(map[string]int)

	for _, letter := range letters {
		_, ok := letterMap[letter]
		if ok {
			return letter
		}
		letterMap[letter] = 1
	}

	// never found one
	return ""
}