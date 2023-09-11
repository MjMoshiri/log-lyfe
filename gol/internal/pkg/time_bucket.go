package pkg

import "time"

// Base time is 1 January 2023, 00:00:00
var baseTime = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

// TimeToBucket converts a time.Time object to its corresponding bucket number.
func TimeToBucket(t time.Time) int64 {
	return int64(t.Sub(baseTime).Hours())
}
