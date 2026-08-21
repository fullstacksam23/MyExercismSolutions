package eliudseggs

func EggCount(displayValue int) int {
    ones := 0
	for displayValue > 0 {
		if displayValue&1 == 1 {
            ones+=1
        }
       	displayValue >>= 1
    }
    return ones
}
