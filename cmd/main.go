package main

import (
	"github.com/joho/godotenv"
	"github.com/polymorph-metadata/app/interface/api"
	"github.com/polymorph-metadata/app/interface/api/routers"
	"os"
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

	apiPort := os.Getenv("API_PORT")
	a := api.NewAPI()

	nodeURL := os.Getenv("CONTRACT_ADDRESS")

	metadataRouter := routers.NewMetadataRouter(ethClient, nodeURL)

	a.AddRouter("/token", metadataRouter)

	a.Start(apiPort)
}
