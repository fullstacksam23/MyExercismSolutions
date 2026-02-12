package isogram

import "unicode"

func IsIsogram(word string) bool {
	set := make(map[rune]struct{})

    for _, letter := range word {
        if letter != ' ' && letter != '-' {
            lower := unicode.ToLower(letter)
            _, found := set[lower]
            if found {
                return false
            }
            set[lower] = struct{}{}
        }
    }
    return true
}
