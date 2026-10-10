// Package gigasecond gives out the exact time 
// one gigasecond from the given time.
package gigasecond

// import path for the time package from the standard library.
import "time"

// AddGigasecond calculates the time one gigasecond
// from the given time.
func AddGigasecond(t time.Time) time.Time {
	result := t.Add(1000000000 * time.Second)
	return result
}
