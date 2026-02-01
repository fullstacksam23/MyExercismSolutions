package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, x := range log {
        if x == '❗' {
            return "recommendation"
        } else if x == '🔍' {
            return "search"
        } else if x == '☀' {
            return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var res []rune
    for _, r := range log {
        if r == oldRune{
            res = append(res, newRune)
        }else {
            res = append(res, r)
        }
    }
    return string(res)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
