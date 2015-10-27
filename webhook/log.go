package webhook

import (
	"log"
	"os"
)

var (
	logger = log.New(os.Stderr, "", log.LstdFlags)
)

func SetLogger(l *log.Logger) {
	logger = l
}

func logErr(err error) {
	if logger == nil {
		return
	}
	logger.Printf("webhoook error: %s", err)
	return
}
