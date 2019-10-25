package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Rule struct {
	Files []string `yaml:"files"`
	Keep  int      `yaml:"keep"`
}

func (r Rule) Validate() error {
	if len(r.Files) == 0 {
		return errors.New("missing 'files' field")
	}
	if r.Keep <= 0 {
		return errors.New("missing 'keep' field")
	}
	return nil
}

func (r Rule) Glob() (out []string, err error) {
	mOut := map[string]bool{}
	var matches []string
	for _, item := range r.Files {
		if matches, err = filepath.Glob(item); err != nil {
			return
		}
		for _, m := range matches {
			mOut[m] = true
		}
	}
	for k := range mOut {
		out = append(out, k)
	}
	return
}

func (r Rule) Check(filename string, now time.Time) (expired bool, ok bool) {
	var date time.Time
	if date, ok = extractDateFromFilename(filename); !ok {
		return
	}
	expired = int(now.Sub(date)/time.Hour*24) > r.Keep
	return
}

type RuleBook struct {
	File  string `yaml:"-"`
	Rules []Rule `yaml:"rules"`
}

func LoadRuleBooks(dir string) ([]RuleBook, error) {
	// file infos
	var err error
	var fis []os.FileInfo
	if fis, err = ioutil.ReadDir(dir); err != nil {
		return nil, err
	}
	var rbs []RuleBook
	// load rule books
outerLoop:
	for _, fi := range fis {
		// check file extension
		if ext := strings.ToLower(filepath.Ext(fi.Name())); ext != ".yaml" && ext != ".yml" {
			continue
		}
		// load rulebook
		fp := filepath.Join(dir, fi.Name())
		log.Printf("loading: %s", fp)
		var buf []byte
		if buf, err = ioutil.ReadFile(fp); err != nil {
			log.Printf("failed: %s", err.Error())
			continue
		}
		var rb RuleBook
		if err = yaml.Unmarshal(buf, &rb); err != nil {
			log.Printf("failed: %s", err.Error())
			continue
		}
		if len(rb.Rules) == 0 {
			log.Println("failed: no rules")
			continue
		}
		for i, r := range rb.Rules {
			if err = r.Validate(); err != nil {
				log.Printf("failed: rule %d invalid: %s", i, err.Error())
				continue outerLoop
			}
		}
		rb.File = fp
		log.Printf("succeeded: %d rules loaded", len(rb.Rules))
		// append to output
		rbs = append(rbs, rb)
	}
	return rbs, nil
}
