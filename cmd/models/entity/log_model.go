package models

import "time"

const (
	danger  = "\033[0;31m"
	warning = "\033[0;33m"
	info    = "\033[0;36m"
	success = "\033[0;32m"
	date    = "\033[0;34m"
	end     = "\033[0m"
)

// time format for log messages
// yyyy-mm-dd hh:mm:ss
var now = time.Now().Format("2006-01-02 15:04:05")

type LogModel struct {
	Message string
	mode    string
}

func (l *LogModel) Print(message string, mode string) {
	l.Message = message
	l.mode = mode

	switch l.mode {
	case "success":
		l.Message = date + now + " " + success + l.Message + end
	case "danger":
		l.Message = date + now + " " + danger + l.Message + end
	case "warning":
		l.Message = date + now + " " + warning + l.Message + end
	case "info":
		l.Message = date + now + " " + info + l.Message + end
	}
	println(l.Message)
}
