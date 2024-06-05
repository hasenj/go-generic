package generic

import (
	"fmt"
	"time"
)

// TimeStamp3Format takes a time and formats it as yyyy-mmdd-hhmm with the
// option to round to a multiple of minutes.
// For example, if you pass "2024-05-28 14:33" and round to 10 minutes, you get
// "2024-0528-1430"
func TimeStamp3Format(now time.Time, roundToMinutes int) string {
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()

	minute = (minute / roundToMinutes) * roundToMinutes

	return fmt.Sprintf("%d-%02d%02d-%02d%02d", year, month, day, hour, minute)
}
