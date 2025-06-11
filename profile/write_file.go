package profile

import (
	"GoJsonWalker/utils"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func logToFile(path string, line string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		utils.GetLogger().Error("Failed to open file", "err", err, "path", path)
		return
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	_, err = f.WriteString(fmt.Sprintf("%s\n", line))
	if err != nil {
		utils.GetLogger().Error("Failed to write to file", "err", err, "path", path, "line", line)
		return
	}
}

func logToCSV(path string, record []string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		utils.GetLogger().Error("Failed to open csv file", "err", err, "path", path)
		return
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	w := csv.NewWriter(f)
	err = w.Write(record)
	if err != nil {
		utils.GetLogger().Error("Failed to write to file", "err", err, "path", path, "record", record)
		return
	}
	w.Flush()
}

func logToJSON(path string, data map[string]interface{}) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		utils.GetLogger().Error("Failed to open json file", "err", err, "path", path)
		return
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	enc := json.NewEncoder(f)
	err = enc.Encode(data)
	if err != nil {
		utils.GetLogger().Error("Failed to write to file", "err", err, "path", path, "data", data)
		return
	}
}
