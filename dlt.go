package functions

import (
	"os"

	"github.com/polymorph-metadata/app/interface/dlt/ethereum"
	log "github.com/sirupsen/logrus"
)

func connectToEthereum() *ethereum.EthereumClient {

	nodeURL := os.Getenv("NODE_URL")

	client, err := ethereum.NewEthereumClient(nodeURL)

	if err != nil {
		log.Fatal(err)
	}

	log.Infoln("Successfully connected to ethereum client")

	return client
}
