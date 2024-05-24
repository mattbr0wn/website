package reading_time

import (
	"fmt"
	"math"
	"strings"
)

func TimeToRead(text string) string {
	wordCount := len(strings.Fields(text))
	wordsPerMinute := 250

	// Calculate total time in minutes as a float.
	totalMinutes := float64(wordCount) / float64(wordsPerMinute)

	// Round using half up rounding
	roundedMinutes := int(math.Round(totalMinutes))

	// Determine the output based on rounded minutes.
	if roundedMinutes == 0 {
		// Calculate seconds when less than half a minute.
		seconds := int(math.Round(float64(wordCount%wordsPerMinute) * 60.0 / float64(wordsPerMinute)))
		return fmt.Sprintf("%d seconds", seconds)
	} else if roundedMinutes == 1 {
		return "1 minute"
	} else {
		return fmt.Sprintf("%d minutes", roundedMinutes)
	}
}
