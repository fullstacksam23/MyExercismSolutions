package scrabble

import "unicode"

//doing this slightly inefficently to save the time of typing out 26 alphabets and their corresponding values
var letterVal = make(map[rune]int)
type combo struct{
    letters string
    value int
}
func init(){
    combos := []combo{
    {"AEIOULNRST", 1},
    		{"DG", 2},
    		{"BCMP", 3},
    		{"FHVWY", 4},
    		{"K", 5},
    		{"JX", 8},
    		{"QZ", 10},
    }
    //populate the map
    for _, c := range combos {
        for _, letter := range c.letters {
            letterVal[letter] = c.value
        }
    } 
}

func Score(word string) int {
    score := 0
	for _, char := range word {
        score += letterVal[unicode.ToUpper(char)]
    }
    return score
}
