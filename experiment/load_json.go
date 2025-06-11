package experiment

import (
	"GoJsonWalker/parser"
	"GoJsonWalker/profile"
	"GoJsonWalker/utils"
)

func ProfileSearchJsonFile() {
	profile.LogMemoryStats("start profile json load")

	filename := "testdata/flat/random_flat.json"
	profile.MonitorFunc("SearchJsonFile", func() {
		result, err := parser.SearchJsonFile(filename, "key_0")
		if err != nil {
			utils.GetLogger().Error("Failed to search file", "err", err, "filename", filename)
		}
		utils.GetLogger().Info("Result of `SearchJsonFile`", "result", result)
	})

	profile.MonitorFunc("SearchJsonFile", func() {
		result, err := parser.SearchJsonFile(filename, "key_0")
		if err != nil {
			utils.GetLogger().Error("Failed to search file", "err", err, "filename", filename)
		}
		utils.GetLogger().Info("Result of `SearchJsonFile`", "result", result)
	})

	profile.MonitorFunc("SearchJsonFile", func() {
		result, err := parser.SearchJsonFile(filename, "key_0")
		if err != nil {
			utils.GetLogger().Error("Failed to search file", "err", err, "filename", filename)
		}
		utils.GetLogger().Info("Result of `SearchJsonFile`", "result", result)
	})

	profile.LogMemoryStats("end profile json load")
}
