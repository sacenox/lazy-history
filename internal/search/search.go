package search

import (
	"sort"

	"github.com/samber/lo"
)

type SearchResult struct {
	Value    string
	Distance int
}

// Search a slice of strings for a substring and sort by levenshtein distance.
func Search(slice []string, substring string) []SearchResult {
	results := lo.Reduce(slice, func(agg []SearchResult, item string, index int) []SearchResult {
		distance := levenshteinDistance(item, substring)
		contains := lo.ContainsBy(agg, func(result SearchResult) bool {
			return result.Value == item
		})
		if distance <= 3 && !contains {
			agg = append(agg, SearchResult{Value: item, Distance: distance})
		}
		return agg
	}, []SearchResult{})

	sort.Slice(results, func(i, j int) bool {
		return results[i].Distance < results[j].Distance
	})

	return results
}

func levenshteinDistance(s, t string) int {
	lenS := len(s)
	lenT := len(t)

	if lenS == 0 {
		return lenT
	}
	if lenT == 0 {
		return lenS
	}

	if s[lenS-1] == t[lenT-1] {
		return levenshteinDistance(s[:lenS-1], t[:lenT-1])
	}

	return 1 + min(
		levenshteinDistance(s[:lenS-1], t),
		levenshteinDistance(s, t[:lenT-1]),
		levenshteinDistance(s[:lenS-1], t[:lenT-1]),
	)
}
