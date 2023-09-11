package eventer

import (
	"testing"
	"time"
)

func TestTimeToBucket(t *testing.T) {
	tests := []struct {
		input    time.Time
		expected int64
	}{
		{baseTime, 0},                                       // Exactly at base time
		{baseTime.Add(1 * time.Hour), 1},                    // 1 hour after base time
		{baseTime.Add(25 * time.Hour), 25},                  // 1 day and 1 hour after base time
		{baseTime.Add(-1 * time.Hour), -1},                  // 1 hour before base time
		{baseTime.AddDate(0, 0, 1), 24},                     // 1 day after base time
		{baseTime.AddDate(0, 0, 1).Add(12 * time.Hour), 36}, // 1 day and 12 hours after base time
	}

	for _, tt := range tests {
		result := TimeToBucket(tt.input)
		if result != tt.expected {
			t.Errorf("For time %v, expected bucket %d but got %d", tt.input, tt.expected, result)
		}
	}
}
