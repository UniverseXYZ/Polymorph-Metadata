package routers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/polymorph-metadata/app/contracts"
	"github.com/polymorph-metadata/app/domain/metadata"
	"github.com/polymorph-metadata/app/interface/dlt/ethereum"
	log "github.com/sirupsen/logrus"
	"math/big"
	"net/http"
	"strconv"
)

func handleMetadataRequest(ethClient *ethereum.EthereumClient, address string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		instance, err := contracts.NewPolymorph(common.HexToAddress(address), ethClient.Client)
		if err != nil {
			render.Status(r, 500)
			log.Fatal(err)
		}

		tokenId := chi.URLParam(r, "tokenId")

		iTokenId, err := strconv.Atoi(tokenId)
		if err != nil {
			render.Status(r, 500)
			log.Fatal(err)
		}

		genomeInt, err := instance.GeneOf(nil, big.NewInt(int64(iTokenId)))
		if err != nil {
			render.Status(r, 500)
			log.Fatal(err)
		}

		g := metadata.Genome(genomeInt.String())
		render.JSON(w, r, (&g).Metadata(tokenId))

	}
}

func NewMetadataRouter(ethClient *ethereum.EthereumClient, address string) http.Handler {
	r := chi.NewRouter()
	r.Get("/{tokenId}", handleMetadataRequest(ethClient, address))
	return r
}
