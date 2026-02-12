package hamming

import "errors"

func Distance(a, b string) (int, error) {
	n1 := len(a)
    n2 := len(b)
    if n1 != n2 {
        return 0, errors.New("hamming distance is defined only for strings of equal length")
    }
    var cnt = 0
	for i := 0; i < n1; i++ {
        if a[i] != b[i]{
            cnt++
        }
    }
    return cnt, nil
}
