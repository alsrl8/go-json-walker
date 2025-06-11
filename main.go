package main

import (
	"GoJsonWalker/parser"
	"GoJsonWalker/profile"
	"GoJsonWalker/utils"
)

func main() {
	utils.GetLogger().Info("Starting...")

	profile.LogMemoryStats("start")

	filename := "testdata/flat/random_flat.json"
	profile.MonitorFunc("ReadJsonFile", func() {
		_, err := parser.ReadJsonFile(filename)
		if err != nil {
			utils.GetLogger().Error("Failed to load file", "err", err, "filename", filename)
		}
	})

	profile.MonitorFunc("SearchJsonFile", func() {
		result, err := parser.SearchJsonFile(filename, "key_0")
		if err != nil {
			utils.GetLogger().Error("Failed to search file", "err", err, "filename", filename)
		}
		utils.GetLogger().Info("Result of `SearchJsonFile`", "result", result)
	})

	profile.LogMemoryStats("end")

	utils.GetLogger().Info("Done.")
}
