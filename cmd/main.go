package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/polymorph-metadata/app/config"
	"github.com/polymorph-metadata/app/interface/api"
	"github.com/polymorph-metadata/app/interface/api/routers"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		godotenv.Load(args[0])
	} else {
		godotenv.Load()
	}

	setupLogger()

	ethClient := connectToEthereum()

	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("API_PORT")
	}
	a := api.NewAPI()

	nodeURL := os.Getenv("CONTRACT_ADDRESS")

	configService := config.NewConfigService("./config.json")

	metadataRouter := routers.NewMetadataRouter(ethClient, nodeURL, configService)

	a.AddRouter("/token", metadataRouter)

	a.Start(port)
}
