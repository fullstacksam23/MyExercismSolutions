package anagram

import (
    "sort"
    "strings"
)

func Detect(subject string, candidates []string) []string {
    lowerSub := strings.ToLower(subject)
    ogString := sortString(lowerSub)
    var sortedCand string
    res := make([]string, 0)
	for _, cand := range candidates {
        lowerCand := strings.ToLower(cand)
        if lowerCand != lowerSub{
            sortedCand = sortString(lowerCand)
            if sortedCand == ogString {
                res = append(res, cand)
            }       
        }
    }
    return res
}
func sortString(s string) string {
    r := []rune(s)
    sort.Slice(r, func(i, j int) bool {
        return r[i] < r[j]
    })
    return string(r)
}