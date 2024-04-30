package main

import (
	"banking/app"
	"banking/logger"
)

func main() {
	// log.Println("Starting the application...")
	logger.Info("Starting the application...")
	app.Start()
}
