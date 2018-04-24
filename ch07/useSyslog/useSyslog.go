package main

import (
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil {
		log.Fatal(err)
	}
	sysLog.Crit("Crit: Logging in Go!")
	sysLog, err = syslog.New(syslog.LOG_ALERT|syslog.LOG_LOCAL7, "Some program")
	if err != nil {
		log.Fatal(err)
	}
	sysLog.Emerg("Emerg: Logging in Go!")
}
