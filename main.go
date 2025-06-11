package main

import (
	"GoJsonWalker/experiment"
	"GoJsonWalker/utils"
)

func main() {
	utils.GetLogger().Info("Starting...")

	experiment.ProfileSearchJsonFile()

	utils.GetLogger().Info("Done.")
}
