package utils

import (
	"log"
	"time"
)

var ViennaLocation *time.Location

func init() {
	vienna, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		log.Fatalln(err)
	}
	ViennaLocation = vienna
}

func WeekStart(t time.Time) time.Time {

	t = Bod(t)

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	return t
}

func WeekStartAndEnd(t time.Time) (time.Time, time.Time) {
	t = WeekStart(t)
	return t, t.Add(time.Hour * 24 * 6).Add(-time.Nanosecond)
}

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}
