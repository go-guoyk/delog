package main

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestExtractDateFromFilename(t *testing.T) {
	date, ok := extractDateFromFilename("project12019-10-10.log")
	require.True(t, ok)
	require.Equal(t, time.Date(2019, 10, 10, 0, 0, 0, 0, time.Local), date)
	date, ok = extractDateFromFilename("project12019.10-10.log")
	require.True(t, ok)
	require.Equal(t, time.Date(2019, 10, 10, 0, 0, 0, 0, time.Local), date)
	date, ok = extractDateFromFilename("project12019.10.10.log")
	require.True(t, ok)
	require.Equal(t, time.Date(2019, 10, 10, 0, 0, 0, 0, time.Local), date)
	date, ok = extractDateFromFilename("project_20191010.log")
	require.True(t, ok)
	require.Equal(t, time.Date(2019, 10, 10, 0, 0, 0, 0, time.Local), date)
}
