package letter

// import "time"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	chans := make(chan FreqMap, len(l))
	out := FreqMap{}
	for _, s := range l {
		go func(text string) {
			chans <- Frequency(text)
		}(s)
	}
	for i := 0; i < len(l); i++ {
		for k, v := range <-chans {
			out[k] += v
		}
	}
	return out
}
