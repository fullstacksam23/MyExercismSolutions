package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	result := make(map[string]int)
    for key, val := range in {
        for _, letter := range val {
            result[strings.ToLower(letter)] = key
        }
    }
    return result
}
