package handling

import (
	"strings"
)

func findKeywords(in, keywords string) (words string, count int) {
	keys := strings.Split(keywords, ",")

	for _, k := range keys {
		k = strings.TrimSpace(k)
		index := strings.Index(in, k)
		if index >= 0 {
			count++
			if len(words) == 0 {
				words = k
			} else {
				words += ", " + k
			}
		}
	}

	return
}

func keywords(data string, controlType string, listWord string) (status int, list string, count int) {
	status = 0

	if data == "" {
		status = 2
		return
	}

	list, count = findKeywords(data, listWord)

	if count == 0 && controlType == "out" || count > 0 && controlType == "in" {
		status = 1
	}

	return
}

func system(data string, listWord string) (status int, list string, count int) {
	status = 0

	if data == "" {
		status = 2
		return
	}

	list, count = findKeywords(data, listWord)

	if count > 0 {
		status = 1
	}

	return
}
