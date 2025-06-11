package main

import (
	"GoJsonWalker/parser"
	"GoJsonWalker/profile"
	"GoJsonWalker/utils"
)

func main() {
	utils.GetLogger().Info("Starting...")

	profile.LogMemoryStats("start")

	filename := "./testdata/large_random_json.json"
	profile.MonitorFunc("ReadJsonFile", func() {
		_, err := parser.ReadJsonFile(filename)
		if err != nil {
			utils.GetLogger().Error("Failed to load file", "err", err, "filename", filename)
		}
	})

	profile.LogMemoryStats("end")

	utils.GetLogger().Info("Done.")
}
