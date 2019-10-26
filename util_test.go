package main

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDataFromFilename(t *testing.T) {
	date, ok := dateFromFilename("project12019-10-10.log")
	require.True(t, ok)
	require.Equal(t, time.Date(2019, 10, 10, 0, 0, 0, 0, time.Local), date)
	date, ok = dateFromFilename("project12019.10-10.log")
	require.True(t, ok)
	require.Equal(t, time.Date(2019, 10, 10, 0, 0, 0, 0, time.Local), date)
	date, ok = dateFromFilename("project12019.10.10.log")
	require.True(t, ok)
	require.Equal(t, time.Date(2019, 10, 10, 0, 0, 0, 0, time.Local), date)
	date, ok = dateFromFilename("project12019.10.11.log")
	require.True(t, ok)
	require.Equal(t, time.Date(2019, 10, 11, 0, 0, 0, 0, time.Local), date)
	date, ok = dateFromFilename("project_20191018.log")
	require.True(t, ok)
	require.Equal(t, time.Date(2019, 10, 18, 0, 0, 0, 0, time.Local), date)
}

func TestBeginningOfDay(t *testing.T) {
	require.Equal(
		t,
		time.Date(2019, 10, 10, 0, 0, 0, 0, time.Local),
		dateMidnight(time.Date(2019, 10, 10, 2, 3, 4, 5, time.Local)))
}
