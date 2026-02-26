// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
    "strings"
    "unicode"
    )
// Abbreviate takes the first letter of each word and returns them concatenated and converted to upper case
func Abbreviate(s string) string {
    var newString string
    newString = strings.ReplaceAll(s, "-", " ")
    newString = strings.ReplaceAll(newString, "_", " ")
    
	words := strings.Split(newString, " ")
    bytes := make([]byte, 0, len(words))
    for _, word := range words {
        if len(word) > 0{
            bytes = append(bytes, byte(unicode.ToUpper(rune(word[0]))))
        }
    }
	return string(bytes)
}
