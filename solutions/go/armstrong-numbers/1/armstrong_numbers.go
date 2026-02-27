package armstrong

import (
    "math"
    )

func IsNumber(n int) bool {
	length := 0
    copy := n
    for copy > 0 {
        copy /= 10
        length++
    }
    copy = n
    sum := 0
    for copy > 0 {
		sum += int(math.Pow(float64(copy%10), float64(length)))
        copy /= 10
    }
    return sum == n
}
