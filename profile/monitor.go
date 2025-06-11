package profile

import (
	"fmt"
	"runtime"
)

func LogMemoryStats(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	logLine := fmt.Sprintf("[MEM] %s - Alloc = %v KiB, TotalAlloc = %v KiB, Sys = %v KiB, NumGC = %v",
		label, m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
	fmt.Println(logLine)
	logToFile("logs/memory.log", logLine)
}

func MonitorFunc(name string, f func()) {
	timer := NewTimer(name)
	LogMemoryStats(name + " (start)")
	f()
	LogMemoryStats(name + " (end)")
	timer.Stop()
}

func MonitorFuncWithResult[T any](name string, f func() T) T {
	timer := NewTimer(name)
	LogMemoryStats(name + " (start)")
	result := f()
	LogMemoryStats(name + " (end)")
	timer.Stop()
	return result
}
