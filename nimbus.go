package nimbus

import (
	"sort"
	"strings"
)

// a singular portion of the entire word cloud
type nimbus struct {
	Word  string
	Count int
}

type ByCount []nimbus

// sorting functions
func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].Count < a[j].Count }

// golang tutorial #45
func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	for _, word := range strings.Fields(s) {
		wordMap[word] += 1
	}

	return wordMap
}

// trim an unsorted word cloud map
func TrimWordCountMap(unsortedWordCloud map[string]int, n int) map[string]int {
	var nimbusSlice = []nimbus{}
	for k, v := range unsortedWordCloud {
		nimbusSlice = append(nimbusSlice, nimbus{Word: k, Count: v})
	}

	// sort by count
	sort.Sort(sort.Reverse(ByCount(nimbusSlice)))

	// ensure no out-of-bounds error
	if n > len(nimbusSlice) {
		return unsortedWordCloud
	}

	// cut off everything after the first n elements of the slice
	nimbusSlice = nimbusSlice[:n]

	// make a new map that contains the slice data, return it
	m := make(map[string]int)
	for _, nim := range nimbusSlice {
		m[nim.Word] += nim.Count
	}

	return m
}
