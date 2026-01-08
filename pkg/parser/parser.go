package parser

// LogEntry представляет одну запись лога Nginx
type LogEntry struct {
    IP         string
    Timestamp  string
    Method     string
    Path       string
    StatusCode int
}