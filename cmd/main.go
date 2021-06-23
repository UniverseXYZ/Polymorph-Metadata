package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/polymorph-metadata/app/config"
	"github.com/polymorph-metadata/app/interface/api"
	"github.com/polymorph-metadata/app/interface/api/processor"
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

	apiPort := os.Getenv("API_PORT")
	queueSize := os.Getenv("QUEUE_SIZE")

	a := api.NewAPI()

	contractAddress := os.Getenv("CONTRACT_ADDRESS")

	configService := config.NewConfigService("./config.json")
	iSize, err := strconv.Atoi(queueSize)
	if err != nil {
		log.Fatalln(err)
	}

	p := processor.NewProcessor(iSize, ethClient, contractAddress, configService)

	metadataRouter := routers.NewMetadataRouter(p)

	a.AddRouter("/token", metadataRouter)

	p.Start()

	a.Start(apiPort)
}
