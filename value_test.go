package log

import (
	"testing"
	"time"
)

func TestDefaultTimestamp(t *testing.T) {
	ts, ok := DefaultTimestamp().(string)
	if !ok {
		t.Error("DefaultTimestamp did not return a value of type string as expected.")
		return
	}

	_, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		t.Error(err)
		return
	}
}
