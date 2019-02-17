package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int
type CFreqMap sync.Map

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency Uses a channel to get each frequency map
// Then count up the final map from the other maps
func ConcurrentFrequency(list []string) FreqMap {
	c := make(chan FreqMap, len(list))
	for _, s := range list {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	m := make(FreqMap)
	for range list {
		for k, v := range <-c {
			m[k] += v
		}
	}
	return m
}
