package search

import (
	"sort"

	"github.com/samber/lo"
)

type SearchResult struct {
	Value    string
	Distance int
}

func Search(slice []string, substring string) []string {
	results := search(slice, substring)

	return lo.Map(results, func(result SearchResult, _ int) string {
		return result.Value
	})
}

// Search a slice of strings for a substring and sort by levenshtein distance.
func search(slice []string, substring string) []SearchResult {
	results := lo.Reduce(slice, func(agg []SearchResult, item string, index int) []SearchResult {
		// Skip empty lines
		if len(item) <= 0 {
			return agg
		}

		// Skip lines that already exist in the results
		contains := lo.ContainsBy(agg, func(result SearchResult) bool {
			return result.Value == item
		})
		if contains {
			return agg
		}

		// Skip lines with a levenshtein distance greater than 3
		distance := levenshteinDistance(item, substring)
		if distance > 4 {
			return agg
		}

		agg = append(agg, SearchResult{Value: item, Distance: distance})

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
