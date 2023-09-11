package eventer

import "time"

// baseTime represents the starting point of 1 January 2023, 00:00:00 UTC.
var baseTime = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

// TimeToBucket calculates the bucket number for a given time by determining the hour difference from the baseTime.
func TimeToBucket(t time.Time) int64 {
	return int64(t.Sub(baseTime).Hours())
}
