package monitor

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"
	"runtime"
	pprof2 "runtime/pprof"
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

// MemoryStats logs memory usage of the current process.
func LogMemoryStats(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	logLine := fmt.Sprintf("[MEM] %s - Alloc = %v KB, TotalAlloc = %v KB, Sys = %v KB, NumGC = %v",
		label, m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
	fmt.Println(logLine)
	logToFile("logs/memory.log", logLine)
}

// ThroughputMeter tracks processed items per second.
type ThroughputMeter struct {
	start time.Time
	count int
	label string
}

func NewThroughputMeter(label string) *ThroughputMeter {
	return &ThroughputMeter{
		start: time.Now(),
		label: label,
	}
}

func (tm *ThroughputMeter) Inc() {
	tm.count++
}

func (tm *ThroughputMeter) Report() {
	duration := time.Since(tm.start).Seconds()
	if duration == 0 {
		duration = 0.001
	}
	rate := float64(tm.count) / duration
	logLine := fmt.Sprintf("[THROUGHPUT] %s - Processed %d items in %.2f seconds (%.2f items/sec)",
		tm.label, tm.count, duration, rate)
	fmt.Println(logLine)
	logToFile("logs/throughput.log", logLine)
	logToCSV("logs/throughput.csv", []string{tm.label, fmt.Sprintf("%d", tm.count), fmt.Sprintf("%.2f", duration), fmt.Sprintf("%.2f", rate)})
	logToJSON("logs/throughput.json", map[string]interface{}{
		"label":     tm.label,
		"count":     tm.count,
		"duration":  duration,
		"rate":      rate,
		"timestamp": time.Now(),
	})
}

// StartPprofServer launches a pprof debug server on the given port.
func StartPprofServer(port string) {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "pprof server running")
		})
		mux.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		mux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		mux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		mux.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		mux.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
		fmt.Printf("[PPROF] Starting pprof server at http://localhost:%s/debug/pprof/\n", port)
		http.ListenAndServe(":"+port, mux)
	}()
}

// SaveCPUProfile writes a CPU profile to the given file.
func SaveCPUProfile(filename string, duration time.Duration) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Failed to create CPU profile file: %v\n", err)
		return
	}
	_ = pprof2.StartCPUProfile(f)
	time.Sleep(duration)
	pprof2.StopCPUProfile()
	fmt.Printf("[PROFILE] CPU profile saved to %s\n", filename)
}

// Helpers
func logToFile(path string, line string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("%s\n", line))
}

func logToCSV(path string, record []string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.Write(record)
	w.Flush()
}

func logToJSON(path string, data map[string]interface{}) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.Encode(data)
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
