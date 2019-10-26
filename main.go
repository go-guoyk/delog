package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	optRulesDir string
	optNoDelete bool
	optBaseDate string
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
	flag.StringVar(&optRulesDir, "d", "/etc/logdel.d", "配置文件目录")
	flag.BoolVar(&optNoDelete, "no-delete", false, "不执行删除操作")
	flag.StringVar(&optBaseDate, "base-date", "", "设置用于计算日志文件过期的基准时间 (格式 YYYY-MM-DD)，默认为当前时间")
	flag.Parse()

	// base date
	var baseDate time.Time
	if len(optBaseDate) != 0 {
		if baseDate, err = time.Parse("2006-01-02", optBaseDate); err != nil {
			return
		}
		baseDate = dateMidnight(baseDate)
	} else {
		baseDate = dateMidnight(time.Now())
	}
	log.Printf("date: %s", baseDate.Format(time.RFC3339))

	// iterate
	if err = ruleIterateDir(optRulesDir, func(rulefile string, line int, pattern string, keep int) {
		var err error
		var files []string
		if files, err = filepath.Glob(pattern); err != nil {
			log.Printf("- line: %d: 'pattern' value invalid", line)
			return
		}
		for _, file := range files {
			var date time.Time
			var ok bool
			if date, ok = dateFromFilename(file); !ok {
				log.Printf("-- unknown: %s", file)
				continue
			}
			if time.Duration(keep)*time.Hour*24 >= baseDate.Sub(date) {
				log.Printf("-- ok: %s", file)
				continue
			}
			if optNoDelete {
				log.Printf("-- delete(dry): %s", file)
				continue
			}
			log.Printf("-- delete: %s", file)
			_ = os.Remove(file)
		}
	}); err != nil {
		return
	}
}
