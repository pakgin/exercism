package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, ch := range log {
		switch ch {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	output := []rune(log)
	for i, ch := range output {
		if ch == oldRune {
			output[i] = newRune
		}
	}
	return string(output)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
