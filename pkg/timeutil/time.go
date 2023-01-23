package timeutil

import (
	"strconv"
	"time"
)

// TryParseUnixTime returns a time pointer if it can parse the value, otherwise returns a  nil pointer.
func TryParseUnixTime(s string) *time.Time {

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil
	}

	if i == 0 {
		return nil
	}

	t := time.Unix(i, 0)
	return &t
}

// FormatTimeP calls FormatTime with the given time t or returns the empty
// string if t is null.
func FormatTimeP(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(time.RFC3339)
}
