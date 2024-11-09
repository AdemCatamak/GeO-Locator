package main

import (
	"GeO-Locator/api"
	"GeO-Locator/internal/config"
	"fmt"
	"log"
)

func main() {
	log.Println("App Is Started")

	configManager := config.GetConfigManager()
	portNumber := configManager.GetString(config.PortKey)

	s := api.NewServer()

	err := s.Run(":" + portNumber)
	if err != nil {
		panic(fmt.Sprintf("Server cannot start: %s", err))
	}
}
