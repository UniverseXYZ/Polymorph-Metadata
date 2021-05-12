package functions

import (
	"github.com/polymorph-metadata/handlers"
	"net/http"
)

func Double(w http.ResponseWriter, r *http.Request) {
	handlers.DoubleNumber(w, r)
}

func TokenMetadata(w http.ResponseWriter, r *http.Request) {
	handlers.TokenMetadata(w, r)
}
