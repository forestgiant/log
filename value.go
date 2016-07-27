package log

import "time"

// A ValueGenerator is a function that generates a dynamic value to be evaluated at the time of a log event
type ValueGenerator func() interface{}

// DefaultTimestamp is a ValueGenerator that returns the current time in RFC3339 format
var DefaultTimestamp ValueGenerator = func() interface{} {
	return time.Now().Format(time.RFC3339)
}
