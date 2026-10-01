package clock

import "fmt"

// Clock represents a time of day without a date.
type Clock struct {
	minutes int
}

// New creates a new Clock given hours and minutes.
func New(hour, minute int) Clock {
	total := (hour*60 + minute) % 1440
	if total < 0 {
		total += 1440
	}
	return Clock{minutes: total}
}

// Add adds a number of minutes to the clock.
func (c Clock) Add(minutes int) Clock {
	return New(0, c.minutes+minutes)
}

// Subtract subtracts a number of minutes from the clock.
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minutes-minutes)
}

// String formats the clock as HH:MM.
func (c Clock) String() string {
	h := c.minutes / 60
	m := c.minutes % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}
