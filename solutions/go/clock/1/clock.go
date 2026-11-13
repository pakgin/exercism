package clock

import (
	"fmt"
	// "math"
)

// Define the Clock type here.
type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	for m < 0 || m >= 60 {
		if m < 0 {
			m += 60
			h--
		} else {
			h++
			m -= 60
		}
	}
	for h < 0 || h >= 24 {
		if h < 0 {
			h += 24
		} else {
			h -= 24
		}
	}
	return Clock{hour: h, minute: m}
}

func (c Clock) Add(m int) Clock {
	new_h, new_m := c.hour, c.minute+m
	if new_m >= 60 {
		new_h += new_m / 60
		new_m %= 60
	}
	if new_h >= 24 {
		new_h %= 24
	}
	return New(new_h, new_m)
}

func (c Clock) Subtract(m int) Clock {
	new_h, new_m := c.hour, c.minute-m
	if new_m < 0 {
		new_h -= (-1 * new_m) / 60
		new_m %= 60
	}
	if new_h < 0 {
		new_h %= 24
	}
	return New(new_h, new_m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
