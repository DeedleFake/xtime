package xtime_test

import (
	"slices"
	"testing"

	"deedles.dev/xtime"
)

func TestExtractTimes(t *testing.T) {
	const str = `11:15 is a time of day that comes before 1532, as well as 13:57`
	actual := []string{"11:15", "1532", "13:57"}
	times := xtime.ExtractTimes(str)
	if !slices.Equal(times, actual) {
		t.Fatal(actual)
	}
}
