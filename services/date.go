package services

import (
	"fmt"
	"strings"
	"time"
)

func ParseDate(dateStr string) (int64, error) {
	dateStr = strings.Split(dateStr, " (")[0]

	layouts := []string{
		"Mon Jan 02 2006 15:04:05 GMT-0700",
		"2006-01-02T15:04:05.999999999Z07:00",
		"2006-01-02T15:04:05Z",
		"2006-01-02 15:04:05",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, dateStr)
		if err == nil {
			return t.Unix(), nil
		}
	}

	return time.Time{}.Unix(), fmt.Errorf("invalid date format: %s", dateStr)
}
