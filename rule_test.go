package main

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

const testRule = `
# this is a test
  /log/test1 : 9 #
 
/log/test2 : 8 #

  ## 
# 
/log/test3:7 #`

func TestRuleIterate(t *testing.T) {
	var lines []int
	var patterns []string
	var keeps []int
	err := ruleIterate(strings.NewReader(testRule), func(line int, pattern string, keep int) {
		lines = append(lines, line)
		patterns = append(patterns, pattern)
		keeps = append(keeps, keep)
	})
	require.NoError(t, err)
	require.Equal(t, []int{3, 5, 9}, lines)
	require.Equal(t, []string{"/log/test1", "/log/test2", "/log/test3"}, patterns)
	require.Equal(t, []int{9, 8, 7}, keeps)
}

func TestRuleDirIterate(t *testing.T) {
	var rulefiles []string
	var lines []int
	var patterns []string
	var keeps []int
	err := ruleIterateDir("testdata/rules", func(rulefile string, line int, pattern string, keep int) {
		rulefiles = append(rulefiles, rulefile)
		lines = append(lines, line)
		patterns = append(patterns, pattern)
		keeps = append(keeps, keep)
	})
	require.NoError(t, err)
	require.Equal(t, []string{"testfile/rules/test", "testfile/rules/test2"}, rulefiles)
	require.Equal(t, []int{3, 4}, lines)
	require.Equal(t, []string{"testdata/logs/**/*.log", "testdata/logs/*.log"}, patterns)
	require.Equal(t, []int{2, 4}, keeps)
}
