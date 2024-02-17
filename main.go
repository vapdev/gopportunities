package main

import (
	"github.com/vapdev/gopportunities.git/config"
	"github.com/vapdev/gopportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	// Initialize router
	router.Init()
}
