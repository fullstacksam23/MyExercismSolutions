package reverse

func Reverse(input string) string {
	r := []rune(input)
    n := len(r)
    for i := 0; i < n/2; i++ {
        r[i], r[n-i-1] = r[n-i-1], r[i]
    }
    return string(r)
}
