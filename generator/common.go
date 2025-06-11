package generator

import (
	"GoJsonWalker/utils"
	"encoding/json"
	"os"
	"path/filepath"
)

func getTestDataDir() string {
	return "testdata/"
}

func writeFile(dirname string, filename string, content interface{}) {
	testDataDir := getTestDataDir()
	outputDir := filepath.Join(testDataDir, dirname)
	outputFile := filepath.Join(outputDir, filename)

	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		utils.GetLogger().Error("Failed to create directory", "err", err, "dir", outputDir)
		return
	}

	file, err := os.Create(outputFile)
	if err != nil {
		utils.GetLogger().Error("Failed to create file", "err", err, "file", outputFile)
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(content)
	if err != nil {
		utils.GetLogger().Error("Failed to write to file", "err", err, "file", outputFile, "content", content)
		return
	}

	utils.GetLogger().Info("JSON file generated", "file", outputFile)
}
