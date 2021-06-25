package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/joho/godotenv"
	"github.com/polymorph-metadata/app/config"
	"github.com/polymorph-metadata/app/interface/api/handlers"
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

	port := os.Getenv("API_PORT")
	if envport := os.Getenv("PORT"); envport == "" {
		port = os.Getenv("API_PORT")
	}

	contractAddress := os.Getenv("CONTRACT_ADDRESS")

	configService := config.NewConfigService("./config.json")

	funcframework.RegisterHTTPFunction("/token", handlers.HandleMetadataRequest(ethClient, contractAddress, configService))

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}

}
