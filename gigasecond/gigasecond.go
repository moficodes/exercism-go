package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// 10^9
const Gigasecond time.Duration = 1e9

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * Gigasecond)
}
