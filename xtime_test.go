package xtime_test

import (
	"slices"
	"testing"
	"time"

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

func TestParseAs(t *testing.T) {
	p, l, err := xtime.ParseAs("11:15", time.ANSIC, time.Stamp, "3:04", time.TimeOnly)
	if err != nil {
		t.Fatal(err)
	}
	if l != "3:04" {
		t.Fatal(l)
	}
	if !p.Equal(time.Date(0, 1, 1, 11, 15, 0, 0, time.UTC)) {
		t.Fatal(p)
	}
}
