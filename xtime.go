// Package xtime provides utility functions to support the time package.
package xtime

import (
	"errors"
	"regexp"
	"time"
)

var timeRE = regexp.MustCompile(`(([1-9]|1[0-9]|2[0-3]):[0-5][0-9])|(([0-1][0-9]|2[0-3])[0-5][0-9])`)

// ExtractTimes returns a slice of substrings of str that contain
// times without dates in either a 12-hour or 24-hour format.
func ExtractTimes(str string) []string {
	return timeRE.FindAllString(str, -1)
}

// ErrNoMatchingLayout is the error returned by [ParseAs] if no
// provided layout was capable of parsing the provided time.
var ErrNoMatchingLayout = errors.New("provided time unparsable with provided layouts")

// ParseAs attempts to parse value using [time.Parse] with each
// provided layout in turn until one is successful, at which point it
// returns the parsed time and the layout that succeeded. If none
// succeed, [ErrNoMatchingLayout] is returned.
func ParseAs(value string, layouts ...string) (time.Time, string, error) {
	for _, layout := range layouts {
		t, err := time.Parse(layout, value)
		if err == nil {
			return t, layout, nil
		}
	}

	return time.Time{}, "", ErrNoMatchingLayout
}

// ParseInLocationAs is exactly like [ParseAs] but uses
// [time.ParseInLocation] when attempting to parse value.
func ParseInLocationAs(value string, loc *time.Location, layouts ...string) (time.Time, string, error) {
	for _, layout := range layouts {
		t, err := time.ParseInLocation(layout, value, loc)
		if err == nil {
			return t, layout, nil
		}
	}

	return time.Time{}, "", ErrNoMatchingLayout
}
