package timequeue

import (
	"testing"
)

func TestPush(t *testing.T) {
	tq := NewTimeQueue()

	tq.Push(0.1, 123, "http://bing.com")

	if tq.Count() != 1 {
		t.Error("Count was not 1 after Push")
	}

	for tq.Count() > 0 {
	}

	if tq.Count() != 0 {
		t.Error("timequeue should be empty")
	}
}
