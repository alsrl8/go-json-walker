package profile

import (
	"fmt"
	"time"
)

// Timer is used to measure elapsed time.
type Timer struct {
	start time.Time
	name  string
}

func NewTimer(name string) *Timer {
	return &Timer{
		start: time.Now(),
		name:  name,
	}
}

func (t *Timer) Stop() {
	elapsed := time.Since(t.start)
	logLine := fmt.Sprintf("[TIMER] %s took %s", t.name, elapsed)
	fmt.Println(logLine)
	logToFile("logs/timers.log", logLine)
}
