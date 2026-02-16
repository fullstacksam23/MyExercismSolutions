package luhn

func Valid(id string) bool {
	sum := 0
	double := false
	digits := 0

	for i := len(id) - 1; i >= 0; i-- {
		ch := id[i]

		if ch == ' ' {
			continue
		}
		if ch < '0' || ch > '9' {
			return false
		}

		num := int(ch - '0')
		digits++

		if double {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}

		sum += num
		double = !double
	}

	return digits > 1 && sum%10 == 0
}
