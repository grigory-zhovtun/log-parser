package parser

import (
	"fmt"
	"regexp"
	"strconv"
)

// LogEntry представляет одну запись лога Nginx
type LogEntry struct {
	IP         string
	Timestamp  string
	Method     string
	Path       string
	StatusCode int
}

func ParseLogLine(line string) (LogEntry, error) {
	re := regexp.MustCompile(`^(\S+) - \[(.+?)\] "(\S+) (\S+) .*?" (\d+)`)

	matches := re.FindStringSubmatch(line)

	if len(matches) < 6 {
		fmt.Println("Ошибка: формат строки не совпадает")
		return
	}

	statusCode, _ := strconv.Atoi(matches[5])

	entry := LogEntry{
		IP:         matches[1],
		Timestamp:  matches[2],
		Method:     matches[3],
		Path:       matches[4],
		StatusCode: statusCode,
	}

	return entry, nil
}
