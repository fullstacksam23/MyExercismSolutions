package pangram

import "strings"
func IsPangram(input string) bool {
	input = strings.ToLower(input)
    var alpha [26]bool
    for _, r := range input {
		if r >= 'a' && r <= 'z' {
			alpha[r-'a'] = true
		}
    }
    for i := 0; i < 26; i++ {
        if !alpha[i] {
            return false
        }
    }
    return true
}
