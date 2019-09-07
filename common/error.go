package common

import (
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
)

var instance = log.New()

// TODO
func LogAndErrorf(format string, a ...interface{}) error {
	instance.Formatter = &log.TextFormatter{
		CallerPrettyfier: func(r *runtime.Frame) (string, string) {
			return "", ""
		},
	}
	instance.Errorf(format, a)
	return fmt.Errorf("", a...)
}
