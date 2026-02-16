package diffsquares

import "math"

func SquareOfSum(n int) int {
	s := n*(n+1)/2
    return s*s
}

func SumOfSquares(n int) int {
	return (n*(n+1)*(2*n+1))/6
}

func Difference(n int) int {
	return int(math.Abs(float64(SquareOfSum(n)-SumOfSquares(n))))
}
