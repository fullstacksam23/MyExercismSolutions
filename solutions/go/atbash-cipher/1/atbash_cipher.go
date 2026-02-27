package atbash

import (
    "unicode/utf8"
    "strings"
)
func Atbash(s string) string {
    result := make([]rune, 0, utf8.RuneCountInString(s))
    s = strings.ToLower(s)
    c := 0
	for _, r := range s {
        if 'a' <= r && r <= 'z' {
            if c % 5 == 0 && c != 0 {
            	result = append(result, ' ')
        	}
            result = append(result, 'a'+25-(r-'a'))
            c+=1
        }else if '0' <= r && r <= '9' {
            if c % 5 == 0 && c != 0{
            	result = append(result, ' ')
        	}
            result = append(result, r)
            c+=1
        }
    }
    return string(result)
}
