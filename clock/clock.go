package clock

import "fmt"

const (
	dayInMin  = 60 * 24
	hourInMin = 60
)

// Clock Represent time of day in hour and minutes only
type Clock struct {
	hour   int
	minute int
}

// New takes hour and minute as int and return a new clock
func New(h, m int) Clock {
	mins := hourInMin*h + m
	mins %= dayInMin

	if mins < 0 {
		mins += dayInMin
	}

	return Clock{mins / hourInMin, mins % hourInMin}
}

func (c Clock) String() string {
	c = New(c.hour, c.minute)
	return fmt.Sprintf("%.2d:%.2d", c.hour, c.minute)
}

// Add takes minute as argument and returns a new clock with minutes added
func (c Clock) Add(minute int) Clock {
	return New(c.hour, c.minute+minute)
}

// Subtract takes minute as argument and returns a new clock with minutes subtracted
func (c Clock) Subtract(minute int) Clock {
	return New(c.hour, c.minute-minute)
}
