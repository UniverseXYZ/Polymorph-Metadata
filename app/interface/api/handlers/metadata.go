package handlers

import (
	"math/big"
	"net/http"
	"strconv"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/render"
	"github.com/polymorph-metadata/app/config"
	"github.com/polymorph-metadata/app/contracts"
	"github.com/polymorph-metadata/app/domain/metadata"
	"github.com/polymorph-metadata/app/interface/dlt/ethereum"
	log "github.com/sirupsen/logrus"
	"github.com/polymorph-metadata/structs"
)

func HandleMetadataRequest(ethClient *ethereum.EthereumClient, address string, configService *config.ConfigService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		instance, err := contracts.NewPolymorph(common.HexToAddress(address), ethClient.Client)
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, err)
			log.Errorln(err)
			return
		}

		tokenId := r.URL.Query().Get("id")

		iTokenId, err := strconv.Atoi(tokenId)

		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, err)
			log.Errorln(err)
			return
		}

		genomeInt, err := instance.GeneOf(nil, big.NewInt(int64(iTokenId)))
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, err)
			log.Errorln(err)
			return
		}
		fmt.Println("genomeInt ", genomeInt)

		rarityResponse := structs.RarityServiceResponse{0,0}

		g := metadata.Genome(genomeInt.String())
		render.JSON(w, r, (&g).Metadata(tokenId, configService, rarityResponse))
	}
}
