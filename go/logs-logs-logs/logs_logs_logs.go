package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, v := range log {
		if v == '❗' {
			return "recommendation"
		}

		if v == '🔍' {
			return "search"
		}

		if v == '☀' {
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog string
	for _, v := range log {
		if v == oldRune {
			newLog += string(newRune)
		} else {
			newLog += string(v)
		}
	}

	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	// a little bit slower
	//return len([]rune(log)) <= limit
	// faster
	return utf8.RuneCountInString(log) <= limit
}
