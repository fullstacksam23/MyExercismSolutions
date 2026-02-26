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
	words := strings.Split(s, " ")
    bytes := make([]byte, 0, len(words))
    for _, word := range words {
        if strings.Contains(word, "-"){
        	subWords := strings.Split(word, "-")
            for _, w := range subWords {
                if len(w) > 0{
                	bytes = append(bytes, byte(unicode.ToUpper(rune(w[0]))))
                }
            }
        }else if strings.Contains(word, "_"){
            subWords := strings.Split(word, "_")
            for _, w := range subWords {
                if len(w) > 0{
                	bytes = append(bytes, byte(unicode.ToUpper(rune(w[0]))))
                }
            }
        }else{
            if len(word) > 0{
                bytes = append(bytes, byte(unicode.ToUpper(rune(word[0]))))
            }
            
        }
    }
	return string(bytes)
}
