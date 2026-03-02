package clock

import "fmt"

type Clock struct {
	hours   int
	minutes int
}

func normalize(total int) (int, int) {
	const day = 24 * 60
	total = ((total % day) + day) % day // handle negatives
	return total / 60, total % 60
}

func New(h, m int) Clock {
	h, m = normalize(h*60 + m)
	return Clock{hours: h, minutes: m}
}

func (c Clock) Add(m int) Clock {
	h, min := normalize(c.hours*60 + c.minutes + m)
	return Clock{hours: h, minutes: min}
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}