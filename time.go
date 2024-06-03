package generic

import "fmt"
import "time"

func TimeStamp3Format(now time.Time, roundToMinutes int) string {
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()

	minute = (minute / roundToMinutes) * roundToMinutes

	return fmt.Sprintf("%d-%02d%02d-%02d%02d", year, month, day, hour, minute)
}
