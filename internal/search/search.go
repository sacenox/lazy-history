package search

import (
	"sort"
	"strings"

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
	// search for the substring in the slice
	matches := lo.Filter(slice, func(item string, _ int) bool {
		return strings.Contains(item, substring)
	})

	results := lo.Reduce(matches, func(agg []SearchResult, item string, index int) []SearchResult {
		// Skip lines that already exist in the results
		contains := lo.ContainsBy(agg, func(result SearchResult) bool {
			return result.Value == item
		})
		if contains {
			return agg
		}
		distance := levenshteinDistance(item, substring)
		agg = append(agg, SearchResult{Value: item, Distance: distance})

		return agg
	}, []SearchResult{})

	sort.Slice(results, func(i, j int) bool {
		return results[i].Distance < results[j].Distance
	})

	return results
}

func levenshteinDistance(s, t string) int {
	// Create matrix of size (lenS+1)*(lenT+1)
	lenS := len(s)
	lenT := len(t)

	// Initialize the matrix
	d := make([][]int, lenS+1)
	for i := range d {
		d[i] = make([]int, lenT+1)
		d[i][0] = i // Fill first column
	}
	for j := range d[0] {
		d[0][j] = j // Fill first row
	}

	// Fill the matrix
	for i := 1; i <= lenS; i++ {
		for j := 1; j <= lenT; j++ {
			cost := 1
			if s[i-1] == t[j-1] {
				cost = 0
			}

			d[i][j] = min(
				d[i-1][j]+1,      // deletion
				d[i][j-1]+1,      // insertion
				d[i-1][j-1]+cost, // substitution
			)
		}
	}

	return d[lenS][lenT]
}
