package main

import (
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

var (
	// DatePattern, supports "2019-10-10", "2019.10.10", or "20191010"
	DatePattern = regexp.MustCompile(`(\d{4})[._-]?(\d{2})[._-]?(\d{2})`)
)

func extractDateFromFilename(filename string) (date time.Time, ok bool) {
	matches := DatePattern.FindStringSubmatch(filepath.Base(filename))
	if len(matches) == 0 {
		return
	}
	year, _ := strconv.Atoi(matches[1])
	month, _ := strconv.Atoi(matches[2])
	day, _ := strconv.Atoi(matches[2])
	date = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	ok = true
	return
}

func beginningOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}
