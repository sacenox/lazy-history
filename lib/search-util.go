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
	// note: am using an explicit sort function because unknown types in actual array, and not using a stable sort function because don't care that items that are equal to one another remain in the same order.
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
