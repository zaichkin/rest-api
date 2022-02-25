package main

import (
	"market/internal/app"
)

const pathConfig = "configs/config.json"

func main() {
	app.Run(pathConfig)
}
