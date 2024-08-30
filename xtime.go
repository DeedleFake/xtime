// Package xtime provides utility functions to support the time package.
package xtime

import "regexp"

var timeRE = regexp.MustCompile(`(([1-9]|1[0-9]|2[0-3]):[0-5][0-9])|(([0-1][0-9]|2[0-3])[0-5][0-9])`)

// ExtractTimes returns a slice of substrings of str that contain
// times without dates in either a 12-hour or 24-hour format.
func ExtractTimes(str string) []string {
	return timeRE.FindAllString(str, -1)
}
