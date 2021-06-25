package routers

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/polymorph-metadata/app/config"
	"github.com/polymorph-metadata/app/contracts"
	"github.com/polymorph-metadata/app/domain/metadata"
	"github.com/polymorph-metadata/app/interface/dlt/ethereum"
	log "github.com/sirupsen/logrus"
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

		tokenId := chi.URLParam(r, "tokenId")

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

		g := metadata.Genome(genomeInt.String())
		render.JSON(w, r, (&g).Metadata(tokenId, configService))

	}
}

func NewMetadataRouter(ethClient *ethereum.EthereumClient, address string, configService *config.ConfigService) http.Handler {
	r := chi.NewRouter()
	r.Get("/{tokenId}", HandleMetadataRequest(ethClient, address, configService))
	return r
}
