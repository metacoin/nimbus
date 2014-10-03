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

// trim a map that is a representation of a word cloud
func TrimWordCountMap(untrimmedWordCloud map[string]int, n int) map[string]int {

	// create a slice of nimbus
	var nimbusSlice = []nimbus{}

	// add each nimbus to the slice
	for k, v := range untrimmedWordCloud {
		nimbusSlice = append(nimbusSlice, nimbus{Word: k, Count: v})
	}

	// do some length calculations
	nsLength := len(nimbusSlice)
	newLength := nsLength - n

	// ensure no out-of-bounds error
	if newLength <= 0 {
		return untrimmedWordCloud
	}

	// sort by count
	sort.Sort(ByCount(nimbusSlice))

	// cut off everything after the first n elements of the slice
	nimbusSlice = nimbusSlice[newLength:]

	// make a new map that contains the slice data, return it
	m := make(map[string]int)
	for _, nim := range nimbusSlice {
		m[nim.Word] += nim.Count
	}

	return m
}
