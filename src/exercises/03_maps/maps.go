package maps

import "strings"

func WordCount(s string) map[string]int {
	list := make(map[string]int)
	for _,w := range strings.Fields(s) {
		list[w]++
	}
	return list
}
