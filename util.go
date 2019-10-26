package main

import (
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

var (
	filenameDatePattern = regexp.MustCompile(`(\d{4})[._-]?(\d{2})[._-]?(\d{2})`)
)

func dateFromFilename(filename string) (date time.Time, ok bool) {
	matches := filenameDatePattern.FindStringSubmatch(filepath.Base(filename))
	if len(matches) == 0 {
		return
	}
	year, _ := strconv.Atoi(matches[1])
	month, _ := strconv.Atoi(matches[2])
	day, _ := strconv.Atoi(matches[3])
	date = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	ok = true
	return
}

func dateMidnight(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}
