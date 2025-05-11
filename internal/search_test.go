package search

import (
	"reflect"
	"testing"
)

// Test the search function.
func TestSearch(t *testing.T) {
	tests := []struct {
		slice     []string
		substring string
		want      []SearchResult
	}{
		{
			slice:     []string{"apple", "banana", "cherry", "date"},
			substring: "ban",
			want:      []SearchResult{{"banana", 3}, {"date", 3}},
		},
		{
			slice:     []string{"apple", "banana", "cherry", "date"},
			substring: "banana",
			want:      []SearchResult{{"banana", 0}},
		},
	}

	for _, test := range tests {
		if got := Search(test.slice, test.substring); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Search(%v, %q) = %v; want %v", test.slice, test.substring, got, test.want)
		}
	}
}

// Test the levenshtein distance function.
func TestLevenshteinDistance(t *testing.T) {
	tests := []struct {
		s, t string
		want int
	}{
		{"ban", "apple", 5},
		{"ban", "banana", 3},
		{"ban", "cherry", 6},
		{"ban", "date", 3},
		{"banana", "banana", 0},
		{"banana", "apple", 5},
		{"banana", "cherry", 6},
		{"banana", "date", 5},
	}

	for _, test := range tests {
		if got := levenshteinDistance(test.s, test.t); got != test.want {
			t.Errorf("levenshteinDistance(%q, %q) = %d; want %d", test.s, test.t, got, test.want)
		}
	}
}
