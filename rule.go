package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type ruleIterator func(line int, pattern string, keep int)

func ruleIterate(r io.Reader, it ruleIterator) error {
	i := 0
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		i++
		line := strings.TrimSpace(s.Text())
		// skip empty
		if len(line) == 0 {
			continue
		}
		// skip comment
		if strings.HasPrefix(line, "#") {
			continue
		}
		// remove tailing comment
		split := strings.SplitN(line, "#", 2)
		if len(split) > 1 {
			line = split[0]
		}
		// extract 'pattern' and 'keep'
		split = strings.SplitN(line, ":", 2)
		if len(split) != 2 {
			log.Printf("- line %d: syntax invalid", i)
			continue
		}
		var err error
		var keep int
		if keep, err = strconv.Atoi(strings.TrimSpace(split[1])); err != nil {
			log.Printf("- line %d: 'keep' value invalid", i)
			continue
		}
		pattern := strings.TrimSpace(split[0])
		// invoke it
		log.Printf("- line %d: %s (keep %d days)", i, pattern, keep)
		it(i, pattern, keep)
	}
	return s.Err()
}

func ruleIterateFile(filename string, it ruleIterator) error {
	var err error
	var f *os.File
	if f, err = os.Open(filename); err != nil {
		return err
	}
	defer f.Close()
	return ruleIterate(f, it)
}

type ruleDirIterator func(rulefile string, line int, pattern string, keep int)

func ruleIterateDir(dir string, it ruleDirIterator) error {
	var err error
	var fis []os.FileInfo
	if fis, err = ioutil.ReadDir(dir); err != nil {
		return err
	}
	sort.SliceStable(fis, func(i, j int) bool {
		return fis[i].Name() < fis[j].Name()
	})
	for _, fi := range fis {
		// skip dot files
		if strings.HasPrefix(fi.Name(), ".") {
			continue
		}
		rf := filepath.Join(dir, fi.Name())
		log.Printf("rule: %s", rf)
		if err = ruleIterateFile(rf, func(line int, pattern string, keep int) {
			it(rf, line, pattern, keep)
		}); err != nil {
			log.Printf("- failed to load: %s", err.Error())
			continue
		}
	}
	return nil
}
