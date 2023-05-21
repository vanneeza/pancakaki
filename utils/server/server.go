package server

import (
	"log"
	"pancakaki/api"
	"pancakaki/config"
	"pancakaki/utils/pkg"
)

func Run() error {

	db, err := config.InitDB()
	if err != nil {
		return err
	}
	defer config.CloseDB(db)

	router := api.Run(db)
	serverAddress := pkg.GetEnv("SERVER_ADDRESS")
	log.Printf("Server is running on address %s\n", serverAddress)
	if err := router.Run(serverAddress); err != nil {
		return err
	}

	return nil
}
