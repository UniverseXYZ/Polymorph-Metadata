package routers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/polymorph-metadata/app/domain/metadata"
	"github.com/polymorph-metadata/app/interface/api/processor"
)

func handleMetadataRequest(p *processor.Processor) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenId := chi.URLParam(r, "tokenId")

		responseC := make(chan *metadata.Metadata)
		errorC := make(chan error)

		go p.HandleProcessRequest(tokenId, responseC, errorC)
		select {
		case m := <-responseC:
			render.JSON(w, r, m)
		case err := <-errorC:
			render.Status(r, 500)
			render.JSON(w, r, err)
			log.Fatal(err)
		}

	}
}

func NewMetadataRouter(p *processor.Processor) http.Handler {
	r := chi.NewRouter()
	r.Get("/{tokenId}", handleMetadataRequest(p))
	return r
}
