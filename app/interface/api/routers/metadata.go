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
	"github.com/polymorph-metadata/app/interface/api/handlers"
	"github.com/polymorph-metadata/app/interface/dlt/ethereum"
	log "github.com/sirupsen/logrus"
)

func HandleMetadataRequest(ethClient *ethereum.EthereumClient, address string, configService *config.ConfigService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		instance, err := contracts.NewPolymorphicFacesRoot(common.HexToAddress(address), ethClient.Client)
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

		ownerOf, err := instance.OwnerOf(nil, big.NewInt(int64(iTokenId)))

		if ownerOf == common.HexToAddress("0x0000000000000000000000000000000000000000") {
			msg := "Query for non-existing token"
			render.Status(r, 404)
			render.JSON(w, r, msg)
			log.Errorln(msg)
			return
		}

		genomeInt, err := instance.GeneOf(nil, big.NewInt(int64(iTokenId)))
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, err)
			log.Errorln(err)
			return
		}

		rarity := handlers.GetRarityById(iTokenId)
		g := metadata.Genome(genomeInt.String())
		render.JSON(w, r, (&g).Metadata(tokenId, configService, rarity))
	}
}

func NewMetadataRouter(ethClient *ethereum.EthereumClient, address string, configService *config.ConfigService) http.Handler {
	r := chi.NewRouter()
	r.Get("/{tokenId}", HandleMetadataRequest(ethClient, address, configService))
	return r
}
