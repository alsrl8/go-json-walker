package experiment

import (
	"GoJsonWalker/parser"
	"GoJsonWalker/profile"
	"GoJsonWalker/utils"
)

func ProfileStreamJsonFile() {
	profile.LogMemoryStats("start profile json stream")

	filename := "testdata/flat/random_flat.json"
	profile.MonitorFunc("StreamJsonFile", func() {
		result, err := parser.StreamJsonFile(filename, "key_1")
		if err != nil {
			utils.GetLogger().Error("Failed to stream file", "err", err, "filename", filename)
		}
		utils.GetLogger().Info("Result of `StreamJsonFile`", "result", result)
	})

	profile.MonitorFunc("StreamJsonFile", func() {
		result, err := parser.StreamJsonFile(filename, "key_5475")
		if err != nil {
			utils.GetLogger().Error("Failed to stream file", "err", err, "filename", filename)
		}
		utils.GetLogger().Info("Result of `StreamJsonFile`", "result", result)
	})

	profile.MonitorFunc("StreamJsonFile", func() {
		result, err := parser.StreamJsonFile(filename, "key_9570")
		if err != nil {
			utils.GetLogger().Error("Failed to stream file", "err", err, "filename", filename)
		}
		utils.GetLogger().Info("Result of `StreamJsonFile`", "result", result)
	})

	profile.LogMemoryStats("end profile json stream")
}
