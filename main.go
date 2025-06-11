package main

import (
	"GoJsonWalker/experiment"
	"GoJsonWalker/utils"
)

func main() {
	utils.GetLogger().Info("Starting...")

	experiment.ProfileSearchJsonFile()
	experiment.ProfileStreamJsonFile()

	utils.GetLogger().Info("Done.")
}
