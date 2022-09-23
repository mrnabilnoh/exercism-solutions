package clock

import (
	"fmt"
)

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

const HOUR int = 24
const MINUTE int = 60

func New(h, m int) Clock {
	// minute
	if m < 0 {
		if m > -MINUTE {
			h--
			m = MINUTE + m
		} else {
			extra_hour := m / MINUTE
			remainder_minute := m % MINUTE
			if remainder_minute != 0 {
				extra_hour--
				m = MINUTE + remainder_minute
			} else {
				m = remainder_minute
			}
			h = h + extra_hour
		}
	} else if m >= MINUTE {
		h = h + (m / MINUTE)
		m = m % MINUTE
	}

	// hour
	if h < 0 {
		h = HOUR + (h % HOUR)
	} else if h > HOUR {
		h = h % HOUR
	} else if h == 0 {
		h = HOUR
	}

	return Clock{h, m}
}

func (c Clock) Add(m int) Clock {
	// only handle positive value
	if m > 0 {
		h := c.h
		m = c.m + m

		if m >= MINUTE {
			extra_hour := m / MINUTE
			remainder_minute := m % MINUTE

			h = (h + extra_hour) % HOUR
			m = remainder_minute
		}

		// handle if zero
		if h == 0 {
			h = HOUR
		}

		c.h = h
		c.m = m
	}

	return c
}

func (c Clock) Subtract(m int) Clock {
	// only handle positive value
	if m > 0 {
		h := c.h
		m = c.m - m

		if m < 0 {
			if m > -MINUTE {
				h--
				m = MINUTE + m
			} else {
				extra_hour := m / MINUTE
				remainder_minute := m % MINUTE
				if remainder_minute != 0 {
					extra_hour--
					m = MINUTE + remainder_minute
				} else {
					m = remainder_minute
				}
				h = h + extra_hour
			}

			// hour
			if h < 0 {
				h = HOUR + (h % HOUR)
			} else if h > HOUR {
				h = h % HOUR
			} else if h == 0 {
				h = HOUR
			}
		}

		c.h = h
		c.m = m
	}
	return c
}

func (c Clock) String() string {
	h := c.h
	m := c.m

	// check if h equal to 24, change to 0
	if h == HOUR {
		h = 0
	}

	return fmt.Sprintf("%02d:%02d", h, m)
}
