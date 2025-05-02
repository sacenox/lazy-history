package lib

import (
	"slices"
	"sort"
)

func FindOne(history []string, query string) (string, int) {
	// search using a binary search
	i := sort.Search(len(history), func(i int) bool {
		Debugf("history[%d]: %q", i, history[i])
		return history[i] >= query
	})
	if i < len(history) {
		Debugf("found %q at index %d\n", query, i)
		return history[i], i
	} else {
		Debugf("%q not found\n", query)
		return "", -1
	}
}

func Search(history []string, query string) []string {
	results := []string{}
	// sort the array
	slices.Sort(history)

	// Keep searching for matches until we don't find any more
	for {
		result, i := FindOne(history, query)
		if result == "" {
			break
		}

		results = append(results, result)
		history = slices.Delete(history, i, i+1)
	}

	return results
}
