package main

import (
	"GoJsonWalker/monitor"
	"fmt"
	"math/rand"
	"time"
)

func loadData() []int {
	time.Sleep(50 * time.Millisecond)
	data := make([]int, 10000)
	for i := range data {
		data[i] = rand.Intn(1000)
	}
	return data
}

func someHeavyFunc() {
	time.Sleep(100 * time.Millisecond)
	sum := 0
	for i := 0; i < 50000; i++ {
		sum += i * i
	}
}

func processBatch(n int) {
	meter := monitor.NewThroughputMeter(fmt.Sprintf("processBatch_%d", n))
	for i := 0; i < n; i++ {
		_ = i * i
		meter.Inc()
	}
	meter.Report()
}

func main() {
	monitor.StartPprofServer("6060")
	monitor.SaveCPUProfile("logs/cpu.pprof", 10*time.Second)
	monitor.LogMemoryStats("start")

	monitor.MonitorFunc("someHeavyFunc", someHeavyFunc)

	data := monitor.MonitorFuncWithResult("loadData", loadData)

	monitor.MonitorFunc("processBatch", func() {
		processBatch(len(data))
	})

	monitor.LogMemoryStats("end")
	fmt.Println("Done. Check logs and http://localhost:6060/debug/pprof/")
	time.Sleep(5 * time.Second)
}
