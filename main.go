package main

import (
	"flag"
	"log"
	"os"
	"time"
)

var (
	optRulesDir string
	optDryRun   bool
	optNow      int64
)

func exit(err *error) {
	if *err != nil {
		log.Printf("exited with error: %s", (*err).Error())
		os.Exit(1)
	}
}

func main() {
	var err error
	defer exit(&err)

	// flags
	flag.StringVar(&optRulesDir, "d", "/etc/delog.d", "rule books directory")
	flag.BoolVar(&optDryRun, "dry", false, "dry run, not actually delete files")
	flag.Int64Var(&optNow, "now", 0, "set the 'now' date, for test only")
	flag.Parse()

	// today
	var today time.Time
	if optNow != 0 {
		today = beginningOfDay(time.Unix(optNow, 0))
	} else {
		today = beginningOfDay(time.Now())
	}
	log.Printf("today: %s", today.Format(time.RFC3339))

	// rule books
	var rbs []RuleBook
	if rbs, err = LoadRuleBooks(optRulesDir); err != nil {
		return
	}

	// run rule books
	for _, rb := range rbs {
		log.Printf("running: %s", rb.File)
		for i, r := range rb.Rules {
			log.Printf("rule: %d", i)
			var files []string
			if files, err = r.Glob(); err != nil {
				log.Printf("failed to glob files: %s", err.Error())
				continue
			}
			for _, file := range files {
				expired, ok := r.Check(file, today)
				log.Printf("check file: %s, matched = %v, expired = %v", file, ok, expired)
				if ok && expired && !optDryRun {
					log.Printf("removed: %s", file)
					_ = os.Remove(file)
				}
			}
		}
	}
}
