package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// LogEntry описывает структуру одной строки лога
type LogEntry struct {
	IP         string
	Timestamp  string
	Method     string
	Path       string
	StatusCode int
}

var logRe = regexp.MustCompile(`^(\S+) - \[(.+?)\] "(\S+) (\S+) .*?" (\d+)`)

func ParseLogLine(line string) (LogEntry, error) {
	matches := logRe.FindStringSubmatch(line)

	if len(matches) < 6 {
		return LogEntry{}, errors.New("неверный формат строки лога")
	}

	statusCode, err := strconv.Atoi(matches[5])
	if err != nil {
		return LogEntry{}, fmt.Errorf("некорректный status code: %w", err)
	}

	return LogEntry{
		IP:         matches[1],
		Timestamp:  matches[2],
		Method:     matches[3],
		Path:       matches[4],
		StatusCode: statusCode,
	}, nil
}
