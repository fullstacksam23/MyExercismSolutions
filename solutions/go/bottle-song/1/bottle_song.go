package bottlesong

import "strings"
func Recite(startBottles, takeDown int) []string {
	var res []string
    for i := 0; i < takeDown; i++ {
		res = append(res, Capitalize(numToWord(startBottles))+" green "+bottlePlurality(startBottles)+" hanging on the wall,")
        res = append(res, Capitalize(numToWord(startBottles))+" green "+bottlePlurality(startBottles)+" hanging on the wall,")
		res = append(res, "And if one green bottle should accidentally fall,")
        startBottles--
		res = append(res, "There'll be "+numToWord(startBottles)+" green "+bottlePlurality(startBottles)+" hanging on the wall.")
        if i != takeDown-1{
            res = append(res, "")
        }
    }
    return res
}
func bottlePlurality(num int) string{
    if num > 1 || num == 0{
        return "bottles"
    }else{
        return "bottle"
    }
}
func numToWord(num int) string {
    switch num {
        case 0:
        	return "no"
        case 1 : 
        	return "one"
    	case 2:
        	return "two"
        case 3:
        	return "three"
    	case 4:
        	return "four"
    	case 5:
        	return "five"
    	case 6:
        	return "six"
    	case 7:
        	return "seven"
    	case 8:
        	return "eight"
    	case 9:
        	return "nine"
        case 10:
        	return "ten"
    }
    return ""
}
func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(word[:1]) + word[1:]
}